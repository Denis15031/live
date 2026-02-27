package main

import "fmt"

func RemoveUnordered[T any](s []T, i int) []T {
	if i < 0 || i >= len(s) {
		return s
	}
	s[i] = s[len(s)-1]  //копируем последний элемент на место удаляемого
	return s[:len(s)-1] //обрезаем слайс, последний элемент дубликат, его убираем
}

func RemoveOrdered[T any](s []T, i int) []T {
	if i < 0 || i >= len(s) {
		return s
	}
	copy(s[i:], s[i+1:])
	return s[:len(s)-1]
}

func RemoveAllByValue[T comparable](s []T, value T) []T {
	result := s[:0]
	for _, v := range s {
		if v != value {
			result = append(result, v)
		}
	}
	return result
}

func RemoveDuplicates[T comparable](s []T) []T {
	if len(s) == 0 {
		return s
	}
	seen := make(map[T]struct{}, len(s))
	result := s[:0]
	for _, v := range s {
		if _, exists := seen[v]; !exists {
			seen[v] = struct{}{}
			result = append(result, v)
		}
	}
	return result
}

func RemoveIf[T any](s []T, predicate func(T) bool) []T {
	result := s[:0]
	for _, v := range s {
		if !predicate(v) {
			result = append(result, v)
		}
	}
	return result
}

func RemoveOrderedWithNil[T any](s []*T, i int) []*T {
	if i < 0 || i >= len(s) {
		return s
	}
	s[i] = nil
	copy(s[i:], s[i+1:])

	s[len(s)-1] = nil
	return s[:len(s)-1]
}

func ShrinkCapacity[T any](s []T) []T {
	if cap(s) > 2*len(s) && len(s) > 0 {
		newS := make([]T, len(s))
		copy(newS, s)
		return newS
	}
	return s
}

func main() {
	fmt.Println("SLICE TASK 3: Удаление элементов")

	fmt.Println("1.RemoveUnordered:")
	s1 := []int{10, 20, 30, 40}
	fmt.Printf("   исходный: %v\n", s1)
	s1 = RemoveUnordered(s1, 1) // удаляем 20
	fmt.Printf("   после удаления индекса 1: %v (порядок изменён!)\n\n", s1)

	fmt.Println("2.RemoveOrdered:")
	s2 := []string{"a", "b", "c", "d"}
	fmt.Printf("   исходный: %v\n", s2)
	s2 = RemoveOrdered(s2, 1) // удаляем "b"
	fmt.Printf("   после удаления индекса 1: %v (порядок сохранён)\n\n", s2)

	fmt.Println("3.RemoveAllByValue:")
	s3 := []int{1, 2, 3, 2, 4, 2}
	fmt.Printf("   исходный: %v\n", s3)
	s3 = RemoveAllByValue(s3, 2)
	fmt.Printf("   после удаления всех 2: %v\n\n", s3)

	fmt.Println("4.RemoveDuplicates:")
	s4 := []string{"go", "rust", "go", "python", "rust"}
	fmt.Printf("   исходный: %v\n", s4)
	s4 = RemoveDuplicates(s4)
	fmt.Printf("   уникальные: %v\n\n", s4)

	fmt.Println("5.RemoveIf (удалить чётные):")
	s5 := []int{1, 2, 3, 4, 5, 6}
	fmt.Printf("   исходный: %v\n", s5)
	s5 = RemoveIf(s5, func(n int) bool { return n%2 == 0 })
	fmt.Printf("   только нечётные: %v\n\n", s5)

	fmt.Println("6.RemoveOrderedWithNil (указатели):")
	type Data struct{ Value int }
	ptrs := []*Data{{1}, {2}, {3}, {4}}
	fmt.Printf("   исходный: [%v, %v, %v, %v]\n",
		ptrs[0].Value, ptrs[1].Value, ptrs[2].Value, ptrs[3].Value)
	ptrs = RemoveOrderedWithNil(ptrs, 1) // удаляем {2}
	fmt.Printf("   после удаления: len=%d, cap=%d\n", len(ptrs), cap(ptrs))
	fmt.Printf("   проверяем что ptrs[1] != nil: %v\n\n", ptrs[1] != nil)

	fmt.Println("7.ShrinkCapacity:")
	s7 := make([]int, 0, 100)
	for i := 0; i < 10; i++ {
		s7 = append(s7, i)
	}
	fmt.Printf("   до: len=%d, cap=%d\n", len(s7), cap(s7))
	s7 = ShrinkCapacity(s7)
	fmt.Printf("   после: len=%d, cap=%d (экономия памяти!)\n\n", len(s7), cap(s7))

	fmt.Println("8.Цепочка операций:")
	data := []int{1, 2, 2, 3, 4, 4, 5}
	data = RemoveDuplicates(data)
	data = RemoveIf(data, func(n int) bool { return n > 3 })
	data = RemoveOrdered(data, 0)
	fmt.Printf("   результат: %v\n", data)
}
