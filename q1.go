package main

import (
	"fmt"
	"os"
	"quack/quack"
)

const (
	nothing = iota
	pushNext
	pullNext
	popNext
)

var opNames = []string{"nothing", "push", "pull", "pop"}

func main() {
	q := quack.NewStackQuack()
	n := 1

	commands := os.Args[n:]

	for i := 0; i < len(commands); {
		str := commands[i]
		switch str {
		case "push":
			q.Push(commands[i+1])
			i += 2
		case "pop":
			popped := q.Pop()
			fmt.Printf("pop %v\n", popped)
			i++
		case "pull":
			pulled := q.Pull()
			fmt.Printf("pull %v\n", pulled)
			i++
		case "quack":
			q.Print(os.Stdout)
			i++
		default:
			fmt.Printf("noop\n")
			i++
		}
	}
}
