package initialize

import (
	"encoding/json"
	"flag"

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
}
