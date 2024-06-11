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

func TestShowService_Random(t *testing.T) {
	data, err := os.ReadFile("data/shows/random_limit.json")
	assert.NoError(t, err)

	ts, bc := setup(t, fmt.Sprintf("/%s", "shows/random?nb=2"), string(data))
	defer ts.Close()

	shows, err := bc.Shows.Random(context.Background(), ShowsRandomParams{
		Number: Int(2),
	})
	assert.NoError(t, err)

	expected := []Show{
		{
			ID:          2152,
			TheTvdbID:   164521,
			ImdbID:      "tt1592254",
			MoviedbID:   32736,
			Title:       "The Defenders",
			Seasons:     1,
			Episodes:    18,
			Followers:   706,
			Creation:    2010,
			Poster:      "https://pictures.betaseries.com/fonds/poster/880f2fd6315a539f8ab251bf92147af9.jpg",
			Platforms:   nil,
			ResourceURL: "https://www.betaseries.com/serie/the-defenders-2010",
		},
		{
			ID:          32716,
			TheTvdbID:   420382,
			ImdbID:      "",
			MoviedbID:   113193,
			Title:       "Heroes (2022)",
			Seasons:     1,
			Episodes:    38,
			Followers:   43,
			Creation:    2022,
			Poster:      "https://pictures.betaseries.com/fonds/poster/8a34e1a3b9fd6b39e5ee693c4fec1caa.jpg",
			Platforms:   nil,
			ResourceURL: "https://www.betaseries.com/serie/heroes-2022",
		},
	}

	assert.Equal(t, expected, shows)
}

func TestShowService_Episodes(t *testing.T) {
	data, err := os.ReadFile("data/shows/episodes.json")
	assert.NoError(t, err)

	ts, bc := setup(t, fmt.Sprintf("/%s", "shows/episodes?episode=1&id=1161&season=1&subtitles=true"), string(data))
	defer ts.Close()

	episodes, err := bc.Shows.Episodes(context.Background(), ShowsEpisodesParams{
		ID:        Int(1161),
		Season:    Int(1),
		Episode:   Int(1),
		Subtitles: Bool(true),
	})
	assert.NoError(t, err)

	d, err := time.Parse("2006-01-02", "2011-04-17")
	assert.NoError(t, err)

	subtitleDT, err := time.Parse("2006-01-02 15:04:05", "2023-01-22 15:29:05")
	assert.NoError(t, err)

	assert.Equal(t, 1, len(episodes))

	assert.Equal(t, 281009, episodes[0].ID)
	assert.Equal(t, "Winter Is Coming", episodes[0].Title)
	assert.Equal(t, 1, episodes[0].Season)
	assert.Equal(t, 1, episodes[0].Episode)
	assert.Equal(t, "S01E01", episodes[0].Code)
	assert.Equal(t, Date(d), episodes[0].Date)

	assert.Equal(t, "", episodes[0].Show.Description)

	assert.Equal(t, 975356, episodes[0].Subtitles[0].ID)
	assert.Equal(t, "VO", episodes[0].Subtitles[0].Language)
	assert.Equal(t, "opensubtitles", episodes[0].Subtitles[0].Source)
	assert.Equal(t, "game.of.thrones.s01.e01.winter.is.coming.(2011).eng.1cd.(9403451).zip", episodes[0].Subtitles[0].File)
	assert.Equal(t, DateTime(subtitleDT), episodes[0].Subtitles[0].Date)
	assert.Equal(t, "https://www.betaseries.com/srt/975356", episodes[0].Subtitles[0].URL)
	assert.Equal(t, 1, episodes[0].Subtitles[0].Quality)

	assert.Equal(t, 8499, episodes[0].Note.Total)
	assert.Equal(t, 4.4326, episodes[0].Note.Mean)

	assert.Equal(t, 31, episodes[0].PlaformLinks[0].PlatformID)
	assert.Equal(t, "Amazon Video", episodes[0].PlaformLinks[0].Platform)
	assert.Equal(t, "#3B8DD0", episodes[0].PlaformLinks[0].Color)
	assert.Equal(t, "vod", episodes[0].PlaformLinks[0].Type)
	assert.Equal(t, "vod", episodes[0].PlaformLinks[0].Type)
	assert.Equal(t, "vod", episodes[0].PlaformLinks[0].Type)
	assert.Equal(t, "https://www.betaseries.com/link/2327460", episodes[0].PlaformLinks[0].Link)
}
