package rsbe

import (
	"fmt"
	"testing"
)

// var batchToSEListEntry = BatchToSEListEntry{}
// var batchToSEShow = BatchToSEEntry{}

var seToCreateForBatchToSETest = SEEntry{
	ID:           "cee91db3-ee73-4953-a05e-98f043d44f97",
	CollectionID: "b9612d5d-619a-4ceb-b620-d816e4b4340b",
	DigiID:       "wakawaka",
	DOType:       "audio",
	Phase:        "digitization",
	Step:         "digitization",
	Status:       "canceled",
}

var batchToSEToCreate = BatchToSEEntry{
	BatchID: "32626389-c942-4e71-9b5a-5d7c7ca4d389",
	SEID:    "8c258cb2-d700-43be-8773-a61a7b9cd668",
	Phase:   "transcoding",
	Step:    "trimming",
	Status:  "active",
	Notes:   "amazing notes, as always",
}

var batchToSEToCreate2 = BatchToSEEntry{
	BatchID: "32626389-c942-4e71-9b5a-5d7c7ca4d389",
	SEID:    "cee91db3-ee73-4953-a05e-98f043d44f97",
	Phase:   "baking",
	Step:    "measuring-flour",
	Status:  "covered-in-dust",
	Notes:   "I wish I had a useful recipe",
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

		if len(list) != 1 {
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

	t.Run("check that proper attribute values are returned when an empty BatchToSEListEntry is passed", func(t *testing.T) {
		list, err := BatchToSEList(BatchToSEListEntry{})
		if err != nil {
			t.Errorf("Unexpected error: %s", err)
		}

		if len(list) != 1 {
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

	t.Run("check that proper list entries when a populated BatchToSEListEntry is passed", func(t *testing.T) {
		// setup
		err := seToCreateForBatchToSETest.Create()
		if err != nil {
			t.Errorf("Error setting up test: %s", err)
		}
		defer seToCreateForBatchToSETest.Delete()

		err = batchToSEToCreate2.Create()
		if err != nil {
			t.Errorf("Error setting up test: %s", err)
		}
		defer batchToSEToCreate2.Delete()

		// filter by BatchID
		list, err := BatchToSEList(BatchToSEListEntry{BatchID: "32626389-c942-4e71-9b5a-5d7c7ca4d389"})
		if err != nil {
			t.Errorf("Unexpected error: %s", err)
		}

		if len(list) != 2 {
			t.Errorf("Mismatch: want: \"%v\", got: \"%v\"", 2, len(list))
		}

		// filter by SEID
		list, err = BatchToSEList(BatchToSEListEntry{SEID: "8c258cb2-d700-43be-8773-a61a7b9cd668"})
		if err != nil {
			t.Errorf("Unexpected error: %s", err)
		}

		if len(list) != 1 {
			t.Errorf("Mismatch: want: \"%v\", got: \"%v\"", 1, len(list))
		}

		if list[0].SEID != "8c258cb2-d700-43be-8773-a61a7b9cd668" {
			t.Errorf("ID mismatch: want: \"8c258cb2-d700-43be-8773-a61a7b9cd668\", got: \"%v\"", list[0].SEID)
		}

		// filter by Phase
		list, err = BatchToSEList(BatchToSEListEntry{Phase: "baking"})
		if err != nil {
			t.Errorf("Unexpected error: %s", err)
		}

		if len(list) != 1 {
			t.Errorf("Mismatch: want: \"%v\", got: \"%v\"", 1, len(list))
		}

		if list[0].SEID != "cee91db3-ee73-4953-a05e-98f043d44f97" {
			t.Errorf("Status filter error. Got unexpected SEID: %v", list[0].SEID)
		}

		// filter by Step
		list, err = BatchToSEList(BatchToSEListEntry{Step: "trimming"})
		if err != nil {
			t.Errorf("Unexpected error: %s", err)
		}

		if len(list) != 1 {
			t.Errorf("Mismatch: want: \"%v\", got: \"%v\"", 1, len(list))
		}

		if list[0].SEID != "8c258cb2-d700-43be-8773-a61a7b9cd668" {
			t.Errorf("ID mismatch: want: \"8c258cb2-d700-43be-8773-a61a7b9cd668\", got: \"%v\"", list[0].SEID)
		}

		// filter by Status
		list, err = BatchToSEList(BatchToSEListEntry{Status: "covered-in-dust"})
		if err != nil {
			t.Errorf("Unexpected error: %s", err)
		}

		if len(list) != 1 {
			t.Errorf("Mismatch: want: \"%v\", got: \"%v\"", 1, len(list))
		}

		if list[0].SEID != "cee91db3-ee73-4953-a05e-98f043d44f97" {
			t.Errorf("Status filter error. Got unexpected SEID: %v", list[0].SEID)
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

		expect = "http://localhost:3000/api/v0/batch_to_ses"
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
