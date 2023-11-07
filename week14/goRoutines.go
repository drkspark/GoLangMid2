package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	runtime.GOMAXPROCS(3)

	var processTest sync.WaitGroup

	processTest.Add(3)

	go func() {
		defer processTest.Done()

		for i := 0; i < 30; i++ {
			for j := 5; j <= 100; j++ {
				fmt.Printf(" %d", j)
				if j == 10 {
					fmt.Println()
				}
			}
		}

	}()

	go func() {
		defer processTest.Done()

		for j := 0; j < 100; j++ {
			for char := 'A'; char < 'A'+26; char++ {
				fmt.Printf("%c ", char)
				if char == 'Z' {
					fmt.Println()
				}
			}
		}

	}()

	go func() {
		defer processTest.Done()

		for i := 0; i < 30; i++ {
			for j := 0; j <= 50; j++ {
				fmt.Printf(" %d", j)
				if j == 50 {
					fmt.Println()
				}
			}
		}

	}()

	processTest.Wait()
}

/*
3. `runtime.GOMAXPROCS(3)`: This sets the maximum number of operating system threads that can simultaneously execute
							Go code to 3. This means that the program will utilize up to 3 CPU cores concurrently.

4. `var processTest sync.WaitGroup`: This declares a variable named `processTest` of type `sync.WaitGroup`.
									 `WaitGroup` is used to wait for a collection of goroutines to finish their work.

5. `processTest.Add(3)`: This increments the counter of the `WaitGroup` by 3. This indicates that there are 3 goroutines
						 that the `WaitGroup` will wait for.

6. Three anonymous functions are defined using the `go` keyword. Each of these functions represents a goroutine.
   - The first goroutine prints numbers from 5 to 100 and prints a new line after reaching 10. It does this 30 times.
   - The second goroutine prints uppercase letters from 'A' to 'Z' and prints a new line after reaching 'Z'. It does this 100 times.
   - The third goroutine prints numbers from 0 to 50 and prints a new line after reaching 50. It does this 30 times.

   Each of these goroutines is wrapped in a `defer processTest.Done()` statement, which decrements the `WaitGroup` counter when the goroutine completes its work.

7. `processTest.Wait()`: This statement blocks the main goroutine until the `WaitGroup` counter becomes zero. In other
						 words, it waits for all three goroutines to finish their work.

8. Once all goroutines are done, the program exits.
*/
