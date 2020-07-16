package rsbe

import (
	"encoding/json"
//	"fmt"
)

// #  id           :uuid             not null, primary key
// #  eid          :uuid             not null
// #  etype        :string(255)      not null
// #  fmd_id       :uuid             not null
// #  role         :string(255)
// #  lock_version :integer
// #  created_at   :datetime
// #  updated_at   :datetime

//  json.extract! etofmd, :id, :etype, :eid, :role, :fmd_id
//  json.url api_v0_etofmd_url(etofmd, format: :json)

type EtoFMDListEntry struct {
	ID    string `json:"id,omitempty"`
	EID   string `json:"eid,omitempty"`
	EType string `json:"etype,omitempty"`
	FMDID string `json:"fmd_id,omitempty"`
	Role  string `json:"role,omitempty"`
	URL   string `json:"url,omitempty"`
}

// json.extract! @etofmd, :id, :etype, :eid, :role, :fmd_id, :created_at, :updated_at
// case @etofmd.etype
// when 'se'
//   json.eurl api_v0_se_url(@etofmd.eid, format: :json)
// when 'ie'
//   json.eurl api_v0_ie_url(@etofmd.eid, format: :json)
// end
// if @current_user_privs == :read_write
//   json.extract! @etofmd, :lock_version
// end

type EtoFMDEntry struct {
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

func EtoFMDsList() (list []EtoFMDListEntry, err error) {

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

// func EtoFMDGet(id string) (item FMDEntry, err error) {
// 	path := "/api/v0/fmds/" + id

// 	body, err := GetBody(path)
// 	if err != nil {
// 		return item, err
// 	}

// 	err = json.Unmarshal(body, &item)
// 	if err != nil {
// 		return item, err
// 	}

// 	return item, nil
// }

// func (p *FMDEntry) Get() (err error) {
// 	path := "/api/v0/fmds/" + p.ID

// 	body, err := GetBody(path)
// 	if err != nil {
// 		return err
// 	}

// 	err = json.Unmarshal(body, p)
// 	if err != nil {
// 		return err
// 	}

// 	return nil
// }

// func (p *FMDEntry) Create() (err error) {
// 	path := "/api/v0/fmds"

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

// func (c *FMDEntry) Update() (err error) {
// 	path := "/api/v0/fmds/" + c.ID

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

// func FMDDelete(id string) (err error) {
// 	path := "/api/v0/fmds/" + id

// 	err = Delete(path)
// 	if err != nil {
// 		return err
// 	}
// 	return nil
// }

// func (c *FMDEntry) Delete() (err error) {
// 	return FMDDelete(c.ID)
// }

// func (e FMDListEntry) ToString() string {
// 	s := fmt.Sprintf("ID: %s, Name: %s, Size: %d, Status: %s, FileMTime: %s, Data: %v, URL: %s , PartnerURL: %s, CollectionURL: %s",
// 		e.ID, e.Name, e.Size, e.Status, e.FileMTime, e.Data, e.URL, e.PartnerURL, e.CollectionURL)
// 	return s
// }

// func (e FMDEntry) ToString() string {
// 	s := fmt.Sprintf("ID: %s, PartnerID: %s, CollectionID: %s, XIPID: %s, Size: %d, PresLevel: %s, PresCommitment: %s, Status: %s, FormatValid: %v, FormatAcceptable: %v, OriginalName: %s, Name: %s, Extension: %s, FileMTime: %s, HashMD5: %s, HashSHA1: %s, HashSHA256: %s, HashSHA512: %s, CreatedAt: %s, UpdatedAt: %s, Formats: %v, Data: %v, PartnerURL: %s, CollectionURL: %s, LockVersion: %d",
// 		e.ID, e.PartnerID, e.CollectionID, e.XIPID, e.Size, e.PresLevel, e.PresCommitment, e.Status, e.FormatValid, e.FormatAcceptable, e.OriginalName, e.Name, e.Extension, e.FileMTime, e.HashMD5, e.HashSHA1, e.HashSHA256, e.HashSHA512, e.CreatedAt, e.UpdatedAt, e.Formats, e.Data, e.PartnerURL, e.CollectionURL, e.LockVersion)
// 	return s
// }
