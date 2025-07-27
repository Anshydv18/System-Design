package main

import (
	"fmt"
	"sync"
)

type singleton struct {
}

var lock sync.Mutex

var instance *singleton

func GetInstance() *singleton {

	if instance == nil {
		lock.Lock()
		defer lock.Unlock()
		if instance == nil {
			fmt.Println("Creating new instance of singleton.")
			instance = &singleton{}
		} else {
			fmt.Println("Instance already created, returning existing instance 1.")
		}
	} else {
		fmt.Println("Instance already created, returning existing instance 2.")
	}
	return instance
}

func main() {
	wg := &sync.WaitGroup{}
	wg.Add(30)
	for i := 0; i < 30; i++ {
		wg.Done()
		go GetInstance()
		fmt.Printf("Goroutine %d started.\n", i+1)
	}

	wg.Wait()
	fmt.Println("All goroutines finished execution.")
}
