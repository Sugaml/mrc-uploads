package main

import (
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
	"github.com/sugaml/mrc-uploads/internal/controller"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		logrus.Fatalf("error in load env variables  %v", err)
	}
	logrus.Info("Successfully loaded env variables.")

	ctl := controller.NewController()
	ctl.Run()
}
