/*
Package egkingpin provides you with a library to seamlessly intergrate `kingpin` (https://github.com/alecthomas/kingpin) with `egcmd` (https://github.com/matthewdunsdon/egcmd).

Currently, this library is not intended for use in production, predominantly as it has been created learning golang.
*/
package egkingpin

import (
	"errors"
	"text/template"

	"github.com/matthewdunsdon/egcmd"
	kingpin "gopkg.in/alecthomas/kingpin.v3-unstable"
)

var (
	kingpinUsageTemplate = `{{define "FormatExamples" -}}
{{if .Examples -}}
Examples:
{{$context:=.Context -}}
{{range .Examples -}}
{{Indent 1 -}}{{.Cli $context}}
{{.Description|Wrap 4}}

{{end -}}
{{end -}}
{{end -}}

{{if .Context.SelectedCommand -}}
{{template "FormatExamples" findExamples .Context.SelectedCommand.FullCommand -}}
{{else -}}
{{template "FormatExamples" findAppExamples -}}
{{end}}`
)

// ExamplesFinder is the interface that wraps the Find method.
//
// Find returns the examples that belong to a particular command.
type ExamplesFinder interface {
	Find(command string) egcmd.ExamplesFound
}

type findExamplesAction func(commandName string) egcmd.ExamplesFound
type findAppExamplesAction func() egcmd.ExamplesFound

// New creates both kingpin and egcmd application, which have the usage context
// configured so that examples are shown.
func New(name, help string) (app *kingpin.Application, appExamples *egcmd.App) {
	app = kingpin.New(name, help)
	appExamples = egcmd.New(name)
	usageContext := &kingpin.UsageContext{}

	updateTemplate(usageContext)
	updateTemplateFunctions(usageContext, appExamples)

	app.UsageContext(usageContext)

	return
}

// UpdateUsageContext takes an existing kingpin usage context along with an app examples instance
// and appends to the template and template functions so that it will find and render examples.
func UpdateUsageContext(usageContext *kingpin.UsageContext, finder ExamplesFinder) (err error) {

	if usageContext == nil {
		return errors.New("No usage context supplied to UpdateUsageContext")
	}

	if finder == nil {
		return errors.New("No app examples instance supplied to UpdateUsageContext")
	}

	updateTemplate(usageContext)
	updateTemplateFunctions(usageContext, finder)

	return nil
}

func updateTemplate(usageContext *kingpin.UsageContext) {

	if usageContext.Template == "" {
		usageContext.Template = kingpin.DefaultUsageTemplate
	}

	usageContext.Template += kingpinUsageTemplate
}

func updateTemplateFunctions(usageContext *kingpin.UsageContext, finder ExamplesFinder) {
	var (
		findExamples    findExamplesAction
		findAppExamples findAppExamplesAction
	)

	findExamples = func(commandName string) egcmd.ExamplesFound {
		return finder.Find(commandName)
	}
	findAppExamples = func() egcmd.ExamplesFound {
		return finder.Find("")
	}

	if usageContext.Funcs == nil {
		usageContext.Funcs = template.FuncMap{
			"findExamples":    findExamples,
			"findAppExamples": findAppExamples,
		}
	} else {
		usageContext.Funcs["findExamples"] = findExamples
		usageContext.Funcs["findAppExamples"] = findAppExamples
	}

}
