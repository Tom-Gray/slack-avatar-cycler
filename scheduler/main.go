package scheduler

import (
	"fmt"
	"time"

	"github.com/spf13/viper"
)

type Config struct {
	Schedule []map[string]SchedulePeriods `mapstructure:"schedule"`
}

type SchedulePeriods struct {
	Image string `mapstructure:"image"`
	Time  string `mapstructure:"endTime"`
}

func main() {
	viper.SetConfigName("config")
	viper.AddConfigPath(".")
	if err := viper.ReadInConfig(); err != nil {
		fmt.Println(err)
		return
	}
	var config Config
	if err := viper.Unmarshal(&config); err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(config)
	currentSchedule := EvaluateSchedule((config))
	fmt.Println(currentSchedule)

}

func checkTime(schedule string) bool {
	theTime := time.Now().Format(time.Kitchen)
	now, _ := time.Parse(time.Kitchen, theTime)
	timeToBeEvaluated, err := time.Parse(time.Kitchen, schedule)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("It is currently %v\n", now.Format(time.Kitchen))
	fmt.Printf("is %v before %v?\n", now.Format(time.Kitchen), timeToBeEvaluated.Format(time.Kitchen))
	if now.Before(timeToBeEvaluated) {
		fmt.Printf("it is not past  %v\n", schedule)
		return true
	}

	return false
}

// returns the current schedule.
func EvaluateSchedule(config Config) string {

	for schedules := range config.Schedule {
		for k, v := range config.Schedule[schedules] {
			//fmt.Printf("key: %v Value: %v", k, v.Time)
			if checkTime(v.Time) {
				fmt.Printf("It is not yet past %v\n", v.Time)
				return k
			}
		}
	}
	return ""
}
