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

type Client struct {
	baseURL    url.URL
	userAgent  string
	apiKey     string
	httpClient *http.Client

	common service
	Shows  showInterface
}

func NewClient(apiKey string) *Client {
	u, err := url.Parse(baseURL)
	if err != nil {
		return nil
	}

	httpClient := &http.Client{
		Timeout: time.Second * 30,
	}

	c := &Client{
		baseURL:    *u,
		userAgent:  "gotaseries/" + version,
		apiKey:     apiKey,
		httpClient: httpClient,
	}

	c.common.client = c
	c.Shows = (*showService)(&c.common)

	return c
}

func (c *Client) newRequest(method, url string) (*http.Request, error) {
	u, err := c.baseURL.Parse(url)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest(method, u.String(), nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("User-Agent", c.userAgent)
	req.Header.Set("X-BetaSeries-Version", version)
	req.Header.Set("X-BetaSeries-Key", c.apiKey)

	return req, nil
}

func (c *Client) do(req *http.Request, v interface{}) error {
	res, err := c.httpClient.Do(req)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	err = json.NewDecoder(res.Body).Decode(v)

	return err
}
