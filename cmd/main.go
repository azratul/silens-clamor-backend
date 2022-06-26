package main

import (
	"log"
	"os"

	"github.com/azratul/silens-clamor-backend/api"
	"github.com/azratul/silens-clamor-backend/config"

	"github.com/gin-gonic/gin"
)

func main() {
	if os.Getenv("LASTFM_APIKEY") == "" || os.Getenv("LASTFM_URL") == "" {
		log.Println("All env vars must be set!!!")
		return
	}

	config.Conf.LastFM.ApiKey = os.Getenv("LASTFM_APIKEY")
	config.Conf.LastFM.Url = os.Getenv("LASTFM_URL")

	r := gin.Default()
	prefix := "/api/v1"
	v1 := r.Group(prefix)

	v1.GET("/info", api.GetAlbumArt)
	r.Run(":8010")
}
