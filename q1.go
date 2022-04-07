package main

import (
	"flag"
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

func main() {
	useList := flag.Bool("list", false, "use a quack implemented by list, default stack based")
	useFdq := flag.Bool("fdq", false, "use a quack implemented by function dequeue, default stack based")
	flag.Parse()

	var q quack.Quack

	if *useList {
		q = quack.NewListQuack()
	} else if *useFdq {
		q = quack.NewFdqQuack()
	} else {
		q = quack.NewStackQuack()
	}

	commands := flag.Args()

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
