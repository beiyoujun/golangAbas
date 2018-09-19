package main

import "fmt"

func main() {

	array := []int{1,2,3,4,5,6,7,8,1,3,1}
	target := 1
	var newArray []int;
	for v :=  range array {
		if(v != target){
			newArray = append(newArray, v)
		}
	}
	fmt.Println(newArray)
}
