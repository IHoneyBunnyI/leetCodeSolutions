package main

import (
	"fmt"
	"sort"
)

type Line struct {
	start int
	end   int
	size  int
}

func dotContainsInLine(dot int, line Line) bool {
	if dot >= line.start && dot <= line.end {
		return true
	}
	return false
}

func main() {
	var n int
	fmt.Scan(&n)
	lines := make([]Line, n)
	for i := 0; i < n; i++ {
		start := 0
		end := 0
		fmt.Scan(&start, &end)
		lines[i].start = start
		lines[i].end = end
		lines[i].size = end - start + 1
	}

	res := make([]int, 0)
	sort.Slice(lines, func(i, j int) bool { return lines[i].end < lines[j].end })
	//fmt.Println(lines)
	for len(lines) != 0 {
		dot := lines[0].end
		i := 1
		for i < len(lines) && dotContainsInLine(dot, lines[i]) {
			i++
		}
		res = append(res, dot)
		lines = lines[i:]
	}
	fmt.Println(len(res))
	for i := 0; i < len(res); i++ {
		fmt.Print(res[i], " ")
	}
	//fmt.Println(minLine.size, minLine.start, minLine.end)
	//res := make([]int, 0)
	//for i := minLine.start; i <= minLine.end; i++ {
	//if containsInLines(i, lines) {
	//res = append(res, i)
	//}
	//}
}
