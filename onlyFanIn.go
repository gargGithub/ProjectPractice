package main

import (
	"fmt"

)

func main() {

	c:=fanIN(boring("foo"),boring("out"))
	for i:=0;i<20;i++{

		fmt.Println(<-c)
	}
}




func boring(msg string) chan string{
	out:=make(chan string)
	go func(){
		for i:=0;;i++{
			out <- fmt.Sprintf("%s : %d",msg,i)
		}

	}()
	return out
}


func fanIN(input1,input2 chan string) chan string{
	out:=make(chan string)
	go func(){
		for{
			out<- <-input1

		}

	}()
	go func(){
		for{
			out<- <-input2

		}

	}()

	return out
}
