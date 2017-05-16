package egcmd

import (
	"testing"
)

func TestExampleCli(t *testing.T) {
	testCases := []struct {
		name    string
		example Example
		want    string
	}{
		{
			"Simple",
			Example{},
			"my-app init",
		},
		{
			"WithArguments",
			Example{Arguments: "--force"},
			"my-app init --force",
		},
		{
			"WithEnvVars",
			Example{EnvVars: "ENV=PROD"},
			"ENV=PROD my-app init",
		},
		{
			"WithArgumentsAndEnvVars",
			Example{EnvVars: "ENV=PROD", Arguments: "--force"},
			"ENV=PROD my-app init --force",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			if got := tc.example.Cli("my-app init"); tc.want != got {
				t.Errorf("Expected cli to return %q, got %q", tc.want, got)
			}
		})
	}
}
