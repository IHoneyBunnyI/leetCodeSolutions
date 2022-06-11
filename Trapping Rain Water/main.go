package main

//height = [0,1,0,2,1,0,1,3,2,1,2,1] => 6
//height = [4,2,0,3,2,5] => 9

func main() {
	input := []int{2, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}
	Slice(input).Print()
	//PrintSlice(input)
	//slice := [][]byte{{'.', '.', '#'}, {'.', '#', '#'}, {'.', '.', '#'}}
	//for i := range slice {
	//for j := range slice {
	//fmt.Print(string(slice[j][i]))
	//}
	//fmt.Println()
	//}
}
