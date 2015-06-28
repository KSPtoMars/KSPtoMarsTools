#
# KSPtoMars Windows Modpack v1.5.0
# Written by Sven Frenzel (sven@frenzel.dk) with some contributions by Darko Pilav (darko.pilav@gmail.com)
#
# The MIT License (MIT)
# 
# Copyright (c) 2015 Sven Frenzel (sven@frenzel.dk)
#
# Permission is hereby granted, free of charge, to any person obtaining a copy
# of this software and associated documentation files (the "Software"), to deal
# in the Software without restriction, including without limitation the rights
# to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
# copies of the Software, and to permit persons to whom the Software is
# furnished to do so, subject to the following conditions:
#  
# The above copyright notice and this permission notice shall be included in
# all copies or substantial portions of the Software.
# 
# THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
# IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
# FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
# AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
# LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
# OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
# THE SOFTWARE.
#
#
#
# ChangeLog
#
# v1.5.2 (2015-06-28)
#  - started work on removing further unnecessary parts. So far removed:
#    * EngineIgnitor Parts 
#    * DMagic Universal Storage Parts
#
# v1.5.1 (2015-06-28)
#  - updated all mods to their newest versions  
#

[CmdletBinding()]
Param(
  [Parameter(Mandatory=$True,Position=1)]
  [string]$k,
  [switch]$b,
  [switch]$c,
  [switch]$f
)

#Definition of function for easy unzipping later on
function unzip($file) {
  Add-Type -assembly "system.io.compression.filesystem"
  [io.compression.zipfile]::ExtractToDirectory($file, $pwd)
}

Write-Output "`r`nThis is v1.5.0 of the ksp2mars modpack script for windows.`r`n`r`n"

$startingPath = $PWD

if (Test-Path $k/GameData) {
  Set-Location $k
  Move-Item GameData/Squad Squad_bak
  Remove-Item -Recurse -Force GameData
  new-item -itemtype directory GameData > $null
  Move-Item Squad_bak GameData/Squad
}else{
  Write-Output "The specified path does not seem to contain a valid install of KSP."
  exit
}

# Create folders
if (Test-Path ksp2m_mods){
  Remove-Item -Recurse -Force ksp2m_mods
}

new-item -itemtype directory ksp2m_mods > $null
Set-Location ksp2m_mods

If ($b){
  Write-Output "Preparing beauty install."
}ElseIf ($c){
  Write-Output "Preparing base install."
}ElseIf ($f){
  Write-Output "Preparing full install."
}Else{
  Write-Output "Preparing developer install."
}

# Download base mods!
Write-Output "`r`nDownloading all mods. This will take a while."

