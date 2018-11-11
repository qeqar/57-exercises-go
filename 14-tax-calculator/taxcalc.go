package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main(){
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("What is the orderamount? ")
	amountRead, _ := reader.ReadString('\n')
	amount, _ := strconv.ParseFloat(strings.TrimSuffix(amountRead, "\n"), 64)

	fmt.Print("What is the state? ")
	stateRead, _ := reader.ReadString('\n')
	state := strings.TrimSuffix(stateRead, "\n")

	tax := float64(0)

	if strings.ToLower(state) == "wi" || strings.ToLower(state) == "wisconsin"{
		taxRate := 0.055
		tax = amount * taxRate
		fmt.Printf("The Subtotal is %.2f\n", amount)
		fmt.Printf("The Tax is %.2f\n", tax)
	}

	fmt.Printf("The total ist %.2f\n", amount + tax)
}
