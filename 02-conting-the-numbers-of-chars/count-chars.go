package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Waht is the input string: ")
	input, _ := reader.ReadString('\n')
	input = strings.TrimSuffix(input, "\n")

	fmt.Printf("%s, has %d charackters", input, len(input))
}
