package rsbe

import (
	"net/http/httptest"
	"testing"

	"github.com/nyudlts/go-rsbe-client/rsbe/internal/testutils"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
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
		require.NoError(t, err, "Unexpected error: %s", err)
		assert.Len(t, got, 3, "Result Length Mismatch: want: 3, got: %d", len(got))
		assert.Equal(t, want.ID, got[0].ID, "ID Mismatch: want: \"%v\", got: \"%v\"", want.ID, got[0].ID)
		assert.Equal(t, want.Name, got[0].Name, "Name Mismatch: want: \"%v\", got: \"%v\"", want.Name, got[0].Name)
		assert.Equal(t, want.Size, got[0].Size, "Size Mismatch: want: \"%v\", got: \"%v\"", want.Size, got[0].Size)
		assert.Equal(t, want.Status, got[0].Status, "Status Mismatch: want: \"%v\", got: \"%v\"", want.Status, got[0].Status)
		assert.Equal(t, want.MTime, got[0].MTime, "MTime Mismatch: want: \"%v\", got: \"%v\"", want.MTime, got[0].MTime)
		assert.Equal(t, want.Data.Searchable, got[0].Data.Searchable, "Data.Searchable Mismatch: want: \"%v\", got: \"%v\"", want.Data.Searchable, got[0].Data.Searchable)
		assert.Equal(t, want.PartnerURL, got[0].PartnerURL, "PartnerURL Mismatch: want: \"%v\", got: \"%v\"", want.PartnerURL, got[0].PartnerURL)
		assert.Equal(t, want.CollectionURL, got[0].CollectionURL, "CollectionURL Mismatch: want: \"%v\", got: \"%v\"", want.CollectionURL, got[0].CollectionURL)
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
		require.NoError(t, err, "Unexpected error: %s", err)
		assert.Equal(t, want.ID, got.ID, "ID Mismatch: want: \"%v\", got: \"%v\"", want.ID, got.ID)
		assert.Equal(t, want.PartnerID, got.PartnerID, "PartnerID Mismatch: want: \"%v\", got: \"%v\"", want.PartnerID, got.PartnerID)
		assert.Equal(t, want.CollectionID, got.CollectionID, "CollectionID Mismatch: want: \"%v\", got: \"%v\"", want.CollectionID, got.CollectionID)
		assert.Equal(t, want.Size, got.Size, "Size Mismatch: want: \"%v\", got: \"%v\"", want.Size, got.Size)
		assert.Equal(t, want.Status, got.Status, "Status Mismatch: want: \"%v\", got: \"%v\"", want.Status, got.Status)
		assert.Equal(t, want.OriginalName, got.OriginalName, "OriginalName Mismatch: want: \"%v\", got: \"%v\"", want.OriginalName, got.OriginalName)
		assert.Equal(t, want.Name, got.Name, "Name Mismatch: want: \"%v\", got: \"%v\"", want.Name, got.Name)
		assert.Equal(t, want.Extension, got.Extension, "Extension Mismatch: want: \"%v\", got: \"%v\"", want.Extension, got.Extension)
		assert.Equal(t, want.MTime, got.MTime, "MTime Mismatch: want: \"%v\", got: \"%v\"", want.MTime, got.MTime)
		assert.Equal(t, want.HashMD5, got.HashMD5, "HashMD5 Mismatch: want: \"%v\", got: \"%v\"", want.HashMD5, got.HashMD5)
		assert.Equal(t, want.HashSHA1, got.HashSHA1, "HashSHA1 Mismatch: want: \"%v\", got: \"%v\"", want.HashSHA1, got.HashSHA1)
		assert.Equal(t, want.HashSHA256, got.HashSHA256, "HashSHA256 Mismatch: want: \"%v\", got: \"%v\"", want.HashSHA256, got.HashSHA256)
		assert.Equal(t, want.HashSHA512, got.HashSHA512, "HashSHA512 Mismatch: want: \"%v\", got: \"%v\"", want.HashSHA512, got.HashSHA512)
		assert.Equal(t, want.Formats.PRONOMID, got.Formats.PRONOMID, "Formats.PRONOMID Mismatch: want: \"%v\", got: \"%v\"", want.Formats.PRONOMID, got.Formats.PRONOMID)
		assert.Equal(t, want.Formats.MIMEType, got.Formats.MIMEType, "Formats.MIMEType Mismatch: want: \"%v\", got: \"%v\"", want.Formats.MIMEType, got.Formats.MIMEType)
		assert.Equal(t, want.Data.Searchable, got.Data.Searchable, "Data.Searchable Mismatch: want: \"%v\", got: \"%v\"", want.Data.Searchable, got.Data.Searchable)
		assert.Equal(t, want.Data.Duration, got.Data.Duration, "Data.Duration Mismatch: want: \"%v\", got: \"%v\"", want.Data.Duration, got.Data.Duration)
		assert.Equal(t, want.Data.Bitrate, got.Data.Bitrate, "Data.Bitrate Mismatch: want: \"%v\", got: \"%v\"", want.Data.Bitrate, got.Data.Bitrate)
		assert.Equal(t, want.Data.Width, got.Data.Width, "Data.Width Mismatch: want: \"%v\", got: \"%v\"", want.Data.Width, got.Data.Width)
		assert.Equal(t, want.Data.Height, got.Data.Height, "Data.Height Mismatch: want: \"%v\", got: \"%v\"", want.Data.Height, got.Data.Height)
		assert.Equal(t, want.Data.AspectRatio, got.Data.AspectRatio, "Data.AspectRatio Mismatch: want: \"%v\", got: \"%v\"", want.Data.AspectRatio, got.Data.AspectRatio)
		assert.Equal(t, want.Data.XMLSchema, got.Data.XMLSchema, "Data.XMLSchema Mismatch: want: \"%v\", got: \"%v\"", want.Data.XMLSchema, got.Data.XMLSchema)
		assert.Equal(t, want.Data.TranscriptionID, got.Data.TranscriptionID, "Data.TranscriptionID Mismatch: want: \"%v\", got: \"%v\"", want.Data.TranscriptionID, got.Data.TranscriptionID)
		assert.Equal(t, want.PartnerURL, got.PartnerURL, "PartnerURL Mismatch: want: \"%v\", got: \"%v\"", want.PartnerURL, got.PartnerURL)
		assert.Equal(t, want.CollectionURL, got.CollectionURL, "CollectionURL Mismatch: want: \"%v\", got: \"%v\"", want.CollectionURL, got.CollectionURL)
		assert.Equal(t, want.LockVersion, got.LockVersion, "LockVersion Mismatch: want: \"%v\", got: \"%v\"", want.LockVersion, got.LockVersion)

		testutils.AssertEquivalentTimestamps(t, want.CreatedAt, got.CreatedAt)
		testutils.AssertEquivalentTimestamps(t, want.UpdatedAt, got.UpdatedAt)

	})

}

