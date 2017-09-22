package main

import "fmt"

type gator int
type flamingo bool

func (gator) greeting(){
	fmt.Println("this is method")
}
func (flamingo) greeting(){
	fmt.Println("this is flamingo method")
}
type swampCreature interface {
	greeting()
}
func bayou(swamp swampCreature){
	swamp.greeting()
}
func main() {

	var g1 gator
	g1 =17
	fmt.Println(g1)
	fmt.Printf("%T\n",g1)

 var c gator
	var c1 flamingo
	bayou(c1)
	bayou(c)


}
