package repository

import (
	"database/sql"
	"fmt"
	"log"
	"strconv"
)

// ConvertIMSIToGlobalID ...
func ConvertIMSIToGlobalID(a string) {
	// Перевод для номера eNodeB
	c := a[5:10]
	ii, _ := strconv.ParseInt(c, 16, 64)
	mm := strconv.FormatInt(ii, 10)

	// Перевод для сектора eNodeB
	b := a[10:]
	i, _ := strconv.ParseInt(b, 16, 64)
	m := strconv.FormatInt(i, 10)

	db, err := sql.Open("mysql", "beclouderp:becloud$erp@tcp(192.168.37.65:3306)/beCloud_database")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	rows, err := db.Query("SELECT * FROM beCloud_database.eNodeB WHERE number LIKE concat('%',?,'%')", mm)
	if err != nil {
		panic(err)
	}
	defer rows.Close()
	products := []eNodeb{}

	for rows.Next() {
		p := eNodeb{}
		err := rows.Scan(
			&p.id,
			&p.number,
			&p.dismantling,
			&p.area,
			&p.district,
			&p.city,
			&p.address,
			&p.vendor,
			&p.location,
			&p.mts,
			&p.life,
			&p.a1,
		)
		if err != nil {
			log.Fatal(err)
			continue
		}
		products = append(products, p)
	}
	if len(products) == 0 {
		// dt := time.Now()
		// d := dt.Format("01.02.2006")
		// fmt.Println(mm, "не в коммерции на", d)
	} else {
		for _, p := range products {
			if p.dismantling != "NULL" {
			} else {
				fmt.Printf("Коллеги, данный клиент подключен к %v сектору eNodeB %v %v",
					m,
					p.number,
					p.address,
				)
				fmt.Println()
			}
		}
	}
}
