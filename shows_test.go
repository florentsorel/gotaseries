package gotaseries

import (
	"context"
	"fmt"
	"os"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestShowService_Note(t *testing.T) {
	data, err := os.ReadFile("data/shows/note_post.json")
	assert.NoError(t, err)

	ts, bc := setup(t, "POST", fmt.Sprintf("/%s", "shows/note?id=1161&note=4"), string(data))
	defer ts.Close()

	show, err := bc.Shows.AddNote(context.Background(), ShowsAddNoteParams{
		ID:   Int(1161),
		Note: 4,
	})
	assert.NoError(t, err)

	assert.Equal(t, 4, *show.Note.User)
}

func TestShowService_DeleteNote(t *testing.T) {
	data, err := os.ReadFile("data/shows/note_delete.json")
	assert.NoError(t, err)

	ts, bc := setup(t, "DELETE", fmt.Sprintf("/%s", "shows/note?id=1161"), string(data))
	defer ts.Close()

	show, err := bc.Shows.DeleteNote(context.Background(), ShowsDeleteNoteParams{
		ID: Int(1161),
	})
	assert.NoError(t, err)

	assert.Equal(t, 0, *show.Note.User)
}

func TestShowService_Search(t *testing.T) {
	data, err := os.ReadFile("data/shows/search.json")
	assert.NoError(t, err)

	ts, bc := setup(t, "GET", fmt.Sprintf("/%s", "shows/search?platforms=1%2C221&title=game"), string(data))
	defer ts.Close()

	shows, err := bc.Shows.Search(context.Background(), ShowsSearchParams{
		Title:     String("game"),
		Platforms: []int{1, 221},
	})
	assert.NoError(t, err)

	assert.Equal(t, 2, len(shows))

	assert.Equal(t, "GameStop : Les geeks défient Wall Street", shows[0].Title)
	assert.Equal(t, 1, shows[0].Platforms.Svod.ID)

	assert.Equal(t, "Gamers!", shows[1].Title)
	assert.Equal(t, 221, shows[1].Platforms.Svod.ID)
}

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

			ts, bc := setup(t, "GET", fmt.Sprintf("/%s", tc.url), string(data))
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

func TestShowService_DisplayWithToken(t *testing.T) {
	data, err := os.ReadFile("data/shows/display_with_token.json")
	assert.NoError(t, err)

	ts, bc := setup(t, "GET", fmt.Sprintf("/%s", "shows/display?id=807"), string(data))
	defer ts.Close()

	show, err := bc.Shows.Display(context.Background(), ShowsDisplayParams{
		ID: Int(807),
	})
	assert.NoError(t, err)

	_, err = time.Parse("2006-01-02", show.User.Next.Date.String())
	assert.NoError(t, err)

	assert.Equal(t, "Modern Family", show.Title)
	assert.Equal(t, true, show.InAccount)
	assert.Equal(t, false, show.User.Archived)
	assert.Equal(t, 87.599999999999994315658113919198513031005859375, show.User.Status)
	assert.Equal(t, "S10E09", show.User.Last)
	assert.Equal(t, Int(1216277), show.User.Next.ID)
	assert.Equal(t, "S10E10", show.User.Next.Code)
	assert.Equal(t, "2018-12-12", show.User.Next.Date.String())
	assert.Equal(t, "test user", show.User.FriendsWatching[0].Login)
}

func TestShowService_DisplayNotFound(t *testing.T) {
	data, err := os.ReadFile("data/shows/no_series_found.json")
	assert.NoError(t, err)

	ts, bc := setup(t, "GET", fmt.Sprintf("/%s", "shows/display"), string(data))
	defer ts.Close()

	_, err = bc.Shows.Display(context.Background(), ShowsDisplayParams{})
	assert.Error(t, err)

	assert.Equal(t, err.Error(), "Code: 4001, Message: No series found.\n")
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

			ts, bc := setup(t, "GET", fmt.Sprintf("/%s", tc.url), string(data))
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

	ts, bc := setup(t, "GET", fmt.Sprintf("/%s", "shows/random?nb=2"), string(data))
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

	ts, bc := setup(t, "GET", fmt.Sprintf("/%s", "shows/episodes?episode=1&id=1161&season=1&subtitles=true"), string(data))
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

func TestShowService_EpisodesNotFound(t *testing.T) {
	data, err := os.ReadFile("data/shows/episodes_not_found.json")
	assert.NoError(t, err)

	ts, bc := setup(t, "GET", fmt.Sprintf("/%s", "shows/episodes"), string(data))
	defer ts.Close()

	_, err = bc.Shows.Episodes(context.Background(), ShowsEpisodesParams{})
	assert.Error(t, err)

	assert.Equal(t, err.Error(), "Code: 0, Message: You must send an \"id\" or a \"thetvdb_id\" parameter to this API request.\n")
}

func TestShowService_Add(t *testing.T) {
	data, err := os.ReadFile("data/shows/add_post.json")
	assert.NoError(t, err)

	ts, bc := setup(t, "POST", fmt.Sprintf("/%s", "shows/show?episode_id=24341&id=2"), string(data))
	defer ts.Close()

	show, err := bc.Shows.Add(context.Background(), ShowsAddParams{
		ID:        Int(2),
		EpisodeID: Int(24341),
	})
	assert.NoError(t, err)

	assert.Equal(t, "S01E03", show.User.Last)
	assert.Equal(t, Int(24342), show.User.Next.ID)
	assert.Equal(t, "S01E04", show.User.Next.Code)
}

func TestShowService_Add_AlreadyInAccount(t *testing.T) {
	data, err := os.ReadFile("data/shows/add_post_already_in_account.json")
	assert.NoError(t, err)

	ts, bc := setup(t, "POST", fmt.Sprintf("/%s", "shows/show?episode_id=24341&id=2"), string(data))
	defer ts.Close()

	_, err = bc.Shows.Add(context.Background(), ShowsAddParams{
		ID:        Int(2),
		EpisodeID: Int(24341),
	})
	assert.Error(t, err)

	assert.Equal(t, err.Error(), "Code: 2003, Message: L'utilisateur a déjà cette série dans son compte.\n")
}

func TestShowService_Delete(t *testing.T) {
	data, err := os.ReadFile("data/shows/delete_delete.json")
	assert.NoError(t, err)

	ts, bc := setup(t, "DELETE", fmt.Sprintf("/%s", "shows/show?id=2"), string(data))
	defer ts.Close()

	show, err := bc.Shows.Delete(context.Background(), ShowsDeleteParams{
		ID: Int(2),
	})
	assert.NoError(t, err)

	assert.Equal(t, "S01E03", show.User.Last)
	assert.Equal(t, Int(24342), show.User.Next.ID)
	assert.Equal(t, "S01E04", show.User.Next.Code)
}

func TestShowService_Archive(t *testing.T) {
	data := `{"show": {"user":{"archived":true}}}`

	ts, bc := setup(t, "POST", fmt.Sprintf("/%s", "shows/archive?id=2"), data)
	defer ts.Close()

	show, err := bc.Shows.Archive(context.Background(), ShowsArchiveParams{
		ID: Int(2),
	})
	assert.NoError(t, err)

	assert.Equal(t, true, show.User.Archived)
}

func TestShowService_Unarchive(t *testing.T) {
	data := `{"show": {"user":{"archived":false}}}`
	ts, bc := setup(t, "DELETE", fmt.Sprintf("/%s", "shows/archive?id=2"), data)
	defer ts.Close()

	show, err := bc.Shows.Unarchive(context.Background(), ShowsUnarchiveParams{
		ID: Int(2),
	})
	assert.NoError(t, err)

	assert.Equal(t, false, show.User.Archived)
}

func TestShowService_Recommendation(t *testing.T) {
	testCases := []struct {
		title         string
		url           string
		params        ShowsCreateRecommendationParams
		file          string
		expected      any
		expectedError bool
	}{
		{
			title: "OK",
			url:   "shows/recommendation?id=1456&to=1234",
			params: ShowsCreateRecommendationParams{
				ID: Int(1456),
				To: 1234,
			},
			file: "data/shows/recommendation_post.json",
			expected: &Recommendation{
				ID:      106602,
				From:    5678,
				To:      1234,
				ShowId:  1456,
				Status:  RecommendationStatusWait,
				Comment: String("Test"),
			},
			expectedError: false,
		},
		{
			title: "Already recommended",
			url:   "shows/recommendation?id=1456&to=1234",
			params: ShowsCreateRecommendationParams{
				ID: Int(1456),
				To: 1234,
			},
			file:          "data/shows/recommendation_post_already_recommended.json",
			expected:      "Code: 0, Message: L'utilisateur a déjà recommandé cette série à ce membre.\n",
			expectedError: true,
		},
		{
			title: "In account",
			url:   "shows/recommendation?id=1456&to=1234",
			params: ShowsCreateRecommendationParams{
				ID: Int(1456),
				To: 1234,
			},
			file:          "data/shows/recommendation_post_in_account.json",
			expected:      "Code: 2003, Message: L'utilisateur a déjà cette série dans son compte.\n",
			expectedError: true,
		},
		{
			title: "Not friends",
			url:   "shows/recommendation?id=1456&to=1234",
			params: ShowsCreateRecommendationParams{
				ID: Int(1456),
				To: 1234,
			},
			file:          "data/shows/recommendation_post_not_friends.json",
			expected:      "Code: 0, Message: Les membres ne sont pas amis entre eux.\n",
			expectedError: true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.title, func(t *testing.T) {
			data, err := os.ReadFile(tc.file)
			assert.NoError(t, err)

			ts, bc := setup(t, "POST", fmt.Sprintf("/%s", tc.url), string(data))
			defer ts.Close()

			recommendation, err := bc.Shows.CreateRecommendation(context.Background(), tc.params)
			if tc.expectedError {
				assert.Error(t, err)
				assert.Equal(t, err.Error(), tc.expected)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tc.expected, recommendation)
			}
		})
	}
}

func TestShowService_UpdateRecommendation(t *testing.T) {
	testCases := []struct {
		title         string
		url           string
		params        ShowsUpdateRecommendationParams
		file          string
		expected      any
		expectedError bool
	}{
		{
			title: "Accepted",
			url:   "shows/recommendation?id=106618&status=accept",
			params: ShowsUpdateRecommendationParams{
				ID:     106618,
				Status: RecommendationStatusAccept,
			},
			file: "data/shows/recommendation_put_accept.json",
			expected: &Recommendation{
				ID:      106618,
				From:    1234,
				To:      5678,
				ShowId:  22028,
				Status:  RecommendationStatusAccept,
				Comment: nil,
			},
			expectedError: false,
		},
		{
			title: "Declined",
			url:   "shows/recommendation?id=106618&status=decline",
			params: ShowsUpdateRecommendationParams{
				ID:     106618,
				Status: RecommendationStatusDecline,
			},
			file: "data/shows/recommendation_put_decline.json",
			expected: &Recommendation{
				ID:      106617,
				From:    1234,
				To:      5678,
				ShowId:  22028,
				Status:  "decline",
				Comment: nil,
			},
			expectedError: false,
		},
		{
			title: "Already in account",
			url:   "shows/recommendation?id=106618&status=decline",
			params: ShowsUpdateRecommendationParams{
				ID:     106618,
				Status: "decline",
			},
			file:          "data/shows/recommendation_put_already_in_account.json",
			expected:      "Code: 2003, Message: L'utilisateur a déjà cette série dans son compte.\n",
			expectedError: true,
		},
		{
			title: "Not exists",
			url:   "shows/recommendation?id=106619&status=decline",
			params: ShowsUpdateRecommendationParams{
				ID:     106619,
				Status: RecommendationStatusDecline,
			},
			file:          "data/shows/recommendation_put_not_exists.json",
			expected:      "Code: 4005, Message: La recommandation de série avec l'ID 106619 n'existe pas.\n",
			expectedError: true,
		},
		{
			title: "Not intended for user",
			url:   "shows/recommendation?id=1&status=decline",
			params: ShowsUpdateRecommendationParams{
				ID:     1,
				Status: RecommendationStatusDecline,
			},
			file:          "data/shows/recommendation_put_not_intended_for_user.json",
			expected:      "Code: 0, Message: The recommendation is not intended for this user.\n",
			expectedError: true,
		},
		{
			title: "Invalid status",
			url:   "shows/recommendation?id=1&status=test",
			params: ShowsUpdateRecommendationParams{
				ID:     1,
				Status: "test",
			},
			file:          "data/shows/recommendation_put_invalid_status.json",
			expected:      "Code: 0, Message: Wrong value for status variable.\n",
			expectedError: true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.title, func(t *testing.T) {
			data, err := os.ReadFile(tc.file)
			assert.NoError(t, err)

			ts, bc := setup(t, "PUT", fmt.Sprintf("/%s", tc.url), string(data))
			defer ts.Close()

			recommendation, err := bc.Shows.UpdateRecommendation(context.Background(), tc.params)
			if tc.expectedError {
				assert.Error(t, err)
				assert.Equal(t, err.Error(), tc.expected)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tc.expected, recommendation)
			}
		})
	}
}

