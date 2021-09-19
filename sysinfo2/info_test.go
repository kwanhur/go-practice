package sysinfo2

import "testing"

func TestMachineID(t *testing.T)  {
	id, err := MachineID()
	if err != nil {
		t.Error(err)
	}

	t.Log(id)
}