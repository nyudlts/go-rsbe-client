package rsbe

import (
	"testing"
)

var batchToSEListEntry = BatchToSEListEntry{
}

var batchToSEShow = BatchToSEEntry{
}

var batchToSEToCreate = BatchToSEEntry{
	BatchID:     "32626389-c942-4e71-9b5a-5d7c7ca4d389",
	SEID: "8c258cb2-d700-43be-8773-a61a7b9cd668",
	Phase: "transcoding",
	Step: "trimming",
	Status: "active",
}

func TestBatchToSECreateFunc(t *testing.T) {
	setupLocalhostClient()

	err := batchToSEToCreate.Create()
	if err != nil {
		t.Errorf("Unexpected error: %s", err)
	}

	t.Run("confirm that attributes updated", func(t *testing.T) {
		if batchToSEToCreate.ID == "" {
			t.Errorf("ID not updated")
		}

		if batchToSEToCreate.CreatedAt == "" {
			t.Errorf("CreatedAt not updated")
		}

		if batchToSEToCreate.UpdatedAt == "" {
			t.Errorf("UpdatedAt not updated")
		}
	})
}

// func TestBatchToSEList(t *testing.T) {

// 	mux := setupMux("/api/v0/batchToSEs", "testdata/batchToSE-list.json")
// 	ts := httptest.NewServer(mux)
// 	defer ts.Close()

// 	setupTestServerClient(ts)

// 	t.Run("result", func(t *testing.T) {
// 		want := batchToSEListEntry
// 		got, err := BatchToSEList()
// 		if err != nil {
// 			t.Errorf("Unexpected error: %s", err)
// 		}

// 		if 4 != len(got) {
// 			t.Errorf("Result Length Mismatch: want: 4, got: %d", len(got))
// 		}

// 		if batchToSEListEntry != got[0] {
// 			t.Errorf("Mismatch: want: \"%v\", got: \"%v\"", want, got)
// 		}
// 	})

// }

// func TestBatchToSEGetFunc(t *testing.T) {

// 	mux := setupMux("/api/v0/batchToSEs/3ca8ecaf-6fae-48a5-8441-5a96e119ad28", "testdata/batchToSE-get.json")
// 	ts := httptest.NewServer(mux)
// 	defer ts.Close()

// 	setupTestServerClient(ts)

// 	t.Run("result", func(t *testing.T) {
// 		want := batchToSEShow
// 		got := BatchToSEEntry{ID: "3ca8ecaf-6fae-48a5-8441-5a96e119ad28"}

// 		err := got.Get()
// 		if err != nil {
// 			t.Errorf("Unexpected error: %s", err)
// 		}

// 		if want.ID != got.ID {
// 			t.Errorf("Mismatch: want: \"%v\", got: \"%v\"", want, got)
// 		}
// 	})

// }

// func TestBatchToSEUpdateFunc(t *testing.T) {
// 	setupLocalhostClient()

// 	_ = batchToSEToCreate.Get()

// 	if batchToSEToCreate.Role != "notes" {
// 		t.Errorf("variable already updated: %s", batchToSEToCreate.ToString())
// 	}

// 	batchToSEToCreate.Role = "waffles"

// 	err := batchToSEToCreate.Update()
// 	if err != nil {
// 		t.Errorf("Unexpected error: %s", err)
// 	}

// 	_ = batchToSEToCreate.Get()

// 	t.Run("confirm that elements updated", func(t *testing.T) {
// 		if batchToSEToCreate.Role != "waffles" {
// 			t.Errorf("Role was not updated: got: %s", batchToSEToCreate.Role)
// 		}

// 		if batchToSEToCreate.CreatedAt == batchToSEToCreate.UpdatedAt {
// 			t.Errorf("UpdatedAt not updated")
// 		}
// 	})
// }

// func TestBatchToSEDeleteFunc(t *testing.T) {
// 	setupLocalhostClient()

// 	_ = batchToSEToCreate.Get()

// 	id := batchToSEToCreate.ID

// 	err := batchToSEToCreate.Delete()
// 	if err != nil {
// 		t.Errorf("Unexpected error: %s", err)
// 	}

// 	t.Run("confirm that deleted item not found", func(t *testing.T) {
// 		// should not be found, so err should NOT be nil
// 		_, err = BatchToSEGet(id)

// 		if err == nil {
// 			t.Errorf("err was nil")
// 		}

// 	})
// }
