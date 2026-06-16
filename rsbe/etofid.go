package rsbe

import (
	"github.com/google/go-querystring/query"
	"github.com/google/uuid"

	"encoding/json"
	"fmt"
)

type EToFIDListEntry struct {
	ID       string `json:"id,omitempty"        url:"-"`
	EType    string `json:"etype,omitempty"      url:"etype,omitempty"`
	EID      string `json:"eid,omitempty"        url:"eid,omitempty"`
	FIDType  string `json:"fid_type,omitempty"   url:"fid_type,omitempty"`
	FIDValue string `json:"fid_value,omitempty" url:"fid_value,omitempty"`
	URL      string `json:"url,omitempty"        url:"-"`
	EURL     string `json:"eurl,omitempty"       url:"-"`
}

type EToFIDEntry struct {
	ID          string `json:"id,omitempty"`
	EID         string `json:"eid,omitempty"`
	EType       string `json:"etype,omitempty"`
	FIDType     string `json:"fid_type,omitempty"`
	FIDValue    string `json:"fid_value,omitempty"`
	CreatedAt   string `json:"created_at,omitempty"`
	UpdatedAt   string `json:"updated_at,omitempty"`
	DeletedAt   string `json:"deleted_at,omitempty"`
	EURL        string `json:"eurl,omitempty"`
	LockVersion int    `json:"lock_version"`
}

// Get a list of EToFIDListEntry objects
// Function accepts 0 or 1 EToFIDListEntry parameters.
//
// If an EToFIDListEntry parameter is passed, the EID, EType,
// FIDType, and FIDValue fields are added to the RSBE query as query
// params.
func EToFIDList(e ...EToFIDListEntry) (list []EToFIDListEntry, err error) {

	path := "/api/v0/etofids"

	// check if there are any query parameters
	switch len(e) {
	case 0:
		// noop
	case 1:
		// extract url.Values
		v, err := query.Values(e[0])
		if err != nil {
			return nil, err
		}

		// if encoded values are not empty, append to path
		if v.Encode() != "" {
			path += "?" + v.Encode()
		}
	default:
		return list, fmt.Errorf("error: can only accept 0 or 1 EToFIDListEntry arguments")
	}

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

func EToFIDGet(id string) (item EToFIDEntry, err error) {
	_, err = uuid.Parse(id)
	if err != nil {
		return EToFIDEntry{}, fmt.Errorf("id is not a UUID: %s", err.Error())
	}

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
	_, err = uuid.Parse(p.ID)
	if err != nil {
		return fmt.Errorf("ID is not a UUID: %s", err.Error())
	}

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

func (p *EToFIDEntry) Create() (err error) {
	path := "/api/v0/etofids"

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

func (c *EToFIDEntry) Update() (err error) {
	_, err = uuid.Parse(c.ID)
	if err != nil {
		return fmt.Errorf("ID is not a UUID: %s", err.Error())
	}

	path := "/api/v0/etofids/" + c.ID

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

func EToFIDDelete(id string) (err error) {
	_, err = uuid.Parse(id)
	if err != nil {
		return fmt.Errorf("id is not a UUID: %s", err.Error())
	}

	path := "/api/v0/etofids/" + id

	err = Delete(path)
	if err != nil {
		return err
	}
	return nil
}

func (c *EToFIDEntry) Delete() (err error) {
	return EToFIDDelete(c.ID)
}

func EToFIDPurge(id string) (err error) {
	_, err = uuid.Parse(id)
	if err != nil {
		return fmt.Errorf("id is not a UUID: %s", err.Error())
	}

	path := "/api/v0/etofids/" + id

	err = Purge(path)
	if err != nil {
		return err
	}
	return nil
}

func (c *EToFIDEntry) Purge() (err error) {
	return EToFIDPurge(c.ID)
}

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
