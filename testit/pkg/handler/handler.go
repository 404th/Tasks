package handler

import (
	"github.com/404th/testit/pkg/service"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{
		services: services,
	}
}

func (h *Handler) InitRoute() *gin.Engine {
	r := gin.New()

	// Auth
	r.POST("/ctc/create", h.createContact)
	r.GET("/ctc/contact/:id", h.getContact)
	r.GET("/ctc/contacts", h.getAllContacts)
	r.PUT("/ctc/contact/update", h.updateContact)
	r.DELETE("/ctc/delete/:id", h.deleteContact)

	return r
}
