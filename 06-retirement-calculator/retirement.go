package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

func main(){
	currentTime := time.Now()
	var year int = currentTime.Year()

	reader := bufio.NewReader(os.Stdin)
	fmt.Print("How old are you: ")
	oldRead, _ := reader.ReadString('\n')
	fmt.Print("When do you want to retire: ")
	retRead, _ := reader.ReadString('\n')

	//convert to int
	old, _ := strconv.Atoi(strings.TrimSuffix(oldRead, "\n"))
	ret, _ := strconv.Atoi(strings.TrimSuffix(retRead, "\n"))

	var yearsToCome int = ret - old
	var retyear int = year + yearsToCome

	fmt.Printf("you have %d years to work and the year of our retirement is %d \n", yearsToCome, retyear)
}
