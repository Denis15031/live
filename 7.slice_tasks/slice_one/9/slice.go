package main

import (
	"fmt"
)

func main() {
	slice := make([]int, 3, 4) //len: 3, cap:4
	appendingSlice(slice[:1])  //len:1, cap:4
	fmt.Println(slice)         //[0 1 0] - (в slice len:3, но массив был изменен в функции)
}

func appendingSlice(slice []int) {
	slice = append(slice, 1) //len:1, cap:4, места для аппенда хватает
}