func TestFMDCreateFunc(t *testing.T) {
	setupLocalhostClient()

	err := fmdToCreate.Create()
	if err != nil {
		t.Errorf("Unexpected error: %s", err)
	}
	defer fmdToCreate.Purge()

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

		if fmdToCreate.CollectionURL == "" {
			t.Errorf("CollectionURL mismatch: want a non-empty value, got: \"%v\"", fmdToCreate.CollectionURL)
		}

		if fmdToCreate.PartnerURL == "" {
			t.Errorf("PartnerURL mismatch: want a non-empty value, got: \"%v\"", fmdToCreate.PartnerURL)
		}
	})

	t.Run("confirm that all fields are as expected", func(t *testing.T) {
		want := fmdToCreate
		got, err := FMDGet(fmdToCreate.ID)
		if err != nil {
			t.Errorf("Unexpected error: %s", err)
		}

		compareFMDValues(t, got, want, false)
	})
}

func TestFMDUpdateFunc(t *testing.T) {
	setupLocalhostClient()

	err := fmdToCreate.Create()
	if err != nil {
		t.Errorf("Unexpected error: %s", err)
	}
	defer fmdToCreate.Purge()

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

		compareFMDValues(t, got, want, true)
	})
}

func TestFMDDeleteFunc(t *testing.T) {
	setupLocalhostClient()

	err := fmdToCreate.Create()
	if err != nil {
		t.Errorf("Unexpected error: %s", err)
	}
	// the Purge() call is needed because Delete() is a soft delete,
	// the record will still exist and needs to be purged
	defer fmdToCreate.Purge()

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

// ------------------------------------------------------------------------------
// Helper Functions
// ------------------------------------------------------------------------------
func compareFMDValues(t *testing.T, got, want FMDEntry, updated bool) {

	assert.Equal(t, want.ID, got.ID, "ID mismatch")
	assert.Equal(t, want.PartnerID, got.PartnerID, "PartnerID mismatch")
	assert.Equal(t, want.CollectionID, got.CollectionID, "CollectionID mismatch")
	assert.Equal(t, want.XIPID, got.XIPID, "XIPID mismatch")
	assert.Equal(t, want.PresLevel, got.PresLevel, "PresLevel mismatch")
	assert.Equal(t, want.PresCommitment, got.PresCommitment, "PresCommitment mismatch")
	assert.Equal(t, want.Size, got.Size, "Size mismatch")
	assert.Equal(t, want.Status, got.Status, "Status mismatch")
	assert.Equal(t, want.FormatAcceptable, got.FormatAcceptable, "FormatAcceptable mismatch")
	assert.Equal(t, want.FormatValid, got.FormatValid, "FormatValid mismatch")
	assert.Equal(t, want.OriginalName, got.OriginalName, "OriginalName mismatch")
	assert.Equal(t, want.Name, got.Name, "Name mismatch")
	assert.Equal(t, want.Extension, got.Extension, "Extension mismatch")
	testutils.AssertEquivalentTimestamps(t, want.MTime, got.MTime, "MTime mismatch")
	assert.Equal(t, want.HashMD5, got.HashMD5, "HashMD5 mismatch")
	assert.Equal(t, want.HashSHA1, got.HashSHA1, "HashSHA1 mismatch")
	assert.Equal(t, want.HashSHA256, got.HashSHA256, "HashSHA256 mismatch")
	assert.Equal(t, want.HashSHA512, got.HashSHA512, "HashSHA512 mismatch")
	assert.Equal(t, want.Formats.MIMEType, got.Formats.MIMEType, "Formats.MIMEType mismatch")
	assert.Equal(t, want.Formats.PRONOMID, got.Formats.PRONOMID, "Formats.PRONOMID Mismatch")
	assert.Equal(t, want.Data.Bitrate, got.Data.Bitrate, "Data.Bitrate mismatch")
	assert.Equal(t, want.Data.Width, got.Data.Width, "Data.Width mismatch")
	assert.Equal(t, want.Data.Height, got.Data.Height, "Data.Height mismatch")
	assert.Equal(t, want.Data.AspectRatio, got.Data.AspectRatio, "Data.AspectRatio mismatch")
	assert.Equal(t, want.Data.XMLSchema, got.Data.XMLSchema, "Data.XMLSchema mismatch")
	assert.Equal(t, want.Data.TranscriptionID, got.Data.TranscriptionID, "Data.TranscriptionID Mismatch")
	assert.Equal(t, want.PartnerURL, got.PartnerURL, "PartnerURL mismatch")
	assert.Equal(t, want.CollectionURL, got.CollectionURL, "CollectionURL mismatch")

	testutils.AssertEquivalentTimestamps(t, want.CreatedAt, got.CreatedAt)
	if updated {
		testutils.AssertNotEquivalentTimestamps(t, want.UpdatedAt, got.UpdatedAt, "Expected UpdatedAt to have changed")
		assert.Equal(t, want.LockVersion+1, got.LockVersion, "LockVersion mismatch")
	} else {
		testutils.AssertEquivalentTimestamps(t, want.UpdatedAt, got.UpdatedAt, "Expected UpdatedAt to be unchanged")
		assert.Equal(t, want.LockVersion, got.LockVersion, "LockVersion mismatch")
	}

}
