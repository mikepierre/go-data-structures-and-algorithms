package main

import (
	"fmt"
)

var stack = [8]int{}
var top int = -1

func main() {
	fmt.Println(top)
	isEmpty()
	isFull()
	push(1)
	pop()
}

func isEmpty() int {
	if top == -1 {
		return 1
	} else {
		return 0
	}
}

func isFull() int {
	if top == 8 {
		return 1
	} else {
		return 0
	}
}

func push(data int) int {
	if isFull() != 8 {
		top = top + 1
		stack[top] = data
		return 1
	} else {
		return 0
	}
}

func pop() int {
	if isEmpty() == 0 {
		var data int = stack[top]
		top = top - 1
		return data
	} else {
		return 0
	}
}
