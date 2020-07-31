package rsbe

import (
	"encoding/json"
	"fmt"
//	"path/filepath"
)

type BatchListEntry struct {
	ID            string `json:"id,omitempty"`
	BatchType     string `json:"batch_type,omitempty"`
	BatchNumber   uint   `json:"batch_number,omitempty"`
	Name          string `json:"name,omitempty"`
	SourceFile    string `json:"source_file,omitempty"`
	CreatedAt     string `json:"created_at,omitempty"`
	UpdatedAt     string `json:"updated_at,omitempty"`
	CollectionID  string `json:"coll_id,omitempty"`
	URL           string `json:"url,omitempty"`
	CollectionURL string `json:"coll_url,omitempty"`
}

type BatchEntry struct {
	ID            string `json:"id,omitempty"`
	Name          string `json:"name,omitempty"`
	SourceFile    string `json:"source_file,omitempty"`
	CollectionID  string `json:"coll_id,omitempty"`
	BatchType     string `json:"batch_type,omitempty"`
	BatchNumber   uint   `json:"batch_number,omitempty"`
	Notes         string `json:"notes,omitempty"`
	CreatedAt     string `json:"created_at,omitempty"`
	UpdatedAt     string `json:"updated_at,omitempty"`
	CollectionURL string `json:"coll_url,omitempty"`
	BatchesURL    string `json:"batches_url,omitempty"`
	LockVersion   int    `json:"lock_version,omitempty"`
}

//    api_v0_batches GET    /api/v0/batches(.:format)                    api/v0/batches#index {:format=>"json"}
//                   POST   /api/v0/batches(.:format)                    api/v0/batches#create {:format=>"json"}
//      api_v0_batch GET    /api/v0/batches/:id(.:format)                api/v0/batches#show {:format=>"json"}
//                   PATCH  /api/v0/batches/:id(.:format)                api/v0/batches#update {:format=>"json"}
//                   PUT    /api/v0/batches/:id(.:format)                api/v0/batches#update {:format=>"json"}
//                   DELETE /api/v0/batches/:id(.:format)                api/v0/batches#destroy {:format=>"json"}

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

// func BatchGetByDigiID(digiID string) (item BatchEntry, err error) {
// 	path := fmt.Sprintf("/api/v0/search?scope=ses&digi_id=%s", digiID)

// 	var searchResult SearchResult

// 	body, err := GetBody(path)
// 	if err != nil {
// 		return item, err
// 	}

// 	err = json.Unmarshal(body, &searchResult)
// 	if err != nil {
// 		return item, err
// 	}

// 	if searchResult.Response.NumFound != 1 {
// 		return item, fmt.Errorf("Incorrect number of results. Expected 1, found %d", searchResult.Response.NumFound)
// 	}

// 	id := filepath.Base(searchResult.Response.Docs[0].URL)
// 	return BatchGet(id)
// }

// func BatchDelete(id string) (err error) {
// 	path := "/api/v0/ses/" + id

// 	err = Delete(path)
// 	if err != nil {
// 		return err
// 	}
// 	return nil
// }

// func (c *BatchEntry) Get() (err error) {
// 	path := fmt.Sprintf("/api/v0/ses/%s", c.ID)

// 	body, err := GetBody(path)
// 	if err != nil {
// 		return err
// 	}

// 	err = json.Unmarshal(body, c)
// 	if err != nil {
// 		return err
// 	}

// 	return nil
// }

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

// func (c *BatchEntry) Update() (err error) {
// 	path := "/api/v0/ses/" + c.ID

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

// func (c *BatchEntry) Delete() (err error) {
// 	return BatchDelete(c.ID)
// }

// func (e BatchListEntry) ToString() string {
// 	s := fmt.Sprintf("ID: %s, DigiID: %s, DOType: %s, Phase: %s, Step: %s, Status: %s, Label: %s, CreatedAt: %s , UpdatedAt: %s, URL: %s, CollectionURL: %s",
// 		e.ID, e.DigiID, e.DOType, e.Phase, e.Step, e.Status, e.Label, e.CreatedAt, e.UpdatedAt, e.URL, e.CollectionURL)

// 	return s
// }

// func (e BatchEntry) ToString() string {
// 	s := fmt.Sprintf("ID: %s, DigiID: %s, DOType: %s, Phase: %s, Step: %s, Status: %s, Label: %s, Title: %s, CreatedAt: %s , UpdatedAt: %s, BDIURL: %s, FMDsURL: %s, CollectionURL: %s, LockVersion: %d, Notes: %s", e.ID, e.DigiID, e.DOType, e.Phase, e.Step, e.Status, e.Label, e.Title, e.CreatedAt, e.UpdatedAt, e.BDIURL, e.FMDsURL, e.CollectionURL, e.LockVersion, e.Notes)

// 	return s
// }
