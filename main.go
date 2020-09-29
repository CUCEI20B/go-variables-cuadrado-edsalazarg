package main

import "fmt"

func main()  {
	var num uint64

	fmt.Scanln(&num)

	area := num * num

	fmt.Println(area)
}