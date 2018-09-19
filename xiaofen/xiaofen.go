package main

import "fmt"

func main() {
	left := 10000
	right := 100
	leftNow,rightNow := xiaofen(left,right)
	leftDown := left-leftNow
	rightDown := right-rightNow
	downCount := (leftDown+rightDown) / 900
	fmt.Printf("LEFT:%d\tRIGHT%d\tDOWNCOUNT:%d",leftDown,rightDown,downCount);


}


func xiaofen(left int,right int) (leftDown int,rightDown int){
	if(left >= right){
		if(left >= 600 && right >= 300){
			left -= 600
			right -= 300
			return xiaofen(left,right)
		}else{
			return left,right
		}
	}else{
		if(right >= 600 && left >=300){
			right -= 600
			left -= 300
			return xiaofen(left,right)
		}else{
			return left,right
		}
	}
}
