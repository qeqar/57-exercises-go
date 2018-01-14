/*
pre exercises Tip Calculator

 */
package main

import (
	"fmt"
	"bufio"
	"os"
	"strconv"
	"strings"
)

func Round(x, unit float64) float64 {
	return float64(int64(x/unit+0.5)) * unit
}

func main() {
	// read data
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("What is the bill amount: ")
	billAmountRead, _ := reader.ReadString('\n')
	fmt.Print("What is the tip rate: ")
	tipRateRead, _ := reader.ReadString('\n')

	//convert to int
	billAmount, _ := strconv.ParseFloat(strings.TrimSuffix(billAmountRead, "\n"), 32)
	tipRate, _ := strconv.ParseFloat(strings.TrimSuffix(tipRateRead, "\n"), 64)

	//calculate
	var tip float64 = Round(billAmount * (tipRate / 100), 0.01)
	var total float64 = billAmount + tip

    // print
	fmt.Printf("Tip: %.2f \n", tip)
	fmt.Printf("Total: %.2f \n", total)
}