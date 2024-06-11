package gotaseries

type Subtitle struct {
	ID       int      `json:"id"`
	Language string   `json:"language"`
	Source   string   `json:"source"`
	File     string   `json:"file"`
	Date     DateTime `json:"date"`
	URL      string   `json:"url"`
	Quality  int      `json:"quality"`
}
