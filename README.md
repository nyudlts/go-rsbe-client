## `go-rsbe-client`

This is an [`rsbe`](https://github.com/nyudlts/rsbe) API client package
written in Golang.

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

The test suite will validate that the `environment` field is set to `test` to prevent accidental modification of non-test databases.

