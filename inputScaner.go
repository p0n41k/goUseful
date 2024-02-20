package gouseful

import (
	"bufio"
	"fmt"
	"os"

	"golang.org/x/crypto/ssh/terminal"
)

func InputScaner() string {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	line := scanner.Text()
	if err := scanner.Err(); err != nil {
		fmt.Println("Ошибка при считывании:", err)
	}
	return line
}

func AnonimInput() string {
	password, err := terminal.ReadPassword(int(os.Stdin.Fd()))
	if err != nil {
		os.Exit(1)
	}

	return string(password)
}
