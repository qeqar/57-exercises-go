package main

import (
	"bufio"
	"fmt"
	"math"
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

	fmt.Print("Enter the number of times the intrest in compoundet per year? ")
	compRead, _ := reader.ReadString('\n')
	comp, _ := strconv.Atoi(strings.TrimSuffix(compRead, "\n"))

	amount := float64(principal) * math.Pow((1 + (intrest / float64(comp))), (float64(comp) * float64(years)))

	fmt.Printf("After %d years at %d%% with %d compounts per yeas, the investment will be worth %.2f€", years, intrestPercent, comp,	amount)
}
