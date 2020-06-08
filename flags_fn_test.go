package flags

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFlagsInit(t *testing.T) {
	flags := Flags{}
	appName := "appName"
	appDescription := "appDescription"

	flags.Init(appName, appDescription)
	assert.Equal(t, appName, flags.AppName)
	assert.Equal(t, appDescription, flags.AppDescription)
}

func TestFlagsWithCommands(t *testing.T) {
	flags := Flags{}
	cmd1 := Command{}
	cmd1Name := "cmdName1"

	cmd2Name := "cmdName2"
	cmd2 := Command{Name: cmd2Name}

	flags.WithCommands(&cmd1, &cmd2)
	cmd1.Name = cmd1Name

	assert.Equal(t, cmd1Name, flags.Commands[0].Name)
	assert.Equal(t, cmd2Name, flags.Commands[1].Name)
}

func TestFlagsWithOptions(t *testing.T) {
	flags := Flags{}
	opt1 := Option{}
	opt1Short := 'a'
	opt2Short := 'b'
	opt2 := Option{Short: opt2Short}

	flags.WithOptions(&opt1, &opt2)
	opt1.Short = opt1Short

	assert.Equal(t, opt1Short, flags.Options[0].Short)
	assert.Equal(t, opt2Short, flags.Options[1].Short)
}

func TestFlagsGetCalledCommand(t *testing.T) {
	flags := Flags{}
	cmdName := "build"
	cmd := Command{Name: cmdName}

	flags.WithCommands(&cmd)
	assert.Nil(t, flags.GetCalledCommand())

	cmd.Called = true
	calledCmd := flags.GetCalledCommand()
	assert.NotNil(t, calledCmd)
	assert.Equal(t, cmdName, calledCmd.Name)
}

func TestFlagsParseEmptyArgs(t *testing.T) {
	flags := Flags{}
	assert.NoError(t, flags.ParseArgs([]string{}, false))
}

func TestFlagsParseRootOption(t *testing.T) {
	option := &Option{Short: 't', Long: "test", Value: &Bool{}}
	flags := Flags{}
	flags.WithOptions(option)
	assert.NoError(t, flags.ParseArgs([]string{"-t"}, false))

	val, err := BoolValue(option)
	assert.NoError(t, err)
	assert.True(t, val)
}

func TestFlagsParseRootCommand(t *testing.T) {
	cmd := &Command{Name: "cmd"}
	flags := Flags{}
	flags.WithCommands(cmd)
	assert.NoError(t, flags.ParseArgs([]string{"cmd"}, false))
	assert.True(t, cmd.Called)
}

func TestFlagsParseBoolOptionWithNextArg(t *testing.T) {
	flags := Flags{}
	opt := NewBool("opt", 'o', "", false)

	flags.WithOptions(opt)
	assert.NoError(t, flags.ParseArgs([]string{"-o", "true"}, false))

	val, err := BoolValue(opt)
	assert.NoError(t, err)
	assert.True(t, val)

	flags.WithOptions(opt)
	assert.NoError(t, flags.ParseArgs([]string{"-o", "false"}, false))

	val, err = BoolValue(opt)
	assert.NoError(t, err)
	assert.False(t, val)

	flags.WithOptions(opt)
	assert.NoError(t, flags.ParseArgs([]string{"--opt", "true"}, false))

	val, err = BoolValue(opt)
	assert.NoError(t, err)
	assert.True(t, val)
}

func TestFlagsParseLongBoolOptionNoArgs(t *testing.T) {
	flags := Flags{}
	opt := NewBool("opt", 'o', "", false)

	flags.WithOptions(opt)
	assert.NoError(t, flags.ParseArgs([]string{"--opt"}, false))

	val, err := BoolValue(opt)
	assert.NoError(t, err)
	assert.True(t, val)
}

func TestFlagsParseLongOptionNoBool(t *testing.T) {
	flags := Flags{}
	opt := NewString("opt", 'o', "", "")
	expectedVal := "abc"

	flags.WithOptions(opt)
	assert.NoError(t, flags.ParseArgs([]string{"--opt", expectedVal}, false))

	val, err := StringValue(opt)
	assert.NoError(t, err)
	assert.Equal(t, expectedVal, val)
}

