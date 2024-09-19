package rsbe

import (
	//	"net/http/httptest"
	"testing"
)

var ieToSEToCreate = IEToSEEntry{
	ID:      "06de6d7a-89cd-476c-9e1d-55fdfabc3094",
	IEID:    "9ea98441-b6b6-46cf-b6c8-91dff385c6c8",
	SEID:    "8c258cb2-d700-43be-8773-a61a7b9cd668",
	Order:   1,
	Section: 23,
	Notes:   "IE to SE Notes",
}

func TestIEToSEList(t *testing.T) {

	setupLocalhostClient()
	err := ieToSEToCreate.Create()
	if err != nil {
		t.Errorf("Unexpected error: %s", err)
	}
	defer ieToSEToCreate.Delete()

	t.Run("result", func(t *testing.T) {
		want := ieToSEToCreate
		got, err := IEToSEList()
		if err != nil {
			t.Errorf("Unexpected error: %s", err)
		}

		if len(got) != 1 {
			t.Errorf("Mismatch: want: \"%v\", got: \"%v\"", want, got)
		}

		if got[0].ID != want.ID {
			t.Errorf("ID Mismatch: want:\n\"%v\", got:\n\"%v\"", want.ID, got[0].ID)
		}

		if got[0].IEID != want.IEID {
			t.Errorf("IEID Mismatch: want:\n\"%v\", got:\n\"%v\"", want.IEID, got[0].IEID)
		}

		if got[0].SEID != want.SEID {
			t.Errorf("SEID Mismatch: want:\n\"%v\", got:\n\"%v\"", want.SEID, got[0].SEID)
		}

		if got[0].Order != want.Order {
			t.Errorf("Order Mismatch: want:\n\"%v\", got:\n\"%v\"", want.Order, got[0].Order)
		}

		if got[0].Section != want.Section {
			t.Errorf("Section Mismatch: want: %v, got: %v", want.Section, got[0].Section)
		}

		if got[0].CreatedAt == "" {
			t.Errorf("CreatedAt is empty")
		}

		if got[0].UpdatedAt == "" {
			t.Errorf("UpdatedAt is empty")
		}

		url := "http://localhost:3000/api/v0/ie_to_ses/06de6d7a-89cd-476c-9e1d-55fdfabc3094"
		if got[0].URL != url {
			t.Errorf("URL Mismatch: want: %v, got: %v", url, got[0].URL)
		}
	})
}

func TestIEToSEListWithFilters(t *testing.T) {

	setupLocalhostClient()
	err := ieToSEToCreate.Create()
	if err != nil {
		t.Errorf("Unexpected error: %s", err)
	}
	defer ieToSEToCreate.Delete()

	ieToCreate.ID = "56c61005-ba14-47dc-a073-a03f66cf84e6"
	err = ieToCreate.Create()
	if err != nil {
		t.Errorf("Unexpected error: %s", err)
	}
	defer ieToCreate.Delete()

	ieToSEToCreate1 := IEToSEEntry{
		ID:      "eff4ef7e-961a-4687-8707-990584fa6660",
		IEID:    "56c61005-ba14-47dc-a073-a03f66cf84e6",
		SEID:    ieToSEToCreate.SEID, // reuse existing SE
		Order:   1000,
		Section: 1,
		Notes:   "blah blah blah",
	}
	err = ieToSEToCreate1.Create()
	if err != nil {
		t.Errorf("Unexpected error: %s", err)
	}
	defer ieToSEToCreate1.Delete()

	t.Run("check that proper results are returned", func(t *testing.T) {

		// run without filter
		got, err := IEToSEList()
		if err != nil {
			t.Errorf("Unexpected error: %s", err)
		}

		if len(got) != 2 {
			t.Errorf("Mismatch: want: \"%v\", got: \"%v\"", 2, len(got))
		}

		// run with filter
		filter := IEToSEListEntry{IEID: "56c61005-ba14-47dc-a073-a03f66cf84e6"}
		want := ieToSEToCreate1

		got, err = IEToSEList(filter)
		if err != nil {
			t.Errorf("Unexpected error: %s", err)
		}

		if len(got) != 1 {
			t.Errorf("Mismatch: want: \"%v\", got: \"%v\"", 1, len(got))
		}

		if got[0].ID != want.ID {
			t.Errorf("ID Mismatch: want:\n\"%v\", got:\n\"%v\"", want.ID, got[0].ID)
		}

		if got[0].IEID != want.IEID {
			t.Errorf("IEID Mismatch: want:\n\"%v\", got:\n\"%v\"", want.IEID, got[0].IEID)
		}

		if got[0].SEID != want.SEID {
			t.Errorf("SEID Mismatch: want:\n\"%v\", got:\n\"%v\"", want.SEID, got[0].SEID)
		}

		if got[0].Order != want.Order {
			t.Errorf("Order Mismatch: want:\n\"%v\", got:\n\"%v\"", want.Order, got[0].Order)
		}

		if got[0].Section != want.Section {
			t.Errorf("Section Mismatch: want: %v, got: %v", want.Section, got[0].Section)
		}

		if got[0].CreatedAt == "" {
			t.Errorf("CreatedAt is empty")
		}

		if got[0].UpdatedAt == "" {
			t.Errorf("UpdatedAt is empty")
		}

		url := "http://localhost:3000/api/v0/ie_to_ses/eff4ef7e-961a-4687-8707-990584fa6660"
		if got[0].URL != url {
			t.Errorf("URL Mismatch: want: %v, got: %v", url, got[0].URL)
		}
	})
}

