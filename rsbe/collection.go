package rsbe

import (
	"encoding/json"
	"fmt"
)

type CollectionListEntry struct {
	ID             string `json:"id"`
	PartnerID      string `json:"partner_id"`
	OwnerID        string `json:"owner_id"`
	Code           string `json:"code"`
	DisplayCode    string `json:"display_code"`
	Name           string `json:"name"`
	Type           string `json:"coll_type"`
	Classification string `json:"classification"`
	CreatedAt      string `json:"created_at"`
	UpdatedAt      string `json:"updated_at"`
	URL            string `json:"url"`
	PartnerURL     string `json:"partner_url"`
	OwnerURL       string `json:"owner_url"`
}

type CollectionEntry struct {
	ID              string `json:"id,omitempty"`
	PartnerID       string `json:"partner_id"`         // REQUIRED for CREATE
	OwnerID         string `json:"owner_id"`           // REQUIRED    ""
	Code            string `json:"code"`               // REQUIRED    ""
	DisplayCode     string `json:"display_code"`       // REQUIRED    ""
	Name            string `json:"name,omitempty"`     // optional    ""
	Type            string `json:"coll_type"`          // REQUIRED    ""  (origin, virtual)
	Classification  string `json:"classification"`     // REQUIRED    ""
	Quota           int    `json:"quota"`              // REQUIRED    ""
	RelPath         string `json:"rel_path,omitempty"` // REQUIRED    ""
	ReadyForContent bool   `json:"ready_for_content"`  // REQUIRED    ""
	PartnerURL      string `json:"partner_url,omitempty"`
	OwnerURL        string `json:"owner_url,omitempty"`
	SEsURL          string `json:"ses_url,omitempty"`
	IEsURL          string `json:"ies_url,omitempty"`
	LockVersion     int    `json:"lock_version,omitempty"`
	CreatedAt       string `json:"created_at,omitempty"`
	UpdatedAt       string `json:"updated_at,omitempty"`
}

func PartnerCollectionList(partnerID string) (collections []CollectionListEntry, err error) {
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

func OwnerCollectionList(ownerID string) (collections []CollectionListEntry, err error) {
	path := fmt.Sprintf("/api/v0/owners/%s/colls", ownerID)

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

func CollectionDelete(id string) (err error) {
	path := "/api/v0/colls/" + id

	err = Delete(path)
	if err != nil {
		return err
	}
	return nil
}

func (c *CollectionEntry) Get() (err error) {
	path := fmt.Sprintf("/api/v0/colls/%s", c.ID)

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

func (c *CollectionEntry) Create() (err error) {
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

func (c *CollectionEntry) Update() (err error) {
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

func (c *CollectionEntry) Delete() (err error) {
	return CollectionDelete(c.ID)
}

func (e CollectionListEntry) ToString() string {
	s := fmt.Sprintf("ID: %s, PartnerID: %s, OwnerID: %s, "+
		"Code: %s, DisplayCode: %s, Name: %s, Type: %s, Classification: %s, "+
		"CreatedAt: %s, UpdatedAt: %s, URL: %s, PartnerURL: %s, OwnerURL: %s",
		e.ID, e.PartnerID, e.OwnerID,
		e.Code, e.DisplayCode, e.Name, e.Type, e.Classification,
		e.CreatedAt, e.UpdatedAt, e.URL, e.PartnerURL, e.OwnerURL)

	return s
}

func (e CollectionEntry) ToString() string {
	s := fmt.Sprintf("ID: %s, PartnerID: %s, OwnerID: %s, Code: %s, "+
		"DisplayCode: %s, Name: %s, Type: %s, Classification: %s, "+
		"CreatedAt: %s, UpdatedAt: %s, Quota: %d, ReadyForContent: %v, "+
		"PartnerURL: %s, OwnerURL: %s, SEsURL: %s, IEsURL: %s, "+
		"LockVersion: %d, RelPath: %s", e.ID, e.PartnerID, e.OwnerID, e.Code,
		e.DisplayCode, e.Name, e.Type, e.Classification,
		e.CreatedAt, e.UpdatedAt, e.Quota, e.ReadyForContent,
		e.PartnerURL, e.OwnerURL, e.SEsURL, e.IEsURL, e.LockVersion, e.RelPath)
	return s
}
