package main

import (
	"fmt"
	"regexp"
)

func isMatch(s string, p string) bool {
	patttern := "^" + p + "$"
	matched, _ := regexp.MatchString(patttern, s)
	return matched
}

func main() {
	if r := isMatch("aa", "s"); !r {
		fmt.Println("OK")
	} else {
		fmt.Println("KO")
	}

	if r := isMatch("aa", "a"); !r {
		fmt.Println("OK")
	} else {
		fmt.Println("KO")
	}

	if r := isMatch("aa", "a*"); r {
		fmt.Println("OK")
	} else {
		fmt.Println("KO")
	}

	//"mississippi"
	//"mis*is*p*."
	// false
	if r := isMatch("mississippi", "mis*is*p*."); !r {
		fmt.Println("OK")
	} else {
		fmt.Println("KO")
	}

	//"mississippi"
	//"mis*is*ip*."
	//true
	if r := isMatch("mississippi", "mis*is*ip*."); r {
		fmt.Println("OK")
	} else {
		fmt.Println("KO")
	}
}
