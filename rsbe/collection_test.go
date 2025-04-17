package rsbe

import (
	"net/http/httptest"
	"testing"
)

var collectionListEntry = CollectionListEntry{
	ID:          "b9612d5d-619a-4ceb-b620-d816e4b4340b",
	PartnerID:   "e6517775-6277-4e25-9373-ee7738e820b5",
	OwnerID:     "1ca830b5-6a2b-43f9-b6bc-4dfeac3ee178",
	Code:        "test",
	DisplayCode: "T.C",
	Name:        "Test Collection",
	Type:        "origin",
	CreatedAt:   "2020-05-30T01:58:38.431Z",
	UpdatedAt:   "2024-08-19T14:53:30.432Z",
	URL:         "http://localhost:3000/api/v0/colls/b9612d5d-619a-4ceb-b620-d816e4b4340b",
	PartnerURL:  "http://localhost:3000/api/v0/partners/e6517775-6277-4e25-9373-ee7738e820b5",
	OwnerURL:    "http://localhost:3000/api/v0/owners/1ca830b5-6a2b-43f9-b6bc-4dfeac3ee178",
}

var collectionShow = CollectionEntry{
	ID:              "b9612d5d-619a-4ceb-b620-d816e4b4340b",
	PartnerID:       "e6517775-6277-4e25-9373-ee7738e820b5",
	OwnerID:         "1ca830b5-6a2b-43f9-b6bc-4dfeac3ee178",
	Code:            "test",
	DisplayCode:     "T.C",
	Name:            "Test Collection",
	Type:            "origin",
	CreatedAt:       "2020-05-30T01:58:38.431Z",
	UpdatedAt:       "2024-08-19T14:53:30.432Z",
	Quota:           500,
	ReadyForContent: true,
	PartnerURL:      "http://localhost:3000/api/v0/partners/e6517775-6277-4e25-9373-ee7738e820b5",
	OwnerURL:        "http://localhost:3000/api/v0/owners/1ca830b5-6a2b-43f9-b6bc-4dfeac3ee178",
	SEsURL:          "http://localhost:3000/api/v0/colls/b9612d5d-619a-4ceb-b620-d816e4b4340b/ses",
	IEsURL:          "http://localhost:3000/api/v0/colls/b9612d5d-619a-4ceb-b620-d816e4b4340b/ies",
	RelPath:         "content/dlts/test",
	LockVersion:     1,
}

var collectionToCreate = CollectionEntry{
	ID:              "044f2d9d-ee53-478a-89a1-a95f0d0c5888",
	PartnerID:       collectionShow.PartnerID,
	OwnerID:         collectionShow.OwnerID,
	Code:            "waffles",
	DisplayCode:     "Big.Waffles",
	Name:            "The Amazing Breakfast Collection",
	Type:            "virtual",
	Classification:  "mixed_open_closed",
	RelPath:         "content/dlts/waffles",
	ReadyForContent: true,
	Quota:           1000,
}

func TestPartnerCollectionList(t *testing.T) {

	mux := setupMux("/api/v0/partners/e6517775-6277-4e25-9373-ee7738e820b5/colls/", "testdata/collection-list.json")
	ts := httptest.NewServer(mux)
	defer ts.Close()

	setupTestServerClient(ts)

	t.Run("result", func(t *testing.T) {
		want := collectionListEntry
		got, err := PartnerCollectionList(collectionShow.PartnerID)
		if err != nil {
			t.Errorf("Unexpected error: %s", err)
		}

		if len(got) != 1 {
			t.Errorf("Mismatch: want: \"%v\", got: \"%v\"", want, got)
		}

		if collectionListEntry != got[0] {
			t.Errorf("Mismatch: want: \n\"%v\", \ngot: \n\"%v\"", want, got[0])
		}
	})

}

func TestOwnerCollectionList(t *testing.T) {

	mux := setupMux("/api/v0/owners/1ca830b5-6a2b-43f9-b6bc-4dfeac3ee178/colls/", "testdata/collection-list.json")
	ts := httptest.NewServer(mux)
	defer ts.Close()

	setupTestServerClient(ts)

	t.Run("result", func(t *testing.T) {
		want := collectionListEntry
		got, err := OwnerCollectionList(collectionShow.OwnerID)
		if err != nil {
			t.Errorf("Unexpected error: %s", err)
		}

		if len(got) != 1 {
			t.Errorf("Mismatch: want: \"%v\", got: \"%v\"", want, got)
		}

		if collectionListEntry != got[0] {
			t.Errorf("Mismatch: want: \n\"%v\", \ngot: \n\"%v\"", want, got[0])
		}
	})

}

