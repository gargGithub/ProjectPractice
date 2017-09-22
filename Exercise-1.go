package main

import "fmt"

type person struct{
	fname string
	lname string
	favFood []string
}

func (p person) walk() string {
	return fmt.Sprintln(p.fname," is walking")

}
func main() {
	arr:=[]int{2,3,4,5,6}
	for i,v:=range arr{
		fmt.Println(i,v)
	}

	m:=map[string]int{
		"Shubham":22,
		"Kannade":21,
		"Aman":21,
		"mukul":22,
	}
	for i,v :=range m {
		fmt.Println(i,":",v)
	}

	p1:=person{"Shubham","Garg",[]string{"Pizza","burger"}}
	for _,v:=range p1.favFood{
		fmt.Println(v)
	}

	s:=p1.walk()
	fmt.Println(s)
}
