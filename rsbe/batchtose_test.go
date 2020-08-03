package rsbe

import (
	"fmt"
	"testing"
)

var batchToSEListEntry = BatchToSEListEntry{}

var batchToSEShow = BatchToSEEntry{}

var batchToSEToCreate = BatchToSEEntry{
	BatchID: "32626389-c942-4e71-9b5a-5d7c7ca4d389",
	SEID:    "8c258cb2-d700-43be-8773-a61a7b9cd668",
	Phase:   "transcoding",
	Step:    "trimming",
	Status:  "active",
	Notes:   "amazing notes, as always",
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

func TestBatchToSEList(t *testing.T) {
	setupLocalhostClient()

	t.Run("check that proper attribute values are returned", func(t *testing.T) {
		list, err := BatchToSEList()
		if err != nil {
			t.Errorf("Unexpected error: %s", err)
		}

		if 1 != len(list) {
			t.Errorf("Mismatch: want: \"%v\", got: \"%v\"", 1, len(list))
		}

		want := batchToSEToCreate
		got := list[0]

		if want.ID != got.ID {
			t.Errorf("ID mismatch: want: \"%v\", got: \"%v\"", want.ID, got.ID)
		}

		if want.BatchID != got.BatchID {
			t.Errorf("BatchID mismatch: want: \"%v\", got: \"%v\"", want.BatchID, got.BatchID)
		}

		if want.SEID != got.SEID {
			t.Errorf("SEID mismatch: want: \"%v\", got: \"%v\"", want.SEID, got.SEID)
		}

		if want.Phase != got.Phase {
			t.Errorf("Phase mismatch: want: \"%v\", got: \"%v\"", want.Phase, got.Phase)
		}

		if want.Step != got.Step {
			t.Errorf("Step mismatch: want: \"%v\", got: \"%v\"", want.Step, got.Step)
		}

		if want.Status != got.Status {
			t.Errorf("Status mismatch: want: \"%v\", got: \"%v\"", want.Status, got.Status)
		}

		if want.CreatedAt != got.CreatedAt {
			t.Errorf("CreatedAt mismatch: want: \"%v\", got: \"%v\"", want.CreatedAt, got.CreatedAt)
		}

		if want.UpdatedAt != got.UpdatedAt {
			t.Errorf("UpdatedAt mismatch: want: \"%v\", got: \"%v\"", want.UpdatedAt, got.UpdatedAt)
		}
	})
}

func TestBatchToSEGetFunc(t *testing.T) {
	setupLocalhostClient()

	t.Run("check that proper attribute values are returned", func(t *testing.T) {
		want := batchToSEToCreate
		got, err := BatchToSEGet(want.ID)
		if err != nil {
			t.Errorf("Unexpected error: %s", err)
		}

		if want.ID != got.ID {
			t.Errorf("ID mismatch: want: \"%v\", got: \"%v\"", want.ID, got.ID)
		}

		if want.BatchID != got.BatchID {
			t.Errorf("BatchID mismatch: want: \"%v\", got: \"%v\"", want.BatchID, got.BatchID)
		}

		if want.SEID != got.SEID {
			t.Errorf("SEID mismatch: want: \"%v\", got: \"%v\"", want.SEID, got.SEID)
		}

		if want.Phase != got.Phase {
			t.Errorf("Phase mismatch: want: \"%v\", got: \"%v\"", want.Phase, got.Phase)
		}

		if want.Step != got.Step {
			t.Errorf("Step mismatch: want: \"%v\", got: \"%v\"", want.Step, got.Step)
		}

		if want.Status != got.Status {
			t.Errorf("Status mismatch: want: \"%v\", got: \"%v\"", want.Status, got.Status)
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

		if want.LockVersion != got.LockVersion {
			t.Errorf("LockVersion mismatch: want: \"%v\", got: \"%v\"", want.LockVersion, got.LockVersion)
		}

		expect := fmt.Sprintf("http://localhost:3000/api/v0/ses/%s", want.SEID)
		if expect != got.SEURL {
			t.Errorf("SEURL mismatch: want: \"%v\", got: \"%v\"", expect, got.SEURL)
		}

		expect = fmt.Sprintf("http://localhost:3000/api/v0/batches/%s", want.BatchID)
		if expect != got.BatchURL {
			t.Errorf("BatchURL mismatch: want: \"%v\", got: \"%v\"", expect, got.BatchURL)
		}

		expect = fmt.Sprintf("http://localhost:3000/api/v0/batch_to_ses")
		if expect != got.BatchToSEsURL {
			t.Errorf("BatchToSEsURL mismatch: want: \"%v\", got: \"%v\"", expect, got.BatchToSEsURL)
		}

	})
}

func TestBatchToSEUpdateFunc(t *testing.T) {
	setupLocalhostClient()

	err := batchToSEToCreate.Get()
	if err != nil {
		t.Errorf("Unexpected error: %s", err)
	}

	if batchToSEToCreate.Status != "active" {
		t.Errorf("variable already updated: %s", batchToSEToCreate.Status)
	}

	batchToSEToCreate.Status = "error"

	err = batchToSEToCreate.Update()
	if err != nil {
		t.Errorf("Unexpected error: %s", err)
	}

	err = batchToSEToCreate.Get()
	if err != nil {
		t.Errorf("Unexpected error: %s", err)
	}

	t.Run("confirm that elements updated", func(t *testing.T) {
		if batchToSEToCreate.Status != "error" {
			t.Errorf("Status was not updated: got: %s", batchToSEToCreate.Status)
		}

		if batchToSEToCreate.CreatedAt == batchToSEToCreate.UpdatedAt {
			t.Errorf("UpeatedAt not updated")
		}
	})
}

func TestBatchToSEDeleteFunc(t *testing.T) {
	setupLocalhostClient()

	err := batchToSEToCreate.Get()
	if err != nil {
		t.Errorf("Unexpected error: %s", err)
	}

	id := batchToSEToCreate.ID

	err = batchToSEToCreate.Delete()
	if err != nil {
		t.Errorf("Unexpected error: %s", err)
	}

	t.Run("confirm that deleted item not found", func(t *testing.T) {
		// should not be found, so err should NOT be nil
		_, err = BatchToSEGet(id)

		if err == nil {
			t.Errorf("err was nil")
		}

	})
}
