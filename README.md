# gofio

File handling with go

- A library to simplify file handling in Golang

## Working

- File Handler initialization
  - Either provide file path for a new file
  - Or a file path to an existing file to work with
  - Extension of the file will be inferred from the file path
- After initialization, Gofio object can be used to perform following I/O operations
  1. Create the file
  2. Read the file contents (returns the string i.e. JSON will also be returned as string hence requiring the consumer to do the Unmarshalling)
  3. Append to file
  4. Delete File
  5. Re-initialize filehandler (to change the file path contained by the Gofio object)

## Demo

- I have built a cli app for the demo
- Run using

```go
go run cmd/demo/main.go
```
