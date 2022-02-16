package main

import (
	"fmt"
	"os"
	"quack/quack"
)

func main() {
	q := quack.NewStackQuack()

	for idx := range os.Args {
		q.Push(os.Args[idx])
	}
	for {
		str := q.Pull()
		if str == nil {
			break
		}
		fmt.Printf("%v\n", str)
	}
}
