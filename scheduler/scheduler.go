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
	fmt.Println(currentPeriod)
	return currentPeriod

}

//EvaluateSchedule determines the current period and returns the
//corresponding image file.
func EvaluateSchedule(config Config) string {
	currentTime := time.Now().Format(time.Kitchen)
	for period := range config.Period {
		for k, v := range config.Period[period] {

			fmt.Printf("Period: %v\nStartTime: %v\nEndTime: %v\n", k, v.StartTime, v.EndTime)
			if evaluatePeriod(currentTime, v.StartTime, v.EndTime, k) != "" {
				fmt.Printf("Current period is: %v", k)
				return v.ImageFile
			}
		}
	}
	return ""
}

func evaluatePeriod(now string, start string, end string, period string) string {
	fmt.Printf("Current time is %v\n", now)
	if now > start && now < end {
		fmt.Printf("We are currently within the %v period\n", period)
		return period
	}
	return ""
}
