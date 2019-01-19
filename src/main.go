package main

import (
	"database/sql"
	"fmt"
	"net/http"

	"github.com/elguapo1611/karaoke/src/helpers"
	"github.com/elguapo1611/karaoke/src/playlist"
	"github.com/elguapo1611/karaoke/src/song"
	"github.com/gin-gonic/gin"
	_ "github.com/mattn/go-sqlite3"
	"github.com/microcosm-cc/bluemonday"
)

var db *sql.DB
var err error

// our main function
func main() {
	db, err = sql.Open("sqlite3", "./db/karaoke.db")

	song.Init(db)
	playlist.Init(db)

	helpers.CheckErr(err)
	if err != nil {
		fmt.Println(err)
	}
	defer db.Close()

	router := gin.Default()

	// search for songs
	router.GET("/search/:term", search)

	// get all songs in the playlist
	router.GET("/playlist", getPlaylist)

	// add a song to the playlist
	router.PUT("/playlist/song/:song_id", getPlaylist)
	// update a song order within the playlist
	router.POST("/playlist/song/:playlist_song_id", getPlaylist)
	// remove a song from the playlist
	router.DELETE("/playlist/song/:playlist_song_id", getPlaylist)
	// reset the playlist and clear all songs
	router.DELETE("/playlist/reset", reset)
	// change order of playlist
	router.GET("/playlist/change_order", changeOrder)

	// page for adding songs and updating the playlist
	router.GET("/remote_control", renderRemoteControl)

	// page that plays the karaoke tracks
	router.GET("/karaoke_room", renderKaraokeRoom)

	router.Run("localhost:3001")

}

func changeOrder(c *gin.Context) {
	c.JSON(http.StatusOK, playlist.ChangeOrder(6, 3))
}
func reset(c *gin.Context) {
	playlist.Reset()
	c.JSON(http.StatusOK, playlist.OK{OK: true})
}

func renderRemoteControl(c *gin.Context) {
}

func renderKaraokeRoom(c *gin.Context) {
}

func search(c *gin.Context) {
	p := bluemonday.StrictPolicy()
	term := p.Sanitize(c.Param("term"))
	songs := song.Search(term)
	c.JSON(http.StatusOK, songs)
}

func getPlaylist(c *gin.Context) {
	playlist := playlist.GetPlaylist(10)
	c.JSON(http.StatusOK, playlist)
}
