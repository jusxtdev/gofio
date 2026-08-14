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
	PARSE
	READ
	APPEND
	DELETE
	EXIT
)

var choices = map[int]string{
	int(CREATE) : "Create a file",
	int(READ): "Read file",
	int(PARSE): "Parse file",
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
		case int(PARSE):
			parse_file(&fh)
		case int(READ):
			read_file(&fh)
		case int(EXIT):
			fmt.Println("Exiting ...")
			return
		default:
			fmt.Println("Not Implemented, Come back later !")
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

func parse_file(fh *gofio.Gofio){
	err := fh.Parse()
	if err != nil {
		fmt.Println(err)
	}
}

func read_file(fh *gofio.Gofio){
	// read file
	data, err := fh.Read()
	if err != nil{
		fmt.Println(err)
		return
	}
	fmt.Println(data)
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