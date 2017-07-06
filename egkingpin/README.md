# egcmd/egkingpin
[![GoDoc](https://godoc.org/github.com/matthewdunsdon/egcmd/egkingpin?status.svg)](https://godoc.org/github.com/matthewdunsdon/egcmd/egkingpin)
[![Build Status](https://travis-ci.org/matthewdunsdon/egcmd.svg?branch=master)](https://travis-ci.org/matthewdunsdon/egcmd)
[![Coverage Status](https://coveralls.io/repos/github/matthewdunsdon/egcmd/badge.svg?branch=master)](https://coveralls.io/github/matthewdunsdon/egcmd?branch=master)
[![Go Report Card](https://goreportcard.com/badge/github.com/matthewdunsdon/egcmd/egkingpin)](https://goreportcard.com/report/github.com/matthewdunsdon/egcmd/egkingpin)

Package `matthewdunsdon/egcmd/egkingpin` provides you with a library to seamlessly intergrate [`kingpin`](https://github.com/alecthomas/kingpin) with [`egcmd`](https://github.com/matthewdunsdon/egcmd).

Currently, this library is not intended for use in production, predominantly as it has been created learning golang.

---

* [Install](#install)
* [Usage](#usage)
* [License](./LICENSE)

---

## Install

With a [correctly configured](https://golang.org/doc/install#testing) Go toolchain:

```sh
go get github.com/matthewdunsdon/egcmd/egkingpin
```

## Usage

The simplist approach is to create a configured kingpin app is to use the `.New()` function:

```go
import (
	"github.com/matthewdunsdon/egcmd/egkingpin"
)

var (
	app, appExamples = egkingpin.New("myapp", "This is my app.")
	_                = appExamples.Example("init", "Ius legimus nonumes te, pri dicat nominavi copiosae id, odio rebum facilis ea pro.")

	initCmd   = app.Command("init", "Initialise cashflow data.")
	initCmdEx = appExamples.Command("init")
	_         = initCmdEx.Example("--yes", "At vis primis debitis, ei verear omittantur signiferumque mei, quo esse aperiri an. Dolore vocent consequuntur pro an, nam no iusto tamquam suscipit.")
)
```

If you have other code the relies on a customised `kingpin.UsageContext`, then you will need to do something like.

```go
import (
	"github.com/matthewdunsdon/egcmd"
	"github.com/matthewdunsdon/egcmd/egkingpin"
	"gopkg.in/alecthomas/kingpin.v3-unstable"
)

var (
	app         = kingpin.New("myapp", "This is my app.")
	appExamples = egcmd.New("myapp")
	_           = appExamples.Example("init", "Ius legimus nonumes te, pri dicat nominavi copiosae id, odio rebum facilis ea pro.")

	initCmd   = app.Command("init", "Initialise cashflow data.")
	initCmdEx = appExamples.Command("init")
	_         = initCmdEx.Example("--yes", "At vis primis debitis, ei verear omittantur signiferumque mei, quo esse aperiri an. Dolore vocent consequuntur pro an, nam no iusto tamquam suscipit.")
)

func init() {
	usageContext := &kingpin.UsageContext{
		// Some configuration here
	}

	// Update the existing usage context so that templates and function definitions
	// have been updated
	egkingpin.UpdateUsageContext(usageContext, appExamples)

	// Apply configuration to kingpin
	app.UsageContext(usageContext)
}
```

## License

MIT licensed. See the LICENSE file for details.
