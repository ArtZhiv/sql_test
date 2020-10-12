package inputsql

import (
	"database/sql"
	"fmt"
	"strconv"
	"strings"

	// "github.com/360EntSecGroup-Skylar/excelize"
	"github.com/360EntSecGroup-Skylar/excelize/v2"
	"github.com/ArtZhiv/sql_test/cmd"
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
	rowsEnb, _ := f.GetRows(enbList)
	var e int
	var mts, life, a1 bool
	for rowEnb := range rowsEnb {
		e = rowEnb
	}
	var count int64
	for i := 2; i <= e; i++ {
		dem, _ := f.GetCellValue(enbList, "M"+strconv.Itoa(i))
		number, _ := f.GetCellValue(enbList, "A"+strconv.Itoa(i))
		// address
		region, _ := f.GetCellValue(enbList, "G"+strconv.Itoa(i))
		district, _ := f.GetCellValue(enbList, "F"+strconv.Itoa(i))
		fAdr, _ := f.GetCellValue(enbList, "B"+strconv.Itoa(i))
		//
		vendor, _ := f.GetCellValue(enbList, "D"+strconv.Itoa(i))
		place, _ := f.GetCellValue(enbList, "U"+strconv.Itoa(i))
		sMts, _ := f.GetCellValue(enbList, "O"+strconv.Itoa(i))
		sLife, _ := f.GetCellValue(enbList, "P"+strconv.Itoa(i))
		sA1, _ := f.GetCellValue(enbList, "Q"+strconv.Itoa(i))

		if err != nil {
			fmt.Println(err.Error())
			return
		}
		sMts = strings.ToUpper(sMts)
		sLife = strings.ToUpper(sLife)
		sA1 = strings.ToUpper(sA1)

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
		// fmt.Println(number, "-|-", fAdr, "-|-", vendor, "-|-", region, "-|-", district, "-|-", dem, "-|-", place, "-|-", mts, "-|-", life, "-|-", a1)
		if dem != "" {
			number, _ := strconv.Atoi(number)
			result, err := db.Exec("INSERT INTO beCloud_database.eNodeB (number, address, vendor, region, district, demolition, mts, life, a1, place) values (?, ?, ?, ?, ?, ?, ?, ?, ?, ?)",
				number, "This eNodeB is dismantled", "???", "???", "???", dem, mts, life, a1, place)
			if err != nil {
				panic(err)
			}
			result.LastInsertId()
		} else {
			number, _ := strconv.Atoi(number)
			result, err := db.Exec("INSERT INTO beCloud_database.eNodeB (number, address, vendor, region, district, demolition, mts, life, a1, place) values (?, ?, ?, ?, ?, ?, ?, ?, ?, ?)",
				number, fAdr, vendor, region, district, "NULL", mts, life, a1, place)
			if err != nil {
				panic(err)
			}
			count, err = result.LastInsertId()
			if err != nil {
				panic(err)
			}
		}
	}
	cmd.ClearCMD()
	fmt.Println("Добавлено: ", count)
}
