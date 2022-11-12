package main

import (
	"os"
)

func WriteFile(s string, filename string) {
	var data []byte = []byte(s)
	err := os.WriteFile(filename, data, 0777)
	if err != nil {
		panic(err)
	}
}

func ReadFile(filename string) string {
	var data []byte
	data, err := os.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	return string(data)
}

func AppendFile(s string, filename string) {
	var file string = ReadFile(filename)
	stream := file + s

	WriteFile(stream, filename)
}
