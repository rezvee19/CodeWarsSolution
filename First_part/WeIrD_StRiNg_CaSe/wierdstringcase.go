package main

import (
	"fmt"
	"strings"
)

func wsc(str string) string {
	s := ""
	t := ""
	checkSpace := 0

	for _, value := range str {
		if value != ' ' {
			if checkSpace%2 == 0 {
				s = strings.ToUpper(string(value))

			} else {
				s = strings.ToLower(string(value))
			}

			t += s
			checkSpace++
		} else {
			t += string(value)
			checkSpace = 0
		}

	}

	return t
}
func main() {

	str := "Weird string case"
	x := wsc(str)
	fmt.Println(x)

}
