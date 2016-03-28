// ConvertMinutesToHours project main.go

/*
Converts raw minutes into hours and minutes

Usage: ./ConvertMinutesToHours <minutes>

Input: Minutes from command line
Output: H hours and M minutes

On Error: Return "not enough args."
*/

package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {

	if len(os.Args) > 1 {
		minArg := os.Args[1]
		convArg, _ := strconv.ParseFloat(minArg, 2)
		hours := convArg / 60
		m := strings.Split(fmt.Sprintf("%f", hours), ".")
		convH, _ := strconv.ParseFloat(m[0], 2)
		decimal := hours - convH
		math := decimal * 60
		fmt.Printf("%v hours and %1.f Minutes\n", convH, math)
	} else {
		fmt.Println("Not enough args.")
	}

}