func TestShowService_DeleteRecommendation(t *testing.T) {
	testCases := []struct {
		title         string
		url           string
		params        ShowsDeleteRecommendationParams
		file          string
		expected      any
		expectedError bool
	}{
		{
			title: "Accepted",
			url:   "shows/recommendation?id=106602",
			params: ShowsDeleteRecommendationParams{
				ID: 106602,
			},
			file: "data/shows/recommendation_delete.json",
			expected: &Recommendation{
				ID:      106602,
				From:    5678,
				To:      1234,
				ShowId:  1456,
				Status:  "wait",
				Comment: nil,
			},
			expectedError: false,
		},
		{
			title: "Declined",
			url:   "shows/recommendation?id=106619",
			params: ShowsDeleteRecommendationParams{
				ID: 106619,
			},
			file:          "data/shows/recommendation_delete_not_exists.json",
			expected:      "Code: 4005, Message: La recommandation de série avec l'ID 106619 n'existe pas.\n",
			expectedError: true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.title, func(t *testing.T) {
			data, err := os.ReadFile(tc.file)
			assert.NoError(t, err)

			ts, bc := setup(t, "DELETE", fmt.Sprintf("/%s", tc.url), string(data))
			defer ts.Close()

			recommendation, err := bc.Shows.DeleteRecommendation(context.Background(), tc.params)
			if tc.expectedError {
				assert.Error(t, err)
				assert.Equal(t, err.Error(), tc.expected)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tc.expected, recommendation)
			}
		})
	}
}

