// ConvertMinutesToHours project main.go
package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {

	fmt.Print("Enter the number of minutes to convert: ")
	var minutes float64
	fmt.Scanf("%f", &minutes)

	hours := minutes / 60

	m := strings.SplitAfter(fmt.Sprintf("%f", hours), "4")
	n := strings.SplitAfter(fmt.Sprintf("%f", hours), ".")

	convMin, _ := strconv.ParseFloat(m[1], 2)
	convH, _ := strconv.ParseFloat(n[0], 2)

	math := convMin * 60

	fmt.Printf("%v hours and %1.f Minutes\n", convH, math)

}
