package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Task Tracker CLI")
		fmt.Println("Usage: task-cli <command> [arguments]")
		return
	}

	command := os.Args[1]
	fmt.Println("Command received:", command)
}
