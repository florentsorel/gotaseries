package gotaseries

import (
	"context"
	"net/http"
	"time"
)

const (
	OrderAlphabetical OrderType = "alphabetical"
	OrderPopularity   OrderType = "popularity"
	OrderFollowers    OrderType = "followers"
)

type ShowService Service

type OrderType string

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

type ShowsDisplayParams struct {
	ID        *int        `url:"id"`
	TheTvdbID *int        `url:"thetvdb_id"`
	ImbdID    *string     `url:"imdb_id"`
	URL       *string     `url:"url"`
	Locale    *LocaleType `url:"locale"`
}

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

type ShowsRandomParams struct {
	Number  *int        `url:"nb"`
	Summary *bool       `url:"summary"`
	Locale  *LocaleType `url:"locale"`
}

type ShowsEpisodesParams struct {
	ID        *int        `url:"id"`
	TheTvdbID *int        `url:"thetvdb_id"`
	Season    *int        `url:"season"`
	Episode   *int        `url:"episode"`
	Subtitles *bool       `url:"subtitles"`
	Locale    *LocaleType `url:"locale"`
}

type ShowsSimilarsParams struct {
	ID        *int        `url:"id"`
	TheTvdbID *int        `url:"thetvdb_id"`
	Details   *bool       `url:"details"`
	Locale    *LocaleType `url:"locale"`
}

type ShowsVideosParams struct {
	ID        *int           `url:"id"`
	TheTvdbID *int           `url:"thetvdb_id"`
	Order     *OrderDateType `url:"order"`
	Start     *int           `url:"start"`
	Limit     *int           `url:"limit"`
	Details   *bool          `url:"details"`
	Locale    *LocaleType    `url:"locale"`
}

type ShowsCharactersParams struct {
	ID        *int        `url:"id"`
	TheTvdbID *int        `url:"thetvdb_id"`
	Locale    *LocaleType `url:"locale"`
}

type ShowsPicturesParams struct {
	ID        *int           `url:"id"`
	TheTvdbID *int           `url:"thetvdb_id"`
	Order     *OrderDateType `url:"order"`
	Start     *int           `url:"start"`
	Limit     *int           `url:"limit"`
	Format    *FormatType    `url:"format"`
	Locale    *LocaleType    `url:"locale"`
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
	var res showResponse
	if err := s.doRequest(ctx, http.MethodGet, "/shows/display", params, &res); err != nil {
		return nil, err
	}
	return &res.Show, nil
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
	var res showsResponse
	if err := s.doRequest(ctx, http.MethodGet, "/shows/list", params, &res); err != nil {
		return nil, err
	}
	return res.Shows, nil
}

// Random returns a list of random series.
func (s *ShowService) Random(ctx context.Context, params ShowsRandomParams) ([]Show, error) {
	var res showsResponse
	if err := s.doRequest(ctx, http.MethodGet, "/shows/random", params, &res); err != nil {
		return nil, err
	}
	return res.Shows, nil
}

// Episodes returns a list of episodes for the series.
func (s *ShowService) Episodes(ctx context.Context, params ShowsEpisodesParams) ([]Episode, error) {
	var res episodesResponse
	if err := s.doRequest(ctx, http.MethodGet, "/shows/episodes", params, &res); err != nil {
		return nil, err
	}
	return res.Episodes, nil
}

// Similars returns a list of characters for the series.
func (s *ShowService) Similars(ctx context.Context, params ShowsSimilarsParams) ([]SimilarShow, error) {
	var res similarsResponse
	if err := s.doRequest(ctx, http.MethodGet, "/shows/similars", params, &res); err != nil {
		return nil, err
	}
	return res.Similars, nil
}

// Videos returns a list of videos for the series.
func (s *ShowService) Videos(ctx context.Context, params ShowsVideosParams) ([]VideoShow, error) {
	var res videosShowResponse
	if err := s.doRequest(ctx, http.MethodGet, "/shows/videos", params, &res); err != nil {
		return nil, err
	}
	return res.Videos, nil
}

// Characters returns a list of characters for the series.
func (s *ShowService) Characters(ctx context.Context, params ShowsCharactersParams) ([]CharacterShow, error) {
	var res charactersShowResponse
	if err := s.doRequest(ctx, http.MethodGet, "/shows/characters", params, &res); err != nil {
		return nil, err
	}
	return res.Characters, nil
}

// Pictures returns a list of pictures for the series.
func (s *ShowService) Pictures(ctx context.Context, params ShowsPicturesParams) ([]PictureShow, error) {
	var res picturesShowResponse
	if err := s.doRequest(ctx, http.MethodGet, "/shows/pictures", params, &res); err != nil {
		return nil, err
	}
	return res.Pictures, nil
}
