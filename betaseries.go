package gotaseries

import (
	"context"
	"encoding/json"
	"net/http"
	"net/url"
	"reflect"
	"strconv"
	"time"
)

const (
	baseURL = "https://api.betaseries.com/api"
	version = "3.0"
)

type service struct {
	client *Client
}

// Client represents a Betaseries client.
type Client struct {
	baseURL    url.URL
	userAgent  string
	apiKey     string
	Locale     locale
	httpClient *http.Client

	common service
	Shows  *ShowService
}

// NewClient returns a new Betaseries client. You need to provide an API key.
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
		Locale:     "",
		httpClient: httpClient,
	}

	c.common.client = c
	c.Shows = (*ShowService)(&c.common)

	return c
}

func (c *Client) newRequest(ctx context.Context, method, url string, params any) (*http.Request, error) {
	u, err := c.buildURL(url, params)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest(method, u.String(), nil)
	if err != nil {
		return nil, err
	}

	req = req.WithContext(ctx)

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

func (c *Client) buildURL(urlStr string, params any) (*url.URL, error) {
	u, err := c.baseURL.Parse(urlStr)
	if err != nil {
		return nil, err
	}

	v := reflect.ValueOf(params)
	paramMap := make(map[string]any, v.NumField())
	for i := 0; i < v.NumField(); i++ {
		if v.Field(i).IsNil() {
			continue
		}
		tag := v.Type().Field(i).Tag.Get("url")
		if tag != "" {
			paramMap[tag] = v.Field(i).Elem().Interface()
		}
	}

	q := u.Query()
	if c.Locale != "" {
		q.Add("locale", c.Locale.String())
	}

	if len(paramMap) > 0 {
		for k, value := range paramMap {
			switch val := value.(type) {
			case int:
				q.Set(k, strconv.Itoa(val))
			case string:
				q.Set(k, val)
			case locale:
				q.Set(k, val.String())
			}
		}
	}

	u.RawQuery = q.Encode()

	return u, nil
}
