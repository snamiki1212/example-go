package main

import "os"

// Write to output file
func write(path, txt string) error {
	if txt == "" { // skip to create file because of empty
		return nil
	}
	file, err := os.Create(path)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	_, err = file.WriteString(txt)
	return err
}
