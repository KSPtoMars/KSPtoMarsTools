package modsources

type TwoDText [][]string //A slice of string slices

var Basemods = TwoDText{
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
  []string{"http://github.com/Swamp-Ig/ProceduralParts/releases/download/v1.1.6/ProceduralParts-1.1.6.zip", "ProceduralParts.zip"},                                //KSP v1.0.4
}

var Devmods = TwoDText{
  []string{"http://github.com/snjo/FShangarExtender/releases/download/v3.3/FShangarExtender_3_3.zip", "FShangarExtender.zip"},                                   //KSP v1.0
  []string{"http://kerbalstuff.com/mod/414/StripSymmetry/download/v1.6", "StripSymmetry.zip"},                                                                   //KSP v1.0
  []string{"http://kerbal.curseforge.com/ksp-mods/220602-rcs-build-aid/files/2243090/download", "RCSbuildAid.zip"},                                              //KSP v1.0.2
  []string{"http://kerbalstuff.com/mod/731/Vessel%20Viewer/download/0.71", "VesselViewer.zip"},                                                                  //KSP v1.0.2
  []string{"http://github.com/Crzyrndm/FilterExtension/releases/download/2.3.0/Filter.Extensions.v2.3.0.1.zip", "Filter.Extensions.zip"},                        //KSP v1.0.3
  []string{"http://github.com/MachXXV/EditorExtensions/releases/download/v2.12/EditorExtensions_v2.12.zip", "EditorExtensions.zip"},                             //KSP v1.0.3
  []string{"http://ksptomars.org/public/HyperEdit-1.4.1_for-KSP-1.0.zip", "HyperEdit.zip"},                                                                      //KSP v?.?.?
  []string{"http://kerbal.curseforge.com/ksp-mods/220530-part-wizard/files/2246104/download", "PartWizard.zip"},                                                 //KSP v1.0.4
  []string{"http://kerbal.curseforge.com/ksp-mods/220221-mechjeb/files/2245658/download", "mechjeb2.zip"},                                                       //KSP v1.0.4
  []string{"http://github.com/nodrog6/LightsOut/releases/download/v0.1.4/LightsOut-v0.1.4.zip", "LightsOut.zip"},                                                //KSP v1.0.4
  []string{"https://github.com/CYBUTEK/KerbalEngineer/releases/download/1.0.17.0/KerbalEngineer-1.0.17.0.zip", "KerbalEngineer.zip"},                            //KSP v1.0.4
  []string{"http://kerbalstuff.com/mod/776/Take%20Command/download/1.1.4", "TakeCommand.zip"},                                                                   //KSP v1.0.4
  []string{"http://github.com/malahx/QuickSearch/releases/download/v1.13/QuickSearch-1.13.zip", "QuickSearch.zip"},                                              //KSP v1.0.x
  []string{"http://github.com/malahx/QuickScroll/releases/download/v1.31/QuickScroll-1.31.zip", "QuickScroll.zip"},                                              //KSP v1.0.x
}

var Beautymods = TwoDText{
  []string{"http://kerbal.curseforge.com/ksp-mods/224876-planetshine/files/2237465/download", "PlanetShine.zip"},                                                //KSP v1.0
  []string{"http://kerbalstuff.com/mod/224/Rover%20Wheel%20Sounds/download/1.2", "RoverWheelSounds.zip"},                                                        //KSP v1.0
  []string{"http://kerbalstuff.com/mod/190/Camera%20Tools/download/v1.3", "CameraTools.zip"},                                                                    //KSP v1.0.2
  []string{"http://kerbalstuff.com/mod/381/Collision%20FX/download/3.2", "CollisionFX.zip"},                                                                     //KSP v1.0.2
  []string{"http://kerbalstuff.com/mod/700/Scatterer/download/0.151", "Scatterer.zip"},                                                                          //KSP v1.0.2
  []string{"http://github.com/KSP-RO/RSS-Textures/releases/download/v10.0/8192.zip", "8192.zip"},                                                                //KSP v?.?.?
  []string{"http://github.com/MOARdV/DistantObject/releases/download/v1.5.7/DistantObject_1.5.7.zip", "DistantObject.zip"},                                      //KSP v1.0.4
  []string{"http://beta.kerbalstuff.com/mod/124/Chatterer/download/0.9.5", "Chatterer.zip"},                                                                     //KSP v1.0.4
  []string{"http://kerbalstuff.com/mod/817/EngineLighting/download/1.4.0", "EngineLighting.zip"},                                                                //KSP v1.0.4
  []string{"http://kerbal.curseforge.com/ksp-mods/220207-hotrockets-particle-fx-replacement/files/2244672/download", "hotrocket.zip"},                           //KSP v1.0.4
  []string{"http://kerbalstuff.com/mod/743/Improved%20Chase%20Camera/download/v1.5.1", "ImprovedChaseCam.zip"},                                                  //KSP v1.0.4
  []string{"http://github.com/richardbunt/Telemachus/releases/download/v1.4.30.0/Telemachus_1_4_30_0.zip", "Telemachus.zip"},                                    //KSP v1.0.4
  []string{"https://ksp.sarbian.com/jenkins/job/SmokeScreen/44/artifact/SmokeScreen-2.6.6.0.zip", "SmokeScreen.zip"},                                            //KSP v1.0.x
  []string{"http://github.com/HappyFaceIndustries/BetterTimeWarp/releases/download/2.0/BetterTimeWarp_2.0.zip", "BetterTimeWarp.zip"},                           //KSP v1.0.x
}
