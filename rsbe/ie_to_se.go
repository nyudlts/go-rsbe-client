package rsbe

import (
	"github.com/google/go-querystring/query"

	"encoding/json"
	"fmt"
)

type IEToSEListEntry struct {
	ID        string `json:"id,omitempty"         url:"id,omitempty"`
	IEID      string `json:"ie_id,omitempty"      url:"ie_id,omitempty"`   // REQUIRED
	SEID      string `json:"se_id,omitempty"      url:"se_id,omitempty"`   // REQUIRED
	Order     int    `json:"order,omitempty"      url:"order,omitempty"`   // REQUIRED if != 1
	Section   int    `json:"section,omitempty"    url:"section,omitempty"` // REQUIRED if != 1
	CreatedAt string `json:"created_at,omitempty" url:"-"`
	UpdatedAt string `json:"updated_at,omitempty" url:"-"`
	URL       string `json:"url,omitempty"        url:"-"`
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
	IEToSEsURL  string `json:"ie_to_ses_url,omitempty"`
}

// Get a list of IEToSEListEntry objects
// Function accepts 0 or 1 IEToSEListEntry parameters.

// If an IEToSEListEntry argument is passed, the populated fields in the argument
// are added to the RSBE query as query params as per the "url:" struct tags.
func IEToSEList(filter ...IEToSEListEntry) (list []IEToSEListEntry, err error) {

	path := "/api/v0/ie_to_ses"

	// check if there are any query parameters
	switch len(filter) {
	case 0:
		// noop
	case 1:
		// extract url.Values
		v, err := query.Values(filter[0])
		if err != nil {
			return nil, err
		}

		// if encoded values are not empty, append to path
		if v.Encode() != "" {
			path += "?" + v.Encode()
		}
	default:
		return list, fmt.Errorf("error: can only accept 0 or 1 IEToSEListEntry arguments")
	}

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

func IEToSEGet(id string) (item IEToSEEntry, err error) {
	path := fmt.Sprintf("/api/v0/ie_to_ses/%s", id)

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

func IEToSEDelete(id string) (err error) {
	path := "/api/v0/ie_to_ses/" + id

	err = Delete(path)
	if err != nil {
		return err
	}
	return nil
}

func (c *IEToSEEntry) Get() (err error) {
	path := fmt.Sprintf("/api/v0/ie_to_ses/%s", c.ID)

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

func (c *IEToSEEntry) Create() (err error) {
	path := "/api/v0/ie_to_ses"

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

func (c *IEToSEEntry) Update() (err error) {
	path := "/api/v0/ie_to_ses/" + c.ID

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

func (c *IEToSEEntry) Delete() (err error) {
	return IEToSEDelete(c.ID)
}

func (e IEToSEListEntry) ToString() string {
	s := fmt.Sprintf("ID: %s, IEID: %s, SEID: %s, Order: %d, Section: %d, CreatedAt: %s , UpdatedAt: %s, URL: %s",
		e.ID, e.IEID, e.SEID, e.Order, e.Section, e.CreatedAt, e.UpdatedAt, e.URL)

	return s
}

func (e IEToSEEntry) ToString() string {
	s := fmt.Sprintf("ID: %s, IEID: %s, SEID: %s, Order: %d, Section: %d, Notes: %s, CreatedAt: %s , UpdatedAt: %s, IEURL: %s, SEURL: %s, IEToSEsURL: %s, LockVersion: %d", e.ID, e.IEID, e.SEID, e.Order, e.Section, e.Notes, e.CreatedAt, e.UpdatedAt, e.IEURL, e.SEURL, e.IEToSEsURL, e.LockVersion)

	return s
}
