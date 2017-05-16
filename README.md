# egcmd
[![GoDoc](https://godoc.org/github.com/matthewdunsdon/egcmd?status.svg)](https://godoc.org/github.com/matthewdunsdon/egcmd)
[![Build Status](https://travis-ci.org/matthewdunsdon/egcmd.svg?branch=master)](https://travis-ci.org/matthewdunsdon/egcmd)
[![Coverage Status](https://coveralls.io/repos/github/matthewdunsdon/egcmd/badge.svg?branch=master)](https://coveralls.io/github/matthewdunsdon/egcmd?branch=master)
[![Go Report Card](https://goreportcard.com/badge/github.com/matthewdunsdon/egcmd)](https://goreportcard.com/report/github.com/matthewdunsdon/egcmd)

Package `matthewdunsdon/egcmd` provides you with a library to include examples for your commands in your CLI golang app.

Currently, this library is not intended for use in production, predominantly as it has been created learning golang.

The main features are:

* It defines a `egcmd.Example` type to describe an example CLI command, including arguments and environment variables
* It allows examples to be associated with the top level of your CLI app

---

* [Install](#install)
* [Usage](#usage)
* [License](./LICENSE)

---

## Install

With a [correctly configured](https://golang.org/doc/install#testing) Go toolchain:

```sh
go get github.com/matthewdunsdon/egcmd
```

## Usage

To get started you should create an `examples.go` file in your project as the same level as `package main`.  In this file you will want to set the correct package and import, which should look something like:

```go
package main

import (
	"github.com/matthewdunsdon/egcmd"
)
```

Typically you can add your cli examples as variables, like shown below:

```go
var (
	myAppEx = egcmd.New("myapp")
	_       = myAppEx.Example("init", "Initalise app data.")
	_       = myAppEx.Example("version --json", "Get application version details in json format.")

	initCmdEx = myAppEx.Command("init")
	_         = initCmdEx.Example("--defaults", "Initalise app data using the recommended defaults")
	_         = initCmdEx.Envs("MYAPP_PATH=~/Documents/", "", "Initialise app data to custom location using MYAPP_PATH")
)
```

You are by no means forced to organise your examples like this;

- You could add these to `func init()` or to a function of you own.
- Your examples do have to be in the main package, but you will need to link up your CLI help code with your egcmd app instance(s).


## License

MIT licensed. See the LICENSE file for details.
