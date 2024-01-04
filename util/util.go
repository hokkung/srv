package util

import (
	"os"
	"os/signal"
	"syscall"
)

func WaitForSignalToShutdown() {
	quit := make(chan os.Signal)
	
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	<-quit

}
