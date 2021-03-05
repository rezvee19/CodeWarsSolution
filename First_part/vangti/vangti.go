// The new "Avengers" movie has just been released! There are a lot of people at the cinema box office standing in a huge line. Each of them has a single 100, 50 or 25 dollar bill. An "Avengers" ticket costs 25 dollars.

// Vasya is currently working as a clerk. He wants to sell a ticket to every single person in this line.

// Can Vasya sell a ticket to every person and give change if he initially has no money and sells the tickets strictly in the order people queue?

// Return YES, if Vasya can sell a ticket to every person and give change with the bills he has at hand at that moment. Otherwise return NO.

// Examples:
// Tickets([]int{25, 25, 50}) // => YES
// Tickets([]int{25, 100}) // => NO. Vasya will not have enough money to give change to 100 dollars
// Tickets([]int{25, 25, 50, 50, 100}) // => NO. Vasya will not have the right bills to give 75 dollars of change (you can't make two bills of 25 from one of 50)
package main

import (
	"fmt"
)

// Tickets can take only 25,50,100 dollers bill
func Tickets(peopleInLine []int) string {
	var temp1 int = 0
	var temp2 int = 0

	for _, b := range peopleInLine {

		if b == 25 {
			temp1 = temp1 + 1
		} else if b == 50 {
			if temp1 >= 1 {

				temp2 = temp2 + 1
				temp1 = temp1 - 1
			} else {

				return "NO. Vasya will not have enough money to give change to 50 dollars"

			}

		} else if b == 100 {
			if temp2 >= 1 && temp1 >= 1 {

				temp2 = temp2 - 1
				temp1 = temp1 - 1
			} else if temp1 >= 3 {
				temp1 = temp1 - 3

			} else if temp2 >= 2 && temp1 == 0 {

				return "NO. Vasya will not have the right bills to give 75 dollars of change (you can't make two bills of 25 from one of 50)"

			} else {
				return "NO. Vasya will not have enough money to give change to 100 dollars"

			}
		}
	}
	return "YES"
}

func main() {
	x := Tickets([]int{25, 25, 50, 50, 100})

	fmt.Println(x)

}
