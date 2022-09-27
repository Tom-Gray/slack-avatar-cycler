package main

import (
	"net/http"
	slack "slack-avatar-cycler/client"
	"time"
)

var client = slack.SlackClient{
	BaseURL: "https://slack.com/api/users.setPhoto",
	HTTPClient: &http.Client{
		Timeout: time.Minute,
	},
	Token: "",
}

func main() {
	// Read Config

	// Create slack client
	client := slack.New(client)
	// evaluate time and select image
	// update image
	slack.SetProfileImage(client, "/Users/grayt5/dev/Homestar/open_mouth.png")
	// retain state somehow?

}
