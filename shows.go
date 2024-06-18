package gotaseries

import (
	"context"
	"errors"
	"net/http"
	"time"
)

const (
	OrderAlphabetical OrderType = "alphabetical"
	OrderTitle        OrderType = "title"
	OrderPopularity   OrderType = "popularity"
	OrderFollowers    OrderType = "followers"

	OrderShowMemberAlphabetical      OrderShowMemberType = "alphabetical"
	OrderShowMemberProgression       OrderShowMemberType = "progression"
	OrderShowMemberRemainingTime     OrderShowMemberType = "remaining_time"
	OrderShowMemberRemainingEpisodes OrderShowMemberType = "remaining_episodes"
	OrderShowMemberLastSeen          OrderShowMemberType = "last_seen"
	OrderShowMemberLastAdded         OrderShowMemberType = "last_added"
	OrderShowMemberRating            OrderShowMemberType = "rating"
	OrderShowMemberAvgRating         OrderShowMemberType = "avg_rating"
	OrderShowMemberCustom            OrderShowMemberType = "custom"
	OrderShowMemberNextDate          OrderShowMemberType = "next_date"

	StatusShowMemberCurrent                 StatusShowMemberType = "current"
	StatusShowMemberActive                  StatusShowMemberType = "active"
	StatusShowMemberArchived                StatusShowMemberType = "archived"
	StatusShowMemberArchivedAndCompleted    StatusShowMemberType = "archived_and_completed"
	StatusShowMemberArchivedAndNotCompleted StatusShowMemberType = "archived_and_not_completed"
	StatusShowMemberCompleted               StatusShowMemberType = "completed"
	StatusShowMemberActiveAndCompleted      StatusShowMemberType = "active_and_completed"
	StatusShowMemberNotStarted              StatusShowMemberType = "not_started"
	StatusShowMemberStopped                 StatusShowMemberType = "stopped"
)

type ShowService Service

type OrderType string

type OrderShowMemberType string

func (osm *OrderShowMemberType) IsValid() error {
	switch *osm {
	case OrderShowMemberAlphabetical, OrderShowMemberProgression, OrderShowMemberRemainingTime, OrderShowMemberRemainingEpisodes, OrderShowMemberLastSeen, OrderShowMemberLastAdded, OrderShowMemberRating, OrderShowMemberAvgRating, OrderShowMemberCustom, OrderShowMemberNextDate:
		return nil
	}
	return errors.New("invalid OrderShowMemberType")
}

type StatusShowMemberType string

func (ssm *StatusShowMemberType) IsValid() error {
	switch *ssm {
	case StatusShowMemberCurrent, StatusShowMemberActive, StatusShowMemberArchived, StatusShowMemberArchivedAndCompleted, StatusShowMemberArchivedAndNotCompleted, StatusShowMemberCompleted, StatusShowMemberActiveAndCompleted, StatusShowMemberNotStarted, StatusShowMemberStopped:
		return nil
	}
	return errors.New("invalid OrderShowMemberType")
}

type showsResponse struct {
	Shows  []Show `json:"shows"`
	Errors Errors `json:"errors"`
}

type showResponse struct {
	Show   Show   `json:"show"`
	Errors Errors `json:"errors"`
}

type ShowsMemberResponse struct {
	Shows             []Show `json:"shows"`
	Total             int    `json:"total"`
	TotalMissingShows int    `json:"totalMissingShows"`
	Errors            Errors `json:"errors"`
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

type Tags []string

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
	Note           Note            `json:"notes"`
	InAccount      bool            `json:"in_account"`
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
	User struct {
		Archived  bool    `json:"archived"`
		Favorited bool    `json:"favorited"`
		Remaining int     `json:"remaining"`
		Status    float64 `json:"status"`
		Last      string  `json:"last"`
		Tags      Tags    `json:"tags"`
		Next      struct {
			ID    *int    `json:"id"`
			Code  string  `json:"code"`
			Date  *Date   `json:"date"`
			Title *string `json:"title"`
			Image *string `json:"image"`
		} `json:"next"`
		FriendsWatching []struct {
			ID     int    `json:"id"`
			Login  string `json:"login"`
			Note   *int   `json:"note"`
			Avatar string `json:"avatar"`
		} `json:"friends_watching"`
	} `json:"user"`
	NextTrailer     *string    `json:"next_trailer"`
	NextTrailerHost *string    `json:"next_trailer_host"`
	ResourceURL     string     `json:"resource_url"`
	Platforms       *Platforms `json:"Platforms"`
}

