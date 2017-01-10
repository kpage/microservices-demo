var path = require('path');
var express = require('express');
var webpack = require('webpack');
var config = require('./webpack.dev.config');
var history = require('connect-history-api-fallback');
var proxy = require('http-proxy-middleware');

var app = express();
var compiler = webpack(config);

app.use(history());
app.use(proxy('/api',
	      {
		  target: 'http://rest-api:3000',
		  changeOrigin: true,
		  xfwd: true
	      }));

app.use(require('webpack-dev-middleware')(compiler, {
    noInfo: true,
    publicPath: config.output.publicPath
}));

app.use(require('webpack-hot-middleware')(compiler));

app.listen(4000, '0.0.0.0', (err) => {
    if (err) {
      console.log(err);
      return;
    }
  
    console.log('Listening at http://localhost:4000');
});
