package rsbe

import (
	"net/http/httptest"
	"testing"
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

		if ownerListEntry != got[0] {
			t.Errorf("Mismatch: wanXZt: \n\"%v\", \ngot: \n\"%v\"", want, got)
		}
	})

}

func TestOwnerGetFunc(t *testing.T) {

	setupLocalhostClient()
	t.Run("result", func(t *testing.T) {
		want := ownerShow
		got := OwnerEntry{ID: "1ca830b5-6a2b-43f9-b6bc-4dfeac3ee178"}

		err := got.Get()
		if err != nil {
			t.Errorf("Unexpected error: %s", err)
		}

		if got != want {
			t.Errorf("Mismatch: want: \"%v\", got: \"%v\"", want, got)
		}
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
		if err != nil {
			t.Errorf("Unexpected error: %s", err)
		}

		if got != want {
			t.Errorf("Mismatch: want: \"%v\", got: \"%v\"", want, got)
		}
	})
}

func TestOwnerCreateFunc(t *testing.T) {
	setupLocalhostClient()

	err := ownerToCreate.Create()
	if err != nil {
		t.Errorf("Unexpected error: %s", err)
	}

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

	_ = ownerToCreate.Get()

	if ownerToCreate.Name != "Test Owner 2" {
		t.Errorf("variable already updated: %s", ownerToCreate.ToString())
	}

	ownerToCreate.Name = "WAFFLES WAFFLES WAFFLES"

	err := ownerToCreate.Update()
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

	_ = ownerToCreate.Get()

	id := ownerToCreate.ID

	err := ownerToCreate.Delete()
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
