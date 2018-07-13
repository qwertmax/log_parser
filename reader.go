package main

import (
	"fmt"
	"os"
	"time"
)

type LogRader struct {
	File *os.File
	Line chan []byte
	Size int64
	Quit chan bool
}

func (reader *LogRader) Read(path string) {
	inFile, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer inFile.Close()

	reader.checkSize(inFile)
	reader.Line = make(chan []byte, 10)
	reader.Quit = make(chan bool)

	go func() {
		for {
			select {
			case line := <-reader.Line:
				fmt.Printf("%s", line)
			case <-time.After(500 * time.Millisecond):
				reader.checkSize(inFile)
				line := make([]byte, reader.Size)
				length, _ := inFile.Read(line)

				if length == 0 {
					continue
				}

				reader.Line <- line
			}
		}
	}()

	<-reader.Quit
}

func (reader *LogRader) Stop() {
	reader.Quit <- true
}

func (reader *LogRader) checkSize(f *os.File) {
	stat, err := f.Stat()
	reader.Size = stat.Size()
	if err != nil {
		panic(err)
	}

	return
}

func (reader *LogRader) ReadLine() (line []byte, length int) {
	reader.checkSize(reader.File)
	line = make([]byte, reader.Size)
	length, _ = reader.File.Read(line)
	return
}
