package flags

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCommandWithOptions(t *testing.T) {
	opt1 := NewBool("opt1", EmptyShort, "", false)
	opt2 := NewString("opt2", EmptyShort, "", "")

	cmd := Command{}
	cmd.WithOptions(opt1, opt2)

	assert.Len(t, cmd.Options, 2)
}

func TestCommandWithCommands(t *testing.T) {
	cmd := Command{}
	cmd.WithCommands(&Command{}, &Command{})

	assert.Len(t, cmd.SubCommands, 2)
}
