package gotaseries

import (
	"context"
	"net/http"
)

type BadgeService Service

type badgeResponse struct {
	Badge  Badge  `json:"badge"`
	Errors Errors `json:"errors"`
}

type Badge struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Image       string `json:"picture_url"`
	Width       int    `json:"width"`
	Height      int    `json:"height"`
	Level       *int   `json:"level"`
}

type BadgesBadgeParams struct {
	ID     int         `url:"id"`
	Locale *LocaleType `url:"locale"`
}

// Badge returns a badge details.
func (b *BadgeService) Badge(ctx context.Context, params BadgesBadgeParams) (*Badge, error) {
	var res badgeResponse
	if err := b.client.doRequest(ctx, http.MethodGet, "/badges/badge", params, &res); err != nil {
		return nil, err
	}
	return &res.Badge, nil
}
