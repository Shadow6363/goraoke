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

// ChangeOrder changes the order of the playlist
func ChangeOrder(playlistSongID int, sortOrder int) []PlaylistSong {
	playlistSong := getPlaylistSong(playlistSongID)

	var sqlStatment string
	// sort.Ints(sortOrders)
	if sortOrder > playlistSong.SortOrder {
		sqlStatment = `
            UPDATE playlist_songs 
            SET sort_order = sort_order - 1
            WHERE sort_order >= ? and sort_order <= ?;
        `
	} else {
		sqlStatment = `
            UPDATE playlist_songs 
            SET sort_order = sort_order + 1
            WHERE sort_order <= ? AND sort_order >= ?;
        `
	}
	_, err := DB.Exec(sqlStatment, playlistSong.SortOrder, sortOrder)
	helpers.CheckErr(err)

	sqlStatment = `UPDATE playlist_songs
        SET sort_order = ?
        WHERE id = ?
    `
	_, err = DB.Exec(sqlStatment, sortOrder, playlistSongID)
	helpers.CheckErr(err)

	return GetPlaylist(100)
}

// Reset Removes all songs in the playlist
func Reset() {
	// delete
	sqlStatment := "delete from playlist_songs"
	fmt.Println("deleting all playlist songs")
	_, err := DB.Exec(sqlStatment)
	helpers.CheckErr(err)
}

// DeletePlaylistSong removes a song from the playlist
func DeletePlaylistSong(playlistSongID int) {
	sqlStatment := `
            DELETE FROM playlist_songs 
            WHERE id = ?
    `
	_, err := DB.Exec(sqlStatment, playlistSongID)
	helpers.CheckErr(err)
}

// GetPlaylist Returns all songs in the playlist
func GetPlaylist(limit int) []PlaylistSong {
	fmt.Println("called get playlist")
	// Make sure to return an empty array when marshalling.  []PlaylistSong{} correctly translates to an empty array when marshalled.
	playlistSongs := []PlaylistSong{}
	rows, err := DB.Query(`
        SELECT ps.id as playlistSongID, ps.song_id as songID, ps.user_id as userID, ps.sort_order as sortOrder, s.name, s.artist, s.source, s.language, s.filename, s.uuid
        FROM playlist_songs ps
        JOIN songs s
            ON ps.song_id = s.id
        ORDER BY sortOrder ASC
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

	rows.Close()
	fmt.Println(playlistSongs)
	return playlistSongs
}

// AddSong adds a song to the playlist at the last position
func AddSong(songID int) {

	// get max sort order
	sqlStatement := `
      SELECT sort_order as maxSortOrder
      FROM playlist_songs
      ORDER BY sort_order DESC
      LIMIT 1
    `
	rows, err := DB.Query(sqlStatement)
	helpers.CheckErr(err)

	var maxSortOrder, nextSortOrder int
	for rows.Next() {
		err = rows.Scan(&maxSortOrder)
		helpers.CheckErr(err)
		nextSortOrder = maxSortOrder + 1
	}
	rows.Close()

	sqlStatement = `
        INSERT INTO playlist_songs(song_id, user_id, sort_order)
            VALUES(?, 1, ?);
    `
	_, err = DB.Exec(sqlStatement, songID, nextSortOrder)

}

func getPlaylistSong(playlistSongID int) PlaylistSong {
	var playlistSongs []PlaylistSong
	rows, err := DB.Query(`
    SELECT song_id as SongID, user_id as userID, sort_order as sortOrder
    FROM playlist_songs
    WHERE id = ?
  `, playlistSongID)
	helpers.CheckErr(err)

	var sortOrder, songID, userID int
	for rows.Next() {
		err = rows.Scan(&songID, &userID, &sortOrder)
		helpers.CheckErr(err)
		playlistSongs = append(playlistSongs, PlaylistSong{
			ID:        playlistSongID,
			SongID:    songID,
			UserID:    userID,
			SortOrder: sortOrder,
		})
	}

	rows.Close()

	return playlistSongs[0]
}
