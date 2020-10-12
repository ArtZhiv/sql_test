package inputsql

import (
	"database/sql"
	"fmt"
	"io"
	"os"
	"os/exec"
	"strings"
	"time"

	"github.com/ArtZhiv/sql_test/cmd"
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

// Search ...
func Search(nummmm string) {
	db, err := sql.Open("mysql", "Artem:Artem$mena@tcp(192.168.37.64:3306)/beCloud_database")

	if err != nil {
		panic(err)
	}
	defer db.Close()

	fmt.Println()

	rows, err := db.Query("SELECT * FROM beCloud_database.eNodeB WHERE number LIKE concat('%',?,'%')", nummmm)
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

				rows, err := db.Query("select * from beCloud_database.sector where number = ?", p.number)
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

// SearchList ...
func SearchList() {
	db, err := sql.Open("mysql", "Artem:Artem$mena@tcp(192.168.37.64:3306)/beCloud_database")

	if err != nil {
		panic(err)
	}
	defer db.Close()

	fmt.Println()
	fmt.Println("Все номера БС вносятся через пробел, когда ввели все номера ставим * и жмём ENTER")
	fmt.Println()
	fmt.Print("Введите список eNodeB: ")

	var nummmm []string
	var a string
	for {
		fmt.Scan(&a)
		if a == "*" {
			break
		} else {
			nummmm = append(nummmm, a)
		}
	}
	fmt.Println()

	for _, elem := range nummmm {
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

					rows, err := db.Query("select * from beCloud_database.sector where number = ?", p.number)
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
}

// SearchRegion ...
func SearchRegion(nummmm string) {
	db, err := sql.Open("mysql", "Artem:Artem$mena@tcp(192.168.37.64:3306)/beCloud_database")

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
	db, err := sql.Open("mysql", "Artem:Artem$mena@tcp(192.168.37.64:3306)/beCloud_database")

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
	db, err := sql.Open("mysql", "Artem:Artem$mena@tcp(192.168.37.64:3306)/beCloud_database")

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
		err := rows.Scan(&p.id, &p.number, &p.address, &p.vendor, &p.region, &p.province, &p.demolition, &p.place)
		if err != nil {
			fmt.Println(err)
			continue
		}
		products = append(products, p)
	}
	for _, p := range products {
		if p.demolition != "NULL" {
			fmt.Println("| ДЕМОНТИРОВАНА--|--", p.number, "--|--", p.demolition, "--|")
		}
	}
	fmt.Println()
}

// SearchMTS ...
func SearchMTS() {
	db, err := sql.Open("mysql", "Artem:Artem$mena@tcp(192.168.37.64:3306)/beCloud_database")

	if err != nil {
		panic(err)
	}
	defer db.Close()

	cmd.ClearCMD()
	fmt.Print("Введите список eNodeB: ")

	var nummmm []string
	var a string
	for {
		fmt.Scan(&a)
		if a == "*" {
			break
		} else {
			nummmm = append(nummmm, a)
		}
	}
	fmt.Println()

	for _, elem := range nummmm {
		if strings.Contains(elem, "LTE") {
			if (strings.Contains(elem, "(LTE),") || strings.Contains(elem, "(LTE).")) && len(elem) == 10 {
				a := "0" + elem[:4]
				FindMTSforText(a)
			} else if (strings.Contains(elem, "(LTE),") || strings.Contains(elem, "(LTE).")) && len(elem) == 9 {
				a := "00" + elem[:3]
				FindMTSforText(a)
			} else if strings.Contains(elem, "(LTE)") {
				if len(elem) == 9 {
					a := "0" + elem[:4]
					FindMTSforText(a)
				} else {
					a := "00" + elem[:3]
					FindMTSforText(a)
				}
			} else {
				fmt.Println()
			}
		} else if strings.Contains(elem, "UMTS") || strings.Contains(elem, "IP") || strings.Contains(elem, "FTTX") {
			continue
		} else {
			if len(elem) == 4 {
				a := "0" + elem
				FindMTSforText(a)
			} else if len(elem) == 3 {
				a := "00" + elem
				FindMTSforText(a)
			} else {
				FindMTSforText(elem)
			}
		}
	}
	fmt.Println()
}

// TextSearchMTS ...
func TextSearchMTS() {
	fmt.Println()
	path := "files/workMts.txt"

	fileCreate, err := os.Create(path)
	if err != nil {
		panic(err)
		os.Exit(1)
	} else {
		fileCreate.Close()
	}

	fmt.Println("Сейчас откроется текстовый файл вставьте то что прислал МТС и сохраните файл и нажмите ENTER")
	cmd := exec.Command("powershell", "/c", "./files/workMts.txt")
	cmd.Stdout = os.Stdout
	cmd.Run()
	fmt.Scanln()
	fmt.Scanln()

	file, err := os.Open(path)
	if err != nil {
		panic(err)
		os.Exit(1)
	}

	data := make([]byte, 1024)
	var n int
	for {
		n, err = file.Read(data)
		if err == io.EOF {
			break
		}
		vvv := string(data[:n])
		fmt.Println()
		w := strings.Split(vvv, " ")
		for _, elem := range w {
			if strings.Contains(elem, "LTE") {
				if len(elem) == 10 {
					a := "0" + elem[:4]
					FindMTSforText(a)
				} else {
					a := "00" + elem[:3]
					FindMTSforText(a)
				}
			} else {
				continue
			}
		}
	}
	file.Close()
	fmt.Println()

	os.Remove(path)
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
