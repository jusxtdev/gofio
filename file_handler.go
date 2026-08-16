package gofio

import (
	"errors"
	"fmt"
	"os"
	"path"
)

// Supported file extension
const (
	TXT      string = "txt"
	MARKDOWN string = "md"
	JSON     string = "json"
)

type Gofio struct {
	extension string
	filepath  string
	data      string
	is_parsed bool
}

func (fh *Gofio) Initialize(filepath string) error {
	// check whether file extension is supported or not
	ext := path.Ext(filepath)
	if ext == "" {
		return errors.New("invalid filepath")
	}
	ext = ext[1:]
	if !is_valid_file_ext(ext) {
		return errors.New("file extension not supported")
	}

	fh.filepath = filepath
	fh.extension = ext
	fh.is_parsed = false
	return nil
}

func (fh *Gofio) Get_filepath() string {
	return fh.filepath
}

func (fh *Gofio) Get_file_extension() string {
	return fh.extension
}

func (fh *Gofio) Get_file_data() (string, error) {
	if !fh.is_parsed {
		return "", errors.New("file not parsed")
	}
	return fh.data, nil
}

func (fh Gofio) Check_parsed() bool {
	return fh.is_parsed

}

func (fh *Gofio) Create() error {
	var file *os.File
	var err error

	// check if file exists
	_, err = os.Stat(fh.filepath)
	if !os.IsNotExist(err) {
		return nil
	}

	// else create one
	file, err = os.Create(fh.filepath)
	if err != nil {
		return err
	} else {
		fmt.Println("File created")
	}

	defer func() {
		err = file.Close()
	}()

	// initialize file
	var data string
	if fh.extension == JSON {
		data = "[]"
	} else {
		data = ""
	}
	// store initial data into memory
	fh.data = data

	// write inital content
	_, err = file.Write([]byte(data))
	if err != nil {
		return err
	}

	fh.is_parsed = true

	return nil
}

func (fh *Gofio) Parse() error {
	data, err := os.ReadFile(fh.filepath)
	if err != nil {
		return err
	}
	fh.data = string(data)
	fh.is_parsed = true
	return nil
}

func (fh *Gofio) Read() (string, error) {
	data, err := fh.Get_file_data()
	if err != nil {
		return "", err
	}

	return data, nil
}

func (fh *Gofio) Append(content string) error {
	if !fh.is_parsed {
		return errors.New("file not parsed")
	}

	fh.data = fh.data + content
	return nil
}

func (fh Gofio) Save() error {
	if !fh.is_parsed {
		return errors.New("file not parsed")
	}

	content := []byte(fh.data)
	err := os.WriteFile(fh.filepath, content, 0644)
	if err != nil {
		return err
	}
	return nil
}

func (fh Gofio) Delete() error {
	err := os.Remove(fh.filepath)
	if err != nil {
		return err
	}
	return nil
}

/*  HELPERS  */
func is_valid_file_ext(ext string) bool {
	switch ext {
	case TXT, MARKDOWN, JSON:
		return true
	default:
		return false
	}
}
