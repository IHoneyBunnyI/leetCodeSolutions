package main

import (
	"fmt"
	"sort"
)

type Item struct {
	value  int
	weight int
	k      float64
}

func main() {
	var n, capacity int
	fmt.Scan(&n, &capacity)
	items := make([]Item, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&items[i].value, &items[i].weight)
		items[i].k = float64(items[i].value) / float64(items[i].weight)
	}
	sort.Slice(items, func(i, j int) bool { return items[i].k > items[j].k })
	//fmt.Println(n, capacity, items)
	var res float64
	for _, item := range items {
		if capacity >= item.weight {
			capacity -= item.weight
			res += float64(item.value)
		} else {
			res += float64(item.k * float64(capacity))
			break
		}
	}
	fmt.Printf("%.3f\n", res)
}
