package main

import (
	"database/sql"
	"os"

	"github.com/elguapo1611/karaoke/src/helpers"
	"github.com/elguapo1611/karaoke/src/playlist"
	"github.com/elguapo1611/karaoke/src/song"
	_ "github.com/mattn/go-sqlite3"
)

var db *sql.DB
var err error
var hub *Hub
var env string

// our main function
func main() {
	env = os.Getenv("ENV")
	db, err = sql.Open("sqlite3", "./db/karaoke.db")
	hub = newHub()

	go hub.run()

	song.Init(db)
	playlist.Init(db)

	helpers.CheckErr(err)

	defer db.Close()

	loadRoutes()
}
