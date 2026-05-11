package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func main() {
	buffer := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("$ ")
		input, err := buffer.ReadString('\n')
		if err != nil {
			fmt.Println(err)
		}
		input = strings.TrimSpace(input)
		inputWords := strings.Fields(input)
		if len(inputWords) < 1 {
			continue
		}
		command := inputWords[0]
		arguments := inputWords[1:]
		var output string

		if command == "exit" {
			break
		} else if command == "echo" {
			output = strings.Join(arguments, " ")
		} else if command == "type" {
			if len(arguments) == 0 {
				continue
			}
			checkedCommand := arguments[0]
			if checkedCommand == "echo" || checkedCommand == "exit" || checkedCommand == "type" {
				output = fmt.Sprintf("%s is a shell builtin", arguments[0])
			} else if path, err := exec.LookPath(arguments[0]); err == nil {
				output = fmt.Sprintf("%s is %s", checkedCommand, path)
			} else {
				output = fmt.Sprintf("%s: not found", arguments[0])
			}
		} else if path, err := exec.LookPath(command); err == nil {
			cmd := exec.Command(path, arguments...)
			out, err := cmd.Output()
			if err != nil {
				output = err.Error()
			} else {
				output = string(out)
			}
		} else {
			output = fmt.Sprintf("%s: command not found", input)
		}
		fmt.Println(strings.TrimSpace(output))
	}
}
