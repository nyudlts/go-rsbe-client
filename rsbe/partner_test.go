package rsbe

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGet(t *testing.T) {

	mux := http.NewServeMux()
	mux.HandleFunc("/api/v0/partners/e6517775-6277-4e25-9373-ee7738e820b5", func(w http.ResponseWriter, _ *http.Request) {
		data, _ := ioutil.ReadFile("testdata/partner-get.json")
		w.Write(data)
	})

	ts := httptest.NewServer(mux)
	defer ts.Close()

	c := new(Config)
	c.BaseURL = ts.URL
	c.User = "foo"
	c.Password = "bar"

	ConfigureClient(c)

	t.Run("result", func(t *testing.T) {
		want := PartnerEntry{
			ID:             "e6517775-6277-4e25-9373-ee7738e820b5",
			Code:           "dlts",
			Name:           "nyu dlts",
			CreatedAt:      "2020-05-30T01:56:01.603Z",
			UpdatedAt:      "2020-05-30T01:56:01.603Z",
			PartnersURL:    "http://localhost:3000/api/v0/partners",
			CollectionsURL: "http://localhost:3000/api/v0/partners/e6517775-6277-4e25-9373-ee7738e820b5/colls",
			LockVersion:    0,
			RelPath:        "content/dlts",
		}
		got := PartnerEntry{ID: "e6517775-6277-4e25-9373-ee7738e820b5"}

		err := got.Get()
		if err != nil {
			t.Errorf("Unexpected error: %s", err)
		}

		if got != want {
			t.Errorf("Expected CollectionId to == \"%v\", but got \"%v\"", want, got)
		}
	})

}

func TestPartnerGet(t *testing.T) {

	mux := http.NewServeMux()
	mux.HandleFunc("/api/v0/partners/e6517775-6277-4e25-9373-ee7738e820b5", func(w http.ResponseWriter, _ *http.Request) {
		data, _ := ioutil.ReadFile("testdata/partner-get.json")
		w.Write(data)
	})

	ts := httptest.NewServer(mux)
	defer ts.Close()

	c := new(Config)
	c.BaseURL = ts.URL
	c.User = "foo"
	c.Password = "bar"

	ConfigureClient(c)

	t.Run("result", func(t *testing.T) {
		want := PartnerEntry{
			ID:             "e6517775-6277-4e25-9373-ee7738e820b5",
			Code:           "dlts",
			Name:           "nyu dlts",
			CreatedAt:      "2020-05-30T01:56:01.603Z",
			UpdatedAt:      "2020-05-30T01:56:01.603Z",
			PartnersURL:    "http://localhost:3000/api/v0/partners",
			CollectionsURL: "http://localhost:3000/api/v0/partners/e6517775-6277-4e25-9373-ee7738e820b5/colls",
			LockVersion:    0,
			RelPath:        "content/dlts",
		}
		got, err := PartnerGet("e6517775-6277-4e25-9373-ee7738e820b5")
		if err != nil {
			t.Errorf("Unexpected error: %s", err)
		}

		if got != want {
			t.Errorf("Expected CollectionId to == \"%v\", but got \"%v\"", want, got)
		}
	})

}
