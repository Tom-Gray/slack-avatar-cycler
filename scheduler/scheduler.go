package scheduler

import (
	"fmt"
	"github.com/spf13/viper"
	"time"
)

type Config struct {
	Period []map[string]PeriodSchedules `mapstructure:"schedule"`
}

type PeriodSchedules struct {
	ImageFile string `mapstructure:"imageFile"`
	EndTime   string `mapstructure:"endTime"`
	StartTime string `mapstructure:"startTime"`
}

func GetImageForSchedulePeriod() string {
	viper.SetConfigName("config")
	viper.AddConfigPath(".")
	if err := viper.ReadInConfig(); err != nil {
		fmt.Println(err)

	}
	var config Config
	if err := viper.Unmarshal(&config); err != nil {
		fmt.Println(err)
		return ""

	}
	fmt.Println(config)
	currentPeriod := EvaluateSchedule(config)
	fmt.Printf("current period is: %v\n", currentPeriod)
	return currentPeriod

}

// EvaluateSchedule determines the current period and returns the
// corresponding image file.
func EvaluateSchedule(config Config) string {

	for period := range config.Period {
		for k, v := range config.Period[period] {

			startTime := constructTimeValues(v.StartTime)
			endTime := constructTimeValues(v.EndTime)
			currentTime := time.Now()
			if evaluatePeriodNew(currentTime, startTime, endTime, k) != "" {
				fmt.Printf("Current period is: %v\n", k)
				return v.ImageFile
			}
		}
	}
	return ""
}

// constructTimeValues converts the stringified "kitchen" value to a time object.
func constructTimeValues(timeString string) time.Time {
	now := time.Now()
	timeValue, err := time.ParseInLocation("2006-01-02 15:04",
		fmt.Sprintf("%d-%d-%d %s", now.Year(), now.Month(), now.Day(), timeString), time.Local)
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	return timeValue
}

func evaluatePeriodNew(now time.Time, start time.Time, end time.Time, period string) string {
	fmt.Printf("Current time is %v\n", now)

	// validate timings.
	if start.After(end) {
		fmt.Printf("period %v probably extends past midnight. Adjusting.\n ", period)
		// extend the enddate 24 hours to signal that it ends tomorrow
		end = end.Add(24 * time.Hour)
	}

	// if startTime is before now
	if start.Before(now) {
		if end.After(now) {
			fmt.Printf("we are in the period: %v\v", period)
			return period
		}

		fmt.Printf("We are not in period: %v\n", period)
		return ""
	}

	return "default"
}
