package rsbe

import (
	"encoding/json"
	"fmt"
)

type SEListEntry struct {
	ID            string `json:"id"`
	DigiID        string `json:"digi_id"`
	DOType        string `json:"dotype"`
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
	DOType        string `json:"dotype"`
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

func CollectionSEList(collectionID string) (list []SEListEntry, err error) {
	path := fmt.Sprintf("/api/v0/colls/%s/ses", collectionID)

	s, err := GetBodyTextString(path)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal([]byte(s), &list)
	if err != nil {
		return nil, err
	}

	return list, nil
}

func SEShow(id string) (item SEShowEntry, err error) {
	path := fmt.Sprintf("/api/v0/ses/%s", id)

	s, err := GetBodyTextString(path)
	if err != nil {
		return item, err
	}

	err = json.Unmarshal([]byte(s), &item)
	if err != nil {
		return item, err
	}

	return item, nil
}

func (li SEListEntry) ToString() string {
	s := fmt.Sprintf("ID: %s, DigiID: %s, DOType: %s, Phase: %s, Step: %s, Status: %s, Label: %s, Created_at: %s , Updated_at: %s, URL: %s, CollectionURL: %s",
		li.ID, li.DigiID, li.DOType, li.Phase, li.Step, li.Status, li.Label, li.Created_at, li.Updated_at, li.URL, li.CollectionURL)

	return s
}

func (i SEShowEntry) ToString() string {
	s := fmt.Sprintf("ID: %s, DigiID: %s, DOType: %s, Phase: %s, Step: %s, Status: %s, Label: %s, Title: %s, Created_at: %s , Updated_at: %s, BDIURL: %s, FMDsURL: %s, CollectionURL: %s, LockVersion: %d, Notes: %s", i.ID, i.DigiID, i.DOType, i.Phase, i.Step, i.Status, i.Label, i.Title, i.Created_at, i.Updated_at, i.BDIURL, i.FMDsURL, i.CollectionURL, i.LockVersion, i.Notes)

	return s
}
