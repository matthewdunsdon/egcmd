package egcmd

// An Example provides the details of an cli command in a given context,
// whether this is at the top-level or within a command.
type Example struct {

	// Arguments contains the arguments and flags associated with the command.
	Arguments string

	// EnvVars contains any environment variables and command prefixes needed for the example
	EnvVars string

	// Description is the textual description of the example,
	// often explaining why the command should be used
	// and what the outcomes the command produces.
	Description string
}

// Cli generates the command line input for a given command
func (e Example) Cli(command string) (cli string) {
	cli = command
	if e.EnvVars != "" {
		cli = e.EnvVars + " " + cli
	}
	if e.Arguments != "" {
		cli = cli + " " + e.Arguments
	}
	return
}
