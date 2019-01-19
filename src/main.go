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
	router.PUT("/playlist/song", addPlaylistSong)
	// remove a song from the playlist
	router.DELETE("/playlist/song", deletePlaylistSong)
	// reset the playlist and clear all songs
	router.DELETE("/playlist/reset", reset)
	// change order of playlist
	router.POST("/playlist/change_order", changeOrder)

	// page for adding songs and updating the playlist
	router.GET("/remote_control", renderRemoteControl)

	// page that plays the karaoke tracks
	router.GET("/karaoke_room", renderKaraokeRoom)

	router.Run("localhost:3001")

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

type changeOrderParams struct {
	PlaylistSongID int `json:"playlist_song_id" binding:"required"`
	SortOrder      int `json:"sort_order" binding:"required"`
}

// curl -i -X  DELETE http://localhost:3000/playlist/change_order \
//   -H "Accept: application/json" -H "Content-Type: application/json" \
//   -d '{ "playlist_song_id": 6, "sort_order": 3 }'
func changeOrder(c *gin.Context) {
	var params changeOrderParams
	c.BindJSON(&params)
	c.JSON(http.StatusOK, playlist.ChangeOrder(params.PlaylistSongID, params.SortOrder))
}

type deletePlaylistSongParams struct {
	PlaylistSongID int `json:"playlist_song_id" binding:"required"`
}

// curl -i -X  DELETE http://localhost:3000/playlist/song \
//   -H "Accept: application/json" -H "Content-Type: application/json" \
//   -d '{ "playlist_song_id": 6 }'
func deletePlaylistSong(c *gin.Context) {
	var params deletePlaylistSongParams
	c.BindJSON(&params)
	playlist.DeletePlaylistSong(params.PlaylistSongID)
	c.JSON(http.StatusOK, helpers.NewOK())
}

type addPlaylistSongParams struct {
	SongID int `json:"song_id" binding:"required"`
}

// curl -i -X PUT http://localhost:3000/playlist/song \
//   -H "Accept: application/json" -H "Content-Type: application/json" \
//   -d '{ "song_id": 6 }'
func addPlaylistSong(c *gin.Context) {
	var params addPlaylistSongParams
	c.BindJSON(&params)
	playlist.AddSong(params.SongID)
	c.JSON(http.StatusOK, helpers.NewOK())
}

func reset(c *gin.Context) {
	playlist.Reset()
	c.JSON(http.StatusOK, helpers.NewOK())
}

func getPlaylist(c *gin.Context) {
	playlist := playlist.GetPlaylist(10)
	c.JSON(http.StatusOK, playlist)
}
