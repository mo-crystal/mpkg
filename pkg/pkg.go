package pkg

import "fmt"

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
