package main

import (
	//"errors"
	"math"
	"fmt"
	"log"
)
/*var (
	ErrNorGateMath=errors.New("Square root of negative number is not allowed!!!!")
)*/
func sqrt(n float64) (float64, error){
	 if(n<0){
		 ErrNorGateMath:=fmt.Errorf("Square root of negative number: %v",n)
	 	return 0, ErrNorGateMath
	 }

	 c:=math.Sqrt(n)
	 return c,nil
}


func main() {
	num,err:=sqrt(-4.23)

	if(err!=nil) {
		log.Fatal(err)
	}else {
		fmt.Println(num)
	}
}
