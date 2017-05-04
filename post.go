package entrez

// EPost accepts a list of UIDs from a given database, stores the set on the
// History Server, and responds with a query key and web environment for the
// uploaded dataset.
const EPostURL = BaseURL + "epost.fcgi"

func EPost(uIDs []UID) (QueryKey, error) {
	return QueryKey(""), nil
}
