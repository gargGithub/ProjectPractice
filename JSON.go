package main

import (
	"encoding/json"
	"fmt"
	//"math"
)

type person struct {
	Name string
	Age int
	notExported int
}

func main() {
    var p2 person
	p1:=person{"shubham garg",22,055}
	bs,_:=json.Marshal(p1)
	fmt.Println(bs)

	fmt.Printf("%T\n",bs)
	fmt.Println(p1.Age)
	bs1:=[]byte(`{"Name":"kiran","Age":24}`)
	json.Unmarshal(bs1,&p2)
	fmt.Println(p2)
}
