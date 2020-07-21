package rsbe

import (
	"encoding/json"
	"path"
	"testing"
)

func TestClientGet(t *testing.T) {

	setupLocalhostClient()
	t.Run("confirm OK status response", func(t *testing.T) {
		want := 200
		resp, err := Get("/api/v0/partners/e6517775-6277-4e25-9373-ee7738e820b5")
		if err != nil {
			t.Errorf("Unexpected error: %s", err)
		}

		got := resp.StatusCode
		if want != got {
			t.Errorf("Mismatch: want: \"%v\", got: \"%v\"", want, got)
		}
	})

	t.Run("confirm bad request response", func(t *testing.T) {
		want := 400
		resp, err := Get("/api/v0/partners/x")
		if err == nil {
			t.Errorf("err should NOT be nil: %v", err)
		}

		got := resp.StatusCode
		if want != got {
			t.Errorf("Mismatch: want: \"%v\", got: \"%v\"", want, got)
		}
	})

}

func TestClientGetBody(t *testing.T) {

	setupLocalhostClient()
	t.Run("confirm OK status response", func(t *testing.T) {
		body, err := GetBody("/api/v0/partners/e6517775-6277-4e25-9373-ee7738e820b5")
		if err != nil {
			t.Errorf("Unexpected error: %s", err)
		}

		if len(body) == 0 {
			t.Errorf("Body should not be empty")
		}
	})

	t.Run("confirm bad request response", func(t *testing.T) {
		body, err := GetBody("/api/v0/partners/x")
		if err == nil {
			t.Errorf("err should NOT be nil: %v", err)
		}

		if len(body) != 0 {
			t.Errorf("Body should be empty")
		}
	})

}

func TestClientPost(t *testing.T) {

	data := "{\"code\":\"canteloupe\",\"name\":\"Can elope\",\"rel_path\":\"content/canteloupe\"}"
	ppath := "/api/v0/partners"

	setupLocalhostClient()
	t.Run("confirm OK status response on successful POST", func(t *testing.T) {

		want := 201
		resp, err := Post(ppath, []byte(data))
		if err != nil {
			t.Errorf("Unexpected error: %s", err)
		}

		got := resp.StatusCode
		if want != got {
			t.Errorf("Mismatch: want: \"%v\", got: \"%v\"", want, got)
		}

		id := path.Base(resp.Header.Get("Location"))
		if id == "." {
			t.Errorf("Unable to find created partner to delete.")
		}
		err = PartnerDelete(id)
		if err != nil {
			t.Errorf("Unexpected error: %s", err)
		}
	})

	t.Run("confirm bad request response", func(t *testing.T) {

		_, err := Post("/api/v0/partners", []byte("{}"))
		if err == nil {
			t.Errorf("err should NOT be nil: %v", err)
		}
	})

}

func TestClientPostReturnBody(t *testing.T) {

	data := "{\"code\":\"canteloupe\",\"name\":\"Can elope\",\"rel_path\":\"content/canteloupe\"}"
	ppath := "/api/v0/partners"

	setupLocalhostClient()
	t.Run("confirm OK status response on successful POST", func(t *testing.T) {

		body, err := PostReturnBody(ppath, []byte(data))
		if err != nil {
			t.Errorf("Unexpected error: %s", err)
		}

		if err != nil {
			t.Errorf("err should be nil: %s", err)
		}

		if len(body) == 0 {
			t.Errorf("Body should not be empty")
		}

		var p PartnerEntry
		err = json.Unmarshal(body, &p)
		if err != nil {
			t.Errorf("Error parsing body.")
		}

		err = PartnerDelete(p.ID)
		if err != nil {
			t.Errorf("Unexpected error: %s", err)
		}
	})

	t.Run("confirm bad request response", func(t *testing.T) {

		_, err := PostReturnBody("/api/v0/partners", []byte("{}"))
		if err == nil {
			t.Errorf("err should NOT be nil: %v", err)
		}
	})

}
