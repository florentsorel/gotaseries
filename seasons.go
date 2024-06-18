package gotaseries

type seasonsResponse struct {
	Seasons []Season `json:"seasons"`
	Errors  Errors   `json:"errors"`
}

type Season struct {
	Number       int     `json:"number"`
	Episodes     int     `json:"episodes"`
	Seen         bool    `json:"seen"`
	Hidden       bool    `json:"hidden"`
	Image        *string `json:"image"`
	HasSubtitles bool    `json:"subtitles"`
	Notes        Note    `json:"notes"`
}
