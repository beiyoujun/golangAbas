package main

import "fmt"

func main() {

	fmt.Printf(jiujiu(1,1,"")) //递归方法

	funcTwo() // for嵌套的方法都是

}

func jiujiu(i int,a int,result string) (string) {
	if(i<10){
		if(a<=i){
			return jiujiu(i,a+1,fmt.Sprintf("%s%dx%d=%d\t",result,a,i,a*i))
		}else{
			return jiujiu(i+1,1,fmt.Sprintf("%s%s",result,"\n"))
		}
	}else{
		return result
	}
}

func funcTwo()  {
	for i:=1 ; i<10 ;i++  {
		for j:=1; j<=i;j++  {
			fmt.Printf("%d*%d=%d\t",j,i,i*j)
		}
		fmt.Printf("\n")
	}
}