$baseModPack = @(
  @("http://github.com/camlost2/AJE/releases/download/2.2.1/Advanced_Jet_Engine-2.2.1.zip", "Advanced_Jet_Engine.zip"),
  @("http://github.com/ferram4/BetterBuoyancy/releases/download/v1.3/BetterBuoyancy_v1.3.zip", "BetterBuoyancy.zip"),
  @("http://github.com/BobPalmer/CommunityResourcePack/releases/download/0.4.2/CRP_0.4.2.zip", "CRP.zip"),
  @("http://github.com/codepoetpbowden/ConnectedLivingSpace/releases/download/1.1.3.1/Connected_Living_Space-1.1.3.1.zip", "Connected_Living_Space.zip"),
  @("http://github.com/NathanKell/CrossFeedEnabler/releases/download/v3.3/CrossFeedEnabler_v3.3.zip", "CrossFeedEnabler.zip"),
  @("http://github.com/Starwaster/DeadlyReentry/releases/download/v7.1.0/DeadlyReentry_7.1.0_The_Melificent_Edition.zip", "DeadlyReentry.zip"),
  @("http://kerbalstuff.com/mod/27/FASA/download/5.35", "FASA.zip"),
  @("http://github.com/ferram4/Ferram-Aerospace-Research/releases/download/v0.15_3_1_Garabedian/FAR_0_15_3_1_Garabedian.zip", "FAR.zip"),
  @("http://github.com/ferram4/Kerbal-Joint-Reinforcement/releases/download/v3.1.4/KerbalJointReinforcement_v3.1.4.zip", "KerbalJointReinforcement.zip"),
  @("http://kerbal.curseforge.com/ksp-mods/220462-ksp-avc-add-on-version-checker/files/2216818/download", "ksp-avc.zip"),
  @("http://beta.kerbalstuff.com/mod/67/KW%20Rocketry/download/2.7", "KWRocketry.zip"),
  @("http://github.com/KSP-RO/SolverEngines/releases/download/v1.4/SolverEngines_v1.4.zip", "SolverEngines.zip"),
  @("http://kerbalstuff.com/mod/26/NovaPunch/download/2.09", "NovaPunch2.zip"),
  @("http://github.com/e-dog/ProceduralFairings/releases/download/v3.15/ProcFairings_3.15.zip", "ProcFairings.zip"),
  @("http://github.com/Mihara/RasterPropMonitor/releases/download/v0.20.0/RasterPropMonitor.0.20.0.zip", "RasterPropMonitor.zip"),
  @("http://kerbalstuff.com/mod/71/RealChute%20Parachute%20Systems/download/1.3.2.3", "RealChute.zip"),
  @("http://github.com/NathanKell/ModularFuelSystem/releases/download/rf-v10.4.1/RealFuels_v10.4.1.zip", "RealFuels.zip"),
  @("http://github.com/KSP-RO/RealismOverhaul/releases/download/v10.0.0/RealismOverhaul-v10.0.0.zip", "RealismOverhaul.zip"),
  @("http://github.com/KSP-RO/RealSolarSystem/releases/download/v10.0.2/RealSolarSystem_v10.0.2.zip", "RealSolarSystem.zip"),
  @("http://github.com/RemoteTechnologiesGroup/RemoteTech/releases/download/1.6.5/RemoteTech-1.6.5.zip", "RemoteTech.zip"),
  @("http://github.com/KSP-RO/RSS-Textures/releases/download/v10.0/2048.zip", "2048.zip"),
  @("http://github.com/Crzyrndm/RW-Saturatable/releases/download/1.10.1/Saturatable.RW.v1.10.1.zip", "Saturatable.RW.zip"),
  @("http://github.com/taraniselsu/TacLifeSupport/releases/download/v0.11.1.20/TacLifeSupport_0.11.1.20.zip", "TacLifeSupport.zip"),
  @("http://github.com/ducakar/TextureReplacer/releases/download/v2.4.6/TextureReplacer-2.4.6.zip", "TextureReplacer.zip"),
  @("http://blizzy.de/toolbar/Toolbar-1.7.9.zip", "Toolbar.zip"),
  @("http://kerbal.curseforge.com/ksp-mods/220213-taurus-hcv-3-75-m-csm-system/files/2239849/download", "Taurus.zip"),
  @("http://github.com/DMagic1/Orbital-Science/releases/download/v1.0.4/DMagic_Orbital_Science-1.0.4.zip", "DMagic_Orbital_Science.zip"),
  @("http://kerbalstuff.com/mod/84/Engine%20Ignitor/download/3.4.1.1", "EngineIgnitor.zip"),
  @("http://ksptomars.org/public/EngineIgnitor.dll", "EngineIgnitor.dll"),
  @("http://kerbal.curseforge.com/ksp-mods/228561-kerbal-inventory-system-kis/files/2240842/download", "KIS.zip"),
  @("http://kerbal.curseforge.com/ksp-mods/223900-kerbal-attachment-system-kas/files/2240844/download", "KAS.zip"),
  @("http://github.com/timmersuk/Timmers_KSP/releases/download/0.7.3.2/KeepFit-0.7.3.2.zip", "KeepFit.zip"),
  @("http://kerbalstuff.com/mod/8/Magic%20Smoke%20Industries%20Infernal%20Robotics/download/0.21.2", "InfernalRobotics.zip"),
  @("http://github.com/ClawKSP/KSP-Stock-Bug-Fix-Modules/releases/download/v1.0.2e.3/StockBugFixModules.v1.0.2e.3.zip", "StockBugFixModules.zip"),
  @("http://github.com/ClawKSP/KSP-Stock-Bug-Fix-Modules/releases/download/v1.0.2e.3/StockPlusController.cfg", "StockPlusController.cfg"),
  @("http://github.com/KSP-KOS/KOS/releases/download/v0.17.3/kOS-v0.17.3.zip", "kOS.zip"),
  @("http://github.com/UbioWeldingLtd/UbioWeldContinued/releases/download/2.1.3/UbioWeldContinued-2.1.3.zip", "UbioWeldContinued.zip"),
  @("http://kerbalstuff.com/mod/250/Universal%20Storage/download/1.1.0.6", "UniversalStorage.zip"),
  @("http://kerbalstuff.com/mod/344/TweakScale%20-%20Rescale%20Everything%21/download/v2.2.1", "TweakScale.zip"),
  @("http://github.com/BobPalmer/MKS/releases/download/0.31.4/UKS_0.31.4.zip", "UKS.zip"),
  @("http://kerbalstuff.com/mod/668/PersistentRotation/download/0.5.3", "PersistentRotation.zip"),
  @("http://kerbalstuff.com/mod/515/B9%20Aerospace%20Procedural%20Parts/download/0.40", "B9ProcParts.zip"),
  @("http://kerbalstuff.com/mod/255/TweakableEverything/download/1.12", "TweakableEverything.zip"),
  @("http://ksptomars.org/public/HabitatPack_04.1.zip", "HabitatPack.zip"),
  @("http://kerbalstuff.com/mod/450/Hullcam%20VDS/download/0.40", "HullcaMove-ItemDS.zip"),
  @("http://dl.orangedox.com/ilvCeXLsPxxWNdz1VY/JDiminishingRTG_v1.3a.zip?dl=1", "JDiminishingRTG.zip"),
  @("http://ksptomars.org/public/AIES_Aerospace151.zip", "AIES_Aerospace151.zip"),
  @("http://dl.dropboxusercontent.com/u/72893034/AIES_Patches/AIES_Node_Patch.cfg.zip", "AIES_Node_Patch.cfg.zip"),
  @("http://kerbalstuff.com/mod/361/NEBULA%20Decals/download/1.01", "NebulaDecals.zip")
)

