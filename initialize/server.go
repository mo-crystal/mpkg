package initialize

import (
	"encoding/json"
	"flag"
	"strings"

	"github.com/mo-crystal/mpkg/config"
	"github.com/mo-crystal/mpkg/pkg"
	"github.com/mo-crystal/mpkg/utils"
)

func Server() {
	server := flag.String("s", "http://mpkg.takemeto.icu", "server address")
	flag.Parse()
	config.Server = *server
}

func Packages() {
	resp := utils.Get("/list")

	pkgJson, err := json.Marshal(resp.Data)
	if err != nil {
		panic(err)
	}

	err = json.Unmarshal(pkgJson, &pkg.Packages)
	if err != nil {
		panic(err)
	}

	for i := range pkg.Packages {
		for i1 := range pkg.Packages[i].Download {
			pkg.Packages[i].Download[i1] = strings.ReplaceAll(pkg.Packages[i].Download[i1], "${mocdir}", config.MocDir)
		}

		for i1 := range pkg.Packages[i].Update {
			pkg.Packages[i].Update[i1] = strings.ReplaceAll(pkg.Packages[i].Update[i1], "${mocdir}", config.MocDir)
		}

		for i1 := range pkg.Packages[i].Build {
			pkg.Packages[i].Build[i1] = strings.ReplaceAll(pkg.Packages[i].Build[i1], "${mocdir}", config.MocDir)
		}

		for i1 := range pkg.Packages[i].Install {
			pkg.Packages[i].Install[i1] = strings.ReplaceAll(pkg.Packages[i].Install[i1], "${mocdir}", config.MocDir)
		}

		for i1 := range pkg.Packages[i].Headers {
			pkg.Packages[i].Headers[i1] = strings.ReplaceAll(pkg.Packages[i].Headers[i1], "${mocdir}", config.MocDir)
		}

		for i1 := range pkg.Packages[i].Library {
			pkg.Packages[i].Library[i1] = strings.ReplaceAll(pkg.Packages[i].Library[i1], "${mocdir}", config.MocDir)
		}

		for i1 := range pkg.Packages[i].Binary {
			pkg.Packages[i].Binary[i1] = strings.ReplaceAll(pkg.Packages[i].Binary[i1], "${mocdir}", config.MocDir)
		}
	}
}
