package main

import (
	"fmt"
)

type Integer int

func (a *Integer) ToPrint() {
	fmt.Println(a)
}

func main() {
	var a Integer
	a.ToPrint()            //方法值形式 (method value)
	(*Integer).ToPrint(&a) //方法表达式形式 (method expression)
}
