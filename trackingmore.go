package trackingmore

import (
	"errors"
	"net/http"
	"time"
)

type Client struct {
	apiKey     string
	httpClient *http.Client
}

// NewClient returns the TrackingMore client
func NewClient(apiKey string) (*Client, error) {
	if apiKey == "" {
		return nil, errors.New(ErrEmptyAPIKey)
	}
	return &Client{
		apiKey: apiKey,
		httpClient: &http.Client{
			Timeout: 10 * time.Second,
		},
	}, nil
}
