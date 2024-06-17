package gotaseries

import (
	"encoding/json"
	"strings"
	"time"
)

type Alias map[int]string

func (a *Alias) UnmarshalJSON(data []byte) error {
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

type Genres []string

func (genres *Genres) UnmarshalJSON(data []byte) error {
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

type BoolFromInt bool

func (b *BoolFromInt) UnmarshalJSON(data []byte) error {
	var intVal int
	if err := json.Unmarshal(data, &intVal); err != nil {
		return err
	}
	*b = intVal != 0
	return nil
}

type Date time.Time

func (d *Date) UnmarshalJSON(data []byte) error {
	var date string
	if err := json.Unmarshal(data, &date); err != nil {
		return err
	}

	t, err := time.Parse("2006-01-02", date)
	if err != nil {
		return err
	}

	*d = Date(t)

	return nil
}

func (d *Date) String() string {
	return time.Time(*d).Format("2006-01-02")
}

type DateTime time.Time

func (d *DateTime) UnmarshalJSON(data []byte) error {
	var date string
	if err := json.Unmarshal(data, &date); err != nil {
		return err
	}

	t, err := time.Parse("2006-01-02 15:04:05", date)
	if err != nil {
		return err
	}

	*d = DateTime(t)

	return nil
}

func (d *DateTime) String() string {
	return time.Time(*d).Format("2006-01-02 15:04:05")
}

func (rs *RecommendationStatus) UnmarshalJSON(data []byte) error {
	var status string
	if err := json.Unmarshal(data, &status); err != nil {
		return err
	}
	tempStatus := RecommendationStatus(status)
	if err := tempStatus.IsValid(); err != nil {
		return err
	}
	*rs = tempStatus
	return nil
}

func (t *Tags) UnmarshalJSON(data []byte) error {
	if string(data) == "\"\"" {
		*t = []string{}
		return nil
	}

	var tags string
	if err := json.Unmarshal(data, &tags); err != nil {
		return err
	}

	result := strings.Split(tags, ", ")

	*t = result

	return nil
}
