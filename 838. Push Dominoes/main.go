package main

import (
	"fmt"
)

var countTest int = 0

// R.........L..........L..........R

func pushDominoes(domino string) string {
	dominoes := []byte(domino)
	right := -1
	for counter := 0; counter < len(dominoes); counter++ {
		if dominoes[counter] == 'L' {
			if right == -1 {
				for i := counter - 1; i >= 0 && dominoes[i] == '.'; i-- {
					dominoes[i] = 'L'
				}
			} else {
				i := right + 1
				j := counter - 1
				for i < j {
					dominoes[i] = 'R'
					dominoes[j] = 'L'
					i++
					j--
				}
			}
			right = -1
		} else if dominoes[counter] == 'R' {
			if right != -1 {
				for i := right + 1; i < counter; i++ {
					dominoes[i] = 'R'
				}
			}
			right = counter
		}
	}
	if right != -1 {
		for i := right + 1; i < len(dominoes); i++ {
			dominoes[i] = 'R'
		}
	}
	return string(dominoes)
}

func test(dominoes string) {
	fmt.Print(countTest, ") ", dominoes, "= ")
	res := pushDominoes(dominoes)
	fmt.Print(res, "\n")
	countTest++
}

func main() {
	test("....L....R......")
	test("L.R.LR")
	test("R...LR")
}
