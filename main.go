package main

import (
	"bb-server/src/router"
	u "bb-server/src/utils"
	"flag"
)

// Command-line flags.
var (
	httpAddr = flag.String("http", ":8080", "Listen address")
)

func main() {

	flag.Parse()

	u.InfoLogger.Println("Starting the application...")

	// Set gin mode.
	// gin.SetMode(viper.GetString("runmode"))

	// Routes.
	g := router.Load()

	// ListenAndServe
	u.ErrorLogger.Println(g.Run(*httpAddr).Error())

}
