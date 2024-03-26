package gotaseries

type showResponse struct {
	Shows []*Show `json:"shows"`
}
type ShowResponse struct {
	Show Show `json:"show"`
}

type Show struct {
	ID            int    `json:"id"`
	Title         string `json:"title"`
	OriginalTitle string `json:"original_title"`
}
