package main

//model aligned with v2-series-transformer
type series struct {
	UUID        string       `json:"uuid"`
	ProperName  string       `json:"properName"`
	Type        string       `json:"type"`
	Identifiers []identifier `json:"identifiers,omitempty"`
}

type identifier struct {
	Authority       string `json:"authority"`
	IdentifierValue string `json:"identifierValue"`
}

type seriesLink struct {
	APIURL string `json:"apiUrl"`
}
