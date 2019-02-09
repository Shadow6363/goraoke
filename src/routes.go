package main

import (
	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"
)

func loadRoutes() {

	router := gin.Default()

	// middleware for serving static assets
	router.Use(static.Serve("/", static.LocalFile("./public", true)))

	// Fix this for the staticly compiled assets
	router.GET("/public/:path")
	// set html template directory
	// router.LoadHTMLGlob("templates/*")

	// search for songs
	router.POST("/api/search", search)
	router.GET("/api/search", search)

	// get all songs in the playlist
	router.GET("/api/playlist", getPlaylist)

	// add a song to the playlist
	router.PUT("/api/playlist/song", addPlaylistSong)
	// remove a song from the playlist
	router.DELETE("/api/playlist/song", deletePlaylistSong)
	// reset the playlist and clear all songs
	router.DELETE("/api/playlist/reset", reset)
	// change order of playlist
	router.POST("/api/playlist/change_order", changeOrder)

	// websocket connection
	router.GET("/api/ws", handleWebsocket)

	// page for adding songs and updating the playlist
	router.GET("/", renderApp)

	// Gin proxies from port 3000 to 3001 during development for hot reloading
	router.Run("localhost:3001")
}
