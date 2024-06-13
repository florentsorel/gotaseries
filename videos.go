package gotaseries

const (
	OrderDateASC  OrderDateType = "date"
	OrderDateDESC OrderDateType = "-date"
)

type OrderDateType string

type videosShowResponse struct {
	Videos []VideoShow `json:"videos"`
	Errors Errors      `json:"errors"`
}

type VideoShow struct {
	ID      int      `json:"id"`
	ShowID  int      `json:"show_id"`
	Host    string   `json:"host"`
	Slug    string   `json:"slug"`
	URL     string   `json:"url"`
	Date    DateTime `json:"date"`
	Title   string   `json:"title"`
	Type    string   `json:"type"`
	Season  int      `json:"season"`
	Episode int      `json:"episode"`
}
