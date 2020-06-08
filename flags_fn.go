package flags

import (
	"fmt"
	"io"
	"os"
	"regexp"
	"strings"
	"text/tabwriter"
)

// Init set the basic information
func (flags *Flags) Init(appName string, appDescription string) {
	flags.AppName = appName
	flags.AppDescription = appDescription
}

// WithCommands add a command to the main help object
func (flags *Flags) WithCommands(cmds ...*Command) {
	if flags.Commands == nil {
		flags.Commands = []*Command{}
	}

	flags.Commands = append(flags.Commands, cmds...)
}

// WithOptions add one or more options to the main help object
func (flags *Flags) WithOptions(opts ...*Option) {
	if flags.Options == nil {
		flags.Options = []*Option{}
	}

	flags.Options = append(flags.Options, opts...)
}

// GetCalledCommand Get the called command, if any
func (flags *Flags) GetCalledCommand() *Command {
	for _, cmd := range flags.Commands {
		if cmd.Called {
			return cmd
		}
	}

	return nil
}

// Parse parse the application's arguments
func (flags *Flags) Parse(printHelpOnError bool) error {
	return flags.ParseArgs(os.Args[1:], printHelpOnError) // first element is the app's name
}

// ParseArgs parse arbitrary arguments
func (flags *Flags) ParseArgs(args []string, printHelpOnError bool) error {
	// guard condition
	if len(args) == 0 {
		// Print the help if the option is triggered
		if HelpOption.Value.(*Bool).Value {
			flags.PrintHelp()
		}

		return nil
	}

	commands := flags.Commands
	options := flags.Options

	if flags.currentCommand != nil {
		commands = flags.currentCommand.SubCommands
		options = flags.currentCommand.Options
	}

	arg := args[0]
	argConsumed := false
	nextArg := ""
	nextArgConsumed := false

	if len(args) > 1 {
		nextArg = args[1]
	}

	if isOption(arg) {
		trueFalseRegexp := regexp.MustCompile("(?i)(true|false)")

		if isShortOption(arg) {
			// possibly multiple flags, i.e. -abc
			for i := 1; i < len(arg); i++ {
				subArg := rune(arg[i])
				canAccessNextArg := i == (len(arg) - 1)

				for _, option := range options {
					if option.Short == subArg {
						if option.Value.IsBoolValue() {
							if canAccessNextArg && trueFalseRegexp.MatchString(nextArg) {
								err := option.Value.Set(strings.ToLower(nextArg))
								if err != nil {
									return err
								}

								nextArgConsumed = true
								argConsumed = true

								break
							}

							if err := option.Value.Set("true"); err != nil {
								return err
							}

							argConsumed = true

							break
						}

						if !canAccessNextArg || nextArg == "" {
							if printHelpOnError {
								flags.PrintHelpWithArgs([]string{}, os.Stdout)
								os.Exit(1)
							}

							return fmt.Errorf("Option '%c' expects a value", subArg)
						}

						nextArgConsumed = true
						err := option.Value.Set(nextArg)

						if err != nil {
							return err
						}

						argConsumed = true

						break
					}
				}
			}

			if argConsumed {
				if nextArgConsumed {
					return flags.ParseArgs(args[2:], printHelpOnError)
				}

				return flags.ParseArgs(args[1:], printHelpOnError)
			}
		}

		// Long option
		argName, err := getOptionName(arg)
		if err != nil {
			if printHelpOnError {
				flags.PrintHelpWithArgs([]string{}, os.Stdout)
				os.Exit(1)
			}

			return err
		}

		for _, option := range options {
			if option.Long == argName {
				if option.Value.IsBoolValue() {
					if trueFalseRegexp.MatchString(nextArg) {
						if err := option.Value.Set(strings.ToLower(nextArg)); err != nil {
							return err
						}

						nextArgConsumed = true
						argConsumed = true

						break
					}

					if err := option.Value.Set("true"); err != nil {
						return err
					}

					argConsumed = true

					break
				}

				if nextArg == "" {
					if printHelpOnError {
						flags.PrintHelpWithArgs([]string{}, os.Stdout)
						os.Exit(1)
					}

					return fmt.Errorf("Option '%s' expects a value", arg)
				}

				if err := option.Value.Set(nextArg); err != nil {
					return err
				}

				nextArgConsumed = true
				argConsumed = true

				break
			}
		}

		if argConsumed {
			if nextArgConsumed {
				return flags.ParseArgs(args[2:], printHelpOnError)
			}

			return flags.ParseArgs(args[1:], printHelpOnError)
		}
	}

	// Not an Option, so it's a Command
	for _, command := range commands {
		if command.Name == arg {
			command.Called = true
			flags.currentCommand = command

			return flags.ParseArgs(args[1:], printHelpOnError)
		}
	}

	if printHelpOnError {
		flags.PrintHelpWithArgs([]string{}, os.Stdout)
		os.Exit(1)
	}

	return fmt.Errorf(`"%s" is not a registered command nor an option`, arg)
}

