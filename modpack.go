package main
import (
"fmt"
"net/http"
"io"
"os"
)

func download(A [][]string) {
  for i := 0; i < cap(A); i++ {
    uri := A[i][0]
    file := A[i][1]
    fmt.Println("[", i+1, " of ", cap(A), "]: ", file)
    out, _ := os.Create(file)
    defer out.Close()
    resp, _ := http.Get(uri)
    defer resp.Body.Close()
    n, _ := io.Copy(out, resp.Body)
    fmt.Println (n)
  }
}

type TwoDText [][]string //A slice of string slices

var coremods = TwoDText{
  []string{"http://kerbalstuff.com/mod/361/NEBULA%20Decals/download/1.01", "NebulaDecals.zip"},                                                                    //KSP v0.25
  []string{"http://github.com/NathanKell/CrossFeedEnabler/releases/download/v3.3/CrossFeedEnabler_v3.3.zip", "CrossFeedEnabler.zip"},                              //KSP v1.0
  []string{"http://github.com/Starwaster/DeadlyReentry/releases/download/v7.1.0/DeadlyReentry_7.1.0_The_Melificent_Edition.zip", "DeadlyReentry.zip"},             //KSP v1.0
  []string{"http://github.com/BobPalmer/CommunityResourcePack/releases/download/0.4.3/CRP_0.4.3.zip", "CRP.zip"},                                                  //KSP v1.0.4
  []string{"http://github.com/codepoetpbowden/ConnectedLivingSpace/releases/download/1.1.3.1/Connected_Living_Space-1.1.3.1.zip", "Connected_Living_Space.zip"},   //KSP v1.0.2
  []string{"http://beta.kerbalstuff.com/mod/67/KW%20Rocketry/download/2.7", "KWRocketry.zip"},                                                                     //KSP v1.0.2
  []string{"http://kerbalstuff.com/mod/26/NovaPunch/download/2.09", "NovaPunch2.zip"},                                                                             //KSP v1.0.2
  []string{"http://github.com/Mihara/RasterPropMonitor/releases/download/v0.21.0/RasterPropMonitor.0.21.0.zip", "RasterPropMonitor.zip"},                          //KSP v1.0.4
  []string{"http://kerbalstuff.com/mod/71/RealChute%20Parachute%20Systems/download/1.3.2.3", "RealChute.zip"},                                                     //KSP v1.0.2
  []string{"http://github.com/Crzyrndm/RW-Saturatable/releases/download/1.10.1/Saturatable.RW.v1.10.1.0.zip", "Saturatable.RW.zip"},                               //KSP v1.0.2
  []string{"http://github.com/taraniselsu/TacLifeSupport/releases/download/v0.11.1.20/TacLifeSupport_0.11.1.20.zip", "TacLifeSupport.zip"},                        //KSP v1.0.2
  []string{"http://blizzy.de/toolbar/Toolbar-1.7.9.zip", "Toolbar.zip"},                                                                                           //KSP v1.0.2
  []string{"http://kerbal.curseforge.com/ksp-mods/228561-kerbal-inventory-system-kis/files/2240842/download", "KIS.zip"},                                          //KSP v1.0.2
  []string{"http://kerbal.curseforge.com/ksp-mods/223900-kerbal-attachment-system-kas/files/2240844/download", "KAS.zip"},                                         //KSP v1.0.2
  []string{"http://github.com/UbioWeldingLtd/UbioWeldContinued/releases/download/2.1.3/UbioWeldContinued-2.1.3.zip", "UbioWeldContinued.zip"},                     //KSP v1.0.2
  []string{"http://kerbalstuff.com/mod/668/PersistentRotation/download/0.5.3", "PersistentRotation.zip"},                                                          //KSP v1.0.2
  []string{"http://kerbalstuff.com/mod/450/Hullcam%20VDS/download/0.40", "HullcaMove-ItemDS.zip"},                                                                 //KSP v1.0.2
  []string{"http://dl.orangedox.com/ilvCeXLsPxxWNdz1VY/JDiminishingRTG_v1.3a.zip?dl=1", "JDiminishingRTG.zip"},                                                    //KSP v1.0.2
  []string{"http://github.com/ferram4/BetterBuoyancy/releases/download/v1.3/BetterBuoyancy_v1.3.zip", "BetterBuoyancy.zip"},                                       //KSP v1.0.3
  []string{"http://github.com/ferram4/Ferram-Aerospace-Research/releases/download/v0.15_3_1_Garabedian/FAR_0_15_3_1_Garabedian.zip", "FAR.zip"},                   //KSP v1.0.3
  []string{"http://github.com/ferram4/Kerbal-Joint-Reinforcement/releases/download/v3.1.4/KerbalJointReinforcement_v3.1.4.zip", "KerbalJointReinforcement.zip"},   //KSP v1.0.3
  []string{"http://github.com/KSP-RO/RSS-Textures/releases/download/v10.0/2048.zip", "2048.zip"},                                                                  //KSP v?.?.?
  []string{"http://github.com/BobPalmer/MKS/releases/download/0.31.4/UKS_0.31.4.zip", "UKS.zip"},                                                                  //KSP v?.?.?
  []string{"http://ksptomars.org/public/HabitatPack_04.1.zip", "HabitatPack.zip"},                                                                                 //KSP v?.?.?
  []string{"http://ksptomars.org/public/AIES_Aerospace151.zip", "AIES_Aerospace151.zip"},                                                                          //KSP v?.?.?
  []string{"http://dl.dropboxusercontent.com/u/72893034/AIES_Patches/AIES_Node_Patch.cfg.zip", "AIES_Node_Patch.cfg.zip"},                                         //KSP v?.?.?
  []string{"http://ksptomars.org/public/KSPtoMars.zip", "KSPtoMars.zip"},                                                                                          //KSP v?.?.?
  []string{"http://github.com/camlost2/AJE/releases/download/2.2.1/Advanced_Jet_Engine-2.2.1.zip", "Advanced_Jet_Engine.zip"},                                     //KSP v1.0.4
  []string{"http://kerbalstuff.com/mod/27/FASA/download/5.35", "FASA.zip"},                                                                                        //KSP v1.0.4
  []string{"http://kerbal.curseforge.com/ksp-mods/220462-ksp-avc-add-on-version-checker/files/2216818/download", "ksp-avc.zip"},                                   //KSP v1.0.4
  []string{"http://github.com/KSP-RO/SolverEngines/releases/download/v1.5/SolverEngines_v1.5.zip", "SolverEngines.zip"},                                           //KSP v1.0.4
  []string{"http://github.com/e-dog/ProceduralFairings/releases/download/v3.15/ProcFairings_3.15.zip", "ProcFairings.zip"},                                        //KSP v1.0.4
  []string{"http://github.com/NathanKell/ModularFuelSystem/releases/download/rf-v10.4.4/RealFuels_v10.4.4.zip", "RealFuels.zip"},                                  //KSP v1.0.4
  []string{"http://github.com/KSP-RO/RealismOverhaul/releases/download/v10.1.0/RealismOverhaul-v10.1.0.zip", "RealismOverhaul.zip"},                               //KSP v1.0.4
  []string{"http://github.com/KSP-RO/RealSolarSystem/releases/download/v10.1/RealSolarSystem_v10.1.zip", "RealSolarSystem.zip"},                                   //KSP v1.0.4
  []string{"http://github.com/RemoteTechnologiesGroup/RemoteTech/releases/download/1.6.7/RemoteTech-1.6.7.zip", "RemoteTech.zip"},                                 //KSP v1.0.4
  []string{"http://github.com/ducakar/TextureReplacer/releases/download/v2.4.7/TextureReplacer-2.4.7.zip", "TextureReplacer.zip"},                                 //KSP v1.0.4
  []string{"http://kerbal.curseforge.com/ksp-mods/220213-taurus-hcv-3-75-m-csm-system/files/2244776/download", "Taurus.zip"},                                      //KSP v1.0.4
  []string{"http://github.com/DMagic1/Orbital-Science/releases/download/v1.0.7/DMagic_Orbital_Science-1.0.7.zip", "DMagic_Orbital_Science.zip"},                   //KSP v1.0.4
  []string{"http://github.com/timmersuk/Timmers_KSP/releases/download/0.7.3.3/KeepFit-0.7.3.3.zip", "KeepFit.zip"},                                                //KSP v1.0.4
  []string{"http://kerbalstuff.com/mod/8/Magic%20Smoke%20Industries%20Infernal%20Robotics/download/0.21.3", "InfernalRobotics.zip"},                               //KSP v1.0.4
  []string{"http://github.com/ClawKSP/KSP-Stock-Bug-Fix-Modules/releases/download/v1.0.4a.1/StockBugFixModules.v1.0.4a.1.zip", "StockBugFixModules.zip"},          //KSP v1.0.4
  []string{"http://github.com/ClawKSP/KSP-Stock-Bug-Fix-Modules/releases/download/v1.0.4a.1/StockPlusController.zip", "StockPlusController.cfg"},                  //KSP v1.0.4
  []string{"http://github.com/KSP-KOS/KOS/releases/download/v0.17.3/kOS-v0.17.3.zip", "kOS.zip"},                                                                  //KSP v1.0.4
  []string{"http://kerbalstuff.com/mod/250/Universal%20Storage/download/1.1.0.6", "UniversalStorage.zip"},                                                         //KSP v1.0.4
  []string{"http://kerbalstuff.com/mod/344/TweakScale%20-%20Rescale%20Everything%21/download/v2.2.1", "TweakScale.zip"},                                           //KSP v1.0.4
  []string{"http://kerbalstuff.com/mod/515/B9%20Aerospace%20Procedural%20Parts/download/0.40", "B9ProcParts.zip"},                                                 //KSP v1.0.4
  []string{"http://kerbalstuff.com/mod/255/TweakableEverything/download/1.12", "TweakableEverything.zip"},                                                         //KSP v1.0.4
  []string{"http://github.com/Swamp-Ig/ProceduralParts/releases/download/v1.1.6/ProceduralParts-1.1.6.zip", "ProceduralParts.zip"},                                 //KSP v1.0.4
}

func main() {
download(coremods)

}