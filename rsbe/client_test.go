package rsbe

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"path"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestClientGet(t *testing.T) {

	setupLocalhostClient()
	t.Run("confirm OK status response", func(t *testing.T) {
		want := 200
		resp, err := Get("/api/v0/partners/e6517775-6277-4e25-9373-ee7738e820b5")
		if err != nil {
			t.Errorf("Unexpected error: %s", err)
		}

		got := resp.StatusCode
		if want != got {
			t.Errorf("Mismatch: want: \"%v\", got: \"%v\"", want, got)
		}
	})

	t.Run("confirm bad request response", func(t *testing.T) {
		want := 400
		resp, err := Get("/api/v0/partners/x")
		if err == nil {
			t.Errorf("err should NOT be nil: %v", err)
		}

		got := resp.StatusCode
		if want != got {
			t.Errorf("Mismatch: want: \"%v\", got: \"%v\"", want, got)
		}
	})

}

func TestClientGetBody(t *testing.T) {

	setupLocalhostClient()
	t.Run("confirm OK status response", func(t *testing.T) {
		body, err := GetBody("/api/v0/partners/e6517775-6277-4e25-9373-ee7738e820b5")
		if err != nil {
			t.Errorf("Unexpected error: %s", err)
		}

		if len(body) == 0 {
			t.Errorf("Body should not be empty")
		}
	})

	t.Run("confirm bad request response", func(t *testing.T) {
		body, err := GetBody("/api/v0/partners/x")
		if err == nil {
			t.Errorf("err should NOT be nil: %v", err)
		}

		if len(body) != 0 {
			t.Errorf("Body should be empty")
		}
	})

}

func TestClientPost(t *testing.T) {

	data := "{\"code\":\"canteloupe\",\"name\":\"Can elope\",\"rel_path\":\"content/canteloupe\"}"
	ppath := "/api/v0/partners"

	setupLocalhostClient()
	t.Run("confirm OK status response on successful POST", func(t *testing.T) {

		want := 201
		resp, err := Post(ppath, []byte(data))
		if err != nil {
			t.Errorf("Unexpected error: %s", err)
		}

		got := resp.StatusCode
		if want != got {
			t.Errorf("Mismatch: want: \"%v\", got: \"%v\"", want, got)
		}

		id := path.Base(resp.Header.Get("Location"))
		if id == "." {
			t.Errorf("Unable to find created partner to delete.")
		}
		err = PartnerDelete(id)
		if err != nil {
			t.Errorf("Unexpected error: %s", err)
		}
	})

	t.Run("confirm bad request response", func(t *testing.T) {

		_, err := Post("/api/v0/partners", []byte("{}"))
		if err == nil {
			t.Errorf("err should NOT be nil: %v", err)
		}
	})

}

func TestClientPostReturnBody(t *testing.T) {

	data := "{\"code\":\"canteloupe\",\"name\":\"Can elope\",\"rel_path\":\"content/canteloupe\"}"
	ppath := "/api/v0/partners"

	setupLocalhostClient()
	t.Run("confirm no error and non-empty response body on successful POST", func(t *testing.T) {

		body, err := PostReturnBody(ppath, []byte(data))
		if err != nil {
			t.Errorf("Unexpected error: %s", err)
		}

		if err != nil {
			t.Errorf("err should be nil: %s", err)
		}

		if len(body) == 0 {
			t.Errorf("Body should not be empty")
		}

		var p PartnerEntry
		err = json.Unmarshal(body, &p)
		if err != nil {
			t.Errorf("Error parsing body.")
		}

		err = PartnerDelete(p.ID)
		if err != nil {
			t.Errorf("Unexpected error: %s", err)
		}
	})

	t.Run("confirm bad request response", func(t *testing.T) {

		_, err := PostReturnBody("/api/v0/partners", []byte("{}"))
		if err == nil {
			t.Errorf("err should NOT be nil: %v", err)
		}
	})

}

