package main

import (
	"fmt"
	"strings"
)

func lengthOfLongestSubstring(s string) int {
	if s == "" {
		return 0
	}
	maxLen := 1
	for i := range s {
		str := ""
		str += string(s[i])
		for j := i + 1; j < len(s); j++ {
			if strings.Contains(str, string(s[j])) {
				//fmt.Println(str)
				if len(str) > maxLen {
					maxLen = len(str)
					str = ""
				}
				break
			} else {
				str += string(s[j])
			}
			if len(str) > maxLen {
				maxLen = len(str)
			}
		}
	}
	return maxLen
}

func main() {
	a := 0
	a = lengthOfLongestSubstring("abcabcbb")
	fmt.Println(a)
	a = lengthOfLongestSubstring("bbbbb")
	fmt.Println(a)
	a = lengthOfLongestSubstring("pwwkew")
	fmt.Println(a)
	a = lengthOfLongestSubstring("a")
	fmt.Println(a)
	a = lengthOfLongestSubstring("abbbbced")
	fmt.Println(a)
}
