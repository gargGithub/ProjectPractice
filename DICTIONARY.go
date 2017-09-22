package main

import (
	"fmt"

)

func main() {

	dictionary:= map[string]map[int]string{
		"a": map[int]string{
			1: "array",
			2: "algorithm",
			3: "awareness",
		},

		"b": map[int]string{
			1: "bore",
			2: "ball",
			3: "basket",
		},
	}
    var ch string
	fmt.Scanf("%s",&ch)



	for i,n:=range dictionary[ch]{

		fmt.Println(i,":",n)
	}

	var num int
	fmt.Println("Enter the number corresponding to the required word: ")
	fmt.Scanln(&num)


	fmt.Println(dictionary[ch][num])




}