// PrintHelp print the help information
func (flags *Flags) PrintHelp() {
	flags.PrintHelpWithArgs(os.Args, os.Stdout)
}

// PrintHelpWithArgs print the help information
func (flags *Flags) PrintHelpWithArgs(args []string, output io.Writer) {
	var lastCommand *Command

	commandChain := []string{}
	commands := flags.Commands
	options := flags.Options

	for _, arg := range args {
		if isOption(arg) {
			continue
		}

		for _, command := range commands {
			if command.Name == arg {
				commandChain = append(commandChain, command.Name)
				lastCommand = command
				commands = command.SubCommands
				options = command.Options

				break
			}
		}
	}

	if flags.AppName != "" {
		fmt.Fprint(output, flags.AppName)

		if flags.AppVersion != "" {
			fmt.Fprintf(output, " version %s", flags.AppVersion)
		}

		fmt.Fprintln(output, "")
		fmt.Fprintln(output, "")
	}

	if flags.AppDescription != "" {
		lines, _ := textSplit(flags.AppDescription, 76)
		for _, line := range lines {
			fmt.Fprintln(output, line)
		}
	}

	if len(commandChain) > 0 {
		fmt.Fprintln(output, "")
		fmt.Fprintln(output, "Details for command: "+strings.Join(commandChain, " "))
		fmt.Fprintln(output, "")

		description, _ := textSplit(lastCommand.Description, 76)
		fmt.Fprintln(output, strings.Join(description, "\n"))
	}

	tabWriter := tabwriter.NewWriter(output, 7, 8, 7, '\t', 0)

	if len(options) > 0 {
		fmt.Fprintln(output, "")
		fmt.Fprintln(output, "Available options.")
		fmt.Fprintln(output, "")

		for _, opt := range options {
			description := opt.Description
			defaultValue := opt.Value.DefaultValueString()

			if defaultValue != "" {
				description = fmt.Sprintf(`%s (default value: "%s")`, description, defaultValue)
			}

			descriptionLines, _ := textSplit(description, 64)
			line := ""

			if opt.Long != "" {
				line = fmt.Sprintf("--%s", opt.Long)
			}

			if opt.Short != EmptyShort {
				line = fmt.Sprintf("%s\t-%c", line, opt.Short)
			} else {
				line = fmt.Sprintf("%s\t", line)
			}

			line = fmt.Sprintf("%s\t%s\n", line, strings.Join(descriptionLines, "\n\t\t"))

			fmt.Fprintf(tabWriter, "%s", line)
		}

		tabWriter.Flush()
	}

	if len(commands) > 0 {
		fmt.Fprintln(output, "")
		fmt.Fprintln(output, "Available commands.")
		fmt.Fprintln(output, "Use --help {command} {subcommand} for details.")
		fmt.Fprintln(output, "")

		for _, cmd := range commands {
			descriptionLines, _ := textSplit(cmd.Description, 64)
			fmt.Fprintf(tabWriter, "- %s\t%s\n", cmd.Name, strings.Join(descriptionLines, "\n\t"))
		}

		tabWriter.Flush()
	}
}
