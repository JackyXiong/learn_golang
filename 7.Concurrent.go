package main

import (
	"fmt"
	"runtime"
)

func say(msg string) {
	for i := 0; i < 5; i++ {
		runtime.Gosched()
		fmt.Println(msg, i)
	}
}

func test1() {
	go say("tom")
	say("hello")
}

func test2() {
	ch := make(chan int, 3) // 双向通道。chan<- 只能接受数据，<-chan 只能发送数据
	fmt.Println(len(ch))    // 0
	ch <- 1
	ch <- 2
	fmt.Println(len(ch)) // 2
	fmt.Println(cap(ch)) // 3
	ch <- 3
	// ch <- 4 // goroutine 1 [chan send]:
	fmt.Println(<-ch) // 1
	close(ch)
	// ch <- 1
	v, ok := <-ch
	fmt.Println(v, ok)
	fmt.Println(<-ch) // 1
	fmt.Println(<-ch) // 1
	// c := make(<-chan int, 2)
	// close(c)
}

func main() {
	// test1()
	test2()
}
