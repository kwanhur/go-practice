package foreach

import (
	"fmt"
)

type Student struct {
	Age  int
	Name string
}

/*
m before: map[]
stus : [{Age:10 Name:n1} {Age:12 Name:n2} {Age:13 Name:n3}]
stu:n1 ptr:0xc42000a180 val:{Age:10 Name:n1}
stu:n2 ptr:0xc42000a180 val:{Age:12 Name:n2}
stu:n3 ptr:0xc42000a180 val:{Age:13 Name:n3}
m after: map[n1:0xc420092120 n2:0xc420092120 n3:0xc420092120]
stu:n1 ptr:0xc420098140 val:{Age:10 Name:n1}
stu:n2 ptr:0xc420098158 val:{Age:12 Name:n2}
stu:n3 ptr:0xc420098170 val:{Age:13 Name:n3}
m rectify: map[n1:0xc420098140 n2:0xc420098158 n3:0xc420098170]
*/
func ParseStudent() {
	m := make(map[string]*Student)

	stus := []Student{
		{Name: "n1", Age: 10},
		{Name: "n2", Age: 12},
		{Name: "n3", Age: 13},
	}

	fmt.Printf("m before: %+v\n", m)
	fmt.Printf("stus : %+v\n", stus)
	for _, stu := range stus {
		// value copy but the same memory address
		fmt.Printf("stu:%s ptr:%p val:%+v\n", stu.Name, &stu, stu)
		m[stu.Name] = &stu
	}
	fmt.Printf("m after: %+v\n", m)

	// should write it like this
	for i := 0; i < len(stus); i++ {
		fmt.Printf("stu:%s ptr:%p val:%+v\n", stus[i].Name, &stus[i], stus[i])
		m[stus[i].Name] = &stus[i]
	}
	fmt.Printf("m rectify: %+v\n", m)
}