Write-Output "Downloading Base Mods."
$index = 0
foreach ($baseMod in $baseModPack) {
  $index = $index+1
  Write-Output "[$index of $($baseModPack.count)]: $($baseMod[1])"
  Invoke-WebRequest -Uri $baseMod[0] -OutFile $baseMod[1]
}

# Dev mods!
if (-not $b -and -not $c){
  $devModPack = @(
    @("http://kerbal.curseforge.com/ksp-mods/220221-mechjeb/files/2238724/download", "mechjeb2.zip"),
    @("http://github.com/malahx/QuickSearch/releases/download/v1.13/QuickSearch-1.13.zip", "QuickSearch.zip"),
    @("http://github.com/snjo/FShangarExtender/releases/download/v3.3/FShangarExtender_3_3.zip", "FShangarExtender.zip"),
    @("http://github.com/Crzyrndm/FilterExtension/releases/download/2.3.0/Filter.Extensions.v2.3.0.1.zip", "Filter.Extensions.zip"),
    @("http://github.com/Swamp-Ig/ProceduralParts/releases/download/v1.1.6/ProceduralParts-1.1.6.zip", "ProceduralParts.zip"),
    @("http://ksptomars.org/public/HyperEdit-1.4.1_for-KSP-1.0.zip", "HyperEdit.zip"),
    @("http://github.com/malahx/QuickScroll/releases/download/v1.31/QuickScroll-1.31.zip", "QuickScroll.zip"),
    @("http://github.com/MachXXV/EditorExtensions/releases/download/v2.12/EditorExtensions_v2.12.zip", "EditorExtensions.zip"),
    @("http://github.com/nodrog6/LightsOut/releases/download/v0.1.4/LightsOut-v0.1.4.zip", "LightsOut.zip"),
    @("http://kerbal.curseforge.com/ksp-mods/220602-rcs-build-aid/files/2243090/download", "RCSbuildAid.zip"),
    @("http://kerbalstuff.com/mod/414/StripSymmetry/download/v1.6", "StripSymmetry.zip"),
    @("http://kerbal.curseforge.com/ksp-mods/220530-part-wizard/files/2237849/download", "PartWizard.zip"),
    @("http://kerbalstuff.com/mod/731/Vessel%20Viewer/download/0.71", "VesselViewer.zip"),
    @("http://kerbalstuff.com/mod/776/Take%20Command/download/1.1.4", "TakeCommand.zip")
  )

  Write-Output "Downloading Dev Mods."
  $index = 0
  foreach ($devMod in $devModPack) {
	$index = $index+1
    Write-Output "[$index of $($devModPack.count)]: $($devMod[1])"
    Invoke-WebRequest -Uri $devMod[0] -OutFile $devMod[1]
  }
}

