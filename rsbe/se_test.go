package rsbe

import (
	"net/http/httptest"
	"testing"
)

var seListEntry = SEListEntry{
	ID:            "8c258cb2-d700-43be-8773-a61a7b9cd668",
	DigiID:        "foo",
	DOType:        "video",
	Phase:         "digitization",
	Step:          "digitization",
	Status:        "active",
	CreatedAt:     "2020-05-30T02:07:17.846Z",
	UpdatedAt:     "2020-05-30T02:07:17.846Z",
	URL:           "http://localhost:3000/api/v0/ses/8c258cb2-d700-43be-8773-a61a7b9cd668",
	CollectionURL: "http://localhost:3000/api/v0/colls/b9612d5d-619a-4ceb-b620-d816e4b4340b",
}

var seShow = SEEntry{
	ID:            "8c258cb2-d700-43be-8773-a61a7b9cd668",
	CollectionID:  "b9612d5d-619a-4ceb-b620-d816e4b4340b",
	DigiID:        "foo",
	DOType:        "video",
	Phase:         "digitization",
	Step:          "digitization",
	Status:        "active",
	CreatedAt:     "2020-05-30T02:07:17.846Z",
	UpdatedAt:     "2020-05-30T02:07:17.846Z",
	FMDsURL:       "http://localhost:3000/api/v0/ses/8c258cb2-d700-43be-8773-a61a7b9cd668/fmds",
	CollectionURL: "http://localhost:3000/api/v0/colls/b9612d5d-619a-4ceb-b620-d816e4b4340b",
	LockVersion:   0,
}

var seToCreate = SEEntry{
	CollectionID: seShow.CollectionID,
	DigiID:       "temporary_item",
	DOType:       "video",
	Phase:        "digitization",
	Step:         "digitization",
	Status:       "queued",
}

func TestCollectionSEList(t *testing.T) {

	mux := setupMux("/api/v0/colls/b9612d5d-619a-4ceb-b620-d816e4b4340b/ses", "testdata/se-list.json")
	ts := httptest.NewServer(mux)
	defer ts.Close()

	setupTestServerClient(ts)

	t.Run("result", func(t *testing.T) {
		want := seListEntry
		got, err := CollectionSEList(seShow.CollectionID)
		if err != nil {
			t.Errorf("Unexpected error: %s", err)
		}

		if 1 != len(got) {
			t.Errorf("Mismatch: want: \"%v\", got: \"%v\"", want, got)
		}

		if seListEntry != got[0] {
			t.Errorf("Mismatch: want: \"%v\", got: \"%v\"", want, got)
		}
	})

}

func TestSEGetFunc(t *testing.T) {

	mux := setupMux("/api/v0/ses/8c258cb2-d700-43be-8773-a61a7b9cd668", "testdata/se-get.json")
	ts := httptest.NewServer(mux)
	defer ts.Close()

	setupTestServerClient(ts)

	t.Run("result", func(t *testing.T) {
		want := seShow
		got := SEEntry{ID: "8c258cb2-d700-43be-8773-a61a7b9cd668"}

		err := got.Get()
		if err != nil {
			t.Errorf("Unexpected error: %s", err)
		}

		if got != want {
			t.Errorf("Mismatch: want: \"%v\", got: \"%v\"", want, got)
		}
	})

}

func TestSEGet(t *testing.T) {

	mux := setupMux("/api/v0/ses/8c258cb2-d700-43be-8773-a61a7b9cd668", "testdata/se-get.json")
	ts := httptest.NewServer(mux)
	defer ts.Close()

	setupTestServerClient(ts)

	t.Run("confirm that expected partner was retrieved", func(t *testing.T) {
		want := seShow
		got, err := SEGet("8c258cb2-d700-43be-8773-a61a7b9cd668")
		if err != nil {
			t.Errorf("Unexpected error: %s", err)
		}

		if got != want {
			t.Errorf("Mismatch: want: \"%v\", got: \"%v\"", want, got)
		}
	})
}

func TestSECreateFunc(t *testing.T) {
	setupLocalhostClient()

	err := seToCreate.Create()
	if err != nil {
		t.Errorf("Unexpected error: %s", err)
	}

	t.Run("confirm that attributes updated", func(t *testing.T) {
		if seToCreate.ID == "" {
			t.Errorf("ID not updated")
		}

		if seToCreate.CreatedAt == "" {
			t.Errorf("CreatedAt not updated")
		}

		if seToCreate.UpdatedAt == "" {
			t.Errorf("UpdatedAt not updated")
		}
	})
}

func TestSEUpdateFunc(t *testing.T) {
	setupLocalhostClient()

	_ = seToCreate.Get()

	if seToCreate.DigiID != "temporary_item" {
		t.Errorf("variable already updated: %s", seToCreate.ToString())
	}

	seToCreate.DigiID = "DogBiscuit"

	err := seToCreate.Update()
	if err != nil {
		t.Errorf("Unexpected error: %s", err)
	}

	_ = seToCreate.Get()

	t.Run("confirm that elements updated", func(t *testing.T) {
		if seToCreate.DigiID != "DogBiscuit" {
			t.Errorf("DigiID was not updated: got: %s", seToCreate.DigiID)
		}

		if seToCreate.CreatedAt == seToCreate.UpdatedAt {
			t.Errorf("UpeatedAt not updated")
		}
	})
}

func TestSEDeleteFunc(t *testing.T) {
	setupLocalhostClient()

	_ = seToCreate.Get()

	id := seToCreate.ID

	err := seToCreate.Delete()
	if err != nil {
		t.Errorf("Unexpected error: %s", err)
	}

	t.Run("confirm that deleted item not found", func(t *testing.T) {
		// should not be found, so err should NOT be nil
		_, err = SEGet(id)

		if err != nil {
			t.Errorf("err was nil")
		}

	})
}
