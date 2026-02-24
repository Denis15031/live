package main

import (
	"fmt"
	"sort"
	"strings"
	"unicode"
)

func WordFrequency(text string) map[string]int {

	words := strings.Fields(text)
	freq := make(map[string]int, len(words)/2+1)

	for _, word := range words {
		cleaned := cleanWord(word)

		if cleaned != "" {
			freq[cleaned]++
		}
	}
	return freq
}

func cleanWord(word string) string {
	cleaned := strings.TrimFunc(word, func(r rune) bool {
		return !unicode.IsLetter(r) && !unicode.IsNumber(r)
	})
	return strings.ToLower(cleaned)
}

type WordCount struct {
	Word  string
	Count int
}

func PrintWordFrequency(freqMap map[string]int) {

	if len(freqMap) == 0 {
		fmt.Println("Нет данных для анализа")
		return
	}
	words := make([]WordCount, 0, len(freqMap))
	for word, count := range freqMap {
		words = append(words, WordCount{Word: word, Count: count})
	}
	sort.Slice(words, func(i, j int) bool {
		if words[i].Count != words[j].Count {
			return words[i].Count > words[j].Count
		}
		return words[i].Word < words[j].Word
	})

	fmt.Println("Частотный анализ слов:")
	fmt.Println(strings.Repeat("-", 30))
	fmt.Printf("%-15s | %s", "Cлово", "Частота")
	fmt.Println(strings.Repeat("-", 30))

	for _, wc := range words {
		fmt.Printf("%-15s | %d", wc.Word, wc.Count)
	}
	fmt.Println(strings.Repeat("-", 30))

}

func main() {

	text := "golang is great and golang is fast"

	fmt.Println("Исходный текст:")
	fmt.Printf("%s", text)

	freqMap := WordFrequency(text)

	PrintWordFrequency(freqMap)

	fmt.Println("\n" + strings.Repeat("=", 40))

	examples := []struct {
		desc string
		text string
	}{
		{"Разный регистр", "Go GO go Rust rust"},
		{"С пунктуацией", "Hello, world! Hello Go."},
		{"Пустая строка", ""},
		{"Только знаки", "!!! ?? ..."},
	}
	for _, ex := range examples {
		fmt.Printf("%s: %s", ex.desc, ex.text)
		PrintWordFrequency(WordFrequency(ex.text))
	}
}
