package installsteps

import (
  "fmt"
  "os"
  "io/ioutil"
  "path/filepath"
  "github.com/KSPtoMars/KSPtoMarsTools/libraries/argumenthandler"
  "github.com/KSPtoMars/KSPtoMarsTools/libraries/helpers"
  "github.com/KSPtoMars/KSPtoMarsTools/libraries/modsources"
)


type Paths struct {
  KspPath, GameDataPath, Ksp2mModsPath string
}

func SetupPaths(inputArguments *argumenthandler.Arguments) Paths {
  fmt.Println("\nSetting up installation paths")

  var relevantPaths Paths
  relevantPaths.KspPath = inputArguments.Path
  relevantPaths.GameDataPath = filepath.Join(inputArguments.Path, "/GameData")
  relevantPaths.Ksp2mModsPath = filepath.Join(inputArguments.Path, "/ksp2mMods")

  if helpers.DoesDirExist(relevantPaths.Ksp2mModsPath) {
    os.RemoveAll(relevantPaths.Ksp2mModsPath)
  }

  os.MkdirAll(relevantPaths.Ksp2mModsPath, 0755)

  return relevantPaths
}

func DownloadNecessaryMods(inputArguments *argumenthandler.Arguments, relevantPaths *Paths) error {
  fmt.Println("\nDownloading all mods. This will take a while.")

  // Core mods
  fmt.Println("\nDownloading Base Mods")
  if err := helpers.Download(modsources.Basemods, relevantPaths.Ksp2mModsPath); err != nil {
    return err
  }

  // Dev mods
  if (inputArguments.DevFlag || inputArguments.FullFlag) {
    fmt.Println("\nDownloading Dev Mods")
    if err := helpers.Download(modsources.Devmods, relevantPaths.Ksp2mModsPath); err != nil {
      return err
    }
  }

  // Beauty mods
  if (inputArguments.BeautyFlag || inputArguments.FullFlag) {
    // Remove low resolution RSS textures.
    if (helpers.DoesFileExist(filepath.Join(relevantPaths.Ksp2mModsPath, "2048.zip"))) {
      os.Remove(filepath.Join(relevantPaths.Ksp2mModsPath, "2048.zip"))
    }

    fmt.Println("\nDownloading Beauty Mods")
    if err := helpers.Download(modsources.Beautymods, relevantPaths.Ksp2mModsPath); err != nil {
      return err
    }
  }
  return nil
}

func UnpackAllZipFiles(relevantPaths *Paths) error {
  fmt.Println("\nUnpacking mods")

  files, err := ioutil.ReadDir(filepath.Join(relevantPaths.Ksp2mModsPath));
  if err != nil {
    return err
  }

  var maxMessageSize int = 0
  for i, f := range files {
    if (filepath.Ext(f.Name()) != ".zip") {
      continue
    }

    messageSize, _ := fmt.Printf("\rUnzipping [%d of %d]: %s", i+1, len(files), f.Name())
    if (messageSize > maxMessageSize) {
      maxMessageSize = messageSize
    } else {
      for i := messageSize; i < maxMessageSize; i++ {
        fmt.Printf(" ")
      }
    }
    fmt.Printf("\r")

    var fileToExtract = filepath.Join(relevantPaths.Ksp2mModsPath, f.Name())
    var pathToExtractTo = filepath.Join(relevantPaths.Ksp2mModsPath, f.Name()[0:len(f.Name()) - 4])
    if err := helpers.Unzip(fileToExtract, pathToExtractTo); err != nil {
      return err
    }
  }

  fmt.Println("")
  return nil
}

func RemoveOldDependencies(relevantPaths *Paths) {
  fmt.Println("\nRemoving outdated dependencies")

  var foldersToDelete = []string {
    "/UKS/GameData/CommunityResourcePack",
    "/UKS/GameData/Firespitter",
    "/Advanced_Jet_Engine/GameData/SolverEngines",
    "/B9ProcParts/GameData/CrossFeedEnabler",
    "/FAR/GameData/ModularFlightIntegrator",
    "/FASA/GameData/JSI",
    "/RealFuels/CommunityResourcePack",
    "/RealFuels/SolverEngines",
    "/RealHeat/ModularFlightIntegrator",
    "/UniversalStorage/CommunityResourcePack",
  }

  for _, folder := range foldersToDelete {
    if err := os.RemoveAll(filepath.Join(relevantPaths.Ksp2mModsPath, folder)); err != nil {
      fmt.Println(err)
    }
  }
}

