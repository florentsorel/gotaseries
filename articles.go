package gotaseries

type articlesResponse struct {
	Articles []Article `json:"articles"`
	Errors   Errors    `json:"errors"`
}

type Article struct {
	ID      int            `json:"id,string"`
	Date    DateTime       `json:"date"`
	Title   string         `json:"title"`
	Excerpt *string        `json:"excerpt"`
	Content string         `json:"content"`
	Slug    string         `json:"slug"`
	Image   string         `json:"image"`
	Sticky  BoolFromString `json:"sticky"`
}
