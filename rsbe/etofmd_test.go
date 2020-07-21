package rsbe

import (
	"net/http/httptest"
	"testing"
)

var etofmdListEntry = EtoFMDListEntry{
	ID:    "3ca8ecaf-6fae-48a5-8441-5a96e119ad28",
	EType: "se",
	EID:   "8c258cb2-d700-43be-8773-a61a7b9cd668",
	FMDID: "4a3f8f8c-6dbe-4d7c-bff1-1b973f9f615c",
	Role:  "master",
	URL:   "http://localhost:3000/api/v0/etofmds/3ca8ecaf-6fae-48a5-8441-5a96e119ad28",
}

var etofmdShow = EtoFMDEntry{
	ID:          "3ca8ecaf-6fae-48a5-8441-5a96e119ad28",
	EType:       "se",
	EID:         "8c258cb2-d700-43be-8773-a61a7b9cd668",
	Role:        "master",
	FMDID:       "4a3f8f8c-6dbe-4d7c-bff1-1b973f9f615c",
	CreatedAt:   "2020-05-31T20:37:30.747Z",
	UpdatedAt:   "2020-05-31T20:37:30.747Z",
	EURL:        "http://localhost:3000/api/v0/ses/8c258cb2-d700-43be-8773-a61a7b9cd668",
	LockVersion: 0,
}

var etofmdToCreate = EtoFMDEntry{
	EType:       "ie",
	EID:         "9ea98441-b6b6-46cf-b6c8-91dff385c6c8",
	Role:        "notes",
	FMDID:       "f9f38cc5-0728-4f1a-85ec-e4cb6906d304",
}

func TestEtoFMDList(t *testing.T) {

	mux := setupMux("/api/v0/etofmds", "testdata/etofmd-list.json")
	ts := httptest.NewServer(mux)
	defer ts.Close()

	setupTestServerClient(ts)

	t.Run("result", func(t *testing.T) {
		want := etofmdListEntry
		got, err := EtoFMDList()
		if err != nil {
			t.Errorf("Unexpected error: %s", err)
		}

		if 4 != len(got) {
			t.Errorf("Result Length Mismatch: want: 4, got: %d", len(got))
		}

		if etofmdListEntry != got[0] {
			t.Errorf("Mismatch: want: \"%v\", got: \"%v\"", want, got)
		}
	})

}

func TestEtoFMDGetFunc(t *testing.T) {

	mux := setupMux("/api/v0/etofmds/3ca8ecaf-6fae-48a5-8441-5a96e119ad28", "testdata/etofmd-get.json")
	ts := httptest.NewServer(mux)
	defer ts.Close()

	setupTestServerClient(ts)

	t.Run("result", func(t *testing.T) {
		want := etofmdShow
		got := EtoFMDEntry{ID: "3ca8ecaf-6fae-48a5-8441-5a96e119ad28"}

		err := got.Get()
		if err != nil {
			t.Errorf("Unexpected error: %s", err)
		}

		if want.ID != got.ID {
			t.Errorf("Mismatch: want: \"%v\", got: \"%v\"", want, got)
		}
	})

}

func TestEtoFMDCreateFunc(t *testing.T) {
	setupLocalhostClient()

	err := etofmdToCreate.Create()
	if err != nil {
		t.Errorf("Unexpected error: %s", err)
	}

	t.Run("confirm that attributes updated", func(t *testing.T) {
		if etofmdToCreate.ID == "" {
			t.Errorf("ID not updated")
		}

		if etofmdToCreate.CreatedAt == "" {
			t.Errorf("CreatedAt not updated")
		}

		if etofmdToCreate.UpdatedAt == "" {
			t.Errorf("UpdatedAt not updated")
		}
	})
}

func TestEtoFMDUpdateFunc(t *testing.T) {
	setupLocalhostClient()

	_ = etofmdToCreate.Get()

	if etofmdToCreate.Role != "notes" {
		t.Errorf("variable already updated: %s", etofmdToCreate.ToString())
	}

	etofmdToCreate.Role = "waffles"

	err := etofmdToCreate.Update()
	if err != nil {
		t.Errorf("Unexpected error: %s", err)
	}

	_ = etofmdToCreate.Get()

	t.Run("confirm that elements updated", func(t *testing.T) {
		if etofmdToCreate.Role != "waffles" {
			t.Errorf("Role was not updated: got: %s", etofmdToCreate.Role)
		}

		if etofmdToCreate.CreatedAt == etofmdToCreate.UpdatedAt {
			t.Errorf("UpdatedAt not updated")
		}
	})
}

func TestEtoFMDDeleteFunc(t *testing.T) {
	setupLocalhostClient()

	_ = etofmdToCreate.Get()

	id := etofmdToCreate.ID

	err := etofmdToCreate.Delete()
	if err != nil {
		t.Errorf("Unexpected error: %s", err)
	}

	t.Run("confirm that deleted item not found", func(t *testing.T) {
		// should not be found, so err should NOT be nil
		_, err = EtoFMDGet(id)

		if err == nil {
			t.Errorf("err was nil")
		}

	})
}
