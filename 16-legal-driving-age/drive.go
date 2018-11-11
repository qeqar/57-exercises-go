package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var legalDrivingAge int = 16

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("What is your age? ")
	ageRead, _ := reader.ReadString('\n')
	age, _ := strconv.Atoi(strings.TrimSuffix(ageRead, "\n"))

	if age >= legalDrivingAge {
		fmt.Println("You are old enough to legally drive")
	} else {
		fmt.Println("You are not old enough to legally drive")
	}
}
