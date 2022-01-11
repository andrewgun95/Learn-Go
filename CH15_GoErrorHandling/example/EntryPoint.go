package main

import (
	"io"
	"log"
	"os"
)

type ErrorWriter struct {
	w   io.Writer
	err error
}

func (e *ErrorWriter) Write(p []byte) {
	if e.err != nil {
		// Do nothing if error occurs
		return
	}

	_, e.err = e.w.Write(p)
}

func main() {
	file, err := os.Create("file.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	errWriter := &ErrorWriter{w: file}

	errWriter.Write([]byte("Hello"))
	errWriter.Write([]byte(", World!"))

	if errWriter.err != nil {
		log.Fatal(errWriter.err)
	}
}
