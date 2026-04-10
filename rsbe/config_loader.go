package rsbe

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v3"
)

const UseDefaultConfig string = "use_default_config"

// TestConfig represents the structure of the test configuration file
type TestConfig struct {
	Environment   string                 `yaml:"environment"`
	DefaultConfig string                 `yaml:"default_config"`
	Configs       map[string]ConfigEntry `yaml:"configs"`
}

// ConfigEntry represents a configuration entry for a specific auth type
type ConfigEntry struct {
	BaseURL   string `yaml:"BaseURL"`
	User      string `yaml:"User"`
	Password  string `yaml:"Password"`
	AuthType  string `yaml:"AuthType,omitempty"`
	LoginPath string `yaml:"LoginPath,omitempty"`
}

var Cfg *TestConfig

// LoadConfig loads the test configuration from the file specified by APP_ENV_FILE_PATH
func LoadConfig() error {
	configPath := os.Getenv("APP_ENV_FILE_PATH")
	if configPath == "" {
		return fmt.Errorf("APP_ENV_FILE_PATH environment variable is not set")
	}

	data, err := os.ReadFile(configPath)
	if err != nil {
		return fmt.Errorf("failed to read config file: %w", err)
	}

	var cfg TestConfig
	if err := yaml.Unmarshal(data, &cfg); err != nil {
		return fmt.Errorf("failed to parse config file: %w", err)
	}

	Cfg = &cfg
	return nil
}

// GetConfig returns a Config struct for the specified configuration key
// passing the special key UseDefaultConfig to GetConfig selects the
// default configuration specified in the config file under the "default_config" key.
// This allows tests to specify which configuration (and therefore auth method)
// to use without changing the test code.
func GetConfig(key string) (*Config, error) {

	if Cfg == nil {
		return nil, fmt.Errorf("configuration not loaded. Call LoadConfig() first")
	}

	if key == UseDefaultConfig {
		if Cfg.DefaultConfig == "" {
			return nil, fmt.Errorf("default_config not specified in config file")
		}
		key = Cfg.DefaultConfig
	}

	entry, ok := Cfg.Configs[key]
	if !ok {
		return nil, fmt.Errorf("configuration key '%s' not found", key)
	}

	config := &Config{
		BaseURL:  entry.BaseURL,
		User:     entry.User,
		Password: entry.Password,
	}

	// Set AuthType if provided, otherwise default to basic
	if entry.AuthType != "" {
		config.AuthType = entry.AuthType
	} else {
		config.AuthType = AuthTypeBasic
	}

	// Set LoginPath if provided
	if entry.LoginPath != "" {
		config.LoginPath = entry.LoginPath
	}

	return config, nil
}
