package main

import (
	"fmt"

)

func generate(nums...int) chan int{
	out:=make(chan int)
	go func(){
		for _,n:=range nums{
			out<-n
		}
		close(out)
	}()
	return out
}


func fetchSquare(c chan int) chan int{
	out:=make(chan int)
	go func(){
		var square int
		for n:=range c{
			square = n*n
			out<-square
		}

		close(out)
	}()
	return out
}
func main() {

	c1:=generate(2,3,4)
	c2:=fetchSquare(c1)
	out:=fetchSquare(c2)
	for n:=range out{                  /* OR REFACTED CODE:      for n:= range fetchSquare(fetchSquare(generate(2,3,4))){
		                                                                   fmt.Println(n)
	                                                                          }*/
	fmt.Println(n)
}

}
