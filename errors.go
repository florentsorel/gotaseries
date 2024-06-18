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

func (p *picturesShowResponse) GetErrors() Errors {
	return p.Errors
}

func (s *ShowsMemberResponse) GetErrors() Errors {
	return s.Errors
}

func (v *videosShowResponse) GetErrors() Errors {
	return v.Errors
}

func (c *charactersShowResponse) GetErrors() Errors {
	return c.Errors
}

func (e *episodesResponse) GetErrors() Errors {
	return e.Errors
}

func (s *similarsResponse) GetErrors() Errors {
	return s.Errors
}

func (r *recommendationsResponse) GetErrors() Errors {
	return r.Errors
}

func (r *recommendationResponse) GetErrors() Errors {
	return r.Errors
}

func (f *FavoritesResponse) GetErrors() Errors {
	return f.Errors
}

func (g *genresResponse) GetErrors() Errors {
	return g.Errors
}

func (s *seasonsResponse) GetErrors() Errors {
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
