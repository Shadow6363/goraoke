### Development

# install golang
# https://golang.org/doc/install

setup.sh

CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main .
docker build -t example-scratch -f Dockerfile.scratch .
docker run -it example-scratch

# curl commands
```
curl -i -X  POST http://localhost:3000/playlist/change_order \
   -H "Accept: application/json" -H "Content-Type: application/json" \
   -d '{ "playlist_song_id": 6, "sort_order": 3 }'
```
```
curl -i -X  DELETE http://localhost:3000/playlist/song \
   -H "Accept: application/json" -H "Content-Type: application/json" \
   -d '{ "playlist_song_id": 2 }'
```
```
curl -i -X PUT http://localhost:3000/playlist/song \
   -H "Accept: application/json" -H "Content-Type: application/json" \
   -d '{ "song_id": 1000 }'
```