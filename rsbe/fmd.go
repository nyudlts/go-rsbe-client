package rsbe

import (
	"encoding/json"
	"fmt"
)

type FMDListEntry struct {
	ID            string                 `json:"id,omitempty"`
	Name          string                 `json:"name,omitempty"`
	Size          uint64                 `json:"size,omitempty"`
	Status        string                 `json:"status,omitempty"`
	FileMTime     string                 `json:"file_mtime,omitempty"`
	DataHash      map[string]interface{} `json:"data,omitempty"`
	URL           string                 `json:"url,omitempty"`
	PartnerURL    string                 `json:"partner_url,omitempty"`
	CollectionURL string                 `json:"coll_url,omitempty"`
}

type FMDEntry struct {
	ID               string                 `json:"id,omitempty"`
	PartnerID        string                 `json:"partner_id,omitempty"`      // REQUIRED
	CollectionID     string                 `json:"coll_id,omitempty"`         // REQUIRED
	XIPID            string                 `json:"xip_id,omitempty"`          /// ---
	Size             uint64                 `json:"size,omitempty"`            // REQUIRED
	PresLevel        string                 `json:"pres_level,omitempty"`      // ---
	PresCommitment   string                 `json:"pres_commitment,omitempty"` // ---
	Status           string                 `json:"status,omitempty"`          // ---
	FormatValid      bool                   `json:"fmt_valid,omitempty"`       // ---
	FormatAcceptable bool                   `json:"fmt_acceptable,omitempty"`  // ---
	OriginalName     string                 `json:"original_name,omitempty"`   // REQUIRED
	Name             string                 `json:"name,omitempty"`            // REQUIRED
	Extension        string                 `json:"extension,omitempty"`       // (optional: file may not have an extension)
	FileMTime        string                 `json:"file_mtime,omitempty"`      // REQUIRED
	HashMD5          string                 `json:"hash_md5,omitempty"`        // ---
	HashSHA1         string                 `json:"hash_sha1,omitempty"`       // ---
	HashSHA256       string                 `json:"hash_sha256,omitempty"`     // ---
	HashSHA512       string                 `json:"hash_sha512,omitempty"`     // ---
	CreatedAt        string                 `json:"created_at,omitempty"`      // ---
	UpdatedAt        string                 `json:"updated_at,omitempty"`      // ---
	FormatsHash      map[string]interface{} `json:"fmts,omitempty"`            // ---
	DataHash         map[string]interface{} `json:"data,omitempty"`            // ---
	PartnerURL       string                 `json:"partner_url,omitempty"`     // ---
	CollectionURL    string                 `json:"coll_url,omitempty"`        // ---
	LockVersion      int                    `json:"lock_version,omitempty"`    // ---
}

func SEFMDsList(eID string) (list []FMDListEntry, err error) {
	path := fmt.Sprintf("/api/v0/ses/%s/fmds", eID)
	return fmdsList(path)
}

func IEFMDsList(eID string) (list []FMDListEntry, err error) {
	path := fmt.Sprintf("/api/v0/ies/%s/fmds", eID)
	return fmdsList(path)
}

func fmdsList(path string) (list []FMDListEntry, err error) {
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

func FMDGet(id string) (item FMDEntry, err error) {
	path := "/api/v0/fmds/" + id

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

func (p *FMDEntry) Get() (err error) {
	path := "/api/v0/fmds/" + p.ID

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

func (e FMDListEntry) ToString() string {
	s := fmt.Sprintf("ID: %s, Name: %s, Size: %d, Status: %s, FileMTime: %s, DataHash: %v, URL: %s , PartnerURL: %s, CollectionURL: %s",
		e.ID, e.Name, e.Size, e.Status, e.FileMTime, e.DataHash, e.URL, e.PartnerURL, e.CollectionURL)
	return s
}

func (e FMDEntry) ToString() string {
	s := fmt.Sprintf("ID: %s, PartnerID: %s, CollectionID: %s, XIPID: %s, Size: %d, PresLevel: %s, PresCommitment: %s, Status: %s, FormatValid: %v, FormatAcceptable: %v, OriginalName: %s, Name: %s, Extension: %s, FileMTime: %s, HashMD5: %s, HashSHA1: %s, HashSHA256: %s, HashSHA512: %s, CreatedAt: %s, UpdatedAt: %s, FormatsHash: %v, DataHash: %v, PartnerURL: %s, CollectionURL: %s, LockVersion: %d",
		e.ID, e.PartnerID, e.CollectionID, e.XIPID, e.Size, e.PresLevel, e.PresCommitment, e.Status, e.FormatValid, e.FormatAcceptable, e.OriginalName, e.Name, e.Extension, e.FileMTime, e.HashMD5, e.HashSHA1, e.HashSHA256, e.HashSHA512, e.CreatedAt, e.UpdatedAt, e.FormatsHash, e.DataHash, e.PartnerURL, e.CollectionURL, e.LockVersion)
	return s
}
