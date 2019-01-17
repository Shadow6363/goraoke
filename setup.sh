# install dependencies

# create sqlite tables


go get github.com/kardianos/govendor

go get github.com/codegangsta/gin

# not sure if we need this yet



# run boostrap

sqlite3 ./src/db/karaoke.db < ./db/migrate/schema.sql
go run --tags "fts5" ./src/db/seed.go
