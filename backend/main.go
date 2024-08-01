package main

import (
	"os"

	"github.com/ShahSau/CareFlow/backend/database"
	"github.com/ShahSau/CareFlow/backend/router"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	// databsace connection
	database.ConnectDB()

	// router
	router.ClientRoutes()
}
