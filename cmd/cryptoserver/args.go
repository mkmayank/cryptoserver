package main

import (
	util "cc/internal/utils"
	"flag"
	"fmt"
	"os"
	"strings"
)

type cmdArgs struct {
	configFile string

	verboseLevel int
	console      bool
	debug        bool

	serverPort int

	// this symbol list may be configured via rest api too
	symbols []string
}

func parseArgs() (args cmdArgs) {

	args.symbols = []string{"BTCUSDT", "ETHBTC"}

	flag.StringVar(&args.configFile, "c", "", "config file path")

	flag.BoolVar(&args.console, "v", false, "log on console")
	flag.IntVar(&args.verboseLevel, "vv", 0, "verbose logs level")
	flag.BoolVar(&args.debug, "debug", false, "log debug messages")

	flag.IntVar(&args.serverPort, "port", 8000, "server port to run this server")

	var symbolsFlag util.ArrayFlags
	flag.Var(&symbolsFlag, "s",
		fmt.Sprintf("symbols to subscribe for ticker, default: %s", strings.Join(args.symbols, " ")))

	flag.Parse()

	if len(symbolsFlag) != 0 {
		args.symbols = symbolsFlag
	}

	if args.serverPort == 0 {
		fmt.Println("server port cannot be 0")
		os.Exit(2)
	}
	return
}
