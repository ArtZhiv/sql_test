package inputsql

import (
	"database/sql"
	"fmt"
	"strconv"

	"github.com/360EntSecGroup-Skylar/excelize"
	// "github.com/360EntSecGroup-Skylar/excelize/v2"
)

// InputSEC ...
func InputSEC() {
	db, err := sql.Open("mysql", "Artem:Artem$mena@tcp(192.168.37.64:3306)/beCloud_database")

	if err != nil {
		panic(err)
	}
	defer db.Close()

	// здесь производится очистка таблицы
	del, err := db.Exec("TRUNCATE TABLE beCloud_database.sector")
	if err != nil {
		panic(err)
	} else {
		fmt.Println("Table sector is CLEARED!!!", del)
	}

	f, err := excelize.OpenFile("//192.168.37.222/24x7/Макросы и шаблоны/Комерция для шаблонов и макросов/Commercial BS.xlsm")
	if err != nil {
		fmt.Println(err)
		return
	}

	// Get value from cell by given worksheet name and axis.
	sectList := f.GetSheetName(1)
	rowsSect := f.GetRows(sectList)
	var s int
	for rowSect := range rowsSect {
		s = rowSect
	}
	for i := 2; i <= s+1; i++ {
		// sector
		sNumber := f.GetCellValue(sectList, "A"+strconv.Itoa(i))
		sector := f.GetCellValue(sectList, "B"+strconv.Itoa(i))
		sBand := f.GetCellValue(sectList, "C"+strconv.Itoa(i))
		sMts := f.GetCellValue(sectList, "E"+strconv.Itoa(i))
		sLife := f.GetCellValue(sectList, "F"+strconv.Itoa(i))
		sA1 := f.GetCellValue(sectList, "G"+strconv.Itoa(i))
		sBecloud := f.GetCellValue(sectList, "H"+strconv.Itoa(i))
		//
		var mts, life, a1, becloud bool
		ssBand, _ := strconv.Atoi(sBand)

		if err != nil {
			fmt.Println(err.Error())
			return
		}

		if sMts == "Да" {
			mts = true
		} else {
			mts = false
		}
		if sLife == "Да" {
			life = true
		} else {
			life = false
		}
		if sA1 == "Да" {
			a1 = true
		} else {
			a1 = false
		}
		if sBecloud == "Да" {
			becloud = true
		} else {
			becloud = false
		}

		number, _ := strconv.Atoi(sNumber)
		result, err := db.Exec("INSERT INTO beCloud_database.sector (number, sector, band, mts, life, a1, beCloud) values (?, ?, ?, ?, ?, ?, ?)",
			number, sector, ssBand, mts, life, a1, becloud)
		if err != nil {
			panic(err)
		}

		fmt.Println(result.LastInsertId())
	}
}
