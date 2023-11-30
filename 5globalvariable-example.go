package main

import "fmt"

var activeUserCount int

// your code goes here

func entry() {
	activeUserCount++
}

func exit() {
	activeUserCount--
}

func main() {
	entry()
	entry()
	exit()
	exit()
	entry()
	entry()
	fmt.Println(activeUserCount)
}
