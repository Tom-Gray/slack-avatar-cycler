package main

import (
	"fmt"
	"github.com/spf13/viper"
	"time"
)

type Config struct {
	Schedule []map[string]SchedulePeriods `mapstructure:"schedule"`
}

type SchedulePeriods struct {
	Image string `mapstructure:"image"`
	Time  string `mapstructure:"time"`
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

}

func checkTime(schedule string) bool {
	now := time.Now()

	if now.Format(time.Kitchen) < schedule {
		fmt.Printf("it is past %v\n", schedule)
		return false
	}

	return true
}
