package utils

import (
	"fmt"
	"os"

	"golang.org/x/term"
)

func Clear() {
	fmt.Print("\033[H\033[2J")
}

func Size() (int, int) {
	w, h, err := term.GetSize(int(os.Stdout.Fd()))
	if err != nil {
		return 0, 0
	}
	return w, h
}
