package cmds

import (
	"flag"
	"fmt"
	"os"
	"strings"
)

type Command struct {
	Run func(cmd *Command, args []string)

	Usage string

	Short string

	Long string

	Flag flag.FlagSet
}

func (c *Command) Name() string {
	name := c.Usage
	i := strings.Index(name, " ")
	if i >= 0 {
		name = name[:i]
	}
	return name
}

func (c *Command) UsageExit() {
	fmt.Fprintf(os.Stderr, "Usage: microdef %s\n\n", c.Usage)
	fmt.Fprintf(os.Stderr, "Run 'microdef help %s' for help.\n", c.Name())
	os.Exit(2)
}