func TestShowService_Recommendations(t *testing.T) {
	data := `{
  "recommendations": [
    {
      "id": 106623,
      "from_id": 1234,
      "to_id": 5678,
      "show_id": 1456,
      "status": "wait",
      "comments": "I recommend this show to you"
    },
    {
      "id": 106630,
      "from_id": 5678,
      "to_id": 1234,
      "show_id": 22028,
      "status": "wait",
      "comments": null
    },
    {
      "id": 106631,
      "from_id": 5678,
      "to_id": 1234,
      "show_id": 23439,
      "status": "wait",
      "comments": null
    }
  ],
  "errors": []
}`

	ts, bc := setup(t, "GET", fmt.Sprintf("/%s", "shows/recommendations"), data)
	defer ts.Close()

	recommendations, err := bc.Shows.Recommendations(context.Background(), ShowsRecommendationsParams{})
	assert.NoError(t, err)

	assert.Equal(t, 3, len(recommendations))
}

func TestShowService_Similars(t *testing.T) {
	data, err := os.ReadFile("data/shows/similars.json")
	assert.NoError(t, err)

	ts, bc := setup(t, "GET", fmt.Sprintf("/%s", "shows/similars?thetvdb_id=121361"), string(data))
	defer ts.Close()

	similars, err := bc.Shows.Similars(context.Background(), ShowsSimilarsParams{
		TheTvdbID: Int(121361),
	})
	assert.NoError(t, err)

	witcher := similars[2]

	assert.Equal(t, 45, len(similars))

	assert.Equal(t, 21591, witcher.ID)
	assert.Equal(t, "The Witcher", witcher.Title)
	assert.Equal(t, 20999, witcher.ShowID)
	assert.Equal(t, 362696, witcher.TheTvdbID)
	assert.Nil(t, witcher.Notes)
	assert.Nil(t, witcher.Show)
}

