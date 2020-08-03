package rsbe

import (
	"encoding/json"
		"fmt"
)

type BatchToSEListEntry struct {
	ID        string `json:"id,omitempty"`
	BatchID   string `json:"batch_id,omitempty"`
	SEID      string `json:"se_id,omitempty"`
	Phase     string `json:"phase,omitempty"`
	Step      string `json:"step,omitempty"`
	Status    string `json:"status,omitempty"`
	CreatedAt string `json:"created_at,omitempty"`
	UpdatedAt string `json:"updated_at,omitempty"`
	URL       string `json:"url,omitempty"`
}

type BatchToSEEntry struct {
	ID            string `json:"id,omitempty"`
	BatchID       string `json:"batch_id,omitempty"`
	SEID          string `json:"se_id,omitempty"`
	Phase         string `json:"phase,omitempty"`
	Step          string `json:"step,omitempty"`
	Status        string `json:"status,omitempty"`
	Notes         string `json:"notes,omitempty"`
	CreatedAt     string `json:"created_at,omitempty"`
	UpdatedAt     string `json:"updated_at,omitempty"`
	BatchURL      string `json:"batch_url,omitempty"`
	SEURL         string `json:"se_url,omitempty"`
	BatchToSEsURL string `json:"batch_to_ses_url"`
	LockVersion   int    `json:"lock_version,omitempty"`
}

func BatchToSEList() (list []BatchToSEListEntry, err error) {

	body, err := GetBody("/api/v0/batch_to_ses")
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(body, &list)
	if err != nil {
		return nil, err
	}

	return list, nil
}

func BatchToSEGet(id string) (item BatchToSEEntry, err error) {
	path := "/api/v0/batch_to_ses/" + id

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

func (p *BatchToSEEntry) Get() (err error) {
	path := "/api/v0/batch_to_ses/" + p.ID

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

func (p *BatchToSEEntry) Create() (err error) {
	path := "/api/v0/batch_to_ses"

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

func (c *BatchToSEEntry) Update() (err error) {
	path := "/api/v0/batch_to_ses/" + c.ID

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

func BatchToSEDelete(id string) (err error) {
	path := "/api/v0/batch_to_ses/" + id

	err = Delete(path)
	if err != nil {
		return err
	}
	return nil
}

func (c *BatchToSEEntry) Delete() (err error) {
	return BatchToSEDelete(c.ID)
}

func (e BatchToSEListEntry) ToString() string {
	s := fmt.Sprintf("ID: %s, BatchID: %s, SEID: %s, Phase: %s, Step: %s, Status: %s, CreatedAt: %s, UpdatedAt: %s, URL: %s",
		e.ID, e.BatchID, e.SEID, e.Phase, e.Step, e.Status, e.CreatedAt, e.UpdatedAt, e.URL)
	return s
}

func (e BatchToSEEntry) ToString() string {
	s := fmt.Sprintf("ID: %s, BatchID: %s, SEID: %s, Phase: %s, Step: %s, Status: %s, Notes: %s, CreatedAt: %s, UpdatedAt: %s, BatchURL: %s, BatchToSEsURL: %s, LockVersion: %d",
		e.ID, e.BatchID, e.SEID, e.Phase, e.Step, e.Status, e.Notes, e.CreatedAt, e.UpdatedAt, e.BatchURL, e.BatchToSEsURL, e.LockVersion)
	return s
}
