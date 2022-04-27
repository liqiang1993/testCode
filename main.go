package main

import (
	"fmt"
)

func main() {
	m := 4
	temp := "0110"
	res := queryString(temp, m)
	fmt.Printf("%v\n", res)
}
