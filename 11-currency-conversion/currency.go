package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("How many euros are you exchanging? ")
	euroRead, _ := reader.ReadString('\n')
	euro, _ := strconv.ParseFloat(strings.TrimSuffix(euroRead, "\n"), 64)

	fmt.Print("What is the exchange rate? ")
	rateRead, _ := reader.ReadString('\n')
	rate, _ := strconv.ParseFloat(strings.TrimSuffix(rateRead, "\n"), 64)

	dollar := euro * rate
	fmt.Printf("%.2f euros at exchange rate of %.2f is %.2f dollars.", euro, rate, math.Ceil(dollar * 100) / 100)
}
