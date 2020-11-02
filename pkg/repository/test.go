package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/user"
	"path/filepath"
	"strconv"
	"strings"

	_ "github.com/go-sql-driver/mysql" // ...
)

// type enb struct {
// 	ID         int    `db:"id"`
// 	number     int    `db:"id"`
// 	address    string `db:"id"`
// 	vendor     string `db:"id"`
// 	region     string `db:"id"`
// 	province   string `db:"district"`
// 	demolition string `db:"id"`
// 	mts        bool   `db:"id"`
// 	life       bool   `db:"id"`
// 	a1         bool   `db:"id"`
// 	place      string `db:"id"`
// }

// OpenTestConnect ...
// func OpenTestConnect() {
// 	csvFile, _ := os.Open("../files/atm_anomaly_console_1414585557.csv")
// 	reader := csv.NewReader(csvFile)

// 	var ddos []Ddos
// 	for {
// 		line, error := reader.Read()
// 		if error == io.EOF {
// 			break
// 		} else if error != nil {
// 			log.Fatalln(error)
// 		}
// 		ddos = append(ddos, Ddos{

// 		})
// 	}
// }

func listDirByWalk(path string) {
	filepath.Walk(path, func(wPath string, info os.FileInfo, err error) error {
		if wPath != path {
			fmt.Println(wPath)
		}
		return nil
	})
}

func main() {
	cur, err := user.Current()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Current user:",
		cur.HomeDir)
	fmt.Println()

	var f string
	f = cur.HomeDir + "\\Downloads"
	fmt.Println(f)
	//----------------------------------------
	files, err := ioutil.ReadDir(f)
	if err != nil {
		log.Fatal("error")
	}
	// fmt.Println(files)
	var dirDdos = map[string]int{}

	// for key, value := range people {
	// 	fmt.Println("key", value)
	// }

	for _, file := range files {
		// fmt.Println(file.Name(), "---", file.ModTime().Format("02012006150405"))
		if strings.Contains(file.Name(), "atm_anomaly") {
			// fmt.Printf("%s --- %v\n", file.Name(), file.ModTime().Format("20060102150405"))
			p, _ := strconv.Atoi(file.ModTime().Format("20060102150405"))
			dirDdos[file.Name()] = p
		}
		for key, value := range dirDdos {
			fmt.Println(key, "-||-", value)
		}
	}
}
