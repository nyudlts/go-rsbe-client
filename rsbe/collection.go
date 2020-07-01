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

type CollectionEntry struct {
	ID              string `json:"id,omitempty"`
	PartnerID       string `json:"partner_id,omitempty"`        // REQUIRED
	Code            string `json:"code,omitempty"`              // REQUIRED
	Name            string `json:"name,omitempty"`              // optional
	CollType        string `json:"coll_type,omitempty"`         // REQUIRED (origin, virtual)
	Quota           int    `json:"quota,omitempty"`             // REQUIRED
	RelPath         string `json:"rel_path,omitempty"`          // REQUIRED
	ReadyForContent bool   `json:"ready_for_content,omitempty"` // optional
	PartnerURL      string `json:"partner_url,omitempty"`
	SEsURL          string `json:"ses_url,omitempty"`
	IEsURL          string `json:"ies_url,omitempty"`
	LockVersion     int    `json:"lock_version,omitempty"`
	Created_at      string `json:"created_at,omitempty"`
	Updated_at      string `json:"updated_at,omitempty"`
}

func PartnerCollectionsList(partnerID string) (collections []CollectionListEntry, err error) {
	path := fmt.Sprintf("/api/v0/partners/%s/colls", partnerID)

	body, err := GetBody(path)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(body, &collections)
	if err != nil {
		return nil, err
	}

	return collections, nil
}

func CollectionGet(id string) (collection CollectionEntry, err error) {
	path := fmt.Sprintf("/api/v0/colls/%s", id)

	body, err := GetBody(path)
	if err != nil {
		return collection, err
	}

	err = json.Unmarshal(body, &collection)
	if err != nil {
		return collection, err
	}

	return collection, nil
}

func CollectionCreate(c *CollectionEntry) (err error) {
	path := "/api/v0/colls"

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

func CollectionUpdate(c *PartnerEntry) (err error) {
	path := "/api/v0/colls/" + c.ID

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

func CollectionDelete(id string) (err error) {
	path := "/api/v0/colls/" + id

	err = Delete(path)
	if err != nil {
		return err
	}
	return nil
}

func (e CollectionListEntry) ToString() string {
	s := fmt.Sprintf("ID: %s, PartnerID: %s, Code: %s, Name: %s, CollType: %s, Created_at: %s , Updated_at: %s, URL: %s, PartnerURL: %s",
		e.ID, e.PartnerID, e.Code, e.Name, e.CollType, e.Created_at, e.Updated_at, e.URL, e.PartnerURL)

	return s
}

func (e CollectionEntry) ToString() string {
	s := fmt.Sprintf("ID: %s, PartnerID: %s, Code: %s, Name: %s, CollType: %s, Created_at: %s , Updated_at: %s, Quota: %d, ReadyForContent: %v, PartnerURL: %s, SEsURL: %s, IEsURL: %s, LockVersion: %d, RelPath: %s", e.ID, e.PartnerID, e.Code, e.Name, e.CollType, e.Created_at, e.Updated_at, e.Quota, e.ReadyForContent, e.PartnerURL, e.SEsURL, e.IEsURL, e.LockVersion, e.RelPath)

	return s
}
