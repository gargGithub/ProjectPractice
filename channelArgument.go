package main

import "fmt"

func main() {


	c:=incrementor()
	csum:=puller(c)
	fmt.Print(<-csum)


}

func incrementor() <-chan int{
	out:=make(chan int)
	go func(){
		for i:=1;i<11;i++{
			out<-i
		}
		close(out)
	}()
	return out
}

func puller(c <-chan int) <-chan int{
	out:=make(chan int)

	go func(){
		var sum int
		for n:=range c{
			sum+=n


		}
		out<-sum
		close(out)
	}()
	return out
}
