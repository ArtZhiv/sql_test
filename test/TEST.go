package test

import (
	"fmt"
	"io"
	"os"
	"strings"
)

// Test ...
func Test() {
	// var nummmm []string
	// nummmm := make([]string, 1024)
	// var a string

	// fmt.Println("Сейчас откроется текстовый файл вставьте то что прислал МТС и сохраните файл")
	// cmd := exec.Command("powershell", "/c", "./test.txt")
	// cmd.Stdout = os.Stdout
	// cmd.Run()

	file, err := os.Open("test.txt")
	if err != nil {
		panic(err)
		os.Exit(1)
	}
	defer file.Close()

	data := make([]byte, 1024)
	var n int
	// var nummmm []string

	for {
		n, err = file.Read(data)
		if err == io.EOF {
			break
		}
		vvv := string(data[:n])
		fmt.Println()
		// fmt.Println(vvv)
		w := strings.Split(vvv, ", ")
		for _, elem := range w {
			fmt.Println("e:", elem)
		}
	}
}
