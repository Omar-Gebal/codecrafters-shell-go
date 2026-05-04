package main

import (
	"bufio"
	"fmt"
	"os"
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
		if command == "exit" {
			break
		} else if command == "echo" {
			var output string
			for i, v := range arguments {
				output += v
				if i != len(arguments)-1 {
					output += " "
				}
			}
			fmt.Println(output)
		} else {
			fmt.Printf("%s: command not found\n", input)
		}
	}
}
