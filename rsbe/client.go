package rsbe

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/cookiejar"
	"net/url"
	"sync"
	"time"
)

const (
	AuthTypeBasic  string = "basic"
	AuthTypeCookie string = "cookie"
)

type ErrMsg struct {
	Error string `json:"error"`
}

type Config struct {
	BaseURL   string
	User      string
	Password  string
	AuthType  string
	LoginPath string
	// AutoRefreshOn401 controls automatic cookie re-authentication on 401.
	// nil means enabled by default for cookie auth.
	AutoRefreshOn401 *bool
	// MaxAuthRetries is the number of retry attempts after re-authentication.
	// Values <= 0 default to 1.
	MaxAuthRetries int
	// AuthRefreshCooldownSeconds skips redundant refresh if a successful refresh
	// happened recently. Values <= 0 disable cooldown checks.
	AuthRefreshCooldownSeconds int
}

var conf *Config
var client *http.Client
var authRefreshMu sync.Mutex
var authRefreshCond *sync.Cond
var authRefreshInFlight bool
var lastAuthRefreshAt time.Time
var lastAuthRefreshErr error

func ConfigureClient(c *Config) error {
	conf = c
	resetAuthRefreshState()
	return nil
}

// Login performs authentication for cookie-based auth and stores session cookies
// in the client's cookie jar.
func Login() error {
	if conf == nil {
		return fmt.Errorf("client is not configured")
	}

	if conf.AuthType != AuthTypeCookie {
		return nil
	}

	if err := login(); err != nil {
		return fmt.Errorf("cookie authentication failed: %w", err)
	}

	return nil
}

// Logout clears cookie-based authentication state from the client's cookie jar.
func Logout() error {
	if conf == nil {
		return fmt.Errorf("client is not configured")
	}

	if conf.AuthType != AuthTypeCookie {
		return nil
	}

	if err := logout(); err != nil {
		return fmt.Errorf("cookie logout failed: %w", err)
	}

	return nil
}

func logout() error {

	if conf.LoginPath == "" {
		return fmt.Errorf("LoginPath is required for cookie authentication")
	}

	logoutURL := fmt.Sprintf("%s%s", conf.BaseURL, conf.LoginPath)
	req, err := http.NewRequest("DELETE", logoutURL, nil)
	if err != nil {
		return err
	}

	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK && resp.StatusCode != http.StatusAccepted && resp.StatusCode != http.StatusNoContent {
		respBody, readErr := io.ReadAll(resp.Body)
		if readErr != nil {
			return fmt.Errorf("logout failed: %d (unable to read response body: %v)", resp.StatusCode, readErr)
		}
		var eMsg ErrMsg
		if err := json.Unmarshal(respBody, &eMsg); err != nil {
			return fmt.Errorf("logout failed: %d (response: %s)", resp.StatusCode, string(respBody))
		}
		return fmt.Errorf("logout failed: %d ; %v", resp.StatusCode, eMsg.Error)
	}

	baseURL, err := url.Parse(conf.BaseURL)
	if err != nil {
		return fmt.Errorf("invalid BaseURL: %w", err)
	}

	// Recreate the cookie jar to ensure all session cookies are cleared.
	jar, err := cookiejar.New(nil)
	if err != nil {
		return fmt.Errorf("failed to reset cookie jar: %w", err)
	}
	client.Jar = jar
	client.Jar.SetCookies(baseURL, []*http.Cookie{})

	return nil
}

func init() {
	// cookiejar.New with nil options never returns an error according to the API docs
	jar, err := cookiejar.New(nil)
	if err != nil {
		// This should never happen with nil options, but handle defensively
		panic(fmt.Sprintf("failed to create cookie jar: %v", err))
	}
	client = &http.Client{
		Jar: jar,
	}
	authRefreshCond = sync.NewCond(&authRefreshMu)
}

func resetAuthRefreshState() {
	authRefreshMu.Lock()
	defer authRefreshMu.Unlock()
	authRefreshInFlight = false
	lastAuthRefreshAt = time.Time{}
	lastAuthRefreshErr = nil
}

func shouldAutoRefreshOn401() bool {
	if conf == nil || conf.AuthType != AuthTypeCookie {
		return false
	}
	if conf.AutoRefreshOn401 == nil {
		return true
	}
	return *conf.AutoRefreshOn401
}

func getMaxAuthRetries() int {
	if conf == nil || conf.MaxAuthRetries <= 0 {
		return 1
	}
	return conf.MaxAuthRetries
}

func getAuthRefreshCooldown() time.Duration {
	if conf == nil || conf.AuthRefreshCooldownSeconds <= 0 {
		return 0
	}
	return time.Duration(conf.AuthRefreshCooldownSeconds) * time.Second
}

func shouldRetryUnauthorized(statusCode int, attempt int) bool {
	if statusCode != http.StatusUnauthorized {
		return false
	}
	if !shouldAutoRefreshOn401() {
		return false
	}
	return attempt < getMaxAuthRetries()
}

