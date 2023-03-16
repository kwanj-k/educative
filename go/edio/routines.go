package main

import (
	"fmt"
	"testing"
)

func main() {
	fmt.Println(" sync", testing.Benchmark(BenchmarkChannelSync).String())
	fmt.Println("buffered", testing.Benchmark(BenchmarkChannelBuffered).String())
}

func BenchmarkChannelSync(b *testing.B) { // makes buffered channel
	ch := make(chan int)
	go func() {
		for i := 0; i < b.N; i++ {
			ch <- i
		}
		close(ch)
	}()
	for _ = range ch { // iterating over channel without doing anything
	}
}

func BenchmarkChannelBuffered(b *testing.B) { // makes buffered channel with capacity of 128
	ch := make(chan int, 128)
	go func() {
		for i := 0; i < b.N; i++ {
			ch <- i
		}
		close(ch)

	}()
	for _ = range ch {
	}
}

//package main
//
//import (
//	"flag"
//	"fmt"
//)
//
//var ngoroutine = flag.Int("n", 100000, "how many goroutines")
//
//func f(left, right chan int) { left <- 1 + <-right }
//
//func main() {
//	flag.Parse()
//	leftmost := make(chan int)
//	var left, right chan int = nil, leftmost
//	for i := 0; i < *ngoroutine; i++ {
//		left, right = right, make(chan int)
//		go f(left, right)
//	}
//	right <- 0      // start the chaining
//	x := <-leftmost // wait for completion
//	fmt.Println(x)  // 100000, approx. 1.5 s
//}

//package main
//
//import (
//	"fmt"
//)
//
//var resume chan int
//
//func integers() chan int {
//	yield := make(chan int)
//	count := 0
//	go func() {
//		for {
//			yield <- count
//			count++
//		}
//	}()
//	return yield
//}
//
//func generateInteger() int {
//	return <-resume
//}
//
//func main() {
//	resume = integers()
//	fmt.Println(generateInteger()) //=> 0
//	fmt.Println(generateInteger()) //=> 1
//	fmt.Println(generateInteger()) //=> 2
//}

//package main
//
//import "fmt"
//import "os"
//import "strconv"
//
//func iter(b int, c chan int) {
//	for i := 0; i < b; i++ {
//		c <- i // put integers from 0 to n-1 on channel c
//	}
//	close(c)
//}
//
//func main() {
//	n, _ := strconv.Atoi(os.Args[1]) // line arg. converted to integer
//	a := make(chan int)
//	b := make(chan int)
//	go iter(n, a)
//	go iter(n, b)
//	for a != nil || b != nil {
//		select {
//		case x, ok := <-a: // takes int value from channel a
//			if ok {
//				go fmt.Println(x)
//			} else { // if channel a is closed
//				a = nil
//			}
//		case y, ok := <-b: // takes int value from channel b
//			if ok {
//				fmt.Println(y)
//			} else { // if channel b is closed
//				b = nil
//			}
//		}
//	}
//}

//package main
//
//import (
//	"fmt"
//	"runtime"
//	"time"
//)
//
//func main() {
//	runtime.GOMAXPROCS(2)
//	ch1 := make(chan int)
//	ch2 := make(chan int)
//	go pump1(ch1)
//	go pump2(ch2)
//	go suck(ch1, ch2)
//	time.Sleep(1e9)
//}
//
//func pump1(ch chan int) {
//	for i := 0; ; i++ {
//		ch <- i * 2
//	}
//}
//
//func pump2(ch chan int) {
//	for i := 0; ; i++ {
//		ch <- i + 5
//	}
//}
//
//func suck(ch1 chan int, ch2 chan int) {
//	for {
//		select {
//
//		case v := <-ch1:
//			fmt.Printf("Received on channel 1: %d\n", v)
//		case v := <-ch2:
//			fmt.Printf("Received on channel 2: %d\n", v)
//		}
//	}
//}

//package main
//
//import "fmt"
//
//func main() {
//	ch := make(chan string)
//	go sendData(ch)
//	getData(ch)
//}
//
//func sendData(ch chan string) {
//	ch <- "Washington"
//	ch <- "Tripoli"
//	ch <- "London"
//	ch <- "Beijing"
//	ch <- "Tokyo"
//	close(ch)
//}
//
//func getData(ch chan string) {
//	for {
//		input, open := <-ch
//		if !open {
//			break
//		}
//		fmt.Printf("%s ", input)
//	}
//}

//package main
//
//import (
//	"fmt"
//	"time"
//)
//
//func main() {
//	ch1 := make(chan int)
//	go pump(ch1)       // pump hangs
//	fmt.Println(<-ch1) // prints only 0
//	time.Sleep(1e9)
//}
//
//func pump(ch chan int) {
//	for i := 0; ; i++ {
//		ch <- i
//	}
//}

//package main
//
//import (
//	"fmt"
//	"time"
//)
//
//func main() {
//	ch := make(chan string)
//	go sendData(ch) // calling goroutine to send the data
//	go getData(ch)  // calling goroutine to receive the data
//	time.Sleep(1e9)
//}
//
//func sendData(ch chan string) { // sending data to ch channel
//	ch <- "Washington"
//	ch <- "Tripoli"
//	ch <- "London"
//	ch <- "Beijing"
//	ch <- "Tokyo"
//}
//
//func getData(ch chan string) {
//	var input string
//	for {
//		input = <-ch // receiving data sent to ch channel
//		fmt.Printf("%s ", input)
//	}
//	close(ch) // closed the channel
//}

//package main
//
//import (
//	"fmt"
//	"sync"
//)
//
//func HeavyFunction1(wg *sync.WaitGroup) {
//	defer wg.Done()
//	// Do a lot of stuff
//}
//
//func HeavyFunction2(wg *sync.WaitGroup) {
//	defer wg.Done()
//	// Do a lot of stuff
//}
//
//func main() {
//	wg := new(sync.WaitGroup)
//	wg.Add(2)
//	go HeavyFunction1(wg)
//	go HeavyFunction2(wg)
//	wg.Wait()
//	fmt.Printf("All Finished!")
//}

//package main
//
//import (
//	"fmt"
//	"time"
//)
//
//func main() {
//	fmt.Println("In main()")
//	go longWait()
//	go shortWait()
//	fmt.Println("About to sleep in main()")
//	time.Sleep(10 * 1e9) // sleep works with a Duration in nanoseconds (ns) !
//	fmt.Println("At the end of main()")
//}
//
//func longWait() {
//	fmt.Println("Beginning longWait()")
//	time.Sleep(5 * 1e9) // sleep for 5 seconds
//	fmt.Println("End of longWait()")
//}
//
//func shortWait() {
//	fmt.Println("Beginning shortWait()")
//	time.Sleep(2 * 1e9) // sleep for 2 seconds
//	fmt.Println("End of shortWait()")
//}
