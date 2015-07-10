package helpers

import (
  "fmt"
  "crypto/tls"
  "net/http"
  "io"
  "os"
  "path/filepath"
)

func CopyFile(source string, target string) error {
  sourceFile, err := os.Open(source)
  if err != nil {
    return err
  }
  defer sourceFile.Close()

  targetFile, err := os.Create(target)
  if err != nil {
    return err
  }
  defer targetFile.Close()

  if _, err := io.Copy(targetFile, sourceFile); err == nil {
    sourceInfo, err := os.Stat(source)
    if err != nil {
      err = os.Chmod(target, sourceInfo.Mode())
    }
  }

  return err
}

func CopyDir(source string, target string) error {
  // get properties of source dir
  sourceInfo, err := os.Stat(source)
  if err != nil {
    return err
  }

  // create target Directory
  if DoesDirExist(target) == false {
    if err := os.MkdirAll(target, sourceInfo.Mode()); err != nil {
      return err
    }
  }

  sourceDir, _ := os.Open(source)
  sourceChildren, err := sourceDir.Readdir(-1)

  for _, child := range sourceChildren {
    sourcePath := filepath.Join(source, child.Name())
    targetPath := filepath.Join(target, child.Name())

    if child.IsDir() {
      // create sub-directories - recursively
      if err := CopyDir(sourcePath, targetPath); err != nil {
        fmt.Println(err)
      }
    } else {
      // perform copy
      if err := CopyFile(sourcePath, targetPath); err != nil {
        fmt.Println(err)
      }
    }
  }
  return err
}


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
