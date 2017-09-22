package main

import (
	"encoding/json"
	"os"
	"strings"
	"fmt"
)

type person struct {
	Fname string
	Lname string
	Age int
	notexported int
}
func main() {

	p1:=person{"shubham","garg",22,055}
	json.NewEncoder(os.Stdout).Encode(p1)

	var p2 person
	rdr:= strings.NewReader(`{"Fname":"kiran","Lname":"kannade","Age	":23}`)
	json.NewDecoder(rdr).Decode(&p2)
	fmt.Println(p2)
}
