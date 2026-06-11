# `go-rsbe-client`

This repo contains a client package for the R* Back End (`RSBE`) API written in Go.  
The library supports two forms of authentication:  

* `HTTP Basic Auth`  (used with the [Rails-based `RSBE` implementation](https://github.com/nyudlts/rsbe))
* cookie based authentication (used with the [Go-based `RSBE` implementation](https://github.com/nyudlts/go-rsbe))

Before placing requests, call `rsbe.ConfigureClient(&cfg)` with a properly initialized `Config` variable.
For cookie auth, call `rsbe.Login()` after configuration.

`ConfigureClient()` now only configures the client and does not perform network authentication.
If you are using cookie auth, login and logout are explicit operations:

* `rsbe.Login()` authenticates and stores session cookies.
* `rsbe.Logout()` sends a `DELETE` request to `LoginPath` and clears client cookies.

```go
type Config struct {
  BaseURL   string
  User      string
  Password  string
  AuthType  string
  LoginPath string
}

// Basic auth
if err := rsbe.ConfigureClient(&cfg); err != nil {
  panic(err)
}

// Cookie auth
if err := rsbe.ConfigureClient(&cfg); err != nil {
  panic(err)
}
if err := rsbe.Login(); err != nil {
  panic(err)
}

// ... make API calls ...

if err := rsbe.Logout(); err != nil {
  panic(err)
}
```

Migration note: if existing code relied on `ConfigureClient()` to automatically authenticate
cookie sessions, add an explicit `rsbe.Login()` call immediately after configuration.

Please reference the [`cfg/config.yaml.example`](./cfg/config.yaml.example) file for sample values.

## Testing

The test suite requires a configuration file to be specified via the `APP_ENV_FILE_PATH` environment variable.

### Setup Test Configuration

1. Create a test configuration file (e.g., `.env.test.yaml`) based on [`cfg/config.yaml.example`](./cfg/config.yaml.example)
2. Ensure the configuration file has `environment: test` set
3. Update the `rails` and `gorsbe` configurations as needed

Example configuration:

```yaml
environment: test

configs:
  rails:
    BaseURL:  "http://localhost:3000/"
    User:     "foo"
    Password: "bar"
  
  gorsbe:
    BaseURL:   "http://localhost:3000/"
    User:      "foo"
    Password:  "bar"
    AuthType:  "cookie"
    LoginPath: "/api/v0/sessions"
```

### Running Tests

Export the configuration file path and run the tests:

```bash
export APP_ENV_FILE_PATH=/path/to/.env.test.yaml
go test ./...
```

The test suite will validate that the `environment` field is set to `test`  
to prevent accidental modification of non-test databases.
