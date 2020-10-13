package rsbe

import (
	"net/http/httptest"
	"testing"
)

var ieListEntry = IEListEntry{
	ID:            "9ea98441-b6b6-46cf-b6c8-91dff385c6c8",
	SysNum:        "123456",
	Phase:         "registration",
	Step:          "loading",
	Status:        "done",
	Title:         "The White Whale",
	CreatedAt:     "2020-05-31T20:57:58.618Z",
	UpdatedAt:     "2020-05-31T20:57:58.618Z",
	URL:           "http://localhost:3000/api/v0/ies/9ea98441-b6b6-46cf-b6c8-91dff385c6c8",
	CollectionURL: "http://localhost:3000/api/v0/colls/b9612d5d-619a-4ceb-b620-d816e4b4340b",
}

var ieShow = IEEntry{
	ID:            "9ea98441-b6b6-46cf-b6c8-91dff385c6c8",
	CollectionID:  "b9612d5d-619a-4ceb-b620-d816e4b4340b",
	SysNum:        "123456",
	Phase:         "registration",
	Step:          "loading",
	Status:        "done",
	Title:         "The White Whale",
	Notes:         "glorious notes",
	CreatedAt:     "2020-05-31T20:57:58.618Z",
	UpdatedAt:     "2020-05-31T20:57:58.618Z",
	FMDsURL:       "http://localhost:3000/api/v0/ies/9ea98441-b6b6-46cf-b6c8-91dff385c6c8/fmds",
	CollectionURL: "http://localhost:3000/api/v0/colls/b9612d5d-619a-4ceb-b620-d816e4b4340b",
	LockVersion:   0,
}

var ieToCreate = IEEntry{
	CollectionID: ieShow.CollectionID,
	SysNum:       "b123888",
	Phase:        "registration",
	Step:         "loading",
	Status:       "done",
	Title:        "One Fish, Two Fish, Red Fish, Blue Fish",
}

func TestCollectionIEList(t *testing.T) {
	mux := setupMux("/api/v0/colls/b9612d5d-619a-4ceb-b620-d816e4b4340b/ies", "testdata/ie-list.json")
	ts := httptest.NewServer(mux)
	defer ts.Close()

	setupTestServerClient(ts)

	t.Run("result", func(t *testing.T) {
		want := ieListEntry
		got, err := CollectionIEList(ieShow.CollectionID)
		if err != nil {
			t.Errorf("Unexpected error: %s", err)
		}

		if 1 != len(got) {
			t.Errorf("Mismatch: want: \"%v\", got: \"%v\"", want, got)
		}

		if ieListEntry != got[0] {
			t.Errorf("Mismatch: want: \"%v\", got: \"%v\"", want, got)
		}
	})

}

// test that struct method Get works
func TestIEGetFunc(t *testing.T) {

	mux := setupMux("/api/v0/ies/9ea98441-b6b6-46cf-b6c8-91dff385c6c8", "testdata/ie-get.json")
	ts := httptest.NewServer(mux)
	defer ts.Close()

	setupTestServerClient(ts)

	t.Run("result", func(t *testing.T) {
		want := ieShow
		got := IEEntry{ID: "9ea98441-b6b6-46cf-b6c8-91dff385c6c8"}

		err := got.Get()
		if err != nil {
			t.Errorf("Unexpected error: %s", err)
		}

		if got != want {
			t.Errorf("Mismatch: want:\n\"%v\", got:\n\"%v\"", want, got)
		}
	})

}

// test that model-level Get works
func TestIEGet(t *testing.T) {

	mux := setupMux("/api/v0/ies/9ea98441-b6b6-46cf-b6c8-91dff385c6c8", "testdata/ie-get.json")
	ts := httptest.NewServer(mux)
	defer ts.Close()

	setupTestServerClient(ts)

	t.Run("confirm that expected resource was retrieved", func(t *testing.T) {
		want := ieShow
		got, err := IEGet("9ea98441-b6b6-46cf-b6c8-91dff385c6c8")
		if err != nil {
			t.Errorf("Unexpected error: %s", err)
		}

		if got != want {
			t.Errorf("Mismatch: want: \"%v\", got: \"%v\"", want, got)
		}
	})
}

func TestIECreateFunc(t *testing.T) {
	setupLocalhostClient()

	err := ieToCreate.Create()
	if err != nil {
		t.Errorf("Unexpected error: %s", err)
	}

	t.Run("confirm that attributes updated", func(t *testing.T) {
		if ieToCreate.ID == "" {
			t.Errorf("ID not updated")
		}

		if ieToCreate.CreatedAt == "" {
			t.Errorf("CreatedAt not updated")
		}

		if ieToCreate.UpdatedAt == "" {
			t.Errorf("UpdatedAt not updated")
		}
	})
}

func TestIEUpdateFunc(t *testing.T) {
	setupLocalhostClient()

	_ = ieToCreate.Get()

	if ieToCreate.SysNum != "b123888" {
		t.Errorf("variable already updated: %s", ieToCreate.ToString())
	}

	ieToCreate.SysNum = "x9988771"
	ieToCreate.Title = "Hop on Pop!"

	err := ieToCreate.Update()
	if err != nil {
		t.Errorf("Unexpected error: %s", err)
	}

	_ = ieToCreate.Get()

	t.Run("confirm that elements updated", func(t *testing.T) {
		if ieToCreate.SysNum != "x9988771" {
			t.Errorf("SysNum was not updated: got: %s", ieToCreate.SysNum)
		}

		if ieToCreate.Title != "Hop on Pop!" {
			t.Errorf("Title was not updated: got: %s", ieToCreate.Title)
		}

		if ieToCreate.CreatedAt == ieToCreate.UpdatedAt {
			t.Errorf("UpdatedAt not updated")
		}
	})
}

func TestIEDeleteFunc(t *testing.T) {
	setupLocalhostClient()

	_ = ieToCreate.Get()

	id := ieToCreate.ID

	err := ieToCreate.Delete()
	if err != nil {
		t.Errorf("Unexpected error: %s", err)
	}

	t.Run("confirm that deleted item not found", func(t *testing.T) {
		// should not be found, so err should NOT be nil
		_, err = IEGet(id)

		if err == nil {
			t.Errorf("err was nil")
		}

	})
}
