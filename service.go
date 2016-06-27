package main

import (
	"net/http"

	"github.com/Financial-Times/tme-reader/tmereader"
	log "github.com/Sirupsen/logrus"
)

type httpClient interface {
	Do(req *http.Request) (resp *http.Response, err error)
}

type seriesService interface {
	getSeries() ([]seriesLink, bool)
	getSeriesByUUID(uuid string) (series, bool)
	checkConnectivity() error
}

type seriesServiceImpl struct {
	repository    tmereader.Repository
	baseURL       string
	seriesMap     map[string]series
	seriesLinks   []seriesLink
	taxonomyName  string
	maxTmeRecords int
}

func newSeriesService(repo tmereader.Repository, baseURL string, taxonomyName string, maxTmeRecords int) (seriesService, error) {
	s := &seriesServiceImpl{repository: repo, baseURL: baseURL, taxonomyName: taxonomyName, maxTmeRecords: maxTmeRecords}
	err := s.init()
	if err != nil {
		return &seriesServiceImpl{}, err
	}
	return s, nil
}

func (s *seriesServiceImpl) init() error {
	s.seriesMap = make(map[string]series)
	responseCount := 0
	log.Printf("Fetching series from TME\n")
	for {
		terms, err := s.repository.GetTmeTermsFromIndex(responseCount)
		if err != nil {
			return err
		}

		if len(terms) < 1 {
			log.Printf("Finished fetching series from TME\n")
			break
		}
		s.initSeriesMap(terms)
		responseCount += s.maxTmeRecords
	}
	log.Printf("Added %d series links\n", len(s.seriesLinks))

	return nil
}

func (s *seriesServiceImpl) getSeries() ([]seriesLink, bool) {
	if len(s.seriesLinks) > 0 {
		return s.seriesLinks, true
	}
	return s.seriesLinks, false
}

func (s *seriesServiceImpl) getSeriesByUUID(uuid string) (series, bool) {
	series, found := s.seriesMap[uuid]
	return series, found
}

func (s *seriesServiceImpl) checkConnectivity() error {
	// TODO: Can we just hit an endpoint to check if TME is available? Or do we need to make sure we get genre taxonmies back? Maybe a healthcheck or gtg endpoint?
	// TODO: Can we use a count from our responses while actually in use to trigger a healthcheck?
	//	_, err := s.repository.GetTmeTermsFromIndex(1)
	//	if err != nil {
	//		return err
	//	}
	return nil
}

func (s *seriesServiceImpl) initSeriesMap(terms []interface{}) {
	for _, iTerm := range terms {
		t := iTerm.(term)
		top := transformSeries(t, s.taxonomyName)
		s.seriesMap[top.UUID] = top
		s.seriesLinks = append(s.seriesLinks, seriesLink{APIURL: s.baseURL + top.UUID})
	}
}
