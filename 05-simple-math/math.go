package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	// read data
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("First Number: ")
	firstRead, _ := reader.ReadString('\n')
	fmt.Print("Second Number: ")
	secondRead, _ := reader.ReadString('\n')

	//convert to int
	first, _ := strconv.Atoi(strings.TrimSuffix(firstRead, "\n"))
	second, _ := strconv.Atoi(strings.TrimSuffix(secondRead, "\n"))

	var sum int = first + second
	var dif int = first - second
	var prod int = first + second
	var quo int = first / second

	fmt.Println(strconv.Itoa(first) + " + " + strconv.Itoa(second) + " = " + strconv.Itoa(sum) + "\n" + strconv.Itoa(first) + " - " + strconv.Itoa(second) + " = " + strconv.Itoa(dif) + "\n" + strconv.Itoa(first) + " * " + strconv.Itoa(second) + " = " + strconv.Itoa(prod) + "\n" + strconv.Itoa(first) + " / " + strconv.Itoa(second) + " = " + strconv.Itoa(quo))
}
