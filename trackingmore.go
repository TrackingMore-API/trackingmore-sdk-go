package trackingmore

import (
	"crypto/tls"
	"errors"
	"net/http"
	"net/url"
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

	client := &http.Client{
		Timeout: 10 * time.Second,
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		},
	}
	proxyURL, _ := url.Parse("http://192.168.2.198:7890")
	client.Transport.(*http.Transport).Proxy = http.ProxyURL(proxyURL)
	return &Client{
		apiKey:     apiKey,
		httpClient: client,
	}, nil
}
