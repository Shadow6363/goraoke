import React from "react";
import { Provider } from 'react-redux';
import Playlist from './components/playlist'
import Complete from './components/search'
import Search from './components/search'
import { BrowserRouter as Router, Route, Link } from "react-router-dom";
import "antd/dist/antd.css";
import './App.css';
import store from './store';


function KaraokeRoom() {
  return (
    <Provider store={store}>
      <Router>
        <div>
          <Route exact path="/" component={Home} />
          <Route exact path="/playlist" component={RenderPlaylist} />
        </div>
      </Router>
    </Provider>
  );
}

function Home() {
  return (
    <div>
      <h2>Home</h2>
    </div>
  );
}

function RenderPlaylist() {
  return (
    <div>
      <Search />
      <h2>Playlist</h2>
      <Playlist removePlaylistSong={() => {}}/>
    </div>
  );
}

export default KaraokeRoom;