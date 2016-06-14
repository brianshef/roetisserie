// This is main process of Electron, started as first thing when your
// app starts. This script is running through entire life of your application.
// It doesn't have any windows which you can see on screen, but we can open
// window from here.

import { app, Menu } from 'electron';
import { devMenuTemplate } from './helpers/dev_menu_template';
import { editMenuTemplate } from './helpers/edit_menu_template';
import createWindow from './helpers/window';

// Special module holding environment variables which you declared
// in config/env_xxx.json file.
import env from './env';

//
// SERVER
//

var pkg         = require('./package.json');
var exp         = require('express');
var express     = exp();
var router      = exp.Router();
var serveStatic = require('serve-static');
var hbs         = require('hbs');
var port        = 3000;
var viewsDir    = __dirname + '/views/';
var contextDir  = __dirname + '/context/';
var server;
//  Set up Handlebars templates and rendering
//  https://github.com/donpark/hbs
express.set('view engine', 'hbs');
express.set('views', viewsDir);
hbs.registerPartials(viewsDir + 'partials/');
var site = {
  "index": {
    "route": "/",
    "template": "index",
    "context": "index.js"
  },
  "inventory": {
    "route": "/inventory",
    "template": "inventory",
    "context": "inventory.js"
  },
}

express.all('*', function(req, res, next) {
  logRequest(req);
  next();
});
//  Set up site routes for pages
express.get(site.index.route, function(req, res) {
  res.render(site.index.template, getContext(site.index.context));
});
express.get(site.inventory.route, function(req, res) {
  res.render(site.inventory.template, getContext(site.inventory.context));
});

function logRequest(req) {
  console.log('[REQEST]', req.protocol, req.method, req.path);
}

//  Return the context data object for use in templates
function getContext(context) {
  return require(contextDir + context);
}

var startServer = function (server) {
  server  = express.listen(port, function() {
    console.log("Server started; listening on port", port);
  });
}

var stopServer = function () {

}

//
// CLIENT
//

var mainWindow;

var setApplicationMenu = function () {
    var menus = [editMenuTemplate];
    if (env.name !== 'production') {
        menus.push(devMenuTemplate);
    }
    Menu.setApplicationMenu(Menu.buildFromTemplate(menus));
};

app.on('ready', function () {
    startServer(server);

    setApplicationMenu();

    var mainWindow = createWindow('main', {
        width: 1000,
        height: 600
    });

    mainWindow.loadURL('http://localhost:3000');
});

app.on('window-all-closed', function () {
    stopServer();
    app.quit();
});
