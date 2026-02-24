package main

import (
	"strings"
	"testing"
)

// Проверяет базовый подсчёт частоты
func TestWordFrequency_Basic(t *testing.T) {
	text := "go go rust rust rust"
	got := WordFrequency(text)

	want := map[string]int{
		"go":   2,
		"rust": 3,
	}

	// Проверяем, что все ожидаемые слова есть с правильной частотой
	for word, expectedCount := range want {
		if got[word] != expectedCount {
			t.Errorf("WordFrequency(%q)[%q] = %d; want %d",
				text, word, got[word], expectedCount)
		}
	}

	// Проверяем, что нет лишних слов
	if len(got) != len(want) {
		t.Errorf("Expected %d unique words, got %d", len(want), len(got))
	}
}

// Проверяет регистронезависимость
func TestWordFrequency_CaseInsensitive(t *testing.T) {
	tests := []struct {
		input    string
		expected map[string]int
	}{
		{
			input:    "Go GO go",
			expected: map[string]int{"go": 3},
		},
		{
			input:    "HELLO Hello hello",
			expected: map[string]int{"hello": 3},
		},
	}

	for _, tt := range tests {
		got := WordFrequency(tt.input)
		for word, expCount := range tt.expected {
			if got[word] != expCount {
				t.Errorf("Input %q: word %q count = %d; want %d",
					tt.input, word, got[word], expCount)
			}
		}
	}
}

// Проверяет очистку знаков препинания
func TestWordFrequency_WithPunctuation(t *testing.T) {
	text := "Hello, world! Hello. Go? GO!"
	got := WordFrequency(text)

	expected := map[string]int{
		"hello": 2,
		"world": 1,
		"go":    2,
	}

	for word, expCount := range expected {
		if got[word] != expCount {
			t.Errorf("Word %q: got %d, want %d", word, got[word], expCount)
		}
	}
}

// Проверяет обработку пустой строки
func TestWordFrequency_Empty(t *testing.T) {
	got := WordFrequency("")
	if len(got) != 0 {
		t.Errorf("Expected empty map for empty input, got %v", got)
	}
}

// Проверяет ввод только со знаками
func TestWordFrequency_OnlyPunctuation(t *testing.T) {
	inputs := []string{"!!!", "??? ...", ",.,.!", " - _ "}
	for _, input := range inputs {
		got := WordFrequency(input)
		if len(got) != 0 {
			t.Errorf("Input %q: expected empty map, got %v", input, got)
		}
	}
}

// Проверяет, что цифры сохраняются в словах
func TestWordFrequency_Numbers(t *testing.T) {
	text := "go123 go123 test42"
	got := WordFrequency(text)

	if got["go123"] != 2 {
		t.Errorf("Expected 'go123' count = 2, got %d", got["go123"])
	}
	if got["test42"] != 1 {
		t.Errorf("Expected 'test42' count = 1, got %d", got["test42"])
	}
}

// Проверяет, что вывод не паникует
func TestPrintWordFrequency_NoPanic(t *testing.T) {
	// Эти вызовы не должны вызывать panic
	defer func() {
		if r := recover(); r != nil {
			t.Errorf("PrintWordFrequency panicked: %v", r)
		}
	}()

	PrintWordFrequency(nil)
	PrintWordFrequency(map[string]int{})
	PrintWordFrequency(map[string]int{"test": 1})
	PrintWordFrequency(map[string]int{"a": 5, "b": 3, "c": 5})
}

// Проверяет вспомогательную функцию нормализации
func TestCleanWord(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{"Hello", "hello"},
		{"HELLO", "hello"},
		{"Hello,", "hello"},
		{"...test!", "test"},
		{"!!!", ""},
		{"go123", "go123"},
		{"", ""},
	}

	for _, tt := range tests {
		got := cleanWord(tt.input)
		if got != tt.expected {
			t.Errorf("cleanWord(%q) = %q; want %q",
				tt.input, got, tt.expected)
		}
	}
}

// Измеряет производительность функции
func BenchmarkWordFrequency(b *testing.B) {
	text := "golang is great and golang is fast " +
		"performance matters and golang performs well"

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = WordFrequency(text)
	}
}

// Измеряет производительность на большом тексте
func BenchmarkWordFrequency_Large(b *testing.B) {
	// Генерируем большой текст (10000 слов)
	var text strings.Builder
	words := []string{"go", "rust", "python", "java", "cpp", "is", "great", "fast"}
	for i := 0; i < 10000; i++ {
		text.WriteString(words[i%len(words)])
		text.WriteString(" ")
	}
	largeText := text.String()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = WordFrequency(largeText)
	}
}
