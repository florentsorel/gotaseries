package gotaseries

// ShowsDisplayParams represents a parameter to the `Display` method.
type ShowsDisplayParams struct {
	ID        *int    `url:"id"`
	TheTvdbID *int    `url:"thetvdb_id"`
	ImbdID    *string `url:"imdb_id"`
	URL       *string `url:"url"`
	Locale    *locale `url:"locale"`
}

// ShowsListParams represents a parameter to the `List` method.
type ShowsListParams struct {
	Order  *string `url:"order"`
	Locale *locale `url:"locale"`
}