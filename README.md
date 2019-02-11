### Development

# install golang
# https://golang.org/doc/install

setup.sh

# running locally each service independently
```
gin --buildArgs '--tags "sqlite_fts5"' run main.go
npm start
```

# running locally
```
bash start.sh
```

This application uses gin to autoreload.  The extra build tags include full text search as a dependency in the sqlite golang package.

# add code dependencies
This project uses govendor to vendor all dependencies locally.  https://github.com/kardianos/govendor

# curl commands
```
curl -i -X  POST http://localhost:3000/api/playlist/change_order \
   -H "Accept: application/json" -H "Content-Type: application/json" \
   -d '{ "playlist_song_id": 6, "sort_order": 3 }'
```
```
curl -i -X  DELETE http://localhost:3000/api/playlist/song \
   -H "Accept: application/json" -H "Content-Type: application/json" \
   -d '{ "playlist_song_id": 2 }'
```
```
curl -i -X PUT http://localhost:3000/api/playlist/song \
   -H "Accept: application/json" -H "Content-Type: application/json" \
   -d '{ "song_id": 1000 }'
```

# TODOS
### Data ingestion:
  - https://www.discogs.com/developers/#
### Create components:
  - PlaylistSongs
  - Song
  - Tabs (playlist, search)
  - CDG player
  - Next 3 songs
  - Toastr
  - Pause/Play, Fast Forward
  - Search/Autocomplete with debounce
### Actions
  - Add song
  - Remove Song
  - Move Song
  - Get playlist
  - Fast forward/skip
  - Pause/Play
### Reducers
  - Playlist Reducer
  - Search Reducer
  - Room Reducer
  - Toast Reducer
### Compile
  - Packr all js assets into binary
  - Test sqlite connectivity
  - Cross compile binary for raspberrypi
  