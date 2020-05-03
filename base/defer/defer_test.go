package deferT_test

import (
	"fmt"
	"testing"
	"time"
)

func TestDefer(t *testing.T) {
	doSomething()
}

func doSomething() {
	defer countTime("doSomething")()

	time.Sleep(3 * time.Second)
	fmt.Println("   done")
}

func countTime(msg string) func() {
	start := time.Now()
	fmt.Printf("run func:%s", msg)
	return func() {
		fmt.Printf("func name: %s run time: %f s \n", msg, time.Since(start).Seconds())
	}
}
