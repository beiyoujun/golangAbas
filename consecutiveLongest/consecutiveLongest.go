package main

import (
	"fmt"
)

func main() {
	array := []int{5,2,3,2,3,10,1,12,4}


	myMap := make(map[int]int)

	for _,value := range array {
		myMap[value] = 1
	}
	longest := 0

	for _,i := range array{
		length := 1
		for j:=i-1; myMap[j] == 1; j--  {
			delete(myMap, j)
			length++
		}
		for j :=i+1;myMap[j] == 1 ;j++  {
			delete(myMap,j)
			length++
		}
		if(length > longest){
			longest = length
		}
	}
	fmt.Println(longest)

}
