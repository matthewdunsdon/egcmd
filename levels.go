package egcmd

import (
	"strings"
)

// A Level provides a context for examples to be found,
// whether this is at the top-level of your app or at a command.
type Level struct {
	name     string
	examples []*Example
}

// Example adds a new example to the command.
func (l *Level) Example(args, decription string) (example *Example) {
	example = &Example{
		Arguments:   args,
		Description: decription,
	}
	l.examples = append(l.examples, example)
	return
}

// Envs adds a new top-level example that includes env details.
func (l *Level) Envs(env, args, decription string) (example *Example) {
	example = &Example{
		EnvVars:     env,
		Arguments:   args,
		Description: decription,
	}
	l.examples = append(l.examples, example)
	return
}

// App provides a container for top-level and command examples to be instantiated
type App struct {
	Level
	commands map[string]*Command
}

// Command adds a new top-level command.
func (a *App) Command(name string) (command *Command) {
	command, ok := a.commands[name]
	if !ok {
		command = &Command{Level: Level{name: name}}
		a.commands[name] = command
	}

	return
}

// Find returns the examples that belong to a particular level
func (a *App) Find(search string) (examples []*Example) {
	if search == a.name {
		examples = a.examples
	} else if strings.HasPrefix(search, a.name) {
		command := search[len(a.name)+1:]

		if value, ok := a.commands[command]; ok {
			examples = value.examples
		}
	}

	return
}

// Command provides a container for top-level and command examples to be instantiated
type Command struct {
	Level
}
