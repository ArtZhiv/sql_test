package inputsql

import (
	"database/sql"
	"fmt"
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

type sec struct {
	id      int
	number  int
	sector  string
	bant    int
	mts     bool
	life    bool
	a1      bool
	beCloud bool
}

const (
	layoutISO = "02-01-2006"
)

// Search ...
func Search() {
	db, err := sql.Open("mysql", "Artem:Artem$mena@tcp(192.168.37.64:3306)/beCloud_database")

	if err != nil {
		panic(err)
	}
	defer db.Close()

	var nummmm int
	fmt.Scan(&nummmm)

	rows, err := db.Query("select * from beCloud_database.eNodeB where number = ?", nummmm)
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
		d := dt.Format("01.02.2006")
		fmt.Println(nummmm, "не в коммерции на", d)
	} else {
		for _, p := range products {
			if p.demolition != "___" {
				len := len(p.demolition) - 4
				fmt.Println("+ -------------- + ---------- +", strings.Repeat("-", len), "+")
				fmt.Println("| ДЕМОНТИРОВАНА--|--", p.number, "--|--", p.demolition, "--|")
				fmt.Println("+ -------------- + ---------- +", strings.Repeat("-", len), "+")
				fmt.Println()
			} else {
				fmt.Println(p.number, p.address)
				fmt.Println("\t", p.region, ", ", p.address)
				fmt.Println("\t Vendor: ", p.vendor, "На площадке: ", p.place)
				fmt.Println()

				rows, err := db.Query("select * from beCloud_database.sector where number = ?", nummmm)
				if err != nil {
					panic(err)
				}
				defer rows.Close()
				slector := []sec{}
				for rows.Next() {
					l := sec{}
					err := rows.Scan(&l.id, &l.number, &l.sector, &l.bant, &l.mts, &l.life, &l.a1, &l.beCloud)
					if err != nil {
						fmt.Println(err)
						continue
					}
					slector = append(slector, l)
				}
				// var a int
				var nMts, nLife, nA1, nBecloud string
				for _, l := range slector {
					if l.mts == true {
						nMts = "MTS"
					} else {
						nMts = "---"
					}

					if l.life == true {
						nLife = "LIFE"
					} else {
						nLife = "----"
					}

					if l.a1 == true {
						nA1 = "A1"
					} else {
						nA1 = "--"
					}

					if l.beCloud == true {
						nBecloud = "beCloud"
					} else {
						nBecloud = "-------"
					}
					if len(l.sector) == 8 {
						fmt.Println("\t", " | ", l.sector, "  | band:", l.bant, " | ", nMts, " | ", nA1, " | ", nLife, " | ", nBecloud, " |")
					} else {
						fmt.Println("\t", " | ", l.sector, " | band:", l.bant, " | ", nMts, " | ", nA1, " | ", nLife, " | ", nBecloud, " |")
						// a := len(fmt.Println("\t", " | ", l.sector, "  | ", nMts, " | ", nA1, " | ", nLife, " | ", nBecloud, " | band:", l.bant))
					}
				}
				m := strings.Repeat("-", len(nMts)+2)
				l := strings.Repeat("-", len(nLife)+2)
				a := strings.Repeat("-", len(nA1)+2)
				b := strings.Repeat("-", len(nBecloud)+2)
				fmt.Println("\t  + ----------- + ----------- +", m, "+", a, "+", l, "+", b, "+")
				fmt.Println()
			}
		}
	}

}
