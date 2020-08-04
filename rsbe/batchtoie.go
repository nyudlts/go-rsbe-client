package rsbe

import (
	"encoding/json"
		"fmt"
)

type BatchToIEListEntry struct {
	ID        string `json:"id,omitempty"`
	BatchID   string `json:"batch_id,omitempty"`
	IEID      string `json:"ie_id,omitempty"`
	Phase     string `json:"phase,omitempty"`
	Step      string `json:"step,omitempty"`
	Status    string `json:"status,omitempty"`
	CreatedAt string `json:"created_at,omitempty"`
	UpdatedAt string `json:"updated_at,omitempty"`
	URL       string `json:"url,omitempty"`
}

type BatchToIEEntry struct {
	ID            string `json:"id,omitempty"`
	BatchID       string `json:"batch_id,omitempty"`
	IEID          string `json:"ie_id,omitempty"`
	Phase         string `json:"phase,omitempty"`
	Step          string `json:"step,omitempty"`
	Status        string `json:"status,omitempty"`
	Notes         string `json:"notes,omitempty"`
	CreatedAt     string `json:"created_at,omitempty"`
	UpdatedAt     string `json:"updated_at,omitempty"`
	BatchURL      string `json:"batch_url,omitempty"`
	IEURL         string `json:"ie_url,omitempty"`
	BatchToIEsURL string `json:"batch_to_ies_url"`
	LockVersion   int    `json:"lock_version,omitempty"`
}

func BatchToIEList() (list []BatchToIEListEntry, err error) {

	body, err := GetBody("/api/v0/batch_to_ies")
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(body, &list)
	if err != nil {
		return nil, err
	}

	return list, nil
}

func BatchToIEGet(id string) (item BatchToIEEntry, err error) {
	path := "/api/v0/batch_to_ies/" + id

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

func (p *BatchToIEEntry) Get() (err error) {
	path := "/api/v0/batch_to_ies/" + p.ID

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

func (p *BatchToIEEntry) Create() (err error) {
	path := "/api/v0/batch_to_ies"

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

func (c *BatchToIEEntry) Update() (err error) {
	path := "/api/v0/batch_to_ies/" + c.ID

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

func BatchToIEDelete(id string) (err error) {
	path := "/api/v0/batch_to_ies/" + id

	err = Delete(path)
	if err != nil {
		return err
	}
	return nil
}

func (c *BatchToIEEntry) Delete() (err error) {
	return BatchToIEDelete(c.ID)
}

func (e BatchToIEListEntry) ToString() string {
	s := fmt.Sprintf("ID: %s, BatchID: %s, IEID: %s, Phase: %s, Step: %s, Status: %s, CreatedAt: %s, UpdatedAt: %s, URL: %s",
		e.ID, e.BatchID, e.IEID, e.Phase, e.Step, e.Status, e.CreatedAt, e.UpdatedAt, e.URL)
	return s
}

func (e BatchToIEEntry) ToString() string {
	s := fmt.Sprintf("ID: %s, BatchID: %s, IEID: %s, Phase: %s, Step: %s, Status: %s, Notes: %s, CreatedAt: %s, UpdatedAt: %s, BatchURL: %s, BatchToIEsURL: %s, LockVersion: %d",
		e.ID, e.BatchID, e.IEID, e.Phase, e.Step, e.Status, e.Notes, e.CreatedAt, e.UpdatedAt, e.BatchURL, e.BatchToIEsURL, e.LockVersion)
	return s
}
