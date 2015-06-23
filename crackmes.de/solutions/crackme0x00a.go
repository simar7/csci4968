package main

import (
	"fmt"
	"io"
	"os"
	"os/exec"
)

func main() {
	subProcess := exec.Command("../crackme0x00a")

	stdin, err := subProcess.StdinPipe()
	if err != nil {
		fmt.Println(err)
	}

	defer stdin.Close()

	subProcess.Stdout = os.Stdout
	subProcess.Stderr = os.Stderr

	fmt.Println("STARTING CRACKME0x00a")
	if err = subProcess.Start(); err != nil {
		fmt.Println("An error has occured: ", err)
	}

	io.WriteString(stdin, "g00dJ0B!")
	fmt.Println("END")
}
