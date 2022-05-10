pipe
---

Small utility for getting pipes from stdin

# exports

```go
// get data from stdin
func Get() []byte
```

# usage

```go
package main

import "github.com/kendfss/pipe"


func main() {
    if data := pipe.Get(); data != nil {
        processInput(data) // this exercise is left to the reader
    } 
}
```
