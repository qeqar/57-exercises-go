package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Waht is the quote: ")
	quote, _ := reader.ReadString('\n')
	fmt.Print("Who said this: ")
	author, _ := reader.ReadString('\n')

	fmt.Println(strings.TrimSuffix(author, "\n") + " says, \"" + strings.TrimSuffix(quote, "\n") + "\"")
}
