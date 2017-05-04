package entrez

// ESearch responds to a text query with the list of matching UIDs in a given
// database (for later use in ESummary, EFetch or ELink), along with the term
// translations of the query.
const ESearchURL = BaseURL + "esearch.fcgi"
