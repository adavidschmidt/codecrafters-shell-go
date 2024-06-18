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
		command, err := reader.ReadString('\n')
		command = strings.TrimSpace(command)
		if err != nil {
			fmt.Printf("%s: command not found\n", strings.TrimRight(command, "\n"))
		}
		
	}
}
