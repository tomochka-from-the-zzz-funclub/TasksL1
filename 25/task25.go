package main

import (
	"fmt"
	"time"
)

func Sleep1(d time.Duration) {
	<-time.After(d)
}

func Sleep2(duration time.Duration) {
	start := time.Now()
	for {
		if time.Since(start) >= duration {
			break
		}
	}
}

func main() {
	fmt.Print("Before sleep1 function\n")
	Sleep1(5 * time.Second)
	fmt.Print("After sleep1 function\n")
	Sleep2(5 * time.Second)
	fmt.Print("After sleep2 function\n")

}
