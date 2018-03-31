package main

import "fmt"

func main() {
	teste1 := []int{2, 4, 1, 6, 8, 5, 3, 7}

	fmt.Println(teste1)
	fmt.Println(MergeSort(teste1))

}

func MergeSort(list []int) []int {
	listSize := len(list)

	if listSize < 2 {
		return list
	}

	mid := listSize / 2
	left := make([]int, mid)
	right := make([]int, listSize-mid)

	for i := 0; i < mid; i++ {
		left[i] = list[i]
	}

	for j := mid; j < listSize; j++ {
		right[j-mid] = list[j]
	}

	fmt.Println("mid", mid)
	fmt.Println("left", left)
	fmt.Println("right", right)

	result := merge(MergeSort(left), MergeSort(right), make([]int, listSize))

	return result

}

func merge(left []int, right []int, list []int) []int {

	nLeft := len(left)
	nRight := len(right)

	index := 0
	indexLeft := 0
	indexRight := 0

	for {
		if indexLeft < nLeft && indexRight < nRight {
			if left[indexLeft] <= right[indexRight] {
				list[index] = left[indexLeft]
				indexLeft++
			} else {
				list[index] = right[indexRight]
				indexRight++
			}
			index++
		} else {
			break
		}
	}

	for {
		if indexLeft < nLeft {
			list[index] = left[indexLeft]
			indexLeft++
			index++
		} else {
			break
		}
	}

	for {
		if indexRight < nRight {
			list[index] = right[indexRight]
			indexRight++
			index++
		} else {
			break
		}
	}
	return list
}
