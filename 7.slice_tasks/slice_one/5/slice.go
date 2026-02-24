package main

import "fmt"

func main() {
	arr := []int{1, 2, 3} //len:3, cap:3 [1,2,3]
	src := arr[:1]        //len:1, cap:3 [1]
	foo(src)
	fmt.Println(src) //[1] (src видит только индекс 0)
	fmt.Println(arr) //[1 5 3] (len:3, arr видит все 3 элемента, включая перзаписанный)
}

func foo(src []int) {
	src = append(src, 5) //у src len:2, но результат присваивается локальной переменной, в main src остается с len:1
}
