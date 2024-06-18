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
		if "echo" in command {
			fmt.Printf("%s", strings.TrimRight(command, " "))
		}
		fmt.Printf("%s: command not found\n", strings.TrimRight(command, "\n"))
	}
}
