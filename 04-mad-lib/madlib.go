package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter a noun: ")
	noun, _ := reader.ReadString('\n')
	fmt.Print("Enter a verb: ")
	verb, _ := reader.ReadString('\n')
	fmt.Print("Enter a adjective: ")
	adjec, _ := reader.ReadString('\n')
	fmt.Print("Enter a adverb: ")
	adverb, _ := reader.ReadString('\n')

	fmt.Println("Every time you " + strings.TrimSuffix(verb, "\n") + " into the " + strings.TrimSuffix(noun, "\n") + " someone " + strings.TrimSuffix(adjec, "\n") + " balwill miss a " + strings.TrimSuffix(adverb, "\n") + " wish")
}
