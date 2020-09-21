package goroutines

import "fmt"

func declaredGoroutine(s string) {
	for i := 0; i < 10; i++ {
		fmt.Println(fmt.Sprintf("%s : %d", s, i))
	}
}

func runFunctions(concurrencyEnabled bool) {
	if concurrencyEnabled {
		go declaredGoroutine("Goroutine with concurrency")
		go func() {
			for i := 0; i < 10; i++ {
				fmt.Println(fmt.Sprintf("On the fly : %d", i))
			}
		}()
	} else {
		declaredGoroutine("Goroutine without concurrency")
		func() {
			for i := 0; i < 10; i++ {
				fmt.Println(fmt.Sprintf("On the fly : %d", i))
			}
		}()
	}
}

func Goroutines_Solution() {
	runFunctions(false)
	runFunctions(true)
}
