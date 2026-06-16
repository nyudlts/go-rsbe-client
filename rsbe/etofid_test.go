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

		if len(got) != 1 {
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
	defer etofidToCreate.Purge()

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

	err := etofidToCreate.Create()
	if err != nil {
		t.Errorf("Unexpected error: %s", err)
	}
	defer etofidToCreate.Purge()

	_ = etofidToCreate.Get()

	if etofidToCreate.FIDValue != "7284f" {
		t.Errorf("variable already updated: %s", etofidToCreate.ToString())
	}

	etofidToCreate.FIDValue = "waffles"

	err = etofidToCreate.Update()
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

	err := etofidToCreate.Create()
	if err != nil {
		t.Errorf("Unexpected error: %s", err)
	}
	// the Purge() call is needed because Delete() is a soft delete,
	// the record will still exist and needs to be purged
	defer etofidToCreate.Purge()

	_ = etofidToCreate.Get()

	id := etofidToCreate.ID

	err = etofidToCreate.Delete()
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

func TestEToFIDListFiltering(t *testing.T) {
	setupLocalhostClient()

	// Create test data with different values for filtering
	etofid1 := EToFIDEntry{
		EType:    "ie",
		EID:      "9ea98441-b6b6-46cf-b6c8-91dff385c6c8",
		FIDType:  "handle",
		FIDValue: "test.1/abc123",
	}

	etofid2 := EToFIDEntry{
		EType:    "se",
		EID:      "8c258cb2-d700-43be-8773-a61a7b9cd668",
		FIDType:  "citation_url",
		FIDValue: "https://hdl.handle.net/test.2/def456",
	}

	etofid3 := EToFIDEntry{
		EType:    "ie",
		EID:      "9ea98441-b6b6-46cf-b6c8-91dff385c6c8",
		FIDType:  "noid",
		FIDValue: "ghi789",
	}

	// Create all test entries
	if err := etofid1.Create(); err != nil {
		t.Errorf("Failed to create test entry 1: %s", err)
	}
	defer etofid1.Purge()

	if err := etofid2.Create(); err != nil {
		t.Errorf("Failed to create test entry 2: %s", err)
	}
	defer etofid2.Purge()

	if err := etofid3.Create(); err != nil {
		t.Errorf("Failed to create test entry 3: %s", err)
	}
	defer etofid3.Purge()

	t.Run("list with no filters returns all entries", func(t *testing.T) {
		list, err := EToFIDList()
		if err != nil {
			t.Errorf("Unexpected error: %s", err)
		}

		if len(list) < 3 {
			t.Errorf("Expected at least 3 entries, got: %d", len(list))
		}
	})

	t.Run("list with empty filter returns all entries", func(t *testing.T) {
		list, err := EToFIDList(EToFIDListEntry{})
		if err != nil {
			t.Errorf("Unexpected error: %s", err)
		}

		if len(list) < 3 {
			t.Errorf("Expected at least 3 entries, got: %d", len(list))
		}
	})

	t.Run("filter by EID", func(t *testing.T) {
		list, err := EToFIDList(EToFIDListEntry{EID: etofid1.EID})
		if err != nil {
			t.Errorf("Unexpected error: %s", err)
		}

		if len(list) < 1 {
			t.Errorf("Expected at least 1 entry, got: %d", len(list))
		}

		found := false
		for _, entry := range list {
			if entry.EID == etofid1.EID {
				found = true
				break
			}
		}

		if !found {
			t.Errorf("Expected to find entry with EID: %s", etofid1.EID)
		}
	})

	t.Run("filter by EType", func(t *testing.T) {
		list, err := EToFIDList(EToFIDListEntry{EType: "ie"})
		if err != nil {
			t.Errorf("Unexpected error: %s", err)
		}

		if len(list) < 2 {
			t.Errorf("Expected at least 2 IE entries, got: %d", len(list))
		}

		for _, entry := range list {
			if entry.EType != "ie" {
				t.Errorf("Expected EType 'ie', got: %s", entry.EType)
			}
		}
	})

	t.Run("filter by FIDType", func(t *testing.T) {
		list, err := EToFIDList(EToFIDListEntry{FIDType: "handle"})
		if err != nil {
			t.Errorf("Unexpected error: %s", err)
		}

		if len(list) < 1 {
			t.Errorf("Expected at least 1 entry with FIDType 'handle', got: %d", len(list))
		}

		found := false
		for _, entry := range list {
			if entry.FIDType == "handle" && entry.FIDValue == etofid1.FIDValue {
				found = true
				break
			}
		}

		if !found {
			t.Errorf("Expected to find entry with FIDType 'handle' and FIDValue: %s", etofid1.FIDValue)
		}
	})

	t.Run("filter by FIDValue", func(t *testing.T) {
		list, err := EToFIDList(EToFIDListEntry{FIDValue: etofid2.FIDValue})
		if err != nil {
			t.Errorf("Unexpected error: %s", err)
		}

		if len(list) < 1 {
			t.Errorf("Expected at least 1 entry, got: %d", len(list))
		}

		found := false
		for _, entry := range list {
			if entry.FIDValue == etofid2.FIDValue {
				found = true
				break
			}
		}

		if !found {
			t.Errorf("Expected to find entry with FIDValue: %s", etofid2.FIDValue)
		}
	})

	t.Run("filter by multiple parameters", func(t *testing.T) {
		list, err := EToFIDList(EToFIDListEntry{
			EType:   "se",
			FIDType: "citation_url",
		})
		if err != nil {
			t.Errorf("Unexpected error: %s", err)
		}

		if len(list) < 1 {
			t.Errorf("Expected at least 1 entry, got: %d", len(list))
		}

		found := false
		for _, entry := range list {
			if entry.EType == "se" && entry.FIDType == "citation_url" {
				found = true
				break
			}
		}

		if !found {
			t.Errorf("Expected to find entry with EType 'se' and FIDType 'citation_url'")
		}
	})
}
