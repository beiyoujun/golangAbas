package main

import (
	"fmt"
)

type integer int

func (a *integer) toPrint() {
	fmt.Println(a)
}

func main() {
	//fmt
	var a integer
	a.toPrint()            //方法值形式 (method value)
	(*integer).toPrint(&a) //方法表达式形式 (method expression)
}
