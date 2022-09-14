package client

import (
	"net/http"
)

type SlackClient struct {
	BaseURL    string
	HTTPClient *http.Client
	Token      string
}

func New(client SlackClient) *SlackClient {
	newClient := &SlackClient{
		BaseURL:    client.BaseURL,
		HTTPClient: client.HTTPClient,
		Token:      client.Token,
	}
	return newClient
}
