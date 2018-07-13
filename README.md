# Log Parser

read log files in gorutine and than do whatere you need. 


```golang
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
```

main.go

```golang
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
```