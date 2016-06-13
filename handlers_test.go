package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

const testUUID = "bba39990-c78d-3629-ae83-808c333c6dbc"
const getSeriesResponse = "[{\"apiUrl\":\"http://localhost:8080/transformers/series/bba39990-c78d-3629-ae83-808c333c6dbc\"}]\n"
const getSeriesByUUIDResponse = "{\"uuid\":\"bba39990-c78d-3629-ae83-808c333c6dbc\",\"properName\":\"European Union\",\"type\":\"Series\",\"identifiers\":[{" +
	"\"authority\":\"http://api.ft.com/system/FT-TME\"," +
	"\"identifierValue\":\"MTE3-U3ViamVjdHM=\"" +
	"}]}\n"

func TestHandlers(t *testing.T) {
	assert := assert.New(t)
	tests := []struct {
		name         string
		req          *http.Request
		dummyService seriesService
		statusCode   int
		contentType  string // Contents of the Content-Type header
		body         string
	}{
		{"Success - get series by uuid", newRequest("GET", fmt.Sprintf("/transformers/series/%s", testUUID)), &dummyService{found: true, initialised: true, series: []series{series{UUID: testUUID, ProperName: "European Union", Identifiers: []identifier{identifier{Authority: "http://api.ft.com/system/FT-TME", IdentifierValue: "MTE3-U3ViamVjdHM="}}, Type: "Series"}}}, http.StatusOK, "application/json", getSeriesByUUIDResponse},
		{"Not found - get series by uuid", newRequest("GET", fmt.Sprintf("/transformers/series/%s", testUUID)), &dummyService{found: false, initialised: true, series: []series{series{}}}, http.StatusNotFound, "application/json", ""},
		{"Service unavailable - get series by uuid", newRequest("GET", fmt.Sprintf("/transformers/series/%s", testUUID)), &dummyService{found: false, initialised: false, series: []series{}}, http.StatusServiceUnavailable, "application/json", ""},
		{"Success - get series", newRequest("GET", "/transformers/series"), &dummyService{found: true, initialised: true, series: []series{series{UUID: testUUID}}}, http.StatusOK, "application/json", getSeriesResponse},
		{"Not found - get series", newRequest("GET", "/transformers/series"), &dummyService{found: false, initialised: true, series: []series{}}, http.StatusNotFound, "application/json", ""},
		{"Service unavailable - get series", newRequest("GET", "/transformers/series"), &dummyService{found: false, initialised: false, series: []series{}}, http.StatusServiceUnavailable, "application/json", ""},
	}

	for _, test := range tests {
		rec := httptest.NewRecorder()
		router(test.dummyService).ServeHTTP(rec, test.req)
		assert.True(test.statusCode == rec.Code, fmt.Sprintf("%s: Wrong response code, was %d, should be %d", test.name, rec.Code, test.statusCode))
		assert.Equal(test.body, rec.Body.String(), fmt.Sprintf("%s: Wrong body", test.name))
	}
}

func newRequest(method, url string) *http.Request {
	req, err := http.NewRequest(method, url, nil)
	if err != nil {
		panic(err)
	}
	return req
}

func router(s seriesService) *mux.Router {
	m := mux.NewRouter()
	h := newSeriesHandler(s)
	m.HandleFunc("/transformers/series", h.getSeries).Methods("GET")
	m.HandleFunc("/transformers/series/{uuid}", h.getOrgByUUID).Methods("GET")
	return m
}

type dummyService struct {
	found       bool
	series        []series
	initialised bool
}

func (s *dummyService) getSeries() ([]seriesLink, bool) {
	var seriesLinks []seriesLink
	for _, sub := range s.series {
		seriesLinks = append(seriesLinks, seriesLink{APIURL: "http://localhost:8080/transformers/series/" + sub.UUID})
	}
	return seriesLinks, s.found
}

func (s *dummyService) getOrgByUUID(uuid string) (series, bool, error) {
	return s.series[0], s.found, nil
}

func (s *dummyService) isInitialised() bool {
	return s.initialised
}
