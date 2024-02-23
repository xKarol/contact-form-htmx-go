package main

import (
	"app/internal/handlers"

	"app/internal/utils"

	"github.com/gin-gonic/gin"
)

func main() {
	utils.LoadEnv()
	router := gin.Default()

	router.Static("/static", "./dist")

	handlers.Init(router)

	router.Run()
}
