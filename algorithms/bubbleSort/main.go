package main

import "fmt"

func main() {

	teste1 := []int{2, 4, 1, 6, 8, 5, 3, 7}

	fmt.Println(teste1)

	bubbleSort(teste1)

	fmt.Println(teste1)

}

func bubbleSort(array []int) int {
	numSwaps := 0
	n := len(array)

	for i := 0; i < n; i++ {
		for j := 0; j < n-1; j++ {
			if array[j] > array[j+1] {
				numSwaps++
				swap(&array[j], &array[j+1])
			}
		}
	}

	return numSwaps

}

func swap(a *int, b *int) {
	*(b), *(a) = *(a), *(b)
}