// test that struct method Get works
func TestIEToSEGetFunc(t *testing.T) {

	setupLocalhostClient()
	err := ieToSEToCreate.Create()
	if err != nil {
		t.Errorf("Unexpected error: %s", err)
	}
	defer ieToSEToCreate.Delete()

	t.Run("result", func(t *testing.T) {
		want := ieToSEToCreate
		got := IEToSEEntry{ID: ieToSEToCreate.ID}

		err := got.Get()
		if err != nil {
			t.Errorf("Unexpected error: %s", err)
		}

		if got.ID != want.ID {
			t.Errorf("ID Mismatch: want:\n\"%v\", got:\n\"%v\"", want.ID, got.ID)
		}

		if got.IEID != want.IEID {
			t.Errorf("IEID Mismatch: want:\n\"%v\", got:\n\"%v\"", want.IEID, got.IEID)
		}

		if got.SEID != want.SEID {
			t.Errorf("SEID Mismatch: want:\n\"%v\", got:\n\"%v\"", want.SEID, got.SEID)
		}

		if got.Order != want.Order {
			t.Errorf("Order Mismatch: want:\n\"%v\", got:\n\"%v\"", want.Order, got.Order)
		}

		if got.Section != want.Section {
			t.Errorf("Section Mismatch: want: %v, got: %v", want.Section, got.Section)
		}

		if got.Notes != want.Notes {
			t.Errorf("Notes Mismatch: want: %v, got: %v", want.Notes, got.Notes)
		}

		url := "http://localhost:3000/api/v0/ies/9ea98441-b6b6-46cf-b6c8-91dff385c6c8"
		if got.IEURL != url {
			t.Errorf("IEURL Mismatch: want: %v, got: %v", url, got.IEURL)
		}

		url = "http://localhost:3000/api/v0/ses/8c258cb2-d700-43be-8773-a61a7b9cd668"
		if got.SEURL != url {
			t.Errorf("SEURL Mismatch: want: %v, got: %v", url, got.SEURL)
		}

		url = "http://localhost:3000/api/v0/ie_to_ses"
		if got.IEToSEsURL != url {
			t.Errorf("SEURL Mismatch: want: %v, got: %v", url, got.IEToSEsURL)
		}

		if got.CreatedAt == "" {
			t.Errorf("CreatedAt is empty")
		}

		if got.UpdatedAt == "" {
			t.Errorf("UpdatedAt is empty")
		}
	})
}

