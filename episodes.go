package gotaseries

type episodesResponse struct {
	Episodes []Episode `json:"episodes"`
	Errors   Errors    `json:"errors"`
}

type PlatformLinks struct {
	PlatformID int     `json:"platform_id,string"`
	Platform   string  `json:"platform"`
	Color      string  `json:"color"`
	Type       string  `json:"type"`
	Logo       *string `json:"logo"`
	Link       string  `json:"link"`
}

type Episode struct {
	ID          int         `json:"id"`
	TheTvdbID   int         `json:"thetvdb_id"`
	YoutubeID   *string     `json:"youtube_id"`
	Title       string      `json:"title"`
	Season      int         `json:"season"`
	Episode     int         `json:"episode"`
	Code        string      `json:"code"`
	Global      int         `json:"global"`
	Description string      `json:"description"`
	Director    string      `json:"director"`
	Writers     []string    `json:"writers"`
	Special     BoolFromInt `json:"special"`
	Comments    int         `json:"comments"`
	ShowSlug    string      `json:"show_slug"`
	ResourceURL string      `json:"resource_url"`
	Note        Note        `json:"note"`
	Date        Date        `json:"date"`
	SeenTotal   int         `json:"seen_total"`
	Show        struct {
		ID          int    `json:"id"`
		TheTvdbID   int    `json:"thetvdb_id"`
		Title       string `json:"title"`
		Slug        string `json:"slug"`
		Creation    Year   `json:"creation,string"`
		Status      string `json:"status"`
		Description string `json:"description"`
	}
	PlaformLinks []PlatformLinks `json:"platform_links"`
	Subtitles    []Subtitle      `json:"subtitles"`
}
