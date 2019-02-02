package main

import (
	"bytes"
	"net/http"
	"strconv"

	"github.com/elguapo1611/karaoke/src/helpers"
	"github.com/elguapo1611/karaoke/src/playlist"
	"github.com/elguapo1611/karaoke/src/song"
	"github.com/gin-gonic/gin"
	"github.com/microcosm-cc/bluemonday"
)

func handleWebsocket(c *gin.Context) {
	serveWs(hub, c.Writer, c.Request)
}

// Used for rendering the html when the app is compiled
func renderApp(c *gin.Context) {

}

type searchParams struct {
	Term string `json:"term" binding:"required"`
}

// curl -i -X  POST http://localhost:4000/api/search \
//   -H "Accept: application/json" -H "Content-Type: application/json" \
//   -d '{ "term": "radio"}'
func search(c *gin.Context) {
	var params searchParams
	c.BindJSON(&params)

	p := bluemonday.StrictPolicy()
	term := p.Sanitize(params.Term)
	songs := song.Search(term)
	c.JSON(http.StatusOK, songs)
}

type changeOrderParams struct {
	PlaylistSongID int `json:"playlist_song_id" binding:"required"`
	SortOrder      int `json:"sort_order" binding:"required"`
}

// curl -i -X  POST http://localhost:4000/api/playlist/change_order \
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

// curl -i -X  DELETE http://localhost:4000/api/playlist/song \
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

// curl -i -X PUT http://localhost:4000/api/playlist/song \
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
	hub.broadcast <- []byte(msg)
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
