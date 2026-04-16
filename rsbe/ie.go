package rsbe

import (
	"encoding/json"
	"fmt"

	"github.com/google/uuid"
)

type IEListEntry struct {
	ID            string `json:"id,omitempty"`
	CollectionID  string `json:"coll_id,omitempty"` // REQUIRED
	SysNum        string `json:"sys_num,omitempty"` // REQUIRED
	Phase         string `json:"phase,omitempty"`   // REQUIRED
	Step          string `json:"step,omitempty"`    // REQUIRED
	Status        string `json:"status,omitempty"`  // REQUIRED
	Title         string `json:"title,omitempty"`
	CreatedAt     string `json:"created_at,omitempty"`
	UpdatedAt     string `json:"updated_at,omitempty"`
	DeletedAt     string `json:"deleted_at,omitempty"`
	URL           string `json:"url,omitempty"`
	CollectionURL string `json:"coll_url,omitempty"`
	LockVersion   int    `json:"lock_version"`
}

type IEEntry struct {
	ID            string        `json:"id,omitempty"`
	CollectionID  string        `json:"coll_id,omitempty"` // REQUIRED
	SysNum        string        `json:"sys_num,omitempty"` // REQUIRED
	Phase         string        `json:"phase,omitempty"`   // REQUIRED
	Step          string        `json:"step,omitempty"`    // REQUIRED
	Status        string        `json:"status,omitempty"`  // REQUIRED
	Title         string        `json:"title,omitempty"`
	Notes         string        `json:"notes,omitempty"`
	CreatedAt     string        `json:"created_at,omitempty"`
	UpdatedAt     string        `json:"updated_at,omitempty"`
	DeletedAt     string        `json:"deleted_at,omitempty"`
	Fids          JSONMap       `json:"fids,omitempty"`
	SEs           []SEForIEShow `json:"ses,omitempty"`
	FMDsURL       string        `json:"fmds_url,omitempty"`
	CollectionURL string        `json:"coll_url,omitempty"`
	LockVersion   int           `json:"lock_version"`
}

type IEforSEShow struct {
	SysNum    string    `json:"sys_num"`
	Order     int       `json:"order"`
	Section   int       `json:"section"`
	IEID      uuid.UUID `json:"ie_id"`
	Notes     string    `json:"notes"`
	IEURL     string    `json:"ie_url"`
	IEToSEURL string    `json:"ie_to_se_url"`
}

func CollectionIEList(collectionID string) (list []IEListEntry, err error) {
	path := fmt.Sprintf("/api/v0/colls/%s/ies", collectionID)

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
	_, err = uuid.Parse(id)
	if err != nil {
		return IEEntry{}, fmt.Errorf("id is not a UUID: %s", err.Error())
	}

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
	_, err = uuid.Parse(id)
	if err != nil {
		return fmt.Errorf("id is not a UUID: %s", err.Error())
	}

	path := "/api/v0/ies/" + id

	err = Delete(path)
	if err != nil {
		return err
	}
	return nil
}

func IEPurge(id string) (err error) {
	_, err = uuid.Parse(id)
	if err != nil {
		return fmt.Errorf("id is not a UUID: %s", err.Error())
	}

	path := "/api/v0/ies/" + id

	err = Purge(path)
	if err != nil {
		return err
	}
	return nil
}

func (c *IEEntry) Get() (err error) {
	_, err = uuid.Parse(c.ID)
	if err != nil {
		return fmt.Errorf("ID is not a UUID: %s", err.Error())
	}

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
	_, err = uuid.Parse(c.ID)
	if err != nil {
		return fmt.Errorf("ID is not a UUID: %s", err.Error())
	}

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

func (c *IEEntry) Purge() (err error) {
	return IEPurge(c.ID)
}

func (e IEListEntry) ToString() string {
	s := fmt.Sprintf("ID: %s, SysNum: %s, Phase: %s, Step: %s, Status: %s, Title: %s, CreatedAt: %s, UpdatedAt: %s, URL: %s, CollectionURL: %s",
		e.ID, e.SysNum, e.Phase, e.Step, e.Status, e.Title, e.CreatedAt, e.UpdatedAt, e.URL, e.CollectionURL)

	return s
}

func (e IEEntry) ToString() string {
	s := fmt.Sprintf("ID: %s, CollectionID: %s, SysNum: %s, Phase: %s, Step: %s, Status: %s, Title: %s, Notes: %s, CreatedAt: %s, UpdatedAt: %s, FMDsURL: %s, CollectionURL: %s, LockVersion: %d", e.ID, e.CollectionID, e.SysNum, e.Phase, e.Step, e.Status, e.Title, e.Notes, e.CreatedAt, e.UpdatedAt, e.FMDsURL, e.CollectionURL, e.LockVersion)

	return s
}