func TestShowService_SimilarsWithShow(t *testing.T) {
	data, err := os.ReadFile("data/shows/similars_with_show.json")
	assert.NoError(t, err)

	ts, bc := setup(t, "GET", fmt.Sprintf("/%s", "shows/similars?details=true&id=1161"), string(data))
	defer ts.Close()

	similars, err := bc.Shows.Similars(context.Background(), ShowsSimilarsParams{
		ID:      Int(1161),
		Details: Bool(true),
	})
	assert.NoError(t, err)

	witcher := similars[2]

	assert.Equal(t, 45, len(similars))

	assert.Equal(t, 21591, witcher.ID)
	assert.Equal(t, "The Witcher", witcher.Title)
	assert.Equal(t, 20999, witcher.ShowID)
	assert.Equal(t, 362696, witcher.TheTvdbID)
	assert.Nil(t, witcher.Notes)
	assert.NotNil(t, witcher.Show)
	assert.Equal(t, 3, len(witcher.Show.SeasonsDetails))
	assert.Equal(t, 33, witcher.Show.Episodes)
}

func TestShowService_Videos(t *testing.T) {
	data, err := os.ReadFile("data/shows/videos.json")
	assert.NoError(t, err)

	ts, bc := setup(t, "GET", fmt.Sprintf("/%s", "shows/videos?id=1161"), string(data))
	defer ts.Close()

	videos, err := bc.Shows.Videos(context.Background(), ShowsVideosParams{
		ID: Int(1161),
	})
	assert.NoError(t, err)

	assert.Equal(t, 41, len(videos))

	video := videos[0]

	datetime, err := time.Parse("2006-01-02 15:04:05", "2011-04-18 22:20:02")
	assert.NoError(t, err)

	assert.Equal(t, 1161, video.ShowID)
	assert.Equal(t, "dailymotion", video.Host)
	assert.Equal(t, "x87hfr9", video.Slug)
	assert.Equal(t, "https://www.dailymotion.com/video/x87hfr9", video.URL)
	assert.Equal(t, DateTime(datetime), video.Date)
	assert.Equal(t, 0, video.Season)
	assert.Equal(t, 0, video.Episode)
}

