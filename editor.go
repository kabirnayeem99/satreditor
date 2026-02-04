package main

import (
	"bufio"
	"fmt"
	"os"

	"golang.org/x/term"
)

func enterRawMode(termFd int) (restore func(), err error) {
	oState, err := term.MakeRaw(termFd)

	if err != nil {
		return nil, fmt.Errorf("enter raw mode: %w", err)
	}

	restore = func() {
		term.Restore(termFd, oState)
	}

	return
}

func debugKey(b byte) {
	fmt.Printf("\r\nkey: 0x%02x (%q)\r\n", b, b)
}

func clearScreen() {
	fmt.Print("\x1b[2J\x1b[H")
}

func readByte(in *bufio.Reader) (byte, error) {
	b, err := in.ReadByte()

	if err != nil {
		return 0, fmt.Errorf("read byte: %w", err)
	}
	return b, nil
}

func printBanner() {
	fmt.Print("satr (raw mode) - ctrl+q to exit\r\n")
}

func RunEditor() error {
	termFd := int(os.Stdin.Fd())
	restore, err := enterRawMode(termFd)

	if err != nil {
		return fmt.Errorf("enter raw mode: %w", err)
	}

	defer restore()

	in := bufio.NewReader(os.Stdin)

	clearScreen()
	printBanner()

	line := make([]byte, 0, 128)

	redraw := func() {
		fmt.Print("\r\x1b[2K")
		fmt.Print(string(line))
	}

	for {

		b, err := readByte(in)

		if err != nil {
			return fmt.Errorf("failed run editor: %w", err)
		}

		if b == CtrlQ || b == CtrlC {
			return nil
		}

		switch b {
		case BsBs, BsDel:
			if len(line) > 0 {
				line = line[:len(line)-1]
			}
			redraw()
		case Enter, EnterA:
			fmt.Print("\r\n")
			line = line[:0]
		default:
			fmt.Print(string(b))
			// debugKey(b)
		}
	}

}
