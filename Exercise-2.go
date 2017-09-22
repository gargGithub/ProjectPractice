package main

import "fmt"

type vehicle struct {
	doors int
	color string
}
type truck struct {
	vehicle
	fourWheel bool
}
type sedan struct {
	vehicle
	luxury bool
}
type transportation interface {
	transportationDevice() string

}

func report(trans...transportation){
	for _, v:=range trans{
	    fmt.Println(v.transportationDevice())
	}
}
func (t truck) transportationDevice() string{
	return fmt.Sprintln("truck with ",t.doors," is us	ed for commercial transporation")
}
func (s sedan) transportationDevice() string{
	return fmt.Sprintln("sedan with ",s.doors," is luxurious transport medium")
}
func main() {
	t:=truck{
		vehicle{4,"black"},
		true,
	}
	S:=sedan{
		vehicle{4,"White"},
		false,
	}
	tck:=t.transportationDevice()
	car:=S.transportationDevice()
	fmt.Println(tck)
	fmt.Println(car)
  report(t,S)

}
