package main

import "fmt"

func insertSort(arr []int) []int {
	for i := 1; i < len(arr); i++ {
		key := arr[i]
		j := i - 1
		for j >= 0 && arr[j] > key {
			arr[j+1] = arr[j]
			j--
		}
		arr[j+1] = key
	}
	return arr
}
func checkDistance(arr []int) (bool, int) {
	if len(arr) <= 2 {
		return true, 0
	}
	distance := arr[1] - arr[0]
	for i := 1; i < len(arr)-1; i++ {
		if (arr[i+1] - arr[i]) != distance {
			return false, 0
		}
	}
	return true, distance
}
func main() {
	var num int
	var data []int
	for {
		fmt.Scan(&num)
		if num < 0 {
			break
		}
		data = append(data, num)
	}
	sortedData := insertSort(data)
	for i, val := range sortedData {
		if i > 0 {
			fmt.Print(" ")
		}
		fmt.Print(val)
	}
	fmt.Println()
	isConstant, distance := checkDistance(sortedData)
	if isConstant {
		fmt.Printf("Data berjarak %d\n", distance)
	} else {
		fmt.Println("Data berjarak yang tidak tetap")
	}
}
