// Утилита cut

// Реализовать утилиту аналог консольной команды cut (man cut). Утилита должна принимать строки через STDIN, разбивать по разделителю (TAB) на колонки и выводить запрошенные.

// Реализовать поддержку утилитой следующих ключей:

// -f - "fields" - выбрать поля (колонки)

// -d - "delimiter" - использовать другой разделитель

// -s - "separated" - только строки с разделителем'
package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func CheckErorr(err error) {
	if err != nil {
		log.Fatalln(err.Error())
	}
}

func CheckFlags(fields int, delimiter string, separated bool) bool {
	if fields == 0 {
		return false
	}
	return true
}
func main() {
	fields := flag.String("f", "0", "Choose field")
	delimiter := flag.String("d", "\t", "Choose field")
	separated := flag.Bool("s", false, "Choose field")
	flag.Parse()
	args := flag.Args()
	file, err := os.Open(args[0])
	CheckErorr(err)
	if *fields == "0" {
		fmt.Println("Error: Wrong flags")
		os.Exit(0)
	}
	fields_tmp := strings.Split(*fields, ",")
	fields_int := make([]int, 0, len(fields_tmp))

	for i := 0; i < len(fields_tmp); i++ {
		num, err := strconv.Atoi(fields_tmp[i])
		if num == 0 {
			fmt.Println("cut: fields are numbered from 1")
			os.Exit(0)
		}
		CheckErorr(err)
		fields_int = append(fields_int, num - 1)
	}
	// fmt.Println(fields_int)
	// fmt.Println(file.Name())

 scaner := bufio.NewScanner(file)
	for scaner.Scan() {
		line := scaner.Text()
		line_arr := strings.Split(line, *delimiter) //Если не найдет разделитель в строке вернет всю строку в 0 элемент слайса

		if line == "" || (line_arr[0] == line && *separated) {
			continue
		}
		if line_arr[0] == line {
			fmt.Println(line)
		} else {
			for i := 0; i < len(fields_int); i++ {

				if fields_int[i] < len(line_arr) {
					fmt.Print(line_arr[fields_int[i]])
				}
				fmt.Print("\n")
			}
		}
	}
}
