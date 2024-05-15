package main

import (
	"flag"

	"github.com/mo-crystal/mpkg/initialize"
	"github.com/mo-crystal/mpkg/processor"
	"github.com/mo-crystal/mpkg/utils"
)

func main() {
	initialize.RootDir()
	initialize.Moc()
	initialize.Server()
	initialize.IncludeDir()
	initialize.LibDir()
	initialize.Packages()

	commands := flag.Args()

	if len(commands) < 1 {
		utils.Exit("no command")
	}

	switch commands[0] {
	case "version":
		processor.Version(commands[1:])
	case "search":
		processor.Search(commands[1:])
	case "info":
		processor.Info(commands[1:])
	case "reset":
		processor.Reset()
	default:
		utils.Exit("unknown command: " + commands[0])
	}
}
