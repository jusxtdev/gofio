package gofio

import "os"

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
}

func (fh *Gofio) Initialize(extension string, filepath string) {
	fh.extension = extension
	fh.filepath = filepath
}

func (fh *Gofio) Create() error {
	var file *os.File
	var err error

	// check if file exists
	_, err = os.Stat(fh.filepath)
	if !os.IsNotExist(err) {
		return nil
	} else if err != nil {
		return err
	}

	// else create one
	file, err = os.Create(fh.filepath)
	if err != nil {
		return err
	}

	defer func() {
		err = file.Close()
	}()

	if err != nil {
		return err
	}

	// initialize file
	var data string
	if fh.extension == JSON {
		data = "[]"
		// store initial data into memory
		fh.data = data

		data = ""
	}
	// write inital content
	_, err = file.Write([]byte(data))
	if err != nil {
		return err
	}

	return nil
}
