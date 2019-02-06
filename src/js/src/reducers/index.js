import { combineReducers } from 'redux';
import playlistReducer from './playlistReducer';
import searchReducer from './searchReducer'

export default combineReducers({
  playlist: playlistReducer,
  search: searchReducer
});
