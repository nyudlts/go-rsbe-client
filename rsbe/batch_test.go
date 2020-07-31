package rsbe

import (
	"testing"
)

var batchToCreate = BatchEntry{
	Name:         "a super cool batch",
	SourceFile:   "foo.xlsx",
	CollectionID: "b9612d5d-619a-4ceb-b620-d816e4b4340b",
	BatchType:    "transcoding",
	BatchNumber:  1,
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

		if 1 != len(list) {
			t.Errorf("Mismatch: want: \"%v\", got: \"%v\"", 1, len(list))
		}

		want := batchToCreate
		got := list[0]

		if want.ID != got.ID {
			t.Errorf("ID mismatch: want: \"%v\", got: \"%v\"", want.ID, got.ID)
		}

		if want.BatchType != got.BatchType {
			t.Errorf("BatchType mismatch: want: \"%v\", got: \"%v\"", want.BatchType, got.BatchType)
		}

		if want.BatchNumber != got.BatchNumber {
			t.Errorf("BatchNumber mismatch: want: \"%v\", got: \"%v\"", want.BatchNumber, got.BatchNumber)
		}

		if want.Name != got.Name {
			t.Errorf("Name mismatch: want: \"%v\", got: \"%v\"", want.Name, got.Name)
		}

		if want.SourceFile != got.SourceFile {
			t.Errorf("SourceFile mismatch: want: \"%v\", got: \"%v\"", want.SourceFile, got.SourceFile)
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

		if want.BatchType != got.BatchType {
			t.Errorf("BatchType mismatch: want: \"%v\", got: \"%v\"", want.BatchType, got.BatchType)
		}

		if want.BatchNumber != got.BatchNumber {
			t.Errorf("BatchNumber mismatch: want: \"%v\", got: \"%v\"", want.BatchNumber, got.BatchNumber)
		}

		if want.Name != got.Name {
			t.Errorf("Name mismatch: want: \"%v\", got: \"%v\"", want.Name, got.Name)
		}

		if want.SourceFile != got.SourceFile {
			t.Errorf("SourceFile mismatch: want: \"%v\", got: \"%v\"", want.SourceFile, got.SourceFile)
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

// func TestBatchGet(t *testing.T) {

// 	mux := setupMux("/api/v0/ses/8c258cb2-d700-43be-8773-a61a7b9cd668", "testdata/se-get.json")
// 	ts := httptest.NewServer(mux)
// 	defer ts.Close()

// 	setupTestServerClient(ts)

// 	t.Run("confirm that expected partner was retrieved", func(t *testing.T) {
// 		want := seShow
// 		got, err := BatchGet("8c258cb2-d700-43be-8773-a61a7b9cd668")
// 		if err != nil {
// 			t.Errorf("Unexpected error: %s", err)
// 		}

// 		if got != want {
// 			t.Errorf("Mismatch: want: \"%v\", got: \"%v\"", want, got)
// 		}
// 	})
// }

// func TestBatchUpdateFunc(t *testing.T) {
// 	setupLocalhostClient()

// 	_ = batchToCreate.Get()

// 	if batchToCreate.DigiID != "temporary_item" {
// 		t.Errorf("variable already updated: %s", batchToCreate.ToString())
// 	}

// 	batchToCreate.DigiID = "DogBiscuit"

// 	err := batchToCreate.Update()
// 	if err != nil {
// 		t.Errorf("Unexpected error: %s", err)
// 	}

// 	_ = batchToCreate.Get()

// 	t.Run("confirm that elements updated", func(t *testing.T) {
// 		if batchToCreate.DigiID != "DogBiscuit" {
// 			t.Errorf("DigiID was not updated: got: %s", batchToCreate.DigiID)
// 		}

// 		if batchToCreate.CreatedAt == batchToCreate.UpdatedAt {
// 			t.Errorf("UpeatedAt not updated")
// 		}
// 	})
// }

// func TestBatchDeleteFunc(t *testing.T) {
// 	setupLocalhostClient()

// 	_ = batchToCreate.Get()

// 	id := batchToCreate.ID

// 	err := batchToCreate.Delete()
// 	if err != nil {
// 		t.Errorf("Unexpected error: %s", err)
// 	}

// 	t.Run("confirm that deleted item not found", func(t *testing.T) {
// 		// should not be found, so err should NOT be nil
// 		_, err = BatchGet(id)

// 		if err == nil {
// 			t.Errorf("err was nil")
// 		}

// 	})
// }

// func TestBatchGetByDigiID(t *testing.T) {
// 	setupLocalhostClient()

// 	want, err := BatchGet("8c258cb2-d700-43be-8773-a61a7b9cd668")
// 	if err != nil {
// 		t.Errorf("Unexpected error: %s", err)
// 	}

// 	t.Run("confirm that exiting Batch is returned", func(t *testing.T) {
// 		got, err := BatchGetByDigiID("foo")
// 		if err != nil {
// 			t.Errorf("Unexpected error: %s", err)
// 		}

// 		if got != want {
// 			t.Errorf("Mismatch: want: \"%v\", got: \"%v\"", want, got)
// 		}
// 	})

// 	t.Run("confirm error is returned when not found", func(t *testing.T) {
// 		_, err := BatchGetByDigiID("fooX")
// 		if err == nil {
// 			t.Errorf("Expected search for non-existant Batch to return error, but err was nil")
// 		}
// 	})

// }
