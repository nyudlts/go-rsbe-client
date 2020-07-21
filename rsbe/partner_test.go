package rsbe

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

func setupMux(apiPath string, filePath string) (mux *http.ServeMux) {
	mux = http.NewServeMux()
	mux.HandleFunc(apiPath, func(w http.ResponseWriter, _ *http.Request) {
		data, _ := ioutil.ReadFile(filePath)
		w.Write(data)
	})

	return mux
}

func setupTestServerClient(ts *httptest.Server) {
	c := new(Config)
	c.BaseURL = ts.URL
	c.User = "foo"
	c.Password = "bar"

	ConfigureClient(c)
}

func setupLocalhostClient() {
	c := new(Config)
	c.BaseURL = "http://localhost:3000"
	c.User = "foo"
	c.Password = "bar"

	ConfigureClient(c)
}

var partnerListEntry = PartnerListEntry{
	ID:        "e6517775-6277-4e25-9373-ee7738e820b5",
	Code:      "dlts",
	Name:      "nyu dlts",
	CreatedAt: "2020-05-30T01:56:01.603Z",
	UpdatedAt: "2020-05-30T01:56:01.603Z",
	URL:       "http://localhost:3000/api/v0/partners/e6517775-6277-4e25-9373-ee7738e820b5",
}

var partnerShow = PartnerEntry{
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

var partnerToCreate = PartnerEntry{
	Code:    "waffles",
	Name:    "Waffles and Syrup",
	RelPath: "content/waffles",
}

func TestPartnerList(t *testing.T) {

	setupLocalhostClient()
	t.Run("result", func(t *testing.T) {
		want := partnerListEntry
		got, err := PartnerList()
		if err != nil {
			t.Errorf("Unexpected error: %s", err)
		}

		if 1 != len(got) {
			t.Errorf("Mismatch: want: \"%v\", got: \"%v\"", want, got)
		}

		if partnerListEntry != got[0] {
			t.Errorf("Mismatch: want: \"%v\", got: \"%v\"", want, got)
		}
	})

}

func TestPartnerGetFunc(t *testing.T) {

	// mux := setupMux("/api/v0/partners/e6517775-6277-4e25-9373-ee7738e820b5", "testdata/partner-get.json")
	// ts := httptest.NewServer(mux)
	// defer ts.Close()

	// setupTestServerClient(ts)

	setupLocalhostClient()
	t.Run("result", func(t *testing.T) {
		want := partnerShow
		got := PartnerEntry{ID: "e6517775-6277-4e25-9373-ee7738e820b5"}

		err := got.Get()
		if err != nil {
			t.Errorf("Unexpected error: %s", err)
		}

		if got != want {
			t.Errorf("Mismatch: want: \"%v\", got: \"%v\"", want, got)
		}
	})

}

func TestPartnerGet(t *testing.T) {

	mux := setupMux("/api/v0/partners/e6517775-6277-4e25-9373-ee7738e820b5", "testdata/partner-get.json")
	ts := httptest.NewServer(mux)
	defer ts.Close()

	setupTestServerClient(ts)

	t.Run("confirm that expected partner was retrieved", func(t *testing.T) {
		want := partnerShow
		got, err := PartnerGet("e6517775-6277-4e25-9373-ee7738e820b5")
		if err != nil {
			t.Errorf("Unexpected error: %s", err)
		}

		if got != want {
			t.Errorf("Mismatch: want: \"%v\", got: \"%v\"", want, got)
		}
	})
}

func TestPartnerCreateFunc(t *testing.T) {
	setupLocalhostClient()

	err := partnerToCreate.Create()
	if err != nil {
		t.Errorf("Unexpected error: %s", err)
	}

	t.Run("confirm that attributes updated", func(t *testing.T) {
		if partnerToCreate.ID == "" {
			t.Errorf("ID not updated")
		}

		if partnerToCreate.CreatedAt == "" {
			t.Errorf("CreatedAt not updated")
		}

		if partnerToCreate.UpdatedAt == "" {
			t.Errorf("UpdatedAt not updated")
		}
	})
}

func TestPartnerUpdateFunc(t *testing.T) {
	setupLocalhostClient()

	_ = partnerToCreate.Get()

	if partnerToCreate.Name != "Waffles and Syrup" {
		t.Errorf("variable already updated: %s", partnerToCreate.ToString())
	}

	partnerToCreate.Name = "WAFFLES WAFFLES WAFFLES"

	err := partnerToCreate.Update()
	if err != nil {
		t.Errorf("Unexpected error: %s", err)
	}

	_ = partnerToCreate.Get()

	t.Run("confirm that elements updated", func(t *testing.T) {
		if partnerToCreate.Name != "WAFFLES WAFFLES WAFFLES" {
			t.Errorf("Name was not updated: got: %s", partnerToCreate.Name)
		}

		if partnerToCreate.CreatedAt == partnerToCreate.UpdatedAt {
			t.Errorf("UpeatedAt not updated")
		}
	})
}

func TestPartnerDeleteFunc(t *testing.T) {
	setupLocalhostClient()

	_ = partnerToCreate.Get()

	id := partnerToCreate.ID

	err := partnerToCreate.Delete()
	if err != nil {
		t.Errorf("Unexpected error: %s", err)
	}

	t.Run("confirm that deleted item not found", func(t *testing.T) {
		// should not be found, so err should NOT be nil
		_, err = PartnerGet(id)

		if err == nil {
			t.Errorf("err was nil")
		}

	})
}
