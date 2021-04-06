package main

import (
	"database/sql"
	"fmt"
	"os"
	"os/user"
	"strings"

	"github.com/ArtZhiv/sql_test/pkg/repository"
	"github.com/c-bata/go-prompt"
	"github.com/fatih/color"
	_ "github.com/go-sql-driver/mysql"
)

var (
	y = color.New(color.FgYellow).Add(color.Underline)
)

func main() {
	if repository.Compare() == false {
		usr, err := user.Current()
		if err != nil {
			repository.Error.Printf("ERROR")
			fmt.Printf(" ошибка определения активного пользователя %v\n", err)
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
				repository.Error.Printf("ERROR")
				fmt.Printf(" ошибка определения активного пользователя %v\n", err)
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
			repository.Error.Printf("ERROR")
			fmt.Printf(" ошибка определения активного пользователя %v\n", err)
		}
		c := color.New(color.FgCyan).Add(color.Underline)
		c.Printf("Bye %v!\n", usr.Name)
		os.Exit(0)
		return
	case "test", "t":
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
	case "upd":
		repository.UpdateTable()
		y.Printf("Обновлена таблица eNodeB, добавлено %v записей\n",
			repository.UpdateBaseStation(initDB(repository.ConnDB)))
		y.Printf("Обновлена таблица с секторами, добавлено %v записей\n",
			repository.UpdateSector(initDB(repository.ConnDB)))
	case "clear":
		repository.ClearCMD()
		usr, err := user.Current()
		if err != nil {
			repository.Error.Printf("ERROR")
			fmt.Printf(" ошибка определения активного пользователя %v\n", err)
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
		repository.Error.Printf("ERROR")
		fmt.Printf(" connected DB %v\n", err)
	}
	if err = db.Ping(); err != nil {
		repository.Error.Printf("ERROR")
		fmt.Printf(" DB Ping %v\n", err)
	}
	return db
}
