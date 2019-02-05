import fetch from 'isomorphic-fetch';
import {
  IS_SEARCHING,
  SEARCHES_RECEIVED,
  PLAYLIST_RECEIVED,
  PLAYLIST_IS_FETCHING,
  PLAYLIST_SONG_ADDED,
  PLAYLIST_UPDATED,
  PLAYLIST_CLEARED
} from '../constants/reducer-actions.const';

// handle results of a generic response and dispatch
function handleGenericResponse(dispatch, response, func) {
  if (response.ok) {
    response.json().then(function(body) {
      dispatch(func(body));
    });
  } else {
    dispatch(authenticationFailed(response));
  }
}

function isFetching() {
  return {type: PLAYLIST_IS_FETCHING};
}


// post api/search
export function search(term) {

}


function playlistReceived(response) {
  return Object.assign({}, {
    type: PLAYLIST_RECEIVED
  }, response);
}

// get api/playlist
export function getPlaylist() {
  return (dispatch) => {
    dispatch(isFetching());
    fetch('/api/playlist')
      .then( function(response) {
        handleGenericResponse(dispatch, response, playlistReceived);
      });
  };
}

// put api/playlist/song
export function addPlaylistSong(songId) {

}

// delete api/playlist/song
// playlist_song_id": 1 
export function deletePlaylistSong(playlistSongId) {

}

// delete api/playlist/reset
export function resetPlaylist() {

}

// post api/playlist/change_order
// playlist_song_id": 6, "sort_order": 3 
export function playlistChangeOrder(playlistSongId, sortOrder) {

}





function addedSong(response) {
  return Object.assign({}, {
    type: ADDED_SONG
  }, response);
}

export function addSong(songId) {
  return (dispatch) => {
    /*jshint camelcase: false */
    dispatch(isFetching());
    fetch('/api/rooms/add_song',
      requestOptions({
        method: 'POST',
        body: JSON.stringify({
          song_id: songId
        })
      }))
      .then( function(response) {
        handleGenericResponse(dispatch, response, addedSong);
      });
  };
}