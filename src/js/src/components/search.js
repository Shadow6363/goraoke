import React, { Component } from 'react';
import { connect } from 'react-redux';
import { Input, Affix } from 'antd';
import PropTypes from 'prop-types';
import {
  search
} from '../actions/playlist.action';


const SearchInput = Input.Search;

class Search extends Component {
  render() {
    return (
      <div style={{padding: "5px", background: "#fff"}}>
      <Affix offsetTop={5}>
        <SearchInput
          placeholder="Song or Artist"
          enterButton="Search"
          size="large"
          onSearch={value => this.props.search(value)}
        />
      </Affix>
      </div>
    );
  }
}

const mapStateToProps = state => ({ 
  search: PropTypes.func.isRequired
});

export default connect(mapStateToProps, {search})(Search);
