package gotaseries

import (
	"bytes"
	"errors"
	"fmt"
)

type errorableResponse interface {
	GetErrors() Errors
}

func (s *showsResponse) GetErrors() Errors {
	return s.Errors
}

func (s *showResponse) GetErrors() Errors {
	return s.Errors
}

func (s *picturesShowResponse) GetErrors() Errors {
	return s.Errors
}

func (s *videosShowResponse) GetErrors() Errors {
	return s.Errors
}

func (s *charactersShowResponse) GetErrors() Errors {
	return s.Errors
}

func (e *episodesResponse) GetErrors() Errors {
	return e.Errors
}

func (s *similarsResponse) GetErrors() Errors {
	return s.Errors
}

func (s *recommendationsResponse) GetErrors() Errors {
	return s.Errors
}

func (s *recommendationResponse) GetErrors() Errors {
	return s.Errors
}

type Error struct {
	Code    int    `json:"code"`
	Message string `json:"text"`
}

type Errors []Error

func (errs Errors) Err() error {
	if len(errs) == 0 {
		return nil
	}

	b := bytes.NewBuffer(nil)
	for _, e := range errs {
		_, _ = fmt.Fprintf(b, "Code: %d, Message: %s\n", e.Code, e.Message)
	}

	return errors.New(b.String())
}
