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
		needReplace := [][]string{
			pkg.Packages[i].Download,
			pkg.Packages[i].Update,
			pkg.Packages[i].Build,
			pkg.Packages[i].Install,
			pkg.Packages[i].Headers,
			pkg.Packages[i].Library,
			pkg.Packages[i].Binary,
		}

		for i1 := range needReplace {
			for i2 := range needReplace[i1] {
				needReplace[i1][i2] = strings.ReplaceAll(needReplace[i1][i2], "${mocdir}", config.MocDir)
			}
		}
	}
}
