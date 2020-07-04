package rsbe

import (
	"encoding/json"
	"fmt"
)

type PartnerListEntry struct {
	ID        string `json:"id"`
	Code      string `json:"code"`
	Name      string `json:"name"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
	URL       string `json:"url"`
}

type PartnerEntry struct {
	ID             string `json:"id,omitempty"`
	Code           string `json:"code,omitempty"`
	Name           string `json:"name,omitempty"`
	CreatedAt      string `json:"created_at,omitempty"`
	UpdatedAt      string `json:"updated_at,omitempty"`
	PartnersURL    string `json:"partners_url,omitempty"`
	CollectionsURL string `json:"colls_url,omitempty"`
	LockVersion    int    `json:"lock_version,omitempty"`
	RelPath        string `json:"rel_path,omitempty"`
}

func PartnerList() (partners []PartnerEntry, err error) {

	body, err := GetBody("/api/v0/partners")
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(body, &partners)
	if err != nil {
		return nil, err
	}

	return partners, nil
}

func PartnerGet(id string) (partner PartnerEntry, err error) {
	path := "/api/v0/partners/" + id

	body, err := GetBody(path)
	if err != nil {
		return partner, err
	}

	err = json.Unmarshal(body, &partner)
	if err != nil {
		return partner, err
	}

	return partner, nil
}

func (p *PartnerEntry) Get() (err error) {
	path := "/api/v0/partners/" + p.ID

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

func (p *PartnerEntry) Create() (err error) {
	path := "/api/v0/partners"

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

func (p *PartnerEntry) Update() (err error) {
	path := "/api/v0/partners/" + p.ID

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

func PartnerDelete(id string) (err error) {
	path := "/api/v0/partners/" + id

	err = Delete(path)
	if err != nil {
		return err
	}
	return nil
}

func (p *PartnerEntry) Delete() (err error) {
	return PartnerDelete(p.ID)
}

func (e PartnerListEntry) ToString() string {
	s := fmt.Sprintf("ID: %s, Code: %s, Name: %s, CreatedAt: %s , UpdatedAt: %s, URL: %s",
		e.ID, e.Code, e.Name, e.CreatedAt, e.UpdatedAt, e.URL)

	return s
}

func (e PartnerEntry) ToString() string {
	s := fmt.Sprintf("ID: %s, Code: %s, Name: %s, CreatedAt: %s , UpdatedAt: %s, PartnersURL: %s, CollectionsURL: %s, LockVersion: %d, RelPath: %s",
		e.ID, e.Code, e.Name, e.CreatedAt, e.UpdatedAt, e.PartnersURL, e.CollectionsURL, e.LockVersion, e.RelPath)

	return s
}
