package main
import (
  "fmt"
  "os"
  "io/ioutil"
  "path/filepath"
  "argumenthandler"
  "helpers"
  "modsources"
)

type paths struct {
  kspPath, gameDataPath, ksp2mModsPath string
}

func main() {

  // Parse input Arguments. Check path and Flags
  inputArguments, err := argumenthandler.CheckArguments()
  if err != nil {
    fmt.Println(err)
    os.Exit(1)
  }

  // Setup all necessary paths
  relevantPaths := setupPaths(inputArguments)

  // Start with mod installation
  if (inputArguments.BeautyFlag) {
    fmt.Println("Preparing beauty install.")
  } else if (inputArguments.CoreFlag){
    fmt.Println("Preparing base install.")
  } else if (inputArguments.FullFlag){
    fmt.Println("Preparing full install.")
  } else {
    fmt.Println("Preparing developer install.")
  }

  // Download necessary mods
  downloadNecessaryMods(inputArguments, &relevantPaths)

  // Unpack all zip files
  unpackAllZipFiles(&relevantPaths)
}

func setupPaths(inputArguments *argumenthandler.Arguments) paths {
  var relevantPaths paths
  relevantPaths.kspPath = inputArguments.Path
  relevantPaths.gameDataPath = filepath.Join(inputArguments.Path, "/GameData")
  relevantPaths.ksp2mModsPath = filepath.Join(inputArguments.Path, "/ksp2mMods")

  if helpers.DoesDirExist(relevantPaths.ksp2mModsPath) {
    os.RemoveAll(relevantPaths.ksp2mModsPath)
  }

  os.MkdirAll(relevantPaths.ksp2mModsPath, 0755)

  return relevantPaths
}

func downloadNecessaryMods(inputArguments *argumenthandler.Arguments, relevantPaths *paths) {
  fmt.Println("Downloading all mods. This will take a while.")

  // Core mods
  fmt.Println("Downloading Base Mods")
  helpers.Download(modsources.Basemods, relevantPaths.ksp2mModsPath)

  // Dev mods
  if (inputArguments.DevFlag || inputArguments.FullFlag) {
    fmt.Println("Downloading Dev Mods")
    helpers.Download(modsources.Devmods, relevantPaths.ksp2mModsPath)
  }

  // Beauty mods
  if (inputArguments.BeautyFlag || inputArguments.FullFlag) {
    // Remove low resolution RSS textures.
    if (helpers.DoesFileExist(filepath.Join(relevantPaths.ksp2mModsPath, "2048.zip"))) {
      os.Remove(filepath.Join(relevantPaths.ksp2mModsPath, "2048.zip"))
    }

    fmt.Println("Downloading Beauty Mods")
    helpers.Download(modsources.Beautymods, relevantPaths.ksp2mModsPath)
  }
}

func unpackAllZipFiles(relevantPaths *paths) {
  files, err := ioutil.ReadDir(filepath.Join(relevantPaths.ksp2mModsPath));
  if err != nil {
    fmt.Println(err)
  }
  for _, f := range files {
    if (filepath.Ext(f.Name()) != ".zip") {
      continue
    }

    var fileToExtract = filepath.Join(relevantPaths.ksp2mModsPath, f.Name())
    var pathToExtractTo = filepath.Join(relevantPaths.ksp2mModsPath, f.Name()[0:len(f.Name()) - 4])
    if err := helpers.Unzip(fileToExtract, pathToExtractTo); err != nil {
      fmt.Println("Error while unzipping " + f.Name())
      fmt.Println(err)
    }
  }
}
