package main

import (
	"fmt"
	"os"
	"os/exec"
)

func main() {
	fmt.Print("\x1b[8m")
	exec.Command("stty", "cbreak").Run()

	var b []byte = make([]byte, 1)
	for {
		os.Stdin.Read(b)
		fmt.Print("\n" + "\x1b[28m" + "I got the byte", b, "("+string(b)+")" + "\x1b[8m")
	}
}