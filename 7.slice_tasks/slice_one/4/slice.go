package main

import "fmt"

func main() {
	a := []int{1, 2, 3} // len: 3, cap:3
	b := a[:2]          //len:2, cap:3
	b = append(b, 4)    // места хватает, релокация не нужна
	fmt.Println(b)      //[1 2 4]
	fmt.Println(a)      //[1 2 4] (массив тот же, который и перезатерли, поэтому "4" а не "3")
}
