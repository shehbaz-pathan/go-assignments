package main

import (
	"fmt"
	"time"
)

func printAlphabet(n int) {
	fmt.Printf("%c ",n+96)
}

func main() {
	var start, end int
	fmt.Println("Enter starting position")
	fmt.Scanf("%d",&start)
	fmt.Println("Enter ending position")
	fmt.Scanf("%d",&end)
	for i:=start;i<=end;i++ {
		go printAlphabet(i)
	}
	time.Sleep(100*time.Millisecond)
}
