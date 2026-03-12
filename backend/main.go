package main

import (
	"github.com/Ferhan0/project-management-system/initializers"
	"github.com/Ferhan0/project-management-system/routes"
	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDb()
	initializers.SyncToDb()
}

func main() {
	router := gin.Default()
	routes.SetupRoutes(router)
	router.Run()
}
