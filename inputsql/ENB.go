package inputsql

import (
	"database/sql"
	"fmt"
	"strconv"

	"github.com/360EntSecGroup-Skylar/excelize"
	// "github.com/360EntSecGroup-Skylar/excelize/v2"
)

// InputENB ...
func InputENB() {
	db, err := sql.Open("mysql", "Artem:Artem$mena@tcp(192.168.37.64:3306)/beCloud_database")

	if err != nil {
		panic(err)
	}
	defer db.Close()

	// здесь производится очистка таблицы
	del, err := db.Exec("TRUNCATE TABLE beCloud_database.eNodeB")
	if err != nil {
		panic(err)
	} else {
		fmt.Println("Table eNodeB is CLEARED!!!", del)
	}

	f, err := excelize.OpenFile("//192.168.37.222/24x7/Макросы и шаблоны/Комерция для шаблонов и макросов/Commercial BS.xlsm")
	if err != nil {
		fmt.Println(err)
		return
	}

	// Get value from cell by given worksheet name and axis.
	enbList := f.GetSheetName(0)
	fmt.Println(enbList)
	rowsEnb := f.GetRows(enbList)
	var e int
	for rowEnb := range rowsEnb {
		e = rowEnb
	}

	for i := 2; i <= e; i++ {
		dem := f.GetCellValue(enbList, "M"+strconv.Itoa(i))
		number := f.GetCellValue(enbList, "A"+strconv.Itoa(i))
		// address
		region := f.GetCellValue(enbList, "G"+strconv.Itoa(i))
		city := f.GetCellValue(enbList, "F"+strconv.Itoa(i))
		fAdr := f.GetCellValue(enbList, "B"+strconv.Itoa(i))
		//
		vendor := f.GetCellValue(enbList, "D"+strconv.Itoa(i))
		operator := f.GetCellValue(enbList, "U"+strconv.Itoa(i))

		if err != nil {
			fmt.Println(err.Error())
			return
		}

		if dem != "" {
			number, _ := strconv.Atoi(number)
			result, err := db.Exec("INSERT INTO beCloud_database.eNodeB (number, address, vendor, region, province, demolition, place) values (?, ?, ?, ?, ?, ?, ?)",
				number, "This eNodeB is dismantled", "???", "???", "???", dem, operator)
			if err != nil {
				panic(err)
			}
			fmt.Println(result.LastInsertId())
		} else {
			number, _ := strconv.Atoi(number)
			result, err := db.Exec("INSERT INTO beCloud_database.eNodeB (number, address, vendor, region, province, demolition, place) values (?, ?, ?, ?, ?, ?, ?)",
				number, fAdr, vendor, region, city, "___", operator)
			if err != nil {
				panic(err)
			}
			fmt.Println(result.LastInsertId())
		}
	}
}
