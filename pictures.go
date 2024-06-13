package gotaseries

const (
	FormatHD  FormatType = "hd"
	FormatAll FormatType = "all"
)

type FormatType string

type picturesShowResponse struct {
	Pictures []PictureShow `json:"pictures"`
	Errors   Errors        `json:"errors"`
}

type PictureShow struct {
	ID     int      `json:"id"`
	ShowID int      `json:"show_id"`
	URL    string   `json:"url"`
	Width  int      `json:"width"`
	Height int      `json:"height"`
	Date   DateTime `json:"date"`
	Picked string   `json:"picked"`
}
