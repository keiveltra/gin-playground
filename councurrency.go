package main

import (
	"fmt"
	"strconv"
	"sync"
)

var mu sync.Mutex

func testConcurrency() {
	xs := [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	xs1 := xs[0:5]
	xs2 := xs[5:10]

	var wg sync.WaitGroup
	wg.Add(2)

	go execProcessWithLock(&wg, xs1)
	go execProcessWithLock(&wg, xs2)

	wg.Wait()

	fmt.Print("[")
	for i, val := range xs {
		if i > 0 {
			fmt.Print(", ")
		}
		fmt.Print(val)
	}
	fmt.Println("]")
}

func execProcessWithLock(wg *sync.WaitGroup, xs []int) {
	defer wg.Done()
        mu.Lock()
	for i := range xs {
		xs[i]++
		fmt.Println("execProcessWithLock: " + strconv.Itoa(xs[i]))
	}
        mu.Unlock()
}
