package main

import (
	"fmt"
	"sync"
	"time"
	"runtime"
	"sync/atomic"
)
var mutex sync.Mutex
var wg sync.WaitGroup
var counter int64

func init() {
	runtime.GOMAXPROCS(runtime.NumCPU())
}
func foo(){
	i:=1
	for i<=10{
		//mutex.Lock()
		atomic.AddInt64(&counter,1)
		fmt.Println("foo: ",i, "Counter: ",counter)
		//mutex.Unlock()
		time.Sleep(time.Duration(time.Millisecond*60))
		i++
	}
	wg.Done()

}
func bar(){
	i:=11
	for i<=20{

		//mutex.Lock()
		atomic.AddInt64(&counter,1)
		fmt.Println("bar: ",i, "Counter: ",counter)
		//mutex.Unlock()
		time.Sleep(time.Duration(time.Millisecond*60))
		i++
	}
	wg.Done()

}

func main() {
    wg.Add(2)
	go foo()
	go bar()
	wg.Wait()
	fmt.Println("final value of counter is: ",counter)
	fmt.Println("this is the end of main")
}
