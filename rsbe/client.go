package rsbe

import (
	"fmt"
	"log"
	"net/http"
	"io/ioutil"
	"os"
)

type Config struct {
	rsbe_url      string
	rsbe_user     string
	rsbe_password string
}

var conf Config
var client *http.Client

func init() {
	init_global_config()
	client = &http.Client{}
}

func lookupEnvKeyOrDefault(env_key string, default_value string) string {
	val, found := os.LookupEnv(env_key)

	if !found {
		val = default_value
	}

	return val
}

func init_global_config() {
	conf.rsbe_url = lookupEnvKeyOrDefault("FLOW_UI_RSBE_URL", "http://localhost:3000")
	conf.rsbe_user = lookupEnvKeyOrDefault("FLOW_UI_RSBE_USER", "foo")
	conf.rsbe_password = lookupEnvKeyOrDefault("FLOW_UI_RSBE_PASSWORD", "bar")
}

func Get(path string) (resp *http.Response, err error) {
	url := fmt.Sprintf("%s%s", conf.rsbe_url, path)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	req.SetBasicAuth(conf.rsbe_user, conf.rsbe_password)
	return client.Do(req)
}

func GetBodyTextString(path string) (s string, err error) {
	s = ""
	url := fmt.Sprintf("%s%s", conf.rsbe_url, path)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatal(err)
	}
	req.SetBasicAuth(conf.rsbe_user, conf.rsbe_password)

	var resp *http.Response
	resp, err = client.Do(req)
	if err != nil {
		return s, err
	}

	bodyText, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return s, err
	}

	
	return string(bodyText), nil

	// resp, err := Get(path)
	// if err != nil {
	// 	return partner, err
	// }

	// bodyText, err := ioutil.ReadAll(resp.Body)
	// if err != nil {
	// 	return partner, err
	// }

	// s := string(bodyText)

}
