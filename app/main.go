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
		command, err := bufio.NewReader(os.Stdin).ReadString('\n')
		if err != nil {
			fmt.Println(err)
		}
		command = strings.TrimSpace(command)
		if command == "exit" {
			break
		}
		fmt.Printf("%s: command not found\n", command)
	}
}
