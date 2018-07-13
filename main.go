package main

import (
	"fmt"
)

func main() {
	logReader := LogRader{}
	go logReader.Read("q.log")

	logReader2 := LogRader{}
	go logReader2.Read("qq.log")

	fmt.Printf("%s\n", "main")

	logReader.Stop()
	logReader2.Stop()
}
