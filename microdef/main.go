package main

import (
	"flag"
	"fmt"
	"github.com/microdef/cmds"
	"github.com/microdef/cmds/microdef"
	"io"
	"os"
	"strings"
	"text/template"
)

var commands = []*cmds.Command{
	microdef.CmdBuild,
}

func main() {
	flag.Usage = usageExit
	flag.Parse()
	args := flag.Args()
	if len(args) < 1 {
		usageExit()
	}

	if args[0] == "help" {
		help(args[1:])
		return
	}

	for _, cmd := range commands {
		if cmd.Name() == args[0] {
			cmd.Flag.Usage = func() { cmd.UsageExit() }
			cmd.Run(cmd, cmd.Flag.Args())
			return
		}
	}

	fmt.Fprintf(os.Stderr, "microdef: unknown command %q\n", args[0])
	fmt.Fprintf(os.Stderr, "Run 'microdef help' for usage.\n")
	os.Exit(2)
}

var usageTemplate = `
Microdef is a tool for??? 

Usage:

	microdef command [arguments]

The commands are:
{{range .}}
    {{.Name | printf "%-8s"}} {{.Short}}{{end}}

Use "microdef help [command]" for more information about a command.
`

var helpTemplate = `
Usage: microdef {{.Usage}}

{{.Long | trim}}
`

func help(args []string) {
	if len(args) == 0 {
		printUsage(os.Stdout)
		return
	}
	if len(args) != 1 {
		fmt.Fprintf(os.Stderr, "usage: microdef help command\n\n")
		fmt.Fprintf(os.Stderr, "Too many arguments given.\n")
		os.Exit(2)
	}
	for _, cmd := range commands {
		if cmd.Name() == args[0] {
			tmpl(os.Stdout, helpTemplate, cmd)
			return
		}
	}
}

func usageExit() {
	printUsage(os.Stderr)
	os.Exit(2)
}

func printUsage(w io.Writer) {
	tmpl(w, usageTemplate, commands)
}

func tmpl(w io.Writer, text string, data interface{}) {
	t := template.New("top")
	t.Funcs(template.FuncMap{
		"trim": strings.TrimSpace,
	})
	template.Must(t.Parse(strings.TrimSpace(text) + "\n\n"))
	if err := t.Execute(w, data); err != nil {
		panic(err)
	}
}
