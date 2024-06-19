package main

import (
	"bufio"
	"strings"
	"fmt"
	"os"
	"os/exec"
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
	b := strings.Join(a, " ")
	fmt.Printf("%s\n", b)
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
				fmt.Printf("%s is %s\n", a, dir)
				return
			}
		}
		fmt.Printf("%s: not found\n", a)
	} 	
}

func runCode(a string, b string) {
	paths := strings.Split(os.Getenv("PATH"), ":")
		for _, path := range paths {
			dir := path + "/" + a
			if _, err := os.Stat(a, b); err == nil {
				cmd := exec.Command(dir)
				cmd.Stdout = os.Stdout
				cmd.Stdin = os.Stdin
				if err := cmd.Run(); err != nil {
					fmt.Printf("Error: %s", err)
				}
				break
			}
		}
		fmt.Printf("%s: not found\n", a)
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
			runCode(tokens[0], tokens[1])
			fmt.Printf("%s: command not found\n", command)
		}
	}
}
