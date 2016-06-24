package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
)

const testUUID = "bba39990-c78d-3629-ae83-808c333c6dbc"
const getSeriesResponse = `[{"apiUrl":"http://localhost:8080/transformers/series/bba39990-c78d-3629-ae83-808c333c6dbc"}]`
const getSeriesByUUIDResponse = `{"uuid":"bba39990-c78d-3629-ae83-808c333c6dbc","alternativeIdentifiers":{"TME":["MTE3-U3ViamVjdHM="],"uuids":["bba39990-c78d-3629-ae83-808c333c6dbc"]},"prefLabel":"Global Series","type":"Series"}`

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
		{"Success - get series by uuid", newRequest("GET", fmt.Sprintf("/transformers/series/%s", testUUID)), &dummyService{found: true, series: []series{getDummySeries(testUUID, "Global Series", "MTE3-U3ViamVjdHM=")}}, http.StatusOK, "application/json", getSeriesByUUIDResponse},
		{"Not found - get series by uuid", newRequest("GET", fmt.Sprintf("/transformers/series/%s", testUUID)), &dummyService{found: false, series: []series{series{}}}, http.StatusNotFound, "application/json", ""},
		{"Success - get series", newRequest("GET", "/transformers/series"), &dummyService{found: true, series: []series{series{UUID: testUUID}}}, http.StatusOK, "application/json", getSeriesResponse},
		{"Not found - get series", newRequest("GET", "/transformers/series"), &dummyService{found: false, series: []series{}}, http.StatusNotFound, "application/json", ""},
	}

	for _, test := range tests {
		rec := httptest.NewRecorder()
		router(test.dummyService).ServeHTTP(rec, test.req)
		assert.True(test.statusCode == rec.Code, fmt.Sprintf("%s: Wrong response code, was %d, should be %d", test.name, rec.Code, test.statusCode))
		assert.Equal(strings.TrimSpace(test.body), strings.TrimSpace(rec.Body.String()), fmt.Sprintf("%s: Wrong body", test.name))
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
	m.HandleFunc("/transformers/series/{uuid}", h.getSeriesByUUID).Methods("GET")
	return m
}

type dummyService struct {
	found  bool
	series []series
}

func (s *dummyService) getSeries() ([]seriesLink, bool) {
	var seriesLinks []seriesLink
	for _, sub := range s.series {
		seriesLinks = append(seriesLinks, seriesLink{APIURL: "http://localhost:8080/transformers/series/" + sub.UUID})
	}
	return seriesLinks, s.found
}

func (s *dummyService) getSeriesByUUID(uuid string) (series, bool) {
	return s.series[0], s.found
}

func (s *dummyService) checkConnectivity() error {
	return nil
}
