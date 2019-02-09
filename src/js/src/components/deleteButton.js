import React from 'react';
import PropTypes from 'prop-types';

export default class DeleteButton extends React.Component {
  handleClick() {
    this.props.removePlaylistSong(this.props.playlistSongId);
    console.log('this is:', this.props.playlistSongId);
  }

  render() {
    // This syntax ensures `this` is bound within handleClick
    return (
      <button onClick={(e) => this.handleClick(e)}>
        Delete
      </button>
    );
  }
}

DeleteButton.propTypes = {
  removePlaylistSong: PropTypes.func.isRequired,
  playlistSongId: PropTypes.number.isRequired
};
