package processor

import (
	"fmt"
	"os"

	"github.com/mo-crystal/mpkg/config"
)

func Reset() {
	fmt.Println("Removing all mpkg files...")
	err := os.Remove(config.MocDir)
	if err != nil {
		panic(err)
	}

	fmt.Println("done")
}
