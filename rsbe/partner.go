package rsbe

import (
	"encoding/json"
)

type PartnerIndexEntry struct {
	Id         string `json:"id"`
	Code       string `json:"code"`
	Name       string `json:"name"`
	Created_at string `json:"created_at"`
	Updated_at string `json:"updated_at"`
	Url        string `json:"url"`
}

type PartnerShowEntry struct {
	Id             string `json:"id"`
	Code           string `json:"code"`
	Name           string `json:"name"`
	Created_at     string `json:"created_at"`
	Updated_at     string `json:"updated_at"`
	PartnersURL    string `json:"partners_url"`
	CollectionsURL string `json:"colls_url"`
	LockVersion    int    `json:"lock_version"`
	RelPath        string `json:"rel_path"`
}

func PartnerIndex() (partners []PartnerIndexEntry, err error) {

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

func PartnerShow(id string) (partner PartnerShowEntry, err error) {
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
