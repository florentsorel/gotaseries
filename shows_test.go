package gotaseries

import (
	"context"
	"fmt"
	"os"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestShowService_Display(t *testing.T) {
	testCases := []struct {
		url      string
		params   map[string]string
		file     string
		expected Show
	}{
		{
			url:    "shows/display?id=1161",
			params: nil,
			file:   "data/shows/display_1161.json",
			expected: Show{
				ID:          1161,
				Title:       "Game of Thrones",
				Seasons:     8,
				Description: "Il y a très longtemps, à une époque oubliée, une force a détruit l'équilibre des saisons.",
			},
		},
		{
			url:    "shows/display?id=1161&locale=en",
			params: map[string]string{"locale": "en"},
			file:   "data/shows/display_1161_en.json",
			expected: Show{
				ID:          1161,
				Title:       "Game of Thrones",
				Seasons:     8,
				Description: "Seven noble families fight for control of the mythical land of Westeros",
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.url, func(t *testing.T) {
			data, err := os.ReadFile(tc.file)
			assert.NoError(t, err)

			ts, bc := setup(t, fmt.Sprintf("/%s", tc.url), string(data))
			defer ts.Close()

			show, err := bc.Shows.Display(1161, tc.params)
			assert.NoError(t, err)

			assert.Equal(t, tc.expected.ID, show.ID)
			assert.Equal(t, tc.expected.Title, show.Title)
			assert.Contains(t, show.Description, tc.expected.Description)
			assert.Equal(t, tc.expected.Seasons, show.Seasons)

			ctx, cancel := context.WithTimeout(context.Background(), 1*time.Microsecond)
			defer cancel()

			_, err = bc.Shows.DisplayWithCtx(ctx, 1161, tc.params)
			assert.Contains(t, err.Error(), "context deadline exceeded")
		})
	}
}
