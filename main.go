package main

import "fmt"

func main() {
	s := []int{1, 2, 3}
	s = s[:len(s)]
	fmt.Println(s)
}