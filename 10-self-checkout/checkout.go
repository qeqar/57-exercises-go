package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var taxRate = 0.55

type Pair struct {
	num, quantity interface{}
}

func main(){
	reader := bufio.NewReader(os.Stdin)
	var items [3]Pair //struct{num; quantity}

	for i := 0; i < 3; i++ {
		fmt.Printf("Enter the price of item %d: ", i + 1)
		numRead, _ := reader.ReadString('\n')
		fmt.Printf("Enter the quantity of item %d: ", i + 1 )
		quantityRead, _ := reader.ReadString('\n')

		num, _ := strconv.ParseFloat(strings.TrimSuffix(numRead, "\n"), 64)
		quantity, _ := strconv.ParseFloat(strings.TrimSuffix(quantityRead, "\n"), 64)
		m := Pair{num, quantity}
		items[i] = m
	}

	subTotal := 0.00
	for i := range items {
		subTotal += items[i].num.(float64) * items[i].quantity.(float64)

	}
	tax := subTotal * taxRate
	total := subTotal + tax

	fmt.Printf("Subtotal: %.2f\n", subTotal)
	fmt.Printf("Tax: %.2f\n", tax)
	fmt.Printf("Total: %.2f\n", total)
}
