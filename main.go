package main

import (
	"fmt"

	"./inputsql"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	var vvod int
	fmt.Print("Введите цифру требуемого действия:\n\t1-обновление eNB;\n\t2-обновление секторов;\n\t3-вывод информации о eNB;\n_:  ")
	fmt.Scan(&vvod)
	switch {
	case vvod == 1:
		inputsql.InputENB()
	case vvod == 2:
		inputsql.InputSEC()
	case vvod == 3:
		fmt.Print("Введите номер eNodeB: ")
		inputsql.Search()
	}
}
