
import IS_SEARCHING from './constants/reducer-actions.const';


class App extends React.Component {
  render() {
    return (<h1>heyo</h1>);
  }
}

ReactDOM.render(<App />, document.getElementById('app'));

// import React from 'react';
// import { render } from 'react-dom';
// // import { Provider } from 'react-redux';

// // import configureStore from './stores/configureStore';
// // import { Switch, Router, Route } from 'react-router';
// // import createHistory from 'history/createBrowserHistory';
// // import injectTapEventPlugin from 'react-tap-event-plugin';

// // import Home from './components/Home';
// // import Navbar from './components/Navbar';

// // injectTapEventPlugin();

// // export const appStore = configureStore();

// render(
//   <Provider store={appStore}>
//    <Router history={createHistory()}>
//      <div>
//        <Navbar />
//        <Switch>
//          <Route path='/' component={Home}></Route>
//         </Switch>
//       </div>
//     </Router>
//   </Provider>,
//   document.getElementById('root')
// );
