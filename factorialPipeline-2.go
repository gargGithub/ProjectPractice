package main

import "fmt"

func main() {

	c1:=generate()
	c2:=get(c1)
	for n:=range c2{
		fmt.Println(n)
	}
}


func generate() chan int{
	out:=make(chan int)
	go func(){
		for i:=0;i<10;i++{
			for j:=0;j<10;j++{
				out<-j
			}
		}
		close(out)
	}()
	return out
}

func get(c chan int) chan int{
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
	var total int =1
	for i:=n;i>0;i--{
		total*=i
	}
	return total
}



