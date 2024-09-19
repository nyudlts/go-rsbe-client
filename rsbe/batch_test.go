package rsbe

import (
	"net/http/httptest"
	"sort"
	"testing"
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

		if len(list) != 2 {
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
		got := report.TimeStamp
		if want != got {
			t.Errorf("Mismatch: want: \"%v\", got: \"%v\"", want, got)
		}

		want = "c44e95e9-5cca-4c26-8e52-12773334dc95"
		got = report.Info.ID
		if want != got {
			t.Errorf("Mismatch: want: \"%v\", got: \"%v\"", want, got)
		}

		want = "publication"
		got = report.Info.Type
		if want != got {
			t.Errorf("Mismatch: want: \"%v\", got: \"%v\"", want, got)
		}

		wantUInt := uint(1)
		gotUInt := report.Info.Number
		if wantUInt != gotUInt {
			t.Errorf("Mismatch: want: \"%v\", got: \"%v\"", wantUInt, gotUInt)
		}

		want = "HIDVL HTML5 Migration"
		got = report.Info.Name
		if want != got {
			t.Errorf("Mismatch: want: \"%v\", got: \"%v\"", want, got)
		}

		want = "2020-10-18T02:52:27.827Z"
		got = report.Info.CreatedAt
		if want != got {
			t.Errorf("Mismatch: want: \"%v\", got: \"%v\"", want, got)
		}

		want = "2020-11-13T05:03:27.638Z"
		got = report.Info.UpdatedAt
		if want != got {
			t.Errorf("Mismatch: want: \"%v\", got: \"%v\"", want, got)
		}

		want = "https://rsbe.dlib.nyu.edu/api/v0/batches/c44e95e9-5cca-4c26-8e52-12773334dc95"
		got = report.Info.URL
		if want != got {
			t.Errorf("Mismatch: want: \"%v\", got: \"%v\"", want, got)
		}

		wantUInt = 1140
		gotUInt = report.Info.Stats.Total
		if wantUInt != gotUInt {
			t.Errorf("Mismatch: want: \"%v\", got: \"%v\"", wantUInt, gotUInt)
		}

		wantUInt = 1139
		gotUInt = report.Info.Stats.SEStats.Total
		if wantUInt != gotUInt {
			t.Errorf("Mismatch: want: \"%v\", got: \"%v\"", wantUInt, gotUInt)
		}

		want = "programmatic_edit/qc/active"
		got = report.Info.Stats.SEStats.Groups[3].PhaseStepStatus
		if want != got {
			t.Errorf("Mismatch: want: \"%v\", got: \"%v\"", want, got)
		}

		wantUInt = 218
		gotUInt = report.Info.Stats.SEStats.Groups[3].Count
		if wantUInt != gotUInt {
			t.Errorf("Mismatch: want: \"%v\", got: \"%v\"", wantUInt, gotUInt)
		}

		want = "https://rsbe.dlib.nyu.edu/api/v0/batch_to_ses?batch_id=c44e95e9-5cca-4c26-8e52-12773334dc95&phase=programmatic_edit&status=active&step=qc"
		got = report.Info.Stats.SEStats.Groups[3].URL
		if want != got {
			t.Errorf("Mismatch: want: \"%v\", got: \"%v\"", want, got)
		}

		want = "1ns1rppm"
		got = report.SEs[4].DigiID
		if want != got {
			t.Errorf("Mismatch: want: \"%v\", got: \"%v\"", want, got)
		}

		want = "ip_prep"
		got = report.SEs[4].Phase
		if want != got {
			t.Errorf("Mismatch: want: \"%v\", got: \"%v\"", want, got)
		}

		want = "derivative_generation"
		got = report.SEs[4].Step
		if want != got {
			t.Errorf("Mismatch: want: \"%v\", got: \"%v\"", want, got)
		}

		want = "active"
		got = report.SEs[4].Status
		if want != got {
			t.Errorf("Mismatch: want: \"%v\", got: \"%v\"", want, got)
		}

		want = "2020-11-13T15:01:16.509Z"
		got = report.SEs[4].CreatedAt
		if want != got {
			t.Errorf("Mismatch: want: \"%v\", got: \"%v\"", want, got)
		}

		want = "2020-11-25T22:30:11.034Z"
		got = report.SEs[4].UpdatedAt
		if want != got {
			t.Errorf("Mismatch: want: \"%v\", got: \"%v\"", want, got)
		}

		want = "https://rsbe.dlib.nyu.edu/api/v0/ses/72b8f70b-9a69-446d-9f22-05460252e07f"
		got = report.SEs[4].SEURL
		if want != got {
			t.Errorf("Mismatch: want: \"%v\", got: \"%v\"", want, got)
		}

		wantUInt = 1
		gotUInt = report.Info.Stats.IEStats.Total
		if wantUInt != gotUInt {
			t.Errorf("Mismatch: want: \"%v\", got: \"%v\"", wantUInt, gotUInt)
		}

		want = "003348201"
		got = report.IEs[0].SysNum
		if want != got {
			t.Errorf("Mismatch: want: \"%v\", got: \"%v\"", want, got)
		}

		want = "test title"
		got = report.IEs[0].Title
		if want != got {
			t.Errorf("Mismatch: want: \"%v\", got: \"%v\"", want, got)
		}

		want = "rework"
		got = report.IEs[0].Phase
		if want != got {
			t.Errorf("Mismatch: want: \"%v\", got: \"%v\"", want, got)
		}

		want = "content_staging"
		got = report.IEs[0].Step
		if want != got {
			t.Errorf("Mismatch: want: \"%v\", got: \"%v\"", want, got)
		}

		want = "queued"
		got = report.IEs[0].Status
		if want != got {
			t.Errorf("Mismatch: want: \"%v\", got: \"%v\"", want, got)
		}

		want = "2020-11-13T18:30:34.172Z"
		got = report.IEs[0].CreatedAt
		if want != got {
			t.Errorf("Mismatch: want: \"%v\", got: \"%v\"", want, got)
		}

		want = "2020-11-21T17:48:07.129Z"
		got = report.IEs[0].UpdatedAt
		if want != got {
			t.Errorf("Mismatch: want: \"%v\", got: \"%v\"", want, got)
		}

		want = "https://rsbe.dlib.nyu.edu/api/v0/ies/dcb20119-272d-46e5-8677-e07d14b964bc"
		got = report.IEs[0].IEURL
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
