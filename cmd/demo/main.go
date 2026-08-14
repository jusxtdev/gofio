package main

import (
	"fmt"
	"gofio"
	"log"
	"slices"
)

type Option int

const (
	CREATE Option = iota + 1
	RE_INITIALIZE
	FILE_INFO
	READ
	APPEND
	DELETE
	EXIT
)

var choices = map[int]string{
	int(CREATE) : "Create a file",
	int(READ): "Read file",
	int(APPEND) : "Append to file",
	int(DELETE) : "Delete file",
	int(RE_INITIALIZE) : "Re-Initialize file",
	int(EXIT) : "Exit",
	int(FILE_INFO): "Print file info",
}

func main() {
	fh := gofio.Gofio{}

	fmt.Printf("Initialize the file : \n")
	init_fh(&fh)

	for {
		choice := menu()

		switch choice {
		case int(CREATE):
			create(&fh)
		case int(FILE_INFO):
			file_info(&fh)
		}

	}
}

func menu() int {
	var sorted_keys []int = get_sorted_keys()

	for _, key := range sorted_keys {
		fmt.Printf("%d : %s\n", key, choices[key])
	}

	var choice int

	fmt.Printf("Enter Choice : ")
	_, err := fmt.Scanln(&choice)
	if err != nil {
		log.Fatal(err)
	}

	return choice
}

/*  FUNCTIONALITIES  */

func init_fh(fh *gofio.Gofio) {
	var ext string

	var filepath string
	fmt.Printf("Enter filepath : ")
	_, err := fmt.Scanln(&filepath)
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

func file_info(fh *gofio.Gofio){
	fp := fh.Get_filepath()
	fext := fh.Get_file_extension()
	if fp == "" || fext == ""{
		fmt.Println("File not initialzed")
		return
	}
	fmt.Printf("Filepath : %s\nFile Extension : %s\n", fp, fext)
}

/*    HELPERS    */
func get_sorted_keys() []int {
	// choices map[int]string => we have to sort according to int here
	var keys []int

	// extract keys from choices
	for key := range choices {
		keys = append(keys, key)
	}

	// sort keys
	slices.Sort(keys)
	return keys
}