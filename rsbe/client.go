package rsbe

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

type Config struct {
	BaseURL  string
	User     string
	Password string
}

var conf *Config
var client *http.Client

func ConfigureClient(c *Config) {
	conf = c
}

func init() {
	client = &http.Client{}
}

func Get(path string) (resp *http.Response, err error) {
	url := fmt.Sprintf("%s%s", conf.BaseURL, path)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	req.SetBasicAuth(conf.User, conf.Password)
	return client.Do(req)
}

func GetBodyTextString(path string) (s string, err error) {
	s = ""

	resp, err := Get(path)
	if err != nil {
		return s, err
	}

	bodyText, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return s, err
	}

	return string(bodyText), nil
}
