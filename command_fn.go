package flags

// WithOptions add multiple options at once
func (cmd *Command) WithOptions(opts ...*Option) {
	if cmd.Options == nil {
		cmd.Options = []*Option{}
	}

	cmd.Options = append(cmd.Options, opts...)
}

// WithCommands add multiple sub-commands at once
func (cmd *Command) WithCommands(cmds ...*Command) {
	if cmd.SubCommands == nil {
		cmd.SubCommands = []*Command{}
	}

	cmd.SubCommands = append(cmd.SubCommands, cmds...)
}
