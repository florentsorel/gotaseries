package gotaseries

import (
	"errors"
)

const (
	RecommendationStatusWait    RecommendationStatus = "wait"
	RecommendationStatusAccept  RecommendationStatus = "accept"
	RecommendationStatusDecline RecommendationStatus = "decline"
)

type RecommendationStatus string

type recommendationResponse struct {
	Recommendation Recommendation `json:"recommendation"`
	Errors         Errors         `json:"errors"`
}

func (rs *RecommendationStatus) IsValid() error {
	switch *rs {
	case RecommendationStatusWait, RecommendationStatusAccept, RecommendationStatusDecline:
		return nil
	}
	return errors.New("invalid RecommendationStatus")
}

type Recommendation struct {
	ID      int                  `json:"id"`
	From    int                  `json:"from_id"`
	To      int                  `json:"to_id"`
	ShowId  int                  `json:"show_id"`
	Status  RecommendationStatus `json:"status"`
	Comment *string              `json:"comments"`
}
