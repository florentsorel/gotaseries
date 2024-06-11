package gotaseries

import (
	"context"
	"net/http"
	"time"
)

type ShowService Service

type OrderType string

const (
	OrderAlphabetical OrderType = "alphabetical"
	OrderPopularity   OrderType = "popularity"
	OrderFollowers    OrderType = "followers"
)

type showsResponse struct {
	Shows  []Show `json:"shows"`
	Errors Errors `json:"errors"`
}

type showResponse struct {
	Show   Show   `json:"show"`
	Errors Errors `json:"errors"`
}

type Year int

type SeasonsDetail struct {
	Number   int `json:"number"`
	Episodes int `json:"episodes"`
}

type Showrunner struct {
	ID      int    `json:"id,string"`
	Name    string `json:"name"`
	Picture string `json:"picture"`
}

type Platforms struct {
	Svods []Svod `json:"svods"`
	Svod  *Svod  `json:"Svod"`
	Vods  []Svod `json:"vod"`
}

type Available struct {
	Last  int `json:"last,omitempty"`
	First int `json:"first,omitempty"`
}

type Svod struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	Tag       *string   `json:"tag"`
	Color     string    `json:"color"`
	LinkURL   string    `json:"link_url"`
	Available Available `json:"Available"`
	Logo      *string   `json:"logo"`
}

type Show struct {
	ID             int             `json:"id"`
	TheTvdbID      int             `json:"thetvdb_id"`
	ImdbID         string          `json:"imdb_id"`
	MoviedbID      int             `json:"themoviedb_id"`
	Slug           string          `json:"slug"`
	Title          string          `json:"title"`
	OriginalTitle  string          `json:"original_title"`
	Description    string          `json:"description"`
	Seasons        int             `json:"seasons,string"`
	SeasonsDetails []SeasonsDetail `json:"seasons_details"`
	Episodes       int             `json:"episodes,string"`
	Followers      int64           `json:"followers,string"`
	Poster         string          `json:"poster"`
	Comments       int64           `json:"comments"`
	Similars       int             `json:"similars,string"`
	Characters     int             `json:"characters,string"`
	Creation       Year            `json:"creation,string"`
	Showrunner     *Showrunner     `json:"Showrunner"`
	Showrunners    []Showrunner    `json:"showrunners"`
	Genres         Genres          `json:"Genres"`
	Length         int             `json:"length,string"`
	Network        string          `json:"network"`
	Country        string          `json:"country"`
	Rating         string          `json:"rating"`
	Status         string          `json:"status"`
	Language       string          `json:"language"`
	Notes          Note            `json:"notes"`
	Image          struct {
		Show   string `json:"show"`
		Banner string `json:"banner"`
		Box    string `json:"box"`
		Poster string `json:"poster"`
		Logo   *struct {
			Url    string `json:"url"`
			Width  int    `json:"width,string"`
			Height int    `json:"height,string"`
		} `json:"clearlogo"`
	} `json:"images"`
	Aliases     Alias `json:"aliases"`
	SocialLinks []struct {
		Type       string `json:"type"`
		ExternalID string `json:"external_id"`
	} `json:"social_links"`
	NextTrailer     *string    `json:"next_trailer"`
	NextTrailerHost *string    `json:"next_trailer_host"`
	ResourceURL     string     `json:"resource_url"`
	Platforms       *Platforms `json:"Platforms"`
}

// ShowsDisplayParams represents parameter to the [ShowService.Display] method.
type ShowsDisplayParams struct {
	ID        *int        `url:"id"`
	TheTvdbID *int        `url:"thetvdb_id"`
	ImbdID    *string     `url:"imdb_id"`
	URL       *string     `url:"url"`
	Locale    *LocaleType `url:"locale"`
}

// ShowsListParams represents parameter to the [ShowService.List] method.
type ShowsListParams struct {
	Order     *OrderType  `url:"order"`
	Since     *time.Time  `url:"since"`
	Recent    *bool       `url:"recent"`
	Starting  *string     `url:"starting"`
	Start     *int        `url:"start"`
	Limit     *int        `url:"limit"`
	Filter    *string     `url:"filter"`
	Platforms *int        `url:"Platforms"`
	Country   *string     `url:"country"`
	Summary   *bool       `url:"summary"`
	Locale    *LocaleType `url:"locale"`
}

// ShowsRandomParams represents parameter to the [ShowService.Random] method.
type ShowsRandomParams struct {
	Number  *int        `url:"nb"`
	Summary *bool       `url:"summary"`
	Locale  *LocaleType `url:"locale"`
}

// ShowsEpisodesParams represents parameter to the [ShowService.Episodes] method.
type ShowsEpisodesParams struct {
	ID        *int        `url:"id"`
	TheTvdbID *int        `url:"thetvdb_id"`
	Season    *int        `url:"season"`
	Episode   *int        `url:"episode"`
	Subtitles *bool       `url:"subtitles"`
	Locale    *LocaleType `url:"locale"`
}

// Display returns information about a series.
//
// Example:
//
//	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
//	defer cancel()
//
//	shows, err := client.Shows.Display(ctx, ShowsDisplayParams{
//		ID: gotaseries.Int(1161),
//	})
//	if err != nil {
//		log.Fatal(err)
//	}
//	fmt.Printf("%+v\n", show)
func (s *ShowService) Display(ctx context.Context, params ShowsDisplayParams) (*Show, error) {
	req, err := s.client.newRequest(ctx, http.MethodGet, "/shows/display", params)
	if err != nil {
		return nil, err
	}

	var show showResponse
	err = s.client.do(req, &show)
	if err != nil {
		return nil, err
	}

	if err = show.Errors.Err(); err != nil {
		return nil, err
	}

	return &show.Show, nil
}

// List returns a list of all series.
//
// Example:
//
//	shows, err := client.Shows.List(context.Background(), ShowsListParams{
//		Order: gotaseries.String("popularity"),
//	})
//	if err != nil {
//		log.Fatal(err)
//	}
//	for _, show := range shows {
//		fmt.Printf("%s\n", show.Title)
//	}
func (s *ShowService) List(ctx context.Context, params ShowsListParams) ([]Show, error) {
	req, err := s.client.newRequest(ctx, http.MethodGet, "/shows/list", params)
	if err != nil {
		return nil, err
	}

	var shows showsResponse
	err = s.client.do(req, &shows)
	if err != nil {
		return nil, err
	}

	if err = shows.Errors.Err(); err != nil {
		return nil, err
	}

	return shows.Shows, nil
}

// Random returns random series.
func (s *ShowService) Random(ctx context.Context, params ShowsRandomParams) ([]Show, error) {
	req, err := s.client.newRequest(ctx, http.MethodGet, "/shows/random", params)
	if err != nil {
		return nil, err
	}

	var shows showsResponse
	err = s.client.do(req, &shows)
	if err != nil {
		return nil, err
	}

	if err = shows.Errors.Err(); err != nil {
		return nil, err
	}

	return shows.Shows, nil
}

// Episodes returns a list of episodes for the series.
func (s *ShowService) Episodes(ctx context.Context, params ShowsEpisodesParams) ([]Episode, error) {
	req, err := s.client.newRequest(ctx, http.MethodGet, "/shows/episodes", params)
	if err != nil {
		return nil, err
	}

	var episodes episodesResponse
	err = s.client.do(req, &episodes)
	if err != nil {
		return nil, err
	}

	if err = episodes.Errors.Err(); err != nil {
		return nil, err
	}

	return episodes.Episodes, nil
}
