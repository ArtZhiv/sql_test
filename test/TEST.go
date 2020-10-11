package test

import (
	"database/sql"
	"encoding/csv"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
	"time"
)

type enb struct {
	id         int
	number     int
	address    string
	vendor     string
	region     string
	province   string
	demolition string
	place      string
}

type test struct {
	id      int
	number  int
	sector  string
	bant    int
	mts     bool
	life    bool
	a1      bool
	beCloud bool
}

// Test ...
func Test() {
	db, err := sql.Open("mysql", "Artem:Artem$mena@tcp(192.168.37.64:3306)/beCloud_database")

	if err != nil {
		panic(err)
	}
	defer db.Close()

	var a, b, c string
	files, err := ioutil.ReadDir("./files/")
	if err != nil {
		panic(err)
	}
	for _, file := range files {
		if strings.Contains(file.Name(), "AlarmLogs") {
			a = file.Name()
		}
		if b < a[9:26] {
			b = a[9:26]
			c = a
		} else {
			continue
		}
	}
	fmt.Println()
	path := "files/" + c
	file, err := os.Open(path)
	if err != nil {
		panic(err)
		os.Exit(1)
	}
	defer file.Close()

	reader := csv.NewReader(file)
	reader.FieldsPerRecord = -1

	for {
		record, err := reader.Read()
		if err != nil {
			panic(err)
		}
		for _, elem := range record {
			if elem == "S1ap Link Down" {
				w := strings.Split(record[7], ", ")
				for idx, elem := range w {
					if idx == 8 {
						clearedOn, err := time.ParseInLocation("02.01.2006 15:04:05", record[9], time.Local)
						if err != nil {
							panic(err)
						}
						occurredOn, err := time.ParseInLocation("02.01.2006 15:04:05", record[8], time.Local)
						if err != nil {
							panic(err)
						}
						FindMTSforText(elem[11:])
						fmt.Print(occurredOn.Format("02.01.2006 15:04:05"), "||", clearedOn.Format("02.01.2006 15:04:05"), " : ", record[10])
					}
				}
				fmt.Println()
			}
		}
	}
}

// FindMTSforText ...
func FindMTSforText(elem string) {
	db, err := sql.Open("mysql", "Artem:Artem$mena@tcp(192.168.37.64:3306)/beCloud_database")

	if err != nil {
		panic(err)
	}
	defer db.Close()

	rows, err := db.Query("SELECT * FROM beCloud_database.eNodeB WHERE number LIKE concat('%',?,'%')", elem)
	if err != nil {
		panic(err)
	}
	defer rows.Close()
	products := []enb{}

	for rows.Next() {
		p := enb{}
		err := rows.Scan(&p.id, &p.number, &p.address, &p.vendor, &p.region, &p.province, &p.demolition, &p.place)
		if err != nil {
			fmt.Println(err)
			continue
		}
		products = append(products, p)
	}
	if len(products) == 0 {
		dt := time.Now()
		d := dt.Format("02.01.2006")
		fmt.Println("на", d, "eNodeB", elem, "не в коммерции")
	} else {
		for _, p := range products {
			if p.demolition != "___" {
				fmt.Println("+ -------------- + ---------- +")
				fmt.Println("| ДЕМОНТИРОВАНА--|--", p.number, "--|--", p.demolition, "--|")
				fmt.Println("+ -------------- + ---------- +")
				fmt.Println()
			} else {
				fmt.Print(p.number, " ", p.address, " ")
			}
		}
	}
}
