package rsbe

import (
	"encoding/json"
	"fmt"
)

type EToFMDListEntry struct {
	ID    string `json:"id,omitempty"`
	EID   string `json:"eid,omitempty"`
	EType string `json:"etype,omitempty"`
	FMDID string `json:"fmd_id,omitempty"`
	Role  string `json:"role,omitempty"`
	URL   string `json:"url,omitempty"`
}

type EToFMDEntry struct {
	ID          string `json:"id,omitempty"`
	EID         string `json:"eid,omitempty"`
	EType       string `json:"etype,omitempty"`
	FMDID       string `json:"fmd_id,omitempty"`
	Role        string `json:"role,omitempty"`
	CreatedAt   string `json:"created_at,omitempty"`
	UpdatedAt   string `json:"updated_at,omitempty"`
	EURL        string `json:"eurl,omitempty"`         // ---
	LockVersion int    `json:"lock_version,omitempty"` // ---
}

func EToFMDList() (list []EToFMDListEntry, err error) {

	body, err := GetBody("/api/v0/etofmds")
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(body, &list)
	if err != nil {
		return nil, err
	}

	return list, nil
}

func EToFMDGet(id string) (item EToFMDEntry, err error) {
	path := "/api/v0/etofmds/" + id

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

func (p *EToFMDEntry) Get() (err error) {
	path := "/api/v0/etofmds/" + p.ID

	body, err := GetBody(path)
	if err != nil {
		return err
	}

	err = json.Unmarshal(body, p)
	if err != nil {
		return err
	}

	return nil
}

func (p *EToFMDEntry) Create() (err error) {
	path := "/api/v0/etofmds"

	data, err := json.Marshal(p)
	if err != nil {
		return err
	}

	body, err := PostReturnBody(path, data)
	if err != nil {
		return err
	}

	err = json.Unmarshal(body, p)
	if err != nil {
		return err
	}

	return nil
}

func (c *EToFMDEntry) Update() (err error) {
	path := "/api/v0/etofmds/" + c.ID

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

func EToFMDDelete(id string) (err error) {
	path := "/api/v0/etofmds/" + id

	err = Delete(path)
	if err != nil {
		return err
	}
	return nil
}

func (c *EToFMDEntry) Delete() (err error) {
	return EToFMDDelete(c.ID)
}

func (e EToFMDListEntry) ToString() string {
	s := fmt.Sprintf("ID: %s, EID: %s, EType: %s, FMDID: %s, Role: %s, URL: %s",
		e.ID, e.EID, e.EType, e.FMDID, e.Role, e.URL)
	return s
}

func (e EToFMDEntry) ToString() string {
	s := fmt.Sprintf("ID: %s, EID: %s, EType: %s, FMDID: %s, Role: %s, CreatedAt: %s, UpdatedAt: %s, EURL: %s, LockVersion: %d",
		e.ID, e.EID, e.EType, e.FMDID, e.Role, e.CreatedAt, e.UpdatedAt, e.EURL, e.LockVersion)
	return s
}