func TestCollectionGetFunc(t *testing.T) {

	mux := setupMux("/api/v0/colls/b9612d5d-619a-4ceb-b620-d816e4b4340b", "testdata/collection-get.json")
	ts := httptest.NewServer(mux)
	defer ts.Close()

	setupTestServerClient(ts)

	t.Run("result", func(t *testing.T) {
		want := collectionShow
		got := CollectionEntry{ID: "b9612d5d-619a-4ceb-b620-d816e4b4340b"}

		err := got.Get()
		if err != nil {
			t.Errorf("Unexpected error: %s", err)
		}

		if got != want {
			t.Errorf("Mismatch: want: \n\"%v\", \ngot: \n\"%v\"", want, got)
		}
	})

}

func TestCollectionGet(t *testing.T) {

	mux := setupMux("/api/v0/colls/b9612d5d-619a-4ceb-b620-d816e4b4340b", "testdata/collection-get.json")
	ts := httptest.NewServer(mux)
	defer ts.Close()

	setupTestServerClient(ts)

	t.Run("confirm that expected collection was retrieved", func(t *testing.T) {
		want := collectionShow
		got, err := CollectionGet("b9612d5d-619a-4ceb-b620-d816e4b4340b")
		if err != nil {
			t.Errorf("Unexpected error: %s", err)
		}

		if got != want {
			t.Errorf("Mismatch: want: \n\"%v\", \ngot: \n\"%v\"", want, got)
		}
	})
}

func TestCollectionCreateFunc(t *testing.T) {
	setupLocalhostClient()

	err := collectionToCreate.Create()
	if err != nil {
		t.Errorf("Unexpected error: %s", err)
	}
	defer collectionToCreate.Delete()

	defer t.Run("confirm that attributes updated", func(t *testing.T) {
		if collectionToCreate.CreatedAt == "" {
			t.Errorf("CreatedAt not updated")
		}

		if collectionToCreate.UpdatedAt == "" {
			t.Errorf("UpdatedAt not updated")
		}
	})
}

func TestCollectionUpdateFunc(t *testing.T) {
	setupLocalhostClient()
	err := collectionToCreate.Create()
	if err != nil {
		t.Errorf("Unexpected error: %s", err)
	}
	defer collectionToCreate.Delete()

	err = collectionToCreate.Get()
	if err != nil {
		t.Errorf("Unexpected error: %s", err)
	}

	collectionToCreate.Code = "dogbiscuit"
	collectionToCreate.DisplayCode = "D.B"
	collectionToCreate.Name = "Dog Biscuit"
	collectionToCreate.Type = "origin"
	collectionToCreate.RelPath = "content/dlts/dogbiscuit"
	collectionToCreate.Quota = 0
	collectionToCreate.ReadyForContent = false

	err = collectionToCreate.Update()
	if err != nil {
		t.Errorf("Unexpected error: %s", err)
	}

	err = collectionToCreate.Get()
	if err != nil {
		t.Errorf("Unexpected error: %s", err)
	}

	checkStringContains(t, collectionToCreate.Code, "dogbiscuit")
	checkStringContains(t, collectionToCreate.DisplayCode, "D.B")
	checkStringContains(t, collectionToCreate.Name, "Dog Biscuit")
	checkStringContains(t, collectionToCreate.Type, "origin")
	checkStringContains(t, collectionToCreate.RelPath, "content/dlts/dogbiscuit")
	if collectionToCreate.Quota != 0 {
		t.Errorf("Quota was not updated: got: %d", collectionToCreate.Quota)
	}
	if collectionToCreate.ReadyForContent != false {
		t.Errorf("ReadyForContent was not updated: got: %t", collectionToCreate.ReadyForContent)
	}

	if collectionToCreate.CreatedAt == collectionToCreate.UpdatedAt {
		t.Errorf("UpdatedAt not updated")
	}
}

func TestCollectionDeleteFunc(t *testing.T) {
	setupLocalhostClient()
	err := collectionToCreate.Create()
	if err != nil {
		t.Errorf("Unexpected error: %s", err)
	}

	err = collectionToCreate.Get()
	if err != nil {
		t.Errorf("Unexpected error: %s", err)
	}

	id := collectionToCreate.ID

	err = collectionToCreate.Delete()
	if err != nil {
		t.Errorf("Unexpected error: %s", err)
	}

	t.Run("confirm that deleted item not found", func(t *testing.T) {
		// should not be found, so err should NOT be nil
		_, err := CollectionGet(id)

		if err == nil {
			t.Errorf("err was nil")
		}
	})
}
