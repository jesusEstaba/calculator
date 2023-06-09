package entities

import "github.com/jesusEstaba/calculator/pkg/domain"

type ErrorResponse struct {
	Error string `json:"error"`
}

type LoginResponse struct {
	Token string `json:"token"`
}

type SearchRecordsResponse struct {
	Records []*domain.Record `json:"records"`
}
