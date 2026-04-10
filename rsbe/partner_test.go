package rsbe

import (
	"net/http/httptest"
	"testing"

	"github.com/nyudlts/go-rsbe-client/rsbe/internal/testutils"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

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

		if len(got) != 1 {
			t.Errorf("Mismatch: want: \"%v\", got: \"%v\"", want, got)
		}

		assert.Equal(t, partnerListEntry.ID, got[0].ID)
		assert.Equal(t, partnerListEntry.Code, got[0].Code)
		assert.Equal(t, partnerListEntry.Name, got[0].Name)
		testutils.AssertEquivalentTimestamps(t, partnerListEntry.CreatedAt, got[0].CreatedAt)
		testutils.AssertEquivalentTimestamps(t, partnerListEntry.UpdatedAt, got[0].UpdatedAt)
		assert.Contains(t, got[0].URL, "/api/v0/partners/e6517775-6277-4e25-9373-ee7738e820b5")
	})
}

func TestPartnerGetFunc(t *testing.T) {

	setupLocalhostClient()
	t.Run("result", func(t *testing.T) {
		want := partnerShow
		got := PartnerEntry{ID: "e6517775-6277-4e25-9373-ee7738e820b5"}

		require.NoError(t, got.Get(), "Partner Get failed")

		assert.Equal(t, want.ID, got.ID)
		assert.Equal(t, want.Code, got.Code)
		assert.Equal(t, want.Name, got.Name)
		testutils.AssertEquivalentTimestamps(t, want.CreatedAt, got.CreatedAt)
		testutils.AssertEquivalentTimestamps(t, want.UpdatedAt, got.UpdatedAt)
		assert.Contains(t, got.PartnersURL, "/api/v0/partners")
		assert.Contains(t, got.CollectionsURL, "/api/v0/partners/e6517775-6277-4e25-9373-ee7738e820b5/colls")
		assert.Equal(t, want.LockVersion, got.LockVersion)
		assert.Equal(t, want.RelPath, got.RelPath)
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
		require.NoError(t, err, "Partner Get failed")

		assert.Equal(t, want.ID, got.ID)
		assert.Equal(t, want.Code, got.Code)
		assert.Equal(t, want.Name, got.Name)
		testutils.AssertEquivalentTimestamps(t, want.CreatedAt, got.CreatedAt)
		testutils.AssertEquivalentTimestamps(t, want.UpdatedAt, got.UpdatedAt)
		assert.Contains(t, got.PartnersURL, "/api/v0/partners")
		assert.Contains(t, got.CollectionsURL, "/api/v0/partners/e6517775-6277-4e25-9373-ee7738e820b5/colls")
		assert.Equal(t, want.LockVersion, got.LockVersion)
		assert.Equal(t, want.RelPath, got.RelPath)
	})
}

func TestPartnerCreateFunc(t *testing.T) {
	setupLocalhostClient()

	err := partnerToCreate.Create()
	require.NoError(t, err, "Partner Create failed")
	defer partnerToCreate.Purge()

	require.NoError(t, partnerToCreate.Get(), "Partner Get failed")

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

	err := partnerToCreate.Create()
	require.NoError(t, err, "Partner Create failed")
	defer partnerToCreate.Purge()

	err = partnerToCreate.Get()
	require.NoError(t, err, "Partner Get failed")

	if partnerToCreate.Name != "Waffles and Syrup" {
		t.Errorf("variable already updated: %s", partnerToCreate.ToString())
	}

	partnerToCreate.Name = "WAFFLES WAFFLES WAFFLES"

	err = partnerToCreate.Update()
	if err != nil {
		t.Errorf("Unexpected error: %s", err)
	}

	_ = partnerToCreate.Get()

	t.Run("confirm that elements updated", func(t *testing.T) {
		if partnerToCreate.Name != "WAFFLES WAFFLES WAFFLES" {
			t.Errorf("Name was not updated: got: %s", partnerToCreate.Name)
		}

		if partnerToCreate.CreatedAt == partnerToCreate.UpdatedAt {
			t.Errorf("UpdatedAt not updated")
		}
	})
}

func TestPartnerDeleteFunc(t *testing.T) {
	setupLocalhostClient()

	err := partnerToCreate.Create()
	require.NoError(t, err, "Partner Create failed")
	// the Purge() call is needed because Delete() is a soft delete,
	// the record will still exist and needs to be purged
	defer partnerToCreate.Purge()

	err = partnerToCreate.Get()
	require.NoError(t, err, "Partner Get failed")

	id := partnerToCreate.ID

	err = partnerToCreate.Delete()
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
