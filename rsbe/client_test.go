package rsbe

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"path"
	"strings"
	"sync"
	"sync/atomic"
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
		err = PartnerPurge(id)
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

		err = PartnerPurge(p.ID)
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
		defer PartnerPurge(id)

		data := "{\"code\":\"bananananana\",\"name\":\"Can elope\",\"rel_path\":\"content/canteloupe\",\"lock_version\":0}"
		err = Put(ppath+"/"+id, []byte(data))
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

		// still need to purge here because GORM soft-delete only marks the Partner as deleted,
		// it doesn't actually remove it from the database
		PartnerPurge(id)
	})

	t.Run("confirm bad request response", func(t *testing.T) {

		err := Delete("/api/v0/partners/adflasdfj")
		if err == nil {
			t.Errorf("err should NOT be nil: %v", err)
		}
	})

}

func TestCookieAuth(t *testing.T) {
	t.Run("test gorsbe config with cookie auth", func(t *testing.T) {
		c, err := GetConfig("gorsbe")
		assert.NoError(t, err, "Failed to get gorsbe config")

		assert.NoError(t, ConfigureClient(c), "ConfigureClient should only return an error if the server is down")
		assert.NoError(t, Login(), "Login should succeed for cookie auth when LoginPath is configured")
		assert.Equal(t, conf.AuthType, AuthTypeCookie, "AuthType should be set to cookie after ConfigureClient")
		assert.Equal(t, conf.User, c.User, "Configured User should match config User")
		assert.Equal(t, conf.Password, c.Password, "Configured Password should match config Password")
		assert.NotEmpty(t, conf.LoginPath, "Expected LoginPath to be set")
	})

	t.Run("test rails config with basic auth", func(t *testing.T) {
		c, err := GetConfig("rails")
		assert.NoError(t, err, "Failed to get rails config")

		assert.NoError(t, ConfigureClient(c), "ConfigureClient should only return an error if the server is down")
		assert.Equal(t, conf.AuthType, AuthTypeBasic, "Expected AuthType to be basic after ConfigureClient")
		assert.Equal(t, conf.User, c.User, "Configured User should match config User")
		assert.Equal(t, conf.Password, c.Password, "Configured Password should match config Password")
		assert.Empty(t, conf.LoginPath, "Expected LoginPath to be empty for basic auth")
	})

	t.Run("test gorsbe config with cookie auth without LoginPath returns error on Login", func(t *testing.T) {
		c, err := GetConfig("gorsbe")
		assert.NoError(t, err, "Failed to get gorsbe config")

		// Clear LoginPath to test error handling
		c.LoginPath = ""

		assert.NoError(t, ConfigureClient(c), "ConfigureClient should not fail when LoginPath is missing")
		assert.Error(t, Login(), "Expected Login to return error when LoginPath is not set for cookie auth")
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
		c, err := GetConfig("rails")
		if err != nil {
			t.Fatalf("Failed to get rails config: %v", err)
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
	gorsbeConfig, err := GetConfig("gorsbe")
	if err != nil {
		t.Fatalf("Failed to get gorsbe config: %v", err)
	}

	// Create a mock server
	loginCalled := false
	logoutCalled := false
	apiCalled := false
	sessionCookie := "test-session-cookie"

	mux := http.NewServeMux()

	// Mock login/logout endpoint
	mux.HandleFunc("/api/v0/login", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case "POST":
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

			if loginReq.Email != gorsbeConfig.User || loginReq.Password != gorsbeConfig.Password {
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
		case "DELETE":
			logoutCalled = true
			w.WriteHeader(http.StatusNoContent)
		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
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

	t.Run("test successful gorsbe config with cookie auth flow", func(t *testing.T) {
		disabled := false
		// Get gorsbe config from configuration file
		c, err := GetConfig("gorsbe")
		if err != nil {
			t.Fatalf("Failed to get gorsbe config: %v", err)
		}

		// Override BaseURL to use the test server
		c.BaseURL = ts.URL
		// Set LoginPath for this specific test
		c.LoginPath = "/api/v0/login"
		// Keep this test focused on explicit Login/Logout behavior.
		c.AutoRefreshOn401 = &disabled

		err = ConfigureClient(c)
		if err != nil {
			t.Fatalf("Unexpected error during ConfigureClient: %v", err)
		}

		if loginCalled {
			t.Errorf("Login endpoint should not be called during ConfigureClient")
		}

		err = Login()
		if err != nil {
			t.Fatalf("Unexpected error during Login: %v", err)
		}

		if !loginCalled {
			t.Errorf("Login endpoint was not called during Login")
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

		assert.NoError(t, Logout(), "Logout should clear cookie auth state")
		assert.True(t, logoutCalled, "Expected Logout to call DELETE on the login endpoint")

		_, err = Get("/api/v0/test")
		assert.Error(t, err, "Expected unauthorized error after Logout")
	})
}

func TestAutoRefreshOn401CookieAuth(t *testing.T) {
	t.Run("401 triggers login and retries once", func(t *testing.T) {
		var loginCalls int32
		sessionCookie := "fresh-session"

		mux := http.NewServeMux()
		mux.HandleFunc("/api/v0/login", func(w http.ResponseWriter, r *http.Request) {
			assert.Equal(t, http.MethodPost, r.Method)
			atomic.AddInt32(&loginCalls, 1)
			http.SetCookie(w, &http.Cookie{Name: "session", Value: sessionCookie, Path: "/"})
			w.WriteHeader(http.StatusOK)
			_, _ = w.Write([]byte(`{"status":"ok"}`))
		})
		mux.HandleFunc("/api/v0/protected", func(w http.ResponseWriter, r *http.Request) {
			cookie, err := r.Cookie("session")
			if err != nil || cookie.Value != sessionCookie {
				w.WriteHeader(http.StatusUnauthorized)
				_, _ = w.Write([]byte(`{"error":"unauthorized"}`))
				return
			}
			w.WriteHeader(http.StatusOK)
			_, _ = w.Write([]byte(`{"ok":true}`))
		})

		ts := httptest.NewServer(mux)
		defer ts.Close()

		c := &Config{BaseURL: ts.URL, User: "foo@example.org", Password: "bar", AuthType: AuthTypeCookie, LoginPath: "/api/v0/login"}
		assert.NoError(t, ConfigureClient(c))

		resp, err := Get("/api/v0/protected")
		assert.NoError(t, err)
		assert.Equal(t, http.StatusOK, resp.StatusCode)
		assert.Equal(t, int32(1), atomic.LoadInt32(&loginCalls))
	})

	t.Run("401 and refresh login failure returns refresh context", func(t *testing.T) {
		var loginCalls int32

		mux := http.NewServeMux()
		mux.HandleFunc("/api/v0/login", func(w http.ResponseWriter, r *http.Request) {
			atomic.AddInt32(&loginCalls, 1)
			w.WriteHeader(http.StatusUnauthorized)
			_, _ = w.Write([]byte(`{"error":"invalid credentials"}`))
		})
		mux.HandleFunc("/api/v0/protected", func(w http.ResponseWriter, _ *http.Request) {
			w.WriteHeader(http.StatusUnauthorized)
			_, _ = w.Write([]byte(`{"error":"expired session"}`))
		})

		ts := httptest.NewServer(mux)
		defer ts.Close()

		c := &Config{BaseURL: ts.URL, User: "foo@example.org", Password: "wrong", AuthType: AuthTypeCookie, LoginPath: "/api/v0/login"}
		assert.NoError(t, ConfigureClient(c))

		_, err := Get("/api/v0/protected")
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "auth refresh failed")
		assert.Equal(t, int32(1), atomic.LoadInt32(&loginCalls))
	})

	t.Run("401 then refresh success then retry 401 returns unauthorized", func(t *testing.T) {
		var loginCalls int32

		mux := http.NewServeMux()
		mux.HandleFunc("/api/v0/login", func(w http.ResponseWriter, r *http.Request) {
			atomic.AddInt32(&loginCalls, 1)
			http.SetCookie(w, &http.Cookie{Name: "session", Value: "fresh", Path: "/"})
			w.WriteHeader(http.StatusOK)
			_, _ = w.Write([]byte(`{"status":"ok"}`))
		})
		mux.HandleFunc("/api/v0/protected", func(w http.ResponseWriter, _ *http.Request) {
			w.WriteHeader(http.StatusUnauthorized)
			_, _ = w.Write([]byte(`{"error":"still unauthorized"}`))
		})

		ts := httptest.NewServer(mux)
		defer ts.Close()

		c := &Config{BaseURL: ts.URL, User: "foo@example.org", Password: "bar", AuthType: AuthTypeCookie, LoginPath: "/api/v0/login"}
		assert.NoError(t, ConfigureClient(c))

		resp, err := Get("/api/v0/protected")
		assert.Error(t, err)
		assert.Equal(t, http.StatusUnauthorized, resp.StatusCode)
		assert.Contains(t, err.Error(), "bad response: 401")
		assert.Equal(t, int32(1), atomic.LoadInt32(&loginCalls))
	})

	t.Run("basic auth 401 does not auto refresh", func(t *testing.T) {
		var loginCalls int32

		mux := http.NewServeMux()
		mux.HandleFunc("/api/v0/login", func(w http.ResponseWriter, _ *http.Request) {
			atomic.AddInt32(&loginCalls, 1)
			w.WriteHeader(http.StatusOK)
		})
		mux.HandleFunc("/api/v0/protected", func(w http.ResponseWriter, _ *http.Request) {
			w.WriteHeader(http.StatusUnauthorized)
			_, _ = w.Write([]byte(`{"error":"unauthorized"}`))
		})

		ts := httptest.NewServer(mux)
		defer ts.Close()

		c := &Config{BaseURL: ts.URL, User: "foo", Password: "bar", AuthType: AuthTypeBasic, LoginPath: "/api/v0/login"}
		assert.NoError(t, ConfigureClient(c))

		resp, err := Get("/api/v0/protected")
		assert.Error(t, err)
		assert.Equal(t, http.StatusUnauthorized, resp.StatusCode)
		assert.Equal(t, int32(0), atomic.LoadInt32(&loginCalls))
	})

	t.Run("non-401 does not auto refresh", func(t *testing.T) {
		var loginCalls int32

		mux := http.NewServeMux()
		mux.HandleFunc("/api/v0/login", func(w http.ResponseWriter, _ *http.Request) {
			atomic.AddInt32(&loginCalls, 1)
			w.WriteHeader(http.StatusOK)
		})
		mux.HandleFunc("/api/v0/protected", func(w http.ResponseWriter, _ *http.Request) {
			w.WriteHeader(http.StatusBadRequest)
			_, _ = w.Write([]byte(`{"error":"bad request"}`))
		})

		ts := httptest.NewServer(mux)
		defer ts.Close()

		c := &Config{BaseURL: ts.URL, User: "foo@example.org", Password: "bar", AuthType: AuthTypeCookie, LoginPath: "/api/v0/login"}
		assert.NoError(t, ConfigureClient(c))

		resp, err := Get("/api/v0/protected")
		assert.Error(t, err)
		assert.Equal(t, http.StatusBadRequest, resp.StatusCode)
		assert.Equal(t, int32(0), atomic.LoadInt32(&loginCalls))
	})
}

func TestAutoRefreshBodyReplay(t *testing.T) {
	t.Run("post and put retry preserve body and content type", func(t *testing.T) {
		sessionCookie := "replay-cookie"
		var postBodies [][]byte
		var putBodies [][]byte
		var postContentTypes []string
		var putContentTypes []string
		var postCalls int
		var putCalls int
		var mu sync.Mutex

		mux := http.NewServeMux()
		mux.HandleFunc("/api/v0/login", func(w http.ResponseWriter, _ *http.Request) {
			http.SetCookie(w, &http.Cookie{Name: "session", Value: sessionCookie, Path: "/"})
			w.WriteHeader(http.StatusOK)
		})

		mux.HandleFunc("/api/v0/post-target", func(w http.ResponseWriter, r *http.Request) {
			body, _ := io.ReadAll(r.Body)
			cookie, err := r.Cookie("session")

			mu.Lock()
			postCalls++
			postBodies = append(postBodies, body)
			postContentTypes = append(postContentTypes, r.Header.Get("Content-Type"))
			currentCall := postCalls
			mu.Unlock()

			if currentCall == 1 {
				w.WriteHeader(http.StatusUnauthorized)
				_, _ = w.Write([]byte(`{"error":"expired"}`))
				return
			}

			if err != nil || cookie.Value != sessionCookie {
				w.WriteHeader(http.StatusUnauthorized)
				_, _ = w.Write([]byte(`{"error":"expired"}`))
				return
			}

			w.WriteHeader(http.StatusCreated)
		})

		mux.HandleFunc("/api/v0/put-target", func(w http.ResponseWriter, r *http.Request) {
			body, _ := io.ReadAll(r.Body)
			cookie, err := r.Cookie("session")

			mu.Lock()
			putCalls++
			putBodies = append(putBodies, body)
			putContentTypes = append(putContentTypes, r.Header.Get("Content-Type"))
			currentCall := putCalls
			mu.Unlock()

			if currentCall == 1 {
				w.WriteHeader(http.StatusUnauthorized)
				_, _ = w.Write([]byte(`{"error":"expired"}`))
				return
			}

			if err != nil || cookie.Value != sessionCookie {
				w.WriteHeader(http.StatusUnauthorized)
				_, _ = w.Write([]byte(`{"error":"expired"}`))
				return
			}

			w.WriteHeader(http.StatusNoContent)
		})

		ts := httptest.NewServer(mux)
		defer ts.Close()

		c := &Config{BaseURL: ts.URL, User: "foo@example.org", Password: "bar", AuthType: AuthTypeCookie, LoginPath: "/api/v0/login"}
		assert.NoError(t, ConfigureClient(c))

		postPayload := []byte(`{"a":1,"b":"two"}`)
		resp, err := Post("/api/v0/post-target", postPayload)
		assert.NoError(t, err)
		assert.Equal(t, http.StatusCreated, resp.StatusCode)

		putPayload := []byte(`{"x":123,"y":"abc"}`)
		assert.NoError(t, Put("/api/v0/put-target", putPayload))

		if !assert.Len(t, postBodies, 2) {
			return
		}
		assert.True(t, bytes.Equal(postPayload, postBodies[0]))
		assert.True(t, bytes.Equal(postPayload, postBodies[1]))
		assert.Equal(t, "application/json", postContentTypes[0])
		assert.Equal(t, "application/json", postContentTypes[1])

		if !assert.Len(t, putBodies, 2) {
			return
		}
		assert.True(t, bytes.Equal(putPayload, putBodies[0]))
		assert.True(t, bytes.Equal(putPayload, putBodies[1]))
		assert.Equal(t, "application/json", putContentTypes[0])
		assert.Equal(t, "application/json", putContentTypes[1])
	})
}

func TestAutoRefreshConcurrencySingleLogin(t *testing.T) {
	var loginCalls int32
	sessionCookie := "shared-session"

	mux := http.NewServeMux()
	mux.HandleFunc("/api/v0/login", func(w http.ResponseWriter, _ *http.Request) {
		atomic.AddInt32(&loginCalls, 1)
		http.SetCookie(w, &http.Cookie{Name: "session", Value: sessionCookie, Path: "/"})
		w.WriteHeader(http.StatusOK)
	})
	mux.HandleFunc("/api/v0/protected", func(w http.ResponseWriter, r *http.Request) {
		cookie, err := r.Cookie("session")
		if err != nil || cookie.Value != sessionCookie {
			w.WriteHeader(http.StatusUnauthorized)
			_, _ = w.Write([]byte(`{"error":"expired"}`))
			return
		}
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(`{"ok":true}`))
	})

	ts := httptest.NewServer(mux)
	defer ts.Close()

	c := &Config{BaseURL: ts.URL, User: "foo@example.org", Password: "bar", AuthType: AuthTypeCookie, LoginPath: "/api/v0/login", AuthRefreshCooldownSeconds: 1}
	assert.NoError(t, ConfigureClient(c))

	const goroutines = 12
	errCh := make(chan error, goroutines)
	var wg sync.WaitGroup

	for i := 0; i < goroutines; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			resp, err := Get("/api/v0/protected")
			if err != nil {
				errCh <- err
				return
			}
			if resp.StatusCode != http.StatusOK {
				errCh <- fmt.Errorf("unexpected status: %d", resp.StatusCode)
				return
			}
		}()
	}

	wg.Wait()
	close(errCh)

	for err := range errCh {
		if err != nil {
			t.Fatalf("unexpected concurrent request error: %v", err)
		}
	}

	if atomic.LoadInt32(&loginCalls) != 1 {
		t.Fatalf("expected exactly one login refresh call, got %d", atomic.LoadInt32(&loginCalls))
	}
}

func TestAutoRefreshToggleAndRetryCount(t *testing.T) {
	t.Run("auto refresh disabled prevents login attempts", func(t *testing.T) {
		var loginCalls int32
		disabled := false

		mux := http.NewServeMux()
		mux.HandleFunc("/api/v0/login", func(w http.ResponseWriter, _ *http.Request) {
			atomic.AddInt32(&loginCalls, 1)
			w.WriteHeader(http.StatusOK)
		})
		mux.HandleFunc("/api/v0/protected", func(w http.ResponseWriter, _ *http.Request) {
			w.WriteHeader(http.StatusUnauthorized)
			_, _ = w.Write([]byte(`{"error":"expired"}`))
		})

		ts := httptest.NewServer(mux)
		defer ts.Close()

		c := &Config{BaseURL: ts.URL, User: "foo@example.org", Password: "bar", AuthType: AuthTypeCookie, LoginPath: "/api/v0/login", AutoRefreshOn401: &disabled}
		assert.NoError(t, ConfigureClient(c))

		_, err := Get("/api/v0/protected")
		assert.Error(t, err)
		assert.Equal(t, int32(0), atomic.LoadInt32(&loginCalls))
	})

	t.Run("max auth retries controls retry count", func(t *testing.T) {
		var loginCalls int32
		var protectedCalls int32

		mux := http.NewServeMux()
		mux.HandleFunc("/api/v0/login", func(w http.ResponseWriter, _ *http.Request) {
			atomic.AddInt32(&loginCalls, 1)
			http.SetCookie(w, &http.Cookie{Name: "session", Value: "never-accepted", Path: "/"})
			w.WriteHeader(http.StatusOK)
		})
		mux.HandleFunc("/api/v0/protected", func(w http.ResponseWriter, _ *http.Request) {
			atomic.AddInt32(&protectedCalls, 1)
			w.WriteHeader(http.StatusUnauthorized)
			_, _ = w.Write([]byte(`{"error":"expired"}`))
		})

		ts := httptest.NewServer(mux)
		defer ts.Close()

		c := &Config{BaseURL: ts.URL, User: "foo@example.org", Password: "bar", AuthType: AuthTypeCookie, LoginPath: "/api/v0/login", MaxAuthRetries: 2}
		assert.NoError(t, ConfigureClient(c))

		_, err := Get("/api/v0/protected")
		assert.Error(t, err)
		assert.True(t, strings.Contains(err.Error(), "bad response: 401"))
		assert.Equal(t, int32(2), atomic.LoadInt32(&loginCalls))
		assert.Equal(t, int32(3), atomic.LoadInt32(&protectedCalls))
	})
}

func TestCookieSessionExpiryMiniIntegration(t *testing.T) {
	var loginCalls int32
	var unauthorizedCalls int32
	var protectedCalls int32

	const sessionCookie = "integration-session"

	var mu sync.Mutex
	injectUnauthorizedOnce := true

	mux := http.NewServeMux()
	mux.HandleFunc("/api/v0/login", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}
		atomic.AddInt32(&loginCalls, 1)
		http.SetCookie(w, &http.Cookie{Name: "session", Value: sessionCookie, Path: "/"})
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(`{"status":"ok"}`))
	})

	authorize := func(w http.ResponseWriter, r *http.Request) bool {
		atomic.AddInt32(&protectedCalls, 1)

		mu.Lock()
		if injectUnauthorizedOnce {
			injectUnauthorizedOnce = false
			mu.Unlock()
			atomic.AddInt32(&unauthorizedCalls, 1)
			w.WriteHeader(http.StatusUnauthorized)
			_, _ = w.Write([]byte(`{"error":"expired session"}`))
			return false
		}
		mu.Unlock()

		cookie, err := r.Cookie("session")
		if err != nil || cookie.Value != sessionCookie {
			atomic.AddInt32(&unauthorizedCalls, 1)
			w.WriteHeader(http.StatusUnauthorized)
			_, _ = w.Write([]byte(`{"error":"unauthorized"}`))
			return false
		}

		return true
	}

	mux.HandleFunc("/api/v0/integration/get", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}
		if !authorize(w, r) {
			return
		}
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(`{"ok":true,"method":"get"}`))
	})

	mux.HandleFunc("/api/v0/integration/post", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}
		if !authorize(w, r) {
			return
		}
		w.Header().Set("Location", "/api/v0/integration/post/abc123")
		w.WriteHeader(http.StatusCreated)
	})

	mux.HandleFunc("/api/v0/integration/put", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPut {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}
		if !authorize(w, r) {
			return
		}
		w.WriteHeader(http.StatusNoContent)
	})

	mux.HandleFunc("/api/v0/integration/delete", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodDelete {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}
		if !authorize(w, r) {
			return
		}
		w.WriteHeader(http.StatusNoContent)
	})

	ts := httptest.NewServer(mux)
	defer ts.Close()

	c := &Config{
		BaseURL:   ts.URL,
		User:      "foo@example.org",
		Password:  "bar",
		AuthType:  AuthTypeCookie,
		LoginPath: "/api/v0/login",
	}
	assert.NoError(t, ConfigureClient(c))

	resp, err := Get("/api/v0/integration/get")
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, resp.StatusCode)

	postResp, err := Post("/api/v0/integration/post", []byte(`{"message":"hello"}`))
	assert.NoError(t, err)
	assert.Equal(t, http.StatusCreated, postResp.StatusCode)

	assert.NoError(t, Put("/api/v0/integration/put", []byte(`{"status":"updated"}`)))
	assert.NoError(t, Delete("/api/v0/integration/delete"))

	assert.Equal(t, int32(1), atomic.LoadInt32(&loginCalls), "expected exactly one refresh login")
	assert.Equal(t, int32(1), atomic.LoadInt32(&unauthorizedCalls), "expected one initial unauthorized due to expired session")
	assert.Equal(t, int32(5), atomic.LoadInt32(&protectedCalls), "expected one retry plus three subsequent successful calls")
}
