package gouseful

import (
	"os"
	"os/exec"
)

func ClearTheTerminal() {
	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout
	cmd.Run()
}
