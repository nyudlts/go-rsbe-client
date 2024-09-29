package rsbe

import (
	"net/http/httptest"
	"testing"
)

var fmdListEntry = FMDListEntry{
	ID:     "4a3f8f8c-6dbe-4d7c-bff1-1b973f9f615c",
	Name:   "foo.pdf",
	Size:   1111,
	Status: "ok",
	MTime:  "2020-05-30T02:21:59.710Z",
	Data: FMDData{
		Searchable: true,
	},
	URL:           "http://localhost:3000/api/v0/fmds/4a3f8f8c-6dbe-4d7c-bff1-1b973f9f615c",
	PartnerURL:    "http://localhost:3000/api/v0/partners/e6517775-6277-4e25-9373-ee7738e820b5",
	CollectionURL: "http://localhost:3000/api/v0/colls/b9612d5d-619a-4ceb-b620-d816e4b4340b",
}

var fmdShow = FMDEntry{
	ID:           "221a87ad-99d4-4c61-9dda-e78895755e05",
	PartnerID:    "e6517775-6277-4e25-9373-ee7738e820b5",
	CollectionID: "b9612d5d-619a-4ceb-b620-d816e4b4340b",
	Size:         12342,
	Status:       "ok",
	OriginalName: "maple.pdf",
	Name:         "syrup.pdf",
	Extension:    "pdf",
	MTime:        "2020-06-30T02:21:59.710Z",
	HashMD5:      "6a6735088d582e2b9867542759988d3c",
	HashSHA1:     "7adfb08560ea47856db668fda00276796404a7dc",
	HashSHA256:   "57cb4643e48bdaf4aad877cbd1a5401341207964bbc3195cd798e34ce69f37fb",
	HashSHA512:   "e21baae6bac92cd46cb3fb7d1117d529ee8c3d80f6e1a7c84ee599bc14bb7cd6c538c9161f75bd9d24f1ce714a9c422bedf55a132fb070e0c7a112316bfbc267",
	CreatedAt:    "2020-07-13T02:13:10.297Z",
	UpdatedAt:    "2020-07-13T02:13:10.297Z",
	Formats: FMDFormat{
		PRONOMID: "fmt/14",
		MIMEType: "application/pdf",
	},
	Data: FMDData{
		Searchable:      true,
		Duration:        "00:01:23.456",
		Bitrate:         800000,
		Width:           1920,
		Height:          1080,
		AspectRatio:     "16:9",
		XMLSchema:       "marcxml",
		TranscriptionID: "cd165a2f-f976-4c55-a63c-6b57017eed49",
	},
	CollectionURL: "http://localhost:3000/api/v0/colls/b9612d5d-619a-4ceb-b620-d816e4b4340b",
	PartnerURL:    "http://localhost:3000/api/v0/partners/e6517775-6277-4e25-9373-ee7738e820b5",
	LockVersion:   0,
}

var fmdToCreate = FMDEntry{
	ID:               "64f2d78a-613e-4274-a120-9d984b5ec09a",
	PartnerID:        "e6517775-6277-4e25-9373-ee7738e820b5",
	CollectionID:     "b9612d5d-619a-4ceb-b620-d816e4b4340b",
	XIPID:            "92f51b59-4589-4730-b861-4ca9500f9480",
	Size:             12342,
	Status:           "ok",
	FormatAcceptable: true,
	FormatValid:      true,
	OriginalName:     "maple.pdf",
	Name:             "syrup.pdf",
	Extension:        "pdf",
	MTime:            "2020-06-30T02:21:59.710Z",
	HashMD5:          "6a6735088d582e2b9867542759988d3c",
	HashSHA1:         "7adfb08560ea47856db668fda00276796404a7dc",
	HashSHA256:       "57cb4643e48bdaf4aad877cbd1a5401341207964bbc3195cd798e34ce69f37fb",
	HashSHA512:       "e21baae6bac92cd46cb3fb7d1117d529ee8c3d80f6e1a7c84ee599bc14bb7cd6c538c9161f75bd9d24f1ce714a9c422bedf55a132fb070e0c7a112316bfbc267",
	Formats: FMDFormat{
		PRONOMID: "fmt/14",
	},
	Data: FMDData{
		Searchable: true,
	},
}

