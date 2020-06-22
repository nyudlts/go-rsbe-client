package rsbe

// import (
// 	"encoding/json"
// )

type CollectionIndexEntry struct {
	Id         string `json:"id"`
	PartnerId  string `json:"partner_id"`
	Code       string `json:"code"`
	Name       string `json:"name"`
	CollType   string `json:"coll_type"`
	Created_at string `json:"created_at"`
	Updated_at string `json:"updated_at"`
	Url        string `json:"url"`
	PartnerUrl string `json:"partner_url"`
}


