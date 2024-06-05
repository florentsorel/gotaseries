package gotaseries

type showsResponse struct {
	Shows  []Show `json:"shows"`
	Errors Errors `json:"errors"`
}

type showResponse struct {
	Show   Show   `json:"show"`
	Errors Errors `json:"errors"`
}

type year int

type seasonsDetail struct {
	Number   int `json:"number"`
	Episodes int `json:"episodes"`
}

type showrunner struct {
	ID      int    `json:"id,string"`
	Name    string `json:"name"`
	Picture string `json:"picture"`
}

type platforms struct {
	Svods []svod `json:"svods"`
	Svod  *svod  `json:"svod"`
}

type svod struct {
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

// Show represents a TV show.
type Show struct {
	ID             int             `json:"id"`
	TheTvdbID      int             `json:"thetvdb_id"`
	ImdbID         string          `json:"imdb_id"`
	MoviedbID      int             `json:"themoviedb_id"`
	Slug           string          `json:"slug"`
	Title          string          `json:"title"`
	OriginalTitle  string          `json:"original_title"`
	Description    string          `json:"description"`
	Seasons        int8            `json:"seasons,string"`
	SeasonsDetails []seasonsDetail `json:"seasons_details"`
	Episodes       int             `json:"episodes,string"`
	Followers      int64           `json:"followers,string"`
	Comments       int64           `json:"comments"`
	Similars       int             `json:"similars,string"`
	Characters     int             `json:"characters,string"`
	Creation       year            `json:"creation,string"`
	Showrunner     *showrunner     `json:"showrunner"`
	Showrunners    []showrunner    `json:"showrunners"`
	Genres         genres          `json:"genres"`
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
	Aliases     alias `json:"aliases"`
	SocialLinks []struct {
		Type       string `json:"type"`
		ExternalID string `json:"external_id"`
	} `json:"social_links"`
	NextTrailer     *string    `json:"next_trailer"`
	NextTrailerHost *string    `json:"next_trailer_host"`
	ResourceURL     string     `json:"resource_url"`
	Platforms       *platforms `json:"platforms"`
}
