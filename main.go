package main

import (
	"github.com/MustaphaSakka/traney-lib/logger"
	"github.com/MustaphaSakka/traney/app"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()
	logger.Info("TRANEY is started!")
	app.Start()

}
