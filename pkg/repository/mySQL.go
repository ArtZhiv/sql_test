package repository

import (
	"database/sql"
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"

	"github.com/360EntSecGroup-Skylar/excelize/v2"
)

// Инициализация подключения к БД
func initDB(identificatorDB string) (db *sql.DB) {
	db, err := sql.Open("mysql", identificatorDB)
	if err != nil {
		Error.Printf("ERROR")
		fmt.Printf(" connected DB %v\n", err)
	}
	if err = db.Ping(); err != nil {
		Error.Printf("ERROR")
		fmt.Printf(" DB Ping %v\n", err)
	}
	return db
}

// Обновление таблиц БД: eNodeB и Секторов
// UpdateTable ...
func UpdateTable() {
	var (
		db      = initDB(ConnDB)
		newDate = listDirByReadDir(PathCommercial)
	)
	truncate(db, "commercialUPD")
	_, err := db.Exec("INSERT INTO beCloud_database.commercialUPD (dateCreate) values (?)",
		newDate,
	)
	if err != nil {
		Error.Printf("ERROR")
		fmt.Printf(" inserted DB commercialUPD %v", err)
	}
}

// здесь производится очистка таблицы
func truncate(db *sql.DB, dbName string) {
	dbFull := fmt.Sprintf("TRUNCATE TABLE beCloud_database.%v", dbName)
	_, err := db.Exec(dbFull)
	if err != nil {
		Error.Printf("ERROR")
		fmt.Printf(" очистки таблицы %v\n", err)
	} else {
		fmt.Printf("Table %v is CLEARED!!!\n", dbName)
	}
}

// показывает колличество записей в таблице
func counterRows(db *sql.DB, dbName string) string {
	dbFull := fmt.Sprintf("SELECT count(*) FROM beCloud_database.%v", dbName)
	rows, err := db.Query(dbFull)
	if err != nil {
		Error.Printf("ERROR")
		fmt.Printf(" запроса количества записей %v", err)
	}
	var count string
	counters := []counter{}
	for rows.Next() {
		p := counter{}
		err := rows.Scan(&p.count)
		if err != nil {
			Error.Printf("ERROR")
			fmt.Printf(" декодинга переменной колличества записей %v", err)
			continue
		}
		counters = append(counters, p)
	}
	for _, p := range counters {
		count = p.count
	}
	return count
}

func UpdateBaseStation(db *sql.DB) string {
	truncate(db, "eNodeB")
	defer db.Close()

	f, err := excelize.OpenFile(PathCommercial + "Commercial BS.xlsm")
	if err != nil {
		Error.Printf("ERROR")
		fmt.Printf(" открытия файла Excel %v\n", err)
	}

	// Get value from cell by given worksheet name and axis.
	enbList := f.GetSheetName(0)
	rowsEnb, _ := f.GetRows(enbList)
	var (
		e int

		commercial_mts,
		commercial_life,
		commercial_a1 bool
	)
	for rowEnb := range rowsEnb {
		e = rowEnb
	}
	for i := 2; i <= e; i++ {
		number, _ := f.GetCellValue(enbList, "A"+strconv.Itoa(i))      // номер
		dismantling, _ := f.GetCellValue(enbList, "M"+strconv.Itoa(i)) // демонтаж
		area, _ := f.GetCellValue(enbList, "G"+strconv.Itoa(i))        // область
		district, _ := f.GetCellValue(enbList, "F"+strconv.Itoa(i))    // район
		city, _ := f.GetCellValue(enbList, "E"+strconv.Itoa(i))        // город
		address, _ := f.GetCellValue(enbList, "B"+strconv.Itoa(i))     // адрес
		vendor, _ := f.GetCellValue(enbList, "D"+strconv.Itoa(i))      // вендор
		location, _ := f.GetCellValue(enbList, "V"+strconv.Itoa(i))    // площадка
		sMts, _ := f.GetCellValue(enbList, "O"+strconv.Itoa(i))        // МТС
		sLife, _ := f.GetCellValue(enbList, "P"+strconv.Itoa(i))       // Life
		sA1, _ := f.GetCellValue(enbList, "Q"+strconv.Itoa(i))         // A1

		sMts = strings.ToUpper(sMts)
		sLife = strings.ToUpper(sLife)
		sA1 = strings.ToUpper(sA1)

		if strings.Contains(sMts, "ДА") {
			commercial_mts = true
		} else {
			commercial_mts = false
		}
		if strings.Contains(sLife, "ДА") {
			commercial_life = true
		} else {
			commercial_life = false
		}
		if strings.Contains(sA1, "ДА") {
			commercial_a1 = true
		} else {
			commercial_a1 = false
		}
		// fmt.Println(number, "-|-", fAdr, "-|-", vendor, "-|-", area, "-|-", district, "-|-", dismantling, "-|-", location, "-|-", commercial_mts, "-|-", commercial_life, "-|-", commercial_a1)
		if dismantling != "" {
			number, _ := strconv.Atoi(number)
			_, err := db.Exec("INSERT INTO beCloud_database.eNodeB (number, dismantling, area, district, city, address, vendor, location, commercial_mts, commercial_life, commercial_a1) values (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)",
				number,
				dismantling,
				"???",
				"???",
				"???",
				"This eNodeB is dismantled",
				"???",
				location,
				commercial_mts,
				commercial_life,
				commercial_a1,
			)
			if err != nil {
				Error.Printf("ERROR")
				fmt.Printf(" запроса к БД enodeB %v\n", err)
			}
		} else {
			number, _ := strconv.Atoi(number)
			_, err := db.Exec("INSERT INTO beCloud_database.eNodeB (number, dismantling, area, district, city, address, vendor, location, commercial_mts, commercial_life, commercial_a1) values (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)",
				number,
				"NULL",
				area,
				district,
				city,
				address,
				vendor,
				location,
				commercial_mts,
				commercial_life,
				commercial_a1,
			)
			if err != nil {
				Error.Printf("ERROR")
				fmt.Printf(" запроса к БД enodeB %v\n", err)
			}
		}
	}
	return counterRows(db, "eNodeB")
}

