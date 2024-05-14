package main

import (
	"flag"

	"github.com/mo-crystal/mpkg/initialize"
	"github.com/mo-crystal/mpkg/processor"
	"github.com/mo-crystal/mpkg/utils"
)

func main() {
	initialize.Server()
	initialize.RootDir()
	initialize.Packages()

	commands := flag.Args()

	if len(commands) < 1 {
		utils.Exit("no command")
	}

	switch commands[0] {
	case "version":
		processor.Version(commands[1:])
	default:
		utils.Exit("unknown command: " + commands[0])
	}
}
