package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var slicesPerPizza int = 8

func main(){
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("How many people: ")
	peopleRead, _ := reader.ReadString('\n')
	fmt.Print("How many pizza: ")
	pizzaRead, _ := reader.ReadString('\n')
	people, _ := strconv.Atoi(strings.TrimSuffix(peopleRead, "\n"))
	pizza, _ := strconv.Atoi(strings.TrimSuffix(pizzaRead, "\n"))

	numSlices := pizza * slicesPerPizza

	slicesPerPerson, remainder := numSlices/people, numSlices%people

	fmt.Printf("We have %d slices per person and %d extra sclices", slicesPerPerson, remainder)

	}
