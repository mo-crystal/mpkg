package processor

import (
	"fmt"

	"github.com/mo-crystal/mpkg/config"
)

func Version(commands []string) {
	fmt.Println("mpkg cli: " + config.Version)
}
