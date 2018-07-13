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