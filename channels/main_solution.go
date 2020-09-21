package channels

import (
	"fmt"
)

func declaredGoroutine(s string, c chan string) {
	c <- fmt.Sprintf("Starting : %d", s)
	for i := 0; i < 10; i++ {
		t := fmt.Sprintf("%s : %d", s, i)
		c <- t
	}
}

func Channels_Solution() {
	responses := make(chan string)
	go declaredGoroutine("Goroutine with concurrency", responses)
	go func() {
		for i := 0; i < 10; i++ {
			responses <- fmt.Sprintf("On the fly : %d", i)
		}
	}()

	for i := 0; i < 21; i++ {
		t := <-responses
		fmt.Println(fmt.Sprintf("Response : %s", t))
	}
	close(responses)
}
