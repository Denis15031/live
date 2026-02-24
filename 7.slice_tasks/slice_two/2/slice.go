package main

import (
	"fmt"
	"strings"
)

func chengeSlice(arr []string) {
	arr[0] = "Goodbye"
}

func appendSomeData(arr []string) []string {
	return append(arr, "!") // возвращаем результат из функции append
	//arr = append(arr, "!") - ненужная строка
}

func main() {
	someSlice := []string{"Hello", "World"} //len:2, cap:2; 0-Hello, 1-World
	someSlice = appendSomeData(someSlice)
	chengeSlice(someSlice)                   //изменение видно в main, так как array общий
	appendSomeData(someSlice)                //нет места для аппенда, создается новый array. Не возвращается в main, остается Goodbye World
	fmt.Println(strings.Join(someSlice, "")) // до исправления видит только свой array : goodbye world
	//GoodbyeWorld! - (после прокинутого return в appendSomeData)
}

//someSlice := make([]string, 2, 3) - можно было б и так, задать сразу запас капасити.
