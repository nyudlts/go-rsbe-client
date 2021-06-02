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

// var etofidShow = EToFIDEntry{
// 	ID:          "3ca8ecaf-6fae-48a5-8441-5a96e119ad28",
// 	EType:       "se",
// 	EID:         "8c258cb2-d700-43be-8773-a61a7b9cd668",
// 	Role:        "master",
// 	FMDID:       "4a3f8f8c-6dbe-4d7c-bff1-1b973f9f615c",
// 	CreatedAt:   "2020-05-31T20:37:30.747Z",
// 	UpdatedAt:   "2020-05-31T20:37:30.747Z",
// 	EURL:        "http://localhost:3000/api/v0/ses/8c258cb2-d700-43be-8773-a61a7b9cd668",
// 	LockVersion: 0,
// }

// var etofidToCreate = EToFIDEntry{
// 	EType: "ie",
// 	EID:   "9ea98441-b6b6-46cf-b6c8-91dff385c6c8",
// 	Role:  "notes",
// 	FMDID: "f9f38cc5-0728-4f1a-85ec-e4cb6906d304",
// }

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

// func TestEToFIDGetFunc(t *testing.T) {

// 	mux := setupMux("/api/v0/etofids/3ca8ecaf-6fae-48a5-8441-5a96e119ad28", "testdata/etofid-get.json")
// 	ts := httptest.NewServer(mux)
// 	defer ts.Close()

// 	setupTestServerClient(ts)

// 	t.Run("result", func(t *testing.T) {
// 		want := etofidShow
// 		got := EToFIDEntry{ID: "3ca8ecaf-6fae-48a5-8441-5a96e119ad28"}

// 		err := got.Get()
// 		if err != nil {
// 			t.Errorf("Unexpected error: %s", err)
// 		}

// 		if want.ID != got.ID {
// 			t.Errorf("Mismatch: want: \"%v\", got: \"%v\"", want, got)
// 		}
// 	})

// }

// func TestEToFIDCreateFunc(t *testing.T) {
// 	setupLocalhostClient()

// 	err := etofidToCreate.Create()
// 	if err != nil {
// 		t.Errorf("Unexpected error: %s", err)
// 	}

// 	t.Run("confirm that attributes updated", func(t *testing.T) {
// 		if etofidToCreate.ID == "" {
// 			t.Errorf("ID not updated")
// 		}

// 		if etofidToCreate.CreatedAt == "" {
// 			t.Errorf("CreatedAt not updated")
// 		}

// 		if etofidToCreate.UpdatedAt == "" {
// 			t.Errorf("UpdatedAt not updated")
// 		}
// 	})
// }

// func TestEToFIDUpdateFunc(t *testing.T) {
// 	setupLocalhostClient()

// 	_ = etofidToCreate.Get()

// 	if etofidToCreate.Role != "notes" {
// 		t.Errorf("variable already updated: %s", etofidToCreate.ToString())
// 	}

// 	etofidToCreate.Role = "waffles"

// 	err := etofidToCreate.Update()
// 	if err != nil {
// 		t.Errorf("Unexpected error: %s", err)
// 	}

// 	_ = etofidToCreate.Get()

// 	t.Run("confirm that elements updated", func(t *testing.T) {
// 		if etofidToCreate.Role != "waffles" {
// 			t.Errorf("Role was not updated: got: %s", etofidToCreate.Role)
// 		}

// 		if etofidToCreate.CreatedAt == etofidToCreate.UpdatedAt {
// 			t.Errorf("UpdatedAt not updated")
// 		}
// 	})
// }

// func TestEToFIDDeleteFunc(t *testing.T) {
// 	setupLocalhostClient()

// 	_ = etofidToCreate.Get()

// 	id := etofidToCreate.ID

// 	err := etofidToCreate.Delete()
// 	if err != nil {
// 		t.Errorf("Unexpected error: %s", err)
// 	}

// 	t.Run("confirm that deleted item not found", func(t *testing.T) {
// 		// should not be found, so err should NOT be nil
// 		_, err = EToFIDGet(id)

// 		if err == nil {
// 			t.Errorf("err was nil")
// 		}

// 	})
// }
