package main

import (
	"fmt"
	"strconv"
)

func main() {
	x := 5
	i, err := x, false // <<<<< extractLocal,10,5,10,7,newVar,fail
	if err == true {
		fmt.Println(err)
	}
	fmt.Println("divisible by 5:" + strconv.Itoa(i))
}