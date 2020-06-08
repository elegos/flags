# flags

[![Build Status](https://travis-ci.org/elegos/flags.svg?branch=master)](https://travis-ci.org/elegos/flags)

Flags is partially an alternative, partially an extension to the golang's standard `flags` package.

**Flags** aims to allow a more unix-style development of application's arguments, providing the following features:

- Double dash (`--option`) for long syntax options
- Single dash (`-o`) for short syntax options
- Possibility to accumulate short syntax options (like `-opx`)
- An automatic and opt-in `--help` (or `-h`) common print (column-based, including app name, version, description. See examples)

## Detailed information

This package divides the arguments in two different types:

- `Command`, used to tell the application to run a specific runtime
- `Option`, used to setup the command's workflow

The most common usage of a CLI command is as follows:

```bash
my-binary --option 123 -xy
```

This application thus uses three different (root) options (`--option`, `-x` and `-y`), represented by a long and two short syntaxes.

Another possibility is to tell the application to run a specific runtime (or command):

```bash
my-binary convert --convert-option 1
```

In this case the application will call the `convert` runtime with the `--convert-option` option set to `1`.

With this library you can combine options of commands, subcommands and root options, too. Here is an example:

```bash
my-binary --verbose convert --convert-option 1
```

This will turn the (root) option `--verbose` on, select the `convert` command and set the command's `--convert-option` option to `1`.

The `flags` object accepts both a series of `Command`s and `Option`s, and commands do the same.

See [examples/main.go](examples/main.go) for a complete example.

## Usage

```golang
import (
  "github.com/elegos/flags"
)

// Setup the flags object
flag := flags.Flags{}
flag.Init("app name", "app description")
flag.AppVersion = "1.0.0"

// Add the "help" helper to show the help message
// when --help or -h is passed (root option). This will
// also handle the sub-commands help pages (i.e. --help command)
flag.WithOptions(flags.HelpOption)

// create some options
opt1 := flags.NewBool("bool", 'b', "description", false)
opt2 := flags.NewString("str", 's', "description", "")

// one of the two syntaxes is optional
// flags.EmptyShort is an alias of empty rune type
opt3 := flags.NewBool("bool2", flags.EmptyShort, "no short syntax", false)
opt4 := flags.NewBool("", 'p', "no long syntax", false)

// add them to the application's root
flag.WithOptions(opt1, opt2)

// create a command with an option
cmd1 := &flags.Command{Name: "command"}
cmd1opt1 := flags.NewFloat64("float", 'f', "description", 0.0)

cmd1.WithOptions(cm1opt1)

// add the command to the application's root
flag.WithCommands(cmd1)

// parse the flags
err := flag.Parse(true)
if err != nil {
  panic(err)
}

// value extractors
isOpt1True := flags.BoolValue(opt1)
opt2Value := flags.StringValue(opt2)
cmd1opt1Value := flags.Float64Value(cmd1opt1)

// continue with application's workflow
```

## Option types (out of the box)

This is the series of option types and option builders you can use out of the box (see [option_values.go](option_values.go)):

- `flags.Bool` via `flags.NewBool`
- `flags.Int` via `flags.NewInt`
- `flags.Int64` via `flags.NewInt64`
- `flags.Float32` via `flags.NewFloat32`
- `flags.Float64` via `flags.NewFloat64`
- `flags.String` via `flags.NewString`
- `flags.Uint` via `flags.NewUint`
- `flags.Uint64` via `flags.NewUint64`

## Option types extension

If you need a particular option type, you can easily create a new one. It MUST adhere to the [`flags.Value`](flags.go) interface. Option values SHOULD have a builder function to init an `*flags.Option` and a value getter to easily get the option (see [option_values.go](option_values.go) and [option_values_fn.go](option_values_fn.go)), as follows:

```golang
func NewTypeStruct(long string, short rune, description string, defaultValue Type) *flags.Option {
  return &Option{
    Long:        long,
    Short:       short,
    Description: description,
    Value:       &TypeStruct{ DefaultValue: defaultValue }
  }
}

func TypeStructValue(option *Option) (Type, error) {
	if value, ok := option.Value.(*TypeStruct); ok {
		if setCondition {
			return value.Value, nil
		}

		return value.DefaultValue, nil
	}

	return typeDefault, fmt.Errorf("Not a Type option")
}
```