func TestClientPut(t *testing.T) {

	data := "{\"code\":\"canteloupe\",\"name\":\"Can elope\",\"rel_path\":\"content/canteloupe\"}"
	ppath := "/api/v0/partners"

	setupLocalhostClient()
	t.Run("confirm OK status response on successful POST", func(t *testing.T) {

		// create Partner to update:
		resp, err := Post(ppath, []byte(data))
		if err != nil {
			t.Errorf("Unexpected error: %s", err)
		}

		id := path.Base(resp.Header.Get("Location"))
		if id == "." {
			t.Errorf("Unable to find created partner to delete.")
		}

		data := "{\"code\":\"bananananana\",\"name\":\"Can elope\",\"rel_path\":\"content/canteloupe\"}"
		err = Put(ppath+"/"+id, []byte(data))
		if err != nil {
			t.Errorf("Unexpected error: %s", err)
		}

		err = PartnerDelete(id)
		if err != nil {
			t.Errorf("Unexpected error: %s", err)
		}
	})

	t.Run("confirm bad request response", func(t *testing.T) {

		err := Put("/api/v0/partners/adflasdfj", []byte("{}"))
		if err == nil {
			t.Errorf("err should NOT be nil: %v", err)
		}
	})

}

func TestClientDelete(t *testing.T) {

	data := "{\"code\":\"canteloupe\",\"name\":\"Can elope\",\"rel_path\":\"content/canteloupe\"}"
	ppath := "/api/v0/partners"

	setupLocalhostClient()
	t.Run("confirm OK status response on successful DELETE", func(t *testing.T) {

		// create Partner to delete
		resp, err := Post(ppath, []byte(data))
		if err != nil {
			t.Errorf("Unexpected error: %s", err)
		}

		id := path.Base(resp.Header.Get("Location"))
		if id == "." {
			t.Errorf("Unable to find created partner to delete.")
		}

		err = Delete(ppath + "/" + id)
		if err != nil {
			t.Errorf("Unexpected error: %s", err)
		}
	})

	t.Run("confirm bad request response", func(t *testing.T) {

		err := Delete("/api/v0/partners/adflasdfj")
		if err == nil {
			t.Errorf("err should NOT be nil: %v", err)
		}
	})

}

func TestCookieAuth(t *testing.T) {
	t.Run("test cookie auth config", func(t *testing.T) {
		c, err := GetConfig("cookie")
		assert.NoError(t, err, "Failed to get cookie config")

		assert.NoError(t, ConfigureClient(c), "ConfigureClient should only return an error if the server is down")
		assert.Equal(t, conf.AuthType, AuthTypeCookie, "AuthType should be set to cookie after ConfigureClient")
		assert.Equal(t, conf.User, c.User, "Configured User should match config User")
		assert.Equal(t, conf.Password, c.Password, "Configured Password should match config Password")
		assert.NotEmpty(t, conf.LoginPath, "Expected LoginPath to be set")
	})

	t.Run("test basic auth config", func(t *testing.T) {
		c, err := GetConfig("basic")
		assert.NoError(t, err, "Failed to get basic config")

		assert.NoError(t, ConfigureClient(c), "ConfigureClient should only return an error if the server is down")
		assert.Equal(t, conf.AuthType, AuthTypeBasic, "Expected AuthType to be basic after ConfigureClient")
		assert.Equal(t, conf.User, c.User, "Configured User should match config User")
		assert.Equal(t, conf.Password, c.Password, "Configured Password should match config Password")
		assert.Empty(t, conf.LoginPath, "Expected LoginPath to be empty for basic auth")
	})

	t.Run("test cookie auth without LoginPath returns error", func(t *testing.T) {
		c, err := GetConfig("cookie")
		assert.NoError(t, err, "Failed to get cookie config")

		// Clear LoginPath to test error handling
		c.LoginPath = ""

		assert.Error(t, ConfigureClient(c), "Expected ConfigureClient to return error when LoginPath is not set for cookie auth")
	})
}

