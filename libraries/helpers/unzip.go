package helpers

import (
  "io"
  "os"
  "archive/zip"
  "path/filepath"
)

func Unzip(zipFilePath, targetDir string) error {
  zipReader, err := zip.OpenReader(zipFilePath)
  if err != nil {
    return err
  }

  defer func() {
    if err := zipReader.Close(); err != nil {
      panic(err)
    }
  }()

  if err := os.MkdirAll(targetDir, 0755); err != nil {
    return err
  }


  // Closure to address file descriptors issue with all the deferred .Close() methods
  extractAndWriteFile := func(file *zip.File) error {
    fileReader, err := file.Open()
    if err != nil {
      return err
    }
    defer func() {
      if err := fileReader.Close(); err != nil {
        panic(err)
      }
    }()

    path := filepath.Join(targetDir, file.Name)

    if file.FileInfo().IsDir() {
      if err := os.MkdirAll(path, 0755); err != nil {
        return err
      }
    } else {
      if err := os.MkdirAll(filepath.Dir(path), 0755); err != nil {
        return err
      }
      file, err := os.OpenFile(path, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, file.Mode())
      if err != nil {
        return err
      }
      defer func() {
        if err := file.Close(); err != nil {
          panic(err)
        }
      }()

      _, err = io.Copy(file, fileReader)
      if err != nil {
        return err
      }
    }
    return nil
  }

  for _, file := range zipReader.File {
    err := extractAndWriteFile(file)
    if err != nil {
      return err
    }
  }

  return nil
}
