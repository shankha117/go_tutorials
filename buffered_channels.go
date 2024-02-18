package main

import (
	"fmt"
	"sync"
)

func main() {
	wg := &sync.WaitGroup{}
	// the 2nd parameter is the buffer . The initial size of the buffer is 0 .
	ch := make(chan int, 1)

	wg.Add(2)
	go func(ch <-chan int, wg *sync.WaitGroup) {
		/* this will always print the first inserted data
		as , we put the message 44 1st and 45(44+1) 2nd . WE RECEIVE THE 44 but we never
		receive 45 back
		*/

		if msg, ok := <-ch; ok {
			fmt.Println(msg)
			// fmt.Println(<-ch)
			// close(ch)
			// fmt.Println(<-ch)
		}
		wg.Done()

	}(ch, wg)

	go func(ch chan<- int, wg *sync.WaitGroup) {
		close(ch)
		// val := 0
		// ch <- val
		// fmt.Println("sending", val)
		wg.Done()
	}(ch, wg)

	wg.Wait()
}

/*
chnnel types are 3
	Bi-directional channel

		* Normal channels are Bi-directional channels
			ch := make(chan int)
			 ...
			func Some(ch chan int){...}

	Send-only channel

			func Some(ch chan <- int){...}

	Receive-only channel

			func Some(ch <- chan int){...}


*/