func TestSEFMDList(t *testing.T) {

	mux := setupMux("/api/v0/ses/8c258cb2-d700-43be-8773-a61a7b9cd668/fmds", "testdata/se-fmd-list.json")
	ts := httptest.NewServer(mux)
	defer ts.Close()

	setupTestServerClient(ts)

	t.Run("result", func(t *testing.T) {
		want := fmdListEntry
		got, err := SEFMDList("8c258cb2-d700-43be-8773-a61a7b9cd668")
		if err != nil {
			t.Errorf("Unexpected error: %s", err)
		}

		if len(got) != 3 {
			t.Errorf("Result Length Mismatch: want: 3, got: %d", len(got))
		}

		if fmdListEntry.ID != got[0].ID {
			t.Errorf("Mismatch: want: \"%v\", got: \"%v\"", want, got)
		}

		if fmdListEntry.Name != got[0].Name {
			t.Errorf("Mismatch: want: \"%v\", got: \"%v\"", want, got)
		}

		if fmdListEntry.Size != got[0].Size {
			t.Errorf("Mismatch: want: \"%v\", got: \"%v\"", want, got)
		}

		if fmdListEntry.Status != got[0].Status {
			t.Errorf("Mismatch: want: \"%v\", got: \"%v\"", want, got)
		}

		if fmdListEntry.MTime != got[0].MTime {
			t.Errorf("Mismatch: want: \"%v\", got: \"%v\"", want, got)
		}

		if fmdListEntry.Data.Searchable != got[0].Data.Searchable {
			t.Errorf("Mismatch: want: \"%v\", got: \"%v\"", want, got)
		}

		if fmdListEntry.PartnerURL != got[0].PartnerURL {
			t.Errorf("Mismatch: want: \"%v\", got: \"%v\"", want, got)
		}

		if fmdListEntry.CollectionURL != got[0].CollectionURL {
			t.Errorf("Mismatch: want: \"%v\", got: \"%v\"", want, got)
		}
	})

}

func TestFMDGetFunc(t *testing.T) {

	mux := setupMux("/api/v0/fmds/4a3f8f8c-6dbe-4d7c-bff1-1b973f9f615c", "testdata/fmd-get.json")
	ts := httptest.NewServer(mux)
	defer ts.Close()

	setupTestServerClient(ts)

	t.Run("result", func(t *testing.T) {
		want := fmdShow
		got := FMDEntry{ID: "4a3f8f8c-6dbe-4d7c-bff1-1b973f9f615c"}

		err := got.Get()
		if err != nil {
			t.Errorf("Unexpected error: %s", err)
		}

		if want.ID != got.ID {
			t.Errorf("ID mismatch: want: \"%v\", got: \"%v\"", want.ID, got.ID)
		}

		if want.PartnerID != got.PartnerID {
			t.Errorf("PartnerID mismatch: want: \"%v\", got: \"%v\"", want.PartnerID, got.PartnerID)
		}

		if want.CollectionID != got.CollectionID {
			t.Errorf("CollectionID mismatch: want: \"%v\", got: \"%v\"", want.CollectionID, got.CollectionID)
		}

		if want.Size != got.Size {
			t.Errorf("Size mismatch: want: \"%v\", got: \"%v\"", want.Size, got.Size)
		}

		if want.Status != got.Status {
			t.Errorf("Status mismatch: want: \"%v\", got: \"%v\"", want.Status, got.Status)
		}

		if want.OriginalName != got.OriginalName {
			t.Errorf("OriginalName mismatch: want: \"%v\", got: \"%v\"", want.OriginalName, got.OriginalName)
		}

		if want.Name != got.Name {
			t.Errorf("Name mismatch: want: \"%v\", got: \"%v\"", want.Name, got.Name)
		}

		if want.Extension != got.Extension {
			t.Errorf("Extension mismatch: want: \"%v\", got: \"%v\"", want.Extension, got.Extension)
		}

		if want.MTime != got.MTime {
			t.Errorf("MTime mismatch: want: \"%v\", got: \"%v\"", want.MTime, got.MTime)
		}

		if want.CreatedAt != got.CreatedAt {
			t.Errorf("CreatedAt mismatch: want: \"%v\", got: \"%v\"", want.CreatedAt, got.CreatedAt)
		}

		if want.UpdatedAt != got.UpdatedAt {
			t.Errorf("UpdatedAt mismatch: want: \"%v\", got: \"%v\"", want.UpdatedAt, got.UpdatedAt)
		}

		if want.Formats.PRONOMID != got.Formats.PRONOMID {
			t.Errorf("Formats.PRONOMID Mismatch: want: \"%v\", got: \"%v\"", want.Formats.PRONOMID, got.Formats.PRONOMID)
		}

		if want.Formats.MIMEType != got.Formats.MIMEType {
			t.Errorf("Formats.MIMEType mismatch: want: \"%v\", got: \"%v\"", "foo", got.Formats.MIMEType)
		}

		if want.Data.Duration != got.Data.Duration {
			t.Errorf("Data.Duration mismatch: want: \"%v\", got: \"%v\"", want.Data.Duration, got.Data.Duration)
		}

		if want.Data.Bitrate != got.Data.Bitrate {
			t.Errorf("Data.Bitrate mismatch: want: \"%v\", got: \"%v\"", want.Data.Bitrate, got.Data.Bitrate)
		}

		if want.Data.Width != got.Data.Width {
			t.Errorf("Data.Width mismatch: want: \"%v\", got: \"%v\"", want.Data.Width, got.Data.Width)
		}

		if want.Data.Height != got.Data.Height {
			t.Errorf("Data.Height mismatch: want: \"%v\", got: \"%v\"", want.Data.Height, got.Data.Height)
		}

		if want.Data.AspectRatio != got.Data.AspectRatio {
			t.Errorf("Data.AspectRatio mismatch: want: \"%v\", got: \"%v\"", want.Data.AspectRatio, got.Data.AspectRatio)
		}

		if want.Data.XMLSchema != got.Data.XMLSchema {
			t.Errorf("Data.XMLSchema mismatch: want: \"%v\", got: \"%v\"", want.Data.XMLSchema, got.Data.XMLSchema)
		}

		if want.Data.TranscriptionID != got.Data.TranscriptionID {
			t.Errorf("Data.TranscriptionID Mismatch: want: \"%v\", got: \"%v\"", want.Data.TranscriptionID, got.Data.TranscriptionID)
		}

		if want.PartnerURL != got.PartnerURL {
			t.Errorf("PartnerURL mismatch: want: \"%v\", got: \"%v\"", want.PartnerURL, got.PartnerURL)
		}

		if want.CollectionURL != got.CollectionURL {
			t.Errorf("CollectionURL mismatch: want: \"%v\", got: \"%v\"", want.CollectionURL, got.CollectionURL)
		}

		if want.LockVersion != got.LockVersion {
			t.Errorf("LockVersion mismatch: want: \"%v\", got: \"%v\"", want.LockVersion, got.LockVersion)
		}

	})

}

