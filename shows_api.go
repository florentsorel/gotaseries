package gotaseries

import (
	"fmt"
	"net/http"
)

type showService service

func (s *showService) Display(id int) (*Show, error) {
	req, err := s.client.newRequest(http.MethodGet, fmt.Sprintf("/shows/display?id=%d", id), nil)
	if err != nil {
		return nil, err
	}

	var show ShowResponse
	err = s.client.do(req, &show)
	if err != nil {
		return nil, err
	}

	return &show.Show, nil
}

func (s *showService) List() ([]*Show, error) {
	req, err := s.client.newRequest(http.MethodGet, "shows/list", nil)
	if err != nil {
		return nil, err
	}

	var shows showResponse
	err = s.client.do(req, &shows)
	if err != nil {
		return nil, err
	}

	return shows.Shows, nil
}
