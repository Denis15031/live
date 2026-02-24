package main

import "fmt"

func main() {
	first := []int{10, 20, 30, 40}
	second := make([]*int, len(first))
	for i, v := range first {
		v := v // новая переменная, копия значения
		second[i] = &v
	}
	fmt.Println(*second[0], *second[1]) // старый вывод: 40 40// После цикла: v = 40, *addr_X = 40, Все 4 указателя в second указывают на один адрес!
	//новый вывод: 10 20
}

//можно и по другому, брать адрес по индексу исходного слайса
/* for i := range first {
second[i] = &first[i]  // адрес конкретного элемента массива
*/
