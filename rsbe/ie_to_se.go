package rsbe

import (
	"encoding/json"
	"fmt"
	//	"path/filepath"
)

// # Table name: ie_to_ses
// #
// #  id           :uuid             not null, primary key
// #  ie_id        :uuid             not null
// #  se_id        :uuid             not null
// #  order        :integer          default(1), not null
// #  section      :integer          default(1), not null
// #  notes        :text
// #  lock_version :integer
// #  created_at   :datetime
// #  updated_at   :datetime

type IEToSEListEntry struct {
	ID        string `json:"id,omitempty"`
	IEID      string `json:"ie_id,omitempty"`   // REQUIRED
	SEID      string `json:"se_id,omitempty"`   // REQUIRED
	Order     int    `json:"order,omitempty"`   // REQUIRED if != 1
	Section   int    `json:"section,omitempty"` // REQUIRED if != 1
	CreatedAt string `json:"created_at,omitempty"`
	UpdatedAt string `json:"updated_at,omitempty"`
	URL       string `json:"url",omitempty"`
}

type IEToSEEntry struct {
	ID          string `json:"id,omitempty"`
	IEID        string `json:"ie_id,omitempty"`   // REQUIRED
	SEID        string `json:"se_id,omitempty"`   // REQUIRED
	Order       int    `json:"order,omitempty"`   // REQUIRED if != 1
	Section     int    `json:"section,omitempty"` // REQUIRED if != 1
	Notes       string `json:"notes,omitempty"`
	CreatedAt   string `json:"created_at,omitempty"`
	UpdatedAt   string `json:"updated_at,omitempty"`
	LockVersion int    `json:"lock_version,omitempty"`
	IEURL       string `json:"ie_url,omitempty"`
	SEURL       string `json:"se_url,omitempty"`
	IEToSESURL  string `json:"ie_to_ses_url,omitempty"`
}

func IEToSEList(IEID string) (list []IEListEntry, err error) {
	path := fmt.Sprintf("/api/v0/ies/%s/ses", IEID)

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

func IEGet(id string) (item IEEntry, err error) {
	path := fmt.Sprintf("/api/v0/ies/%s", id)

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

func IEDelete(id string) (err error) {
	path := "/api/v0/ies/" + id

	err = Delete(path)
	if err != nil {
		return err
	}
	return nil
}

func (c *IEEntry) Get() (err error) {
	path := fmt.Sprintf("/api/v0/ies/%s", c.ID)

	body, err := GetBody(path)
	if err != nil {
		return err
	}

	err = json.Unmarshal(body, c)
	if err != nil {
		return err
	}

	return nil
}

func (c *IEEntry) Create() (err error) {
	path := "/api/v0/ies"

	data, err := json.Marshal(c)
	if err != nil {
		return err
	}

	body, err := PostReturnBody(path, data)
	if err != nil {
		return err
	}

	err = json.Unmarshal(body, c)
	if err != nil {
		return err
	}

	return nil
}

func (c *IEEntry) Update() (err error) {
	path := "/api/v0/ies/" + c.ID

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

func (c *IEEntry) Delete() (err error) {
	return IEDelete(c.ID)
}

func (e IEListEntry) ToString() string {
	s := fmt.Sprintf("ID: %s, SysNum: %s, Phase: %s, Step: %s, Status: %s, Title: %s, CreatedAt: %s , UpdatedAt: %s, URL: %s, CollectionURL: %s",
		e.ID, e.SysNum, e.Phase, e.Step, e.Status, e.Title, e.CreatedAt, e.UpdatedAt, e.URL, e.CollectionURL)

	return s
}

func (e IEEntry) ToString() string {
	s := fmt.Sprintf("ID: %s, CollectionID: %s, SysNum: %s, Phase: %s, Step: %s, Status: %s, Title: %s, Notes: %s, CreatedAt: %s , UpdatedAt: %s, FMDsURL: %s, CollectionURL: %s, LockVersion: %d", e.ID, e.CollectionID, e.SysNum, e.Phase, e.Step, e.Status, e.Title, e.Notes, e.CreatedAt, e.UpdatedAt, e.FMDsURL, e.CollectionURL, e.LockVersion)

	return s
}