# Beauty mods!
if ($b -or $f){
  Remove-Item -force 2048.zip #Remove low resolution RSS textures.

  $beautyModPack = @(
    @("http://github.com/KSP-RO/RSS-Textures/releases/download/v10.0/8192.zip", "8192.zip"),
    @("http://kerbal.curseforge.com/ksp-mods/224876-planetshine/files/2237465/download", "PlanetShine.zip"),
    @("http://github.com/MOARdV/DistantObject/releases/download/v1.5.6/DistantObject_1.5.6.zip", "DistantObject.zip"),
    @("http://beta.kerbalstuff.com/mod/124/Chatterer/download/0.9.5", "Chatterer.zip"),
    @("http://kerbalstuff.com/mod/190/Camera%20Tools/download/v1.3", "CameraTools.zip"),
    @("http://kerbalstuff.com/mod/381/Collision%20FX/download/3.2", "CollisionFX.zip"),
    @("http://kerbalstuff.com/mod/817/EngineLighting/download/1.3.6", "EngineLighting.zip"),
    @("http://kerbal.curseforge.com/ksp-mods/220207-hotrockets-particle-fx-replacement/files/2244672/download", "hotrocket.zip"),
    @("http://kerbalstuff.com/mod/743/Improved%20Chase%20Camera/download/v1.5.1", "ImprovedChaseCam.zip"),
    @("http://kerbalstuff.com/mod/224/Rover%20Wheel%20Sounds/download/1.2", "RoverWheelSounds.zip"),
    @("http://kerbalstuff.com/mod/700/Scatterer/download/0.151", "Scatterer.zip"),
    @("http://github.com/HappyFaceIndustries/BetterTimeWarp/releases/download/2.0/BetterTimeWarp_2.0.zip", "BetterTimeWarp.zip"),
    @("http://ksp.sarbian.com/jenkins/job/SmokeScreen/40/artifact/SmokeScreen-2.6.3.0.zip", "SmokeScreen.zip"),
    @("http://github.com/richardbunt/Telemachus/releases/download/v1.4.29.0/Telemachus_1_4_29_0.zip", "Telemachus.zip") 
  )

  Write-Output "Downloading Dev Mods."
  $index = 0
  foreach ($beautyMod in $beautyModPack) {
  	$index = $index+1
    Write-Output "[$index of $($beautyModPack.count)]: $($beautyMod[1])"
    Invoke-WebRequest -Uri $beautyMod[0] -OutFile $beautyMod[1]
  }
}

# Unzip all the mods

Write-Output "Extracting Mods"
$childItems = Get-ChildItem ./ -Filter *.zip
$index = 0
$childItems |
foreach-Object {
  $index = $index + 1
  $dirname = $_.FullName | %{$_ -replace ".zip",""}
  new-item -itemtype directory $dirname > $null
  if ($?){
    Set-Location $dirname
  if ($?){
      Write-Output "[$index of $($childItems.count)]: $_"
      unzip $_.FullName > $null
      Set-Location ..
    }else{
      Write-Output "Could not unpack $_ - Set-Location failed"
    }
  }else{  
    Write-Output "Could not unpack $_ - new-item -itemtype directory failed"
  }
}

