package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func main() {
	for {
		fmt.Print("$ ")
		input, err := bufio.NewReader(os.Stdin).ReadString('\n')
		if err != nil {
			fmt.Println(err)
		}
		input = strings.TrimSpace(input)
		inputWords := strings.Fields(input)
		command := inputWords[0]
		arguments := inputWords[1:]
		var output strings.Builder

		if command == "exit" {
			break
		} else if command == "echo" {
			for i, v := range arguments {
				output.WriteString(v)
				if i != len(arguments)-1 {
					output.WriteString(" ")
				}
			}
		} else if command == "type" {
			checkedCommand := arguments[0]
			if checkedCommand == "echo" || checkedCommand == "exit" || checkedCommand == "type" {
				output.WriteString(fmt.Sprintf("%s is a shell builtin", arguments[0]))
			} else if path, err := exec.LookPath(arguments[0]); err == nil {
				output.WriteString(fmt.Sprintf("%s is %s", checkedCommand, path))

			} else {
				output.WriteString(fmt.Sprintf("%s: not found", arguments[0]))
			}
		} else {
			output.WriteString(fmt.Sprintf("%s: command not found", input))
		}
		fmt.Println(output.String())
	}
}
