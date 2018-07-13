package main

import (
	"fmt"
)

func main() {
	logReader := LogRader{}
	logReader.ParseFunc = func(line []byte) {
		fmt.Printf("f1 => %s", line)
	}
	go logReader.Read("q.log")

	logReader2 := LogRader{}
	logReader2.ParseFunc = func(line []byte) {
		fmt.Printf("f2 => %s", line)
	}
	go logReader2.Read("qq.log")

	fmt.Printf("%s\n", "main")

	logReader.Stop()
	// logReader2.Stop()
}
