# Summit

Boilerplate for making an Electron app with React and Go. It is still early days and this is not fully ready for wide use, it requires lots of manual work to use.

## Getting Started

These instructions require knowledge of using Go and React.

Clone this repo into a project folder locally. You can then customize the `package.json`, `go.mod` and the general Go/React files.

Next installs deps with `npm install`.

Build the Go API/Server with `go install` and adjust the `dev` script in `package.json` so the executable name is correct (folder name by default).

Finally, run `npm run dev` to start up React, wait for the dev server to start and run Electron (which starts the Go executable). This gets everything up and running ready for development.

## TODO

There is currently no handling of production builds and some development areas need more elegant solutions, like running/reloading the Go server.