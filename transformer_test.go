package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTransform(t *testing.T) {
	assert := assert.New(t)
	tests := []struct {
		name   string
		term   term
		series series
	}{
		{"Transform term to series", term{
			CanonicalName: "Africa Series",
			RawID:         "Nstein_GL_AFTM_GL_164835"},
			series{
				UUID:      "56a141a4-9894-3559-b25b-d0142f8148ff",
				PrefLabel: "Africa Series",
				AlternativeIdentifiers: alternativeIdentifiers{
					TME:   []string{"TnN0ZWluX0dMX0FGVE1fR0xfMTY0ODM1-U2VyaWVz"},
					Uuids: []string{"56a141a4-9894-3559-b25b-d0142f8148ff"},
				},
				Type: "Series"}},
	}

	for _, test := range tests {
		expectedSeries := transformSeries(test.term, "Series")
		assert.Equal(test.series, expectedSeries, fmt.Sprintf("%s: Expected series incorrect", test.name))
	}

}
