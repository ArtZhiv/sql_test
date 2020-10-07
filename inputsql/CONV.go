package inputsql

import (
	"database/sql"
	"fmt"
	"strconv"
	"strings"
	"time"
)

// enb2 ...
type enb2 struct {
	id         int
	number     int
	address    string
	vendor     string
	region     string
	province   string
	demolition string
	place      string
}

// Convert ...
func Convert(val string) {
	a, _ := strconv.Atoi(val)
	b := a % 100
	bb := strconv.Itoa(b)

	i, _ := strconv.ParseInt(bb, 16, 64)
	m := strconv.FormatInt(i, 10)

	c := (a / 100) % 100000
	cc := strconv.Itoa(c)

	ii, _ := strconv.ParseInt(cc, 16, 64)
	mm := strconv.FormatInt(ii, 10)

	db, err := sql.Open("mysql", "Artem:Artem$mena@tcp(192.168.37.64:3306)/beCloud_database")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	rows, err := db.Query("select * from beCloud_database.eNodeB where number like ?", mm)
	if err != nil {
		panic(err)
	}
	defer rows.Close()
	products := []enb2{}

	for rows.Next() {
		p := enb2{}
		err := rows.Scan(&p.id, &p.number, &p.address, &p.vendor, &p.region, &p.province, &p.demolition, &p.place)
		if err != nil {
			fmt.Println(err)
			continue
		}
		products = append(products, p)
	}
	if len(products) == 0 {
		dt := time.Now()
		d := dt.Format("01.02.2006")
		fmt.Println(mm, "не в коммерции на", d)
	} else {
		for _, p := range products {
			if p.demolition != "___" {
				len := len(p.demolition) - 4
				fmt.Println("+ -------------- + ---------- +", strings.Repeat("-", len), "+")
				fmt.Println("| ДЕМОНТИРОВАНА--|--", p.number, "--|--", p.demolition, "--|")
				fmt.Println("+ -------------- + ---------- +", strings.Repeat("-", len), "+")
				fmt.Println()
			} else {
				fmt.Println("Коллеги, данный клиент подключен к", m, "сектору на eNodeB", p.number, p.address)
				fmt.Println()
			}
		}
	}
}
