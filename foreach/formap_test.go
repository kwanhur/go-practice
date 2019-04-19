package foreach

import "testing"

func TestParseStudent(t *testing.T) {
	ParseStudent()
}

func TestUpdateName(t *testing.T) {
	stus := []Student{
		{Name: "n1", Age: 10},
		{Name: "n2", Age: 12},
		{Name: "n3", Age: 13},
	}
	UpdateName(stus)
}
