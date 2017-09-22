package main

import (
	"sort"
	"fmt"
)

type person []string

func main() {

	student:=person{"Shubham","Kiran","Ambuj","Ganesh"}
	fmt.Println(student)
	sort.Strings(student)
	fmt.Println(student)

	student2 := []string{"shubham garg","ambuj","shubham jain","mukul"}
	fmt.Println(student2)
	sort.Strings(student2)
	fmt.Println(student2)

	rollno:=[]int{12,34,2,1,78}
	fmt.Println(rollno)
	sort.Ints(rollno)
	fmt.Println(rollno)
	sort.Sort(sort.Reverse(sort.IntSlice(rollno)))
	fmt.Println(rollno)
	sort.Sort(sort.Reverse(sort.StringSlice(student)))
	fmt.Println(student)
	sort.Sort(sort.Reverse(sort.StringSlice(student2)))
	fmt.Println(student2)


}
