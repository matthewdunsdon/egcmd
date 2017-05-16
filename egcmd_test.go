package egcmd

import "testing"

func TestNew(t *testing.T) {
	app := New("app-name")
	expectedName := "app-name"

	if got := app.name; expectedName != got {
		t.Errorf("Expected app name to be %q, got %q", expectedName, got)
	}

	app.Command("cmd")

	if _, ok := app.commands["cmd"]; !ok {
		t.Errorf("Expected command to be added, but it eas not found")
	}

	app.Example("cmd", "A command")
	expectedExampleCount := 1

	if got := len(app.examples); expectedExampleCount != got {
		t.Errorf("Expected example count to be %q, got %q", expectedName, got)
	}

	return
}
