/*
Package egcmd provides you with a library to include examples for your commands in your CLI golang app.

Currently, this library is not intended for use in production, predominantly as it has been created learning golang.

Usage:

To get started you should create an `examples.go` file in your project as the same level as `package main`.  In this file you will want to set the correct package and import, which should look something like:

	package main

	import (
		"github.com/matthewdunsdon/egcmd"
	)

Typically you can add your cli examples as variables, like shown below:

	var (
		myAppEx = egcmd.New("myapp")
		_       = myAppEx.Example("init", "Initalise app data.")
		_       = myAppEx.Example("version --json", "Get application version details in json format.")

		initCmdEx = myAppEx.Command("init")
		_         = initCmdEx.Example("--defaults", "Initalise app data using the recommended defaults")
		_         = initCmdEx.Envs("MYAPP_PATH=~/Documents/", "", "Initialise app data to custom location using MYAPP_PATH")
	)


You are by no means forced to organise your examples like this. You could add these to `func init()` or to a function of you own.
Your examples do have to be in the main package, but you will need to link up your CLI help code with your egcmd app instance(s).
*/
package egcmd

// New creates a new application instance whether examples can be added.
func New(name string) (app *App) {
	app = &App{
		Level:    Level{name: name},
		commands: make(map[string]*Command),
	}
	return
}
