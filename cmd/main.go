package main

import (
	"os"
	"os/signal"
	"syscall"
	"time"
	"code.cloudfoundry.org/lager"
)

func main() {
	logger := lager.NewLogger("signals")
	logger.RegisterSink(lager.NewWriterSink(os.Stdout, lager.DEBUG))
	logger.Info("well, hello")
	var gracefulStop = make(chan os.Signal)
	signal.Notify(gracefulStop, syscall.SIGTERM)
	signal.Notify(gracefulStop, syscall.SIGINT)
	go func() {
		sig := <-gracefulStop
		logger.Info("caught sig: %+v", lager.Data{"signal": sig})
		logger.Info("Wait for 2 second to finish processing")
		time.Sleep(2 * time.Second)
		os.Exit(0)
	}()
	logger.Info("going to sleep")
	time.Sleep(9999 * time.Second)
}