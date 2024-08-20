package rsbe

import (
	"encoding/json"
	"fmt"
)

type OwnerListEntry struct {
	ID        string `json:"id"`
	Code      string `json:"code"`
	Name      string `json:"name"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
	URL       string `json:"url"`
}

type OwnerEntry struct {
	ID             string `json:"id,omitempty"`
	Code           string `json:"code,omitempty"`
	Name           string `json:"name,omitempty"`
	CreatedAt      string `json:"created_at,omitempty"`
	UpdatedAt      string `json:"updated_at,omitempty"`
	OwnersURL      string `json:"owners_url,omitempty"`
	CollectionsURL string `json:"colls_url,omitempty"`
	LockVersion    int    `json:"lock_version,omitempty"`
}

func OwnerList() (owners []OwnerListEntry, err error) {

	body, err := GetBody("/api/v0/owners")
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(body, &owners)
	if err != nil {
		return nil, err
	}

	return owners, nil
}

func OwnerGet(id string) (owner OwnerEntry, err error) {
	path := "/api/v0/owners/" + id

	body, err := GetBody(path)
	if err != nil {
		return owner, err
	}

	err = json.Unmarshal(body, &owner)
	if err != nil {
		return owner, err
	}

	return owner, nil
}

func (p *OwnerEntry) Get() (err error) {
	path := "/api/v0/owners/" + p.ID

	body, err := GetBody(path)
	if err != nil {
		return err
	}

	err = json.Unmarshal(body, &p)
	if err != nil {
		return err
	}

	return nil
}

func (p *OwnerEntry) Create() (err error) {
	path := "/api/v0/owners"

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

func (p *OwnerEntry) Update() (err error) {
	path := "/api/v0/owners/" + p.ID

	data, err := json.Marshal(p)
	if err != nil {
		return err
	}

	err = Put(path, data)
	if err != nil {
		return err
	}

	return nil
}

func OwnerDelete(id string) (err error) {
	path := "/api/v0/owners/" + id

	err = Delete(path)
	if err != nil {
		return err
	}
	return nil
}

func (p *OwnerEntry) Delete() (err error) {
	return OwnerDelete(p.ID)
}

func (e OwnerListEntry) ToString() string {
	s := fmt.Sprintf("ID: %s, Code: %s, Name: %s, CreatedAt: %s , UpdatedAt: %s, URL: %s",
		e.ID, e.Code, e.Name, e.CreatedAt, e.UpdatedAt, e.URL)

	return s
}

func (e OwnerEntry) ToString() string {
	s := fmt.Sprintf("ID: %s, Code: %s, Name: %s, CreatedAt: %s , UpdatedAt: %s, OwnersURL: %s, CollectionsURL: %s, LockVersion: %d",
		e.ID, e.Code, e.Name, e.CreatedAt, e.UpdatedAt, e.OwnersURL, e.CollectionsURL, e.LockVersion)

	return s
}