func TestShowService_VideosNotFound(t *testing.T) {
	data, err := os.ReadFile("data/shows/no_series_found.json")
	assert.NoError(t, err)

	ts, bc := setup(t, "GET", fmt.Sprintf("/%s", "shows/videos"), string(data))
	defer ts.Close()

	_, err = bc.Shows.Videos(context.Background(), ShowsVideosParams{})
	assert.Error(t, err)

	assert.Equal(t, err.Error(), "Code: 4001, Message: No series found.\n")
}

func TestShowService_Characters(t *testing.T) {
	data, err := os.ReadFile("data/shows/characters.json")
	assert.NoError(t, err)

	ts, bc := setup(t, "GET", fmt.Sprintf("/%s", "shows/characters?id=1161"), string(data))
	defer ts.Close()

	characters, err := bc.Shows.Characters(context.Background(), ShowsCharactersParams{
		ID: Int(1161),
	})
	assert.NoError(t, err)

	assert.Equal(t, 545, len(characters))

	assert.Equal(t, 1161, characters[3].ShowID)
	assert.Equal(t, 14605, characters[3].PersonID)
	assert.Equal(t, "Daenerys Targaryen", characters[3].Name)
	assert.Equal(t, "Emilia Clarke", characters[3].Actor)
	assert.Equal(t, "https://pictures.betaseries.com/persons/wb8VfDPGpyqcFltnRcJR1Wj3h4Z.jpg", characters[3].Picture)
}

