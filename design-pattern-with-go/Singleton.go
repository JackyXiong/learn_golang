package main

import (
	"fmt"
	"sync"
)

var (
	once     sync.Once
	instance *Singleton
)

type Singleton struct {
	a string
}

func New() *Singleton {
	fmt.Println("before", instance)
	once.Do(func() {
		instance = new(Singleton)
	})
	return instance
}

func main() {
	a := New()
	fmt.Println("a", a, *a, &a)
	fmt.Println(a.a)
	a.a = "test"
	b := New()
	fmt.Println("b", b, *b, &b)
	fmt.Println(b.a)

	c := make(map[string]int, 10)
	fmt.Println(c, &c)
}
