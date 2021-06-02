package rsbe

import (
	"net/http/httptest"
	"testing"
)

var etofidListEntry = EToFIDListEntry{
	ID:       "5c4a0cf4-78c7-49cc-b227-9b017d85e65f",
	EType:    "se",
	EID:      "8c258cb2-d700-43be-8773-a61a7b9cd668",
	FIDType:  "handle",
	FIDValue: "2333.1/abc123",
	URL:      "http://localhost:3000/api/v0/etofids/5c4a0cf4-78c7-49cc-b227-9b017d85e65f",
}

var etofidShow = EToFIDEntry{
	ID:          "5c4a0cf4-78c7-49cc-b227-9b017d85e65f",
	EType:       "se",
	EID:         "8c258cb2-d700-43be-8773-a61a7b9cd668",
	FIDType:     "handle",
	FIDValue:    "2333.1/abc123",
	CreatedAt:   "2021-05-28T16:39:08.102Z",
	UpdatedAt:   "2021-05-28T16:39:08.102Z",
	EURL:        "http://localhost:3000/api/v0/ses/8c258cb2-d700-43be-8773-a61a7b9cd668",
	LockVersion: 0,
}

var etofidToCreate = EToFIDEntry{
	EType:    "ie",
	EID:      "9ea98441-b6b6-46cf-b6c8-91dff385c6c8",
	FIDType:  "noid",
	FIDValue: "7284f",
}

func TestEToFIDList(t *testing.T) {

	mux := setupMux("/api/v0/etofids", "testdata/etofid-list.json")
	ts := httptest.NewServer(mux)
	defer ts.Close()

	setupTestServerClient(ts)

	t.Run("result", func(t *testing.T) {
		want := etofidListEntry
		got, err := EToFIDList()
		if err != nil {
			t.Errorf("Unexpected error: %s", err)
		}

		if 1 != len(got) {
			t.Errorf("Result Length Mismatch: want: 1, got: %d", len(got))
		}

		if etofidListEntry != got[0] {
			t.Errorf("Mismatch: want: \"%v\", got: \"%v\"", want, got)
		}
	})

}

func TestEToFIDGetFunc(t *testing.T) {

	mux := setupMux("/api/v0/etofids/5c4a0cf4-78c7-49cc-b227-9b017d85e65f", "testdata/etofid-get.json")
	ts := httptest.NewServer(mux)
	defer ts.Close()

	setupTestServerClient(ts)

	t.Run("result", func(t *testing.T) {
		want := etofidShow
		got := EToFIDEntry{ID: "5c4a0cf4-78c7-49cc-b227-9b017d85e65f"}

		err := got.Get()
		if err != nil {
			t.Errorf("Unexpected error: %s", err)
		}

		if want.ID != got.ID {
			t.Errorf("Mismatch: want: \"%v\", got: \"%v\"", want.ID, got.ID)
		}

		if want.EType != got.EType {
			t.Errorf("Mismatch: want: \"%v\", got: \"%v\"", want.EType, got.EType)
		}
		if want.FIDType != got.FIDType {
			t.Errorf("Mismatch: want: \"%v\", got: \"%v\"", want.FIDType, got.FIDType)
		}

		if want.FIDValue != got.FIDValue {
			t.Errorf("Mismatch: want: \"%v\", got: \"%v\"", want.FIDValue, got.FIDValue)
		}

	})

}

func TestEToFIDCreateFunc(t *testing.T) {
	setupLocalhostClient()

	err := etofidToCreate.Create()
	if err != nil {
		t.Errorf("Unexpected error: %s", err)
	}

	t.Run("confirm that attributes updated", func(t *testing.T) {
		if etofidToCreate.ID == "" {
			t.Errorf("ID not updated")
		}

		if etofidToCreate.CreatedAt == "" {
			t.Errorf("CreatedAt not updated")
		}

		if etofidToCreate.UpdatedAt == "" {
			t.Errorf("UpdatedAt not updated")
		}
	})
}

func TestEToFIDUpdateFunc(t *testing.T) {
	setupLocalhostClient()

	_ = etofidToCreate.Get()

	if etofidToCreate.FIDValue != "7284f" {
		t.Errorf("variable already updated: %s", etofidToCreate.ToString())
	}

	etofidToCreate.FIDValue = "waffles"

	err := etofidToCreate.Update()
	if err != nil {
		t.Errorf("Unexpected error: %s", err)
	}

	_ = etofidToCreate.Get()

	t.Run("confirm that elements updated", func(t *testing.T) {
		if etofidToCreate.FIDValue != "waffles" {
			t.Errorf("FIDValue was not updated: got: %s", etofidToCreate.FIDValue)
		}

		if etofidToCreate.CreatedAt == etofidToCreate.UpdatedAt {
			t.Errorf("UpdatedAt not updated")
		}
	})
}

func TestEToFIDDeleteFunc(t *testing.T) {
	setupLocalhostClient()

	_ = etofidToCreate.Get()

	id := etofidToCreate.ID

	err := etofidToCreate.Delete()
	if err != nil {
		t.Errorf("Unexpected error: %s", err)
	}

	t.Run("confirm that deleted item not found", func(t *testing.T) {
		// should not be found, so err should NOT be nil
		_, err = EToFIDGet(id)

		if err == nil {
			t.Errorf("err was nil")
		}

	})
}
