package main

import (
	"fmt"
)

func main() {
	slice := make([]string, 3, 4) // ["","","",_]
	fmt.Println(slice)            // [   ] // (len:3, cap:4, выводит индекса 0, 1, 2)

	appendSlice(slice) // ["","","", "privet']
	fmt.Println(slice) // [   ] //len:3, выводит только 0, 1, 2 индексы, "privet" - индекс 3

	mutareSlice(slice) //["vasya", "", "", "privet"] - изменения видно в main, так как array общий
	fmt.Println(slice) //["vasya"  ] //так же len:3
}

func appendSlice(slice []string) {
	slice = append(slice, "privet") //Внутри: slice.len=3, slice.cap=4, места хватает, append пишет "privet" в array[3]
}
func mutareSlice(slice []string) {
	slice[0] = "vasya" //slice[0] = "vasya" меняет underlying array напрямую
}
