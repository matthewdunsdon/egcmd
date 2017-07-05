package egcmd

import (
	"reflect"
	"testing"
)

func TestLevelExample(t *testing.T) {
	level := Level{name: "root"}
	testCases := []struct {
		name        string
		arguments   string
		description string
		want        Example
	}{
		{
			"NoArguments",
			"",
			"Default behaviour",
			Example{Description: "Default behaviour"},
		},
		{
			"WithArguments",
			"--verbose",
			"verbose output",
			Example{Arguments: "--verbose", Description: "verbose output"},
		},
	}

	for i, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			actual := level.Example(tc.arguments, tc.description)
			count := i + 1

			if got := len(level.examples); count != got {
				t.Errorf("Expected level example count to be %d, got %d", count, got)
			}
			if got := *actual; tc.want != got {
				t.Errorf("Expected returned example to be %v, got %v", tc.want, got)
			}
			if got := *level.examples[i]; tc.want != got {
				t.Errorf("Expected level example %d to be %v, got %v", i, tc.want, got)
			}
		})
	}
}

func TestLevelEnvs(t *testing.T) {
	level := Level{name: "root"}
	testCases := []struct {
		name        string
		envVars     string
		arguments   string
		description string
		want        Example
	}{
		{
			"NoEnvs",
			"",
			"",
			"Default behaviour",
			Example{Description: "Default behaviour"},
		},
		{
			"WithEnvs",
			"LOGGING=verbose",
			"",
			"verbose output",
			Example{EnvVars: "LOGGING=verbose", Description: "verbose output"},
		},
	}

	for i, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			actual := level.Envs(tc.envVars, tc.arguments, tc.description)
			count := i + 1

			if got := len(level.examples); count != got {
				t.Errorf("Expected level example count to be %d, got %d", count, got)
			}
			if got := *actual; tc.want != got {
				t.Errorf("Expected returned example to be %v, got %v", tc.want, got)
			}
			if got := *level.examples[i]; tc.want != got {
				t.Errorf("Expected level example %d to be %v, got %v", i, tc.want, got)
			}
		})
	}
}

func TestAppCommand(t *testing.T) {
	app := App{Level: Level{name: "root"}, commands: make(map[string]*Command)}
	testCases := []struct {
		name             string
		command          string
		exampleToAppend  *Example
		duplicateCommand bool
	}{
		{
			"FirstCommand",
			"init",
			&Example{Description: "first"},
			false,
		},
		{
			"DuplicateCommand",
			"init",
			&Example{Description: "second"},
			true,
		},
		{
			"NextCommand",
			"LOGGING=verbose",
			&Example{Description: "next"},
			false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			actual := app.Command(tc.command)
			isDuplicate := len(actual.examples) > 0
			appCommand, found := app.commands[tc.command]

			actual.examples = append(actual.examples, tc.exampleToAppend)

			if got := actual.name; tc.command != got {
				t.Errorf("Expected returned command to be %q, got %q", tc.command, got)
			}

			if !found {
				t.Errorf("Expected command in app to be present, but was not found")
				return
			}

			if got := appCommand.name; tc.command != got {
				t.Errorf("Expected command in app to have the name %q, got %q", tc.command, got)
			}
			if got := isDuplicate; tc.duplicateCommand != got {
				t.Errorf("Expected command to have an example count of %t, got %t", tc.duplicateCommand, got)
			}

		})
	}
}

func TestAppFind(t *testing.T) {
	appExamples := []*Example{
		{"", "Action: default", ""},
		{"simple", "Action: simple", ""},
		{"complex sub-action", "Action: complex sub-action", ""},
	}
	simpleExamples := []*Example{
		{"", "Default for simple", ""},
		{"--debug", "Simple with debug", ""},
	}
	complexExamples := []*Example{
		{"--json", "Default for complex with json arg", ""},
		{"sub-action", "complex action with sub-action", ""},
	}
	subActionExamples := []*Example{
		{"", "Subaction", "APP_KEY=1234"},
	}

	app := App{
		Level: Level{"root", appExamples},
		commands: map[string]*Command{
			"simple":             {Level: Level{"simple", simpleExamples}},
			"complex":            {Level: Level{"complex", complexExamples}},
			"complex sub-action": {Level: Level{"complex sub-action", subActionExamples}},
		},
	}

	testCases := []struct {
		testName string
		command  string
		want     ExamplesFound
		ok       bool
	}{
		{
			"Root",
			"",
			ExamplesFound{Context: "root", Examples: appExamples},
			true,
		},
		{
			"SimpleCommand",
			"simple",
			ExamplesFound{Context: "root simple", Examples: simpleExamples},
			true,
		},
		{
			"ComplexCommand",
			"complex",
			ExamplesFound{Context: "root complex", Examples: complexExamples},
			true,
		},
		{
			"ComplexSubCommand",
			"complex sub-action",
			ExamplesFound{Context: "root complex sub-action", Examples: subActionExamples},
			true,
		},
		{
			"CommandNotFound",
			"no-match",
			ExamplesFound{},
			false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.testName, func(t *testing.T) {
			got, ok := app.Find(tc.command)

			if tc.ok != ok {
				t.Errorf("Expected ok to be %#v, got %#v", tc.want, got)
			}

			if tc.want.Context != got.Context {
				t.Errorf("Expected context to be %#v, got %#v", tc.want, got)
			}

			if !reflect.DeepEqual(tc.want.Examples, got.Examples) {
				t.Errorf("Expected the examples found to be list %#v, got %#v", tc.want, got)
			}
		})
	}
}