func TestBackwardCompatibility(t *testing.T) {
	t.Run("test config without AuthType defaults to basic auth", func(t *testing.T) {
		// Create a simple test server that checks for basic auth
		ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			user, pass, ok := r.BasicAuth()
			if !ok || user != "foo" || pass != "bar" {
				http.Error(w, "Unauthorized", http.StatusUnauthorized)
				return
			}
			w.WriteHeader(http.StatusOK)
			w.Write([]byte(`{"status":"ok"}`))
		}))
		defer ts.Close()

		// Get basic config and clear AuthType to test default behavior
		c, err := GetConfig("basic")
		if err != nil {
			t.Fatalf("Failed to get basic config: %v", err)
		}
		c.BaseURL = ts.URL
		// Clear AuthType to test default behavior
		c.AuthType = ""

		err = ConfigureClient(c)
		if err != nil {
			t.Fatalf("Unexpected error: %v", err)
		}

		// Test that basic auth is used
		resp, err := Get("/api/v0/test")
		if err != nil {
			t.Errorf("Unexpected error: %s", err)
		}

		if resp.StatusCode != 200 {
			t.Errorf("Expected status 200, got %d", resp.StatusCode)
		}
	})
}

func TestCookieAuthWithMockServer(t *testing.T) {
	// Get expected credentials from config
	cookieConfig, err := GetConfig("cookie")
	if err != nil {
		t.Fatalf("Failed to get cookie config: %v", err)
	}

	// Create a mock server
	loginCalled := false
	apiCalled := false
	sessionCookie := "test-session-cookie"

	mux := http.NewServeMux()

	// Mock login endpoint
	mux.HandleFunc("/api/v0/login", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "POST" {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}

		loginCalled = true

		// Verify the request body contains email and password
		var loginReq struct {
			Email    string `json:"email"`
			Password string `json:"password"`
		}
		if err := json.NewDecoder(r.Body).Decode(&loginReq); err != nil {
			http.Error(w, "Bad request", http.StatusBadRequest)
			return
		}

		if loginReq.Email != cookieConfig.User || loginReq.Password != cookieConfig.Password {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		// Set a session cookie
		http.SetCookie(w, &http.Cookie{
			Name:  "session",
			Value: sessionCookie,
			Path:  "/",
		})

		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"status":"ok"}`))
	})

	// Mock API endpoint that requires authentication
	mux.HandleFunc("/api/v0/test", func(w http.ResponseWriter, r *http.Request) {
		apiCalled = true

		// Check for session cookie
		cookie, err := r.Cookie("session")
		if err != nil || cookie.Value != sessionCookie {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"data":"test"}`))
	})

	ts := httptest.NewServer(mux)
	defer ts.Close()

	t.Run("test successful cookie auth flow", func(t *testing.T) {
		// Get cookie config from configuration file
		c, err := GetConfig("cookie")
		if err != nil {
			t.Fatalf("Failed to get cookie config: %v", err)
		}

		// Override BaseURL to use the test server
		c.BaseURL = ts.URL
		// Set LoginPath for this specific test
		c.LoginPath = "/api/v0/login"

		err = ConfigureClient(c)
		if err != nil {
			t.Fatalf("Unexpected error during ConfigureClient: %v", err)
		}

		if !loginCalled {
			t.Errorf("Login endpoint was not called during ConfigureClient")
		}

		// Test that subsequent API calls use the cookie
		resp, err := Get("/api/v0/test")
		if err != nil {
			t.Errorf("Unexpected error: %s", err)
		}

		if resp.StatusCode != 200 {
			t.Errorf("Expected status 200, got %d", resp.StatusCode)
		}

		if !apiCalled {
			t.Errorf("API endpoint was not called")
		}
	})
}
