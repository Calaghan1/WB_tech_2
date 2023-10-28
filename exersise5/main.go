// Реализовать утилиту фильтрации по аналогии с консольной утилитой (man grep — смотрим описание и основные параметры).

// Реализовать поддержку утилитой следующих ключей:

// -A - "after" печатать +N строк после совпадения

// -B - "before" печатать +N строк до совпадения

// -C - "context" (A+B) печатать ±N строк вокруг совпадения

// -c - "count" (количество строк)

// -i - "ignore-case" (игнорировать регистр)

// -v - "invert" (вместо совпадения, исключать)

// -F - "fixed", точное совпадение со строкой, не паттерн

// -n - "line num", напечатать номер строки

package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
)

func CheckErorr(err error) {
	if err != nil {
		log.Fatalln(err.Error())
	}
}
func FindMathes(pattern, line string, ignoreCaseFlag, invertFlag, fixedFlag bool) bool {
	if fixedFlag {
		return invertFlag != (pattern == line)
	}
	if ignoreCaseFlag {
		return invertFlag != strings.Contains(strings.ToLower(line), strings.ToLower(pattern)) 
	}
	return invertFlag != strings.Contains(line, pattern) 
}

func Save_N_lines(line string, n int, buff []string) []string {
	if n == 0 {
		return buff
	}
	if len(buff) == n {
		for i := 0; i < n - 1; i++ {
			buff[i] = buff[i + 1]
		}
		buff[n - 1] = line
	} else {
		buff = append(buff, line)
	}
	return buff
}

func PintLinesBefore(buff []string) {
	for _, val := range buff {
		fmt.Println(val)
	}
}


func main() {
	afterFlag := flag.Int("A", 0, "Print N lines after match")
    beforeFlag := flag.Int("B", 0, "Print N lines before match")
    contextFlag := flag.Int("C", 0, "Print N lines before and after match")
    countFlag := flag.Bool("c", false, "Print only the count of matching lines")
    ignoreCaseFlag := flag.Bool("i", false, "Case-insensitive search")
    invertFlag := flag.Bool("v", false, "Invert the match")
    fixedFlag := flag.Bool("F", false, "Exact string match, not a pattern")
    lineNumFlag := flag.Bool("n", false, "Print line numbers")
	counter := 0
	if *contextFlag > 0 {
		*afterFlag = *contextFlag
		*beforeFlag = *contextFlag
	}
	Bbuffer := make([]string, 0, *beforeFlag)
    flag.Parse()
    args := flag.Args()
	if len(args) < 1 {
        fmt.Println("Usage: grep [OPTIONS] PATTERN [FILE...]")
        os.Exit(1)
    }
	pattern := args[0]
    files := args[1:]
	fmt.Println(files)
	fmt.Println(pattern)
	if len(files) == 0 {
		log.Fatalln("No files")
	}
	ForcePrintCounter := 0
	for _, name := range files {
		Ncounter := 0
		file, err := os.Open(name)
		CheckErorr(err)
		scanner := bufio.NewScanner(file)
		for scanner.Scan() {
			Ncounter++
			line := scanner.Text()
			Save_N_lines(line, *beforeFlag, Bbuffer)
			if FindMathes(pattern, line, *ignoreCaseFlag, *invertFlag, *fixedFlag) {
				counter++
				ForcePrintCounter = *afterFlag
				if !*countFlag {
					if len(files) > 1 {
						PintLinesBefore(Bbuffer)
						if *lineNumFlag {
							fmt.Printf("%d%s:%s\n",Ncounter, name, line)
						} else {
							fmt.Printf("%s:%s\n", name, line)
						}
					} else {
						PintLinesBefore(Bbuffer)
						
						fmt.Println(line)
					}
				}
			} else if (ForcePrintCounter > 0) {
				ForcePrintCounter--
				fmt.Println(line)
			}
		}
		if *countFlag {
			if len(files) > 1{
				fmt.Printf("%s:%d\n", name, counter)
			} else {
				fmt.Println(counter)
			}
		}
	}

	
}

