package rsbe

type SearchResult struct {
	ResponseHeader SearchResponseHeader `json:"responseHeader,omitempty"`
	Response       SearchResponse       `json:"response,omitempty"`
}

type SearchResponseHeader struct {
	Params SearchParams `json:"responseHeader,omitempty"`
}

type SearchResponse struct {
	NumFound uint64      `json:"numFound,omitempty"`
	Start    uint64      `json:"start,omitempty"`
	Docs     []SearchDoc `json:"docs,omitempty"`
}

type SearchDoc struct {
	URL string `json:"url,omitempty"`
}

type SearchParams struct {
	Scope  string `json:"scope,omitempty"`
	DigiID string `json:"digi_id,omitempty"`
}
