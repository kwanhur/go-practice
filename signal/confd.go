package main

import "time"

type Processor struct {
	stopChan chan bool
	doneChan chan bool
}

func (p *Processor) process() {
	defer close(p.doneChan)
	for {
		time.Sleep(time.Second * 1)
	}
}