func TestShowService_CharactersNotFound(t *testing.T) {
	data, err := os.ReadFile("data/shows/no_series_found.json")
	assert.NoError(t, err)

	ts, bc := setup(t, "GET", fmt.Sprintf("/%s", "shows/characters"), string(data))
	defer ts.Close()

	_, err = bc.Shows.Characters(context.Background(), ShowsCharactersParams{})
	assert.Error(t, err)

	assert.Equal(t, err.Error(), "Code: 4001, Message: No series found.\n")
}

func TestShowService_Pictures(t *testing.T) {
	data, err := os.ReadFile("data/shows/pictures.json")
	assert.NoError(t, err)

	ts, bc := setup(t, "GET", fmt.Sprintf("/%s", "shows/pictures?format=hd&id=1161&order=-date"), string(data))
	defer ts.Close()

	pictures, err := bc.Shows.Pictures(context.Background(), ShowsPicturesParams{
		ID:     Int(1161),
		Order:  OrderDate(OrderDateDESC),
		Format: Format(FormatHD),
	})
	assert.NoError(t, err)

	assert.Equal(t, 222, len(pictures))

	picture := pictures[0]

	datetime, err := time.Parse("2006-01-02 15:04:05", "2023-06-15 14:30:48")
	assert.NoError(t, err)

	assert.Equal(t, 149516, picture.ID)
	assert.Equal(t, 1161, picture.ShowID)
	assert.Equal(t, "https://pictures.betaseries.com/fonds/original/1161_63172304.jpg", picture.URL)
	assert.Equal(t, 1920, picture.Width)
	assert.Equal(t, 1080, picture.Height)
	assert.Equal(t, DateTime(datetime), picture.Date)
	assert.Equal(t, "none", picture.Picked)
}

func TestShowService_PicturesNotFound(t *testing.T) {
	data, err := os.ReadFile("data/shows/no_series_found.json")
	assert.NoError(t, err)

	ts, bc := setup(t, "GET", fmt.Sprintf("/%s", "shows/pictures"), string(data))
	defer ts.Close()

	_, err = bc.Shows.Pictures(context.Background(), ShowsPicturesParams{})
	assert.Error(t, err)

	assert.Equal(t, err.Error(), "Code: 4001, Message: No series found.\n")
}

func TestShowService_Favorites(t *testing.T) {
	data, err := os.ReadFile("data/shows/favorites.json")
	assert.NoError(t, err)

	ts, bc := setup(t, "GET", fmt.Sprintf("/%s", "shows/favorites?id=1&limit=2&order=alphabetical&status=archived"), string(data))
	defer ts.Close()

	shows, err := bc.Shows.Favorites(context.Background(), ShowsFavoritesParams{
		ID:     Int(1),
		Order:  OrderFavorite(OrderFavoriteAlphabetical),
		Limit:  Int(2),
		Status: StatusFavorite(StatusFavoritesArchived),
	})
	assert.NoError(t, err)

	assert.Equal(t, 2, len(shows.Shows))
	assert.Equal(t, 5, shows.Total)
}

func TestShowService_AddFavorite(t *testing.T) {
	data := `{"show": {"user":{"favorited":true}}}`

	ts, bc := setup(t, "POST", fmt.Sprintf("/%s", "shows/favorite?id=2"), data)
	defer ts.Close()

	show, err := bc.Shows.AddFavorite(context.Background(), ShowsAddFavoriteParams{
		ID: 2,
	})
	assert.NoError(t, err)

	assert.Equal(t, true, show.User.Favorited)
}

func TestShowService_DeleteFavorite(t *testing.T) {
	data := `{"show": {"user":{"favorited":false}}}`

	ts, bc := setup(t, "DELETE", fmt.Sprintf("/%s", "shows/favorite?id=2"), data)
	defer ts.Close()

	show, err := bc.Shows.DeleteFavorite(context.Background(), ShowsDeleteFavoriteParams{
		ID: 2,
	})
	assert.NoError(t, err)

	assert.Equal(t, false, show.User.Favorited)
}

