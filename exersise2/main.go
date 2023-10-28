// Задача на распаковку
// Создать Go-функцию, осуществляющую примитивную распаковку строки, содержащую повторяющиеся символы/руны, например:

// "a4bc2d5e" => "aaaabccddddde" 
// "abcd" => "abcd"
// "45" => "" (некорректная строка)
// "" => ""

// Дополнительно 
// Реализовать поддержку escape-последовательностей.

// Например:

// qwe\4\5 => qwe45 (*)
// qwe\45 => qwe44444 (*)
// qwe\\5 => qwe\\\\\ (*)
// В случае если была передана некорректная строка, функция должна возвращать ошибку. Написать unit-тесты.

package main

import (
	"fmt"
	"strings"
	"strconv"
)

func IsDigit(c rune) bool {
	if c >= '0' && c <= '9' {
		return true
	} else {
		return false 
	}
}

func UnpackStr(str string) (string, error) {
	fmt.Println("Working with:", str)
	builder_1 := strings.Builder{}
	builder_2 := strings.Builder{}
	number := ""
	last_simb := ""
	for i := 0; i < len(str); i++ {
		if IsDigit(rune(str[i])) {
			if i == 0 {
				return "", fmt.Errorf("Wrong format")
			}
		for i < len(str)  && IsDigit(rune(str[i])) {
			number += string(str[i])
			i++
		}
		i--
		if number != "" {
			num, err := strconv.Atoi(number)
			if err != nil {
				return "", err
			}
			for j := 0; j < num - 1; j ++ {
				builder_1.Write([]byte(last_simb))
			}
			builder_2.Write([]byte(builder_1.String()))
			builder_1.Reset()
		}
	} else {
		if rune(str[i]) == '\\' {
			if i + 1 < len(str) {
				i++
			} 
		}
		last_simb = string(str[i])
		number = ""
		builder_2.Write([]byte(last_simb))

	}
}
	return builder_2.String(), nil
}
func main() {
	str, err := UnpackStr("qwe\\\\5") 
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(str)
	}
	
}