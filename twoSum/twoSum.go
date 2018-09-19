package main

import "fmt"

func main() {

	array := [10]int{1,2,3,4,5,6,7,8,9,20}
	target := 29
	index,index2 := twoSum(array,target)
	fmt.Println(index)
	fmt.Println(index2)

}

func twoSum(array [10]int,target int)(index int,index2 int){

	myMap := make(map[int]int)

	for i,v := range array {
		myMap[v] = i
	}

	for i,v := range array{
		if(myMap[target-v] > 0){
			return i,myMap[target-v]
		}
	}
	return 0,0
}