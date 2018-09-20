package main

import "fmt"

func main() {
	array := []int{5, 4, 3, 2, 1}
	nextPermutationOne(array)

}

func nextPermutationOne(array []int) {
	k := len(array) - 1
	i := k
	for i >= 0 {
		if i < k && array[i] < array[i+1] {
			j := i
			for j <= k && array[i] <= array[j] {
				j++
			}
			array[i], array[j-1] = array[j-1], array[i]
			break
		}
		i--
	}
	reverse(array, i+1, k)
	fmt.Println(array)
}

func reverse(array []int, begin int, end int) {
	for begin < end {
		array[begin], array[end] = array[end], array[begin]
		end--
		begin++
	}
}
