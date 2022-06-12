package main

import "fmt"

//Symbol       Value
//I             1
//V             5
//X             10
//L             50
//C             100
//D             500
//M             1000

// 1 <= num <= 3999

func intToRoman(num int) string {
	strs := map[int]string{
		1:    "I",
		2:    "II",
		3:    "III",
		4:    "IV",
		5:    "V",
		6:    "VI",
		7:    "VII",
		8:    "VIII",
		9:    "IX",
		10:   "X",
		20:   "XX",
		30:   "XXX",
		40:   "XL",
		50:   "L",
		60:   "LX",
		70:   "LXX",
		80:   "LXXX",
		90:   "XC",
		100:  "C",
		200:  "CC",
		300:  "CCC",
		400:  "CD",
		500:  "D",
		600:  "DC",
		700:  "DCC",
		800:  "DCCC",
		900:  "CM",
		1000: "M",
		2000: "MM",
		3000: "MMM",
	}
	factor := 1
	str := ""
	for num > 0 {
		digit := (num % 10) * factor
		factor *= 10
		num /= 10
		str = strs[digit] + str
	}
	return str
}

func main() {
	fmt.Println(intToRoman(1235))
	fmt.Println(intToRoman(3))
	fmt.Println(intToRoman(4))
	fmt.Println(intToRoman(49))
}
