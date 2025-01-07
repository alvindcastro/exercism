package partyrobot

import (
	"fmt"
	"math"
	"strconv"
)

// Welcome greets a person by name.
func Welcome(name string) string {
	return "Welcome to my party, " + name + "!"
}

// HappyBirthday wishes happy birthday to the birthday person and exclaims their age.
func HappyBirthday(name string, age int) string {
	return "Happy birthday " + name + "! You are now " + strconv.Itoa(age) + " years old!"
}

// AssignTable assigns a table to each guest.
func AssignTable(name string, table int, neighbor, direction string, distance float64) string {
	assignedTable := "You have been assigned to table " + fmt.Sprintf("%03d", table)
	//formatFloat := strconv.FormatFloat(distance, 'f', -1, 64)
	formatFloat := toFixed(distance, 1)

	return "Welcome to my party, " + name + "!" + "\n" +
		assignedTable +
		". Your table is " + direction + ", exactly " + fmt.Sprintf("%.1f", formatFloat) + " meters from here." + "\n" +
		"You will be sitting next to " + neighbor + "."

}

func toFixed(num float64, precision int) float64 {
	output := math.Pow(10, float64(precision))
	return float64(round(num*output)) / output
}

func round(num float64) int {
	return int(num + math.Copysign(0.5, num))
}
