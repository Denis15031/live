package main

import (
	"fmt"
)

func main() {
	nums := make([]int, 1, 3) //len:1, cap:3
	fmt.Println(nums)         // [0]
	appendSlice(nums, 1)
	fmt.Println(nums)            // [0]
	copySlice(nums, []int{2, 3}) //copy работает на прямую с памятью, изменения видны в main
	fmt.Println(nums)            // [2] - (len:1, изменили значение с 0 на 2)
	mutateSlice(nums, 1, 4)      //4 присвоили первому индексу
	fmt.Println(nums)            //[2] - (len по прежнему 1)
}

func appendSlice(sl []int, val int) {
	sl = append(sl, val) // аппенд пишет в array[sl.len], но результат присваивается локальной переменной
}

func copySlice(sl, cp []int) {
	copy(sl, cp) //копирует данные напрямую, без создания нового слайса
}

func mutateSlice(sl []int, idx, val int) {
	sl[idx] = val //прямая индексация меняет массив по указанному индексу
}
