package main

import "fmt"

func main() {
	var a = [10]int{1,2,4,4,5,6,7,8,9,10}
	var b [10]int
	var i = 0
	for _,value := range  a{
		cf := 0
		for x:=0; x<i; x++ {
			if(value == b[x]){
				cf = 1
			}
		}
		if(cf == 0){
			b[i] = value
			i++
		}
	}
	fmt.Println(b)
}