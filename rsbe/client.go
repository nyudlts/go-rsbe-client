package rsbe

import (
	"bytes"
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

// TODO: change fn name? GetBodyText?
// TODO: do you need intermediate bodyText variable, or
//       or can you just use variable "s"?
//       add status code check
//       also might want to just merge all of this into Get?
func GetBody(path string) (body []byte, err error) {

	resp, err := Get(path)
	if err != nil {
		return body, err
	}
	defer resp.Body.Close()

	body, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		return body, err
	}

	return body, nil
}

func Post(path string, data []byte) (resp *http.Response, err error) {
	url := fmt.Sprintf("%s%s", conf.BaseURL, path)

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(data))
	if err != nil {
		return resp, err
	}

	req.Header.Set("Content-Type", "application/json")
	req.SetBasicAuth(conf.User, conf.Password)

	return client.Do(req)
}

func PostReturnBody(path string, data []byte) (body []byte, err error) {
	resp, err := Post(path, data)
	if err != nil {
		return body, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 201 {
		return body, fmt.Errorf("unexpected status code: %s", resp.Status)
	}

	body, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		return body, err
	}

	return body, nil
}

func Put(path string, data []byte) (err error) {
	url := fmt.Sprintf("%s%s", conf.BaseURL, path)

	req, err := http.NewRequest("PUT", url, bytes.NewBuffer(data))
	if err != nil {
		return err
	}

	req.Header.Set("Content-Type", "application/json")
	req.SetBasicAuth(conf.User, conf.Password)

	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 204 {
		return fmt.Errorf("unexpected status code: %s", resp.Status)
	}

	return nil
}

func Delete(path string) (err error) {
	url := fmt.Sprintf("%s%s", conf.BaseURL, path)

	req, err := http.NewRequest("DELETE", url, nil)
	if err != nil {
		return err
	}
	req.SetBasicAuth(conf.User, conf.Password)

	resp, err := client.Do(req)
	if err != nil {
		return err
	}

	if resp.StatusCode != 204 {
		return fmt.Errorf("unexpected status code: %s", resp.Status)
	}

	return nil
}
