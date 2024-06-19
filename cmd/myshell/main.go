package main

import (
	"bufio"
	"strings"
	"fmt"
	"os"
)

func main() {


	// Wait for user input
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Fprint(os.Stdout, "$ ")
		command, _ := reader.ReadString('\n')
		command = strings.TrimSpace(command)
		if command == "exit 0" {
			os.Exit(0)
		}
		if strings.Contains(command, "echo") {
			fmt.Printf("%s\n", strings.SplitAfterN(command, " ", 1))
		} else {
		fmt.Printf("%s: command not found\n", strings.TrimRight(command, "\n"))
		}
	}
}
