package main

import (
	"fmt"
	queue2 "go-structure/queue"
)

func main() {
	queue := queue2.Queue{}
	queue.Push(1)
	queue.Push(2)
	queue.Push(3)
	res, err := queue.Pop()
	if err != nil {
		fmt.Print(err)
	}
	fmt.Println(queue.Cap())
	fmt.Println(res)
}
