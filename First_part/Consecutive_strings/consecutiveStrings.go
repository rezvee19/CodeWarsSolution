// You are given an array(list) strarr of strings and an integer k.
// Your task is to return the first longest string consisting of k consecutive strings taken in the array.
package main

import (
	"fmt"	"strings"
)


func LongestConsec(strarr []string, k int) string {
	l := 0
	largest := ""
	
	for i := 0; i < len(strarr)-k+1; i++ {
		g := strings.Join(strarr[i:i+k], "")
		if len(g) > l {
			l = len(g)
			largest = g
			
		}

	}
	return largest

}

func main() {

	x := []string{"abc", "def", "ghmmmi", "jkl", "mno"}
	
	s := LongestConsec(x, 3)
	fmt.Println(s)
	
}
