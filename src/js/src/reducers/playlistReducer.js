import {
  PLAYLIST_RECEIVED,
  PLAYLIST_IS_FETCHING,
  PLAYLIST_SONG_ADDED,
  PLAYLIST_SONG_REMOVED,
  PLAYLIST_UPDATED,
  PLAYLIST_CLEARED
} from '../constants/reducer-actions.const';

const initialState = {  
  playlistSongs: [],
  isLoading: true
};

export default function(state = initialState, action) {
  switch (action.type) {
    case PLAYLIST_SONG_REMOVED:
      const filteredSongs = state.playlistSongs.filter(playlistSong => playlistSong.ID != action.removedPlaylistSongId);
      return {
        ...state,
        playlistSongs: filteredSongs,
        isLoading: false
      }
    case PLAYLIST_RECEIVED:
      return {
        ...state,
        playlistSongs: action.payload,
        isLoading: false
      }
    case PLAYLIST_UPDATED:
      return {
        ...state,
        playlistSongs: action.payload,
        isLoading: false
      }
    case PLAYLIST_SONG_ADDED:
      return {
        ...state,
        playlistSongs: action.payload,
        isLoading: false
      };
    default:
      return state;
  }
}
