package handler

import "interview/pkg/core/common"

type Handler struct {
	TaxController *TaxController
}

func NewHandler(s common.Services) *Handler {
	return &Handler{TaxController: NewTaxController(s.Cart)}
}
