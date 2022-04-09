package handlers

import (
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/lowswoo/jwt/models"
)

func ShowIndexPage(ctx *gin.Context) {
	render(ctx, gin.H{
		"title":   "Home page",
		"payload": models.GetAllArticles(),
	}, "index.html")
}

func GetArticle(ctx *gin.Context) {

	article_id, err := strconv.Atoi(ctx.Param("article_id"))
	if err != nil {
		ctx.AbortWithError(http.StatusNotFound, err)
	}
	article, err := models.GetArticleById(article_id)
	if err != nil {
		ctx.AbortWithError(http.StatusNotFound, err)
	}

	render(ctx, gin.H{
		"title":   article.Title,
		"payload": article.Content,
	}, "article.html")
}

func CreateArticle(ctx *gin.Context) {
	err := ctx.Request.ParseForm()
	if err != nil {
		log.Fatal("PIZDETS")
		ctx.AbortWithError(http.StatusNotFound, err)
	}
	models.CreateArticle(ctx.Request.FormValue("articleTitle"), ctx.Request.FormValue("articleContent"))
	ctx.Redirect(http.StatusFound, "/")
}
