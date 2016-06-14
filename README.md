# roetisserie
Inventory and sales management for LuLaRoe consultants.

## Development
### About
Built with Electron, the app contains both the client and the server components
together in a single repository. Thanks to Electron, the entire application
can be bundled together and released as a single executable for Linux,
OSX/macOS, and Windows.
### Install
`npm install`
### Build and run locally
`npm start`
### Create a release
**Note** You must create a release in the target environment; Linux releases
can only be built on Linux machines, etc.
`npm release`
#### Windows only
#### Installer
The installer is built using [NSIS](http://nsis.sourceforge.net). You have to install NSIS version 3.0, and add its folder to PATH in Environment Variables, so it is reachable to scripts in this project. For example, `C:\Program Files (x86)\NSIS`.

Special thanks to my wife, Kristen, who works so hard for our family.

---

Comprehensive boilerplate application for [Electron runtime](http://electron.atom.io).  

Scope of this project:

- Provide basic structure of the application so you can much easier grasp what should go where.
- Give you cross-platform development environment, which works the same way on OSX, Windows and Linux.
- Generate ready for distribution installers of your app for all three operating systems.

Not in the scope:

- Imposing on you any framework (e.g. Angular, React). You can integrate the one which makes most sense for you.

# Quick start
The only development dependency of this project is [Node.js](https://nodejs.org). So just make sure you have it installed.
Then type few commands known to every Node developer...
```
git clone https://github.com/szwacz/electron-boilerplate.git
cd electron-boilerplate
npm install
npm start
```
... and boom! You have running desktop application on your screen.

# Structure of the project

There are **two** `package.json` files:  

#### 1. For development
Sits on path: `electron-boilerplate/package.json`. Here you declare dependencies for your development environment and build scripts. **This file is not distributed with real application!**

Also here you declare the version of Electron runtime you want to use:
```json
"devDependencies": {
  "electron-prebuilt": "^0.34.0"
}
```

#### 2. For your application
Sits on path: `electron-boilerplate/app/package.json`. This is **real** manifest of your application. Declare your app dependencies here.

#### OMG, but seriously why there are two `package.json`?
1. Native npm modules (those written in C, not JavaScript) need to be compiled, and here we have two different compilation targets for them. Those used in application need to be compiled against electron runtime, and all `devDependencies` need to be compiled against your locally installed node.js. Thanks to having two files this is trivial.
2. When you package the app for distribution there is no need to add up to size of the app with your `devDependencies`. Here those are always not included (because reside outside the `app` directory).

### Project's folders

- `app` - code of your application goes here.
- `config` - place where you can declare environment specific stuff for your app.
- `build` - in this folder lands built, runnable application.
- `releases` - ready for distribution installers will land here.
- `resources` - resources needed for particular operating system.
- `tasks` - build and development environment scripts.


# Development

#### Installation

```
npm install
```
It will also download Electron runtime, and install dependencies for second `package.json` file inside `app` folder.

#### Starting the app

```
npm start
```

#### Adding npm modules to your app

Remember to add your dependency to `app/package.json` file, so do:
```
cd app
npm install name_of_npm_module --save
```

#### Native npm modules

Want to use [native modules](https://github.com/atom/electron/blob/master/docs/tutorial/using-native-node-modules.md)? This objective needs some extra work (rebuilding module for Electron). In this boilerplate it's fully automated, just use special command instead of standard `npm install something` when want to install native module.
```
npm run install-native -- name_of_native_module
```
This script when run first time will add [electron-rebuild](https://github.com/electronjs/electron-rebuild) to your project. After that everything is wired and no further maintenance is necessary.

#### Working with modules

How about being future proof and using ES6 modules all the time in your app? Thanks to [rollup](https://github.com/rollup/rollup) you can do that. It will transpile the imports to proper `require()` statements, so even though ES6 modules aren't natively supported yet you can start using them today.

You can use it on those kinds of modules:
```js
// Modules authored by you
import { myStuff } from './my_lib/my_stuff';
// Node.js native
import fs from 'fs';
// Electron native
import { app } from 'electron';
// Loaded from npm
import moment from 'moment';
```

#### Including files to your project

The build script copies files from `app` to `build` directory and the application is started from `build`. Therefore if you want to use any special file/folder in your app make sure it will be copied via some of glob patterns in `tasks/build/build.js`:

```js
var paths = {
    copyFromAppDir: [
        './node_modules/**',
        './vendor/**',
        './**/*.html',
        './**/*.+(jpg|png|svg)'
    ],
}
```

#### Unit tests

electron-boilerplate has preconfigured [mocha](https://mochajs.org/) test runner with the [chai](http://chaijs.com/api/assert/) assertion library. To run the tests go with standard:
```
npm test
```
You don't have to declare paths to spec files in any particular place. The runner will search through the project for all `*.spec.js` files and include them automatically.

Those tests can be plugged into [continuous integration system](https://github.com/atom/electron/blob/master/docs/tutorial/testing-on-headless-ci.md).
