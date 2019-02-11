import React from 'react';
import PropTypes from 'prop-types';
import {
  Button,
  Popconfirm,
  Icon
} from 'antd';

export default class DeleteButton extends React.Component {
  handleClick() {
    this.props.removePlaylistSong(this.props.playlistSongId);
    console.log('this is:', this.props.playlistSongId);
  }

  render() {
    // This syntax ensures `this` is bound within handleClick
    return (
      <Popconfirm placement="left" title="Are you sure？" onConfirm={(e) => this.handleClick(e)} icon={<Icon type="question-circle-o" style={{ color: 'red' }} />}>
        <Button type="danger" size="small">delete</Button>
      </Popconfirm>
      
    )
  }
}

DeleteButton.propTypes = {
  removePlaylistSong: PropTypes.func.isRequired,
  playlistSongId: PropTypes.number.isRequired
};
