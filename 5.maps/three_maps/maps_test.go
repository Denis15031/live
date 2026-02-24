package main

import (
	"strconv"
	"strings"
	"testing"
)

// Тесты для FilterByValue

// Базовая фильтрация
func TestFilterByValue_Basic(t *testing.T) {
	m := map[int]string{
		1: "golang",
		2: "python",
		3: "ruby",
		4: "rust",
		5: "java",
		6: "delphi",
	}
	allowed := []string{"golang", "rust"}

	got := FilterByValue(m, allowed)

	if len(got) != 2 {
		t.Errorf("Expected 2 elements, got %d", len(got))
	}

	if got[1] != "golang" {
		t.Errorf("got[1] = %q; want %q", got[1], "golang")
	}
	if got[4] != "rust" {
		t.Errorf("got[4] = %q; want %q", got[4], "rust")
	}

	for _, key := range []int{2, 3, 5, 6} {
		if _, exists := got[key]; exists {
			t.Errorf("Key %d should be filtered out", key)
		}
	}
}

// Пустой список разрешённых значений
func TestFilterByValue_EmptyAllowed(t *testing.T) {
	m := map[int]string{1: "a", 2: "b", 3: "c"}
	got := FilterByValue(m, []string{})

	if len(got) != 0 {
		t.Errorf("Expected empty result, got %v", got)
	}
}

// Пустая входная мапа
func TestFilterByValue_EmptyInput(t *testing.T) {
	got := FilterByValue(map[int]string{}, []string{"a", "b"})

	if len(got) != 0 {
		t.Errorf("Expected empty result for empty input, got %v", got)
	}
}

// Нет совпадений
func TestFilterByValue_NoMatches(t *testing.T) {
	m := map[int]string{1: "go", 2: "rust"}
	got := FilterByValue(m, []string{"python", "java"})

	if len(got) != 0 {
		t.Errorf("Expected empty result (no matches), got %v", got)
	}
}

// Исходная мапа не изменяется
func TestFilterByValue_OriginalUnchanged(t *testing.T) {
	original := map[int]string{1: "a", 2: "b", 3: "c"}
	originalCopy := make(map[int]string, len(original))
	for k, v := range original {
		originalCopy[k] = v
	}

	_ = FilterByValue(original, []string{"a"})

	for k, v := range originalCopy {
		if original[k] != v {
			t.Errorf("Original map was modified at key %d", k)
		}
	}
}

// Дубликаты в allowedValues
func TestFilterByValue_AllowedWithDuplicates(t *testing.T) {
	m := map[int]string{1: "a", 2: "b", 3: "c"}
	allowed := []string{"a", "a", "b", "b"}

	got := FilterByValue(m, allowed)

	if len(got) != 2 {
		t.Errorf("Expected 2 elements, got %d", len(got))
	}
	if got[1] != "a" || got[2] != "b" {
		t.Errorf("Unexpected result: %v", got)
	}
}

// Тесты для InvertMap

// Базовое инвертирование
func TestInvertMap_Basic(t *testing.T) {
	m := map[string]int{
		"asker":  18,
		"alex":   19,
		"epifan": 54,
	}

	got, err := InvertMap(m)
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}

	if got[18] != "asker" {
		t.Errorf("got[18] = %q; want %q", got[18], "asker")
	}
	if got[19] != "alex" {
		t.Errorf("got[19] = %q; want %q", got[19], "alex")
	}
	if got[54] != "epifan" {
		t.Errorf("got[54] = %q; want %q", got[54], "epifan")
	}
}

// Пустая мапа
func TestInvertMap_Empty(t *testing.T) {
	got, err := InvertMap(map[string]int{})
	if err != nil {
		t.Errorf("Unexpected error for empty map: %v", err)
	}
	if len(got) != 0 {
		t.Errorf("Expected empty result, got %v", got)
	}
}

// Дубликаты значений -> ошибка
func TestInvertMap_DuplicateValues_Error(t *testing.T) {
	m := map[string]int{
		"asker":  18,
		"frosya": 18,
		"alex":   53,
	}

	got, err := InvertMap(m)

	if err == nil {
		t.Fatal("Expected error for duplicate values, got nil")
	}

	if got != nil {
		t.Error("Expected nil result on error, got non-nil map")
	}

	errMsg := err.Error()
	if !strings.Contains(errMsg, "18") {
		t.Errorf("Error message should contain value 18, got: %q", errMsg)
	}
	if !strings.Contains(errMsg, "asker") || !strings.Contains(errMsg, "frosya") {
		t.Errorf("Error message should contain both conflicting keys, got: %q", errMsg)
	}
}

// Исходная мапа не изменяется
func TestInvertMap_OriginalUnchanged(t *testing.T) {
	original := map[string]int{"a": 1, "b": 2, "c": 3}
	originalCopy := make(map[string]int, len(original))
	for k, v := range original {
		originalCopy[k] = v
	}

	_, _ = InvertMap(original)

	for k, v := range originalCopy {
		if original[k] != v {
			t.Errorf("Original map was modified at key %q", k)
		}
	}
}

// Один элемент
func TestInvertMap_SingleElement(t *testing.T) {
	m := map[string]int{"solo": 42}

	got, err := InvertMap(m)
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}

	if got[42] != "solo" {
		t.Errorf("Expected got[42] = %q, got %q", "solo", got[42])
	}
}

// Бенчмарки

// Производительность фильтрации
func BenchmarkFilterByValue(b *testing.B) {
	m := make(map[int]string, 1000)
	for i := 0; i < 1000; i++ {
		m[i] = string(rune('a' + i%10))
	}
	allowed := []string{"a", "b", "c", "d", "e"}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = FilterByValue(m, allowed)
	}
}

// Производительность инвертирования (уникальные)
func BenchmarkInvertMap(b *testing.B) {
	m := make(map[string]int, 1000)
	for i := 0; i < 1000; i++ {
		m[strconv.Itoa(i)] = i
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = InvertMap(m)
	}
}

// Производительность с конфликтами
func BenchmarkInvertMap_WithConflict(b *testing.B) {
	m := make(map[string]int, 1000)
	for i := 0; i < 1000; i++ {
		m[strconv.Itoa(i)] = i % 100
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = InvertMap(m)
	}
}
