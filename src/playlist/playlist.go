package playlist

import (
	"database/sql"
	"fmt"

	"github.com/elguapo1611/karaoke/src/helpers"
	"github.com/elguapo1611/karaoke/src/song"
)

// DB used to pass sqlite db object
var DB *sql.DB

// Init Set the db variable as a pointer
func Init(db *sql.DB) {
	fmt.Println("initializing song")
	DB = db
}

// PlaylistSong used as main struct for song data
type PlaylistSong struct {
	ID        int
	SongID    int
	UserID    int
	SortOrder int
	Song      song.Song
}

// GetPlaylist Returns all songs in the playlist
func GetPlaylist(limit int) []PlaylistSong {
	var playlistSongs []PlaylistSong
	rows, err := DB.Query(`
		SELECT ps.id as playlistSongID, ps.song_id as songID, ps.user_id as userID, ps.sort_order as sortOrder, s.name, s.artist, s.source, s.language, s.filename, s.uuid
		FROM playlist_songs ps
		JOIN songs s
			ON ps.song_id = s.id
		ORDER BY sortOrder DESC, playlistSongID ASC
		LIMIT ?
	`, limit)
	helpers.CheckErr(err)

	var name, artist, source, language, filename, uuid string
	var sortOrder, playlistSongID, songID, userID int
	for rows.Next() {
		err = rows.Scan(&playlistSongID, &songID, &userID, &sortOrder, &name, &artist, &source, &language, &filename, &uuid)
		helpers.CheckErr(err)
		playlistSongs = append(playlistSongs, PlaylistSong{
			ID:        playlistSongID,
			SongID:    songID,
			UserID:    userID,
			SortOrder: sortOrder,
			Song: song.Song{
				ID:       songID,
				UUID:     uuid,
				Name:     name,
				Source:   source,
				Artist:   artist,
				Language: language,
				Filename: filename,
			},
		})

	}

	return playlistSongs
}
