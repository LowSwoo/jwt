package main

import (
	"github.com/gin-gonic/gin"
	"github.com/lowswoo/jwt/handlers"
)

var router *gin.Engine

func main() {
	router = gin.Default()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())
	router.Static("/templates/src/", "./templates/src")
	router.LoadHTMLGlob("templates/*.html")
	initializeRoutes()
	router.Run()
}

func initializeRoutes() {
	router.GET("/", handlers.ShowIndexPage)
	router.GET("/article/view/:article_id", handlers.GetArticle)
}
