package helpers

import (
  "fmt"
  "net/http"
  "io"
  "os"
)

func DoesDirExist(path string) bool {
  if _, err := os.Stat(path); os.IsNotExist(err) {
	  return false
  }
  return true
}

func Download(A [][]string) {
  for i := 0; i < cap(A); i++ {
    uri := A[i][0]
    file := A[i][1]
    fmt.Println("[", i+1, " of ", cap(A), "]: ", file)
    out, _ := os.Create(file)
    defer out.Close()
    resp, _ := http.Get(uri)
    defer resp.Body.Close()
    n, _ := io.Copy(out, resp.Body)
    fmt.Println (n)
  }
}
