package flags

// Value all values must adhere to this interface
type Value interface {
	String() string             // String representation of the value
	DefaultValueString() string // String representation of the default value
	// Set - Used to set from the application's arguments.
	// If RequiresValue() == true, ("true" or "") or "false" are expected
	Set(string) error
	IsBoolValue() bool // If it returns true and no parameter is specified, Set will be called with "true"
}

// Option Application or command level option
type Option struct {
	Short       rune   // Short option name (i.e. 'd')
	Long        string // Long option name (i.e. "debug")
	Description string // Option's description (i.e. "Log debug messages")
	Value       Value  // Option's value and default value
}

// Command a command, or subcommand, called by the user
type Command struct {
	Name        string     // Name of the command
	Description string     // Description of the command
	Options     []*Option  // Eventual options bound to the command
	SubCommands []*Command // Eventual sub-commands
	Called      bool
}

// Flags main struct for setting up commands and options
type Flags struct {
	AppName        string     // application's name
	AppVersion     string     // application's version
	AppDescription string     // application's description
	Options        []*Option  // application-level options
	Commands       []*Command // available commands
	currentCommand *Command
}

// EmptyShort the short option name's null-value
var EmptyShort rune

// HelpOption add it to the root to enable the automatic help prompt
var HelpOption = &Option{
	Short:       'h',
	Long:        "help",
	Description: "Show the application's help",
	Value:       &Bool{},
}
