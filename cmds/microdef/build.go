package microdef

import (
	"fmt"
	"github.com/microdef/cmds"
)

var CmdBuild = &cmds.Command{
	Usage: "build",
	Short: "builds this micro service",
	Long: `
Long description of how this command works!		
	`,
	Run: runBuild,
}

func runBuild(cmd *cmds.Command, args []string) {
	fmt.Println("Running the BUILD command!!!")
}
