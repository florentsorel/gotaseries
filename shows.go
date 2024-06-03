package gotaseries

import (
	"context"
	"fmt"
	"net/http"
)

type showService service

type showInterface interface {
	Display(id int) (*Show, error)
	DisplayWithCtx(ctx context.Context, id int) (*Show, error)
	List() ([]Show, error)
	ListWithCtx(ctx context.Context) ([]Show, error)
}

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
	Svod  *Svod  `json:"svod"`
}

type Svod struct {
	ID        int     `json:"id"`
	Name      string  `json:"name"`
	Tag       *string `json:"tag"`
	Color     string  `json:"color"`
	LinkURL   string  `json:"link_url"`
	Available struct {
		Last  int `json:"last,omitempty"`
		First int `json:"first,omitempty"`
	} `json:"available"`
	Logo *string `json:"logo"`
}

type Show struct {
	ID             int             `json:"id"`
	TvdbID         int             `json:"thetvdb_id"`
	ImdbID         string          `json:"imdb_id"`
	MoviedbID      int             `json:"themoviedb_id"`
	Slug           string          `json:"slug"`
	Title          string          `json:"title"`
	OriginalTitle  string          `json:"original_title"`
	Description    string          `json:"description"`
	Seasons        int8            `json:"seasons,string"`
	SeasonsDetails []SeasonsDetail `json:"seasons_details"`
	Episodes       int             `json:"episodes,string"`
	Followers      int64           `json:"followers,string"`
	Comments       int64           `json:"comments"`
	Similars       int             `json:"similars,string"`
	Characters     int             `json:"characters,string"`
	Creation       Year            `json:"creation,string"`
	Showrunner     *Showrunner     `json:"showrunner"`
	Showrunners    []Showrunner    `json:"showrunners"`
	Genres         Genres          `json:"genres"`
	Length         int             `json:"length,string"`
	Network        string          `json:"network"`
	Country        string          `json:"country"`
	Rating         string          `json:"rating"`
	Status         string          `json:"status"`
	Language       string          `json:"language"`
	Notes          struct {
		Total int     `json:"total"`
		Mean  float32 `json:"mean"`
	} `json:"notes"`
	Image struct {
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
	Platforms       *Platforms `json:"platforms"`
}

func (s *showService) display(ctx context.Context, id int) (*Show, error) {
	req, err := s.client.newRequest(http.MethodGet, fmt.Sprintf("/shows/display?id=%d", id))
	if err != nil {
		return nil, err
	}

	if ctx != nil {
		req = req.WithContext(ctx)
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

func (s *showService) Display(id int) (*Show, error) {
	return s.display(context.TODO(), id)
}

func (s *showService) DisplayWithCtx(ctx context.Context, id int) (*Show, error) {
	return s.display(ctx, id)
}

func (s *showService) list(ctx context.Context) ([]Show, error) {
	req, err := s.client.newRequest(http.MethodGet, "/shows/list")
	if err != nil {
		return nil, err
	}

	if ctx != nil {
		req = req.WithContext(ctx)
	}

	var shows showsResponse
	err = s.client.do(req, &shows)
	if err != nil {
		return nil, err
	}

	return shows.Shows, nil
}

func (s *showService) List() ([]Show, error) {
	return s.list(context.TODO())
}

func (s *showService) ListWithCtx(ctx context.Context) ([]Show, error) {
	return s.list(ctx)
}
