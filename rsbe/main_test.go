package rsbe

import (
	"fmt"
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	err := LoadConfig()
	if err != nil {
		errMsg := fmt.Sprintf(
			"Failed to load configuration:\n"+
				"%v\n\n"+
				"To resolve this issue:\n"+
				"1. Create a test configuration file (e.g., .env.test.yaml)\n"+
				"2. Set APP_ENV_FILE_PATH to point to your test config file:\n"+
				"   $ export APP_ENV_FILE_PATH=/path/to/.env.test.yaml\n"+
				"3. Ensure your config file has 'environment: test' set\n"+
				"Then run the tests again.\n",
			err)
		panic(errMsg)
	}

	if Cfg.Environment != "test" {
		errMsg := fmt.Sprintf(
			"Incorrect Config File:\n"+
				"It looks like you're using a non-test config file when testing.\n"+
				"APP_ENV must be set to 'test' when testing but APP_ENV is currently set to '%s'\n"+
				"This check prevents the accidental modification of non-test databases\n"+
				"when testing.\n\n"+
				"To resolve this issue, create a config file for your test environment\n"+
				"and set the environment variable in the test-config file to 'test'.\n\n"+
				"You can do this by creating a .env.test.yaml file in your project root,\n"+
				"or by using an existing config file that has environment: test.\n\n"+
				"Once your test-config file exists, export the APP_ENV_FILE_PATH variable\n"+
				"and set the value to the path of your test-config file.\n"+
				"e.g, \n"+
				"  $ export APP_ENV_FILE_PATH=/path/to/go-rsbe-client/.env.test.yaml\n"+
				"  $ grep environment $APP_ENV_FILE_PATH \n"+
				"  environment: test\n"+
				"  $\n"+
				"Then run the tests again.\n",
			Cfg.Environment)
		panic(errMsg)
	}

	os.Exit(m.Run())
}
