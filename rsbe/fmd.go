package rsbe

import (
	"encoding/json"
	"fmt"
)

type FMDListEntry struct {
	ID            string  `json:"id,omitempty"`
	Name          string  `json:"name,omitempty"`
	Size          uint64  `json:"size,omitempty"`
	Status        string  `json:"status,omitempty"`
	MTime         string  `json:"file_mtime,omitempty"`
	Data          FMDData `json:"data,omitempty"`
	URL           string  `json:"url,omitempty"`
	PartnerURL    string  `json:"partner_url,omitempty"`
	CollectionURL string  `json:"coll_url,omitempty"`
}

type FMDEntry struct {
	ID               string    `json:"id,omitempty"`
	PartnerID        string    `json:"partner_id,omitempty"`      // REQUIRED
	CollectionID     string    `json:"coll_id,omitempty"`         // REQUIRED
	XIPID            string    `json:"xip_id,omitempty"`          /// ---
	Size             uint64    `json:"size,omitempty"`            // REQUIRED
	PresLevel        string    `json:"pres_level,omitempty"`      // ---
	PresCommitment   string    `json:"pres_commitment,omitempty"` // ---
	Status           string    `json:"status,omitempty"`          // ---
	FormatValid      bool      `json:"fmt_valid,omitempty"`       // ---
	FormatAcceptable bool      `json:"fmt_acceptable,omitempty"`  // ---
	OriginalName     string    `json:"original_name,omitempty"`   // REQUIRED
	Name             string    `json:"name,omitempty"`            // REQUIRED
	Extension        string    `json:"extension,omitempty"`       // (optional: file may not have an extension)
	MTime            string    `json:"file_mtime,omitempty"`      // REQUIRED
	HashMD5          string    `json:"hash_md5,omitempty"`        // ---
	HashSHA1         string    `json:"hash_sha1,omitempty"`       // ---
	HashSHA256       string    `json:"hash_sha256,omitempty"`     // ---
	HashSHA512       string    `json:"hash_sha512,omitempty"`     // ---
	CreatedAt        string    `json:"created_at,omitempty"`      // ---
	UpdatedAt        string    `json:"updated_at,omitempty"`      // ---
	Formats          FMDFormat `json:"fmts,omitempty"`            // ---
	Data             FMDData   `json:"data,omitempty"`            // ---
	PartnerURL       string    `json:"partner_url,omitempty"`     // ---
	CollectionURL    string    `json:"coll_url,omitempty"`        // ---
	LockVersion      int       `json:"lock_version,omitempty"`    // ---
}

type FMDFormat struct {
	PRONOM string `json:"pronom,omitempty"`
	MIME   string `json:"mime,omitempty"`
}

type FMDData struct {
	Searchable      bool   `json:"searchable"`
	Duration        string `json:"duration,omitempty"`
	Bitrate         uint64 `json:"bitrate,omitempty"`
	Width           uint   `json:"width,omitempty"`
	Height          uint   `json:"height,omitempty"`
	AspectRatio     string `json:"aspect_ratio,omitempty"`
	XMLSchema       string `json:"xml_schema,omitempty"`
	TranscriptionID string `json:"transcription_id,omitempty"`
}

func SEFMDList(eID string) (list []FMDListEntry, err error) {
	path := fmt.Sprintf("/api/v0/ses/%s/fmds", eID)
	return fmdList(path)
}

func IEFMDList(eID string) (list []FMDListEntry, err error) {
	path := fmt.Sprintf("/api/v0/ies/%s/fmds", eID)
	return fmdList(path)
}

func fmdList(path string) (list []FMDListEntry, err error) {
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

func (p *FMDEntry) Create() (err error) {
	path := "/api/v0/fmds"

	data, err := json.Marshal(p)
	if err != nil {
		return err
	}

	body, err := PostReturnBody(path, data)
	if err != nil {
		return err
	}

	fmt.Printf("%v\n", string(body))

	err = json.Unmarshal(body, p)
	if err != nil {
		return err
	}

	return nil
}

func (c *FMDEntry) Update() (err error) {
	path := "/api/v0/fmds/" + c.ID

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

func FMDDelete(id string) (err error) {
	path := "/api/v0/fmds/" + id

	err = Delete(path)
	if err != nil {
		return err
	}
	return nil
}

func (c *FMDEntry) Delete() (err error) {
	return FMDDelete(c.ID)
}

func (e FMDListEntry) ToString() string {
	s := fmt.Sprintf("ID: %s, Name: %s, Size: %d, Status: %s, MTime: %s, Data: %v, URL: %s , PartnerURL: %s, CollectionURL: %s",
		e.ID, e.Name, e.Size, e.Status, e.MTime, e.Data, e.URL, e.PartnerURL, e.CollectionURL)
	return s
}

func (e FMDEntry) ToString() string {
	s := fmt.Sprintf("ID: %s, PartnerID: %s, CollectionID: %s, XIPID: %s, Size: %d, PresLevel: %s, PresCommitment: %s, Status: %s, FormatValid: %v, FormatAcceptable: %v, OriginalName: %s, Name: %s, Extension: %s, MTime: %s, HashMD5: %s, HashSHA1: %s, HashSHA256: %s, HashSHA512: %s, CreatedAt: %s, UpdatedAt: %s, Formats: %v, Data: %v, PartnerURL: %s, CollectionURL: %s, LockVersion: %d",
		e.ID, e.PartnerID, e.CollectionID, e.XIPID, e.Size, e.PresLevel, e.PresCommitment, e.Status, e.FormatValid, e.FormatAcceptable, e.OriginalName, e.Name, e.Extension, e.MTime, e.HashMD5, e.HashSHA1, e.HashSHA256, e.HashSHA512, e.CreatedAt, e.UpdatedAt, e.Formats, e.Data, e.PartnerURL, e.CollectionURL, e.LockVersion)
	return s
}
