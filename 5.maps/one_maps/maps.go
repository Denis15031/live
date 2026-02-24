package main

import (
	"fmt"
	"sort"
	"sync"
)

var (
	people   = make(map[string]int, 10)
	peopleMu sync.RWMutex
)

func init() {

	AddPerson("Mahmud", 20)
	AddPerson("Dimon", 29)
	AddPerson("Jack", 24)

}

func AddPerson(name string, age int) {
	peopleMu.Lock()
	defer peopleMu.Unlock()

	people[name] = age
	fmt.Printf("Добавлен: %s, возраст: %d\n", name, age)
}

func GetAge(name string) int {
	peopleMu.RLock()
	defer peopleMu.RUnlock()

	age, ok := people[name]
	if !ok {
		fmt.Printf("Человек '%s' не найден", name)
		return 0
	}
	return age
}

func DeletePerson(name string) {
	peopleMu.Lock()
	defer peopleMu.Unlock()

	delete(people, name)
	fmt.Printf("Удален: %s\n", name)
}

func PrintAll() {
	peopleMu.RLock()
	defer peopleMu.RUnlock()

	if len(people) == 0 {
		fmt.Println("Мапа пустая")
		return

	}

	fmt.Println("Все записи:")
	fmt.Println("------")

	//Получаем и сортируем ключи для предсказуемого вывода
	names := make([]string, 0, len(people))
	for name := range people {
		names = append(names, name)
	}
	sort.Strings(names)

	//Выводим отсортированные записи
	for _, name := range names {
		fmt.Printf("%s: %d лет", name, people[name])
	}
	fmt.Println("------")
}

func main() {
	fmt.Println("Запуск тестирования функций работы с Мапой")

	PrintAll()

	fmt.Printf("\n Возраст Mahmud: %d лет\n\n", GetAge("Mahmud"))

	fmt.Printf("\n Возраст Unknown: %d лет\n\n", GetAge("Unknown"))

	fmt.Println("Добавляем новых людей:")
	AddPerson("Saida", 27)
	AddPerson("Oksana", 34)
	fmt.Println()

	PrintAll()

	fmt.Println("Меняем возраст Mahmud:")
	AddPerson("Mahmud", 40)
	PrintAll()

	fmt.Println("Удаляем Dimon:")
	DeletePerson("Dimon")
	PrintAll()
}
