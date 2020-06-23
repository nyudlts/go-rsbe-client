package rsbe

import (
	"encoding/json"
	"fmt"
)

type CollectionListEntry struct {
	ID         string `json:"id"`
	PartnerID  string `json:"partner_id"`
	Code       string `json:"code"`
	Name       string `json:"name"`
	CollType   string `json:"coll_type"`
	Created_at string `json:"created_at"`
	Updated_at string `json:"updated_at"`
	URL        string `json:"url"`
	PartnerURL string `json:"partner_url"`
}

type CollectionShowEntry struct {
	ID              string `json:"id"`
	PartnerID       string `json:"partner_id"`
	Code            string `json:"code"`
	Name            string `json:"name"`
	CollType        string `json:"coll_type"`
	Created_at      string `json:"created_at"`
	Updated_at      string `json:"updated_at"`
	Quota           int    `json:"quota"`
	ReadyForContent bool   `json:"ready_for_content"`
	PartnerURL      string `json:"partner_url"`
	SEsURL          string `json:"ses_url"`
	IEsURL          string `json:"ies_url"`
	LockVersion     int    `json:"lock_version"`
	RelPath         string `json:"rel_path"`
}

func PartnerCollectionsList(partnerID string) (collections []CollectionListEntry, err error) {
	path := fmt.Sprintf("/api/v0/partners/%s/colls", partnerID)

	s, err := GetBodyTextString(path)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal([]byte(s), &collections)
	if err != nil {
		return nil, err
	}

	return collections, nil
}

func CollectionGet(id string) (collection CollectionShowEntry, err error) {
	path := fmt.Sprintf("/api/v0/colls/%s", id)

	s, err := GetBodyTextString(path)
	if err != nil {
		return collection, err
	}

	err = json.Unmarshal([]byte(s), &collection)
	if err != nil {
		return collection, err
	}

	return collection, nil
}

func (e CollectionListEntry) ToString() string {
	s := fmt.Sprintf("ID: %s, PartnerID: %s, Code: %s, Name: %s, CollType: %s, Created_at: %s , Updated_at: %s, URL: %s, PartnerURL: %s",
		e.ID, e.PartnerID, e.Code, e.Name, e.CollType, e.Created_at, e.Updated_at, e.URL, e.PartnerURL)

	return s
}

func (e CollectionShowEntry) ToString() string {
	s := fmt.Sprintf("ID: %s, PartnerID: %s, Code: %s, Name: %s, CollType: %s, Created_at: %s , Updated_at: %s, Quota: %d, ReadyForContent: %v, PartnerURL: %s, SEsURL: %s, IEsURL: %s, LockVersion: %d, RelPath: %s", e.ID, e.PartnerID, e.Code, e.Name, e.CollType, e.Created_at, e.Updated_at, e.Quota, e.ReadyForContent, e.PartnerURL, e.SEsURL, e.IEsURL, e.LockVersion, e.RelPath)

	return s
}