func TestFMDCreateFunc(t *testing.T) {
	setupLocalhostClient()

	err := fmdToCreate.Create()
	if err != nil {
		t.Errorf("Unexpected error: %s", err)
	}
	defer fmdToCreate.Delete()

	t.Run("confirm that attributes updated", func(t *testing.T) {
		if fmdToCreate.ID == "" {
			t.Errorf("ID not updated")
		}

		if fmdToCreate.CreatedAt == "" {
			t.Errorf("CreatedAt not updated")
		}

		if fmdToCreate.UpdatedAt == "" {
			t.Errorf("UpdatedAt not updated")
		}

		if fmdToCreate.CollectionURL != "" {
			t.Errorf("CollectionURL mismatch: want a non-empty value, got: \"%v\"", fmdToCreate.CollectionURL)
		}

		if fmdToCreate.PartnerURL != "" {
			t.Errorf("PartnerURL mismatch: want a non-empty value, got: \"%v\"", fmdToCreate.PartnerURL)
		}
	})

	t.Run("confirm that all fields are as expected", func(t *testing.T) {
		want := fmdToCreate
		got, err := FMDGet(fmdToCreate.ID)
		if err != nil {
			t.Errorf("Unexpected error: %s", err)
		}

		checkStringContains(t, got.ID, want.ID)
		checkStringContains(t, got.PartnerID, want.PartnerID)
		checkStringContains(t, got.CollectionID, want.CollectionID)
		if want.Size != got.Size {
			t.Errorf("Size mismatch: want: \"%v\", got: \"%v\"", want.Size, got.Size)
		}
		checkStringContains(t, got.Status, want.Status)
		checkStringContains(t, got.OriginalName, want.OriginalName)
		checkStringContains(t, got.Name, want.Name)
		checkStringContains(t, got.Extension, want.Extension)
		checkStringContains(t, got.MTime, want.MTime)
		checkStringContains(t, got.CreatedAt, want.CreatedAt)
		checkStringContains(t, got.UpdatedAt, want.UpdatedAt)
		checkStringContains(t, got.Formats.MIMEType, want.Formats.MIMEType)
		checkStringContains(t, got.Formats.PRONOMID, want.Formats.PRONOMID)

		// if want.MTime != got.MTime {
		// 	t.Errorf("MTime mismatch: want: \"%v\", got: \"%v\"", want.MTime, got.MTime)
		// }

		// if want.CreatedAt != got.CreatedAt {
		// 	t.Errorf("CreatedAt mismatch: want: \"%v\", got: \"%v\"", want.CreatedAt, got.CreatedAt)
		// }

		// if want.UpdatedAt != got.UpdatedAt {
		// 	t.Errorf("UpdatedAt mismatch: want: \"%v\", got: \"%v\"", want.UpdatedAt, got.UpdatedAt)
		// }

		// if want.Formats.PRONOMID != got.Formats.PRONOMID {
		// 	t.Errorf("Formats.PRONOMID Mismatch: want: \"%v\", got: \"%v\"", want.Formats.PRONOMID, got.Formats.PRONOMID)
		// }

		// if want.Formats.MIMEType != got.Formats.MIMEType {
		// 	t.Errorf("Formats.MIMEType mismatch: want: \"%v\", got: \"%v\"", "foo", got.Formats.MIMEType)
		// }

		// if want.Data.Duration != got.Data.Duration {
		// 	t.Errorf("Data.Duration mismatch: want: \"%v\", got: \"%v\"", want.Data.Duration, got.Data.Duration)
		// }

		if want.Data.Bitrate != got.Data.Bitrate {
			t.Errorf("Data.Bitrate mismatch: want: \"%v\", got: \"%v\"", want.Data.Bitrate, got.Data.Bitrate)
		}

		if want.Data.Width != got.Data.Width {
			t.Errorf("Data.Width mismatch: want: \"%v\", got: \"%v\"", want.Data.Width, got.Data.Width)
		}

		if want.Data.Height != got.Data.Height {
			t.Errorf("Data.Height mismatch: want: \"%v\", got: \"%v\"", want.Data.Height, got.Data.Height)
		}

		if want.Data.AspectRatio != got.Data.AspectRatio {
			t.Errorf("Data.AspectRatio mismatch: want: \"%v\", got: \"%v\"", want.Data.AspectRatio, got.Data.AspectRatio)
		}

		if want.Data.XMLSchema != got.Data.XMLSchema {
			t.Errorf("Data.XMLSchema mismatch: want: \"%v\", got: \"%v\"", want.Data.XMLSchema, got.Data.XMLSchema)
		}

		if want.Data.TranscriptionID != got.Data.TranscriptionID {
			t.Errorf("Data.TranscriptionID Mismatch: want: \"%v\", got: \"%v\"", want.Data.TranscriptionID, got.Data.TranscriptionID)
		}

		if want.LockVersion != got.LockVersion {
			t.Errorf("LockVersion mismatch: want: \"%v\", got: \"%v\"", want.LockVersion, got.LockVersion)
		}
	})
}

