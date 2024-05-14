package utils

import (
	"fmt"
	"os"
)

func Exit(msg string) {
	fmt.Println(msg)
	os.Exit(1)
}
