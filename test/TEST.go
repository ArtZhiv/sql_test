package test

import (
	"os"
	"os/exec"
)

// Test ...
func Test() {
	// var nummmm []string
	// var a string
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
	cmd := exec.Command("powershell", "/c", "./test.txt")
	cmd.Stdout = os.Stdout
	cmd.Run()
}
