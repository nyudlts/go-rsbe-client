# `go-rsbe-client`

This repo contains a client package for the R* Back End (`RSBE`) API written in Go.  
The library supports two forms of authentication:  

* `HTTP Basic Auth`  (used with the [Rails-based `RSBE` implementation](https://github.com/nyudlts/rsbe))
* cookie based authentication (used with the [Go-based `RSBE` implementation](https://github.com/nyudlts/go-rsbe))

Before placing requests to either implementation, you must call the  
`rsbe.ConfigureClient(&cfg)` function with a properly initialized `Config` variable.

```go
type Config struct {
  BaseURL   string
  User      string
  Password  string
  AuthType  string
  LoginPath string
}
```

Please reference the [`.env.test.yaml.example`](./.env.test.yaml.example) file for sample values.

## Testing

The test suite requires a configuration file to be specified via the `APP_ENV_FILE_PATH` environment variable.

### Setup Test Configuration

1. Create a test configuration file (e.g., `.env.test.yaml`) based on the example in `cfg/config.yaml.example`
2. Ensure the configuration file has `environment: test` set
3. Configure both `basic` and `cookie` authentication entries as needed

Example configuration:

```yaml
environment: test

configs:
  basic:
    BaseURL:  "http://localhost:3000/"
    User:     "foo"
    Password: "bar"
  
  cookie:
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
