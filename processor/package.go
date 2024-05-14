package processor

import (
	"fmt"
	"strings"

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
