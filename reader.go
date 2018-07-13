package main

import (
	"os"
	"time"
)

type LogRader struct {
	File      *os.File
	Line      chan []byte
	Size      int64
	Quit      chan bool
	ParseFunc func([]byte)
}

func NewLogRader(path string, parseFunc func([]byte)) LogRader {
	file, err := os.Open(path)
	if err != nil {
		panic(err)
	}

	return LogRader{
		File:      file,
		ParseFunc: parseFunc,
	}
}

func (reader *LogRader) Read() {
	reader.checkSize()
	reader.Line = make(chan []byte, 10)
	reader.Quit = make(chan bool)

	go func() {
		for {
			select {

			case line := <-reader.Line:
				reader.ParseFunc(line)

			case <-time.After(500 * time.Millisecond):
				reader.checkSize()
				line := make([]byte, reader.Size)
				length, _ := reader.File.Read(line)

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
	// reader.File.Close()
}

func (reader *LogRader) checkSize() {
	stat, err := reader.File.Stat()
	reader.Size = stat.Size()
	if err != nil {
		panic(err)
	}

	return
}

func (reader *LogRader) ReadLine() (line []byte, length int) {
	reader.checkSize()
	line = make([]byte, reader.Size)
	length, _ = reader.File.Read(line)
	return
}
