package main

import (
	"fmt"
	"strconv"
	"time"
)

//
// This is a sample program for the concurrent/multithread programming
// in Golang. It looks so far we do not need such tips but
// during your development if performance concerns requires either
// of concurrent/multithread programming then you can look here for
// example.
//

func testConcurrency() {
	var c chan string = make(chan string)

	go execProcess1(c)
	go execProcess2(c)

	var input string
	fmt.Scanln(&input)
}

func execProcess1(c chan string) {
  for i := 0; i < 10 ; i++ {
    fmt.Println("execProcess1: " + strconv.Itoa(i))
    c <- strconv.Itoa(i)
  }
}

func execProcess2(c chan string) {
  for i := 0; i < 10 ; i++ {
    msg := <- c
    fmt.Println("execProcess2: " + msg)
    time.Sleep(time.Second * 1)
  }
}
