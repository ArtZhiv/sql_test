package main

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/ArtZhiv/sql_test/inputsql"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	cls()

	var vvod int
	fmt.Print("Введите цифру требуемого действия:\n\t1-обновление eNB;\n\t2-обновление секторов;\n\t3-вывод информации о eNB;\n\t4-Перевод для UNET (из HEX в DEC)\n_:  ")
	fmt.Scan(&vvod)
	fmt.Println()
	switch {
	case vvod == 1:
		inputsql.InputENB()
	case vvod == 2:
		inputsql.InputSEC()
	case vvod == 3:
		fmt.Print("Введите номер eNodeB: ")
		inputsql.Search()
	case vvod == 4:
		var inp string
		fmt.Print("Введите значение из запроса: ")
		fmt.Scan(&inp)
		fmt.Println()
		inputsql.Convert(inp)
	}
}

func cls() {
	cmd := exec.Command("powershell", "/c", "cls")
	cmd.Stdout = os.Stdout
	cmd.Run()
}
