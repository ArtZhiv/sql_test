package main

import (
	"database/sql"
	"log"
	"os"
	"os/user"
	"strings"

	"github.com/ArtZhiv/sql_test/pkg/repository"
	"github.com/c-bata/go-prompt"
	"github.com/fatih/color"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	y := color.New(color.FgYellow).Add(color.Underline)
	if repository.Compare() == false {
		usr, err := user.Current()
		if err != nil {
			log.Fatal(err)
		}

		c := color.New(color.FgCyan).Add(color.Underline)
		repository.ClearCMD()
		c.Println("\nStarting the application ...")
		c.Printf("Hello %v!\n", usr.Name)
		repository.InfoV()
		defer c.Printf("Bye %v!\n", usr.Name)

		pt := prompt.New(
			Executor,
			completer,
			prompt.OptionTitle("ERP_CLI_beCloud "+"v."),
			prompt.OptionPrefixTextColor(15),
			prompt.OptionPrefix("INPUT >>> "),
			prompt.OptionInputTextColor(prompt.Cyan),
		)
		pt.Run()
	} else {
		repository.UpdateTable()
		repository.ClearCMD()
		y.Printf("Обновлена таблица eNodeB, добавлено %v записей\n",
			repository.UpdateBaseStation(initDB(repository.ConnDB)))
		y.Printf("Обновлена таблица с секторами, добавлено %v записей\n",
			repository.UpdateSector(initDB(repository.ConnDB)))

		if repository.Compare() == false {
			usr, err := user.Current()
			if err != nil {
				log.Fatal(err)
			}

			c := color.New(color.FgCyan).Add(color.Underline)
			c.Println("\nStarting the application ...")
			c.Printf("Hello %v!\n", usr.Name)
			repository.InfoV()
			defer c.Printf("Bye %v!\n", usr.Name)

			pt := prompt.New(
				Executor,
				completer,
				prompt.OptionTitle("ERP_CLI_beCloud "+"v."),
				prompt.OptionPrefixTextColor(15),
				prompt.OptionPrefix("INPUT >>> "),
				prompt.OptionInputTextColor(prompt.Cyan),
			)
			pt.Run()
		}
	}
}

func completer(d prompt.Document) []prompt.Suggest {
	var s []prompt.Suggest
	// {Text: "exit", Description: "Exit at the program"},
	switch d.Text {
	case "q", "quit":
		s = []prompt.Suggest{
			{Text: "exit", Description: "Exit at the program"},
		}
	case "s":
		s = []prompt.Suggest{
			{Text: "search ", Description: "Search eNodeB by number"},
		}
	case "w":
		s = []prompt.Suggest{
			{Text: "work ", Description: "Вывод списка станций без секторов"},
			{Text: "workFromMts ", Description: "Вывод списка станций по работам МТС."},
		}
	case "u":
		s = []prompt.Suggest{
			{Text: "unet ", Description: "Шаблон письма для Unet"},
		}
	default:
		s = repository.GenYaml()
	}
	return prompt.FilterHasPrefix(s,
		d.GetWordBeforeCursor(),
		true,
	)
}

// Executor ...
func Executor(s string) {
	s = strings.TrimSpace(s)
	setCommand := strings.Split(s, " ")
	switch setCommand[0] {
	case "q", "quit":
		repository.ClearCMD()
		usr, err := user.Current()
		if err != nil {
			log.Fatal(err)
		}
		c := color.New(color.FgCyan).Add(color.Underline)
		c.Printf("Bye %v!\n", usr.Name)
		os.Exit(0)
		return
	case "test", "t":
		repository.RipeRequest()
	case "unet":
		var value string = setCommand[1]
		repository.ConvertIMSIToGlobalID(value)
		return
	case "search":
		layer := setCommand[1:]
		for _, data := range layer {
			repository.Search(data)
		}
	case "open":
		repository.OpenFileMTS()
	case "workFromMts":
		var value string = setCommand[1]
		repository.TextSearchMTS(value)
	case "clear":
		repository.ClearCMD()

		usr, err := user.Current()
		if err != nil {
			log.Fatal(err)
		}

		c := color.New(color.FgCyan).Add(color.Underline)
		c.Println("\nThe application is started by the user ...")
		c.Printf("... %v!\n\n", usr.Name)
	case "info":
		repository.InfoV()
	case "?":
		repository.GetHelp()
	default:
		return
	}
}

// Инициализация подключения к БД
func initDB(identificatorDB string) (db *sql.DB) {
	db, err := sql.Open("mysql", identificatorDB)
	if err != nil {
		log.Fatalf("ERROR connected DB %v\n", err)
	}
	if err = db.Ping(); err != nil {
		log.Fatalf("ERROR DB Ping %v\n", err)
	}
	return db
}

// func main() {
// 	flag.Parse()

// 	var vvod int
// 	repository.ClearCMD()
// 	repository.InfoV()
// 	repository.Info1()
// 	fmt.Scan(&vvod)
// 	fmt.Println()
// 	repository.ClearCMD()
// 	switch {
// 	case vvod == 1:
// 		repository.InputENB()
// 	case vvod == 2:
// 		repository.InputSEC()
// 	case vvod == 3:
// 		repository.ClearCMD()
// 		repository.InfoSearchENB()

// 		var a string
// 		fmt.Scan(&a)
// 		switch {
// 		case a == "*":
// 			repository.SearchList()
// 			fmt.Scanln()
// 			fmt.Scanln()
// 		case a == "г":
// 			var nummmm string
// 			fmt.Print("Введите город: ")
// 			fmt.Scan(&nummmm)
// 			repository.SearchCity(nummmm)
// 			fmt.Scanln()
// 			fmt.Scanln()
// 		case a == "о":
// 			var nummmm string
// 			fmt.Print("Введите область: ")
// 			fmt.Scan(&nummmm)
// 			repository.SearchRegion(nummmm)
// 			fmt.Scanln()
// 			fmt.Scanln()
// 		case a == "д":
// 			repository.SearchDel()
// 			fmt.Scanln()
// 			fmt.Scanln()
// 		default:
// 			repository.Search(a)
// 			fmt.Scanln()
// 			fmt.Scanln()
// 		}
// 		fmt.Scanln()
// 		repository.ClearCMD()

// 	case vvod == 4:
// 		var inp string
// 		fmt.Print("Введите значение из запроса: ")
// 		fmt.Scan(&inp)
// 		fmt.Println()
// 		repository.Convert(inp)
// 		fmt.Scanln()
// 		fmt.Scanln()
// 	case vvod == 5:
// 		fmt.Print("1-ввод БС ручками;\n2-ввод через файл TXT;\n_: ")
// 		var a int
// 		fmt.Scan(&a)
// 		switch {
// 		case a == 1:
// 			repository.SearchMTS()
// 		case a == 2:
// 			repository.TextSearchMTS()
// 		}
// 	case vvod == 666:
// 		// handler.OpenTestConnect()
// 		// repository.dirByName()
// 	case vvod == 0:
// 		fmt.Println("Bye!")
// 		os.Exit(0)
// 		return
// 	}
// }
