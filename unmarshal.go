package gotaseries

import (
	"encoding/json"
)

type alias map[int]string

type genres []string

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

func (g *genres) UnmarshalJSON(data []byte) error {
	if string(data) == "[]" {
		*g = []string{}
		return nil
	}

	var genres map[string]string

	if err := json.Unmarshal(data, &genres); err != nil {
		return err
	}

	var result []string
	for _, value := range genres {
		result = append(result, value)
	}

	*g = result

	return nil
}
