package rsbe

import (
	"net/http/httptest"
	"testing"

	"github.com/nyudlts/go-rsbe-client/rsbe/internal/testutils"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
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

		if len(got) != 1 {
			t.Errorf("Mismatch: want: \"%v\", got: \"%v\"", want, got)
		}

		assert.Equal(t, want.ID, got[0].ID)
		assert.Equal(t, want.CollectionID, got[0].CollectionID)
		assert.Equal(t, want.SysNum, got[0].SysNum)
		assert.Equal(t, want.Phase, got[0].Phase)
		assert.Equal(t, want.Step, got[0].Step)
		assert.Equal(t, want.Status, got[0].Status)
		assert.Equal(t, want.Title, got[0].Title)
		testutils.AssertEquivalentTimestamps(t, want.CreatedAt, got[0].CreatedAt)
		testutils.AssertEquivalentTimestamps(t, want.UpdatedAt, got[0].UpdatedAt)
		assert.Contains(t, got[0].URL, "/api/v0/ies/9ea98441-b6b6-46cf-b6c8-91dff385c6c8")
		assert.Contains(t, got[0].CollectionURL, "/api/v0/colls/b9612d5d-619a-4ceb-b620-d816e4b4340b")
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

		require.NoError(t, got.Get(), "IE Get failed")

		assert.Equal(t, want.ID, got.ID)
		assert.Equal(t, want.CollectionID, got.CollectionID)
		assert.Equal(t, want.SysNum, got.SysNum)
		assert.Equal(t, want.Phase, got.Phase)
		assert.Equal(t, want.Step, got.Step)
		assert.Equal(t, want.Status, got.Status)
		assert.Equal(t, want.Title, got.Title)
		assert.Equal(t, want.Notes, got.Notes)
		testutils.AssertEquivalentTimestamps(t, want.CreatedAt, got.CreatedAt)
		testutils.AssertEquivalentTimestamps(t, want.UpdatedAt, got.UpdatedAt)
		assert.Contains(t, got.FMDsURL, "/api/v0/ies/9ea98441-b6b6-46cf-b6c8-91dff385c6c8/fmds")
		assert.Contains(t, got.CollectionURL, "/api/v0/colls/b9612d5d-619a-4ceb-b620-d816e4b4340b")
		assert.Equal(t, want.LockVersion, got.LockVersion)
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
		require.NoError(t, err, "IE Get failed")

		assert.Equal(t, want.ID, got.ID)
		assert.Equal(t, want.CollectionID, got.CollectionID)
		assert.Equal(t, want.SysNum, got.SysNum)
		assert.Equal(t, want.Phase, got.Phase)
		assert.Equal(t, want.Step, got.Step)
		assert.Equal(t, want.Status, got.Status)
		assert.Equal(t, want.Title, got.Title)
		assert.Equal(t, want.Notes, got.Notes)
		testutils.AssertEquivalentTimestamps(t, want.CreatedAt, got.CreatedAt)
		testutils.AssertEquivalentTimestamps(t, want.UpdatedAt, got.UpdatedAt)
		assert.Contains(t, got.FMDsURL, "/api/v0/ies/9ea98441-b6b6-46cf-b6c8-91dff385c6c8/fmds")
		assert.Contains(t, got.CollectionURL, "/api/v0/colls/b9612d5d-619a-4ceb-b620-d816e4b4340b")
		assert.Equal(t, want.LockVersion, got.LockVersion)
	})
}

func TestIECreateFunc(t *testing.T) {
	setupLocalhostClient()

	err := ieToCreate.Create()
	require.NoError(t, err, "IE Create failed")
	defer ieToCreate.Purge()

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

	err := ieToCreate.Create()
	require.NoError(t, err, "IE Create failed")
	defer ieToCreate.Purge()

	_ = ieToCreate.Get()

	if ieToCreate.SysNum != "b123888" {
		t.Errorf("variable already updated: %s", ieToCreate.ToString())
	}

	ieToCreate.SysNum = "x9988771"
	ieToCreate.Title = "Hop on Pop!"

	err = ieToCreate.Update()
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

	err := ieToCreate.Create()
	require.NoError(t, err, "IE Create failed")
	// the Purge() call is needed because Delete() is a soft delete,
	// the record will still exist and needs to be purged
	defer ieToCreate.Purge()

	_ = ieToCreate.Get()

	id := ieToCreate.ID

	err = ieToCreate.Delete()
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
