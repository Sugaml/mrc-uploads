package main

import (
	"github.com/sugaml/upload-api/internal/controller"
)

func main() {
	ctl := controller.NewController()
	ctl.Run()
}
