package initialize

import (
	"os"
	"os/user"
	"runtime"

	"github.com/mo-crystal/mpkg/config"
)

func RootDir() {
	u, err := user.Current()
	if err != nil {
		panic(err)
	}

	config.MocDir = u.HomeDir + string(os.PathSeparator) + "moc"
}

func IncludeDir() {
	if runtime.GOOS == "windows" {
		config.IncludeDir = []string{config.MocDir + string(os.PathSeparator) + "include"}
	} else {
		config.IncludeDir = []string{"/usr/include"}
	}
}

func LibDir() {
	if runtime.GOOS == "windows" {
		config.LibDir = []string{config.MocDir + string(os.PathSeparator) + "lib"}
		return
	}

	arch := runtime.GOARCH
	if arch == "amd64" {
		arch = "x86_64"
	}

	contents, err := os.ReadDir("/usr/lib/gcc/" + arch + "-linux-gnu")
	if err != nil {
		panic(err)
	}

	for _, file := range contents {
		if file.IsDir() {
			config.LibDir = append(config.LibDir, "/usr/lib/gcc/"+arch+"-linux-gnu/"+file.Name())
		}
	}
}
