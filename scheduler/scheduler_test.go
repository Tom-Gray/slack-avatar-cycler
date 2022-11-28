package scheduler

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestEvaluatePeriod(t *testing.T) {
	tests := []struct {
		name        string
		currentTime string
		startTime   string
		endTime     string
		period      string
	}{
		{
			name:        "should return morning period",
			currentTime: "09:36",
			startTime:   "08:00",
			endTime:     "10:00",
			period:      "morning",
		},
		{
			name:        "should return lunch period",
			currentTime: "12:00",
			startTime:   "10:00",
			endTime:     "13:00",
			period:      "lunch",
		},
		{
			name:        "should return afternoon period",
			currentTime: "14:00",
			startTime:   "13:00",
			endTime:     "17:00",
			period:      "arvo",
		},
		{
			name:        "should return night period",
			currentTime: "23:00",
			startTime:   "22:00",
			endTime:     "06:00",
			period:      "night",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			current := constructTimeValues(test.currentTime)
			start := constructTimeValues(test.startTime)
			end := constructTimeValues(test.endTime)
			out := evaluatePeriodNew(current, start, end, test.period)
			require.Equal(t, test.period, out)
		})
	}
}
