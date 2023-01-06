package main

import (
	"github.com/sugaml/mrc-uploads/internal/controller"
)

func main() {
	ctl := controller.NewController()
	ctl.Run()
}
