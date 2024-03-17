package router

import "github.com/gin-gonic/gin"

func Initialize() {

	// Initialize Server in application
	server := gin.Default()

	// Initialize Routes
	initializerRoutes(server)

	//run Server HTTP
	server.Run(":8089") //listening port 8089
}
