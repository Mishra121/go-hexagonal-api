package main

import (
	"banking/app"
	"banking/logger"
)

func main() {

	logger.Info("Starting the application...")
	app.Start()
}
