package rsbe

import (
	"encoding/json"
	"fmt"
)

type PartnerListEntry struct {
	ID         string `json:"id"`
	Code       string `json:"code"`
	Name       string `json:"name"`
	Created_at string `json:"created_at"`
	Updated_at string `json:"updated_at"`
	URL        string `json:"url"`
}

type PartnerShowEntry struct {
	ID             string `json:"id"`
	Code           string `json:"code"`
	Name           string `json:"name"`
	Created_at     string `json:"created_at"`
	Updated_at     string `json:"updated_at"`
	PartnersURL    string `json:"partners_url"`
	CollectionsURL string `json:"colls_url"`
	LockVersion    int    `json:"lock_version"`
	RelPath        string `json:"rel_path"`
}

func PartnerList() (partners []PartnerListEntry, err error) {

	s, err := GetBodyTextString("/api/v0/partners")
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal([]byte(s), &partners)
	if err != nil {
		return nil, err
	}

	return partners, nil
}

func PartnerGet(id string) (partner PartnerShowEntry, err error) {
	path := "/api/v0/partners/" + id

	s, err := GetBodyTextString(path)
	if err != nil {
		return partner, err
	}

	err = json.Unmarshal([]byte(s), &partner)
	if err != nil {
		return partner, err
	}

	return partner, nil
}

func (e PartnerListEntry) ToString() string {
	s := fmt.Sprintf("ID: %s, Code: %s, Name: %s, Created_at: %s , Updated_at: %s, URL: %s",
		e.ID, e.Code, e.Name, e.Created_at, e.Updated_at, e.URL)

	return s
}

func (e PartnerShowEntry) ToString() string {
	s := fmt.Sprintf("ID: %s, Code: %s, Name: %s, Created_at: %s , Updated_at: %s, PartnersURL: %s, CollectionsURL: %s, LockVersion: %d, RelPath: %s",
		e.ID, e.Code, e.Name, e.Created_at, e.Updated_at, e.PartnersURL, e.CollectionsURL, e.LockVersion, e.RelPath)

	return s
}