func UpdateSector(db *sql.DB) string {
	truncate(db, "sector")
	defer db.Close()

	f, err := excelize.OpenFile(PathCommercial + "Commercial BS.xlsm")
	if err != nil {
		Error.Printf("ERROR")
		fmt.Printf(" открытия файла Excel %v\n", err)
	}

	// Get value from cell by given worksheet name and axis.
	sectList := f.GetSheetName(1)
	rowsSect, _ := f.GetRows(sectList)
	var s int
	for rowSect := range rowsSect {
		s = rowSect
	}
	for i := 2; i <= s+1; i++ {
		basestantion, _ := f.GetCellValue(sectList, "A"+strconv.Itoa(i))
		cell_number, _ := f.GetCellValue(sectList, "B"+strconv.Itoa(i))
		bandwidth, _ := f.GetCellValue(sectList, "C"+strconv.Itoa(i))
		sMts, _ := f.GetCellValue(sectList, "E"+strconv.Itoa(i))
		sLife, _ := f.GetCellValue(sectList, "F"+strconv.Itoa(i))
		sA1, _ := f.GetCellValue(sectList, "G"+strconv.Itoa(i))
		sBecloud, _ := f.GetCellValue(sectList, "H"+strconv.Itoa(i))

		var (
			mts,
			life,
			a1,
			beCloud bool
		)

		cell_numberNum, _ := strconv.Atoi(cell_number[7:])

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
			beCloud = true
		} else {
			beCloud = false
		}

		basestantionNum, _ := strconv.Atoi(basestantion)
		_, err := db.Exec("INSERT INTO beCloud_database.sector (basestantion, cell_number, bandwidth, mts, life, a1, beCloud) values (?, ?, ?, ?, ?, ?, ?)",
			basestantionNum,
			cell_numberNum,
			bandwidth,
			mts,
			life,
			a1,
			beCloud,
		)
		if err != nil {
			Error.Printf("ERROR")
			fmt.Printf(" запроса к БД Sector %v\n", err)
		}
	}
	return counterRows(db, "sector")
}

func listDirByReadDir(path string) string {
	var dateServer string

	lst, err := ioutil.ReadDir(path)
	if err != nil {
		Error.Printf("ERROR")
		fmt.Printf("ERROR read DIR %v\n", err)
	}
	for _, val := range lst {
		if strings.Contains(val.Name(), "Commercial BS.xlsm") {
			dateServer = val.ModTime().String()
		}
	}
	return dateServer
}

func listCommercialDateUPD(identificatorDB string) string {
	var dateDB string

	db := initDB(ConnDB)

	rows, err := db.Query("select * from beCloud_database.commercialUPD")
	if err != nil {
		Error.Printf("ERROR")
		fmt.Printf(" запроса к БД commercialUPD %v\n", err)
	}
	defer rows.Close()

	upd := []commercialUPD{}

	for rows.Next() {
		p := commercialUPD{}
		err := rows.Scan(&p.id, &p.dateCreate)
		if err != nil {
			Error.Printf("ERROR")
			fmt.Printf(" декодировани даты обновления %v\n", err)
			continue
		}
		upd = append(upd, p)
	}
	for _, p := range upd {
		dateDB = p.dateCreate
	}
	return dateDB
}

func Compare() bool {
	dateDB := listCommercialDateUPD(ConnDB)
	dateServer := listDirByReadDir(PathCommercial)

	if dateDB >= dateServer {
		return false
	} else {
		return true
	}
}