# Move all the mods to GameData folder
Write-Output "Moving Mods"
Get-ChildItem ./*/* -Filter GameData |
foreach-Object {
  Copy-Item -force -recurse $_/* ../GameData
}
Set-Location ..

# Custom move for base install
Copy-Item -force -recurse ksp2m_mods/CrossFeedEnabler/* GameData
Copy-Item -force -recurse ksp2m_mods/DeadlyReentry/* GameData
Copy-Item -force -recurse ksp2m_mods/RealFuels/* GameData
Copy-Item -force -recurse ksp2m_mods/RealSolarSystem/* GameData
Copy-Item -force -recurse ksp2m_mods/Toolbar/Toolbar-1.7.9/GameData/* GameData
Copy-Item -force -recurse ksp2m_mods/ksp-avc/* GameData
Copy-Item -force -recurse ksp2m_mods/EngineIgnitor/* GameData
Copy-Item -force -recurse "ksp2m_mods/KWRocketry/KW Release Package v2.7 (Open this, don't extract it)/GameData/*" GameData
Copy-Item -force -recurse ksp2m_mods/UniversalStorage/* GameData
Copy-Item -force -recurse ksp2m_mods/StockBugFixModules/* GameData
Copy-Item -force -recurse ksp2m_mods/AIES_Aerospace151/* GameData
Copy-Item -force -recurse ksp2m_mods/HullcaMove-ItemDS/* GameData
Copy-Item -force -recurse ksp2m_mods/JDiminishingRTG/JDiminishingRTG_v1_3a/GameData/* GameData
Copy-Item -force -recurse ksp2m_mods/NebulaDecals/NEBULA/* GameData

# Custom move for dev
if (-not $b -and -not $c){
Copy-Item -force -recurse ksp2m_mods/mechjeb2/* GameData
Copy-Item -force -recurse ksp2m_mods/VesselViewer/* GameData
Copy-Item -force -recurse ksp2m_mods/FShangarExtender/* GameData
Copy-Item -force -recurse ksp2m_mods/PartWizard/* GameData
Copy-Item -force -recurse ksp2m_mods/RCSbuildAid/* GameData
Copy-Item -force -recurse ksp2m_mods/StripSymmetry/Gamedata/* GameData
Copy-Item -force -recurse ksp2m_mods/EditorExtensions/* GameData
}

# Custom move for beauty
if ($b -or $f){
Copy-Item -force -recurse ksp2m_mods/hotrocket/* GameData
Copy-Item -force -recurse "ksp2m_mods/DistantObject/Alternate Planet Color Configs/Real Solar System (metaphor's PlanetFactory config)/*" GameData
Copy-Item -force -recurse ksp2m_mods/EngineLighting/EngineLight/GameData/* GameData
Copy-Item -force -recurse ksp2m_mods/ImprovedChaseCam/* GameData
Copy-Item -force -recurse "ksp2m_mods/PlanetShine/Alternate Colors/Real Solar System/*" GameData
Copy-Item -force -recurse ksp2m_mods/RoverWheelSounds/* GameData
}

# Fix some configs
Write-Output "Adapting Configs"
Copy-Item -recurse -force ksp2m_mods/RealismOverhaul/GameData/* GameData #We do this to make sure that we use the RO/RSS configs and not the configs provided by plugins installed after RO/RSS
Copy-Item -force ksp2m_mods/RealismOverhaul/GameData/RealismOverhaul/RemoteTech_Settings.cfg GameData/RemoteTech/RemoteTech_Settings.cfg
Copy-Item -force ksp2m_mods/EngineIgnitor.dll GameData/EngineIgnitor/Plugins/EngineIgnitor.dll
Copy-Item -force ksp2m_mods/TextureReplacer/Extras/MM_ReflectionPluginWrapper.cfg GameData
Copy-Item -force ksp2m_mods/StockPlusController.cfg GameData
Copy-Item -force ksp2m_mods/AIES_Node_Patch.cfg/AIES_Node_Patch.cfg GameData

# Clean up
Write-Output "Starting Clean up"
Remove-Item -Recurse -Force ksp2m_mods
Set-Location GameData
new-item -itemtype directory licensesAndReadmes
if (Test-Path *.txt){
Move-Item *.txt licensesAndReadmes
}
if (Test-Path *.md){
Move-Item *.md licensesAndReadmes
}
if (Test-Path *.pdf){
Move-Item *.pdf licensesAndReadmes
}
if (Test-Path *.htm){
Move-Item *.htm licensesAndReadmes
}

#Remove-Item ModuleManager.2.5.1.dll
Write-Output "Removing old ModuleManager Versions"
Remove-Item ModuleManager.2.6.1.dll, ModuleManager.2.6.2.dll, ModuleManager.2.6.3.dll

# Remove unneded parts
Write-Output "Removing unneeded parts"
# HabitatPack
Remove-Item -Recurse -Force HabitatPack/Parts/Basemount

# FASA
Set-Location FASA
Remove-Item -Recurse -Force Agencies, Flags, ICBM, Mercury, Resources
Set-Location Apollo
Remove-Item -Recurse -Force ApolloCSM, FASA_Apollo_Fairings, FASA_Apollo_Str, Science
Set-Location LEM
Remove-Item -Recurse -Force Antennas, AscentStage, DescentStage, DockingCone, InterStage, LandingLegs
Set-Location ../../Gemini2
Remove-Item -Recurse -Force FASA_ASAS_MiniComp, FASA_Fairings_Plate_2m, FASA_Gemini_BigG, FASA_Gemini_Dec_Dark, FASA_Gemini_Engine_Fuel2, FASA_Gemini_LES, FASA_Gemini_LFT, FASA_Gemini_LFTLong, FASA_Gemini_Lander_Eng, FASA_Gemini_Lander_Legs, FASA_Gemini_Lander_Pod, FASA_Gemini_MOL, FASA_Gemini_NoseCone2, FASA_Gemini_Parachute2, FASA_Gemini_Pod2, FASA_Gemini_RCS_Thruster, FASA_Gemini_SAS_RCS, FASA_WingGemini, SmallGearBay
Set-Location ../Probes
Remove-Item -Recurse -Force Explorer, Pioneer, Probe_Parachute_Box
Set-Location ../..

# Engine Ignitor
Set-Location EngineIgnitor
Remove-Item -Recurse -Force Parts
Set-Location ..

# DMagic -> UniversalStorage Parts
Set-Location DMagicOrbitalScience
Remove-Item -Recurse -Force UniversalStorage
Set-Location ..

# KW Rocketry
Set-Location ../../KWRocketry/Parts
Set-Location Fuel
Remove-Item -Recurse -Force KW_Universal_Tanks
Set-Location ../Structural
Remove-Item -Recurse -Force KWFuelAdapter, KWFlatadapter*
Set-Location ../../..

# NovaPunch2 
Set-Location NovaPunch2
Remove-Item -Recurse -Force Agencies, Flags
Set-Location Parts
Remove-Item -Recurse -Force ControlPods, Fairings, FuelTanks, NoseCone, SAS, YawmasterCSM
Set-Location CouplersAndAdapters
Remove-Item -Recurse -Force NP_interstage*
Set-Location ../Freyja
Remove-Item -Recurse -Force FreyjaEng, FreyjaRCS, FreyjaTrunk
Set-Location ../Odin2
Remove-Item -Recurse -Force OdinFairings, OdinPod, OdinRCS, OdinServiceModule
Set-Location ../../..

# UKS/MKS
Set-Location UmbraSpaceIndustries/Kolonization
Remove-Item -Recurse -Force Flags
Set-Location Parts
Remove-Item -force MK3*, MKS_A*, MKS_C*, MKS_D*, MKS_E*, MKS_F*, MKS_K*, MKS_L*, MKS_M*, MKS_P*, MKS_S*, MKS_W*, MKV_Ag*, MKV_B*, MKV_L*, MKV_Pod.cfg, MKV_W*, MiniRover.cfg, OKS_A*, OKS_Col*, OKS_Ha*, OKS_K*, OKS_P*, OKS_S*, OKS_W*, OctoLander.cfg, ScanOMatic.cfg
Set-Location ../../..

# UniversalStorage
Set-Location UniversalStorage
Remove-Item -Recurse -Force Flags
Set-Location ..

Set-Location $startingPath

Write-Output "Finished!"
