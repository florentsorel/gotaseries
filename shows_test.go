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
		params   ShowsDisplayParams
		file     string
		expected Show
	}{
		{
			url: "shows/display?id=1161",
			params: ShowsDisplayParams{
				ID: Int(1161),
			},
			file: "data/shows/display_id.json",
			expected: Show{
				ID:          1161,
				Title:       "Game of Thrones",
				Seasons:     8,
				Description: "Il y a très longtemps, à une époque oubliée, une force a détruit l'équilibre des saisons.",
			},
		},
		{
			url: "shows/display?id=1161&locale=en",
			params: ShowsDisplayParams{
				ID:     Int(1161),
				Locale: Locale(LocaleEN),
			},
			file: "data/shows/display_id_locale.json",
			expected: Show{
				ID:          1161,
				Title:       "Game of Thrones",
				Seasons:     8,
				Description: "Seven noble families fight for control of the mythical land of Westeros",
			},
		},
		{
			url: "shows/display?id=1161&thetvdb_id=81189",
			params: ShowsDisplayParams{
				ID:        Int(1161),
				TheTvdbID: Int(81189),
			},
			file: "data/shows/display_thetvdb_id.json",
			expected: Show{
				ID:          481,
				Title:       "Breaking Bad",
				Seasons:     5,
				Description: "La vie de Walter White, professeur de chimie dans un lycé",
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.url, func(t *testing.T) {
			data, err := os.ReadFile(tc.file)
			assert.NoError(t, err)

			ts, bc := setup(t, fmt.Sprintf("/%s", tc.url), string(data))
			defer ts.Close()

			show, err := bc.Shows.Display(context.Background(), tc.params)
			assert.NoError(t, err)

			assert.Equal(t, tc.expected.ID, show.ID)
			assert.Equal(t, tc.expected.Title, show.Title)
			assert.Contains(t, show.Description, tc.expected.Description)
			assert.Equal(t, tc.expected.Seasons, show.Seasons)

			ctx, cancel := context.WithTimeout(context.Background(), 1*time.Microsecond)
			defer cancel()

			_, err = bc.Shows.Display(ctx, tc.params)
			assert.Contains(t, err.Error(), "context deadline exceeded")
		})
	}
}

func TestShowService_List(t *testing.T) {
	testCases := []struct {
		url    string
		params ShowsListParams
		file   string
	}{
		{
			url: "shows/list?order=popularity",
			params: ShowsListParams{
				Order: Order(OrderPopularity),
			},
			file: "data/shows/list_order_popularity.json",
		},
		{
			url: "shows/list?locale=en&starting=game",
			params: ShowsListParams{
				Starting: String("game"),
				Locale:   Locale(LocaleEN),
			},
			file: "data/shows/list_starting_locale.json",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.url, func(t *testing.T) {
			data, err := os.ReadFile(tc.file)
			assert.NoError(t, err)

			ts, bc := setup(t, fmt.Sprintf("/%s", tc.url), string(data))
			defer ts.Close()

			show, err := bc.Shows.List(context.Background(), tc.params)
			assert.NoError(t, err)

			assert.Equal(t, 2, len(show))
			assert.NotEmpty(t, show)
		})
	}
}