// test that model-level Get works
func TestIEToSEGet(t *testing.T) {

	setupLocalhostClient()
	err := ieToSEToCreate.Create()
	if err != nil {
		t.Errorf("Unexpected error: %s", err)
	}
	defer ieToSEToCreate.Delete()

	t.Run("confirm that expected resource was retrieved", func(t *testing.T) {
		want := ieToSEToCreate
		got, err := IEToSEGet("06de6d7a-89cd-476c-9e1d-55fdfabc3094")
		if err != nil {
			t.Errorf("Unexpected error: %s", err)
		}

		if got.ID != want.ID {
			t.Errorf("ID Mismatch: want:\n\"%v\", got:\n\"%v\"", want.ID, got.ID)
		}

		if got.IEID != want.IEID {
			t.Errorf("IEID Mismatch: want:\n\"%v\", got:\n\"%v\"", want.IEID, got.IEID)
		}

		if got.SEID != want.SEID {
			t.Errorf("SEID Mismatch: want:\n\"%v\", got:\n\"%v\"", want.SEID, got.SEID)
		}

		if got.Order != want.Order {
			t.Errorf("Order Mismatch: want:\n\"%v\", got:\n\"%v\"", want.Order, got.Order)
		}

		if got.Section != want.Section {
			t.Errorf("Section Mismatch: want: %v, got: %v", want.Section, got.Section)
		}

		if got.Notes != want.Notes {
			t.Errorf("Notes Mismatch: want: %v, got: %v", want.Notes, got.Notes)
		}

		url := "http://localhost:3000/api/v0/ies/9ea98441-b6b6-46cf-b6c8-91dff385c6c8"
		if got.IEURL != url {
			t.Errorf("IEURL Mismatch: want: %v, got: %v", url, got.IEURL)
		}

		url = "http://localhost:3000/api/v0/ses/8c258cb2-d700-43be-8773-a61a7b9cd668"
		if got.SEURL != url {
			t.Errorf("SEURL Mismatch: want: %v, got: %v", url, got.SEURL)
		}

		url = "http://localhost:3000/api/v0/ie_to_ses"
		if got.IEToSEsURL != url {
			t.Errorf("SEURL Mismatch: want: %v, got: %v", url, got.IEToSEsURL)
		}

		if got.CreatedAt == "" {
			t.Errorf("CreatedAt is empty")
		}

		if got.UpdatedAt == "" {
			t.Errorf("UpdatedAt is empty")
		}
	})
}

func TestIEToSECreateFunc(t *testing.T) {
	setupLocalhostClient()

	err := ieToSEToCreate.Create()
	if err != nil {
		t.Errorf("Unexpected error: %s", err)
	}

	t.Run("confirm that attributes updated", func(t *testing.T) {
		if ieToSEToCreate.ID == "" {
			t.Errorf("ID not updated")
		}

		if ieToSEToCreate.CreatedAt == "" {
			t.Errorf("CreatedAt not updated")
		}

		if ieToSEToCreate.UpdatedAt == "" {
			t.Errorf("UpdatedAt not updated")
		}
	})
}

func TestIEToSEUpdateFunc(t *testing.T) {
	setupLocalhostClient()
	id := ieToSEToCreate.ID

	sut, err := IEToSEGet(id)
	if err != nil {
		t.Errorf("Unexpected error: %s", err)
	}

	if sut.Order != 1 {
		t.Errorf("variable already updated: %s", sut.ToString())
	}

	sut.Order = 97
	sut.Notes = "Hop on Pop!"

	err = sut.Update()
	if err != nil {
		t.Errorf("Unexpected error: %s", err)
	}

	sut = IEToSEEntry{}
	sut, err = IEToSEGet(id)
	if err != nil {
		t.Errorf("Unexpected error: %s", err)
	}

	t.Run("confirm that elements updated", func(t *testing.T) {
		if sut.Order != 97 {
			t.Errorf("Order was not updated: got: %v", sut.Order)
		}

		if sut.Notes != "Hop on Pop!" {
			t.Errorf("Notes field was not updated: got: %s", sut.Notes)
		}

		if sut.CreatedAt == sut.UpdatedAt {
			t.Errorf("UpdatedAt not updated")
		}
	})
}

func TestIEToSEDeleteFunc(t *testing.T) {
	setupLocalhostClient()

	_ = ieToSEToCreate.Get()

	id := ieToSEToCreate.ID

	err := ieToSEToCreate.Delete()
	if err != nil {
		t.Errorf("Unexpected error: %s", err)
	}

	t.Run("confirm that deleted item not found", func(t *testing.T) {
		// should not be found, so err should NOT be nil
		_, err = IEToSEGet(id)

		if err == nil {
			t.Errorf("err was nil")
		}

	})
}
