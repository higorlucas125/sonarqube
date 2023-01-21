package main

import "fmt"

func main(){
	fmt.Println(Somar(2,5))
}

func Somar(a int,b int) int {
	return a + b ;
}

func Sub(a int,b int) int {
	return a - b ;
}

func Times(a int,b int) int {
	return a / b ;
}

func SumX(a int,b int) int {
	return a * b ;
}