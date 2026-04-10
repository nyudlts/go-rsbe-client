package rsbe

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/cookiejar"
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
}

var conf *Config
var client *http.Client

func ConfigureClient(c *Config) error {
	conf = c

	// If cookie auth, perform login to get session cookie
	if conf.AuthType == AuthTypeCookie {
		if err := login(); err != nil {
			return fmt.Errorf("cookie authentication failed: %w", err)
		}
	}

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
	url := fmt.Sprintf("%s%s", conf.BaseURL, path)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	setAuth(req)

	resp, err = client.Do(req)

	if err != nil {
		return resp, err
	}

	if resp.StatusCode != 200 {
		body, _ := io.ReadAll(resp.Body)
		var eMsg ErrMsg
		_ = json.Unmarshal(body, &eMsg)
		return resp, fmt.Errorf("bad response: %d ; %v", resp.StatusCode, eMsg.Error)
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
	url := fmt.Sprintf("%s%s", conf.BaseURL, path)

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(data))
	if err != nil {
		return resp, err
	}

	req.Header.Set("Content-Type", "application/json")
	setAuth(req)

	resp, err = client.Do(req)
	if err != nil {
		return resp, err
	}

	if resp.StatusCode != 201 {
		body, _ := io.ReadAll(resp.Body)

		var eMsg ErrMsg
		_ = json.Unmarshal(body, &eMsg)
		return resp, fmt.Errorf("bad response: %d ; %v", resp.StatusCode, eMsg.Error)
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
	url := fmt.Sprintf("%s%s", conf.BaseURL, path)

	req, err := http.NewRequest("PUT", url, bytes.NewBuffer(data))
	if err != nil {
		return err
	}

	req.Header.Set("Content-Type", "application/json")
	setAuth(req)

	resp, err := client.Do(req)
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
	url := fmt.Sprintf("%s%s", conf.BaseURL, path)

	req, err := http.NewRequest("DELETE", url, nil)
	if err != nil {
		return err
	}

	setAuth(req)

	resp, err := client.Do(req)
	if err != nil {
		return err
	}

	if resp.StatusCode != 204 {
		body, _ := io.ReadAll(resp.Body)

		var eMsg ErrMsg
		_ = json.Unmarshal(body, &eMsg)
		return fmt.Errorf("bad response: %d ; %v", resp.StatusCode, eMsg.Error)
	}

	return nil
}

func Purge(path string) (err error) {
	url := fmt.Sprintf("%s%s?purge=true", conf.BaseURL, path)

	req, err := http.NewRequest("DELETE", url, nil)
	if err != nil {
		return err
	}

	setAuth(req)

	resp, err := client.Do(req)
	if err != nil {
		return err
	}

	if resp.StatusCode != 204 {
		body, _ := io.ReadAll(resp.Body)

		var eMsg ErrMsg
		_ = json.Unmarshal(body, &eMsg)
		return fmt.Errorf("bad response: %d ; %v", resp.StatusCode, eMsg.Error)
	}

	return nil
}
