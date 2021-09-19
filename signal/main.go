package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	stopChan := make(chan bool)
	doneChan := make(chan bool)

	proc := Processor{stopChan: stopChan, doneChan: doneChan}
	go proc.process()

	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, syscall.SIGINT, syscall.SIGTERM)
	for {
		select {
		//case err := <-errChan:
		//	log.Error(err.Error())
		case s := <-signalChan:
			log.Println(fmt.Sprintf("Captured %v. Exiting...", s))
			close(doneChan)
		case <-doneChan:
			os.Exit(0)
		}
	}
}
