package controller

import (
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func (ctl *Controller) InitializeRoutes() {
	//middleware to allow all clients to communicate using http and allow cors
	ctl.Router.Use(cors.New())

	ctl.Router.Static("/uploads", "./uploads")
	routes := ctl.Router.Group("/uploads")
	routes.Post("/", ctl.handleFileupload)
	// handle image uploading using post request

	// delete uploaded image by providing unique image name

	routes.Delete("/:imageName", ctl.handleDeleteImage)
	// // delete uploaded image by providing unique image name

	//ctl.Router.Delete("/:imageName", handleDeleteImage)
}
