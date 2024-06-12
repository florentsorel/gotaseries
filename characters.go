package gotaseries

type charactersResponse struct {
	Characters []CharacterShow `json:"characters"`
	Errors     Errors          `json:"errors"`
}

type CharacterShow struct {
	ShowID   int    `json:"show_id"`
	PersonID int    `json:"person_id,string"`
	Name     string `json:"name"`
	Actor    string `json:"actor"`
	Picture  string `json:"picture"`
}
