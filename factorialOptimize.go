package main

import (

	"fmt"
)

func factorial(n int) chan int{
	out:=make(chan int)
	go func(){
		total:=1
		for i:=n;i>0;i--{
			total*=i

		}
		out<-total
		close(out)
	}()
	return out
}
func main() {
	c1:=factorial(6)
	for n:=range c1{
		fmt.Println(n)
	}
}
