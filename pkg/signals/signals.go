package signals

import (
	"os"
	"os/signal"
	"syscall"
)

func InitSignals() {
	// handle ^c (os.Interrupt)
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	signal.Notify(c, syscall.SIGTERM)
	go func() {
		<-c
		os.Exit(1)
	}()
}
