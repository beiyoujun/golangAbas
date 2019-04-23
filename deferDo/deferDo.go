package deferDo

import "fmt"

func deferdo() {
	defer func() {
		fmt.Println("c")
		if err := recover(); err != nil {
			fmt.Println(err)
		}
		fmt.Println("d")
	}()

	f()
}

func f() {
	fmt.Println("a")
	panic("55")
	fmt.Println("e")
}

func tr() {
	fmt.Println("1")
	r()
	fmt.Println("2")
}

func r() {
	defer func() {
		fmt.Println("打印前")
	}()
	defer func() {
		fmt.Println("打印后")
	}()
	return
}
