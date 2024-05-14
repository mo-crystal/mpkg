package initialize

import (
	"flag"

	"github.com/mo-crystal/mpkg/config"
)

func Server() {
	server := flag.String("s", "http://mpkg.takemeto.icu", "server address")
	flag.Parse()
	config.Server = *server
}
