package main

import (
	"cc/internal/config"
	"cc/internal/logger"
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
)

var log = logger.Logger{TAG: "cc"}

func main() {

	args := parseArgs()

	config.Init(args.configFile)
	logger.Init(config.PathLogFile("cc"), args.console, args.verboseLevel, args.debug)

	state := state{}
	state.init(args.symbols)

	f, err := os.OpenFile(config.PathLogFile("cc.web"), os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatal(fmt.Sprintf("Failed to open log file : %s", err.Error()))
	}

	gin.DefaultWriter = f

	router := gin.Default()

	router.GET("/currency/:symbol", getCurrencyHandler(state))

	router.Run(fmt.Sprintf(":%d", args.serverPort))
}
