package controller

import (
	"github.com/gofiber/fiber"
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
	ctl.Router.Listen(":4000")
}
