package main

import (
	"fmt"
	"github.com/idasilva/go-jump-start/principles"
)

func main (){

	fmt.Println("principles of the go language")
	variable := principles.Add(5,5)
	fmt.Println(variable)

	boolVariable := principles.Boolean()
	fmt.Println(boolVariable)

	principles.Showvar()

	principles.LoopForTraditional()
	principles.LoopForWithBreak()
	principles.LoopForWithContinue()
	principles.LoopForWithRange()

}