package main

import (
	"fmt"
	"unsafe"
)

func main() {
	//1
	first := []int{1, 2, 3, 4, 5}
	first = nil
	fmt.Println("first:", first, ":", len(first), ":", cap(first)) // first: [] : 0 : 0 //nil срез полностью пустой, так как он обнуляет всё
	//2
	second := []int{1, 2, 3, 4, 5}
	second = second[:0]
	fmt.Println("second:", second, ":", len(second), ":", cap(second)) // second: [] : 0 : 5 // len = 0 != nil, cap и данные остаются

	//3
	third := []int{1, 2, 3, 4, 5}
	clear(third)
	fmt.Println("third:", third, ":", len(third), ":", cap(third)) // third: [0 0 0 0 0] : 5 :5 // clear != удаление, просто обнуляем элементы

	//4
	fourth := []int{1, 2, 3, 4, 5}
	clear(fourth[1:3])
	fmt.Println("fourth:", fourth, ":", len(fourth), ":", cap(fourth)) // fourth: [1 0 0 4 5] : 5 : 5 // clear обнуляет только индексы 1 и 2

	//5
	slice := make([]int, 3, 6)
	array := [3]int(slice[:3])
	slice[0] = 10

	fmt.Println("slice = ", slice, len(slice), cap(slice)) // slice : [10 0 0] : 3 : 6
	fmt.Println("array =", array, len(array), cap(array))  // array: [0 0 0] 3 : 3

	//6 В каких случаях Slice пустой или нулевой
	//1
	var data []string
	fmt.Println("var data []string:")                                                                                           // var data []string:
	fmt.Printf("\tempty=%t nil=%t size=%d data=%p\n", len(data) == 0, data == nil, unsafe.Sizeof(data), unsafe.SliceData(data)) // empty=true nil=true size=24 data=0x0// срез не инициализирован; zero value = nil; фикс размер по 8 байт у ptr, len и cap; указатель nil
	//2
	data = []string(nil)
	fmt.Println("data = []string(nil):")                                                                                        // data = []string(nil):
	fmt.Printf("\tempty=%t nil=%t size=%d data=%p\n", len(data) == 0, data == nil, unsafe.Sizeof(data), unsafe.SliceData(data)) // empty=true nil = true size=24 data=0x0 //нет элементов; явно присвоили nil; осталось так же
	//3
	data = []string{}
	fmt.Println("data = []string{}")                                                                                            // data = []string{}
	fmt.Printf("\tempty=%t nil=%t size=%d data=%p\n", len(data) == 0, data == nil, unsafe.Sizeof(data), unsafe.SliceData(data)) // empty=true nil=false size=24 data=0x7ff697c14bc0 // нет элементов; создан не nil срез; дескриптор тот же; указатель уже не nil
	//4
	data = make([]string, 0)
	fmt.Println("data =make([]string,0)")                                                                                       // data =make([]string, 0)
	fmt.Printf("\tempty=%t nil=%t size=%d data=%p\n", len(data) == 0, data == nil, unsafe.Sizeof(data), unsafe.SliceData(data)) // empty=true nil=false size=24 data=0x7ff697c14bc0 //явно указали; make всегда возвращает non-nil; дескриптор тот же; тот же адрес

	empty := struct{}{}
	fmt.Println("empty struct address ", unsafe.Pointer(&empty)) //empty struct address 0x7ff697c14bc0
}
