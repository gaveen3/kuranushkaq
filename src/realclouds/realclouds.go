package main

import (
	"os"
	"runtime"

	"website"
)

func init() {
	if err :=
		os.Setenv("LANG", "C"); nil != err {
		website.LogFatalln(err)
	}
}

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	website.Run()
}
