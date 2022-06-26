package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"

	"github.com/azratul/silens-clamor-backend/config"
	"github.com/azratul/silens-clamor-backend/lastfm"

	"github.com/gin-gonic/gin"
)

func GetAlbumArt(c *gin.Context) {
	if c.Query("artist") == "" || c.Query("song") == "" {
		c.JSON(http.StatusBadRequest, "Artist & Song are mandatory!")
		return
	}

	artist := url.QueryEscape(c.Query("artist"))
	song := url.QueryEscape(c.Query("song"))
	attr := url.QueryEscape(c.Query("attr"))

	url := fmt.Sprintf("%s/?method=track.getInfo&api_key=%s&artist=%s&track=%s&format=json", config.Conf.LastFM.Url, config.Conf.LastFM.ApiKey, artist, song)

	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)

	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}

	resp, err := client.Do(req)

	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	var info lastfm.Info

	err = json.NewDecoder(resp.Body).Decode(&info)

	defer resp.Body.Close()
	req.Close = true

	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	if attr == "image" {
		c.JSON(http.StatusOK, info.Track.Album.Image)
		return
	}

	c.JSON(http.StatusOK, info)
}
