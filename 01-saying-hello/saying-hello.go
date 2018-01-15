package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	//get name
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("What is your name: ")
	name, _ := reader.ReadString('\n')
	//add name to output string
	output := fmt.Sprintf("Hello, %s, nice  to meet you!", strings.TrimSuffix(name, "\n"))
	//print string
	fmt.Println(output)
}
