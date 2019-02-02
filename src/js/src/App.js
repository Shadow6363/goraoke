import React from "react";
import { BrowserRouter as Router, Route, Link } from "react-router-dom";
import './App.css';

function KaraokeRoom() {
  return (
    <Router>
      <div>
        <Route exact path="/" component={Home} />
        <Route exact path="/playlist" component={Playlist} />
      </div>
    </Router>
  );
}

function Home() {
  return (
    <div>
      <h2>Home</h2>
    </div>
  );
}

function Playlist() {
  return (
    <div>
      <h2>Playlist</h2>
    </div>
  );
}

export default KaraokeRoom;