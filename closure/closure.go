package closure

import (
	"runtime"
	"sync"
	"fmt"
)

/*
PROCS=1, A:10 foreach 10 times,because i address all the same
B:x different print

PROCS=2, A:x or A:10 most,because maybe some goroutine run first then i value diff, not be the end one
B:x different print totally


A: 6
A: 6
A: 7
A: 10
A: 10
A: 10
A: 10
B: 0
A: 10
A: 10
A: 10
B: 0
B: 1
B: 5
B: 3
B: 6
B: 9
B: 7
B: 8
B: 4
*/

func Closure() {
	runtime.GOMAXPROCS(2)
	wg := sync.WaitGroup{}
	wg.Add(20)

	for i := 0; i < 10; i++ {
		go func() {
			fmt.Println("A:", i)
			wg.Done()
		}()
	}

	for j := 0; j < 10; j++ {
		go func(i int) {
			fmt.Println("B:", i)
			wg.Done()
		}(j)
	}
	wg.Wait()
}
