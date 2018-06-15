package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gin-contrib/static"
	"github.com/gin-gonic/contrib/gzip"
	"github.com/gin-gonic/gin"
)

func main() {

	langMap, err := LoadAllLanguageFiles(10)

	if err != nil {

		router := gin.Default()
		router.Use(gzip.Gzip(gzip.DefaultCompression))

		router.Use(gin.Recovery())

		router.Use(static.Serve("/", static.LocalFile("/public", true)))
		router.Use(static.Serve("/public", static.LocalFile("/public", true)))

		api := router.Group("/api")

		api.GET("/ping", func(c *gin.Context) {
			c.String(http.StatusOK, "pong: "+os.Getenv("VERSION"))
		})

		api.POST("/possiblewords", func(c *gin.Context) {
			boggleChars := BoggleChars{}

			if c.ShouldBindJSON(&boggleChars) == nil {
				mapped := ConvertToMapped(boggleChars)
				words, err := GetAllValidWords(langMap[boggleChars.Lang], mapped, 5)
				if err == nil {
					c.JSON(http.StatusOK, words)
				} else {
					log.Println(err.Error())
					c.JSON(http.StatusInternalServerError, gin.H{})
				}
			} else {
				log.Println(err.Error())
				c.JSON(http.StatusInternalServerError, gin.H{})
			}

		})

		router.Run(":" + os.Getenv("PORT"))

	}

}
