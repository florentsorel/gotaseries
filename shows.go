package gotaseries

import (
	"context"
	"net/http"
)

type ShowService Service

type order string

const (
	OrderAlphabetical order = "alphabetical"
	OrderPopularity   order = "popularity"
	OrderFollowers    order = "followers"
)

// Display returns information about a series.
//
// Example:
//
//	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
//	defer cancel()
//
//	shows, err := client.Shows.Display(ctx, ShowsDisplayParams{
//		ID: gotaseries.Int(1161),
//	})
//	if err != nil {
//		log.Fatal(err)
//	}
//	fmt.Printf("%+v\n", show)
func (s *ShowService) Display(ctx context.Context, params ShowsDisplayParams) (*Show, error) {
	req, err := s.client.newRequest(ctx, http.MethodGet, "/shows/display", params)
	if err != nil {
		return nil, err
	}

	var show showResponse
	err = s.client.do(req, &show)
	if err != nil {
		return nil, err
	}

	if err = show.Errors.Err(); err != nil {
		return nil, err
	}

	return &show.Show, nil
}

// List returns a list of all series.
//
// Example:
//
//	shows, err := client.Shows.List(context.Background(), ShowsListParams{
//		Order: gotaseries.String("popularity"),
//	})
//	if err != nil {
//		log.Fatal(err)
//	}
//	for _, show := range shows {
//		fmt.Printf("%s\n", show.Title)
//	}
func (s *ShowService) List(ctx context.Context, params ShowsListParams) ([]Show, error) {
	req, err := s.client.newRequest(ctx, http.MethodGet, "/shows/list", params)
	if err != nil {
		return nil, err
	}

	var shows showsResponse
	err = s.client.do(req, &shows)
	if err != nil {
		return nil, err
	}

	if err = shows.Errors.Err(); err != nil {
		return nil, err
	}

	return shows.Shows, nil
}

// Random returns random series.
func (s *ShowService) Random(ctx context.Context, params ShowsRandomParams) ([]Show, error) {
	req, err := s.client.newRequest(ctx, http.MethodGet, "/shows/random", params)
	if err != nil {
		return nil, err
	}

	var shows showsResponse
	err = s.client.do(req, &shows)
	if err != nil {
		return nil, err
	}

	if err = shows.Errors.Err(); err != nil {
		return nil, err
	}

	return shows.Shows, nil
}