func TestShowService_AddTags(t *testing.T) {
	data := `{"show": {"user":{"tags":"tag1, tag2"}}}`

	ts, bc := setup(t, "POST", fmt.Sprintf("/%s", "shows/tags?id=2&tags=tag1%2Ctag2"), data)
	defer ts.Close()

	show, err := bc.Shows.UpdateTags(context.Background(), ShowsUpdateTagsParams{
		ID:   2,
		Tags: []string{"tag1", "tag2"},
	})
	assert.NoError(t, err)

	assert.Equal(t, Tags{"tag1", "tag2"}, show.User.Tags)
}

func TestShowService_Member(t *testing.T) {
	data, err := os.ReadFile("data/shows/member.json")
	assert.NoError(t, err)

	ts, bc := setup(t, "GET", fmt.Sprintf("/%s", "shows/member?id=1&limit=2&order=progression&status=current"), string(data))
	defer ts.Close()

	shows, err := bc.Shows.Member(context.Background(), ShowsMemberParams{
		ID:     Int(1),
		Order:  OrderShowMember(OrderShowMemberProgression),
		Limit:  Int(2),
		Status: StatusShowMember(StatusShowMemberCurrent),
	})
	assert.NoError(t, err)

	assert.Equal(t, 2, len(shows.Shows))
	assert.Equal(t, 13, shows.Total)
	assert.Equal(t, 150, shows.TotalMissingShows)
}

func TestShowService_Discover(t *testing.T) {
	data, err := os.ReadFile("data/shows/discover.json")
	assert.NoError(t, err)

	ts, bc := setup(t, "GET", fmt.Sprintf("/%s", "shows/discover?limit=2"), string(data))
	defer ts.Close()

	shows, err := bc.Shows.Discover(context.Background(), ShowsDiscoverParams{
		Limit: Int(2),
	})
	assert.NoError(t, err)

	assert.Equal(t, 2, len(shows))
}

func TestShowService_DiscoverPlatform(t *testing.T) {
	ts, bc := setup(t, "GET", fmt.Sprintf("/%s", "shows/discover_platform"), `{"shows": [{"title": "Breaking Bad"}]}`)
	defer ts.Close()

	shows, err := bc.Shows.DiscoverPlatform(context.Background(), ShowsDiscoverPlatformsParams{})
	assert.NoError(t, err)

	assert.Equal(t, 1, len(shows))
}

func TestShowService_Genres(t *testing.T) {
	data, err := os.ReadFile("data/shows/genres.json")
	assert.NoError(t, err)

	ts, bc := setup(t, "GET", fmt.Sprintf("/%s", "shows/genres"), string(data))
	defer ts.Close()

	genres, err := bc.Shows.Genres(context.Background(), ShowsGenreParams{})
	assert.NoError(t, err)

	assert.Equal(t, 35, len(genres))
}

func TestShowService_Seasons(t *testing.T) {
	data, err := os.ReadFile("data/shows/seasons.json")
	assert.NoError(t, err)

	ts, bc := setup(t, "GET", fmt.Sprintf("/%s", "shows/seasons?id=1161"), string(data))
	defer ts.Close()

	seasons, err := bc.Shows.Seasons(context.Background(), ShowsSeasonsParams{
		ID: Int(1161),
	})
	assert.NoError(t, err)

	assert.Equal(t, 8, len(seasons))
}

func TestShowService_Articles(t *testing.T) {
	data, err := os.ReadFile("data/shows/articles.json")
	assert.NoError(t, err)

	ts, bc := setup(t, "GET", fmt.Sprintf("/%s", "shows/articles?id=1456"), string(data))
	defer ts.Close()

	articles, err := bc.Shows.Articles(context.Background(), ShowsArticlesParams{
		ID: 1456,
	})
	assert.NoError(t, err)

	assert.Equal(t, 7, len(articles))

	assert.Equal(t, BoolFromString(true), articles[0].Sticky)
}

func TestShowService_Unrated(t *testing.T) {
	data, err := os.ReadFile("data/shows/unrated.json")
	assert.NoError(t, err)

	ts, bc := setup(t, "GET", fmt.Sprintf("/%s", "shows/unrated?date=all"), string(data))
	defer ts.Close()

	shows, err := bc.Shows.Unrated(context.Background(), ShowsUnratedParams{
		Date: String("all"),
	})
	assert.NoError(t, err)

	assert.Equal(t, 10, len(shows))
}
