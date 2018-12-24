# install dependencies
# mac os only
brew install ffmpeg

# create sqlite tables
go get github.com/codegangsta/gin
go get github.com/giorgisio/goav

# not sure if we need this yet
go get github.com/jinzhu/gorm


# run boostrap

sqlite3 ./db/karaoke.db < ./db/migrate/schema.sql
go run ./db/seed.go