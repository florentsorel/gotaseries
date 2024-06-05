package gotaseries

import (
	"context"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func setup(t *testing.T, expectedURL string, expectedJSON string) (*httptest.Server, *Client) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, expectedURL, r.URL.String())
		_, _ = w.Write([]byte(expectedJSON))
	}))

	mockURL, err := url.Parse(ts.URL)
	assert.NoError(t, err)

	c := &Client{
		baseURL:    *mockURL,
		userAgent:  "test",
		apiKey:     "api_key",
		httpClient: &http.Client{Timeout: time.Second * 30},
	}

	c.common.client = c
	c.Shows = (*showService)(&c.common)

	return ts, c
}

func TestNewClient(t *testing.T) {
	client := NewClient("api_key")

	expectedURL, err := url.Parse("http://api.test.com")
	assert.Nil(t, err)

	client.baseURL = *expectedURL

	assert.NotNil(t, client)
	assert.Equal(t, "api_key", client.apiKey)
	assert.Equal(t, "http://api.test.com", client.baseURL.String())
}

func TestNewRequest(t *testing.T) {
	client := NewClient("api_key")
	apiURL, err := url.Parse("http://api.test.com")
	assert.NoError(t, err)

	client.baseURL = *apiURL
	client.userAgent = "gotaseries-user-agent"
	client.apiKey = "api_key"

	req, err := client.newRequest(context.Background(), "GET", "/test", ShowsListParams{})
	assert.NoError(t, err)

	assert.Equal(t, "http://api.test.com/test", req.URL.String())
	assert.Equal(t, "gotaseries-user-agent", req.Header.Get("User-Agent"))
	assert.Equal(t, "api_key", req.Header.Get("X-BetaSeries-Key"))
	assert.Equal(t, "3.0", req.Header.Get("X-BetaSeries-Version"))
}

func TestNewRequestWithParams(t *testing.T) {
	client := NewClient("api_key")
	apiURL, err := url.Parse("http://api.test.com")
	assert.NoError(t, err)

	client.baseURL = *apiURL
	client.userAgent = "gotaseries-user-agent"
	client.apiKey = "api_key"

	req, err := client.newRequest(context.Background(), "GET", "/test", ShowsDisplayParams{ID: Int(1161)})
	assert.NoError(t, err)

	assert.Equal(t, "http://api.test.com/test?id=1161", req.URL.String())
	assert.Equal(t, "gotaseries-user-agent", req.Header.Get("User-Agent"))
	assert.Equal(t, "api_key", req.Header.Get("X-BetaSeries-Key"))
	assert.Equal(t, "3.0", req.Header.Get("X-BetaSeries-Version"))
}
