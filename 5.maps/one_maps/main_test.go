package main

import (
	"fmt"
	"sync"
	"testing"
)

func resetPeople() {
	peopleMu.Lock()
	defer peopleMu.Unlock()
	people = make(map[string]int, 10)
}

func TestAddPerson(t *testing.T) {
	resetPeople()
	AddPerson("Test", 25)

	peopleMu.RLock()
	age, exists := people["Test"]
	peopleMu.RUnlock()

	if !exists || age != 25 {
		t.Errorf("AddPerson failed: exists=%v, age=%d", exists, age)
	}
}

func TestGetAge_Existing(t *testing.T) {
	resetPeople()
	AddPerson("Alice", 30)
	if got := GetAge("Alice"); got != 30 {
		t.Errorf("GetAge(Alice) = %d, want 30", got)
	}
}

func TestGetAge_NotFound(t *testing.T) {
	resetPeople()
	if got := GetAge("Unknown"); got != 0 {
		t.Errorf("GetAge(Unknown) = %d, want 0", got)
	}
}

func TestDeletePerson(t *testing.T) {
	resetPeople()
	AddPerson("ToDelete", 40)
	DeletePerson("ToDelete")

	peopleMu.RLock()
	_, exists := people["ToDelete"]
	peopleMu.RUnlock()

	if exists {
		t.Error("DeletePerson failed: item still exists")
	}
}

func TestConcurrentAccess(t *testing.T) {
	resetPeople()
	var wg sync.WaitGroup

	for i := 0; i < 100; i++ {
		wg.Add(3)
		go func(n int) {
			defer wg.Done()
			AddPerson(fmt.Sprintf("U%d", n), n)
		}(i)
		go func(n int) {
			defer wg.Done()
			GetAge(fmt.Sprintf("U%d", n%10))
		}(i)
		go func(n int) {
			defer wg.Done()
			DeletePerson(fmt.Sprintf("U%d", n/2))
		}(i)
	}
	wg.Wait()
}
