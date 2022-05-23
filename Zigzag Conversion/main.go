package main

import "fmt"

func convert(s string, numRows int) string {
	var res string
	if numRows == 1 {
		return s
	}
	canvas := make([][]byte, numRows)
	for i := range canvas {
		canvas[i] = make([]byte, len(s))
	}
	x := 0
	y := 0
	up := false
	down := true
	for _, ch := range s {
		if down {
			canvas[y][x] = byte(ch)
			y++
			if y == numRows {
				y -= 2
				x++
				down = false
				up = true
				continue
			}
		}
		//fmt.Println(y)
		//for i := range canvas {
		//fmt.Println(canvas[i])
		//}
		if up {
			canvas[y][x] = byte(ch)
			x++
			y--
			if y == -1 {
				y = 1
				x--
				up = false
				down = true
			}
		}
	}

	for i := range canvas {
		for _, ch := range canvas[i] {
			//fmt.Print(string(ch), "\t")
			if ch != byte(0) {
				res += string(ch)
			}
		}
		//fmt.Print("\n")
	}
	//fmt.Println(res)
	return res
}

func main() {
	var s string

	s = convert("PAYPALISHIRING", 4)
	if fmt.Sprint("PINALSIGYAHRPI") == s {
		fmt.Println("OK")
	} else {
		fmt.Println("KO")
	}
	/*Input: s = "PAYPALISHIRING", numRows = 4
	Output: "PINALSIGYAHRPI"
	Explanation:
	P     I    N
	A   L S  I G
	Y A   H R
	P     I
	*/

	s = convert("PAYPALISHIRING", 3)
	if fmt.Sprint("PAHNAPLSIIGYIR") == s {
		fmt.Println("OK")
	} else {
		fmt.Println("KO")
	}
	/*
		P   A   H   N
		A P L S I I G
		Y   I   R
	*/

	s = convert("A", 1)
	if fmt.Sprint("A") == s {
		fmt.Println("OK")
	} else {
		fmt.Println("KO")
	}

	s = convert("AB", 1)
	if fmt.Sprint("AB") == s {
		fmt.Println("OK")
	} else {
		fmt.Println("KO")
	}
	/*
		numRows = 6
		P         R
		A       I I
		Y     H   N
		P   S     G
		A I
		L
	*/
}
