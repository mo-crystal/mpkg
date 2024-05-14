package initialize

import (
	"os"
	"os/user"

	"github.com/mo-crystal/mpkg/config"
)

func RootDir() {
	u, err := user.Current()
	if err != nil {
		panic(err)
	}

	config.MocDir = u.HomeDir + string(os.PathSeparator) + "moc"
}
