package main

import (
	"fmt"
	"gofio"
	"log"
)

type Option int

const (
	CREATE Option = iota + 1
	RE_INITIALIZE
	READ
	APPEND
	DELETE
	EXIT
)

var choices = map[string]int{
	"Create a file":      int(CREATE),
	"Read file":          int(READ),
	"Append to file":     int(APPEND),
	"Delete file":        int(DELETE),
	"Re-Initialize file": int(RE_INITIALIZE),
	"Exit":               int(EXIT),
}

func main() {
	fh := gofio.Gofio{}

	fmt.Printf("Initialize the file : \n")
	init_fh(&fh)

	for {
		choice := menu()

		switch choice {
		case 1:
			create(&fh)
		}
	}
}

func menu() int {
	for key, value := range choices {
		fmt.Printf("%d : %s\n", value, key)
	}

	var choice int

	fmt.Printf("Enter Choice : ")
	_, err := fmt.Scanln(&choice)
	if err != nil {
		log.Fatal(err)
	}

	return choice
}

func init_fh(fh *gofio.Gofio) {
	var ext string

	fmt.Printf("Enter Extension : ")

	_, err := fmt.Scanln(&ext)
	if err != nil {
		log.Fatal(err)
	}

	var filepath string
	fmt.Printf("Enter filepath : ")
	_, err = fmt.Scanln(&filepath)
	if err != nil {
		log.Fatal(err)
	}

	fh.Initialize(ext, filepath)
}

func create(fh *gofio.Gofio) {
	err := fh.Create()
	if err != nil {
		log.Fatal(err)
	}
}
