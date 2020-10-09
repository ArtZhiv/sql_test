package cmd

import "fmt"

// Info1 ...
func Info1() {
	ClearCMD()
	info := []string{
		"Введите цифру требуемого действия:",
		"\t1-обновление eNB;",
		"\t2-обновление секторов;",
		"\t3-вывод информации о eNB;",
		"\t4-Перевод для UNET (из HEX в DEC);",
		"\t5-РРЛ для СООО <<МТС>>",
	}
	for _, elem := range info {
		fmt.Println(elem)
	}
	fmt.Println()
	fmt.Print("_:  ")
}

// InfoSearchENB ...
func InfoSearchENB() {
	info := []string{
		"\tномер eNodeB если требуется найти только одну станцию |",
		"\t* - для ввода станций списком                         |",
		"\tг - город                                             |",
		"\tо - область                                           |",
		"\tд - демонтированные                                   |",
	}

	fmt.Println("+ ----------------------------------------------------------- +")
	for _, elem := range info {
		fmt.Println(elem)
	}
	fmt.Println("+ ----------------------------------------------------------- +")
	fmt.Println()
	fmt.Print("Введите: ")
}
