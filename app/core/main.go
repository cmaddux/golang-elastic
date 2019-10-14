package main

import (
	"github.com/cmaddux/string_manipulation/elasticsearch"
	"github.com/cmaddux/string_manipulation/handlers"

	"github.com/gin-gonic/gin"
)

func setupRouter() *gin.Engine {
	r := gin.Default()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	r.Use(elasticsearch.Client())

	r.GET("/ok", handlers.OK())
	r.POST("/search", handlers.PostSearch())
	r.GET("/search/:value", handlers.GetSearchValue())
	r.POST("/special/count", handlers.GetSpecialCount())

	return r
}

func main() {
	r := setupRouter()
	r.Run(":8080")
}
