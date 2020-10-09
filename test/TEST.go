package test

import (
	"fmt"
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

	// file, err := os.Open("test.txt")
	// if err != nil {
	// 	panic(err)
	// 	os.Exit(1)
	// }
	// defer file.Close()

	// data := make([]byte, 1024)
	// var n int
	// // var nummmm []string

	// for {
	// 	n, err = file.Read(data)
	// 	if err == io.EOF {
	// 		break
	// 	}
	// 	fmt.Println(n)
	// 	fmt.Println(string(data[6:10]))
	// 	vvv := string(data[:n])
	// 	fmt.Println()
	// 	fmt.Println(vvv)

	// 	for _, elem := range vvv {
	// 		fmt.Println("ELEM:", elem)
	// 	}
	// }

	// for {
	// 	fmt.Scan(&a)
	// 	if a == "*" {
	// 		break
	// 	} else {
	// 		nummmm = append(nummmm, a)
	// 	}
	// }
	// fmt.Println(len(nummmm))
	// for idx, elem := range nummmm {
	// 	fmt.Println("IDX:", idx, "ELEM:", elem)
	// }

	// if strings.Contains(a, "*") {
	// 	fmt.Println("Massive")
	// } else {
	// 	fmt.Println("String", a)
	// 	break
	// }

	// var a string
	// fmt.Scanln(&a)
	// if a == "" {
	// 	fmt.Println("Enter")
	// } else {
	// 	fmt.Println("Error")
	// }
	for i := 1; i < 7; i++ {
		go factorial(i)
	}
	fmt.Scanln()
	fmt.Println("The End")
}

func factorial(n int) {
	if n < 1 {
		fmt.Println("Invalid input number")
		return
	}
	result := 1
	for i := 1; i <= n; i++ {
		result *= i
	}
	fmt.Println(n, "-", result)
}
