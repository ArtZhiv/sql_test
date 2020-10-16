package repository

import (
	"database/sql"
	"fmt"
	"strconv"
	"strings"

	"github.com/360EntSecGroup-Skylar/excelize/v2"
)

// InputSEC ...
func InputSEC() {
	db, err := sql.Open("mysql", "beclouderp:becloud$erp@tcp(192.168.37.65:3306)/beCloud_database")

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
	rowsSect, _ := f.GetRows(sectList)
	var s int
	var count int64
	for rowSect := range rowsSect {
		s = rowSect
	}
	for i := 2; i <= s+1; i++ {
		// sector
		sNumber, _ := f.GetCellValue(sectList, "A"+strconv.Itoa(i))
		sector, _ := f.GetCellValue(sectList, "B"+strconv.Itoa(i))
		sBand, _ := f.GetCellValue(sectList, "C"+strconv.Itoa(i))
		sMts, _ := f.GetCellValue(sectList, "E"+strconv.Itoa(i))
		sLife, _ := f.GetCellValue(sectList, "F"+strconv.Itoa(i))
		sA1, _ := f.GetCellValue(sectList, "G"+strconv.Itoa(i))
		sBecloud, _ := f.GetCellValue(sectList, "H"+strconv.Itoa(i))
		//
		var mts, life, a1, becloud bool
		ssBand, _ := strconv.Atoi(sBand)

		if err != nil {
			fmt.Println(err.Error())
			return
		}

		sMts = strings.ToUpper(sMts)
		sLife = strings.ToUpper(sLife)
		sA1 = strings.ToUpper(sA1)
		sBecloud = strings.ToUpper(sBecloud)

		if strings.Contains(sMts, "ДА") {
			mts = true
		} else {
			mts = false
		}
		if strings.Contains(sLife, "ДА") {
			life = true
		} else {
			life = false
		}
		if strings.Contains(sA1, "ДА") {
			a1 = true
		} else {
			a1 = false
		}
		if strings.Contains(sBecloud, "ДА") {
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

		count, err = result.LastInsertId()
		if err != nil {
			panic(err)
		}
	}
	ClearCMD()
	fmt.Println("Добавлено: ", count)
}