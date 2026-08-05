package main

import (
	"log"

	"rumah-restaurant/config"
	"rumah-restaurant/routes"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {

	err := godotenv.Load()

	if err != nil {
		log.Fatal(err)
	}

	config.ConnectDatabase()

	router := gin.Default()

	router.Static("/uploads", "./uploads")

	routes.SetupRoutes(router)

	router.Run(":8080")
}
