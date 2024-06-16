package gotaseries

const (
	OrderFavoriteAlphabetical      OrderFavoriteType = "alphabetical"
	OrderFavoriteProgression       OrderFavoriteType = "progression"
	OrderFavoriteRemainingTime     OrderFavoriteType = "remaining_time"
	OrderFavoriteRemainingEpisodes OrderFavoriteType = "remaining_episodes"

	StatusFavoritesCurrent  StatusFavoriteType = "current"
	StatusFavoritesActive   StatusFavoriteType = "active"
	StatusFavoritesArchived StatusFavoriteType = "archived"
)

type OrderFavoriteType string

type StatusFavoriteType string

type FavoritesResponse struct {
	Shows  []Show `json:"shows"`
	Total  int    `json:"total"`
	Errors Errors `json:"errors"`
}
