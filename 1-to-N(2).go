	package main

import (

	"fmt"

)


func main() {
	c := make(chan int)
	done := make(chan bool)
	n:=10
	go func() {
		for i := 0; i < 100; i++ {
			c <- i
		}
		close(c)

	}()


	for j:=0;j<n;j++ {
		go func() {

			for n := range c {
				fmt.Println(j,": ", n)
			}
			done <- true
		}()
	}
for j:=0;j<n;j++ {
	<-done
}

}
