package main

import (
	"sync"

	"fmt"
)

func main() {

	in := gen()
	c1 := fanFactorial(in)
	c2 := fanFactorial(in)
	c3 := fanFactorial(in)
	c4 := fanFactorial(in)
	c5 := fanFactorial(in)
	c6 := fanFactorial(in)
	c7 := fanFactorial(in)
	c8 := fanFactorial(in)
	c9 := fanFactorial(in)
	c10 := fanFactorial(in)
       i:=0
	for n := range merge(c1, c2, c3, c4, c5, c6, c7, c8, c9, c10) {
		i++
          fmt.Println(i,": ",n)
	}
}

func gen() chan int{
	out:=make(chan int)
	go func(){
		for i:=0;i<100;i++{
			for j:=4;j<16;j++{
				out<-j
			}
		}
		close(out)

	}()
	return out
}

func fanFactorial(c chan int) chan int{

	out:=make(chan int)
	go func(){
		for n:=range c{
			out<-fact(n)
		}
		close(out)
	}()
	return out
}


func fact(n int) int{
	total:=1
	for i:=n;i>0;i--{
		total*=n
	}
	return total
}

func merge(nums... chan int) chan int {
	var wg sync.WaitGroup
	out := make(chan int)

	output := func(c chan int) {
		for n := range c {
			out <- n
		}
		wg.Done()
	}
	wg.Add(len(nums))
	for _, j := range nums {
		go output(j)
	}

	go func() {
		wg.Wait()
		close(out)
	}()
	return out
}



