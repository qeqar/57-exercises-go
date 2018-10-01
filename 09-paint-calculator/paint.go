package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"math"
)

var squareFeetPerGallon float64 = 350

func main(){
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("What is the length of the room in feet: ")
	lengthRead, _ := reader.ReadString('\n')
	fmt.Print("What is the width of the room in feet: ")
	widthRead, _ := reader.ReadString('\n')

	length, _ := strconv.ParseFloat(strings.TrimSuffix(lengthRead, "\n"), 64)
	width, _ := strconv.ParseFloat(strings.TrimSuffix(widthRead, "\n"), 64)

	areaFeet := length * width

	gallons := areaFeet/squareFeetPerGallon

	fmt.Printf("You need to buy %d gallons of point", int(math.Ceil(gallons)))
}
