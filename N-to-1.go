package main

import (
	"sync"
	"fmt"
)

var wg sync.WaitGroup
func main() {
	c:=make(chan int)
	wg.Add(2)
	go func(){
		for i:=0;i<11;i++{
			c<-i
		}
		wg.Done()
	}()
	go func(){
		for i:=0;i<11;i++{
			c<-i
		}
		wg.Done()
	}()
	go func(){

		wg.Wait()
		close(c)
	}()

	for n:=range c{
		fmt.Println(n)
	}
}
