package main

import (
	"fmt"

	"github.com/ArtZhiv/sql_test/cmd"
	"github.com/ArtZhiv/sql_test/inputsql"
	"github.com/ArtZhiv/sql_test/test"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	var vvod int
	cmd.Info1()
	fmt.Scan(&vvod)
	fmt.Println()
	switch {
	case vvod == 1:
		inputsql.InputENB()
	case vvod == 2:
		inputsql.InputSEC()
	case vvod == 3:
		cmd.ClearCMD()
		cmd.InfoSearchENB()

		var a string
		fmt.Scan(&a)
		switch {
		case a == "*":
			inputsql.SearchList()
		case a == "г":
			var nummmm string
			fmt.Print("Введите город: ")
			fmt.Scan(&nummmm)
			inputsql.SearchCity(nummmm)
		case a == "о":
			var nummmm string
			fmt.Print("Введите область: ")
			fmt.Scan(&nummmm)
			inputsql.SearchRegion(nummmm)
		case a == "д":
			inputsql.SearchDel()
		default:
			inputsql.Search(a)
		}

	case vvod == 4:
		var inp string
		fmt.Print("Введите значение из запроса: ")
		fmt.Scan(&inp)
		fmt.Println()
		inputsql.Convert(inp)
	case vvod == 5:
		inputsql.SearchMTS()
	case vvod == 666:
		fmt.Println("Здесь запускаются тестовые функции")
		test.Test()
	}
}
