package main

import (
	"flag"
	"fmt"
	"os"
)

var version = "master"
var debugMode bool

type serverConfig struct {
	port        int
	credentials []string
	debug       bool
	browser     bool
}

func errorAndExit(err error, code int) {
	fmt.Println(fmt.Errorf("Error: %v", err))
	os.Exit(code)
}

func main() {
	fmt.Println("Search Plugins", "v"+version, "by Doğan Çelik (dogancelik.com)")
	fmt.Println("Starting Search Plugins server…")

	portPtr := flag.Int("p", 8080, "server port")
	credPtr := flag.String("c", "admin:123456", "admin username & password")
	debugPtr := flag.Bool("d", false, "disable or enable logging")
	browserPtr := flag.Bool("o", false, "open browser")
	flag.Parse()

	err, creds := splitCreds(*credPtr)
	if err != nil {
		errorAndExit(err, 1)
	}

	debugMode = *debugPtr
	cfg := serverConfig{*portPtr, creds, debugMode, *browserPtr}

	initWeb()
	startServer(cfg)
}
