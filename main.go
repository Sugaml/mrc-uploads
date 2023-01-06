package main

import (
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
	"github.com/sugaml/mrc-uploads/internal/controller"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		logrus.Fatalf("error in load env file %v", err)
	} else {
		logrus.Info("Successfully loaded env file.")
	}
	ctl := controller.NewController()
	ctl.Run()
}
