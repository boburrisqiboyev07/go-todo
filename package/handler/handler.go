package handler

import (
	"github.com/gin-gonic/gin"

	"github.com/boburrisqiboyev07/go-todo.git/package/service"

)

type Handler struct {
	services *service.Service
}
func NewHandler(services *service.Service) *Handler {
	return &Handler{services: services}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()
	auth := router.Group("/auth")
	{
		auth.POST("/sign-up", h.signUp)
		auth.POST("/sign-in", h.signIn)
	}
	api := router.Group("/api")
	{
		lists := api.Group("/lists")
		{
			lists.POST("/", h.createList)
			lists.GET("/", h.getAllList)
			lists.GET("/:id", h.GetListById)
			lists.PUT("/:id", h.updateList)
			lists.DELETE("/:id", h.deleteList)
			items := lists.Group("/:id/")
			{
				items.POST(":id/", h.createItem)
				items.GET("/", h.getAllItems)
				items.GET("/:item_id", h.GetItemById)
				items.PUT("/:item_id", h.updateItem)
				items.DELETE("/:item_id", h.deleteItem)
			}
		}
	}
	return router
}
