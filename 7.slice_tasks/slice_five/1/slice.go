package main

type stack struct {
	data []int
}

// Создает и возвращает новый экземпляр стека
func New() *stack {
	return &stack{
		data: make([]int, 0),
	}
}

// Добавляет элемент на вершину стека
func (s *stack) Push(v int) {
	s.data = append(s.data, v)
}

// Удаляет и возвращает элемент с вершины стека
func (s *stack) Pop() int {
	if len(s.data) == 0 {
		panic("stack is empty")
	}
	v := s.data[len(s.data)-1]
	s.data = s.data[:len(s.data)-1]
	return v
}

var _Stacker = (*stack)(nil)

type Stacker interface {
	Push(v int)
	Pop() int
}
