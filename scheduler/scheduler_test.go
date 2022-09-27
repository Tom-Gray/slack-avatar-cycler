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
			currentTime: "09:01AM",
			startTime:   "09:00AM",
			endTime:     "10:00AM",
			period:      "morning",
		},
		{
			name:        "should return lunch period",
			currentTime: "12:00PM",
			startTime:   "10:00AM",
			endTime:     "1:00PM",
			period:      "lunch",
		},
		{
			name:        "should return afternoon period",
			currentTime: "2:00PM",
			startTime:   "1:00PM",
			endTime:     "5:00PM",
			period:      "afternoon",
		},
		{
			name:        "should return night period",
			currentTime: "11:00PM",
			startTime:   "10:00PM",
			endTime:     "6:00AM",
			period:      "night",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			out := evaluatePeriod(test.currentTime, test.startTime, test.endTime, test.period)
			require.Equal(t, test.period, out)
		})
	}
}