type ShowsAddNoteParams struct {
	ID        *int        `url:"id"`
	TheTvdbID *int        `url:"thetvdb_id"`
	Note      int         `url:"note"`
	Locale    *LocaleType `url:"locale"`
}

type ShowsDeleteNoteParams struct {
	ID        *int        `url:"id"`
	TheTvdbID *int        `url:"thetvdb_id"`
	Locale    *LocaleType `url:"locale"`
}

type ShowsSearchParams struct {
	Title     *string     `url:"title"`
	Summary   *bool       `url:"summary"`
	Order     *OrderType  `url:"order"`
	Recent    *bool       `url:"recent"`
	Platforms []int       `url:"platforms"`
	Country   *string     `url:"country"`
	PerPage   *int        `url:"nbpp"`
	Page      *int        `url:"page"`
	Locale    *LocaleType `url:"locale"`
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
	Platforms *string     `url:"platforms"`
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

type ShowsAddParams struct {
	ID        *int        `url:"id"`
	TheTvdbID *int        `url:"thetvdb_id"`
	ImdbID    *string     `url:"imdb_id"`
	EpisodeID *int        `url:"episode_id"`
	Locale    *LocaleType `url:"locale"`
}

type ShowsDeleteParams struct {
	ID        *int        `url:"id"`
	TheTvdbID *int        `url:"thetvdb_id"`
	ImdbID    *string     `url:"imdb_id"`
	Locale    *LocaleType `url:"locale"`
}

type ShowsArchiveParams struct {
	ID        *int        `url:"id"`
	TheTvdbID *int        `url:"thetvdb_id"`
	Locale    *LocaleType `url:"locale"`
}

type ShowsUnarchiveParams struct {
	ID        *int        `url:"id"`
	TheTvdbID *int        `url:"thetvdb_id"`
	Locale    *LocaleType `url:"locale"`
}

type ShowsCreateRecommendationParams struct {
	ID        *int        `url:"id"`
	TheTvdbID *int        `url:"thetvdb_id"`
	To        int         `url:"to"`
	Comment   *string     `url:"comments"`
	Locale    *LocaleType `url:"locale"`
}

type ShowsUpdateRecommendationParams struct {
	ID     int                  `url:"id"`
	Status RecommendationStatus `url:"status"`
	Locale *LocaleType          `url:"locale"`
}

type ShowsDeleteRecommendationParams struct {
	ID     int         `url:"id"`
	Locale *LocaleType `url:"locale"`
}

type ShowsRecommendationsParams struct {
	Locale *LocaleType `url:"locale"`
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

type ShowsFavoritesParams struct {
	ID      *int                `url:"id"`
	Order   *OrderFavoriteType  `url:"order"`
	Limit   *int                `url:"limit"`
	Offset  *int                `url:"offset"`
	Status  *StatusFavoriteType `url:"status"`
	Summary *bool               `url:"summary"`
	Locale  *LocaleType         `url:"locale"`
}

type ShowsAddFavoriteParams struct {
	ID     int         `url:"id"`
	Locale *LocaleType `url:"locale"`
}

type ShowsDeleteFavoriteParams struct {
	ID     int         `url:"id"`
	Locale *LocaleType `url:"locale"`
}

type ShowsUpdateTagsParams struct {
	ID   int      `url:"id"`
	Tags []string `url:"tags"`
}

type ShowsMemberParams struct {
	ID               *int                  `url:"id"`
	Order            *OrderShowMemberType  `url:"order"`
	Limit            *int                  `url:"limit"`
	Offset           *int                  `url:"offset"`
	Status           *StatusShowMemberType `url:"status"`
	ExcludedGenres   *string               `url:"excluded_genres"`
	ExcludedNetworks *string               `url:"excluded_networks"`
	ExcludedStatus   *string               `url:"excluded_status"`
	Tags             *string               `url:"tags"`
	ExcludedTags     *string               `url:"excluded_tags"`
	Summary          *bool                 `url:"summary"`
	Platforms        *string               `url:"platforms"`
	Locale           *LocaleType           `url:"locale"`
}

type ShowsDiscoverParams struct {
	Limit   *int        `url:"limit"`
	Offset  *int        `url:"offset"`
	Summary *bool       `url:"summary"`
	Locale  *LocaleType `url:"locale"`
}

type ShowsDiscoverPlatformsParams struct {
	Summary *bool       `url:"summary"`
	Locale  *LocaleType `url:"locale"`
}

type ShowsGenreParams struct {
	Locale *LocaleType `url:"locale"`
}

type ShowsSeasonsParams struct {
	ID        *int        `url:"id"`
	TheTvdbID *int        `url:"thetvdb_id"`
	Locale    *LocaleType `url:"locale"`
}

type ShowsArticlesParams struct {
	ID     int         `url:"id"`
	Locale *LocaleType `url:"locale"`
}

type ShowsUnratedParams struct {
	PerPage *int        `url:"nbpp"`
	Page    *int        `url:"page"`
	Date    *string     `url:"date"`
	Locale  *LocaleType `url:"locale"`
}

// AddNote rate a series.
// Require a valid token.
func (s *ShowService) AddNote(ctx context.Context, params ShowsAddNoteParams) (*Show, error) {
	var res showResponse
	if err := s.doRequest(ctx, http.MethodPost, "/shows/note", params, &res); err != nil {
		return nil, err
	}
	return &res.Show, nil
}

// DeleteNote delete a series rating.
// Require a valid token.
func (s *ShowService) DeleteNote(ctx context.Context, params ShowsDeleteNoteParams) (*Show, error) {
	var res showResponse
	if err := s.doRequest(ctx, http.MethodDelete, "/shows/note", params, &res); err != nil {
		return nil, err
	}
	return &res.Show, nil
}

// Search returns a list of series matching the search query, with member information if a  token is provided.
func (s *ShowService) Search(ctx context.Context, params ShowsSearchParams) ([]Show, error) {
	var res showsResponse
	if err := s.doRequest(ctx, http.MethodGet, "/shows/search", params, &res); err != nil {
		return nil, err
	}
	return res.Shows, nil
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

// Add add a series to the member's account.
// Require a valid token.
func (s *ShowService) Add(ctx context.Context, params ShowsAddParams) (*Show, error) {
	var res showResponse
	if err := s.doRequest(ctx, http.MethodPost, "/shows/show", params, &res); err != nil {
		return nil, err
	}
	return &res.Show, nil
}

// Delete remove a series from the member's account.
// Require a valid token.
func (s *ShowService) Delete(ctx context.Context, params ShowsDeleteParams) (*Show, error) {
	var res showResponse
	if err := s.doRequest(ctx, http.MethodDelete, "/shows/show", params, &res); err != nil {
		return nil, err
	}
	return &res.Show, nil
}

// Archive archive a series in the member's account.
func (s *ShowService) Archive(ctx context.Context, params ShowsArchiveParams) (*Show, error) {
	var res showResponse
	if err := s.doRequest(ctx, http.MethodPost, "/shows/archive", params, &res); err != nil {
		return nil, err
	}
	return &res.Show, nil
}

// Unarchive remove a series from the archives of the member's account.
func (s *ShowService) Unarchive(ctx context.Context, params ShowsUnarchiveParams) (*Show, error) {
	var res showResponse
	if err := s.doRequest(ctx, http.MethodDelete, "/shows/archive", params, &res); err != nil {
		return nil, err
	}
	return &res.Show, nil
}

// CreateRecommendation create a series recommendation from authenticated member to a friend.
// Require a valid token.
func (s *ShowService) CreateRecommendation(ctx context.Context, params ShowsCreateRecommendationParams) (*Recommendation, error) {
	var res recommendationResponse
	if err := s.doRequest(ctx, http.MethodPost, "/shows/recommendation", params, &res); err != nil {
		return nil, err
	}
	return &res.Recommendation, nil
}

// UpdateRecommendation update the status of a recommendation.
// Require a valid token.
func (s *ShowService) UpdateRecommendation(ctx context.Context, params ShowsUpdateRecommendationParams) (*Recommendation, error) {
	var res recommendationResponse
	if err := s.doRequest(ctx, http.MethodPut, "/shows/recommendation", params, &res); err != nil {
		return nil, err
	}
	return &res.Recommendation, nil
}

// DeleteRecommendation delete a recommendation.
// Require a valid token.
func (s *ShowService) DeleteRecommendation(ctx context.Context, params ShowsDeleteRecommendationParams) (*Recommendation, error) {
	var res recommendationResponse
	if err := s.doRequest(ctx, http.MethodDelete, "/shows/recommendation", params, &res); err != nil {
		return nil, err
	}
	return &res.Recommendation, nil
}

// Recommendations returns a list of recommendations send and received for the authenticated member.
// Require a valid token.
func (s *ShowService) Recommendations(ctx context.Context, params ShowsRecommendationsParams) ([]Recommendation, error) {
	var res recommendationsResponse
	if err := s.doRequest(ctx, http.MethodGet, "/shows/recommendations", params, &res); err != nil {
		return nil, err
	}
	return res.Recommendations, nil
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

// Favorites returns a list of favorite series for the authenticated member or ID member. (ID member has priority over token)
func (s *ShowService) Favorites(ctx context.Context, params ShowsFavoritesParams) (*FavoritesResponse, error) {
	var res FavoritesResponse
	if err := s.doRequest(ctx, http.MethodGet, "/shows/favorites", params, &res); err != nil {
		return nil, err
	}
	return &res, nil
}

// AddFavorite add a series to the member's favorite list.
// Require a valid token.
func (s *ShowService) AddFavorite(ctx context.Context, params ShowsAddFavoriteParams) (*Show, error) {
	var res showResponse
	if err := s.doRequest(ctx, http.MethodPost, "/shows/favorite", params, &res); err != nil {
		return nil, err
	}
	return &res.Show, nil
}

// DeleteFavorite delete a series from the member's favorite.
// Require a valid token.
func (s *ShowService) DeleteFavorite(ctx context.Context, params ShowsDeleteFavoriteParams) (*Show, error) {
	var res showResponse
	if err := s.doRequest(ctx, http.MethodDelete, "/shows/favorite", params, &res); err != nil {
		return nil, err
	}
	return &res.Show, nil
}

// UpdateTags update the tags of a series for authenticated member.
// Require a valid token.
func (s *ShowService) UpdateTags(ctx context.Context, params ShowsUpdateTagsParams) (*Show, error) {
	var res showResponse
	if err := s.doRequest(ctx, http.MethodPost, "/shows/tags", params, &res); err != nil {
		return nil, err
	}
	return &res.Show, nil
}

// Member returns a list of series which belongs to the authenticated member or ID member. (ID member has priority over token)
func (s *ShowService) Member(ctx context.Context, params ShowsMemberParams) (*ShowsMemberResponse, error) {
	var res ShowsMemberResponse
	if err := s.doRequest(ctx, http.MethodGet, "/shows/member", params, &res); err != nil {
		return nil, err
	}
	return &res, nil
}

// Discover returns a list of series to discover.
func (s *ShowService) Discover(ctx context.Context, params ShowsDiscoverParams) ([]Show, error) {
	var res showsResponse
	if err := s.doRequest(ctx, http.MethodGet, "/shows/discover", params, &res); err != nil {
		return nil, err
	}
	return res.Shows, nil
}

// DiscoverPlatform returns a list of series to discover on major SVoD platforms.
func (s *ShowService) DiscoverPlatform(ctx context.Context, params ShowsDiscoverPlatformsParams) ([]Show, error) {
	var res showsResponse
	if err := s.doRequest(ctx, http.MethodGet, "/shows/discover_platform", params, &res); err != nil {
		return nil, err
	}
	return res.Shows, nil
}

// Genres returns the list of available series genres.
func (s *ShowService) Genres(ctx context.Context, params ShowsGenreParams) (Genres, error) {
	var res genresResponse
	if err := s.doRequest(ctx, http.MethodGet, "/shows/genres", params, &res); err != nil {
		return nil, err
	}

	return res.Genres, nil
}

// Seasons returns the list of seasons for a series.
func (s *ShowService) Seasons(ctx context.Context, params ShowsSeasonsParams) ([]Season, error) {
	var res seasonsResponse
	if err := s.doRequest(ctx, http.MethodGet, "/shows/seasons", params, &res); err != nil {
		return nil, err
	}
	return res.Seasons, nil
}

// Articles returns the list of articles for a series.
func (s *ShowService) Articles(ctx context.Context, params ShowsArticlesParams) ([]Article, error) {
	var res articlesResponse
	if err := s.doRequest(ctx, http.MethodGet, "/shows/articles", params, &res); err != nil {
		return nil, err
	}
	return res.Articles, nil
}

// Unrated retrieve the list of finished and unrated series for the authenticated member.
// Require a valid token.
func (s *ShowService) Unrated(ctx context.Context, params ShowsUnratedParams) ([]Show, error) {
	var res showsResponse
	if err := s.doRequest(ctx, http.MethodGet, "/shows/unrated", params, &res); err != nil {
		return nil, err
	}
	return res.Shows, nil
}
