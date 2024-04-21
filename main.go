package main

import (
	"github.com/MustaphaSakka/traney/app"
	"github.com/MustaphaSakka/traney/logger"
)

func main() {
	logger.Info("Application is started!")
	app.Start()

}
