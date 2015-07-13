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

func DownloadNecessaryMods(inputArguments *argumenthandler.Arguments, relevantPaths *Paths) {
  fmt.Println("\nDownloading all mods. This will take a while.")

  // Core mods
  fmt.Println("\nDownloading Base Mods")
  helpers.Download(modsources.Basemods, relevantPaths.Ksp2mModsPath)

  // Dev mods
  if (inputArguments.DevFlag || inputArguments.FullFlag) {
    fmt.Println("\nDownloading Dev Mods")
    helpers.Download(modsources.Devmods, relevantPaths.Ksp2mModsPath)
  }

  // Beauty mods
  if (inputArguments.BeautyFlag || inputArguments.FullFlag) {
    // Remove low resolution RSS textures.
    if (helpers.DoesFileExist(filepath.Join(relevantPaths.Ksp2mModsPath, "2048.zip"))) {
      os.Remove(filepath.Join(relevantPaths.Ksp2mModsPath, "2048.zip"))
    }

    fmt.Println("\nDownloading Beauty Mods")
    helpers.Download(modsources.Beautymods, relevantPaths.Ksp2mModsPath)
  }
}

func UnpackAllZipFiles(relevantPaths *Paths) {
  fmt.Println("\nUnpacking mods")

  files, err := ioutil.ReadDir(filepath.Join(relevantPaths.Ksp2mModsPath));
  if err != nil {
    fmt.Println(err)
  }
  for i, f := range files {
    if (filepath.Ext(f.Name()) != ".zip") {
      continue
    }
    fmt.Println("Unzipping [",i+1, "of", len(files),"]: " + f.Name())

    var fileToExtract = filepath.Join(relevantPaths.Ksp2mModsPath, f.Name())
    var pathToExtractTo = filepath.Join(relevantPaths.Ksp2mModsPath, f.Name()[0:len(f.Name()) - 4])
    if err := helpers.Unzip(fileToExtract, pathToExtractTo); err != nil {
      fmt.Println("Error while unzipping " + f.Name())
      fmt.Println(err)
    }
  }
}

func RemoveOldDependencies(relevantPaths *Paths) {
  fmt.Println("\nRemoving outdated dependencies")

  var foldersToDelete = []string {
    "/UKS/GameData/CommunityResourcePack",
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

  var backupPath = filepath.Join(relevantPaths.GameDataPath, "/GameData_Backup_By_KSPtoMars_Modscript")
  os.Rename(relevantPaths.GameDataPath, backupPath)

  os.MkdirAll(filepath.Join(relevantPaths.GameDataPath, "/Squad"), 0775)
  helpers.CopyDir(filepath.Join(backupPath, "/Squad"), filepath.Join(relevantPaths.GameDataPath, "/Squad"))

  return backupPath
}

func RollBack(relevantPaths *Paths, backupPath *string) {
  fmt.Println("\nRolling back")

  os.RemoveAll(filepath.Join(relevantPaths.GameDataPath))
  os.Rename(*backupPath, relevantPaths.GameDataPath)
}


func MoveMods(relevantPaths *Paths) error {
  fmt.Println("\nMoving mods do GameData folder")

  // Generic move for mods
  files, err := ioutil.ReadDir(relevantPaths.Ksp2mModsPath);
  if err != nil {
    fmt.Println(err)
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
      fmt.Println(err)
    }
  }

  // Custom move for mods
  // core install
  var coreCustomFolder = []string{
    "/CrossFeedEnabler",
    "/DeadlyReentry",
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
  }
  for _, folder := range coreCustomFolder {
    var pathToModGameData = filepath.Join(relevantPaths.Ksp2mModsPath, folder)
    if err := helpers.CopyDir(pathToModGameData, relevantPaths.GameDataPath); err != nil {
      fmt.Println(err)
    }
  }

  // dev install
  var devCustomFolder = []string {
    "/mechjeb2",
    "/VesselViewer",
    "/FShangarExtender",
    "/PartWizard",
    "/RCSbuildAid",
    "/StripSymmetry/Gamedata",
    "/EditorExtensions",
    "/KerbalEngineer",
  }
  for _, folder := range devCustomFolder {
    var pathToModGameData = filepath.Join(relevantPaths.Ksp2mModsPath, folder)
    if err := helpers.CopyDir(pathToModGameData, relevantPaths.GameDataPath); err != nil {
      fmt.Println(err)
    }
  }

  // beauty install
  var beautyCustomFolder = []string {
    "/hotrocket",
    "/DistantObject/Alternate Planet Color Configs/Real Solar System (metaphor's PlanetFactory config)",
    "/EngineLighting/EngineLight/GameData",
    "/ImprovedChaseCam",
    "/PlanetShine/Alternate Colors/Real Solar System",
    "/RoverWheelSounds",
  }
  for _, folder := range beautyCustomFolder {
    var pathToModGameData = filepath.Join(relevantPaths.Ksp2mModsPath, folder)
    if err := helpers.CopyDir(pathToModGameData, relevantPaths.GameDataPath); err != nil {
      fmt.Println(err)
    }
  }

  // Fixing Configs
  var configFixes = [][]string {
    []string {filepath.Join(relevantPaths.Ksp2mModsPath, "RealismOverhaul/GameData"), relevantPaths.GameDataPath},
    []string {filepath.Join(relevantPaths.Ksp2mModsPath, "/RealismOverhaul/GameData/RealismOverhaul/RemoteTech_Settings.cfg"), relevantPaths.GameDataPath + "/RemoteTech/RemoteTech_Settings.cfg"},
    []string {filepath.Join(relevantPaths.Ksp2mModsPath, "/TextureReplacer/Extras/MM_ReflectionPluginWrapper.cfg"), relevantPaths.GameDataPath + "/MM_ReflectionPluginWrapper.cfg"},
    []string {filepath.Join(relevantPaths.Ksp2mModsPath, "/StockPlusController.cfg"), relevantPaths.GameDataPath + "/StockPlusController.cfg"},
    []string {filepath.Join(relevantPaths.Ksp2mModsPath, "/AIES_Node_Patch.cfg/AIES_Node_Patch.cfg"), relevantPaths.GameDataPath + "/AIES_Node_Patch.cfg"},
  }
  for _, fix := range configFixes {
    if err := helpers.CopyDir(fix[0], fix[1]); err != nil {
      fmt.Println(err)
    }
  }

  return nil
}

func CleanUp(relevantPaths *Paths, backupPath *string) {
  fmt.Println("\nCleaning up")

  os.RemoveAll(relevantPaths.Ksp2mModsPath)

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
    fmt.Println("Deleting pattern", i+1, "of", len(filesToDelete))
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
