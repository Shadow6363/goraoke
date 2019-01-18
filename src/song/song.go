package song

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/elguapo1611/karaoke/src/helpers"
)

// DB used to pass sqlite db object
var DB *sql.DB

// InitSong Set the db variable as a pointer
func Init(db *sql.DB) {
	fmt.Println("initializing song")
	DB = db
}

// Song used as main struct for song data
type Song struct {
	ID     int
	UUID   string
	Name   string
	Artist string

	Source   string
	Language string
	Filename string
	Enabled  bool
	Keywords string
	// DurationInSeconds sql.NullInt64
	CreatedAt time.Time
	UpdatedAt time.Time
}

// Search returns all songs that match a given term
func Search(term string) []Song {
	var songs []Song
	rows, err := DB.Query(`
    SELECT id, name, artist, source, language, filename, uuid
    FROM songs
    WHERE id IN (
      SELECT rowid
      FROM songs_search
      WHERE songs_search MATCH ? ORDER BY bm25(songs_search)
	)
	LIMIT 50
  `, term)
	helpers.CheckErr(err)

	var name, artist, source, language, filename, uuid string
	var id int
	for rows.Next() {
		err = rows.Scan(&id, &name, &artist, &source, &language, &filename, &uuid)
		helpers.CheckErr(err)
		songs = append(songs, Song{ID: id, UUID: uuid, Name: name, Source: source, Artist: artist, Language: language, Filename: filename})
	}
	return songs
}
