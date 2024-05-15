package pkg

import (
	"fmt"
	"os"
	"os/exec"
	"strings"

	"github.com/mo-crystal/mpkg/config"
	"github.com/mo-crystal/mpkg/utils"
)

func runCommands(commands []string) {
	for _, v := range commands {
		args := strings.Split(v, " ")
		if err := exec.Command(args[0], args[1:]...).Run(); err != nil {
			panic(err)
		}
	}
}

type Package struct {
	Name     string
	Download []string
	Update   []string
	Build    []string
	Install  []string
	Headers  []string
	Library  []string
	Binary   []string
}

func (p *Package) PrintInfo() {
	fmt.Println("Name: " + p.Name)
	fmt.Println("Commands:")
	fmt.Println("  Download: ", p.Download)
	fmt.Println("  Update  : ", p.Update)
	fmt.Println("  Build   : ", p.Build)
	fmt.Println("  Install : ", p.Install)
	fmt.Println("Files:")
	fmt.Println("  Headers: ", p.Headers)
	fmt.Println("  Linrary: ", p.Library)
	fmt.Println("  Binary : ", p.Binary)
}

func (p *Package) Cache() {
	runCommands(p.Download)
}

func (p *Package) Cached() bool {
	_, err := os.Stat(config.MocDir + string(os.PathSeparator) + p.Name)
	if err == nil {
		return true
	}
	return os.IsExist(err)
}

func (p *Package) InstallToEnv() {
	if !p.Cached() {
		p.Cache()
	}

	if !p.Cached() {
		panic("cannot download package: " + p.Name)
	}

	runCommands(p.Build)

	for _, headerFile := range p.Headers {
		for _, includeDir := range config.IncludeDir {
			os.Mkdir(includeDir+string(os.PathSeparator)+p.Name, os.ModePerm)
			utils.Copy(includeDir+string(os.PathSeparator)+p.Name+string(os.PathSeparator), headerFile)
		}
	}

	for _, libFile := range p.Library {
		for _, libDir := range config.LibDir {
			utils.Copy(libDir+string(os.PathSeparator), libFile)
		}
	}
}
