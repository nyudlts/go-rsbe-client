package rsbe

import (
	"encoding/json"
	"fmt"
)

type BatchListEntry struct {
	ID            string `json:"id,omitempty"`
	Type          string `json:"batch_type,omitempty"`
	Number        uint   `json:"batch_number,omitempty"`
	Name          string `json:"name,omitempty"`
	Source        string `json:"source_file,omitempty"`
	CreatedAt     string `json:"created_at,omitempty"`
	UpdatedAt     string `json:"updated_at,omitempty"`
	CollectionID  string `json:"coll_id,omitempty"`
	URL           string `json:"url,omitempty"`
	CollectionURL string `json:"coll_url,omitempty"`
}

type BatchEntry struct {
	ID            string `json:"id,omitempty"`
	Name          string `json:"name,omitempty"`
	Source        string `json:"source_file,omitempty"`
	CollectionID  string `json:"coll_id,omitempty"`
	Type          string `json:"batch_type,omitempty"`
	Number        uint   `json:"batch_number,omitempty"`
	Notes         string `json:"notes,omitempty"`
	CreatedAt     string `json:"created_at,omitempty"`
	UpdatedAt     string `json:"updated_at,omitempty"`
	CollectionURL string `json:"coll_url,omitempty"`
	BatchesURL    string `json:"batches_url,omitempty"`
	LockVersion   int    `json:"lock_version,omitempty"`
}

type BatchReport struct {
	TimeStamp string               `json:"report_timestamp"`
	Info      BatchInfo            `json:"batch"`
	SEs       []BatchToSEListEntry `json:"ses"`
	IEs       []BatchToIEListEntry `json:"ies"`
}

type BatchInfo struct {
	ID string `json:"id"`
}

func BatchList() (list []BatchListEntry, err error) {
	path := fmt.Sprintf("/api/v0/batches")

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

func BatchGet(id string) (item BatchEntry, err error) {
	path := fmt.Sprintf("/api/v0/batches/%s", id)

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

func BatchReportGet(id string) (item BatchReport, err error) {
	path := fmt.Sprintf("/api/v0/batches/%s/report", id)

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

func BatchDelete(id string) (err error) {
	path := "/api/v0/batches/" + id

	err = Delete(path)
	if err != nil {
		return err
	}
	return nil
}

func (c *BatchEntry) Get() (err error) {
	path := fmt.Sprintf("/api/v0/batches/%s", c.ID)

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

func (c *BatchEntry) Create() (err error) {
	path := "/api/v0/batches"

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

func (c *BatchEntry) Update() (err error) {
	path := "/api/v0/batches/" + c.ID

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

func (c *BatchEntry) Delete() (err error) {
	return BatchDelete(c.ID)
}

func (e BatchListEntry) ToString() string {
	s := fmt.Sprintf("ID: %s, Type: %s, Number: %d, Name: %s, Source: %s, CollectionID: %s, CreatedAt: %s , UpdatedAt: %s, URL: %s, CollectionURL: %s",
		e.ID, e.Type, e.Number, e.Name, e.Source, e.CollectionID, e.CreatedAt, e.UpdatedAt, e.URL, e.CollectionURL)
	return s
}

func (e BatchEntry) ToString() string {
	s := fmt.Sprintf("ID: %s, Type: %s, Number: %d, Name: %s, Source: %s, CollectionID: %s, CreatedAt: %s , UpdatedAt: %s, CollectionURL: %s, BatchesURL: %s, LockVersion: %d, Notes: %s",
		e.ID, e.Type, e.Number, e.Name, e.Source, e.CollectionID, e.CreatedAt, e.UpdatedAt, e.CollectionURL, e.BatchesURL, e.LockVersion, e.Notes)
	return s
}
