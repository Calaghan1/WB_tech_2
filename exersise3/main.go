// Утилита sort
// Отсортировать строки в файле по аналогии с консольной утилитой sort (man sort — смотрим описание и основные параметры): на входе подается файл из несортированными строками, на выходе — файл с отсортированными.

// Реализовать поддержку утилитой следующих ключей:

// -k — указание колонки для сортировки (слова в строке могут выступать в качестве колонок, по умолчанию разделитель — пробел)

// -n — сортировать по числовому значению

// -r — сортировать в обратном порядке

// -u — не выводить повторяющиеся строки

// Дополнительно

// Реализовать поддержку утилитой следующих ключей:

// -M — сортировать по названию месяца

// -b — игнорировать хвостовые пробелы

// -c — проверять отсортированы ли данные

// -h — сортировать по числовому значению с учетом суффиксов

package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)
func CheckErorr(er error) {
	if er != nil {
		log.Fatalf(er.Error())
	}
}
func main() {

	
	column := flag.Int("k", -1, "Колонка для сортировки (по умолчанию -1 для всей строки)")
	numeric := flag.Bool("n", false, "Сортировать числа")
	// reverse := flag.Bool("r", false, "Сортировать в обратном порядке")
	// unique := flag.Bool("u", false, "Не выводить повторяющиеся строки")
	
	file_data := make([]string, 0, 1000)
	flag.Parse()
	file_names := flag.Args()
	fmt.Println("COLUMN", *column)
	fmt.Println("FILENAMES", file_names)
	file, err := os.Open(file_names[0])
	CheckErorr(err)
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		file_data = append(file_data, (scanner.Text()))
	}
	if *column != -1 {
		if *numeric {
			sort.SliceStable(file_data, func(i, j int) bool {
				num1, err := strconv.Atoi(strings.Split(file_data[i], " ")[*column])
				CheckErorr(err)
				num2, err := strconv.Atoi(strings.Split(file_data[j], " ")[*column])
				CheckErorr(err)
				return num1 < num2
		})
	} else {
		sort.SliceStable(file_data, func(i, j int) bool {return  strings.Split(file_data[i], " ")[*column] < strings.Split(file_data[j], " ")[*column]})
	}			
	} else {
		sort.SliceStable(file_data, func(i, j int) bool {return file_data[i] < file_data[j]})
	}
	for _, d := range file_data {
		fmt.Println(string(d))
	}
}