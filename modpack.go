package main

import (
  "fmt"
  "os"
  "github.com/KSPtoMars/KSPtoMarsTools/libraries/argumenthandler"
  "github.com/KSPtoMars/KSPtoMarsTools/libraries/installsteps"
)

func main() {

  // Parse input Arguments. Check path and Flags
  inputArguments, err := argumenthandler.CheckArguments()
  if err != nil {
    fmt.Println(err)
    os.Exit(1)
  }

  // Setup all necessary paths
  relevantPaths := installsteps.SetupPaths(inputArguments)

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
  installsteps.DownloadNecessaryMods(inputArguments, &relevantPaths)

  // Unpack all zip files
  installsteps.UnpackAllZipFiles(&relevantPaths)

  // Remove outdated dependencies (especially if dependency will be installed anyway)
  installsteps.RemoveOldDependencies(&relevantPaths)

  // Move mods to GameData Folder
  backupPath := installsteps.CreateBackup(&relevantPaths)

  // Move mods to GameData Folder
  if err := installsteps.MoveMods(&relevantPaths); err != nil {
    fmt.Println ("There has been an error during copying!")
    fmt.Println (err)
    installsteps.RollBack(&relevantPaths, &backupPath)
  }

  // Clean up
  installsteps.CleanUp(&relevantPaths, &backupPath)

  // Remove unneeded Parts
  installsteps.RemoveUnneededParts(&relevantPaths)
}
