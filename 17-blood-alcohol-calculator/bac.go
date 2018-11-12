package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var distributionRate = map[string]float64{
	"w": 0.66, //woman
	"m": 0.73, //men
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("What is your wight? ")
	weightRead, _ := reader.ReadString('\n')
	weight, _ := strconv.ParseFloat(strings.TrimSuffix(weightRead, "\n"), 64)

	fmt.Print("What is your gender? (m/w) ")
	genderRead, _ := reader.ReadString('\n')
	gender := strings.TrimSuffix(genderRead, "\n")

	fmt.Print("How many oz alcohol do you had? ")
	ozRead, _ := reader.ReadString('\n')
	oz, _ := strconv.ParseFloat(strings.TrimSuffix(ozRead, "\n"), 64)

	fmt.Print("With which alcohol volume? ")
	volumeRead, _ := reader.ReadString('\n')
	volume, _ := strconv.Atoi(strings.TrimSuffix(volumeRead, "\n"))
	volumePercent := float64(volume) / 100

	fmt.Print("How many hours ago was your last drink? ")
	hoursRead, _ := reader.ReadString('\n')
	hours, _ := strconv.ParseFloat(strings.TrimSuffix(hoursRead, "\n"), 64)

	bac := (oz*volumePercent + 5.14/weight*distributionRate[gender]) - 0.15*hours

	if bac < 0 {
		bac = 0
	}

	if bac >= 0.08 {
		fmt.Printf("Your BAC is %.2f, you are not allowed to drive", bac)
	} else {
		fmt.Printf("Your BAC is %.2f, you are allowed to drive", bac)

	}

}
