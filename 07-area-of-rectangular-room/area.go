package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

var conversion float64 = 0.09290304

func main() {
	// input
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("What is the length of the room in feet: ")
	lengthRead, _ := reader.ReadString('\n')
	fmt.Print("What is the width of the room in feet: ")
	widthRead, _ := reader.ReadString('\n')

	length, _ := strconv.ParseFloat(strings.TrimSuffix(lengthRead, "\n"), 64)
	width, _ := strconv.ParseFloat(strings.TrimSuffix(widthRead, "\n"), 64)

	areaFeet := length * width
	areaMetric := math.Sqrt(areaFeet * areaFeet * conversion)

	fmt.Printf("You enterd demensions of %.2f feet by %.2f feet\n", length, width)
	fmt.Printf("The area is\n%.2f square feet\n%.2f square meters", areaFeet, areaMetric)
}

