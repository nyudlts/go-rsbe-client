package rsbe

import (
	"encoding/json"
	"fmt"
	"path/filepath"
)

type SEListEntry struct {
	ID            string `json:"id,omitempty"`
	DigiID        string `json:"digi_id,omitempty"`
	DOType        string `json:"do_type,omitempty"`
	Phase         string `json:"phase,omitempty"`
	Step          string `json:"step,omitempty"`
	Status        string `json:"status,omitempty"`
	Label         string `json:"label,omitempty"`
	CreatedAt     string `json:"created_at,omitempty"`
	UpdatedAt     string `json:"updated_at,omitempty"`
	URL           string `json:"url,omitempty"`
	CollectionURL string `json:"coll_url,omitempty"`
}

type SEEntry struct {
	ID            string `json:"id,omitempty"`
	CollectionID  string `json:"coll_id,omitempty"` // REQUIRED
	DigiID        string `json:"digi_id,omitempty"` // REQUIRED
	DOType        string `json:"do_type,omitempty"` // REQUIRED
	Phase         string `json:"phase,omitempty"`   // REQUIRED
	Step          string `json:"step,omitempty"`    // REQUIRED
	Status        string `json:"status,omitempty"`  // REQUIRED
	Notes         string `json:"notes,omitempty"`   
	Label         string `json:"label,omitempty"`
	Title         string `json:"title,omitempty"`
	CreatedAt     string `json:"created_at,omitempty"`
	UpdatedAt     string `json:"updated_at,omitempty"`
	BDIURL        string `json:"bdi_url,omitempty"`
	FMDsURL       string `json:"fmds_url,omitempty"`
	CollectionURL string `json:"coll_url,omitempty"`
	LockVersion   int    `json:"lock_version,omitempty"`
}

func CollectionSEList(collectionID string) (list []SEListEntry, err error) {
	path := fmt.Sprintf("/api/v0/colls/%s/ses", collectionID)

	body, err := GetBody(path)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(body, &list)
	if err != nil {
		return nil, err
	}

	return list, nil
}

func SEGet(id string) (item SEEntry, err error) {
	path := fmt.Sprintf("/api/v0/ses/%s", id)

	body, err := GetBody(path)
	if err != nil {
		return item, err
	}

	err = json.Unmarshal(body, &item)
	if err != nil {
		return item, err
	}

	return item, nil
}

func SEGetByDigiID(digiID string) (item SEEntry, err error) {
	item.DigiID = digiID

	err = item.GetByDigiID()
	return item, err
}

func SEDelete(id string) (err error) {
	path := "/api/v0/ses/" + id

	err = Delete(path)
	if err != nil {
		return err
	}
	return nil
}

func (c *SEEntry) Get() (err error) {
	path := fmt.Sprintf("/api/v0/ses/%s", c.ID)

	body, err := GetBody(path)
	if err != nil {
		return err
	}

	err = json.Unmarshal(body, c)
	if err != nil {
		return err
	}

	return nil
}

func (c *SEEntry) GetByDigiID() (err error) {
	path := fmt.Sprintf("/api/v0/search?scope=ses&digi_id=%s", c.DigiID)

	var searchResult SearchResult

	body, err := GetBody(path)
	if err != nil {
		return err
	}

	err = json.Unmarshal(body, &searchResult)
	if err != nil {
		return err
	}

	if searchResult.Response.NumFound != 1 {
		return fmt.Errorf("Incorrect number of results. Expected 1, found %d", searchResult.Response.NumFound)
	}

	c.ID = filepath.Base(searchResult.Response.Docs[0].URL)
	return c.Get()
}


func (c *SEEntry) Create() (err error) {
	path := "/api/v0/ses"

	data, err := json.Marshal(c)
	if err != nil {
		return err
	}

	body, err := PostReturnBody(path, data)
	if err != nil {
		return err
	}

	err = json.Unmarshal(body, c)
	if err != nil {
		return err
	}

	return nil
}

func (c *SEEntry) Update() (err error) {
	path := "/api/v0/ses/" + c.ID

	data, err := json.Marshal(c)
	if err != nil {
		return err
	}

	err = Put(path, data)
	if err != nil {
		return err
	}

	return nil
}

func (c *SEEntry) Delete() (err error) {
	return SEDelete(c.ID)
}

func (e SEListEntry) ToString() string {
	s := fmt.Sprintf("ID: %s, DigiID: %s, DOType: %s, Phase: %s, Step: %s, Status: %s, Label: %s, CreatedAt: %s , UpdatedAt: %s, URL: %s, CollectionURL: %s",
		e.ID, e.DigiID, e.DOType, e.Phase, e.Step, e.Status, e.Label, e.CreatedAt, e.UpdatedAt, e.URL, e.CollectionURL)

	return s
}

func (e SEEntry) ToString() string {
	s := fmt.Sprintf("ID: %s, DigiID: %s, DOType: %s, Phase: %s, Step: %s, Status: %s, Label: %s, Title: %s, CreatedAt: %s , UpdatedAt: %s, BDIURL: %s, FMDsURL: %s, CollectionURL: %s, LockVersion: %d, Notes: %s", e.ID, e.DigiID, e.DOType, e.Phase, e.Step, e.Status, e.Label, e.Title, e.CreatedAt, e.UpdatedAt, e.BDIURL, e.FMDsURL, e.CollectionURL, e.LockVersion, e.Notes)

	return s
}