func TestFMDUpdateFunc(t *testing.T) {
	setupLocalhostClient()

	err := fmdToCreate.Create()
	if err != nil {
		t.Errorf("Unexpected error: %s", err)
	}
	defer fmdToCreate.Delete()

	err = fmdToCreate.Get()
	if err != nil {
		t.Errorf("Unexpected error: %s", err)
	}

	fmdToCreate.XIPID = "a240d0df-4aa1-4a8d-b62f-8698443e2953"
	fmdToCreate.Size = 0
	fmdToCreate.PresLevel = "foo"
	fmdToCreate.PresCommitment = "bar"
	fmdToCreate.Status = "error"
	fmdToCreate.FormatAcceptable = false
	fmdToCreate.FormatValid = false
	fmdToCreate.OriginalName = "banana.bin"
	fmdToCreate.Name = "plantain.bin"
	fmdToCreate.Extension = "bin"
	fmdToCreate.MTime = "2020-06-30T02:21:59.710Z"
	fmdToCreate.HashMD5 = "0d4b08e85f5a5bd7211fa7f548bfed88"
	fmdToCreate.HashSHA1 = "8ea7303278ca121924eaa6b84b4c424a40b74307"
	fmdToCreate.HashSHA256 = "5f713dcea42922c925d7ebee97f14a48d27631000d09e3245a163db16edc27c2"
	fmdToCreate.HashSHA512 = "adfec7b531c9b676ca7a428d115dbe8a7741b1e99ec28a5c35ec52111ef023ef931d73522ff2db5226291243a7b02379e61bc807bbcadd761f14eb4460529eb3"
	fmdToCreate.Formats.PRONOMID = "fmt/99"
	fmdToCreate.Data.Bitrate = 123456
	fmdToCreate.Data.Width = 1920
	fmdToCreate.Data.Height = 1080
	fmdToCreate.Data.AspectRatio = "16:9"
	fmdToCreate.Data.XMLSchema = "marcxml"
	fmdToCreate.Data.TranscriptionID = "cd165a2f-f976-4c55-a63c-6b57017eed49"

	err = fmdToCreate.Update()
	if err != nil {
		t.Errorf("Unexpected error: %s", err)
	}

	t.Run("confirm that elements updated", func(t *testing.T) {
		want := fmdToCreate
		got, err := FMDGet(fmdToCreate.ID)
		if err != nil {
			t.Errorf("Unexpected error: %s", err)
		}

		checkStringContains(t, got.ID, want.ID)
		checkStringContains(t, got.PartnerID, want.PartnerID)
		checkStringContains(t, got.CollectionID, want.CollectionID)
		checkStringContains(t, got.XIPID, want.XIPID)
		if got.Size != want.Size {
			t.Errorf("Size mismatch: want: \"%v\", got: \"%v\"", want.Size, got.Size)
		}
		checkStringContains(t, got.PresLevel, want.PresLevel)
		checkStringContains(t, got.PresCommitment, want.PresCommitment)
		checkStringContains(t, got.Status, want.Status)
		if got.FormatValid != want.FormatValid {
			t.Errorf("FormatValid mismatch: want: \"%v\", got: \"%v\"", want.FormatValid, got.FormatValid)
		}
		if got.FormatAcceptable != want.FormatAcceptable {
			t.Errorf("FormatAcceptable mismatch: want: \"%v\", got: \"%v\"", want.FormatAcceptable, got.FormatAcceptable)
		}
		checkStringContains(t, got.OriginalName, want.OriginalName)
		checkStringContains(t, got.Name, want.Name)
		checkStringContains(t, got.Extension, want.Extension)
		checkStringContains(t, got.MTime, want.MTime)
		checkStringContains(t, got.HashMD5, want.HashMD5)
		checkStringContains(t, got.HashSHA1, want.HashSHA1)
		checkStringContains(t, got.HashSHA256, want.HashSHA256)
		checkStringContains(t, got.HashSHA512, want.HashSHA512)
		checkStringContains(t, got.Formats.MIMEType, want.Formats.MIMEType)
		checkStringContains(t, got.Formats.PRONOMID, want.Formats.PRONOMID)
		checkStringContains(t, got.CreatedAt, want.CreatedAt)

		if got.UpdatedAt == want.UpdatedAt {
			t.Errorf("UpdatedAt was not changed: want: \"%v\", got: \"%v\"", want.UpdatedAt, got.UpdatedAt)
		}

		if got.Data.Bitrate != want.Data.Bitrate {
			t.Errorf("Data.Bitrate mismatch: want: \"%v\", got: \"%v\"", want.Data.Bitrate, got.Data.Bitrate)
		}

		if got.Data.Width != want.Data.Width {
			t.Errorf("Data.Width mismatch: want: \"%v\", got: \"%v\"", want.Data.Width, got.Data.Width)
		}

		if got.Data.Height != want.Data.Height {
			t.Errorf("Data.Height mismatch: want: \"%v\", got: \"%v\"", want.Data.Height, got.Data.Height)
		}

		if got.Data.AspectRatio != want.Data.AspectRatio {
			t.Errorf("Data.AspectRatio mismatch: want: \"%v\", got: \"%v\"", want.Data.AspectRatio, got.Data.AspectRatio)
		}

		if got.Data.XMLSchema != want.Data.XMLSchema {
			t.Errorf("Data.XMLSchema mismatch: want: \"%v\", got: \"%v\"", want.Data.XMLSchema, got.Data.XMLSchema)
		}

		if got.Data.TranscriptionID != want.Data.TranscriptionID {
			t.Errorf("Data.TranscriptionID Mismatch: want: \"%v\", got: \"%v\"", want.Data.TranscriptionID, got.Data.TranscriptionID)
		}

		if got.LockVersion != want.LockVersion+1 {
			t.Errorf("LockVersion mismatch: want: \"%v\", got: \"%v\"", want.LockVersion, got.LockVersion)
		}
	})
}

func TestFMDDeleteFunc(t *testing.T) {
	setupLocalhostClient()

	err := fmdToCreate.Create()
	if err != nil {
		t.Errorf("Unexpected error: %s", err)
	}

	err = fmdToCreate.Get()
	if err != nil {
		t.Errorf("Unexpected error: %s", err)
	}

	id := fmdToCreate.ID

	err = fmdToCreate.Delete()
	if err != nil {
		t.Errorf("Unexpected error: %s", err)
	}

	t.Run("confirm that deleted item not found", func(t *testing.T) {
		// should not be found, so err should NOT be nil
		_, err = FMDGet(id)

		if err == nil {
			t.Errorf("err was nil")
		}

	})
}