func CreateBackup(relevantPaths *Paths) string {
  fmt.Println("\nCreating backup of GameData folder")

  var backupPath = filepath.Join(relevantPaths.KspPath, "/GameData_Backup_By_KSPtoMars_Modscript")
  if err := os.Rename(relevantPaths.GameDataPath, backupPath); err != nil {
    fmt.Println(err)
  }
  os.MkdirAll(filepath.Join(relevantPaths.GameDataPath, "/Squad"), 0775)
  helpers.CopyDir(filepath.Join(backupPath, "/Squad"), filepath.Join(relevantPaths.GameDataPath, "/Squad"))

  return backupPath
}

func RollBack(relevantPaths *Paths, backupPath *string) {
  fmt.Println("\nRolling back")

  os.RemoveAll(filepath.Join(relevantPaths.GameDataPath))
  if err := os.Rename(*backupPath, relevantPaths.GameDataPath); err != nil {
    fmt.Println(err)
  }
}


func MoveMods(relevantPaths *Paths) error {
  fmt.Println("\nMoving mods to GameData folder")

  // Generic move for mods
  files, err := ioutil.ReadDir(relevantPaths.Ksp2mModsPath);
  if err != nil {
    return err
  }
  for _, f := range files {
    if f.IsDir() == false {
      continue
    }

    var pathToMod = filepath.Join(relevantPaths.Ksp2mModsPath, f.Name())
    var pathToModGameData = filepath.Join(pathToMod, "/GameData")
    if (helpers.DoesDirExist(pathToModGameData) == false) {
      continue
    }

    if err := helpers.CopyDir(pathToModGameData, relevantPaths.GameDataPath); err != nil {
      return err
    }
  }

  // Custom move for mods
  var customFolders = []string {
    // core install
    "/CrossFeedEnabler",
    "/RealFuels",
    "/RealHeat",
    "/RealSolarSystem",
    "/Toolbar/Toolbar-1.7.9/GameData",
    "/ksp-avc",
    "/KWRocketry/KW Release Package v2.7 (Open this, don't extract it)/GameData",
    "/UniversalStorage",
    "/StockBugFixModules",
    "/AIES_Aerospace151",
    "/HullcaMove-ItemDS",
    "/JDiminishingRTG/JDiminishingRTG_v1_3a/GameData",
    "/NebulaDecals",
    "/ModularFlightIntegrator",
    // dev install
    "/mechjeb2",
    "/VesselViewer",
    "/FShangarExtender",
    "/PartWizard",
    "/RCSbuildAid",
    "/StripSymmetry/Gamedata",
    "/EditorExtensions",
    "/KerbalEngineer",
    // beauty install
    "/hotrocket",
    "/DistantObject/Alternate Planet Color Configs/Real Solar System (metaphor's PlanetFactory config)",
    "/EngineLighting/EngineLight/GameData",
    "/ImprovedChaseCam",
    "/PlanetShine/Alternate Colors/Real Solar System",
    "/SpaceY-Lifters",
  }
  if err := customMoveMods(relevantPaths, customFolders); err != nil {
    return err
  }

  // Fixing Configs
  if err := helpers.CopyDir(filepath.Join(relevantPaths.Ksp2mModsPath, "RealismOverhaul/GameData"), relevantPaths.GameDataPath); err != nil {
    return err
  }

  var configFixes = [][]string {
    []string {filepath.Join(relevantPaths.Ksp2mModsPath, "/RealismOverhaul/GameData/RealismOverhaul/RemoteTech_Settings.cfg"), relevantPaths.GameDataPath + "/RemoteTech/RemoteTech_Settings.cfg"},
    []string {filepath.Join(relevantPaths.Ksp2mModsPath, "/TextureReplacer/Extras/MM_ReflectionPluginWrapper.cfg"), relevantPaths.GameDataPath + "/MM_ReflectionPluginWrapper.cfg"},
    []string {filepath.Join(relevantPaths.Ksp2mModsPath, "/StockPlusController.cfg"), relevantPaths.GameDataPath + "/StockPlusController.cfg"},
    []string {filepath.Join(relevantPaths.Ksp2mModsPath, "/AIES_Node_Patch.cfg/AIES_Node_Patch.cfg"), relevantPaths.GameDataPath + "/AIES_Node_Patch.cfg"},
    []string {filepath.Join(relevantPaths.Ksp2mModsPath, "/Firespitter/Firespitter/Plugins/Firespitter.dll"), relevantPaths.GameDataPath + "/Firespitter/Plugins/Firespitter.dll"},
    []string {filepath.Join(relevantPaths.Ksp2mModsPath, "/Firespitter/Firespitter/Plugins/FSfuelSwitchTweakscale.cfg"), relevantPaths.GameDataPath + "/Firespitter/Plugins/FSfuelSwitchTweakscale.cfg"},
    []string {filepath.Join(relevantPaths.Ksp2mModsPath, "/RO_TaurusHCV.cfg"), relevantPaths.GameDataPath + "/RealismOverhaul/REWORK/RO_TaurusHCV.cfg"},
  }
  for _, fix := range configFixes {
    if err := helpers.CopyFile(fix[0], fix[1]); err != nil {
      return err
    }
  }

  return nil
}

