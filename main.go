package main

import (
	"errors"
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

		api.GET("/supportedlanguages", func(c *gin.Context) {
			supportedLanguages := SupportedLanguageResponse{
				SupportedLanguages: []SupportedLanguage{
					{
						Lang:     "en_US",
						Language: "English",
					}, {
						Lang:     "es_ES",
						Language: "Espanol",
					},
				},
			}

			c.JSON(http.StatusOK, supportedLanguages)

		})

		api.POST("/possiblewords", func(c *gin.Context) {
			boggleChars := BoggleChars{}

			if c.ShouldBindJSON(&boggleChars) == nil {
				// check that requested language is supported
				if _, ok := langMap[boggleChars.Lang]; ok {

					// check that boggle board size is supported
					if len(boggleChars.Rows) > 20 ||
						len(boggleChars.Rows[0].Cols) > 20 ||
						len(boggleChars.Rows) < 1 ||
						len(boggleChars.Rows[0].Cols) < 1 {
						err := errors.New("Board size is not supported")
						log.Println(err)
						c.Status(http.StatusBadRequest)
					} else {

						mapped := ConvertToMapped(boggleChars)
						words, err := GetAllValidWords(langMap[boggleChars.Lang], mapped, 5)
						if err == nil {
							c.JSON(http.StatusOK, words)
						} else {
							log.Println(err.Error())
							c.Status(http.StatusInternalServerError)
						}
					}
				} else {
					err := errors.New("Language not supported")
					log.Println(err)
					c.Status(http.StatusBadRequest)
				}
			} else {
				log.Println(err.Error())
				c.Status(http.StatusBadRequest)
			}

		})

		router.Run(":" + os.Getenv("PORT"))

	}

}
