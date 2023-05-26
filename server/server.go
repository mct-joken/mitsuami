package server

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/mct-joken/mitsuami/repository/dummy"
	"github.com/mct-joken/mitsuami/server/controller"
)

func Start() {
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	itemController := controller.NewItemController(
		dummy.NewLogRepository(nil),
		dummy.NewItemRepository(nil),
	)

	api := e.Group("/api/v1")
	{
		items := api.Group("/items")
		{
			items.GET("/", itemController.GetItem)
			items.POST("/", itemController.CreateItem)

			items.GET("/:id", itemController.GetItemByID)
			items.POST("/:id", itemController.StartUsing)
			items.DELETE("/:id", itemController.EndUsing)
		}
	}
	if err := e.Start(":3000"); err != nil {
		panic(err)
	}
}
