package main

import (
	"bytes"
	"database/sql"
	"fmt"
	"net/http"
	"strconv"

	"github.com/elguapo1611/karaoke/src/helpers"
	"github.com/elguapo1611/karaoke/src/playlist"
	"github.com/elguapo1611/karaoke/src/song"
	"github.com/gin-gonic/gin"
	_ "github.com/mattn/go-sqlite3"
	"github.com/microcosm-cc/bluemonday"
)

var db *sql.DB
var err error
var hub *Hub

// our main function
func main() {
	db, err = sql.Open("sqlite3", "./db/karaoke.db")
	hub = newHub()
	go hub.run()

	song.Init(db)
	playlist.Init(db)

	helpers.CheckErr(err)
	if err != nil {
		fmt.Println(err)
	}
	defer db.Close()

	router := gin.Default()

	// set html template directory
	router.LoadHTMLGlob("templates/*")

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

	// websocket connection
	router.GET("/ws", handleWebsocket)

	// page for adding songs and updating the playlist
	router.GET("/remote_control", renderRemoteControl)

	// page that plays the karaoke tracks
	router.GET("/", renderKaraokeRoom)

	router.Run("localhost:3001")

}
func handleWebsocket(c *gin.Context) {
	serveWs(hub, c.Writer, c.Request)
}

func renderRemoteControl(c *gin.Context) {
}

func renderKaraokeRoom(c *gin.Context) {
	fmt.Println("loading karaoke room")
	c.HTML(http.StatusOK, "home.tmpl", gin.H{
		"title": "Karaoke",
	})
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

// curl -i -X  POST http://localhost:3000/playlist/change_order \
//   -H "Accept: application/json" -H "Content-Type: application/json" \
//   -d '{ "playlist_song_id": 6, "sort_order": 3 }'
func changeOrder(c *gin.Context) {
	go publishUpdate("orderChanged")
	var params changeOrderParams
	c.BindJSON(&params)
	c.JSON(http.StatusOK, playlist.ChangeOrder(params.PlaylistSongID, params.SortOrder))
}

type deletePlaylistSongParams struct {
	PlaylistSongID int `json:"playlist_song_id" binding:"required"`
}

// curl -i -X  DELETE http://localhost:3000/playlist/song \
//   -H "Accept: application/json" -H "Content-Type: application/json" \
//   -d '{ "playlist_song_id": 1 }'
func deletePlaylistSong(c *gin.Context) {
	go publishUpdate("songRemoved")
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
//   -d '{ "song_id": 1000 }'
func addPlaylistSong(c *gin.Context) {
	var params addPlaylistSongParams

	c.BindJSON(&params)
	playlist.AddSong(params.SongID)
	var buffer bytes.Buffer
	buffer.WriteString("songAdded: ")
	buffer.WriteString(strconv.Itoa(params.SongID))

	go publishUpdate(buffer.String())
	c.JSON(http.StatusOK, helpers.NewOK())
}

// pushes a text message to all clients
func publishUpdate(msg string) {
	h := hub
	for client := range h.clients {
		fmt.Println("update ws client")
		client.hub.broadcast <- []byte(msg)
	}
}

func reset(c *gin.Context) {
	go publishUpdate("playlistReset")
	playlist.Reset()
	c.JSON(http.StatusOK, helpers.NewOK())
}

func getPlaylist(c *gin.Context) {
	playlist := playlist.GetPlaylist(10)
	c.JSON(http.StatusOK, playlist)
}
