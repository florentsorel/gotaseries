package gotaseries

import "time"

// ShowsDisplayParams represents a parameter to the [ShowService.Display] method.
type ShowsDisplayParams struct {
	ID        *int    `url:"id"`
	TheTvdbID *int    `url:"thetvdb_id"`
	ImbdID    *string `url:"imdb_id"`
	URL       *string `url:"url"`
	Locale    *locale `url:"locale"`
}

// ShowsListParams represents a parameter to the [ShowService.List] method.
type ShowsListParams struct {
	Order     *order     `url:"order"`
	Since     *time.Time `url:"since"`
	Recent    *bool      `url:"recent"`
	Starting  *string    `url:"starting"`
	Start     *int       `url:"start"`
	Limit     *int       `url:"limit"`
	Filter    *string    `url:"filter"`
	Platforms *int       `url:"platforms"`
	Country   *string    `url:"country"`
	Summary   *bool      `url:"summary"`
	Locale    *locale    `url:"locale"`
}
