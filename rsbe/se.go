package rsbe

import (
	"encoding/json"
	"fmt"
)

type SEListEntry struct {
	ID            string `json:"id"`
	DigiID        string `json:"digi_id"`
	DOType        string `json:"do_type"`
	Phase         string `json:"phase"`
	Step          string `json:"step"`
	Status        string `json:"status"`
	Label         string `json:"label"`
	Created_at    string `json:"created_at"`
	Updated_at    string `json:"updated_at"`
	URL           string `json:"url"`
	CollectionURL string `json:"coll_url"`
}

type SEShowEntry struct {
	ID            string `json:"id"`
	CollectionID  string `json:"coll_id"`
	DigiID        string `json:"digi_id"`
	DOType        string `json:"do_type"`
	Phase         string `json:"phase"`
	Step          string `json:"step"`
	Status        string `json:"status"`
	Notes         string `json:"notes"`
	Label         string `json:"label"`
	Title         string `json:"title"`
	Created_at    string `json:"created_at"`
	Updated_at    string `json:"updated_at"`
	BDIURL        string `json:"bdi_url"`
	FMDsURL       string `json:"fmds_url"`
	CollectionURL string `json:"coll_url"`
	LockVersion   int    `json:"lock_version"`
}

func CollectionSEsList(collectionID string) (list []SEListEntry, err error) {
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

func SEGet(id string) (item SEShowEntry, err error) {
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

func (e SEListEntry) ToString() string {
	s := fmt.Sprintf("ID: %s, DigiID: %s, DOType: %s, Phase: %s, Step: %s, Status: %s, Label: %s, Created_at: %s , Updated_at: %s, URL: %s, CollectionURL: %s",
		e.ID, e.DigiID, e.DOType, e.Phase, e.Step, e.Status, e.Label, e.Created_at, e.Updated_at, e.URL, e.CollectionURL)

	return s
}

func (e SEShowEntry) ToString() string {
	s := fmt.Sprintf("ID: %s, DigiID: %s, DOType: %s, Phase: %s, Step: %s, Status: %s, Label: %s, Title: %s, Created_at: %s , Updated_at: %s, BDIURL: %s, FMDsURL: %s, CollectionURL: %s, LockVersion: %d, Notes: %s", e.ID, e.DigiID, e.DOType, e.Phase, e.Step, e.Status, e.Label, e.Title, e.Created_at, e.Updated_at, e.BDIURL, e.FMDsURL, e.CollectionURL, e.LockVersion, e.Notes)

	return s
}