func customMoveMods(relevantPaths *Paths, customPaths []string) error {
  for _, folder := range customPaths {
    var pathToModGameData = filepath.Join(relevantPaths.Ksp2mModsPath, folder)
    if helpers.DoesDirExist(pathToModGameData) == false {
      continue
    }
    if err := helpers.CopyDir(pathToModGameData, relevantPaths.GameDataPath); err != nil {
      return err
    }
  }

  return nil
}

func CleanUp(relevantPaths *Paths, backupPath *string) {
  fmt.Println("\nCleaning up")

  os.MkdirAll(filepath.Join(relevantPaths.GameDataPath, "/licensesAndReadmes"), 0755)

  files, err := ioutil.ReadDir(relevantPaths.GameDataPath);
  if err != nil {
    fmt.Println(err)
  }
  for _, f := range files {
    if filepath.Ext(f.Name()) == ".txt" ||
       filepath.Ext(f.Name()) == ".md" ||
       filepath.Ext(f.Name()) == ".pdf" ||
       filepath.Ext(f.Name()) == ".htm" ||
       f.Name() == "License" {
       if err := os.Rename(relevantPaths.GameDataPath + "/" + f.Name(), relevantPaths.GameDataPath + "/licensesAndReadmes" + "/" + f.Name()); err != nil {
         fmt.Println(err)
       }
    }
  }

  os.RemoveAll(relevantPaths.Ksp2mModsPath)
  os.RemoveAll(*backupPath)
}

func RemoveUnneededParts(relevantPaths *Paths) {
  fmt.Println("\nRemoving unneccessary files")

  deleteOldModuleManagers(relevantPaths)

  deleteListOfFiles(relevantPaths, filesToDelete)
}

func deleteOldModuleManagers(relevantPaths *Paths) error {
  files, err := filepath.Glob(filepath.Join(relevantPaths.GameDataPath, "/ModuleManager*"))
  if err != nil {
    return err
  }

  for j, file := range files {
    if j == len(files) - 1 {
      break
    }
    if err := os.Remove(file); err != nil {
      return err
    }
  }

  return nil
}

func deleteListOfFiles(relevantPaths *Paths, filesToDelete []string) error {
  for i, pattern := range filesToDelete {
    files, err := filepath.Glob(filepath.Join(relevantPaths.GameDataPath, pattern))
    if err != nil {
      fmt.Println("Encountered Error! i =", i,", pattern =", pattern)
      fmt.Println(err)
      return err
    }

    for j, file := range files {
      if err := os.RemoveAll(file); err != nil {
        fmt.Println("Encountered Error! i =", i,", j =", j,", pattern =", pattern,", file = ", file)
        return err
      }
    }
  }

  return nil
}
