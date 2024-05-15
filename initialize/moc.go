package initialize

import (
	"fmt"
	"os"
	"os/user"
	"runtime"

	"github.com/mo-crystal/mpkg/config"
)

func Moc() {
	user, err := user.Current()
	if err != nil {
		panic(err)
	}

	_, err = os.Stat(user.HomeDir + string(os.PathSeparator) + "moc")
	if err == nil {
		return
	}

	fmt.Println("Mpkg not initialized. Initializing...")
	fmt.Println("Making mpkg dir..")
	os.Mkdir(user.HomeDir+string(os.PathSeparator)+"moc", os.ModePerm)
	os.Mkdir(user.HomeDir+string(os.PathSeparator)+"moc"+string(os.PathSeparator)+"include", os.ModePerm)
	os.Mkdir(user.HomeDir+string(os.PathSeparator)+"moc"+string(os.PathSeparator)+"lib", os.ModePerm)
	fmt.Println("ok")

	if runtime.GOOS == "windows" {
		C_INCLUDE_PATH := os.Getenv("C_INCLUDE_PATH")
		if C_INCLUDE_PATH == "" {
			C_INCLUDE_PATH = config.MocDir + string(os.PathSeparator) + "include"
		} else {
			C_INCLUDE_PATH += ";" + config.MocDir + string(os.PathSeparator) + "include"
		}
		err = os.Setenv("C_INCLUDE_PATH", C_INCLUDE_PATH)
		if err != nil {
			panic(err)
		}

		CPLUS_INCLUDE_PATH := os.Getenv("CPLUS_INCLUDE_PATH")
		if CPLUS_INCLUDE_PATH == "" {
			CPLUS_INCLUDE_PATH = config.MocDir + string(os.PathSeparator) + "include"
		} else {
			CPLUS_INCLUDE_PATH += ";" + config.MocDir + string(os.PathSeparator) + "include"
		}
		err = os.Setenv("CPLUS_INCLUDE_PATH", CPLUS_INCLUDE_PATH)
		if err != nil {
			panic(err)
		}

		LIBRARY_PATH := os.Getenv("LIBRARY_PATH")
		if LIBRARY_PATH == "" {
			LIBRARY_PATH = config.MocDir + string(os.PathSeparator) + "lib"
		} else {
			LIBRARY_PATH += ";" + config.MocDir + string(os.PathSeparator) + "lib"
		}
		err = os.Setenv("LIBRARY_PATH", LIBRARY_PATH)
		if err != nil {
			panic(err)
		}
	}
}
