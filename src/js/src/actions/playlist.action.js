import fetch from 'isomorphic-fetch';
import {
  IS_SEARCHING,
  SEARCH_RESULTS_RECEIVED,
  PLAYLIST_RECEIVED,
  PLAYLIST_IS_FETCHING,
  PLAYLIST_SONG_ADDED,
  PLAYLIST_UPDATED,
  PLAYLIST_CLEARED,
  PLAYLIST_SONG_REMOVED
} from '../constants/reducer-actions.const';

// set json content type headers
export const requestOptions = function(options) {
  return Object.assign({}, options, {
    headers: {
      'Accept': 'application/json',
      'Content-Type': 'application/json',
    },
    credentials: 'same-origin'
  });
};

// handle results of a generic response and dispatch
function handleGenericResponse(dispatch, response, func) {
  if (response.ok) {
    response.json().then(function(body) {
      dispatch(func(body));
    });
  } else {
    dispatch(errorOccurred(response));
  }
}

function errorOccurred(response) {

}

function isFetching() {
  return {type: PLAYLIST_IS_FETCHING};
}

function isSearching() {
  return {type: IS_SEARCHING}
}

function searchResultsReceived(response) {
  return Object.assign({}, {
    type: SEARCH_RESULTS_RECEIVED,
    response: response
  })
}

// post api/search
export function search(term) {
  return (dispatch) => {
    dispatch(isSearching());
    fetch('api/playlist/song',
      requestOptions({
        method: 'POST',
        body: JSON.stringify({
          term: term
        })
      }))
      .then( function(response) {
        handleGenericResponse(dispatch, response, searchResultsReceived);
      });
  }
}


function playlistReceived(response) {
  return Object.assign({}, {
    type: PLAYLIST_RECEIVED
  }, {playlistSongs: response});
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

function songAdded(response) {
  return Object.assign({}, {
    type: PLAYLIST_SONG_ADDED
  })
}
// put api/playlist/song
export function addPlaylistSong(songId) {
  return (dispatch) => {
    fetch('api/playlist/song',
      requestOptions({
        method: 'PUT',
        body: JSON.stringify({
          song_id: songId
        })
      }))
      .then( function(response) {
        handleGenericResponse(dispatch, response, songAdded);
      });
  }
}

function playlistSongRemoved(response) {
  return Object.assign({}, {
    type: PLAYLIST_SONG_REMOVED
  }, response )
}
// delete api/playlist/song
// playlist_song_id": 1 
export function removePlaylistSong(playlistSongId) {
  return (dispatch) => {
    fetch('api/playlist/song',
      requestOptions({
        method: 'DELETE',
        body: JSON.stringify({
          playlist_song_id: playlistSongId
        })
      }))
      .then( function(response) {
        handleGenericResponse(dispatch, response, playlistSongRemoved);
      });
  }
}

function playlistReset(response) {
  return Object.assign({}, {
    type: PLAYLIST_CLEARED
  }, response)
}

// delete api/playlist/reset
export function resetPlaylist() {
  return (dispatch) => {
    fetch('api/playlist/reset',
      requestOptions({
        method: 'DELETE'
      }))
      .then( function(response) {
        handleGenericResponse(dispatch, response, playlistReset);
      });
  }
}

function playlistOrderChanged(response) {
  return Object.assign({}, {
    type: PLAYLIST_UPDATED
  }, response)
}

// post api/playlist/change_order
// playlist_song_id": 6, "sort_order": 3 
export function playlistChangeOrder(playlistSongId, sortOrder) {
  return (dispatch) => {
    fetch('api/playlist/reset',
      requestOptions({
        method: 'POST',
        body: JSON.stringify({
          playlist_song_id: playlistSongId,
          sortOrder: sortOrder
        })
      }))
      .then( function(response) {
        handleGenericResponse(dispatch, response, playlistOrderChanged);
      });
  }
}
