import {
  IS_SEARCHING,
  SEARCH_RESULTS_RECEIVED
} from '../constants/reducer-actions.const';

const initialState = {
  songs: []
};

export default function(state = initialState, action) {
  switch (action.type) {
    case SEARCH_RESULTS_RECEIVED:
      return {
        ...state,
        songs: action.payload
      };
    case IS_SEARCHING:
      return {
        ...state,
        songs: action.payload
      };
    default:
      return state;
  }
}
