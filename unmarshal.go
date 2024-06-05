package gotaseries

import (
	"encoding/json"
)

type alias map[int]string

func (a *alias) UnmarshalJSON(data []byte) error {
	if string(data) == "[]" {
		*a = map[int]string{}
		return nil
	}

	var aliases map[int]string

	if err := json.Unmarshal(data, &aliases); err != nil {
		return err
	}

	*a = aliases

	return nil
}

type genres []string

func (genres *genres) UnmarshalJSON(data []byte) error {
	if string(data) == "[]" {
		*genres = []string{}
		return nil
	}

	var g map[string]string

	if err := json.Unmarshal(data, &g); err != nil {
		return err
	}

	var result []string
	for _, value := range g {
		result = append(result, value)
	}

	*genres = result

	return nil
}
