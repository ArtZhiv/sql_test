package test

import (
	"fmt"
	"io/ioutil"
	"strings"
)

// type enb struct {
// 	id         int
// 	number     int
// 	address    string
// 	vendor     string
// 	region     string
// 	province   string
// 	demolition string
// 	place      string
// }

// Test ...
func Test() {
	var a, b string
	files, err := ioutil.ReadDir("./files/")
	if err != nil {
		panic(err)
	}
	for _, file := range files {
		fmt.Println(file.Name(), "-|-", file.IsDir(), "-|-", file.ModTime())
		if strings.Contains(file.Name(), "AlarmLogs") {
			a = file.Name()
			fmt.Println("-----|", a, "|-----")
			fmt.Println()
		}
		b = a[9:26]
		if a[9:26] < b {
			continue
		} else {
			fmt.Println("a__", a)
		}
		// fmt.Println("b:__", b)
	}

	// file, err := os.Open("files/.csv")
	// if err != nil {
	// 	panic(err)
	// 	os.Exit(1)
	// }
	// defer file.Close()

	fmt.Println()
}
