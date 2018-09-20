package main

import (
	"fmt"
	"sort"
)

type Integer []int

func (c Integer) Len() int {
	return len(c)
}

func (c Integer) Less(i, j int) bool {
	return c[i] < c[j]
}

func (c Integer) Swap(i, j int) {
	c[i], c[j] = c[j], c[i]
}

func main() {

	array := Integer{10, 20, 30, 40, -3, -9, -20, 3, 9}
	target := 0

	var result [][4]int

	sort.Sort(array)

	for i := 0; i < len(array)-3; i++ {
		if i > 0 && array[i] == array[i-1] {
			continue
		}
		for j := i + 1; j < len(array)-2; j++ {
			if j > i+1 && array[j] == array[j-1] {
				continue
			}
			k := j + 1
			l := len(array) - 1
			for k < l {
				sum := array[i] + array[j] + array[k] + array[l]
				if sum < target {
					k++
					for k < l && array[k] == array[k-1] {
						k++
					}
				} else if sum > target {
					l--
					for k < l && array[l] == array[l+1] {
						l--
					}
				} else {
					fourInt := [4]int{array[i], array[j], array[k], array[l]}
					result = append(result, fourInt)
					k++
					l--
					for k < l && array[l] == array[l+1] {
						l--
					}
					for k < l && array[k] == array[k-1] {
						k++
					}
				}
			}
		}

	}
	fmt.Println(result)
}
