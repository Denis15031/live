package main

import (
	"fmt"
	"sort"
	"strings"
)

func FilterByValue(m map[int]string, allowedValues []string) map[int]string {
	allowedSet := make(map[string]struct{}, len(allowedValues))
	for _, val := range allowedValues {
		allowedSet[val] = struct{}{}
	}

	result := make(map[int]string, len(m))
	for key, value := range m {
		if _, exists := allowedSet[value]; exists {
			result[key] = value
		}
	}

	return result
}

func InvertMap(m map[string]int) (map[int]string, error) {
	seen := make(map[int]string, len(m))

	for key, value := range m {
		if originalKey, exists := seen[value]; exists {
			return nil, fmt.Errorf(
				"conflict: value %d is associated with both keys %q and %q",
				value, originalKey, key,
			)
		}
		seen[value] = key
	}

	inverted := make(map[int]string, len(m))
	for key, value := range m {
		inverted[value] = key
	}

	return inverted, nil
}

func PrintMapIntString(m map[int]string, title string) {
	if len(m) == 0 {
		fmt.Printf("%s: пустая\n", title)
		return
	}

	fmt.Printf("%s (%d элементов):\n", title, len(m))
	fmt.Println(strings.Repeat("─", 40))

	keys := make([]int, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}
	sort.Ints(keys)

	for _, k := range keys {
		fmt.Printf("  [%d] => %q\n", k, m[k])
	}
	fmt.Println(strings.Repeat("─", 40))
}

func main() {
	fmt.Println("Демонстрация продвинутых операций с Мапой")

	original := map[int]string{
		1: "golang",
		2: "python",
		3: "ruby",
		4: "rust",
		5: "java",
		6: "delphi",
	}

	PrintMapIntString(original, "Исходная Мапа")

	fmt.Println("Тест: FilterByValue")
	allowed := []string{"golang", "rust"}
	filtered := FilterByValue(original, allowed)
	PrintMapIntString(filtered, "После фильтрации (только golang, rust)")

	fmt.Println("Тест: InvertMap (уникальные значения)")
	unique := map[string]int{
		"asker":  18,
		"alex":   19,
		"epifan": 54,
	}

	fmt.Println("Исходная Мапа (имя -> возраст):")
	for name, age := range unique {
		fmt.Printf("  %q => %d\n", name, age)
	}

	inverted, err := InvertMap(unique)
	if err != nil {
		fmt.Printf("Ошибка: %v\n", err)
	} else {
		fmt.Println("Инвертированная map (возраст → имя):")
		ages := make([]int, 0, len(inverted))
		for age := range inverted {
			ages = append(ages, age)
		}
		sort.Ints(ages)
		for _, age := range ages {
			fmt.Printf("  %d => %q\n", age, inverted[age])
		}
	}

	fmt.Println("Тест: InvertMap (дубликаты значений)")
	withDuplicates := map[string]int{
		"asker":  18,
		"frosya": 18,
		"alex":   53,
	}
	_, err = InvertMap(withDuplicates)
	if err != nil {
		fmt.Printf("Ожидаемая ошибка: %v\n", err)
	}
}
