package main
import (
  "fmt"
  "os"
  "io/ioutil"
  "path/filepath"
  "github.com/Orkeren/KSPtoMarsTools/libraries/argumenthandler"
  "github.com/Orkeren/KSPtoMarsTools/libraries/helpers"
  "github.com/Orkeren/KSPtoMarsTools/libraries/modsources"
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

  // Remove outdated dependencies (especially if dependency will be installed anyway)
  removeOldDependencies(&relevantPaths)

  // Move mods to GameData Folder
  moveMods(&relevantPaths)

  // Clean up
  cleanUp(&relevantPaths)

  // Remove unneeded Parts
  removeUnneededParts(&relevantPaths)
}

func removeUnneededParts(relevantPaths *paths) {
}

func moveMods(relevantPaths *paths) {
  // Generic move for mods
  files, err := ioutil.ReadDir(relevantPaths.ksp2mModsPath);
  if err != nil {
    fmt.Println(err)
  }
  for _, f := range files {
    if f.IsDir() == false {
      continue
    }

    var pathToMod = filepath.Join(relevantPaths.ksp2mModsPath, f.Name())
    var pathToModGameData = filepath.Join(pathToMod, "/GameData")
    if (helpers.DoesDirExist(pathToModGameData) == false) {
      continue
    }

    if err := helpers.CopyDir(pathToModGameData, relevantPaths.gameDataPath); err != nil {
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
    "/NebulaDecals/NEBULA",
  }
  for _, folder := range coreCustomFolder {
    var pathToModGameData = filepath.Join(relevantPaths.ksp2mModsPath, folder)
    if err := helpers.CopyDir(pathToModGameData, relevantPaths.gameDataPath); err != nil {
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
    var pathToModGameData = filepath.Join(relevantPaths.ksp2mModsPath, folder)
    if err := helpers.CopyDir(pathToModGameData, relevantPaths.gameDataPath); err != nil {
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
    var pathToModGameData = filepath.Join(relevantPaths.ksp2mModsPath, folder)
    if err := helpers.CopyDir(pathToModGameData, relevantPaths.gameDataPath); err != nil {
      fmt.Println(err)
    }
  }

  // Fixing Configs
  if err := helpers.CopyDir(filepath.Join(relevantPaths.ksp2mModsPath, "RealismOverhaul/GameData"), relevantPaths.gameDataPath); err != nil {
    fmt.Println(err)
  }
  if err := helpers.CopyFile(relevantPaths.ksp2mModsPath + "/RealismOverhaul/GameData/RealismOverhaul/RemoteTech_Settings.cfg", relevantPaths.gameDataPath + "/RemoteTech/RemoteTech_Settings.cfg"); err != nil {
    fmt.Println(err)
  }
  if err := helpers.CopyFile(relevantPaths.ksp2mModsPath + "/TextureReplacer/Extras/MM_ReflectionPluginWrapper.cfg", relevantPaths.gameDataPath); err != nil {
    fmt.Println(err)
  }
  if err := helpers.CopyFile(relevantPaths.ksp2mModsPath + "/StockPlusController.cfg", relevantPaths.gameDataPath); err != nil {
    fmt.Println(err)
  }
  if err := helpers.CopyFile(relevantPaths.ksp2mModsPath + "/StockPlusController.cfg", relevantPaths.gameDataPath); err != nil {
    fmt.Println(err)
  }
}

func cleanUp(relevantPaths *paths) {
  os.RemoveAll(relevantPaths.ksp2mModsPath)
  os.MkdirAll(relevantPaths.gameDataPath + "licensesAndReadmes", 0755)

  files, err := ioutil.ReadDir(relevantPaths.gameDataPath);
  if err != nil {
    fmt.Println(err)
  }
  for _, f := range files {
    if filepath.Ext(f.Name()) == ".txt" ||
       filepath.Ext(f.Name()) == ".md" ||
       filepath.Ext(f.Name()) == ".pdf" ||
       filepath.Ext(f.Name()) == ".htm" {
       if err := os.Rename(relevantPaths.gameDataPath + f.Name(), relevantPaths.gameDataPath + "/licensesAndReadmes" + f.Name()); err != nil {
         fmt.Println(err)
       }
    }
  }
}

func removeOldDependencies(relevantPaths *paths) {
  var foldersToDelete = []string {
    "/UKS/GameData/CommunityResourcePack",
    "/Advanced_Jet_Engine/GameData/SolverEngines",
    "/B9ProcParts/GameData/CrossFeedEnabler",
    "/DeadlyReentry/ModularFlightIntegrator",
    "/FAR/GameData/ModularFlightIntegrator",
    "/FASA/GameData/JSI",
    "/RealFuels/CommunityResourcePack",
    "/RealFuels/SolverEngines",
    "/RealHeat/ModularFlightIntegrator",
    "/UniversalStorage/CommunityResourcePack",
  }

  for _, folder := range foldersToDelete {
    if err := os.RemoveAll(filepath.Join(relevantPaths.ksp2mModsPath, folder)); err != nil {
      fmt.Println(err)
    }
  }
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
