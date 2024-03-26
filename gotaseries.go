package gotaseries

import (
	"encoding/json"
	"net/http"
	"net/url"
	"time"
)

const (
	baseURL = "https://api.betaseries.com/api"
	version = "3.0"
)

type service struct {
	client *Client
}

// Client define Betaseries client to make request on their API
type Client struct {
	baseURL    url.URL
	userAgent  string
	apiKey     string
	httpClient *http.Client

	common service // avoid multiple allocations
	Shows  *showService
}

// NewClient returns a new Betaseries client
func NewClient(apiKey string) *Client {
	u, err := url.Parse(baseURL)
	if err != nil {
		return nil
	}

	c := &Client{
		baseURL: *u,
		apiKey:  apiKey,
		httpClient: &http.Client{
			Timeout: 60 * time.Second,
		},
	}

	c.common.client = c
	c.Shows = (*showService)(&c.common)

	return c
}

// newRequest creates a new http request
func (c *Client) newRequest(method, url string, body any) (*http.Request, error) {
	u, err := c.baseURL.Parse(url)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest(method, u.String(), nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("User-Agent", c.userAgent)
	req.Header.Set("X-Betaseries-Version", version)
	req.Header.Set("X-BetaSeries-Key", c.apiKey)

	return req, nil
}

// do send a request and decode it into v
func (c *Client) do(req *http.Request, v any) error {
	resp, err := c.httpClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	err = json.NewDecoder(resp.Body).Decode(v)

	return err
}
