package controller

func (ctl *Controller) InitializeRoutes() {
	//ctl.Router.Use(cors.New())

	ctl.Router.Static("/uploads", "./uploads")
	routes := ctl.Router.Group("/uploads")
	routes.Post("/", ctl.handleFileupload)
	routes.Post("/multiple", ctl.handleMultipleFileupload)
	routes.Delete("/:imageName", ctl.handleDeleteImage)
}