func refreshCookieAuthSession() error {
	if conf == nil {
		return fmt.Errorf("client is not configured")
	}

	authRefreshMu.Lock()
	for authRefreshInFlight {
		authRefreshCond.Wait()
	}

	cooldown := getAuthRefreshCooldown()
	if cooldown > 0 && !lastAuthRefreshAt.IsZero() && time.Since(lastAuthRefreshAt) < cooldown && lastAuthRefreshErr == nil {
		authRefreshMu.Unlock()
		return nil
	}

	authRefreshInFlight = true
	authRefreshMu.Unlock()

	err := login()

	authRefreshMu.Lock()
	if err == nil {
		lastAuthRefreshAt = time.Now()
	}
	lastAuthRefreshErr = err
	authRefreshInFlight = false
	authRefreshCond.Broadcast()
	authRefreshMu.Unlock()

	if err != nil {
		return fmt.Errorf("cookie authentication failed: %w", err)
	}

	return nil
}

func executeRequest(method string, path string, data []byte, contentType string) (resp *http.Response, err error) {
	if conf == nil {
		return nil, fmt.Errorf("client is not configured")
	}

	requestURL := fmt.Sprintf("%s%s", conf.BaseURL, path)
	maxRetries := getMaxAuthRetries()

	for attempt := 0; attempt <= maxRetries; attempt++ {
		var bodyReader io.Reader
		if len(data) > 0 {
			bodyReader = bytes.NewReader(data)
		}

		req, reqErr := http.NewRequest(method, requestURL, bodyReader)
		if reqErr != nil {
			return nil, reqErr
		}

		if contentType != "" {
			req.Header.Set("Content-Type", contentType)
		}

		setAuth(req)

		resp, err = client.Do(req)
		if err != nil {
			return resp, err
		}

		if shouldRetryUnauthorized(resp.StatusCode, attempt) {
			_ = resp.Body.Close()
			if refreshErr := refreshCookieAuthSession(); refreshErr != nil {
				return resp, fmt.Errorf("bad response: %d ; unauthorized; auth refresh failed: %v", http.StatusUnauthorized, refreshErr)
			}
			continue
		}

		return resp, nil
	}

	return resp, err
}

func badResponseError(resp *http.Response) error {
	body, _ := io.ReadAll(resp.Body)

	var eMsg ErrMsg
	_ = json.Unmarshal(body, &eMsg)
	return fmt.Errorf("bad response: %d ; %v", resp.StatusCode, eMsg.Error)
}

// setAuth applies the appropriate authentication method to the HTTP request
// based on the configured AuthType. For basic auth, sets HTTP Basic Auth headers.
// For cookie auth, cookies are automatically handled by the HTTP client's cookie jar.
func setAuth(req *http.Request) {
	// Default to basic auth if AuthType is not set or is explicitly basic
	if conf.AuthType == AuthTypeBasic || conf.AuthType == "" {
		req.SetBasicAuth(conf.User, conf.Password)
	}
	// Cookie auth: cookies are automatically handled by client.Jar
}

func login() error {
	if conf.LoginPath == "" {
		return fmt.Errorf("LoginPath is required for cookie authentication")
	}

	url := fmt.Sprintf("%s%s", conf.BaseURL, conf.LoginPath)

	body := struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}{
		Email:    conf.User,
		Password: conf.Password,
	}

	data, err := json.Marshal(body)
	if err != nil {
		return err
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(data))
	if err != nil {
		return err
	}

	req.Header.Set("Content-Type", "application/json")

	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK && resp.StatusCode != http.StatusCreated {
		respBody, err := io.ReadAll(resp.Body)
		if err != nil {
			return fmt.Errorf("login failed: %d (unable to read response body: %v)", resp.StatusCode, err)
		}
		var eMsg ErrMsg
		if err := json.Unmarshal(respBody, &eMsg); err != nil {
			return fmt.Errorf("login failed: %d (response: %s)", resp.StatusCode, string(respBody))
		}
		return fmt.Errorf("login failed: %d ; %v", resp.StatusCode, eMsg.Error)
	}

	return nil
}

func Get(path string) (resp *http.Response, err error) {
	resp, err = executeRequest("GET", path, nil, "")
	if err != nil {
		return resp, err
	}

	if resp.StatusCode != 200 {
		return resp, badResponseError(resp)
	}

	return resp, nil
}

func GetBody(path string) (body []byte, err error) {

	resp, err := Get(path)
	if err != nil {
		return body, err
	}
	defer resp.Body.Close()

	body, err = io.ReadAll(resp.Body)
	if err != nil {
		return body, err
	}

	return body, nil
}

func Post(path string, data []byte) (resp *http.Response, err error) {
	resp, err = executeRequest("POST", path, data, "application/json")
	if err != nil {
		return resp, err
	}

	if resp.StatusCode != 201 {
		return resp, badResponseError(resp)
	}

	return resp, nil
}

func PostReturnBody(path string, data []byte) (body []byte, err error) {
	resp, err := Post(path, data)
	if err != nil {
		return body, err
	}
	defer resp.Body.Close()

	body, err = io.ReadAll(resp.Body)
	if err != nil {
		return body, err
	}

	return body, nil
}

func Put(path string, data []byte) (err error) {
	resp, err := executeRequest("PUT", path, data, "application/json")
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 204 {
		return fmt.Errorf("unexpected status code: %s", resp.Status)
	}

	return nil
}

func Delete(path string) (err error) {
	resp, err := executeRequest("DELETE", path, nil, "")
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 204 {
		return badResponseError(resp)
	}

	return nil
}

func Purge(path string) (err error) {
	resp, err := executeRequest("DELETE", fmt.Sprintf("%s?purge=true", path), nil, "")
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 204 {
		return badResponseError(resp)
	}

	return nil
}
