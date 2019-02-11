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

import {
  List,
  Skeleton,
  Row,
  Col,
  Card,
  Icon,
  Button
} from 'antd';

const DragHandle = sortableHandle(() => <Icon type="ellipsis" />);
const SortableItem = sortableElement(({playlistSong, removePlaylistSong, props}) => (
  <Skeleton avatar title={false} loading={props.isLoading} active>
    <List.Item actions={[
      <Button shape="circle" icon="caret-up" />,
      <Button shape="circle" icon="caret-down" />,
      <DeleteButton playlistSongId={playlistSong.ID} removePlaylistSong={removePlaylistSong}/>
    ]}>
      <DragHandle /> <b>{playlistSong.Song.Name}</b> | {playlistSong.Song.Artist}
    </List.Item>
  </Skeleton>
));

const SortableContainer = sortableContainer(({children}) => {
  return <div>{children}</div>;
});

class Playlist extends Component {
  componentWillMount() {
    this.props.getPlaylist();
  }

  onSortEnd = ({oldIndex, newIndex}) => {
    this.props.playlistChangeOrder(this.props.playlistSongs[oldIndex].ID, this.props.playlistSongs[newIndex].SortOrder)
  };

  render() {
    const playlist = this.props.playlistSongs.map((playlistSong, index) => (
      <SortableItem props={this.props} removePlaylistSong={this.props.removePlaylistSong} sortOrder={playlistSong.SortOrder} key={`playlistSongID-${playlistSong.ID}`} index={index} playlistSong={playlistSong} />
    ));
    return (
      
        <Row gutter={16}>
          <Col className="gutter-row" span={24}>
            <Card>
              <SortableContainer onSortEnd={this.onSortEnd} useDragHandle>
                {playlist}
              </SortableContainer>
            </Card>
          </Col>
        </Row>
      
    );
  }
}

Playlist.propTypes = {
  getPlaylist: PropTypes.func.isRequired,
  playlistChangeOrder: PropTypes.func.isRequired,
  removePlaylistSong: PropTypes.func.isRequired,
  playlistSongs: PropTypes.array.isRequired,
  isLoading: PropTypes.bool.isRequired
};

const mapStateToProps = state => ({
  playlistSongs: state.playlistReducer.playlistSongs,
  isLoading: state.playlistReducer.isLoading
});

export default connect(mapStateToProps, { getPlaylist, playlistChangeOrder, removePlaylistSong })(Playlist);


