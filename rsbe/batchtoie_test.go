package rsbe

import (
	"fmt"
	"testing"
)

var batchToIEListEntry = BatchToIEListEntry{}

var batchToIEShow = BatchToIEEntry{}

var batchToIEToCreate = BatchToIEEntry{
	BatchID: "32626389-c942-4e71-9b5a-5d7c7ca4d389",
	IEID:    "9ea98441-b6b6-46cf-b6c8-91dff385c6c8",
	Phase:   "karms",
	Step:    "remediation",
	Status:  "active",
	Notes:   "more amazing notes",
}

func TestBatchToIECreateFunc(t *testing.T) {
	setupLocalhostClient()

	err := batchToIEToCreate.Create()
	if err != nil {
		t.Errorf("Unexpected error: %s", err)
	}

	t.Run("confirm that attributes updated", func(t *testing.T) {
		if batchToIEToCreate.ID == "" {
			t.Errorf("ID not updated")
		}

		if batchToIEToCreate.CreatedAt == "" {
			t.Errorf("CreatedAt not updated")
		}

		if batchToIEToCreate.UpdatedAt == "" {
			t.Errorf("UpdatedAt not updated")
		}
	})
}

func TestBatchToIEList(t *testing.T) {
	setupLocalhostClient()

	t.Run("check that proper attribute values are returned", func(t *testing.T) {
		list, err := BatchToIEList()
		if err != nil {
			t.Errorf("Unexpected error: %s", err)
		}

		if len(list) != 1 {
			t.Errorf("Mismatch: want: \"%v\", got: \"%v\"", 1, len(list))
		}

		want := batchToIEToCreate
		got := list[0]

		if want.ID != got.ID {
			t.Errorf("ID mismatch: want: \"%v\", got: \"%v\"", want.ID, got.ID)
		}

		if want.BatchID != got.BatchID {
			t.Errorf("BatchID mismatch: want: \"%v\", got: \"%v\"", want.BatchID, got.BatchID)
		}

		if want.IEID != got.IEID {
			t.Errorf("IEID mismatch: want: \"%v\", got: \"%v\"", want.IEID, got.IEID)
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

func TestBatchToIEGetFunc(t *testing.T) {
	setupLocalhostClient()

	t.Run("check that proper attribute values are returned", func(t *testing.T) {
		want := batchToIEToCreate
		got, err := BatchToIEGet(want.ID)
		if err != nil {
			t.Errorf("Unexpected error: %s", err)
		}

		if want.ID != got.ID {
			t.Errorf("ID mismatch: want: \"%v\", got: \"%v\"", want.ID, got.ID)
		}

		if want.BatchID != got.BatchID {
			t.Errorf("BatchID mismatch: want: \"%v\", got: \"%v\"", want.BatchID, got.BatchID)
		}

		if want.IEID != got.IEID {
			t.Errorf("IEID mismatch: want: \"%v\", got: \"%v\"", want.IEID, got.IEID)
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

		expect := fmt.Sprintf("http://localhost:3000/api/v0/ies/%s", want.IEID)
		if expect != got.IEURL {
			t.Errorf("IEURL mismatch: want: \"%v\", got: \"%v\"", expect, got.IEURL)
		}

		expect = fmt.Sprintf("http://localhost:3000/api/v0/batches/%s", want.BatchID)
		if expect != got.BatchURL {
			t.Errorf("BatchURL mismatch: want: \"%v\", got: \"%v\"", expect, got.BatchURL)
		}

		expect = fmt.Sprintf("http://localhost:3000/api/v0/batch_to_ies")
		if expect != got.BatchToIEsURL {
			t.Errorf("BatchToIEsURL mismatch: want: \"%v\", got: \"%v\"", expect, got.BatchToIEsURL)
		}

	})
}

func TestBatchToIEUpdateFunc(t *testing.T) {
	setupLocalhostClient()

	err := batchToIEToCreate.Get()
	if err != nil {
		t.Errorf("Unexpected error: %s", err)
	}

	if batchToIEToCreate.Status != "active" {
		t.Errorf("variable already updated: %s", batchToIEToCreate.Status)
	}

	batchToIEToCreate.Status = "error"

	err = batchToIEToCreate.Update()
	if err != nil {
		t.Errorf("Unexpected error: %s", err)
	}

	err = batchToIEToCreate.Get()
	if err != nil {
		t.Errorf("Unexpected error: %s", err)
	}

	t.Run("confirm that elements updated", func(t *testing.T) {
		if batchToIEToCreate.Status != "error" {
			t.Errorf("Status was not updated: got: %s", batchToIEToCreate.Status)
		}

		if batchToIEToCreate.CreatedAt == batchToIEToCreate.UpdatedAt {
			t.Errorf("UpeatedAt not updated")
		}
	})
}

func TestBatchToIEDeleteFunc(t *testing.T) {
	setupLocalhostClient()

	err := batchToIEToCreate.Get()
	if err != nil {
		t.Errorf("Unexpected error: %s", err)
	}

	id := batchToIEToCreate.ID

	err = batchToIEToCreate.Delete()
	if err != nil {
		t.Errorf("Unexpected error: %s", err)
	}

	t.Run("confirm that deleted item not found", func(t *testing.T) {
		// should not be found, so err should NOT be nil
		_, err = BatchToIEGet(id)

		if err == nil {
			t.Errorf("err was nil")
		}

	})
}
