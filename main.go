package main

import (
	"github.com/MustaphaSakka/traney-lib/logger"
	"github.com/MustaphaSakka/traney/app"
)

func main() {
	logger.Info("TRANEY is started!")
	app.Start()

}
