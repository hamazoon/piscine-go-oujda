package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	arg := os.Args[1:]
	if len(arg) == 0 {
		fmt.Println("File name missing")
		return
	}
	if len(arg) > 1 {
		fmt.Println("Too many arguments")
		return
	}
	filename := arg[0]
	file, err := os.Open(filename)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()
	_, err = io.Copy(os.Stdout, file)
	if err != nil {
		fmt.Println("Error reading file:", err)
	}
}
