package selectx

import (
	"runtime"
	"fmt"
)

func Select() {
	runtime.GOMAXPROCS(1)
	intChan := make(chan int, 1)
	strChan := make(chan string, 1)

	intChan <- 1
	strChan <- "hello"

	select {
	case i := <-intChan:
		fmt.Printf("int chan:%d", i)
	case strChan := <-strChan:
		panic(strChan)
	}
}

func Default() {
	intChan := make(chan int, 1)

	go func() {
		intChan <- 10
	}()
	select {
	case i := <-intChan:
		fmt.Println(i)
	default:
		fmt.Println("default call")
	}

}

func ReturnFirst() {
	intChan := make(chan int, 1)
	strChan := make(chan string, 1)
	go func() {
		strChan <- "hello"
	}()
	select {
	case intChan <- 10:
		fmt.Println("return int")
	case str := <-strChan:
		fmt.Println(str)
	}
}
