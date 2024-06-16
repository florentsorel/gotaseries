package gotaseries

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"reflect"
	"strconv"
	"strings"
	"time"
)

const (
	baseURL = "https://api.betaseries.com/api"
	version = "3.0"
)

type Service struct {
	client *Client
}

// Client represents a Betaseries client.
type Client struct {
	baseURL    url.URL
	userAgent  string
	apiKey     string
	Token      string
	Locale     LocaleType
	httpClient *http.Client

	common Service
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
		Token:      "",
		Locale:     "",
		httpClient: httpClient,
	}

	c.common.client = c
	c.Shows = (*ShowService)(&c.common)

	return c
}

func (s *ShowService) doRequest(ctx context.Context, method, urlStr string, params any, response errorableResponse) error {
	req, err := s.client.newRequest(ctx, method, urlStr, params)
	if err != nil {
		return err
	}

	err = s.client.do(req, response)
	if err != nil {
		return err
	}

	if err = response.GetErrors().Err(); err != nil {
		return err
	}

	return nil
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

	if ctx == nil {
		return nil, fmt.Errorf("context cannot be nil")
	}
	req = req.WithContext(ctx)

	req.Header.Set("User-Agent", c.userAgent)
	req.Header.Set("X-BetaSeries-Version", version)
	req.Header.Set("X-BetaSeries-Key", c.apiKey)

	if c.Token != "" {
		req.Header.Set("X-BetaSeries-Token", c.Token)
	}

	return req, nil
}

func (c *Client) do(req *http.Request, v any) error {
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
		if v.Field(i).Kind() != reflect.Ptr {
			tag := v.Type().Field(i).Tag.Get("url")
			if tag != "" {
				paramMap[tag] = v.Field(i).Interface()
			}
			continue
		}

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
			case bool:
				q.Set(k, strconv.FormatBool(val))
			case int:
				q.Set(k, strconv.Itoa(val))
			case []int:
				if len(val) == 0 {
					continue
				}
				var strInts []string
				for _, num := range val {
					strInts = append(strInts, strconv.Itoa(num))
				}
				q.Set(k, strings.Join(strInts, ","))
			case string:
				q.Set(k, val)
			case time.Time:
				timestamp := strconv.Itoa(int(val.Unix()))
				q.Set(k, timestamp)
			case LocaleType:
				q.Set(k, val.String())
			case OrderType:
				q.Set(k, string(val))
			case OrderFavoriteType:
				q.Set(k, string(val))
			case OrderDateType:
				q.Set(k, string(val))
			case FormatType:
				q.Set(k, string(val))
			case RecommendationStatus:
				q.Set(k, string(val))
			case StatusFavoriteType:
				q.Set(k, string(val))
			}
		}
	}

	u.RawQuery = q.Encode()

	return u, nil
}
