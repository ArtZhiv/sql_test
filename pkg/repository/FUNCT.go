package repository

import (
	"os"
	"os/exec"

	_ "github.com/go-sql-driver/mysql"
)

// ClearCMD ...
func ClearCMD() {
	cmd := exec.Command("powershell", "/c", "cls")
	cmd.Stdout = os.Stdout
	cmd.Run()
}
