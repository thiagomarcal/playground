package main

import "fmt"

// Given an array of integers, return indices of the two numbers such that they add up to a specific target.
func main() {

	teste1 := []int{1, 4, 5, 3, 2}
	target := 4

	result := twoSum(teste1, target)
	fmt.Println(result)

}

func twoSum(nums []int, target int) []int {
	mapa := make(map[int]int)
	var result []int

	for index, value := range nums {
		complement := target - value
		i, ok := mapa[complement]
		if ok {
			result = []int{i, index}
			return result
		}
		mapa[value] = index
	}

	return result
}
