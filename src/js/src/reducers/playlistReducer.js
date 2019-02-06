import {
  PLAYLIST_RECEIVED,
  PLAYLIST_IS_FETCHING,
  PLAYLIST_SONG_ADDED,
  PLAYLIST_SONG_REMOVED,
  PLAYLIST_UPDATED,
  PLAYLIST_CLEARED
} from '../constants/reducer-actions.const';

const initialState = {
  playlistSongs: [
    {name: "foo"}
  ]
};

export default function(state = initialState, action) {
  console.log(state)
  switch (action.type) {
    case PLAYLIST_RECEIVED:
      return {
        ...state,
        playlistSongs: action.response
      };
    case PLAYLIST_SONG_ADDED:
      return {
        ...state,
        playlistSongs: action.response
      };
    default:
      return state;
  }
}
