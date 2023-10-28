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
	"testing"
)

func TestUnpac(t *testing.T) {
	str, err := UnpackStr("a4bc2d5e")
	if err != nil {
		t.Errorf("Wrong answer %s expect aaaabccddddde", str)
	}
	if str != "aaaabccddddde" {
		t.Errorf("Wrong answer %s expect aaaabccddddde", str)
	}
	str, err = UnpackStr("abcd") 
	if err != nil {
		t.Errorf("Wrong answer %s expect aaaabccddddde", str)
	}
	if str != "abcd" {
		t.Errorf("Wrong answer %s expect aaaabccddddde", str)
	}
	str, _ = UnpackStr("45") 
	if str != "" {
		t.Errorf("Wrong answer %s expect aaaabccddddde", str)
	}
	str, _ = UnpackStr("") 
	if str != "" {
		t.Errorf("Wrong answer %s expect aaaabccddddde", str)
	}
}
// qwe\4\5 => qwe45 (*)
// qwe\45 => qwe44444 (*)
// qwe\\5 => qwe\\\\\ (*)
func TestEscape(t *testing.T) {
	str, err := UnpackStr("qwe\\4\\5")
	if err != nil {
		t.Errorf("Wrong answer %s expect aaaabccddddde", str)
	}
	if str != "qwe45" {
		t.Errorf("Wrong answer %s expect aaaabccddddde", str)
	}
	str, err = UnpackStr("qwe\\45")
	if err != nil {
		t.Errorf("Wrong answer %s expect aaaabccddddde", str)
	}
	if str != "qwe44444" {
		t.Errorf("Wrong answer %s expect aaaabccddddde", str)
	}
	str, err = UnpackStr("qwe\\\\5")
	if err != nil {
		t.Errorf("Wrong answer %s expect aaaabccddddde", str)
	}
	if str != "qwe\\\\\\\\\\" {
		t.Errorf("Wrong answer %s expect aaaabccddddde", str)
	}
}