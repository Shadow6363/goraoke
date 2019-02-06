const proxy = require('http-proxy-middleware');

module.exports = function(app) {
  app.use(proxy('/api', { target: 'http://localhost:3000/' }));
  app.use(proxy('/api/ws', {headers: {"accept-encoding":""}, target: 'ws://localhost:3000', changeOrigin: true, "ws": true }));
  app.use(proxy("/api/ws", {headers: {"accept-encoding":""}, target: 'ws://localhost:3000', changeOrigin: true, "ws": true }));
  console.log("configuring proxy")
};

// ,
//   "proxy": {
//     "/": {
//       "target": "http://localhost:4000"
//     },
//     "/sockjs-node": {
//       "target": "ws://localhost:4000",
//       "ws": true
//     },
//     "/socket.io": {
//       "target": "ws://localhost:4000",
//       "ws": true
//     }
//   }