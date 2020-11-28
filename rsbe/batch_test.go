package rsbe

import (
	"net/http/httptest"
	"testing"
	"sort"
)

var batchToCreate = BatchEntry{
	Name:         "a super cool batch",
	Source:       "foo.xlsx",
	CollectionID: "b9612d5d-619a-4ceb-b620-d816e4b4340b",
	Type:         "transcoding",
	Number:       2,
	Notes:        "get stuff ready",
}

func TestBatchCreateFunc(t *testing.T) {
	setupLocalhostClient()

	err := batchToCreate.Create()
	if err != nil {
		t.Errorf("Unexpected error: %s", err)
	}

	t.Run("confirm that attributes updated", func(t *testing.T) {
		if batchToCreate.ID == "" {
			t.Errorf("ID not updated")
		}

		if batchToCreate.CreatedAt == "" {
			t.Errorf("CreatedAt not updated")
		}

		if batchToCreate.UpdatedAt == "" {
			t.Errorf("UpdatedAt not updated")
		}
	})
}

func TestBatchList(t *testing.T) {
	setupLocalhostClient()

	t.Run("check that proper attribute values are returned", func(t *testing.T) {
		list, err := BatchList()
		if err != nil {
			t.Errorf("Unexpected error: %s", err)
		}

		if 2 != len(list) {
			t.Errorf("Mismatch: want: \"%v\", got: \"%v\"", 2, len(list))
		}

		// this is needed because the API does not guarantee order or returned elements
		sort.SliceStable(list, func(i, j int) bool {
			return list[i].Number < list[j].Number
		})

		want := batchToCreate
		got := list[1]

		if want.ID != got.ID {
			t.Errorf("ID mismatch: want: \"%v\", got: \"%v\"", want.ID, got.ID)
		}

		if want.Type != got.Type {
			t.Errorf("Type mismatch: want: \"%v\", got: \"%v\"", want.Type, got.Type)
		}

		if want.Number != got.Number {
			t.Errorf("Number mismatch: want: \"%v\", got: \"%v\"", want.Number, got.Number)
		}

		if want.Name != got.Name {
			t.Errorf("Name mismatch: want: \"%v\", got: \"%v\"", want.Name, got.Name)
		}

		if want.Source != got.Source {
			t.Errorf("Source mismatch: want: \"%v\", got: \"%v\"", want.Source, got.Source)
		}

		if want.CollectionID != got.CollectionID {
			t.Errorf("CollectionID mismatch: want: \"%v\", got: \"%v\"", want.CollectionID, got.CollectionID)
		}

		if want.CreatedAt != got.CreatedAt {
			t.Errorf("CreatedAt mismatch: want: \"%v\", got: \"%v\"", want.CreatedAt, got.CreatedAt)
		}

		if want.UpdatedAt != got.UpdatedAt {
			t.Errorf("UpdatedAt mismatch: want: \"%v\", got: \"%v\"", want.UpdatedAt, got.UpdatedAt)
		}

	})

}

func TestBatchReport(t *testing.T) {

	mux := setupMux("/api/v0/batches/c44e95e9-5cca-4c26-8e52-12773334dc95/report", "testdata/batch-report.json")
	ts := httptest.NewServer(mux)
	defer ts.Close()

	setupTestServerClient(ts)

	t.Run("result", func(t *testing.T) {
		report, err := BatchReportGet("c44e95e9-5cca-4c26-8e52-12773334dc95")
		if err != nil {
			t.Errorf("Unexpected error: %s", err)
		}

		want := "2020-11-27T01:05:44Z"
		got  := report.TimeStamp
		if want != got {
			t.Errorf("Mismatch: want: \"%v\", got: \"%v\"", want, got)
		}
	})

}

