package argumenthandler
import (
  "fmt"
  "os"
  "flag"
  "path/filepath"
  "errors"
)

type Arguments struct {
  Path string
  DevFlag, CoreFlag, BeautyFlag, FullFlag bool
}

func CheckArguments() (*Arguments, error) {
  flag.Usage = func() {
    fmt.Printf("Usage: modScript -path=\"<Path to KSP folder>\" [-dev|-core|-beauty|-full]\n\n")
    flag.PrintDefaults()
  }

  inputArguments := new(Arguments)
  var argErr error = nil

  flag.StringVar(&(inputArguments.Path), "path", "./", "Path to KSP")

  flag.BoolVar(&(inputArguments.DevFlag), "dev", false, "Flag for core and dev mods")
  flag.BoolVar(&(inputArguments.CoreFlag), "core", false, "Flag for core mods")
  flag.BoolVar(&(inputArguments.BeautyFlag), "beauty", false, "Flag for core and beauty mods")
  flag.BoolVar(&(inputArguments.FullFlag), "full", false, "Flag for all mods")

  flag.Parse()

  var checkSum int = 0
  var errorEncountered = false

  if (inputArguments.DevFlag) { checkSum += 1 }
  if (inputArguments.CoreFlag) { checkSum += 1 }
  if (inputArguments.BeautyFlag) { checkSum += 1 }
  if (inputArguments.FullFlag) { checkSum += 1 }

  if (checkSum == 0) {
    inputArguments.DevFlag = true
  } else if (checkSum > 1) {
    argErr = errors.New("Please select only one installation type flag")
    errorEncountered = true
  }

  if _, err := os.Stat(filepath.Join(inputArguments.Path, "/GameData/Squad")); os.IsNotExist(err) {
    argErr = errors.New(inputArguments.Path + " doesn't seem to be a valid KSP installation")
    errorEncountered = true
  }

  if (errorEncountered) {
    flag.Usage()
    return nil, argErr
  }

  return inputArguments, argErr
}
