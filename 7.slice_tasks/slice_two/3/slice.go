package main

import "fmt"

func test(testSlice []string) []string {
	return append(testSlice, "Пока") //возвращаем результат из функции

	//testSlice = append(testSlice, "Пока") - ненужная строка
}
func main() {
	testSlice := make([]string, 0, 3)       // [_,_,_]
	testSlice = append(testSlice, "Привет") // ["Привет", _,_]
	testSlice = append(testSlice, "Привет") // ["Привет", "Привет", _]
	testSlice = test(testSlice)
	test(testSlice) //testSlice len:3, но результат присваивается локальной переменной, в main testSlice остается len:2

	fmt.Println(testSlice) //["Привет", "Привет"] - старый вывод
	//["Привет", "Привет", "Пока"] - новый корректный вывод
}
