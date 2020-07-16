package rsbe

import (
	"net/http/httptest"
	"testing"
)

var etofmdListEntry = EtoFMDListEntry{
	ID:    "3ca8ecaf-6fae-48a5-8441-5a96e119ad28",
	EType: "se",
	EID:   "8c258cb2-d700-43be-8773-a61a7b9cd668",
	FMDID: "4a3f8f8c-6dbe-4d7c-bff1-1b973f9f615c",
	Role:  "master",
	URL:   "http://localhost:3000/api/v0/etofmds/3ca8ecaf-6fae-48a5-8441-5a96e119ad28",
}

var etofmdShow = EtoFMDEntry{
	ID:          "3ca8ecaf-6fae-48a5-8441-5a96e119ad28",
	EType:       "se",
	EID:         "8c258cb2-d700-43be-8773-a61a7b9cd668",
	Role:        "master",
	FMDID:       "4a3f8f8c-6dbe-4d7c-bff1-1b973f9f615c",
	CreatedAt:   "2020-05-31T20:37:30.747Z",
	UpdatedAt:   "2020-05-31T20:37:30.747Z",
	EURL:        "http://localhost:3000/api/v0/ses/8c258cb2-d700-43be-8773-a61a7b9cd668",
	LockVersion: 0,
}

// var fmdToCreate = FMDEntry{
// 	PartnerID:    "e6517775-6277-4e25-9373-ee7738e820b5",
// 	CollectionID: "b9612d5d-619a-4ceb-b620-d816e4b4340b",
// 	Size:         12342,
// 	Status:       "ok",
// 	OriginalName: "maple.pdf",
// 	Name:         "syrup.pdf",
// 	Extension:    "pdf",
// 	FileMTime:    "2020-06-30T02:21:59.710Z",
// 	HashMD5:      "6a6735088d582e2b9867542759988d3c",
// 	HashSHA1:     "7adfb08560ea47856db668fda00276796404a7dc",
// 	HashSHA256:   "57cb4643e48bdaf4aad877cbd1a5401341207964bbc3195cd798e34ce69f37fb",
// 	HashSHA512:   "e21baae6bac92cd46cb3fb7d1117d529ee8c3d80f6e1a7c84ee599bc14bb7cd6c538c9161f75bd9d24f1ce714a9c422bedf55a132fb070e0c7a112316bfbc267",
// 	Formats: FMDFormat{
// 		PRONOM: "fmt/14",
// 	},
// 	Data: FMDData{
// 		Searchable: true,
// 	},
// }

func TestEtoFMDsList(t *testing.T) {

	mux := setupMux("/api/v0/etofmds", "testdata/etofmd-list.json")
	ts := httptest.NewServer(mux)
	defer ts.Close()

	setupTestServerClient(ts)

	t.Run("result", func(t *testing.T) {
		want := etofmdListEntry
		got, err := EtoFMDsList()
		if err != nil {
			t.Errorf("Unexpected error: %s", err)
		}

		if 4 != len(got) {
			t.Errorf("Result Length Mismatch: want: 4, got: %d", len(got))
		}

		if etofmdListEntry != got[0] {
			t.Errorf("Mismatch: want: \"%v\", got: \"%v\"", want, got)
		}
	})

}

func TestEtoFMDGetFunc(t *testing.T) {

	mux := setupMux("/api/v0/etofmds/3ca8ecaf-6fae-48a5-8441-5a96e119ad28", "testdata/etofmd-get.json")
	ts := httptest.NewServer(mux)
	defer ts.Close()

	setupTestServerClient(ts)

	t.Run("result", func(t *testing.T) {
		want := etofmdShow
		got := EtoFMDEntry{ID: "3ca8ecaf-6fae-48a5-8441-5a96e119ad28"}

		err := got.Get()
		if err != nil {
			t.Errorf("Unexpected error: %s", err)
		}

		if want.ID != got.ID {
			t.Errorf("Mismatch: want: \"%v\", got: \"%v\"", want, got)
		}
	})

}

// func TestFMDCreateFunc(t *testing.T) {
// 	setupLocalhostClient()

// 	err := fmdToCreate.Create()
// 	if err != nil {
// 		t.Errorf("Unexpected error: %s", err)
// 	}

// 	t.Run("confirm that attributes updated", func(t *testing.T) {
// 		if fmdToCreate.ID == "" {
// 			t.Errorf("ID not updated")
// 		}

// 		if fmdToCreate.CreatedAt == "" {
// 			t.Errorf("CreatedAt not updated")
// 		}

// 		if fmdToCreate.UpdatedAt == "" {
// 			t.Errorf("UpdatedAt not updated")
// 		}
// 	})
// }

// func TestFMDUpdateFunc(t *testing.T) {
// 	setupLocalhostClient()

// 	_ = fmdToCreate.Get()

// 	if fmdToCreate.Formats.PRONOM != "fmt/14" {
// 		t.Errorf("variable already updated: %s", fmdToCreate.ToString())
// 	}

// 	fmdToCreate.Formats.PRONOM = "fmt/99"

// 	err := fmdToCreate.Update()
// 	if err != nil {
// 		t.Errorf("Unexpected error: %s", err)
// 	}

// 	_ = fmdToCreate.Get()

// 	t.Run("confirm that elements updated", func(t *testing.T) {
// 		if fmdToCreate.Formats.PRONOM != "fmt/99" {
// 			t.Errorf("Formats was not updated: got: %s", fmdToCreate.Formats.PRONOM)
// 		}

// 		if fmdToCreate.CreatedAt == fmdToCreate.UpdatedAt {
// 			t.Errorf("UpdatedAt not updated")
// 		}
// 	})
// }

// func TestFMDDeleteFunc(t *testing.T) {
// 	setupLocalhostClient()

// 	_ = fmdToCreate.Get()

// 	id := fmdToCreate.ID

// 	err := fmdToCreate.Delete()
// 	if err != nil {
// 		t.Errorf("Unexpected error: %s", err)
// 	}

// 	t.Run("confirm that deleted item not found", func(t *testing.T) {
// 		// should not be found, so err should NOT be nil
// 		_, err = FMDGet(id)

// 		if err != nil {
// 			t.Errorf("err was nil")
// 		}

// 	})
// }
