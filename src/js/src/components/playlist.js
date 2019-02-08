import {
  sortableContainer,
  sortableElement,
  sortableHandle,
} from 'react-sortable-hoc';

import DeleteButton from './deleteButton'
import React, { Component } from 'react';
import PropTypes from 'prop-types';
import { connect } from 'react-redux';
import {
  removePlaylistSong,
  getPlaylist,
  playlistChangeOrder
} from '../actions/playlist.action';

const DragHandle = sortableHandle(() => <span>::</span>);
const SortableItem = sortableElement(({playlistSong, removePlaylistSong}) => (
  <li>
    <DragHandle />
    <p><b>{playlistSong.Song.Name}</b> | {playlistSong.Song.Artist}</p>
    <p>Sort order: {playlistSong.SortOrder}</p>
    <p>ID: {playlistSong.ID}</p>
    <DeleteButton playlistSongId={playlistSong.ID} removePlaylistSong={removePlaylistSong}/>
    <hr></hr>
  </li>
));

const SortableContainer = sortableContainer(({children}) => {
  return <ul>{children}</ul>;
});

class Playlist extends Component {
  componentWillMount() {
    this.props.getPlaylist();
  }

  onSortEnd = ({oldIndex, newIndex}) => {
    this.props.playlistChangeOrder(this.props.playlistSongs[oldIndex].ID, newIndex+1)
  };

  render() {
    const playlist = this.props.playlistSongs.map((playlistSong, index) => (
      <SortableItem removePlaylistSong={this.props.removePlaylistSong} key={`playlistSongID-${playlistSong.ID}`} index={index} playlistSong={playlistSong} />
    ));
    return (
      <SortableContainer onSortEnd={this.onSortEnd} useDragHandle>
        {playlist}
      </SortableContainer>
    );
  }
}

Playlist.propTypes = {
  getPlaylist: PropTypes.func.isRequired,
  playlistChangeOrder: PropTypes.func.isRequired,
  removePlaylistSong: PropTypes.func.removePlaylistSong,
  playlistSongs: PropTypes.array.isRequired
};

const mapStateToProps = state => ({
  playlistSongs: state.playlistReducer.playlistSongs
});

export default connect(mapStateToProps, { getPlaylist, playlistChangeOrder, removePlaylistSong })(Playlist);