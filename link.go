package entrez

// ELink Responds to a list of UIDs in a given database with either a list of
// related UIDs (and relevancy scores) in the same database or a list of linked
// UIDs in another Entrez database; checks for the existence of a specified link
// from a list of one or more UIDs; creates a hyperlink to the primary LinkOut
// provider for a specific UID and database, or lists LinkOut URLs and
// attributes for multiple UIDs.
const ELinkURL = BaseURL + "elink.fcgi"
