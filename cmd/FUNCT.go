package cmd

import (
	"os"
	"os/exec"
)

// ClearCMD ...
func ClearCMD() {
	cmd := exec.Command("powershell", "/c", "cls")
	cmd.Stdout = os.Stdout
	cmd.Run()
}

// StartProgramm ...
func StartProgramm() {
	cmd := exec.Command("powershell", "/c", "go run main.go")
	cmd.Stdout = os.Stdout
	cmd.Run()
}

/*
	cmd := exec.Command("powershell", "/c", "echo", c)
	cmd.Stdout = os.Stdout
	cmd.Run()
*/
