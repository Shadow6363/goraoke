import React, { Component } from 'react';
import PropTypes from 'prop-types';
import { connect } from 'react-redux';
import { getPlaylist } from '../actions/playlist.action';


class Playlist extends Component {
  componentWillMount() {
    this.props.getPlaylist();
  }

  // componentWillReceiveProps(nextProps) {
  //   if (nextProps.newPost) {
  //     this.props.posts.unshift(nextProps.newPost);
  //   }
  // }

  render() {
    const playlist = this.props.playlistSongs.map(playlistSong => (
      <div key={playlistSong.id}>
        <h3>{playlistSong.name}</h3>
        <p>{playlistSong.artist}</p>
      </div>
    ));
    return (
      <div>
        <h1>Playlist</h1>
        {playlist}
      </div>
    );
  }
}

Playlist.propTypes = {
  getPlaylist: PropTypes.func.isRequired,
  playlistSongs: PropTypes.array.isRequired
};

const mapStateToProps = state => ({
  playlistSongs: state.playlistSongs
});

export default connect(mapStateToProps, { getPlaylist })(Playlist);