func TestFlagsParseMultipleShortOptions(t *testing.T) {
	opt1 := &Option{Short: 'a', Value: &Bool{}}
	opt2 := &Option{Short: 'm', Value: &Bool{}}
	opt3 := &Option{Short: 'z', Value: &String{}}
	opt4 := &Option{Short: 'k', Value: &Bool{}}

	opt3Value := "z value"

	flags := Flags{}
	flags.WithOptions(opt1, opt2, opt3)
	assert.NoError(t, flags.ParseArgs([]string{"-amz", opt3Value}, false))

	calledBoolOptions := []*Option{opt1, opt2}
	for _, opt := range calledBoolOptions {
		val, err := BoolValue(opt)
		assert.NoError(t, err)
		assert.True(t, val)
	}

	val, err := BoolValue(opt4)
	assert.NoError(t, err)
	assert.False(t, val)

	strVal, ok := opt3.Value.(*String)
	assert.True(t, ok)
	assert.Equal(t, opt3Value, strVal.Value)
}

func TestFlagsParseSubCommand(t *testing.T) {
	cmd := &Command{Name: "command"}
	subCmd := &Command{Name: "subcommand"}
	cmd.WithCommands(subCmd)

	flags := Flags{Commands: []*Command{cmd}}
	assert.NoError(t, flags.ParseArgs([]string{cmd.Name, subCmd.Name}, false))
	assert.True(t, cmd.Called)
	assert.True(t, subCmd.Called)
}

func TestFlagsParseSubOption(t *testing.T) {}

func TestFlagsParseUnkownCommand(t *testing.T) {
	flags := Flags{}
	assert.Error(t, flags.ParseArgs([]string{"unknown-command"}, false))
}

func TestFlagsParseUnkownOption(t *testing.T) {
	flags := Flags{}
	assert.Error(t, flags.ParseArgs([]string{"--unknown-option"}, false))
}

func TestFlagsParseUnknownOptionSubcommand(t *testing.T) {
	flags := Flags{}
	cmd := Command{Name: "cmd"}
	flags.WithCommands(&cmd)

	assert.Error(t, flags.ParseArgs([]string{"cmd", "--unknown-option"}, false))
}

func TestFlagsExpectedOptionValue(t *testing.T) {
	flags := Flags{}
	opt := &Option{Long: "str", Value: &String{}}
	flags.WithOptions(opt)
	assert.Error(t, flags.ParseArgs([]string{"--str"}, false))
}

func TestFlagsPrintMainHelp(t *testing.T) {
	testWriter := &testStringWriter{}

	flags := Flags{}
	flags.Init("AppName", "Application description.")
	flags.AppVersion = "1.0.0"

	rootOpt := NewBool("bool", 'b', "Bool value", false)
	rootCmd := &Command{Name: "cmd", Description: "This is the description"}

	flags.WithCommands(rootCmd)
	flags.WithOptions(rootOpt)

	expectedOutput := `AppName version 1.0.0

Application description.

Available options.

--bool		-b		Bool value (default value: "false")

Available commands.
Use --help {command} {subcommand} for details.

- cmd		This is the description
`

	flags.PrintHelpWithArgs([]string{"./app"}, testWriter)

	assert.Equal(t, expectedOutput, testWriter.Value)
}

func TestFlagsPrintSubCommandHelp(t *testing.T) {
	testWriter := &testStringWriter{}

	flags := Flags{}
	flags.Init("AppName", "Application description.")
	flags.AppVersion = "1.0.0"

	rootOpt := NewBool("bool", 'b', "Bool value", false)
	rootCmd := &Command{Name: "cmd-1", Description: "This is the description of cmd-1"}

	subCmd := &Command{Name: "cmd-2", Description: "This is the description of cmd-2"}
	subCmdOpt1 := NewString("str", 's', "String option example", "default")
	subCmdOpt2 := NewBool("bool", 'b', "Bool option example", false)
	subSubCmd := &Command{Name: "cmd-3", Description: "This is the description of cmd-3"}

	rootCmd.WithCommands(subCmd)
	subCmd.WithCommands(subSubCmd)
	subCmd.WithOptions(subCmdOpt1)
	subCmd.WithOptions(subCmdOpt2)

	flags.WithCommands(rootCmd)
	flags.WithOptions(rootOpt)

	expectedOutput := `AppName version 1.0.0

Application description.

Details for command: cmd-1 cmd-2

This is the description of cmd-2

Available options.

--str		-s		String option example (default value: "default")
--bool		-b		Bool option example (default value: "false")

Available commands.
Use --help {command} {subcommand} for details.

- cmd-3		This is the description of cmd-3
`

	flags.PrintHelpWithArgs([]string{"./app", "cmd-1", "cmd-2"}, testWriter)

	assert.Equal(t, expectedOutput, testWriter.Value)
}
