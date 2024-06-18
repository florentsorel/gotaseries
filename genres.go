package gotaseries

type genresResponse struct {
	Genres Genres `json:"genres"`
	Errors Errors `json:"errors"`
}
