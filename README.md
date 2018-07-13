# Log Parser

read log files in gorutine and than do whatere you need. 


```golang
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
```