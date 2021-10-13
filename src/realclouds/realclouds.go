package main

import (
	"os"
	log "rclog"
	"runtime"
	"website"
)

func init() {
	if err :=
		os.Setenv("LANG", "C"); nil != err {
		log.Fatalln(err)
	}
}

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	website.Run()
}
