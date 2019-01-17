package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"crypto/md5"
	"encoding/hex"
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
)

func run() ([]string, error) {
	db, err := sql.Open("sqlite3", "./db/karaoke.db")
	checkErr(err)

	searchDir := getPath()

	fileList := make([]string, 0)
	e := filepath.Walk(searchDir, func(path string, f os.FileInfo, err error) error {
		fileList = append(fileList, path)
		return err
	})
	
	if e != nil {
		panic(e)
	}

	for _, file := range fileList {
		if filepath.Ext(file) == ".mp3" {
			filename := file
			hasher := md5.New()
    	hasher.Write([]byte(file))
    	uuid := hex.EncodeToString(hasher.Sum(nil))

			splitFilename := strings.Split(file, " - ")
			artistMeta := splitFilename[0]
			artistWithCrap := strings.Split(artistMeta, "/")
			artist := strings.TrimSpace(artistWithCrap[len(artistWithCrap)-1])
			songAndMeta := splitFilename[1]
			language := "english"
			name := strings.TrimSpace(strings.Split(songAndMeta, "[")[0])
			meta := GetStringInBetween(songAndMeta, "[", "]")

			var source string

			if meta == "Karaoke" {
				source = "unknown"
			} else {
				source = strings.ToLower(strings.TrimSpace(strings.Split(meta, " ")[0]))
				if source == "italian" || source == "spanish" || source == "german" || source == "french" {
					language = source
				}
			}
			fmt.Println("uuid: ", uuid)
			fmt.Println("filename: ", file)
			fmt.Println("source: ", source)
			fmt.Println("name: ", name)
			fmt.Println("artist: ", artist)
			fmt.Println("language: ", language)
			fmt.Println(songAndMeta)
			fmt.Println()

			// insert record
			stmt, err := db.Prepare("INSERT INTO songs(uuid, filename, source, name, artist, language) values(?,?,?,?,?,?)")
      checkErr(err)
      res, err := stmt.Exec(uuid, filename, source, name, artist, language)
      checkErr(err)
      id, err := res.LastInsertId()
      checkErr(err)
      fmt.Println(id)

		}
	}

	return fileList, nil
}

func checkErr(err error) {
    if err != nil {
        panic(err)
    }
}

func GetStringInBetween(str string, start string, end string) (result string) {
    s := strings.Index(str, start)
    if s == -1 {
      return
    }
    s += len(start)
    e := strings.Index(str, end)
    if e == -1 {
	    return "Karaoke"
    }

    fmt.Println(str)
    final := str[s:e]
    return final
}

func getPath() string {
	defaultSongPath := "/Volumes/external/songs"	
	var searchDir string
	searchDir = os.Getenv("SONG_PATH")
	if searchDir == "" {
	  searchDir = defaultSongPath
	}

	return searchDir
}

func main() {
	run()
}