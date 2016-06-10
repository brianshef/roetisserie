var pkg     = require('./package.json');
var server  = require('./server.js');
var opener  = require('opener');

function unref(app) {
  if(app) {
    app.unref();
    app.stdin.unref();
    app.stdout.unref();
    app.stderr.unref();
  }
}

function run() {
  //  Log the version
  console.log(pkg.name, pkg.version);
  //  Start the server
  server.main();
  //  Open the default OS browser to the server's location
  var app = opener('http://localhost:' + server.port);
  //  Unreference opened app to free up resources
  unref(app);
}

//  Script main entry point
run();
