package main

import (
	"cc/internal/config"
	"cc/internal/logger"
	"context"
	"fmt"
	"net/http"
	"os"
	"time"

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

	server := &http.Server{
		Addr:    fmt.Sprintf(":%d", args.serverPort),
		Handler: router,
	}

	go func() {
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatal(err.Error())
		}
	}()

	<-state.closeChan

	log.Info("stopping server")

	// The context is used to inform the server it has 5 seconds to finish
	// the request it is currently handling
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := server.Shutdown(ctx); err != nil {
		log.Fatal(fmt.Sprintf("Server forced to shutdown: %s", err.Error()))
	}

	log.Info("Server exiting")
}
