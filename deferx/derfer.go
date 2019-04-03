package deferx

import "fmt"

func ParseDefer()  {
	defer func() {
		fmt.Println("call one")
	}()

	defer func() {
		fmt.Println("call two")
	}()

	defer func() {
		fmt.Println("call three")
	}()

	panic("panic here")
}
