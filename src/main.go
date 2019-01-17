package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/elguapo1611/karaoke/src/model"
	"github.com/gorilla/mux"

	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

var db *sql.DB
var err error

// our main function
func main() {
	db, err = sql.Open("sqlite3", "./db/karaoke.db")

	model.InitSong(db)

	if err != nil {
		fmt.Println(err)
	}
	defer db.Close()

	router := mux.NewRouter()
	router.HandleFunc("/songs.json", GetSongs).Methods("POST")
	router.HandleFunc("/songs.json", GetSongs).Methods("GET")
	log.Fatal(http.ListenAndServe("localhost:8000", router))
}
func GetSongs(w http.ResponseWriter, r *http.Request) {
	songs := model.Search("radio")
	fmt.Println(songs)
	w.Header().Set("Content-Type", "application/json")
	rows, err := db.Query("SELECT * FROM songs LIMIT 10")
	fmt.Println(rows)
	if err != nil {
		fmt.Println(err)
	}

	json, err := json.Marshal(songs)
	if err != nil {
		fmt.Println(err)
	}

	// fmt.Println(json)
	w.Write(json)
}
