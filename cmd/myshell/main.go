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

func runEcho(a []string) {
	a = strings.Join(a, " ")
	fmt.Printf("%s\n", a)
}

func runExit() {
	os.Exit(0)
}

func runType(a string, b []string) {
	if stringInSlice(a, b) {
		fmt.Printf("%s is a shell builtin\n", a)
	} else {
		paths := strings.Split(os.Getenv("PATH"), ":")
		for _, path := range paths {
			dir := path + "/" + a
			if _, err := os.Stat(dir); err == nil {
				fmt.Printf("%s is %v\n", a, dir)
				return
			}
		}
		fmt.Printf("%s: not found\n", a)
	} 	
}

func main() {


	// Wait for user input
	reader := bufio.NewReader(os.Stdin)
	listCommands := []string{"echo", "type", "exit"}
	for {
		fmt.Fprint(os.Stdout, "$ ")
		command, _ := reader.ReadString('\n')
		tokens := strings.Fields(strings.TrimSpace(command))
		command = strings.Join(tokens, " ")
		switch {
		case tokens[0] == "exit" && tokens[1] == "0":
			runExit()
			
		case tokens[0] == "echo":
			runEcho(tokens[1:])
		
		case tokens[0] == "type":
			runType(tokens[1], listCommands)
		
		default:
			fmt.Printf("%s: command not found\n", command)
		}
	}
}
