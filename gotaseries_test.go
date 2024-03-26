package gotaseries

import "testing"

func TestNewClient(t *testing.T) {
	c := NewClient("123456789")

	if c == nil {
		t.Error("Expected a client, got nil")
	}

	if c.apiKey != "123456789" {
		t.Errorf("Expected api key to be 123456789, got %s", c.apiKey)
	}

	if c.common.client != c {
		t.Error("Expected common client to be the client")
	}
}
