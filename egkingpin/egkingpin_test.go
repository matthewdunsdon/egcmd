package egkingpin

import (
	"errors"
	"testing"
	"text/template"

	"github.com/matthewdunsdon/egcmd"

	kingpin "gopkg.in/alecthomas/kingpin.v3-unstable"
)

type testExamplesFinder struct {
	lastCommand string
}

func (t *testExamplesFinder) Find(command string) (examples egcmd.ExamplesFound) {
	t.lastCommand = command
	return
}

func TestNew(t *testing.T) {
	var (
		app, appExamples = New("app-name", "help details")
		expectedName     = "app-name"
		expectedHelp     = "help details"
	)

	if got := app.Name; expectedName != got {
		t.Errorf("Expected app name to be %q, got %q", expectedName, got)
	}

	if got := app.Help; expectedHelp != got {
		t.Errorf("Expected app help to be %q, got %q", expectedName, got)
	}

	if got := appExamples.Find("").Context; expectedName != got {
		t.Errorf("Expected app examples name to be %q, got %q", expectedName, got)
	}
}

func TestUpdateUsageContext(t *testing.T) {
	testCases := []struct {
		testName     string
		usageContext *kingpin.UsageContext
		finder       ExamplesFinder
		want         error
	}{
		{
			"AllPresent",
			&kingpin.UsageContext{},
			&testExamplesFinder{},
			nil,
		},
		{
			"NoUsageContext",
			nil,
			&testExamplesFinder{},
			errors.New("No usage context supplied to UpdateUsageContext"),
		},
		{
			"NoAppExamples",
			&kingpin.UsageContext{},
			nil,
			errors.New("No app examples instance supplied to UpdateUsageContext"),
		},
	}

	for _, tc := range testCases {
		t.Run(tc.testName, func(t *testing.T) {
			err := UpdateUsageContext(tc.usageContext, tc.finder)

			if err != nil {
				if tc.want == nil {
					t.Errorf("Expected UpdateUsageContext to return without error, got %#v", err)
				} else if got := err.Error(); tc.want.Error() != got {
					t.Errorf("Expected UpdateUsageContext to return %#v, got %#v", tc.want, err)
				}
			} else if tc.want != nil {
				t.Errorf("Expected UpdateUsageContext to return %#v, but no error returned", tc.want)
			}
		})
	}
}

func TestUpdateTemplate(t *testing.T) {
	customTemplate := `Details:
`
	testCases := []struct {
		testName     string
		usageContext *kingpin.UsageContext
		want         string
	}{
		{
			"EmptyValue",
			&kingpin.UsageContext{},
			kingpin.DefaultUsageTemplate + kingpinUsageTemplate,
		},
		{
			"Extends",
			&kingpin.UsageContext{
				Template: customTemplate,
			},
			customTemplate + kingpinUsageTemplate,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.testName, func(t *testing.T) {
			updateTemplate(tc.usageContext)

			if got := tc.usageContext.Template; tc.want != got {
				t.Errorf("Expected template to be set to %q, got %q", tc.want, got)
			}
		})
	}
}

func TestUpdateTemplateFunctions(t *testing.T) {
	examplesFinder := &testExamplesFinder{}

	testCases := []struct {
		testName     string
		usageContext *kingpin.UsageContext
		want         []string
	}{
		{
			"EmptyValue",
			&kingpin.UsageContext{},
			[]string{
				"findExamples",
				"findAppExamples",
			},
		},
		{
			"Extends",
			&kingpin.UsageContext{
				Funcs: template.FuncMap{
					"existingFunc": func() string { return "Hello" },
				},
			},
			[]string{
				"existingFunc",
				"findExamples",
				"findAppExamples",
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.testName, func(t *testing.T) {
			updateTemplateFunctions(tc.usageContext, examplesFinder)

			if got := len(tc.usageContext.Funcs); len(tc.want) != got {
				t.Errorf("Expected template functions count be to %q, got %q", len(tc.want), got)
			}

			for _, wantedItem := range tc.want {
				if _, ok := tc.usageContext.Funcs[wantedItem]; !ok {
					t.Errorf("Expected template functions to contain %q, but was not present", wantedItem)
				}
			}

		})
	}
}

func TestUpdateTemplateFunctionsFindExamples(t *testing.T) {
	var (
		examplesFinder = &testExamplesFinder{}
		usageContext   = &kingpin.UsageContext{}
		command        = "my-command"
	)

	updateTemplateFunctions(usageContext, examplesFinder)

	findExamples, ok := usageContext.Funcs["findExamples"].(findExamplesAction)

	if !ok {
		t.Errorf("Expected findExamples to be type 'findExamplesAction', got %T", usageContext.Funcs["findExamples"])
		return
	}

	findExamples(command)

	if got := examplesFinder.lastCommand; command != got {
		t.Errorf("Expected findExamples to call ExamplesFinder.Find with %q, got %q", command, got)
	}
}

func TestUpdateTemplateFunctionsFindAppExamples(t *testing.T) {
	var (
		examplesFinder = &testExamplesFinder{lastCommand: "something"}
		usageContext   = &kingpin.UsageContext{}
		want           = ""
	)

	updateTemplateFunctions(usageContext, examplesFinder)

	findAppExamples, ok := usageContext.Funcs["findAppExamples"].(findAppExamplesAction)

	if !ok {
		t.Errorf("Expected findAppExamples to be type 'findAppExamplesAction', got %T", usageContext.Funcs["findAppExamples"])
		return
	}

	findAppExamples()

	if got := examplesFinder.lastCommand; want != got {
		t.Errorf("Expected findAppExamples to call ExamplesFinder.Find with %q, got %q", want, got)
	}
}
