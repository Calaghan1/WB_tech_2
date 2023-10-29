// Утилита cut

// Реализовать утилиту аналог консольной команды cut (man cut). Утилита должна принимать строки через STDIN, разбивать по разделителю (TAB) на колонки и выводить запрошенные.

// Реализовать поддержку утилитой следующих ключей:

// -f - "fields" - выбрать поля (колонки)

// -d - "delimiter" - использовать другой разделитель

// -s - "separated" - только строки с разделителем'
package main

import (
	"bufio"
	// "flag"
	"fmt"
	"os"
	// "strings"
)

func main() {
	// fields := flag.Int("f", 0, "Choose field")
	// delimiter := flag.String("d", " ", "Choose field")
	// separated := flag.Bool("d", false, "Choose field")
 scaner := bufio.NewScanner(os.Stdin)
	for scaner.Scan() {
		text := scaner.Text()

		if text == "" {
			break
		}
		// strings.Split(text, " ")
		fmt.Println(text)
	}
}
