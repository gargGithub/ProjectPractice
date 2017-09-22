package main

import "fmt"

type foo struct {
	fname string
	lname string
	age int
}


func main() {

   p1:=&foo{"shubham","garg",22}
	fmt.Println(p1)
	fmt.Printf("%T\n",p1)
	fmt.Println(p1.fname)
	fmt.Println(&p1.fname)

}
