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

	for i:=0;i<10;i++ {
		for j:=3;j<13;j++{
			c1 := factorial(j)
			c2 := calculate(c1)
			for n := range c2 {
				fmt.Println(n)
			}
		}
	}

}
