package main

import (
	"athena_go/algorithm/template"
	"fmt"
)

func main() {

	// sword2offer.VerifySquenceOfBST()
	println("hello")
	slice := []int{1, 2, 3, 4, 5, 6}
	fmt.Printf("the slice %v \n", slice[1:3])

	c := template.CombinationSum{}
	candidates := []int{2, 3, 6, 7}
	c.CombinationSum2(candidates, 7)
}
