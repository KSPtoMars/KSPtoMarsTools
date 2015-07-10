package main
import (
  "fmt"
  "os"
  "path/filepath"
  "argumenthandler"
  "helpers"
  "modsources"
)

type paths struct {
  kspPath, gameDataPath, ksp2mModsPath string
}

func main() {
  var inputArguments *argumenthandler.Arguments = argumenthandler.CheckArguments()
  var relevantPaths paths
  if inputArguments == nil {
    os.Exit(1)
  } else {
    relevantPaths.kspPath = inputArguments.Path
    relevantPaths.gameDataPath = filepath.Dir(filepath.Join(inputArguments.Path, "GameData"))
    relevantPaths.ksp2mModsPath = filepath.Dir(filepath.Join(inputArguments.Path, "ksp2mMods"))
  }

  if helpers.DoesDirExist(relevantPaths.ksp2mModsPath) {
    os.RemoveAll(relevantPaths.ksp2mModsPath)
  }
  os.MkdirAll(relevantPaths.ksp2mModsPath, 0755)

  if (inputArguments.BeautyFlag) {
    fmt.Println("Preparing beauty install.")
  } else if (inputArguments.CoreFlag){
    fmt.Println("Preparing base install.")
  } else if (inputArguments.FullFlag){
    fmt.Println("Preparing full install.")
  } else {
    fmt.Println("Preparing developer install.")
  }

  fmt.Println("Downloading all mods. This will take a while.")

  fmt.Println("Downloading Base Mods")
  helpers.Download(modsources.Basemods)

  if (inputArguments.DevFlag || inputArguments.FullFlag) {
    fmt.Println("Downloading Dev Mods")
    helpers.Download(modsources.Devmods)
  }

  if (inputArguments.BeautyFlag || inputArguments.FullFlag) {
    // Remove low resolution RSS textures.
    os.Remove(filepath.Join(relevantPaths.ksp2mModsPath, "2048.zip"))

    fmt.Println("Downloading Beauty Mods")
    helpers.Download(modsources.Beautymods)
  }
}
