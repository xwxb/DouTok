package handler

import (
	"github.com/TremblingV5/DouTok/applications/relationDomain/service"
)

type RelationDomainServiceImpl struct {
}

type Handler struct {
	service *service.Service
}

func New(service *service.Service) *Handler {
	return &Handler{
		service: service,
	}
}
