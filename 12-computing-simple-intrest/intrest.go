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
	fmt.Print("Enter the principal? ")
	principalRead, _ := reader.ReadString('\n')
	principal, _ := strconv.Atoi(strings.TrimSuffix(principalRead, "\n"))

	fmt.Print("Enter the rate of intrest? ")
	intrestRead, _ := reader.ReadString('\n')
	intrestPercent, _:= strconv.Atoi(strings.TrimSuffix(intrestRead, "\n"))
	intrest := float64(intrestPercent) / 100

	fmt.Print("Enter the number of years? ")
	yearsRead, _ := reader.ReadString('\n')
	years, _ := strconv.Atoi(strings.TrimSuffix(yearsRead, "\n"))

	amount := float64(principal) * (1 + (intrest * float64(years)))
	fmt.Printf("After %d years at %d%%, the investment will be worth %.2f€", years, intrestPercent,	 amount)

}
