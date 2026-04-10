package rsbe

import (
	"net/http/httptest"
	"testing"

	"github.com/nyudlts/go-rsbe-client/rsbe/internal/testutils"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

var ownerListEntry = OwnerListEntry{
	ID:        "1ca830b5-6a2b-43f9-b6bc-4dfeac3ee178",
	Code:      "test_owner",
	Name:      "Test Owner",
	CreatedAt: "2024-08-19T14:51:30.383Z",
	UpdatedAt: "2024-08-19T14:51:30.383Z",
	URL:       "http://localhost:3000/api/v0/owners/1ca830b5-6a2b-43f9-b6bc-4dfeac3ee178",
}

var ownerShow = OwnerEntry{
	ID:             "1ca830b5-6a2b-43f9-b6bc-4dfeac3ee178",
	Code:           "test_owner",
	Name:           "Test Owner",
	CreatedAt:      "2024-08-19T14:51:30.383Z",
	UpdatedAt:      "2024-08-19T14:51:30.383Z",
	OwnersURL:      "http://localhost:3000/api/v0/owners",
	CollectionsURL: "http://localhost:3000/api/v0/owners/1ca830b5-6a2b-43f9-b6bc-4dfeac3ee178/colls",
	LockVersion:    0,
}

var ownerToCreate = OwnerEntry{
	Code: "test_owner_2",
	Name: "Test Owner 2",
}

func TestOwnerList(t *testing.T) {

	setupLocalhostClient()
	t.Run("result", func(t *testing.T) {
		want := ownerListEntry
		got, err := OwnerList()
		if err != nil {
			t.Errorf("Unexpected error: %s", err)
		}

		if len(got) != 1 {
			t.Errorf("Mismatch: want: \"%v\", got: \"%v\"", 1, len(got))
		}

		assert.Equal(t, want.ID, got[0].ID)
		assert.Equal(t, want.Code, got[0].Code)
		assert.Equal(t, want.Name, got[0].Name)
		testutils.AssertEquivalentTimestamps(t, want.CreatedAt, got[0].CreatedAt)
		testutils.AssertEquivalentTimestamps(t, want.UpdatedAt, got[0].UpdatedAt)
		assert.Contains(t, got[0].URL, "/api/v0/owners/1ca830b5-6a2b-43f9-b6bc-4dfeac3ee178")
	})

}

func TestOwnerGetFunc(t *testing.T) {

	setupLocalhostClient()
	t.Run("result", func(t *testing.T) {
		want := ownerShow
		got := OwnerEntry{ID: "1ca830b5-6a2b-43f9-b6bc-4dfeac3ee178"}

		require.NoError(t, got.Get(), "Owner Get failed")

		assert.Equal(t, want.ID, got.ID)
		assert.Equal(t, want.Code, got.Code)
		assert.Equal(t, want.Name, got.Name)
		testutils.AssertEquivalentTimestamps(t, want.CreatedAt, got.CreatedAt)
		testutils.AssertEquivalentTimestamps(t, want.UpdatedAt, got.UpdatedAt)
		assert.Contains(t, got.OwnersURL, "/api/v0/owners")
		assert.Contains(t, got.CollectionsURL, "/api/v0/owners/1ca830b5-6a2b-43f9-b6bc-4dfeac3ee178/colls")
		assert.Equal(t, want.LockVersion, got.LockVersion)
	})

}

func TestOwnerGet(t *testing.T) {

	mux := setupMux("/api/v0/owners/1ca830b5-6a2b-43f9-b6bc-4dfeac3ee178", "testdata/owner-get.json")
	ts := httptest.NewServer(mux)
	defer ts.Close()

	setupTestServerClient(ts)

	t.Run("confirm that expected owner was retrieved", func(t *testing.T) {
		want := ownerShow
		got, err := OwnerGet("1ca830b5-6a2b-43f9-b6bc-4dfeac3ee178")
		require.NoError(t, err, "Owner Get failed")

		assert.Equal(t, want.ID, got.ID)
		assert.Equal(t, want.Code, got.Code)
		assert.Equal(t, want.Name, got.Name)
		testutils.AssertEquivalentTimestamps(t, want.CreatedAt, got.CreatedAt)
		testutils.AssertEquivalentTimestamps(t, want.UpdatedAt, got.UpdatedAt)
		assert.Contains(t, got.OwnersURL, "/api/v0/owners")
		assert.Contains(t, got.CollectionsURL, "/api/v0/owners/1ca830b5-6a2b-43f9-b6bc-4dfeac3ee178/colls")
		assert.Equal(t, want.LockVersion, got.LockVersion)
	})
}

func TestOwnerCreateFunc(t *testing.T) {
	setupLocalhostClient()

	err := ownerToCreate.Create()
	require.NoError(t, err, "Owner Create failed")
	defer ownerToCreate.Purge()

	require.NoError(t, ownerToCreate.Get(), "Owner Get failed")

	t.Run("confirm that attributes updated", func(t *testing.T) {
		if ownerToCreate.ID == "" {
			t.Errorf("ID not updated")
		}

		if ownerToCreate.CreatedAt == "" {
			t.Errorf("CreatedAt not updated")
		}

		if ownerToCreate.UpdatedAt == "" {
			t.Errorf("UpdatedAt not updated")
		}
	})
}

func TestOwnerUpdateFunc(t *testing.T) {
	setupLocalhostClient()

	err := ownerToCreate.Create()
	require.NoError(t, err, "Owner Create failed")
	defer ownerToCreate.Purge()

	err = ownerToCreate.Get()
	require.NoError(t, err, "Owner Get failed")

	if ownerToCreate.Name != "Test Owner 2" {
		t.Errorf("variable already updated: %s", ownerToCreate.ToString())
	}

	ownerToCreate.Name = "WAFFLES WAFFLES WAFFLES"

	err = ownerToCreate.Update()
	if err != nil {
		t.Errorf("Unexpected error: %s", err)
	}

	_ = ownerToCreate.Get()

	t.Run("confirm that elements updated", func(t *testing.T) {
		if ownerToCreate.Name != "WAFFLES WAFFLES WAFFLES" {
			t.Errorf("Name was not updated: got: %s", ownerToCreate.Name)
		}

		if ownerToCreate.CreatedAt == ownerToCreate.UpdatedAt {
			t.Errorf("UpeatedAt not updated")
		}
	})
}

func TestOwnerDeleteFunc(t *testing.T) {
	setupLocalhostClient()

	err := ownerToCreate.Create()
	require.NoError(t, err, "Owner Create failed")
	// the Purge() call is needed because Delete() is a soft delete,
	// the record will still exist and needs to be purged
	defer ownerToCreate.Purge()

	err = ownerToCreate.Get()
	require.NoError(t, err, "Owner Get failed")

	id := ownerToCreate.ID

	err = ownerToCreate.Delete()
	if err != nil {
		t.Errorf("Unexpected error: %s", err)
	}

	t.Run("confirm that deleted item not found", func(t *testing.T) {
		// should not be found, so err should NOT be nil
		_, err = OwnerGet(id)

		if err == nil {
			t.Errorf("err was nil")
		}

	})
}
