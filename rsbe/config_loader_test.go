package rsbe

import (
	"os"
	"testing"
)

func TestLoadConfig(t *testing.T) {
	// Save original env var
	originalPath := os.Getenv("APP_ENV_FILE_PATH")
	defer func() {
		if originalPath != "" {
			os.Setenv("APP_ENV_FILE_PATH", originalPath)
		}
	}()

	t.Run("test loading config from file", func(t *testing.T) {
		// This test assumes APP_ENV_FILE_PATH is already set by TestMain
		if Cfg == nil {
			t.Fatal("Config should be loaded by TestMain")
		}

		if Cfg.Environment != "test" {
			t.Errorf("Expected environment to be 'test', got '%s'", Cfg.Environment)
		}
	})

	t.Run("test GetConfig for rails config with basic auth", func(t *testing.T) {
		config, err := GetConfig("rails")
		if err != nil {
			t.Fatalf("Unexpected error: %v", err)
		}

		if config.BaseURL != "http://localhost:3000" {
			t.Errorf("Expected BaseURL to be 'http://localhost:3000', got '%s'", config.BaseURL)
		}

		if config.AuthType != AuthTypeBasic {
			t.Errorf("Expected AuthType to be 'basic', got '%s'", config.AuthType)
		}
	})

	t.Run("test GetConfig for gorsbe config with cookie auth", func(t *testing.T) {
		config, err := GetConfig("gorsbe")
		if err != nil {
			t.Fatalf("Unexpected error: %v", err)
		}

		if config.BaseURL != "http://localhost:3001" {
			t.Errorf("Expected BaseURL to be 'http://localhost:3001', got '%s'", config.BaseURL)
		}

		if config.AuthType != AuthTypeCookie {
			t.Errorf("Expected AuthType to be 'cookie', got '%s'", config.AuthType)
		}

		if config.LoginPath != "/api/v0/sessions" {
			t.Errorf("Expected LoginPath to be '/api/v0/sessions', got '%s'", config.LoginPath)
		}
	})

	t.Run("test GetConfig with invalid key", func(t *testing.T) {
		_, err := GetConfig("invalid_key")
		if err == nil {
			t.Error("Expected error for invalid key, got nil")
		}
	})
}
