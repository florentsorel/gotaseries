package gotaseries

import "testing"

func TestShowService_Display(t *testing.T) {
	c := NewClient("123456789")

	show, err := c.Shows.Display(1161)
	if err != nil {
		t.Error(err)
	}

	if show.ID != 1161 {
		t.Errorf("Expected show id to be 1161, got %d", show.ID)
	}
}
