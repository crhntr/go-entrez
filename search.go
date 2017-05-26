package entrez

import (
	"encoding/xml"
	"fmt"
	"net/http"
	"net/url"
)

// ESearch responds to a text query with the list of matching UIDs in a given
// database (for later use in ESummary, EFetch or ELink), along with the term
// translations of the query.
const ESearchURL = BaseURL + "esearch.fcgi"

func ESearch(query string, start int, max int) (ESearchResult, error) {
	res := ESearchResult{}
	query = url.QueryEscape(query)

	r, err := http.Get(
		fmt.Sprintf(
			BaseURL+"esearch.fcgi?db=%s&term=%s&retstart=%d&retmax=%d&usehistory=y",
			DBName, query, start, max))
	if err != nil {
		return res, err
	}

	dec := xml.NewDecoder(r.Body)
	err = dec.Decode(&res)
	return res, err
}

type ESearchResult struct {
	Count    int
	RetMax   int
	RetStart int
	QueryKey string
	WebEnv   string
	IDs      []string `xml:"IdList>Id"`
}
