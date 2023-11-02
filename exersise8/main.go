// Взаимодействие с ОС

// Необходимо реализовать свой собственный UNIX-шелл-утилиту с поддержкой ряда простейших команд:

// - cd <args> - смена директории (в качестве аргумента могут быть то-то и то)

// - pwd - показать путь до текущего каталога

// - echo <args> - вывод аргумента в STDOUT

// - kill <args> - "убить" процесс, переданный в качесте аргумента (пример: такой-то пример)

// - ps - выводит общую информацию по запущенным процессам в формате *такой-то формат*

// Так же требуется поддерживать функционал fork/exec-команд

// Дополнительно необходимо поддерживать конвейер на пайпах (linux pipes, пример cmd1 | cmd2 | .... | cmdN).

// *Шелл — это обычная консольная программа, которая будучи запущенной, в интерактивном сеансе выводит некое приглашение

// в STDOUT и ожидает ввода пользователя через STDIN. Дождавшись ввода, обрабатывает команду согласно своей логике

// и при необходимости выводит результат на экран. Интерактивный сеанс поддерживается до тех пор, пока не будет введена команда выхода (например \quit).
package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)
func CheckErorr(err error) bool {
	if err != nil {
		log.Println(err.Error())
		return false
	} else {
		return true
	}
}
func ProcessCd(path string) {
	err := os.Chdir(path)
	CheckErorr(err)
}
func main() {
	fmt.Println("Shel is ready ot porcces comands")
	scaner := bufio.NewScanner(os.Stdin)
	for scaner.Scan() {
		text := scaner.Text()

		if text == "" {
			break
		}
		buff := strings.Split(text, " ")
		switch buff[0] {
			case "cd":
				ProcessCd(buff[1])
			case "pwd":
				res, err := os.Getwd()
				if CheckErorr(err) {
					fmt.Println(res)
				}
			case "echo":
				fmt.Println(strings.Join(buff[1:], " "))
			case "ps":
				os.Getppid()
		}

	}
}