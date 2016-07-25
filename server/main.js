import { Meteor } from 'meteor/meteor';
var open  = require('open');
var port  = 3000;

Meteor.startup(() => {
  // code to run on server at startup
  open('http://localhost:' + port, function (err) {
    if (err) throw err;
  });
});
