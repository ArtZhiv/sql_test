package repository

import (
	"database/sql"
	"fmt"
	"strconv"
	"time"

	"github.com/360EntSecGroup-Skylar/excelize/v2"
	_ "github.com/mattn/go-sqlite3"
)

type product struct {
	id    int
	count string
}

func clearTab(table string) {
	tableInput := fmt.Sprintf("DELETE from %v", table)
	db, err := sql.Open("sqlite3", "../files/interrupting.db3")
	if err != nil {
		panic(err)
	}
	// defer db.Close()

	_, err = db.Exec(tableInput)
	if err != nil {
		panic(err)
	} else {
		fmt.Printf("Table %v is CLEARED!!!\n", table)
	}

}

// DATA TIME layout			Mon Jan 2 15:04:05 MST 2006
// ДД.ММ.ГГГГ чч:мм:сс
// [чч]:мм:сс;@

func InputToSQL() {
	var table string = "lte"
	clearTab(table)

	db, err := sql.Open("sqlite3", "../files/interrupting.db3")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	f, err := excelize.OpenFile("../files/Журнал простоев 2021.xlsx")
	if err != nil {
		fmt.Println(err)
		return
	}

	// Get value from cell by given worksheet name and axis.
	enbList := f.GetSheetName(0)
	rowsEnb, _ := f.GetRows(enbList)
	var e int
	for rowEnb := range rowsEnb {
		e = rowEnb
	}

	loc, _ := time.LoadLocation("Europe/Minsk")

	for i := 2; i <= e+1; i++ {
		name, _ := f.GetCellValue(enbList, "A"+strconv.Itoa(i))
		startTime, _ := f.GetCellValue(enbList, "B"+strconv.Itoa(i))
		endTime, _ := f.GetCellValue(enbList, "C"+strconv.Itoa(i))
		typeT, _ := f.GetCellValue(enbList, "E"+strconv.Itoa(i))
		responsibility, _ := f.GetCellValue(enbList, "F"+strconv.Itoa(i))
		service, _ := f.GetCellValue(enbList, "G"+strconv.Itoa(i))
		note, _ := f.GetCellValue(enbList, "H"+strconv.Itoa(i))
		alarm, _ := f.GetCellValue(enbList, "I"+strconv.Itoa(i))
		fio, _ := f.GetCellValue(enbList, "J"+strconv.Itoa(i))

		startTimeConvert, _ := time.ParseInLocation("02/01/2006 15:04:05", startTime, loc)
		endTimeConvert, _ := time.ParseInLocation("02/01/2006 15:04:05", endTime, loc)
		timeC := endTimeConvert.Sub(startTimeConvert)

		_, err := db.Exec(
			`INSERT INTO lte (
					Name,
					Start_time,
					End_time,
					Time,
					Type,
					Responsibility,
					Service,
					Note,
					Alarm,
					FIO)
					values (?, ?, ?, ?, ?, ?, ?, ?, ?, ?)`,
			name,
			startTimeConvert.Format("02.01.2006 15:04:05"),
			endTimeConvert.Format("02.01.2006 15:04:05"),
			timeC,
			typeT,
			responsibility,
			service,
			note,
			alarm,
			fio,
		)
		if err != nil {
			panic(err)
		}
	}

	rows, err := db.Query(`SELECT COUNT(*) FROM lte`)
	if err != nil {
		panic(err)
	}
	products := []product{}
	for rows.Next() {
		p := product{}
		err := rows.Scan(&p.count)
		if err != nil {
			fmt.Println(err)
			continue
		}
		products = append(products, p)
	}
	for _, p := range products {
		fmt.Println(p.count)
	}
}
