package main

import (
	"fmt"
	"sort"
)

type SliceInt [10]int

func (c SliceInt)Len()int{
	return len(c)
}

func (c SliceInt)Swap(i,j int){
	c[i],c[j] = c[j],c[i]
}

func (c SliceInt)Less(i,j int) bool {
	return c[i] < c[j]
}

func main() {
	array := SliceInt{1,2,3,4,5,6,7,8,-1,20}

	target := 0

	sort.Sort(array)

	minGap := 10000

	result := 0

	for i := 0; i<len(array)-1; i++  {
		j :=i+1
		k := len(array) - 1

		for j<k{
			threeSum := array[i] + array[j] + array[k]
			gap := getAbs(target - threeSum)
			if(gap < minGap){
				minGap = gap
				result = threeSum
			}
			if(threeSum > target){
				k--
			}else{
				j++
			}
		}
	}
	fmt.Println(result)
}

func getAbs(int2 int)int  {
	if(int2 < 0){
		return -int2
	}else{
		return int2
	}
}

