package gouseful

import (
	"bufio"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"golang.org/x/term"
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
	CtrlCExit()
	password, err := term.ReadPassword(int(os.Stdin.Fd()))
	if err != nil {
		os.Exit(1)
	}

	return string(password)
}

func CtrlCExit() {
	initialState, err := term.GetState(int(os.Stdin.Fd()))
	if err != nil {
		fmt.Println("Ошибка при получении состояния терминала:", err)
		return
	}

	// Обработчик сигнала прерывания (Ctrl + C)
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)

	go func() {
		<-c
		// Восстанавливаем состояние терминала перед выходом
		term.Restore(int(os.Stdin.Fd()), initialState)
		os.Exit(0)
	}()
}
