package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func run() ([]string, error) {

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
			fmt.Println("path: ", file)
			fmt.Println("source: ", source)
			fmt.Println("name: ", name)
			fmt.Println("artist: ", artist)
			fmt.Println("language: ", language)
			fmt.Println(songAndMeta)
			fmt.Println()
		}
	}

	return fileList, nil
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