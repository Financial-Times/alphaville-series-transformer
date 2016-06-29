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
const getAlphavilleSeriesResponse = `[{"apiUrl":"http://localhost:8080/transformers/alphavilleseries/bba39990-c78d-3629-ae83-808c333c6dbc"}]`
const getAlphavilleSeriesByUUIDResponse = `{"uuid":"bba39990-c78d-3629-ae83-808c333c6dbc","alternativeIdentifiers":{"TME":["MTE3-U3ViamVjdHM="],"uuids":["bba39990-c78d-3629-ae83-808c333c6dbc"]},"prefLabel":"Global Alphaville Series","type":"AlphavilleSeries"}`

func TestHandlers(t *testing.T) {
	assert := assert.New(t)
	tests := []struct {
		name         string
		req          *http.Request
		dummyService alphavilleSeriesService
		statusCode   int
		contentType  string // Contents of the Content-Type header
		body         string
	}{
		{"Success - get alphavilleSeries by uuid", newRequest("GET", fmt.Sprintf("/transformers/alphavilleseries/%s", testUUID)), &dummyService{found: true, alphavilleSeries: []alphavilleSeries{getDummyAlphavilleSeries(testUUID, "Global Alphaville Series", "MTE3-U3ViamVjdHM=")}}, http.StatusOK, "application/json", getAlphavilleSeriesByUUIDResponse},
		{"Not found - get alphavilleSeries by uuid", newRequest("GET", fmt.Sprintf("/transformers/alphavilleseries/%s", testUUID)), &dummyService{found: false, alphavilleSeries: []alphavilleSeries{alphavilleSeries{}}}, http.StatusNotFound, "application/json", ""},
		{"Success - get alphavilleSeries", newRequest("GET", "/transformers/alphavilleseries"), &dummyService{found: true, alphavilleSeries: []alphavilleSeries{alphavilleSeries{UUID: testUUID}}}, http.StatusOK, "application/json", getAlphavilleSeriesResponse},
		{"Not found - get alphavilleSeries", newRequest("GET", "/transformers/alphavilleseries"), &dummyService{found: false, alphavilleSeries: []alphavilleSeries{}}, http.StatusNotFound, "application/json", ""},
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

func router(s alphavilleSeriesService) *mux.Router {
	m := mux.NewRouter()
	h := newAlphavilleSeriesHandler(s)
	m.HandleFunc("/transformers/alphavilleseries", h.getAlphavilleSeries).Methods("GET")
	m.HandleFunc("/transformers/alphavilleseries/{uuid}", h.getAlphavilleSeriesByUUID).Methods("GET")
	return m
}

type dummyService struct {
	found            bool
	alphavilleSeries []alphavilleSeries
}

func (s *dummyService) getAlphavilleSeries() ([]alphavilleSeriesLink, bool) {
	var alphavilleSeriesLinks []alphavilleSeriesLink
	for _, sub := range s.alphavilleSeries {
		alphavilleSeriesLinks = append(alphavilleSeriesLinks, alphavilleSeriesLink{APIURL: "http://localhost:8080/transformers/alphavilleseries/" + sub.UUID})
	}
	return alphavilleSeriesLinks, s.found
}

func (s *dummyService) getAlphavilleSeriesByUUID(uuid string) (alphavilleSeries, bool) {
	return s.alphavilleSeries[0], s.found
}

func (s *dummyService) checkConnectivity() error {
	return nil
}
