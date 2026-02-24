package main

import "fmt"

type account struct {
	value int
}

func main() {
	s1 := make([]account, 0, 2) // len 0, cap 2: [_, _]
	s1 = append(s1, account{})  // len 1, cap 2: [{0}, _]
	s2 := append(s1, account{}) //len 2, cap 2: [0,1]
	acc := &s2[0]               //s1 и s2 делят один и тот же массив
	acc.value = 100
	fmt.Println(s1, s2)        // [{100}] [{100} {0}]
	s1 = append(s1, account{}) // s1 len 2, cap 2; s2 len 2, cap 2: [0,1]
	acc.value += 100           // [{200}, {0}] и у s1 и у s2
	fmt.Println(s1, s2)        // [{200}, {0}] [{200}, {0}]
}
