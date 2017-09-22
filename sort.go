package main

import (
	"sort"
	"fmt"
)

type people []string

func (p people)Len() int{
	return len(p)
}

func (p people) Swap(i, j int) {
	p[i],p[j] = p[j],p[i]
}

func (p people) Less(i,j int) bool{
	return p[i]<p[j]
}
func main() {

	studygroup:=people{"shubham","aman","banga","naruto"}
	sort.Sort(studygroup)
	fmt.Println(studygroup)
	fmt.Println("---------------")
	studygroup2:=people{"shubham","aman","banga","naruto","warrior"}
	sort.StringSlice(studygroup2).Sort()
	fmt.Println(studygroup2)
	fmt.Println("---------------")
	sort.Sort(sort.StringSlice(studygroup2))
	fmt.Println(studygroup2)
	fmt.Println("---------------")
	sort.Strings(studygroup2)
	fmt.Println(studygroup2)
	fmt.Println("---------------")
   sort.Sort(sort.Reverse(sort.StringSlice(studygroup)))
	fmt.Println(studygroup)

	s:=[]int{2,6,3,9,1,34}
	sort.Sort(sort.IntSlice(s))
	fmt.Println(s)
	fmt.Println("---------------")
	sort.IntSlice(s).Sort()
	fmt.Println(s)

}
