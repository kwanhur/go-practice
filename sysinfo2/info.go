package sysinfo2

import (
	"github.com/elastic/go-sysinfo/providers/darwin"
)


func MachineID() (string,error) {
	return darwin.MachineID()
}