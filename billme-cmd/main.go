package main

import (
	"billme-cmd/cmd"
	"bufio"
	"fmt"
	"os"
)

func main() {

	reader := bufio.NewReader(os.Stdin)

	inputCommand, err := reader.ReadString('\n')
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
	}

	fmt.Println(inputCommand)

	cmd := cmd.Command{Name: inputCommand}
	cmd.Process()
}
