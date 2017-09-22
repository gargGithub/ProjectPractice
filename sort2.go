package main

import (
	"fmt"
	"sort"
)

type person struct {
	Name string
	Age int
}

func (p person) String() string{
	return fmt.Sprintf("details: %s %d",p.Name,p.Age)
}


type byage []person

func (a byage) Len() int{
	return len(a)
}

func (a byage) Swap(i,j int){
	a[i],a[j]=a[j],a[i]

}

func (a byage) Less(i,j int) bool{
	return a[i].Age<a[j].Age
}

func main() {


	people:=[]person{

		{"shubham",22},
		{"kiran",21},
		{"ambuj",23},
		{"mukul",20},


	}

	fmt.Println(people[0])
	fmt.Println(people)
    sort.Sort(byage(people))
	fmt.Println(people)
}
