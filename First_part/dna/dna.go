package main

import (
	"fmt"
)

func DNAStrand(dna string) string {
	str := ""

	for _, value := range dna {
		if value != 'A' && value != 'T' && value != 'C' && value != 'G' {
			return "Not a DNA Strand"
		}

	}
	for _, value := range dna {
		if value == 'A' {
			str += string('T')

		} else if value == 'T' {
			str += string('A')
		} else if value == 'C' {
			str += string('G')

		} else if value == 'G' {
			str += string('C')

		}

	}

	return str
}

func main() {
	d := "ATTGC"
	x := DNAStrand(d)
	fmt.Println(x)

}
