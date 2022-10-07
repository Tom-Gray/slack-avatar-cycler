package client

type slackError struct {
	Ok    bool   `json:"ok"`
	Error string `json:"error"`
}
