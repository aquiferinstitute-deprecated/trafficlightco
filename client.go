package trafficlightco

import (
	"fmt"
	"net/http"
	"net/url"
	"time"
)

// apiTimeout is the time out of API requests in second
const apiTimeout = 10

// Client contains the setup for the API interactions
type Client struct {
	httpClient *http.Client
	apiURL     *url.URL
	apiKey     string
}

// NewClient returns a new client
func NewClient(u, key string) (*Client, error) {
	uu, err := url.Parse(u)
	if err != nil {
		return nil, fmt.Errorf("NewClient - %s: %v", u, err)
	}
	if key == "" {
		return nil, fmt.Errorf("NewClient: empty api key")
	}
	return &Client{
		apiURL: uu,
		apiKey: key,
		httpClient: &http.Client{
			Timeout: apiTimeout * time.Second,
		},
	}, nil
}
