package rsbe

import (
	"net/http/httptest"
	"testing"

	"github.com/nyudlts/go-rsbe-client/rsbe/internal/testutils"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
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

		if len(got) != 1 {
			t.Errorf("Mismatch: want: \"%v\", got: \"%v\"", want, got)
		}

		assert.Equal(t, want.ID, got[0].ID)
		assert.Equal(t, want.DigiID, got[0].DigiID)
		assert.Equal(t, want.DOType, got[0].DOType)
		assert.Equal(t, want.Phase, got[0].Phase)
		assert.Equal(t, want.Step, got[0].Step)
		assert.Equal(t, want.Status, got[0].Status)
		assert.Equal(t, want.Label, got[0].Label)
		testutils.AssertEquivalentTimestamps(t, want.CreatedAt, got[0].CreatedAt)
		testutils.AssertEquivalentTimestamps(t, want.UpdatedAt, got[0].UpdatedAt)
		assert.Contains(t, got[0].URL, "/api/v0/ses/8c258cb2-d700-43be-8773-a61a7b9cd668")
		assert.Contains(t, got[0].CollectionURL, "/api/v0/colls/b9612d5d-619a-4ceb-b620-d816e4b4340b")
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

		require.NoError(t, got.Get(), "SE Get failed")

		assert.Equal(t, want.ID, got.ID)
		assert.Equal(t, want.CollectionID, got.CollectionID)
		assert.Equal(t, want.DigiID, got.DigiID)
		assert.Equal(t, want.DOType, got.DOType)
		assert.Equal(t, want.Phase, got.Phase)
		assert.Equal(t, want.Step, got.Step)
		assert.Equal(t, want.Status, got.Status)
		testutils.AssertEquivalentTimestamps(t, want.CreatedAt, got.CreatedAt)
		testutils.AssertEquivalentTimestamps(t, want.UpdatedAt, got.UpdatedAt)
		assert.Contains(t, got.FMDsURL, "/api/v0/ses/8c258cb2-d700-43be-8773-a61a7b9cd668/fmds")
		assert.Contains(t, got.CollectionURL, "/api/v0/colls/b9612d5d-619a-4ceb-b620-d816e4b4340b")
		assert.Equal(t, want.LockVersion, got.LockVersion)
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
		require.NoError(t, err, "SE Get failed")

		assert.Equal(t, want.ID, got.ID)
		assert.Equal(t, want.CollectionID, got.CollectionID)
		assert.Equal(t, want.DigiID, got.DigiID)
		assert.Equal(t, want.DOType, got.DOType)
		assert.Equal(t, want.Phase, got.Phase)
		assert.Equal(t, want.Step, got.Step)
		assert.Equal(t, want.Status, got.Status)
		testutils.AssertEquivalentTimestamps(t, want.CreatedAt, got.CreatedAt)
		testutils.AssertEquivalentTimestamps(t, want.UpdatedAt, got.UpdatedAt)
		assert.Contains(t, got.FMDsURL, "/api/v0/ses/8c258cb2-d700-43be-8773-a61a7b9cd668/fmds")
		assert.Contains(t, got.CollectionURL, "/api/v0/colls/b9612d5d-619a-4ceb-b620-d816e4b4340b")
		assert.Equal(t, want.LockVersion, got.LockVersion)
	})
}

func TestSECreateFunc(t *testing.T) {
	setupLocalhostClient()

	err := seToCreate.Create()
	require.NoError(t, err, "SE Create failed")
	defer seToCreate.Purge()

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

	err := seToCreate.Create()
	require.NoError(t, err, "SE Create failed")
	defer seToCreate.Purge()

	_ = seToCreate.Get()

	if seToCreate.DigiID != "temporary_item" {
		t.Errorf("variable already updated: %s", seToCreate.ToString())
	}

	seToCreate.DigiID = "DogBiscuit"

	err = seToCreate.Update()
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

	err := seToCreate.Create()
	require.NoError(t, err, "SE Create failed")
	// the Purge() call is needed because Delete() is a soft delete,
	// the record will still exist and needs to be purged
	defer seToCreate.Purge()

	_ = seToCreate.Get()

	id := seToCreate.ID

	err = seToCreate.Delete()
	if err != nil {
		t.Errorf("Unexpected error: %s", err)
	}

	t.Run("confirm that deleted item not found", func(t *testing.T) {
		// should not be found, so err should NOT be nil
		_, err = SEGet(id)

		if err == nil {
			t.Errorf("err was nil")
		}

	})
}

func TestSEGetByDigiID(t *testing.T) {
	setupLocalhostClient()

	want, err := SEGet("8c258cb2-d700-43be-8773-a61a7b9cd668")
	if err != nil {
		t.Errorf("Unexpected error: %s", err)
	}

	t.Run("confirm that exiting SE is returned", func(t *testing.T) {
		got, err := SEGetByDigiID("foo")
		if err != nil {
			t.Errorf("Unexpected error: %s", err)
		}

		if got != want {
			t.Errorf("Mismatch: want: \"%v\", got: \"%v\"", want, got)
		}
	})

	t.Run("confirm error is returned when not found", func(t *testing.T) {
		_, err := SEGetByDigiID("fooX")
		if err == nil {
			t.Errorf("Expected search for non-existant SE to return error, but err was nil")
		}
	})

}

func TestGetByDigiIDFunc(t *testing.T) {
	setupLocalhostClient()

	want, err := SEGet("8c258cb2-d700-43be-8773-a61a7b9cd668")
	if err != nil {
		t.Errorf("Unexpected error: %s", err)
	}

	t.Run("confirm that exiting SE is returned", func(t *testing.T) {
		var got SEEntry

		got.DigiID = "foo"
		got.GetByDigiID()
		if err != nil {
			t.Errorf("Unexpected error: %s", err)
		}

		if got != want {
			t.Errorf("Mismatch: want: \"%v\", got: \"%v\"", want, got)
		}
	})

	t.Run("confirm error is returned when not found", func(t *testing.T) {
		var got SEEntry

		got.DigiID = "fooX"
		err = got.GetByDigiID()
		if err == nil {
			t.Errorf("Expected search for non-existant SE to return error, but err was nil")
		}
	})

}
