package main

import "fmt"

type MyStruct struct {
	MyInt int
}

func func1() MyStruct {
	return MyStruct{MyInt: 1} // Создает и возвращает новую структуру MyStruct c полем MyInt = 1
}

func func2() *MyStruct {
	return &MyStruct{} // Создает новую структуру MyStruct (с MyInt = 0 по умолчанию), и возвращает указатель на неё
}

func func3(s *MyStruct) {
	s.MyInt = 333 // Принимает указатель на структуру MyStruct. Изменяет значение поля MyInt у структуры, на которую указывает s. Это изменение влияет на оригинальный объект, переданный в функцию
}

func func4(s MyStruct) {
	s.MyInt = 923 // Принимает структуру MyStruct по значению(создается копия). Изменяет MyInt в этой копии. Оригинальный объект, переданный в функцию, не изменяется
}

func func5() *MyStruct {
	return nil // Возвращает nil указатель
}

func main() {
	ms1 := func1()
	fmt.Println(ms1.MyInt) // 1 (func1 возвращает структуру, копируем ее в ms1)

	ms2 := func2()
	fmt.Println(ms2.MyInt) // 0 (func2 возвращает указатель на новую структуру, сохраняем его в ms2. Значение по умолчанию для int)

	func3(ms2)
	fmt.Println(ms2.MyInt) //333 (Передаем указатель ms2 в func3. func3 изменяет поле MyInt структуры, на которую указывает ms2)

	func4(ms1)
	fmt.Println(ms1.MyInt) // 1 (Передаем значение ms1 в func4. func4 изменяет копию, ms1 остается неизменным)

	ms5 := func5()
	fmt.Println(ms5.MyInt) // panic (Попытка обратиться к полю MyInt через указатель ms5, который равен nil. Это попытка получить доступ к памяти по адресу 0, что приводит к панике времени выполнения)
}
