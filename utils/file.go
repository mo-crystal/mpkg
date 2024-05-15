package utils

import (
	"io"
	"os"
	"strings"
)

func Copy(dst, src string) {
	srcF, err := os.Open(src)
	if err != nil {
		panic(err)
	}
	defer srcF.Close()

	fileName := srcF.Name()
	split := strings.Split(fileName, "/")
	fileName = split[len(split)-1]
	split = strings.Split(fileName, string(os.PathSeparator))
	fileName = split[len(split)-1]

	dstF, err := os.OpenFile(dst+string(os.PathSeparator)+fileName, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, os.ModePerm)
	if err != nil {
		panic(err)
	}
	defer dstF.Close()

	_, err = io.Copy(dstF, srcF)
	if err != nil {
		panic(err)
	}
}
