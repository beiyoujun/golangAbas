package main

import "fmt"

func main() {

	array := []int{0,2,0,0,4,5,0,1,2,0,2,4}
	moveZeroOne(array)
	moveZeroTwo(array)
}

func moveZeroOne(array []int)  {
	k := len(array)
	for i := 0; i < k-2; i++ {
		if array[i] == 0 {
			for array[k-1] == 0{
				k--
			}
			array[i],array[k-1] = array[k-1],array[i]
		}
	}
	fmt.Println(array)
}

func moveZeroTwo(array []int)  {
	k := len(array);
	index := 0
	for i := 0; i < k ; i ++  {
		if(array[i] != 0){
			array[index] = array[i]
			index++
		}
	}

	for j := index;j<k;j++{
		array[j] = 0
	}
	fmt.Println(array)
}



