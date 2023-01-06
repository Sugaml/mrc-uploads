package controller

import (
	"os"

	"github.com/gofiber/fiber/v2"
)

type Controller struct {
	Router *fiber.App
}

func NewController() *Controller {
	ctl := &Controller{}
	ctl.Router = fiber.New()
	ctl.InitializeRoutes()
	return ctl
}

func (ctl *Controller) Run() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8083"
	}
	ctl.Router.Listen(":" + port)
}
