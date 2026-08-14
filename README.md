# gofio

A small Go library that simplifies file handling for text-based files.

## Supported formats

- `.txt`
- `.md`
- `.json`

## Usage

```go
package main

import (
	"fmt"
	"gofio"
)

func main() {
	var fh gofio.Gofio
	fh.Initialize("", "notes.txt")

	fh.Create()
	fh.Parse()
	fh.Append("hello world\n")
	fh.Save()

	data, _ := fh.Read()
	fmt.Println(data)

	fh.Delete()
}
```

## API

| Method | Description |
| --- | --- |
| `Initialize(extension, filepath)` | Sets up the handler. The file extension is inferred from the filepath. |
| `Create()` | Creates the file if it does not exist (JSON files start as `[]`). |
| `Parse()` | Reads the file contents into memory. |
| `Read()` | Returns the file contents as a string (requires `Parse()` first). |
| `Append(content)` | Appends content to the in-memory data. |
| `Save()` | Writes the in-memory data to disk. |
| `Delete()` | Removes the file from disk. |
| `Get_filepath()` | Returns the file path. |
| `Get_file_extension()` | Returns the file extension. |
| `Check_parsed()` | Returns whether the file has been parsed. |

Note: JSON is read and written as a raw string, so you are responsible for marshalling/unmarshalling.

## Demo

A CLI demo is available in `cmd/demo`. Run it with:

```sh
go run cmd/demo/main.go
```
