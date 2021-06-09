package handler

import (
	"github.com/404th/contactlist/pkg/service"
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

func (h *Handler) InitRoutes() *gin.Engine {
	r := gin.New()

	// AUTH
	r.POST("/auth/sign-up", h.signUp)
	r.POST("/auth/sign-in", h.signIn)

	// LISTS
	r.POST("/create/lists", h.createList)
	r.GET("/create/lists/:id", h.getList)
	r.GET("/create/lists", h.getALlLists)
	r.PUT("/create/lists/:id", h.updateList)
	r.DELETE("/create/lists/:id", h.deleteList)

	// TASKS
	r.POST("/create/lists/:id/item", h.createItem)
	r.GET("/create/lists/:id/item/:item_id", h.getItem)
	r.GET("/create/lists/:id/item", h.getAllItems)
	r.PUT("/create/lists/:id/item/:item_id", h.updateItem)
	r.DELETE("/create/lists/:id/item/:item_id", h.deleteItem)

	return r
}
