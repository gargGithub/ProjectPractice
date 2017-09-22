package main

import (
	"fmt"

)

func factorial(n int)chan int {
	out := make(chan int)
	go func() {


		for i := n; i > 0; i-- {
			out <- i
		}
		close(out)
	}()

   return out
}

func calculate(c chan int) chan int{
	out:=make(chan int)
	total:=1
	go func(){

		for n:=range c{
			total = total*n

		}
		out<-total
		close(out)
	}()
	return out
}


func main() {
	var n int
	fmt.Scan(&n)
	c1:=factorial(n)
	c2:=calculate(c1)
	for n:=range c2{
		fmt.Print(n)
	}

}