func TestBatchGetFunc(t *testing.T) {
	setupLocalhostClient()

	t.Run("check that proper attribute values are returned", func(t *testing.T) {
		want := batchToCreate
		got, err := BatchGet(want.ID)
		if err != nil {
			t.Errorf("Unexpected error: %s", err)
		}

		if want.ID != got.ID {
			t.Errorf("ID mismatch: want: \"%v\", got: \"%v\"", want.ID, got.ID)
		}

		if want.Type != got.Type {
			t.Errorf("Type mismatch: want: \"%v\", got: \"%v\"", want.Type, got.Type)
		}

		if want.Number != got.Number {
			t.Errorf("Number mismatch: want: \"%v\", got: \"%v\"", want.Number, got.Number)
		}

		if want.Name != got.Name {
			t.Errorf("Name mismatch: want: \"%v\", got: \"%v\"", want.Name, got.Name)
		}

		if want.Source != got.Source {
			t.Errorf("Source mismatch: want: \"%v\", got: \"%v\"", want.Source, got.Source)
		}

		if want.CollectionID != got.CollectionID {
			t.Errorf("CollectionID mismatch: want: \"%v\", got: \"%v\"", want.CollectionID, got.CollectionID)
		}

		if want.CreatedAt != got.CreatedAt {
			t.Errorf("CreatedAt mismatch: want: \"%v\", got: \"%v\"", want.CreatedAt, got.CreatedAt)
		}

		if want.UpdatedAt != got.UpdatedAt {
			t.Errorf("UpdatedAt mismatch: want: \"%v\", got: \"%v\"", want.UpdatedAt, got.UpdatedAt)
		}

		if want.Notes != got.Notes {
			t.Errorf("Notes mismatch: want: \"%v\", got: \"%v\"", want.Notes, got.Notes)
		}
	})

}

func TestBatchGet(t *testing.T) {
	setupLocalhostClient()

	t.Run("confirm that expected batch was retrieved", func(t *testing.T) {
		want := batchToCreate
		got, err := BatchGet(batchToCreate.ID)
		if err != nil {
			t.Errorf("Unexpected error: %s", err)
		}

		if want.ID != got.ID {
			t.Errorf("ID mismatch: want: \"%v\", got: \"%v\"", want.ID, got.ID)
		}

		if want.Type != got.Type {
			t.Errorf("Type mismatch: want: \"%v\", got: \"%v\"", want.Type, got.Type)
		}

		if want.Number != got.Number {
			t.Errorf("Number mismatch: want: \"%v\", got: \"%v\"", want.Number, got.Number)
		}

		if want.Name != got.Name {
			t.Errorf("Name mismatch: want: \"%v\", got: \"%v\"", want.Name, got.Name)
		}

		if want.Source != got.Source {
			t.Errorf("Source mismatch: want: \"%v\", got: \"%v\"", want.Source, got.Source)
		}

		if want.CollectionID != got.CollectionID {
			t.Errorf("CollectionID mismatch: want: \"%v\", got: \"%v\"", want.CollectionID, got.CollectionID)
		}

		if want.CreatedAt != got.CreatedAt {
			t.Errorf("CreatedAt mismatch: want: \"%v\", got: \"%v\"", want.CreatedAt, got.CreatedAt)
		}

		if want.UpdatedAt != got.UpdatedAt {
			t.Errorf("UpdatedAt mismatch: want: \"%v\", got: \"%v\"", want.UpdatedAt, got.UpdatedAt)
		}

		if want.Notes != got.Notes {
			t.Errorf("Notes mismatch: want: \"%v\", got: \"%v\"", want.Notes, got.Notes)
		}

	})
}

func TestBatchUpdateFunc(t *testing.T) {
	setupLocalhostClient()

	err := batchToCreate.Get()
	if err != nil {
		t.Errorf("Unexpected error: %s", err)
	}

	if batchToCreate.Name != "a super cool batch" {
		t.Errorf("variable already updated: %s", batchToCreate.Name)
	}

	batchToCreate.Name = "DogBiscuit"

	err = batchToCreate.Update()
	if err != nil {
		t.Errorf("Unexpected error: %s", err)
	}

	err = batchToCreate.Get()
	if err != nil {
		t.Errorf("Unexpected error: %s", err)
	}

	t.Run("confirm that elements updated", func(t *testing.T) {
		if batchToCreate.Name != "DogBiscuit" {
			t.Errorf("Name was not updated: got: %s", batchToCreate.Name)
		}

		if batchToCreate.CreatedAt == batchToCreate.UpdatedAt {
			t.Errorf("UpeatedAt not updated")
		}
	})
}

func TestBatchDeleteFunc(t *testing.T) {
	setupLocalhostClient()

	err := batchToCreate.Get()
	if err != nil {
		t.Errorf("Unexpected error: %s", err)
	}

	id := batchToCreate.ID

	err = batchToCreate.Delete()
	if err != nil {
		t.Errorf("Unexpected error: %s", err)
	}

	t.Run("confirm that deleted item not found", func(t *testing.T) {
		// should not be found, so err should NOT be nil
		_, err = BatchGet(id)

		if err == nil {
			t.Errorf("err was nil")
		}

	})
}
