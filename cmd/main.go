package main

import (
	"flag"
	"fmt"

	"github.com/artzhiv/sql_test/pkg/repository"
	_ "github.com/go-sql-driver/mysql"
)

var (
	configPath string
)

func init() {
	flag.StringVar(&configPath, "config-path", "configs/apiserver.toml", "path to config file")
}

func main() {
	flag.Parse()

	var vvod int
	repository.ClearCMD()
	repository.InfoV()
	repository.Info1()
	fmt.Scan(&vvod)
	fmt.Println()
	repository.ClearCMD()
	switch {
	case vvod == 1:
		repository.InputENB()
	case vvod == 2:
		repository.InputSEC()
	case vvod == 3:
		repository.ClearCMD()
		repository.InfoSearchENB()

		var a string
		fmt.Scan(&a)
		switch {
		case a == "*":
			repository.SearchList()
			fmt.Scanln()
			fmt.Scanln()
		case a == "г":
			var nummmm string
			fmt.Print("Введите город: ")
			fmt.Scan(&nummmm)
			repository.SearchCity(nummmm)
			fmt.Scanln()
			fmt.Scanln()
		case a == "о":
			var nummmm string
			fmt.Print("Введите область: ")
			fmt.Scan(&nummmm)
			repository.SearchRegion(nummmm)
			fmt.Scanln()
			fmt.Scanln()
		case a == "д":
			repository.SearchDel()
			fmt.Scanln()
			fmt.Scanln()
		default:
			repository.Search(a)
			fmt.Scanln()
			fmt.Scanln()
		}
		fmt.Scanln()
		repository.ClearCMD()

	case vvod == 4:
		var inp string
		fmt.Print("Введите значение из запроса: ")
		fmt.Scan(&inp)
		fmt.Println()
		repository.Convert(inp)
		fmt.Scanln()
		fmt.Scanln()
	case vvod == 5:
		fmt.Print("1-ввод БС ручками;\n2-ввод через файл TXT;\n_: ")
		var a int
		fmt.Scan(&a)
		switch {
		case a == 1:
			repository.SearchMTS()
		case a == 2:
			repository.TextSearchMTS()
		}
	case vvod == 666:
		// handler.OpenTestConnect()
		// repository.dirByName()
	case vvod == 0:
		break
	}
}
