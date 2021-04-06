package repository

import (
	"database/sql"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"strings"
	"time"

	"github.com/gookit/color"
)

var (
	db      *sql.DB = initDB(ConnDB)
	pathMTS         = "../files/workMts.txt"
)

var (
	delete       = color.New(color.FgRed, color.OpBold)
	notDelete    = color.New(color.FgGreen)
	notInComerce = color.New(color.FgYellow)
)

// Search ...
func Search(nummmm string) {
	var (
		nMts,
		nLife,
		nA1,
		nBecloud,
		pMts,
		pA1,
		pLife string
	)

	rows, err := db.Query("SELECT * FROM beCloud_database.eNodeB WHERE number LIKE concat('%',?,'%')",
		nummmm,
	)
	if err != nil {
		Error.Printf("ERROR")
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
			Error.Printf("ERROR")
			fmt.Println(err)
			continue
		}
		products = append(
			products,
			p,
		)
	}
	if len(products) == 0 {
		dt := time.Now()
		d := dt.Format("02.01.2006")
		notInComerce.Printf("%v\nна %v eNodeB не в коммерции\n",
			nummmm,
			d,
		)
	} else {
		for _, p := range products {
			if p.dismantling != "NULL" {
				delete.Println("+ --------------- + ---------- +")
				delete.Printf("| ДЕМОНТИРОВАНА --|-- %v --|-- %v --|\n",
					p.number,
					p.dismantling,
				)
				delete.Println("+ --------------- + ---------- +")
			} else {
				fmt.Println()
				notDelete.Println(
					p.number,
					p.address,
				)
				fmt.Println()
				notDelete.Printf("\t Область: %v и район: %v\n",
					p.area,
					p.district,
				)
				fmt.Printf("\t Vendor:  %v \n\t На площадке:  %v\n",
					p.vendor,
					p.location,
				)

				if p.mts == true {
					pMts = "MTS"
				} else {
					pMts = "---"
				}

				if p.life == true {
					pLife = "LIFE"
				} else {
					pLife = "----"
				}

				if p.a1 == true {
					pA1 = "A1"
				} else {
					pA1 = "--"
				}
				fmt.Println()
				fmt.Printf("\t     Присутствуют:  |  %v  |  %v  |  %v  | \n",
					pMts,
					pA1,
					pLife,
				)
				fmt.Println("\t + ---------------- + ----- + ---- + ------ +")

				rows, err := db.Query("select * from beCloud_database.sector where basestantion = ?", p.number)
				if err != nil {
					Error.Printf("ERROR")
				}
				defer rows.Close()
				slector := []sectors{}
				for rows.Next() {
					l := sectors{}
					err := rows.Scan(
						&l.id,
						&l.basestantion,
						&l.cell_number,
						&l.bandwidth,
						&l.mts,
						&l.life,
						&l.a1,
						&l.beCloud,
					)
					if err != nil {
						Error.Printf("ERROR")
						fmt.Println(err)
						continue
					}
					slector = append(
						slector,
						l,
					)
				}
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
					if len(l.cell_number) == 1 {
						fmt.Printf("\t | 0%v | band: %v  |  %v  |  %v  |  %v  |  %v  |\n",
							l.cell_number,
							l.bandwidth,
							nMts,
							nA1,
							nLife,
							nBecloud,
						)
					} else {
						fmt.Printf("\t | %v | band: %v  |  %v  |  %v  |  %v  |  %v  |\n",
							l.cell_number,
							l.bandwidth,
							nMts,
							nA1,
							nLife,
							nBecloud,
						)
					}
				}
				m := strings.Repeat("-", len(nMts)+2)
				l := strings.Repeat("-", len(nLife)+2)
				a := strings.Repeat("-", len(nA1)+2)
				b := strings.Repeat("-", len(nBecloud)+2)
				fmt.Println("\t + -- + ----------- +", m, "+", a, "+", l, "+", b, "+")
				fmt.Println()
			}
		}
	}
}

// OpenFileMTS ...
func OpenFileMTS() {
	fileCreate, err := os.Create(pathMTS)
	if err != nil {
		Error.Printf("ERROR")
	} else {
		fileCreate.Close()
	}

	fmt.Printf(`
Сейчас откроется текстовый файл в который нужно вставить 
БС из файлов работ от МТС.
Для корректной обработки номеров значения вставляются с закрывающей скобкой
и разделяются пробелами.

Пример: 5764(LTE)
	` + "\n")
	cmd := exec.Command("powershell", "/c", "../files/workMts.txt")
	cmd.Stdout = os.Stdout
	cmd.Run()
}

// TextSearchMTS ...
func TextSearchMTS(abort string) {
	defer os.Remove(pathMTS)
	lst, err := ioutil.ReadDir("../files")
	if err != nil {
		Error.Printf("ERROR")
		log.Fatalf("ERROR read DIR %v\n", err)
	}
	for _, val := range lst {
		if strings.Contains(val.Name(), "workMts.txt") {
			file, err := os.Open(pathMTS)
			if err != nil {
				Error.Printf("ERROR")
			}

			data := make([]byte, 1024)
			var n int
			for {
				n, err = file.Read(data)
				if err == io.EOF {
					Error.Printf("ERROR")
					break
				}
				vvv := string(data[:n])
				fmt.Println()
				w := strings.Split(vvv, " ")

				if strings.Contains(abort, "H") || strings.Contains(abort, "h") {
					fmt.Printf("При проведении работ планируется прерывание сервиса LTE до %v часов на каждой eNodeB из списка:\n",
						abort[:len(abort)-1])
					parsTheDestination(w)
					fmt.Println()
				} else {
					fmt.Printf("При проведении работ планируется прерывание сервиса LTE до %v минут на каждой eNodeB из списка:\n",
						abort[:len(abort)-1])
					parsTheDestination(w)
					fmt.Println()
				}

			}
			file.Close()
		}
	}
}

