package main

import (
	"fmt"

	"github.com/ArtZhiv/sql_test/cmd"
	"github.com/ArtZhiv/sql_test/inputsql"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	cmd.ClearCMD()

	var vvod int
	fmt.Print("Введите цифру требуемого действия:\n\t1-обновление eNB;\n\t2-обновление секторов;\n\t3-вывод информации о eNB;\n\t4-Перевод для UNET (из HEX в DEC);\n\t5-РРЛ для СООО <<МТС>>\n_:  ")
	fmt.Scan(&vvod)
	fmt.Println()
	switch {
	case vvod == 1:
		inputsql.InputENB()
	case vvod == 2:
		inputsql.InputSEC()
	case vvod == 3:
		cmd.ClearCMD()
		fmt.Print("Введите номер eNodeB: ")
		inputsql.Search()
	case vvod == 4:
		var inp string
		fmt.Print("Введите значение из запроса: ")
		fmt.Scan(&inp)
		fmt.Println()
		inputsql.Convert(inp)
	case vvod == 5:
		fmt.Print("Введите список eNodeB: ")
		inputsql.SearchMTS()
	}
}
