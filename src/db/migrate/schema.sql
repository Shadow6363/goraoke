DROP TABLE IF EXISTS playlist_songs;
DROP TABLE IF EXISTS songs;
DROP TABLE IF EXISTS users;
DROP TABLE IF EXISTS songs_search;

CREATE TABLE users (
  id INTEGER PRIMARY KEY,
  email,
  name NOT NULL,
  avatar,
  created_at DATETIME DEFAULT CURRENT_TIMESTAMP NOT NULL,
  updated_at DATETIME DEFAULT CURRENT_TIMESTAMP NOT NULL
);

CREATE TABLE songs (
  id INTEGER PRIMARY KEY,
  uuid NOT NULL,
  name NOT NULL,
  artist NOT NULL,
  source,
  language,
  filename NOT NULL,
  enabled DEFAULT TRUE,
  keywords,
  duration_in_seconds INTEGER,
  created_at DATETIME DEFAULT CURRENT_TIMESTAMP NOT NULL,
  updated_at DATETIME DEFAULT CURRENT_TIMESTAMP NOT NULL
);

CREATE UNIQUE INDEX index_songs_on_uuid ON songs (uuid);

-- create search index on songs
CREATE VIRTUAL TABLE songs_search USING FTS5(
  name,
  artist,
  keywords,
  content=songs,
  content_rowid=id,
  tokenize = 'porter ascii'
);

-- keep songs search index up to date
-- after an insert
CREATE TRIGGER songs_after_insert AFTER INSERT ON songs BEGIN
  INSERT INTO songs_search(rowid, name, artist, keywords) VALUES (new.id, new.name, new.artist, new.keywords);
END;
-- after a delete
CREATE TRIGGER songs_after_delete AFTER DELETE ON songs BEGIN
  INSERT INTO songs_search(songs_search, rowid, name, artist, keywords) VALUES('delete', old.id, old.name, old.artist, old.keywords);
END;
-- after an update
CREATE TRIGGER songs_after_update AFTER UPDATE ON songs BEGIN
  INSERT INTO songs_search(songs_search, rowid, name, artist, keywords) VALUES('delete', old.id, old.name, old.artist, old.keywords);
  INSERT INTO songs_search(rowid, name, artist, keywords) VALUES (new.id, new.name, new.artist, new.keywords);
END;

-- INSERT INTO songs(name, artist, source, language, filename, duration_in_seconds, keywords, uuid)
--   VALUES("Under the sea", "disney", "songfly", "en", "foo.mp3", 200, "foo bar fizz buzz", "foo");

SELECT *
FROM songs
WHERE id IN (SELECT rowid
             FROM songs_search
             WHERE songs_search MATCH 'video'
             ORDER BY bm25(songs_search)
             );


DROP TABLE IF EXISTS playlist_songs;

CREATE TABLE playlist_songs (
  id INTEGER PRIMARY KEY,
  song_id INT NOT NULL,
  user_id INT NOT NULL,
  sort_order INT,
  created_at DATETIME DEFAULT CURRENT_TIMESTAMP NOT NULL,
  updated_at DATETIME DEFAULT CURRENT_TIMESTAMP NOT NULL
);

CREATE INDEX index_playlist_songs_on_room_id_and_song_id_and_sort_order ON playlist_songs (song_id, sort_order);



INSERT INTO playlist_songs(song_id, user_id, sort_order)
   VALUES(1, 1, 1);
INSERT INTO playlist_songs(song_id, user_id, sort_order)
   VALUES(2, 1, 2);
INSERT INTO playlist_songs(song_id, user_id, sort_order)
   VALUES(3, 1, 3);
INSERT INTO playlist_songs(song_id, user_id, sort_order)
   VALUES(4, 1, 4);
INSERT INTO playlist_songs(song_id, user_id, sort_order)
   VALUES(5, 1, 5);
INSERT INTO playlist_songs(song_id, user_id, sort_order)
   VALUES(6, 1, 6);
INSERT INTO playlist_songs(song_id, user_id, sort_order)
   VALUES(7, 1, 7);
INSERT INTO playlist_songs(song_id, user_id, sort_order)
   VALUES(8, 1, 8);
INSERT INTO playlist_songs(song_id, user_id, sort_order)
   VALUES(9, 1, 9);


-- order was 2 now it's 5
UPDATE playlist_songs 
  SET sort_order = sort_order - 1
  WHERE sort_order >= 2 and sort_order <= 5;

UPDATE playlist_songs
  SET order = 5
  WHERE id = 1