func parsTheDestination(w []string) {
	for _, elem := range w {
		if strings.Contains(elem, "LTE") {
			if elem[len(elem)-1:] == "," || elem[len(elem)-1:] == "." {
				if len(elem) == 10 {
					a := "0" + elem[:4]
					findWorkForMTS(a)
				} else {
					a := "00" + elem[:3]
					findWorkForMTS(a)
				}
			} else {
				if len(elem) == 9 {
					a := "0" + elem[:4]
					findWorkForMTS(a)
				} else {
					a := "00" + elem[:3]
					findWorkForMTS(a)
				}
			}
		} else {
			continue
		}
	}
}

func findWorkForMTS(elem string) {
	rows, err := db.Query("SELECT * FROM beCloud_database.eNodeB WHERE number LIKE concat('%',?,'%')",
		elem,
	)
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
			fmt.Println(err)
			continue
		}
		products = append(
			products,
			p,
		)
	}
	if len(products) == 0 {
	} else {
		for _, p := range products {
			if p.dismantling != "NULL" {
				continue
			} else {
				fmt.Println(
					p.number,
					p.address,
				)
			}
		}
	}
}

// SearchRegion ...
/*
func SearchRegion(nummmm string) {
	db, err := sql.Open("mysql", ConnDB)

	if err != nil {
		panic(err)
	}
	defer db.Close()

	fmt.Println()

	rows, err := db.Query("SELECT * FROM beCloud_database.eNodeB WHERE region LIKE concat('%',?,'%')", nummmm)
	if err != nil {
		panic(err)
	}
	defer rows.Close()
	products := []enb{}

	for rows.Next() {
		p := enb{}
		err := rows.Scan(&p.id, &p.number, &p.address, &p.vendor, &p.region, &p.province, &p.demolition, &p.mts, &p.life, &p.a1, &p.place)
		if err != nil {
			fmt.Println(err)
			continue
		}
		products = append(products, p)
	}
	if len(products) == 0 {
		dt := time.Now()
		d := dt.Format("02.01.2006")
		fmt.Println("на", d, "eNodeB", nummmm, "не в коммерции")
	} else {
		for _, p := range products {
			if p.demolition != "NULL" {
				fmt.Println("+ -------------- + ---------- +")
				fmt.Println("| ДЕМОНТИРОВАНА--|--", p.number, "--|--", p.demolition, "--|")
				fmt.Println("+ -------------- + ---------- +")
				fmt.Println()
			} else {
				fmt.Println(p.number, p.address)
			}
		}
	}
}

// SearchCity ...
func SearchCity(nummmm string) {
	db, err := sql.Open("mysql", ConnDB)

	if err != nil {
		panic(err)
	}
	defer db.Close()

	fmt.Println()

	rows, err := db.Query("SELECT * FROM beCloud_database.eNodeB WHERE address LIKE concat('%',?,'%')", nummmm)
	if err != nil {
		panic(err)
	}
	defer rows.Close()
	products := []enb{}

	for rows.Next() {
		p := enb{}
		err := rows.Scan(&p.id, &p.number, &p.address, &p.vendor, &p.region, &p.province, &p.demolition, &p.mts, &p.life, &p.a1, &p.place)
		if err != nil {
			fmt.Println(err)
			continue
		}
		products = append(products, p)
	}
	if len(products) == 0 {
		dt := time.Now()
		d := dt.Format("02.01.2006")
		fmt.Println("на", d, "eNodeB", nummmm, "не в коммерции")
	} else {
		for _, p := range products {
			if p.demolition != "NULL" {
				fmt.Println("+ -------------- + ---------- +")
				fmt.Println("| ДЕМОНТИРОВАНА--|--", p.number, "--|--", p.demolition, "--|")
				fmt.Println("+ -------------- + ---------- +")
				fmt.Println()
			} else {
				fmt.Println(p.number, p.address)
				fmt.Println()
				fmt.Println("\t", p.region, ", ", p.address)
				fmt.Println("\t Vendor: ", p.vendor, "\n\t На площадке: ", p.place)
				fmt.Println()
			}
		}
	}
}

// SearchDel ...
func SearchDel() {
	db, err := sql.Open("mysql", ConnDB)

	if err != nil {
		panic(err)
	}
	defer db.Close()
	fmt.Println()

	rows, err := db.Query("SELECT * FROM beCloud_database.eNodeB")
	if err != nil {
		panic(err)
	}
	defer rows.Close()
	products := []enb{}

	for rows.Next() {
		p := enb{}
		err := rows.Scan(&p.id, &p.number, &p.address, &p.vendor, &p.region, &p.province, &p.demolition, &p.mts, &p.life, &p.a1, &p.place)
		if err != nil {
			fmt.Println(err)
			continue
		}
		products = append(products, p)
	}
	for _, p := range products {
		if p.demolition != "NULL" {
			delete.Println("| ДЕМОНТИРОВАНА--|--", p.number, "--|--", p.demolition, "--|")
		}
	}
	fmt.Println()
}
*/
