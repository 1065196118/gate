package main

import (
	"github.com/gin-gonic/gin"
	"github.com/social-mediam-users/config"
	"github.com/social-mediam-users/middleware"
	"github.com/social-mediam-users/routes"
)

var DB = config.SetupDatabase()

func main() {
	server := gin.Default()

	server.Use(middleware.CORSMiddleware())

	// Users Route
	routes.SetupUserRoute(server)

	server.Run(":6060")
}
