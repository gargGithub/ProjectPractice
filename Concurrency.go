package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

func table2(){
	i:=1
	for i<=10{
		fmt.Println(i*2)
		i++
		time.Sleep(time.Duration(time.Millisecond*60))
	}
	wg.Done()
}
func table3(){
	i:=1
	for i<=10{
		fmt.Println(i*3)
		i++
		time.Sleep(time.Duration(time.Millisecond*60))
	}
	wg.Done()
}

func main() {
    wg.Add(2)
	go table2()
	go table3()
	wg.Wait()
}
