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
      <div key={playlistSong.ID}>
        <h3>{playlistSong.Song.Name}</h3>
        <p>{playlistSong.Song.Artist}</p>
        <p>{playlistSong.Song.ID}</p>
      </div>
    ));
    console.log(playlist);
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
  playlistSongs: state.playlistReducer.playlistSongs
});

export default connect(mapStateToProps, { getPlaylist })(Playlist);