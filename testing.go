package main

import (
	"bytes"
	"fmt"
	"io"
	"os"
)

func createFile(path string) error {
	// detect if file exists
	var _, err = os.Stat(path)

	// create file if not exists
	if os.IsNotExist(err) {
		var file, err = os.Create(path)
		if isError(err) {
			return err
		}
		defer file.Close()
	}
	return nil
}

func writeFile(path string, message string) error {
	// open file using READ & WRITE permission
	var file, err = os.OpenFile(path, os.O_RDWR, 0644)
	if isError(err) {
		return err
	}
	defer file.Close()

	// write some text line-by-line to file
	_, err = file.WriteString(message)
	if isError(err) {
		return err
	}

	// save changes
	err = file.Sync()
	if isError(err) {
		return err
	}
	return nil
}

func readFile(path string) (txt string, error error) {
	// re-open file
	var file, err = os.OpenFile(path, os.O_RDWR, 0644)
	if isError(err) {
		return
	}
	defer file.Close()

	// read file, line by line
	var text = make([]byte, 1024)
	for {
		_, err = file.Read(text)

		// break if finally arrived at end of file
		if err == io.EOF {
			break
		}

		// break if error occured
		if err != nil && err != io.EOF {
			isError(err)
			break
		}
	}
	text = bytes.Trim(text, "\x00")
	return string(text), nil
}

func deleteFile(path string) error {
	// delete file
	var err = os.Remove(path)
	if isError(err) {
		return err
	}
	return nil
}

func isError(err error) bool {
	if err != nil {
		fmt.Println(err.Error())
	}

	return (err != nil)
}
