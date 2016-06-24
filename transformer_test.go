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
				UUID:      "adb4f804-c3b6-3eca-8708-5edeec653a27",
				PrefLabel: "Africa Series",
				AlternativeIdentifiers: alternativeIdentifiers{
					TME:   []string{"TnN0ZWluX0dMX0FGVE1fR0xfMTY0ODM1-U2VjdGlvbnM="},
					Uuids: []string{"adb4f804-c3b6-3eca-8708-5edeec653a27"},
				},
				Type: "Series"}},
	}

	for _, test := range tests {
		expectedSeries := transformSeries(test.term, "Series")
		assert.Equal(test.series, expectedSeries, fmt.Sprintf("%s: Expected series incorrect", test.name))
	}

}
