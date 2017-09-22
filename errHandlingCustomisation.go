package main

import (
	"fmt"
	"log"
	"math"
)

type ErrSqrtNumber struct {
	lat,long string
	err error
}


func(e *ErrSqrtNumber) Error() string{

	return fmt.Sprintf("Error Occured: %v %v %v",e.lat,e.long,e.err)
}


func sqrt(n float64)(float64,error){
	if(n<0){

		nme:=fmt.Errorf("Square root of negative number: %v",n)
	      return 0,&ErrSqrtNumber{"50.4353N","98.3433W",nme}

	}else{
		return math.Sqrt(n),nil
	}
}


func main() {

	num,err:=sqrt(-45)
	if(err!=nil){
		log.Fatal(err)
	}else{
		fmt.Println(num)
	}

}
