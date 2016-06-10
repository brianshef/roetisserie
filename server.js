var pkg         = require('./package.json');
var express     = require('express');
var serveStatic = require('serve-static');
var hbs         = require('hbs');

var port        = 3000;
var app         = express();
var router      = express.Router();
var viewsDir    = __dirname + '/views/';
var contextDir  = __dirname + '/context/';

//  Set up Handlebars templates and rendering
//  https://github.com/donpark/hbs
app.set('view engine', 'hbs');
app.set('views', viewsDir);
hbs.registerPartials(viewsDir + 'partials/');
var site = {
  "index": {
    "route": "/",
    "template": "index",
    "context": "index.js"
  },
  "dashboard": {
    "route": "/dashboard",
    "template": "dashboard",
    "context": "dashboard.js"
  },
}

app.all('*', function(req, res, next) {
  logRequest(req);
  next();
});

//  Set up site routes for pages
app.get(site.index.route, function(req, res) {
  res.render(site.index.template, getContext(site.index.context));
});
app.get(site.dashboard.route, function(req, res) {
  res.render(site.dashboard.template, getContext(site.dashboard.context));
});

function logRequest(req) {
  console.log('[REQEST]', req.protocol, req.method, req.path);
}

//  Return the context data object for use in templates
function getContext(context) {
  return require(contextDir + context);
}

// Logic that occurs when the server first starts up
function start() {
  console.log("Server started; listening on port", port);
}

// Server main entry point
function main() {
  var server  = app.listen(port, start);
}

module.exports = {
  "main": main,
  "port": port
};
