package main

import (
  "path/filepath"
  "fmt"
)

func main() {
  paths, _ := filepath.Glob("/home/darko/[m|D-P]*")

  for _, path := range paths {
    fmt.Println(path)
  }
}
