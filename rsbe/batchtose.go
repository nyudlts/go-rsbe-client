package rsbe

import (
	"encoding/json"
//	"fmt"
)

// #  id           :uuid             not null, primary key
// #  batch_id     :uuid             not null
// #  se_id        :uuid             not null
// #  lock_version :integer
// #  created_at   :datetime
// #  updated_at   :datetime
// #  phase        :string(256)      not null
// #  step         :string(256)      not null
// #  status       :string(256)      not null
// #  notes        :text
// #  data         :text

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
	Data          string `json:"data,omitempty"`
	BatchURL      string `json:"batch_url,omitempty"`
	SEURL         string `json:"se_url,omitempty"`
	BatchToSEsURL string `json:"batch_to_ies_url"`
	LockVersion   int    `json:"lock_version,omitempty"`
}


 // api_v0_batch_to_ses GET    /api/v0/batch_to_ses(.:format)               api/v0/batch_to_ses#index {:format=>"json"}
 //                     POST   /api/v0/batch_to_ses(.:format)               api/v0/batch_to_ses#create {:format=>"json"}
 //  api_v0_batch_to_se GET    /api/v0/batch_to_ses/:id(.:format)           api/v0/batch_to_ses#show {:format=>"json"}
 //                     PATCH  /api/v0/batch_to_ses/:id(.:format)           api/v0/batch_to_ses#update {:format=>"json"}
 //                     PUT    /api/v0/batch_to_ses/:id(.:format)           api/v0/batch_to_ses#update {:format=>"json"}
 //                     DELETE /api/v0/batch_to_ses/:id(.:format)           api/v0/batch_to_ses#destroy {:format=>"json"}


// func BatchToSEList() (list []BatchToSEListEntry, err error) {

// 	body, err := GetBody("/api/v0/batch_to_ses")
// 	if err != nil {
// 		return nil, err
// 	}

// 	err = json.Unmarshal(body, &list)
// 	if err != nil {
// 		return nil, err
// 	}

// 	return list, nil
// }

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

// func (p *BatchToSEEntry) Get() (err error) {
// 	path := "/api/v0/batch_to_ses/" + p.ID

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

// func (c *BatchToSEEntry) Update() (err error) {
// 	path := "/api/v0/batch_to_ses/" + c.ID

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

// func BatchToSEDelete(id string) (err error) {
// 	path := "/api/v0/batch_to_ses/" + id

// 	err = Delete(path)
// 	if err != nil {
// 		return err
// 	}
// 	return nil
// }

// func (c *BatchToSEEntry) Delete() (err error) {
// 	return BatchToSEDelete(c.ID)
// }

// func (e BatchToSEListEntry) ToString() string {
// 	s := fmt.Sprintf("ID: %s, BatchID: %s, EType: %s, FMDID: %s, Role: %s, URL: %s",
// 		e.ID, e.BatchID, e.EType, e.FMDID, e.Role, e.URL)
// 	return s
// }

// func (e BatchToSEEntry) ToString() string {
// 	s := fmt.Sprintf("ID: %s, BatchID: %s, EType: %s, FMDID: %s, Role: %s, CreatedAt: %s, UpdatedAt: %s, EURL: %s, LockVersion: %d",
// 		e.ID, e.BatchID, e.EType, e.FMDID, e.Role, e.CreatedAt, e.UpdatedAt, e.EURL, e.LockVersion)
// 	return s
// }
