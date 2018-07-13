package main

import (
	"fmt"
)

func main() {
	logReader := NewLogRader("q.log", func(line []byte) {
		fmt.Printf("f1 => %s", line)
	})

	logReader2 := NewLogRader("qq.log", func(line []byte) {
		fmt.Printf("f2 => %s", line)
	})

	go logReader.Read()
	go logReader2.Read()

	fmt.Printf("%s\n", "main")

	logReader.Stop()
	logReader2.Stop()
}
