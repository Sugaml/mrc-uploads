package controller

import (
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func (ctl *Controller) InitializeRoutes() {
	//middleware to allow all clients to communicate using http and allow cors
	ctl.Router.Use(cors.New())

	// serve  images from images directory prefixed with /images
	// i.e http://localhost:4000/images/someimage.webp

	//ctl.Router.Static("/images", "./images")

	// handle image uploading using post request
	//ctl.Router.Post("/", handleFileupload)

	// // delete uploaded image by providing unique image name

	//ctl.Router.Delete("/:imageName", handleDeleteImage)

}
