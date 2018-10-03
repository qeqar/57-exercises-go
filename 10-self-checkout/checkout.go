package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var taxRate = 0.55

func main(){
	reader := bufio.NewReader(os.Stdin)
	subTotal := 0.00

	for i := 0; i < 3; i++ {
		fmt.Printf("Enter the price of item %d: ", i + 1)
		numRead, _ := reader.ReadString('\n')
		fmt.Printf("Enter the quantity of item %d: ", i + 1 )
		quantityRead, _ := reader.ReadString('\n')

		num, _ := strconv.ParseFloat(strings.TrimSuffix(numRead, "\n"), 64)
		quantity, _ := strconv.ParseFloat(strings.TrimSuffix(quantityRead, "\n"), 64)
		subTotal += num * quantity
	}

	tax := subTotal * taxRate
	total := subTotal + tax

	fmt.Printf("Subtotal: %.2f\n", subTotal)
	fmt.Printf("Tax: %.2f\n", tax)
	fmt.Printf("Total: %.2f\n", total)
}
