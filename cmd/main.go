package main

import (
	"fmt"
	"net/http"
	"os"
	slack "slack-avatar-cycler/client"
	"slack-avatar-cycler/scheduler"
	"time"
)

var client = slack.SlackClient{
	BaseURL: "https://slack.com/api/users.setPhoto",
	HTTPClient: &http.Client{
		Timeout: time.Minute,
	},
	Token: os.Getenv("SLACK_TOKEN"),
}

func main() {
	// Read Config

	client := slack.New(client)
	image := scheduler.GetImageForSchedulePeriod()
	err := slack.SetProfileImage(client, image)
	if err != nil {
		fmt.Println(err)
	}

}
