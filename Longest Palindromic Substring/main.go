package main

import (
	"fmt"
)

func Reverse(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

func setLongest(s string, longest *string) {
	if s == Reverse(s) {
		if len(*longest) < len(s) {
			*longest = s
		}
		//fmt.Println(*longest)
	}
}

func longestPalindrome(s string) string {
	var longest string
	var str string

	if len(s) == 1 || (len(s) == 2 && s[0] == s[1]) {
		return s
	}
	if len(s) == 2 && s[0] != s[1] {
		return (string(s[0]))
	}
	if s == Reverse(s) {
		return s
	}
	for start := 0; start < len(s)-1; start++ {
		str = string(s[start])
		setLongest(str, &longest)
		if s[start] == s[start+1] {
			for j := 1; start+j < len(s); j++ {
				if s[start] == s[start+j] {
					str = s[start : start+j+1]
					//str += string(s[start + j])
					setLongest(str, &longest)
					//fmt.Print("|", start," ", longest, " ","|")
				} else {
					break
				}
			}
		}
		if len(str)%2 == 0 {
			//fmt.Print("start:", start, "str:", str, " ")
			for j := 1; start+j+1 < len(s) && start-j >= 0; j++ {
				if s[start+1+j] == s[start-j] {
					str = s[start-j : start+j+2]
					//fmt.Print("start:", start, "str:", str, " ")
					setLongest(str, &longest)
					//fmt.Print(start," ", longest, " ")
				} else {
					break
				}
			}
		}
		//if s[start+1] == s[start] {
		//// если четное кол-во
		//str = s[start : start+2]
		//setLongest(str, &longest)
		//for j := 1; start+1+j < len(s) && start-j >= 0; j++ {
		//if s[start+1+j] == s[start-j] {
		//str = s[start-j : start+j+2]
		//setLongest(str, &longest)
		//} else {
		//break
		//}
		//}
		//} else {
		// если не четное кол-во
		for j := 1; start+j < len(s) && start-j >= 0; j++ {
			if s[start+j] == s[start-j] {
				str = s[start-j : start+j+1]
				setLongest(str, &longest)
				//fmt.Print(start," ", longest, " ")
			} else {
				break
			}
		}
	}
	//}
	//fmt.Println("BBB", longest)
	return longest
}
func main() {
	var s string
	s = longestPalindrome("bbcaabaac!aabaaaa")
	fmt.Println(s)
	s = longestPalindrome("aacabdkacaa")
	fmt.Println(s)
	s = longestPalindrome("pwkww")
	fmt.Println(s)
	s = longestPalindrome("a")
	fmt.Println(s)
	s = longestPalindrome("abbbbced")
	fmt.Println(s)
	s = longestPalindrome("caac")
	fmt.Println(s)
	s = longestPalindrome("caxac")
	fmt.Println(s)
	s = longestPalindrome("acb")
	fmt.Println(s)
	s = longestPalindrome("ccc")
	fmt.Println(s)
	s = longestPalindrome("cccc")
	fmt.Println(s)
	s = longestPalindrome("cc")
	fmt.Println(s)
	s = longestPalindrome("accc!ccca")
	fmt.Println(s)
	s = longestPalindrome("bppbgdsooosgq")
	fmt.Println(s)
	s = longestPalindrome("tattarrattat")
	fmt.Println(s)
	s = longestPalindrome("iptmykvjanwiihepqhzupneckpzomgvzmyoybzfynybpfybngttozprjbupciuinpzryritfmyxyppxigitnemanreexcpwscvcwddnfjswgprabdggbgcillisyoskdodzlpbltefiz")
	fmt.Println(s)

}
