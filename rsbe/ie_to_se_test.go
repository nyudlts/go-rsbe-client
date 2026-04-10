package rsbe

import (
	//	"net/http/httptest"
	"testing"

	"github.com/nyudlts/go-rsbe-client/rsbe/internal/testutils"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

var ieToSEToCreate = IEToSEEntry{
	ID:      "06de6d7a-89cd-476c-9e1d-55fdfabc3094",
	IEID:    "9ea98441-b6b6-46cf-b6c8-91dff385c6c8",
	SEID:    "8c258cb2-d700-43be-8773-a61a7b9cd668",
	Order:   1,
	Section: 23,
	Notes:   "IE to SE Notes",
}

func TestIEToSEList(t *testing.T) {

	setupLocalhostClient()
	err := ieToSEToCreate.Create()
	if err != nil {
		t.Errorf("Unexpected error: %s", err)
	}
	defer ieToSEToCreate.Purge()

	t.Run("result", func(t *testing.T) {
		want := ieToSEToCreate
		got, err := IEToSEList()
		require.NoError(t, err, "Unexpected error listing IE to SE entries")
		assert.Equal(t, 1, len(got), "Expected exactly one IE to SE entry in list")
		assert.Equal(t, want.ID, got[0].ID, "ID Mismatch")
		assert.Equal(t, want.IEID, got[0].IEID, "IEID Mismatch")
		assert.Equal(t, want.SEID, got[0].SEID, "SEID Mismatch")
		assert.Equal(t, want.Order, got[0].Order, "Order Mismatch")
		assert.Equal(t, want.Section, got[0].Section, "Section Mismatch")
		assert.NotEmpty(t, got[0].CreatedAt, "CreatedAt is empty")
		assert.NotEmpty(t, got[0].UpdatedAt, "UpdatedAt is empty")
		assert.Contains(t, got[0].URL, "/api/v0/ie_to_ses/06de6d7a-89cd-476c-9e1d-55fdfabc3094", "URL does not contain expected string")
	})
}

func TestIEToSEListWithFilters(t *testing.T) {

	setupLocalhostClient()
	err := ieToSEToCreate.Create()
	if err != nil {
		t.Errorf("Unexpected error: %s", err)
	}
	defer ieToSEToCreate.Purge()

	ieToCreate.ID = "56c61005-ba14-47dc-a073-a03f66cf84e6"
	err = ieToCreate.Create()
	if err != nil {
		t.Errorf("Unexpected error: %s", err)
	}
	defer ieToCreate.Purge()

	ieToSEToCreate1 := IEToSEEntry{
		ID:      "eff4ef7e-961a-4687-8707-990584fa6660",
		IEID:    "56c61005-ba14-47dc-a073-a03f66cf84e6",
		SEID:    ieToSEToCreate.SEID, // reuse existing SE
		Order:   1000,
		Section: 1,
		Notes:   "blah blah blah",
	}
	err = ieToSEToCreate1.Create()
	if err != nil {
		t.Errorf("Unexpected error: %s", err)
	}
	defer ieToSEToCreate1.Purge()

	t.Run("check that proper results are returned", func(t *testing.T) {

		// run without filter
		got, err := IEToSEList()
		require.NoError(t, err, "Unexpected error listing IE to SE entries")
		assert.Equal(t, 2, len(got), "Expected exactly two IE to SE entries in list")

		// run with filter
		filter := IEToSEListEntry{IEID: "56c61005-ba14-47dc-a073-a03f66cf84e6"}
		want := ieToSEToCreate1

		got, err = IEToSEList(filter)
		require.NoError(t, err, "Unexpected error listing IE to SE entries with filter")
		assert.Equal(t, 1, len(got), "Expected exactly one IE to SE entry in list")
		assert.Equal(t, want.ID, got[0].ID, "ID Mismatch")
		assert.Equal(t, want.IEID, got[0].IEID, "IEID Mismatch")
		assert.Equal(t, want.SEID, got[0].SEID, "SEID Mismatch")
		assert.Equal(t, want.Order, got[0].Order, "Order Mismatch")
		assert.Equal(t, want.Section, got[0].Section, "Section Mismatch")
		assert.NotEmpty(t, got[0].CreatedAt, "CreatedAt is empty")
		assert.NotEmpty(t, got[0].UpdatedAt, "UpdatedAt is empty")
		assert.Contains(t, got[0].URL, "/api/v0/ie_to_ses/eff4ef7e-961a-4687-8707-990584fa6660", "URL does not contain expected string")
	})
}

// test that struct method Get works
func TestIEToSEGetFunc(t *testing.T) {

	setupLocalhostClient()
	err := ieToSEToCreate.Create()
	if err != nil {
		t.Errorf("Unexpected error: %s", err)
	}
	defer ieToSEToCreate.Purge()

	t.Run("result", func(t *testing.T) {
		want := ieToSEToCreate
		got := IEToSEEntry{ID: ieToSEToCreate.ID}

		err := got.Get()
		require.NoError(t, err, "Unexpected error getting IE to SE entry by ID")
		assert.Equal(t, want.ID, got.ID, "ID Mismatch")
		assert.Equal(t, want.IEID, got.IEID, "IEID Mismatch")
		assert.Equal(t, want.SEID, got.SEID, "SEID Mismatch")
		assert.Equal(t, want.Order, got.Order, "Order Mismatch")
		assert.Equal(t, want.Section, got.Section, "Section Mismatch")
		assert.Equal(t, want.Notes, got.Notes, "Notes Mismatch")
		assert.NotEmpty(t, got.CreatedAt, "CreatedAt is empty")
		assert.NotEmpty(t, got.UpdatedAt, "UpdatedAt is empty")
		assert.Contains(t, got.IEURL, "/api/v0/ies/9ea98441-b6b6-46cf-b6c8-91dff385c6c8", "IEURL does not contain expected string")
		assert.Contains(t, got.SEURL, "/api/v0/ses/8c258cb2-d700-43be-8773-a61a7b9cd668", "SEURL does not contain expected string")
		assert.Contains(t, got.IEToSEsURL, "/api/v0/ie_to_ses", "IEToSEsURL does not contain expected string")
	})
}

// test that model-level Get works
func TestIEToSEGet(t *testing.T) {

	setupLocalhostClient()
	err := ieToSEToCreate.Create()
	if err != nil {
		t.Errorf("Unexpected error: %s", err)
	}
	defer ieToSEToCreate.Purge()

	t.Run("confirm that expected resource was retrieved", func(t *testing.T) {
		want := ieToSEToCreate
		got, err := IEToSEGet("06de6d7a-89cd-476c-9e1d-55fdfabc3094")
		require.NoError(t, err, "Unexpected error getting IE to SE entry by ID")
		assert.Equal(t, want.ID, got.ID, "ID Mismatch")
		assert.Equal(t, want.IEID, got.IEID, "IEID Mismatch")
		assert.Equal(t, want.SEID, got.SEID, "SEID Mismatch")
		assert.Equal(t, want.Order, got.Order, "Order Mismatch")
		assert.Equal(t, want.Section, got.Section, "Section Mismatch")
		assert.Equal(t, want.Notes, got.Notes, "Notes Mismatch")
		assert.NotEmpty(t, got.CreatedAt, "CreatedAt is empty")
		assert.NotEmpty(t, got.UpdatedAt, "UpdatedAt is empty")
		assert.Contains(t, got.IEURL, "/api/v0/ies/9ea98441-b6b6-46cf-b6c8-91dff385c6c8", "IEURL does not contain expected string")
		assert.Contains(t, got.SEURL, "/api/v0/ses/8c258cb2-d700-43be-8773-a61a7b9cd668", "SEURL does not contain expected string")
		assert.Contains(t, got.IEToSEsURL, "/api/v0/ie_to_ses", "IEToSEsURL does not contain expected string")
	})
}

func TestIEToSECreateFunc(t *testing.T) {
	setupLocalhostClient()

	err := ieToSEToCreate.Create()
	if err != nil {
		t.Errorf("Unexpected error: %s", err)
	}
	defer ieToSEToCreate.Purge()

	t.Run("confirm that attributes updated", func(t *testing.T) {
		assert.NotEmpty(t, ieToSEToCreate.ID, "ID is empty")
		assert.NotEmpty(t, ieToSEToCreate.CreatedAt, "CreatedAt is empty")
		assert.NotEmpty(t, ieToSEToCreate.UpdatedAt, "UpdatedAt is empty")
	})
}

func TestIEToSEUpdateFunc(t *testing.T) {
	setupLocalhostClient()

	err := ieToSEToCreate.Create()
	if err != nil {
		t.Errorf("Unexpected error: %s", err)
	}
	defer ieToSEToCreate.Purge()

	id := ieToSEToCreate.ID

	sut, err := IEToSEGet(id)
	if err != nil {
		t.Errorf("Unexpected error: %s", err)
	}

	if sut.Order != 1 {
		t.Errorf("variable already updated: %s", sut.ToString())
	}

	sut.Order = 97
	sut.Notes = "Hop on Pop!"

	err = sut.Update()
	if err != nil {
		t.Errorf("Unexpected error: %s", err)
	}

	sut = IEToSEEntry{}
	sut, err = IEToSEGet(id)
	if err != nil {
		t.Errorf("Unexpected error: %s", err)
	}

	t.Run("confirm that elements updated", func(t *testing.T) {
		assert.Equal(t, id, sut.ID, "ID should be unchanged")
		assert.Equal(t, 97, sut.Order, "Order was not updated")
		assert.Equal(t, "Hop on Pop!", sut.Notes, "Notes field was not updated")
		testutils.AssertNotEquivalentTimestamps(t, sut.CreatedAt, sut.UpdatedAt, "UpdatedAt not updated")
	})
}

func TestIEToSEDeleteFunc(t *testing.T) {
	setupLocalhostClient()

	err := ieToSEToCreate.Create()
	if err != nil {
		t.Errorf("Unexpected error: %s", err)
	}
	// the Purge() call is needed because Delete() is a soft delete,
	// the record will still exist and needs to be purged
	defer ieToSEToCreate.Purge()

	_ = ieToSEToCreate.Get()

	id := ieToSEToCreate.ID

	err = ieToSEToCreate.Delete()
	if err != nil {
		t.Errorf("Unexpected error: %s", err)
	}

	t.Run("confirm that deleted item not found", func(t *testing.T) {
		// should not be found, so err should NOT be nil
		_, err = IEToSEGet(id)

		if err == nil {
			t.Errorf("err was nil")
		}

	})
}
