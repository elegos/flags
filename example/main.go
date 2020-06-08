package main

import (
	"fmt"

	"github.com/elegos/flags"
)

func main() {
	longText := "This is a description long enough to let it go on a new line: " +
		"this tests the capability of managing columns automatically."
	lipsum := "Lorem ipsum dolor sit amet, consectetur adipiscing elit, " +
		"sed do eiusmod tempor incididunt ut labore et dolore magna aliqua."
	optFlags := flags.Flags{}
	optFlags.Init("AppName", longText)
	optFlags.AppVersion = "0.0.1"

	// Define a bunch of commands
	buildCmd := &flags.Command{Name: "build", Description: lipsum}
	testCmd := &flags.Command{Name: "test", Description: "Do test stuff"}

	// Define a bunch of bool options
	debugOpt := flags.NewBool("debug", 'd', "Enable debug session", false)
	// Options might have an empty short, disabling it entirely
	longOpt := flags.NewBool("very-very-long-option", flags.EmptyShort, "", false)
	// Just like short syntax, an option might miss the Long one
	zOpt := flags.NewBool("", 'z', "Z factor", false)

	// Add some options to the "build" command
	verboseOpt := flags.NewBool("verbose", 'v', "Enable verbose otuput", false)
	dryOpt := flags.NewBool("dry", 'd', "Test the build environment, but do not compile", false)
	buildCmd.WithOptions(verboseOpt, dryOpt)

	// Add (root) options and (root) commands
	optFlags.WithOptions(debugOpt, longOpt, zOpt, flags.HelpOption)
	optFlags.WithCommands(buildCmd, testCmd)

	// parse the command's arguments
	err := optFlags.Parse(true)
	// This happens only if flags are mis-configured
	if err != nil {
		panic(err)
	}

	// 2nd value is error, only if the option is not of type Bool
	if isDebug, _ := flags.BoolValue(debugOpt); isDebug {
		fmt.Println("Debug mode on")
	}

	if zFactor, _ := flags.BoolValue(zOpt); zFactor {
		fmt.Println("Z factor. Party hard.")
	}
}
