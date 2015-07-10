package helpers

import (
  "fmt"
  "crypto/tls"
  "net/http"
  "io"
  "os"
  "path/filepath"
)

func DoesDirExist(path string) bool {
  if _, err := os.Stat(path); os.IsNotExist(err) {
	  return false
  }
  return true
}

func DoesFileExist(path string) bool {
  return DoesDirExist(path)
}


func Download(A [][]string, targetDir string) {
  tr := &http.Transport{
    TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
  }
  client := &http.Client{Transport: tr}

  for i := 0; i < cap(A); i++ {
    uri := A[i][0]
    file := A[i][1]
    fmt.Println("[", i+1, " of ", cap(A), "]: ", file)
    out, _ := os.Create(filepath.Join(targetDir, file))
    defer out.Close()
    resp, err := client.Get(uri)
    if err != nil {
      fmt.Println(err)
      os.Exit (3)
    }
    defer resp.Body.Close()
    io.Copy(out, resp.Body)
  }
}
