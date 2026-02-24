package main

import "fmt"

func main() {
	a := []string{"a", "b", "c"}   //len:3, cap:3 ["a", "b", "c"]
	b := a[1:2]                    // len:1, cap:2 ["b"]
	fmt.Println(b, cap(b), len(b)) //["b"] cap(2), len(1)
	b[0] = "q"                     //так как а и b делят один и тот же массив, изменения видны в обоих слайсах
	fmt.Println(a)                 //["a", "q", "c"] (индекс 0 в b, это первый индекс в a, поэтому q встаёт на место первого индекса в слайсе а)
}
