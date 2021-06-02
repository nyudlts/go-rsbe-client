package rsbe

import (
	"encoding/json"
	"fmt"
)

type EToFIDListEntry struct {
	ID       string `json:"id,omitempty"`
	EID      string `json:"eid,omitempty"`
	EType    string `json:"etype,omitempty"`
	FIDType  string `json:"fid_type,omitempty"`
	FIDValue string `json:"fid_value,omitempty"`
	URL      string `json:"url,omitempty"`
}

type EToFIDEntry struct {
	ID          string `json:"id,omitempty"`
	EID         string `json:"eid,omitempty"`
	EType       string `json:"etype,omitempty"`
	FIDType     string `json:"fid_type,omitempty"`
	FIDValue    string `json:"fid_value,omitempty"`
	CreatedAt   string `json:"created_at,omitempty"`
	UpdatedAt   string `json:"updated_at,omitempty"`
	EURL        string `json:"eurl,omitempty"`         
	LockVersion int    `json:"lock_version,omitempty"` 
}

func EToFIDList() (list []EToFIDListEntry, err error) {

	body, err := GetBody("/api/v0/etofids")
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(body, &list)
	if err != nil {
		return nil, err
	}

	return list, nil
}

func EToFIDGet(id string) (item EToFIDEntry, err error) {
	path := "/api/v0/etofids/" + id

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

func (p *EToFIDEntry) Get() (err error) {
	path := "/api/v0/etofids/" + p.ID

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

// func (p *EToFIDEntry) Create() (err error) {
// 	path := "/api/v0/etofids"

// 	data, err := json.Marshal(p)
// 	if err != nil {
// 		return err
// 	}

// 	body, err := PostReturnBody(path, data)
// 	if err != nil {
// 		return err
// 	}

// 	err = json.Unmarshal(body, p)
// 	if err != nil {
// 		return err
// 	}

// 	return nil
// }

// func (c *EToFIDEntry) Update() (err error) {
// 	path := "/api/v0/etofids/" + c.ID

// 	data, err := json.Marshal(c)
// 	if err != nil {
// 		return err
// 	}

// 	err = Put(path, data)
// 	if err != nil {
// 		return err
// 	}

// 	return nil
// }

// func EToFIDDelete(id string) (err error) {
// 	path := "/api/v0/etofids/" + id

// 	err = Delete(path)
// 	if err != nil {
// 		return err
// 	}
// 	return nil
// }

// func (c *EToFIDEntry) Delete() (err error) {
// 	return EToFIDDelete(c.ID)
// }

func (e EToFIDListEntry) ToString() string {
	s := fmt.Sprintf("ID: %s, EID: %s, EType: %s, FIDType: %s, FIDValue: %s, URL: %s",
		e.ID, e.EID, e.EType, e.FIDType, e.FIDValue, e.URL)
	return s
}

func (e EToFIDEntry) ToString() string {
	s := fmt.Sprintf("ID: %s, EID: %s, EType: %s, FIDType: %s, FIDValue: %s, CreatedAt: %s, UpdatedAt: %s, EURL: %s, LockVersion: %d",
		e.ID, e.EID, e.EType, e.FIDType, e.FIDValue, e.CreatedAt, e.UpdatedAt, e.EURL, e.LockVersion)
	return s
}
