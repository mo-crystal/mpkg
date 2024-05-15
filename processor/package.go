package processor

import (
	"fmt"
	"os"
	"strings"

	"github.com/mo-crystal/mpkg/config"
	"github.com/mo-crystal/mpkg/initialize"
	"github.com/mo-crystal/mpkg/pkg"
	"github.com/mo-crystal/mpkg/utils"
)

func Search(commands []string) {
	if len(commands) < 1 {
		utils.Exit("invalid arg")
	}

	for _, v := range pkg.Packages {
		if strings.Contains(v.Name, commands[0]) {
			fmt.Println(v.Name)
		}
	}
}

func Info(commands []string) {
	if len(commands) < 1 {
		utils.Exit("invalid arg")
	}

	for _, v := range pkg.Packages {
		if v.Name == commands[0] {
			v.PrintInfo()
		}
	}
}

func Update() {
	fmt.Println("Updating package cache...")
	os.Remove(config.MocDir + string(os.PathSeparator) + "packages.json")
	initialize.Packages()
	fmt.Println("done")
}
