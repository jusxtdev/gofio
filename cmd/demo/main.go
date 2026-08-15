package main

import (
	"bufio"
	"errors"
	"fmt"
	"github.com/jusxtdev/gofio"
	"log"
	"os"
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
	SAVE
	DELETE
	EXIT
)

var choices = map[int]string{
	int(CREATE):        "Create a file",
	int(READ):          "Read file",
	int(PARSE):         "Parse file",
	int(APPEND):        "Append to file",
	int(SAVE):          "Save file",
	int(DELETE):        "Delete file",
	int(RE_INITIALIZE): "Re-Initialize file",
	int(EXIT):          "Exit",
	int(FILE_INFO):     "Print file info",
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
		case int(RE_INITIALIZE):
			init_fh(&fh)
		case int(FILE_INFO):
			file_info(&fh)
		case int(PARSE):
			parse_file(&fh)
		case int(READ):
			read_file(&fh)
		case int(APPEND):
			append_to_file(&fh)
		case int(SAVE):
			save_file(fh)
		case int(DELETE):
			delete_file(fh)
		case int(EXIT):
			fmt.Println("Exiting ...")
			return
		default:
			fmt.Println("Not Implemented, Come back later !")
		}

	}
}

func menu() int {
	var sorted_keys = get_sorted_keys()

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
	var filepath string
	fmt.Printf("Enter filepath : ")
	_, err := fmt.Scanln(&filepath)
	if err != nil {
		log.Fatal(err)
	}

	fh.Initialize(filepath)
}

func create(fh *gofio.Gofio) {
	err := fh.Create()
	if err != nil {
		log.Fatal(err)
	}
}

func file_info(fh *gofio.Gofio) {
	fp := fh.Get_filepath()
	fext := fh.Get_file_extension()
	if fp == "" || fext == "" {
		fmt.Println("File not initialzed")
		return
	}
	fmt.Printf("Filepath : %s\nFile Extension : %s\n", fp, fext)
}

func parse_file(fh *gofio.Gofio) {
	err := fh.Parse()
	if err != nil {
		fmt.Println(err)
	}
}

func read_file(fh *gofio.Gofio) {
	// read file
	data, err := fh.Read()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(data)
}

func append_to_file(fh *gofio.Gofio) {
	if !fh.Check_parsed() {
		fmt.Println("File Not parsed !")
		return
	}

	var content string

	scanner := bufio.NewScanner(os.Stdin)

	fmt.Printf("Enter Content to append (EOF or Ctrl+d to stop)\n")
	// this loop reads each line, return `false` if EOF encountered
	fmt.Printf("> ")
	for scanner.Scan() {
		fmt.Printf("> ")
		line := scanner.Text()
		content = content + line + "\n"
	}
	// err check for scanner, it will not throw err if EOF otherwise raises error
	if err := scanner.Err(); err != nil {
		fmt.Println(err)
		return
	}

	err := fh.Append(content)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("Succesfully append %d bytes\n", len(content))
}

func save_file(fh gofio.Gofio) {
	err := fh.Save()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("File saved successfully")
}

func delete_file(fh gofio.Gofio) {
	// confirm before deleting
	prompt := fmt.Sprintf("Are you sure you want to delete : %s ?", fh.Get_filepath())
	confirm, err := yes_no_prompt(prompt)
	if err != nil {
		fmt.Println(err)
		return
	}

	if !confirm {
		return
	}

	err = fh.Delete()
	if err != nil {
		fmt.Println(err)
		return
	}

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

func yes_no_prompt(s string) (bool, error) {
	reader := bufio.NewReader(os.Stdin)
	fmt.Printf("%s [y/n] : ", s)

	char, _, err := reader.ReadRune()
	if err != nil {
		return false, err
	}

	switch char {
	case 'y':
		return true, nil
	case 'n':
		return false, nil
	default:
		return false, errors.New("invalid value")
	}
}

