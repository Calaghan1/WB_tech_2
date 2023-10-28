// Поиск анаграмм по словарю

// Написать функцию поиска всех множеств анаграмм по словарю.

// Например:

// "пятак", "пятка" и "тяпка" - принадлежат одному множеству,

// "листок", "слиток" и "столик" - другому.

// Требования:

// Входные данные для функции: ссылка на массив, каждый элемент которого - слово на русском языке в кодировке utf8
// Выходные данные: ссылка на мапу множеств анаграмм
// Ключ - первое встретившееся в словаре слово из множества. Значение - ссылка на массив, каждый элемент которого,
// слово из множества.

// Массив должен быть отсортирован по возрастанию.
// Множества из одного элемента не должны попасть в результат.
// Все слова должны быть приведены к нижнему регистру.
// В результате каждое слово должно встречаться только один раз.

package main

import (
	"fmt"
	"sort"
)

func MakeCast(s string) []rune{
	buff := []rune(s)
	sort.SliceStable(buff, func(i, j int) bool {return buff[i] < buff[j]})
	return buff
}

func Isequal(s1, s2 []rune) bool{
	if len(s1) == len(s2) {
		for i := 0; i < len(s1); i++ {
			if s1[i] != s2[i] {
				return false
			}
		}
	} else {
		return false
	}
	return true
}

func Isanagram(s1, s2 string) bool {
	s1_b := MakeCast(s1)
	s2_b := MakeCast(s2)
	return Isequal(s1_b, s2_b)
}



func FindAn(data []string) map[string][]string {
	result := make(map[string][]string)
	flag := true
	for i := 0; i < len(data); i ++ {
		flag = true
		for key := range result {
			if key == data[i] {
				flag = false
				break
			}
			if Isanagram(key, data[i]) {
				index := sort.SearchStrings(result[key], data[i]) 
				if index == len(result[key]) {
					result[key] = append(result[key], data[i])
					flag = false
					break
				}
			}
		}
		if flag {result[data[i]] = make([]string, 0, 3)}
	}
	for key := range result {
		value, _ := result[key] 
		if len(value) == 0 {
			delete(result, key)
		}
	}
	return result
}
func main() {
	var data = []string{"пятак", "пятка", "тяпка", "листок", "слиток", "столик", "сосна"}
	fmt.Println(Isanagram("пятак", "листок"))
	res := FindAn(data)
	fmt.Println(res)
}