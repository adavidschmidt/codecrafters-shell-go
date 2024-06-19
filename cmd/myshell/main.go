package main

import (
	"bufio"
	"strings"
	"fmt"
	"os"
)

func stringInSlice(a string, list []string) bool {
    for _, b := range list {
        if b == a {
            return true
        }
    }
    return false
}

func main() {


	// Wait for user input
	reader := bufio.NewReader(os.Stdin)
	listCommands := ["echo", "type", "exit"]
	for {
		fmt.Fprint(os.Stdout, "$ ")
		command, _ := reader.ReadString('\n')
		tokens := strings.Fields(strings.TrimSpace(command))
		command = strings.Join(tokens, " ")
		switch {
		case tokens[0] == "exit" && tokens[1] == "0":
				os.Exit(0)
			
		case tokens[0] == "echo":
			fmt.Printf("%s", strings.Join(tokens[1:], " "))
		
		case tokens[0] == "type" && stringInSlice(tokens[1], listCommands):
			fmt.Printf("%s is a shell builtin\n", tokens[1])
		
		case tokens[0] == "type" && !stringInSlice(tokens[1]. listCommands):
			fmt.Printf("%s: command not found\n", tokens[1])
		
		default:
			fmt.Printf("%s: command not found\n", command)
		}
	}
}
