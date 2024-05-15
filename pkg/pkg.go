package pkg

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"

	"github.com/mo-crystal/mpkg/config"
)

func runCommands(commands []string) {
	for _, v := range commands {
		if err := exec.Command(v).Run(); err != nil {
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

	cp := "cp"
	if runtime.GOOS == "windows" {
		cp = "copy"
	}

	mkdir := "mkdir"
	if runtime.GOOS == "windows" {
		mkdir = "md"
	}

	for _, headerFile := range p.Headers {
		for _, includeDir := range config.IncludeDir {
			if err := exec.Command(mkdir, includeDir+string(os.PathSeparator)+p.Name); err != nil {
				panic(err)
			}
			if err := exec.Command(cp, headerFile, includeDir+string(os.PathSeparator)+p.Name+string(os.PathSeparator)); err != nil {
				panic(err)
			}
		}
	}

	for _, libFile := range p.Library {
		for _, libDir := range config.LibDir {
			if err := exec.Command(cp, libFile, libDir+string(os.PathSeparator)); err != nil {
				panic(err)
			}
		}
	}
}
