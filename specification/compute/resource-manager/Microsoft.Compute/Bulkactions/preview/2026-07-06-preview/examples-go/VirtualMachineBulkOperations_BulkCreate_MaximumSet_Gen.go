package armbulkactions_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/compute/armbulkactions"
)

// Generated from example definition: 2026-07-06-preview/VirtualMachineBulkOperations_BulkCreate_MaximumSet_Gen.json
func ExampleVirtualMachineBulkOperationsClient_BulkCreateOperation_virtualMachineBulkOperationsBulkCreateExample() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armbulkactions.NewClientFactory("1FBA3C66-5C9C-4391-B72F-9F52735FC9F2", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewVirtualMachineBulkOperationsClient().BulkCreateOperation(ctx, "rgBulkactions", "useast2euap", armbulkactions.ExecuteCreateContent{
		ResourceConfigParameters: &armbulkactions.ResourceProvisionPayload{
			BaseProfile: map[string]any{
				"plan": map[string]any{
					"name":          "iemasqqkbixbewezyrhnpntjd",
					"publisher":     "bvggylbvfstnscuupuithafvvgc",
					"product":       "bguuzrknnuohugjhernflurpx",
					"promotionCode": "bxgonranwqoryfkhkfaumdgz",
					"version":       "uyxetqmmzvqianqv",
				},
				"zones": []any{
					"wczj",
				},
				"identity": map[string]any{
					"type": "SystemAssigned",
					"userAssignedIdentities": map[string]any{
						"key7": map[string]any{},
					},
				},
				"extendedLocation": map[string]any{
					"name": "gbnxzymbdkxhwjpqkur",
					"type": "EdgeZone",
				},
				"placement": map[string]any{
					"zonePlacementPolicy": "Any",
					"includeZones": []any{
						"inagtbtedobdea",
					},
					"excludeZones": []any{
						"pvvwrhuhdpvbacwmesblpgwzk",
					},
				},
				"tags": map[string]any{
					"key6824": "cefndldgkx",
				},
				"properties": map[string]any{
					"scheduledEventsPolicy": map[string]any{
						"userInitiatedRedeploy": map[string]any{
							"automaticallyApprove": true,
						},
						"userInitiatedReboot": map[string]any{
							"automaticallyApprove": true,
						},
						"scheduledEventsAdditionalPublishingTargets": map[string]any{
							"eventGridAndResourceGraph": map[string]any{
								"enable":                    true,
								"scheduledEventsApiVersion": "lifncbftlkounuyfn",
							},
						},
						"allInstancesDown": map[string]any{
							"automaticallyApprove": true,
						},
					},
					"storageProfile": map[string]any{
						"imageReference": map[string]any{
							"publisher":               "ojlplghybdamadvsrq",
							"offer":                   "uvnqoxhkxefqwbsvjgbswqy",
							"sku":                     "hajdxhjmlkx",
							"version":                 "u",
							"sharedGalleryImageId":    "fz",
							"communityGalleryImageId": "tsfpcq",
							"id":                      "cdbrkpdicibtlliq",
						},
						"osDisk": map[string]any{
							"osType": "Windows",
							"encryptionSettings": map[string]any{
								"diskEncryptionKey": map[string]any{
									"secretUrl": "vzkogocyw",
									"sourceVault": map[string]any{
										"id": "lvzxxyypkeqlflftmfn",
									},
								},
								"keyEncryptionKey": map[string]any{
									"keyUrl": "mjjkvgpoohatw",
									"sourceVault": map[string]any{
										"id": "lvzxxyypkeqlflftmfn",
									},
								},
								"enabled": true,
							},
							"name": "opogpznvctmraoajgizcyrfvpt",
							"vhd": map[string]any{
								"uri": "elpzggtxubepzgjqvdbjmbu",
							},
							"image": map[string]any{
								"uri": "elpzggtxubepzgjqvdbjmbu",
							},
							"caching":                 "None",
							"writeAcceleratorEnabled": true,
							"diffDiskSettings": map[string]any{
								"option":    "Local",
								"placement": "CacheDisk",
							},
							"createOption": "FromImage",
							"diskSizeGB":   2,
							"managedDisk": map[string]any{
								"storageAccountType": "Standard_LRS",
								"diskEncryptionSet": map[string]any{
									"id": "magvkzhdmzhktjlqkkk",
								},
								"securityProfile": map[string]any{
									"securityEncryptionType": "VMGuestStateOnly",
									"diskEncryptionSet": map[string]any{
										"id": "magvkzhdmzhktjlqkkk",
									},
								},
								"id": "numddbqmkxuu",
							},
							"deleteOption": "Delete",
						},
						"dataDisks": []any{
							map[string]any{
								"lun":  7,
								"name": "nbthfzqsxyqvqnbgcljxbwyyoj",
								"vhd": map[string]any{
									"uri": "elpzggtxubepzgjqvdbjmbu",
								},
								"image": map[string]any{
									"uri": "elpzggtxubepzgjqvdbjmbu",
								},
								"caching":                 "None",
								"writeAcceleratorEnabled": true,
								"createOption":            "FromImage",
								"diskSizeGB":              19,
								"managedDisk": map[string]any{
									"storageAccountType": "Standard_LRS",
									"diskEncryptionSet": map[string]any{
										"id": "magvkzhdmzhktjlqkkk",
									},
									"securityProfile": map[string]any{
										"securityEncryptionType": "VMGuestStateOnly",
										"diskEncryptionSet": map[string]any{
											"id": "magvkzhdmzhktjlqkkk",
										},
									},
									"id": "numddbqmkxuu",
								},
								"sourceResource": map[string]any{
									"id": "qnukyordmomtjjqabovlsxl",
								},
								"toBeDetached": true,
								"detachOption": "ForceDetach",
								"deleteOption": "Delete",
							},
						},
						"diskControllerType": "SCSI",
					},
					"hardwareProfile": map[string]any{
						"vmSize": "szrnjqwbruz",
						"vmSizeProperties": map[string]any{
							"vCpusAvailable": 24,
							"vCpusPerCore":   6,
						},
					},
					"additionalCapabilities": map[string]any{
						"ultraSSDEnabled":    true,
						"hibernationEnabled": true,
					},
					"osProfile": map[string]any{
						"computerName":  "bplxnfp",
						"adminUsername": "fxzbi",
						"adminPassword": "<a-password-goes-here>",
						"customData":    "hbdlirohsgnbrahscboc",
						"windowsConfiguration": map[string]any{
							"provisionVMAgent":       true,
							"enableAutomaticUpdates": true,
							"timeZone":               "t",
							"additionalUnattendContent": []any{
								map[string]any{
									"passName":      "OobeSystem",
									"componentName": "Microsoft-Windows-Shell-Setup",
									"settingName":   "AutoLogon",
									"content":       "rguazthnx",
								},
							},
							"patchSettings": map[string]any{
								"patchMode":         "Manual",
								"enableHotpatching": true,
								"assessmentMode":    "ImageDefault",
								"automaticByPlatformSettings": map[string]any{
									"rebootSetting": "Unknown",
									"bypassPlatformSafetyChecksOnUserSchedule": true,
								},
							},
							"winRM": map[string]any{
								"listeners": []any{
									map[string]any{
										"protocol":       "Http",
										"certificateUrl": "quhfapfpyeeocwvwtvuggoqqwt",
									},
								},
							},
						},
						"linuxConfiguration": map[string]any{
							"disablePasswordAuthentication": true,
							"ssh": map[string]any{
								"publicKeys": []any{
									map[string]any{
										"path":    "mrdfxnfjazxog",
										"keyData": "wfhrknkehgesontscqyrewfmhgwt",
									},
								},
							},
							"provisionVMAgent": true,
							"patchSettings": map[string]any{
								"patchMode":      "ImageDefault",
								"assessmentMode": "ImageDefault",
								"automaticByPlatformSettings": map[string]any{
									"rebootSetting": "Unknown",
									"bypassPlatformSafetyChecksOnUserSchedule": true,
								},
							},
							"enableVMAgentPlatformUpdates": true,
						},
						"secrets": []any{
							map[string]any{
								"sourceVault": map[string]any{
									"id": "lvzxxyypkeqlflftmfn",
								},
								"vaultCertificates": []any{
									map[string]any{
										"certificateUrl":   "crgbpfdvlohwkupdjp",
										"certificateStore": "hyx",
									},
								},
							},
						},
						"allowExtensionOperations":    true,
						"requireGuestProvisionSignal": true,
					},
					"networkProfile": map[string]any{
						"networkInterfaces": []any{
							map[string]any{
								"properties": map[string]any{
									"primary":      true,
									"deleteOption": "Delete",
								},
								"id": "ymfxctb",
							},
						},
						"networkApiVersion": "2020-11-01",
						"networkInterfaceConfigurations": []any{
							map[string]any{
								"name": "qrkzoctmzjketostzabnra",
								"properties": map[string]any{
									"primary":                     true,
									"deleteOption":                "Delete",
									"enableAcceleratedNetworking": true,
									"disableTcpStateTracking":     true,
									"enableFpga":                  true,
									"enableIPForwarding":          true,
									"networkSecurityGroup": map[string]any{
										"id": "lvzxxyypkeqlflftmfn",
									},
									"dnsSettings": map[string]any{
										"dnsServers": []any{
											"tqcqopnanyyiavfwhqbkarxtrfqbww",
										},
									},
									"ipConfigurations": []any{
										map[string]any{
											"name": "gqymuvgzzfmxqvdadx",
											"properties": map[string]any{
												"subnet": map[string]any{
													"id": "lvzxxyypkeqlflftmfn",
												},
												"primary": true,
												"publicIPAddressConfiguration": map[string]any{
													"name": "cwxsqjijtwbsyqdwht",
													"properties": map[string]any{
														"idleTimeoutInMinutes": 17,
														"deleteOption":         "Delete",
														"dnsSettings": map[string]any{
															"domainNameLabel":      "fampou",
															"domainNameLabelScope": "TenantReuse",
														},
														"ipTags": []any{
															map[string]any{
																"ipTagType": "hkjoxhqadudjartwooezaxl",
																"tag":       "xywunkjglkmmwfpf",
															},
														},
														"publicIPPrefix": map[string]any{
															"id": "lvzxxyypkeqlflftmfn",
														},
														"publicIPAddressVersion":   "IPv4",
														"publicIPAllocationMethod": "Dynamic",
													},
													"sku": map[string]any{
														"name": "Basic",
														"tier": "Regional",
													},
													"tags": map[string]any{
														"key5442": "qhpwpnylvmdthxazhxamnbhdfpf",
													},
												},
												"privateIPAddressVersion": "IPv4",
												"applicationSecurityGroups": []any{
													map[string]any{
														"id": "lvzxxyypkeqlflftmfn",
													},
												},
												"applicationGatewayBackendAddressPools": []any{
													map[string]any{
														"id": "lvzxxyypkeqlflftmfn",
													},
												},
												"loadBalancerBackendAddressPools": []any{
													map[string]any{
														"id": "lvzxxyypkeqlflftmfn",
													},
												},
											},
										},
									},
									"dscpConfiguration": map[string]any{
										"id": "lvzxxyypkeqlflftmfn",
									},
									"auxiliaryMode": "None",
									"auxiliarySku":  "None",
								},
								"tags": map[string]any{
									"key9436": "bjbadzbfvpszbsickv",
								},
							},
						},
					},
					"securityProfile": map[string]any{
						"uefiSettings": map[string]any{
							"secureBootEnabled": true,
							"vTpmEnabled":       true,
						},
						"encryptionAtHost": true,
						"securityType":     "TrustedLaunch",
						"encryptionIdentity": map[string]any{
							"userAssignedIdentityResourceId": "tnajlgbwcepmhytzb",
						},
						"proxyAgentSettings": map[string]any{
							"enabled":          true,
							"mode":             "Audit",
							"keyIncarnationId": 4,
							"wireServer": map[string]any{
								"mode":                                "Audit",
								"inVMAccessControlProfileReferenceId": "xvlzroy",
							},
							"imds": map[string]any{
								"mode":                                "Audit",
								"inVMAccessControlProfileReferenceId": "xvlzroy",
							},
							"addProxyAgentExtension": true,
						},
					},
					"diagnosticsProfile": map[string]any{
						"bootDiagnostics": map[string]any{
							"enabled":    true,
							"storageUri": "pxuhtzehlfsqolbdleirgj",
						},
					},
					"licenseType":          "ymwuemwuntbignqyvzqflvjpcdus",
					"extensionsTimeBudget": "dnyqmcijikzkltjav",
					"scheduledEventsProfile": map[string]any{
						"terminateNotificationProfile": map[string]any{
							"notBeforeTimeout": "owbwifqrlsdmm",
							"enable":           true,
						},
						"osImageNotificationProfile": map[string]any{
							"notBeforeTimeout": "ataqykjdakdvyyzdspaqnhd",
							"enable":           true,
						},
					},
					"userData": "nwjvxe",
					"capacityReservation": map[string]any{
						"capacityReservationGroup": map[string]any{
							"id": "lvzxxyypkeqlflftmfn",
						},
					},
					"applicationProfile": map[string]any{
						"galleryApplications": []any{
							map[string]any{
								"tags":                            "cmygipvpkegyclvpznfu",
								"order":                           8,
								"packageReferenceId":              "afrfkjdrtzftmwramfyu",
								"configurationReference":          "nmfaspclhidtznslsps",
								"treatFailureAsDeploymentFailure": true,
								"enableAutomaticUpgrade":          true,
							},
						},
					},
					"vmExtensions": []any{
						map[string]any{
							"name": "jkpmcxwuahpzwkvexgzpypk",
							"properties": map[string]any{
								"forceUpdateTag":          "dockqxgatsfzhctxrncuw",
								"publisher":               "qesyfldbfoaexyoywhcxafdtdwcg",
								"type":                    "ptlmlzpbpbkfbu",
								"typeHandlerVersion":      "crllsludntz",
								"autoUpgradeMinorVersion": true,
								"enableAutomaticUpgrade":  true,
								"settings":                map[string]any{},
								"protectedSettings":       map[string]any{},
								"suppressFailures":        true,
								"protectedSettingsFromKeyVault": map[string]any{
									"secretUrl": "vzkogocyw",
									"sourceVault": map[string]any{
										"id": "lvzxxyypkeqlflftmfn",
									},
								},
								"provisionAfterExtensions": []any{
									"onbtyoeifafiktrkmal",
								},
							},
						},
					},
				},
				"computeApiVersion": "axcvphjtsdjzcwqczcglmq",
				"name":              "dbozdvegpdvqxltqipvmqsfgunpe",
			},
			ResourceOverrides: []map[string]any{
				{
					"plan": map[string]any{
						"name":          "iemasqqkbixbewezyrhnpntjd",
						"publisher":     "bvggylbvfstnscuupuithafvvgc",
						"product":       "bguuzrknnuohugjhernflurpx",
						"promotionCode": "bxgonranwqoryfkhkfaumdgz",
						"version":       "uyxetqmmzvqianqv",
					},
					"zones": []any{
						"wczj",
					},
					"identity": map[string]any{
						"type": "SystemAssigned",
						"userAssignedIdentities": map[string]any{
							"key7": map[string]any{},
						},
					},
					"extendedLocation": map[string]any{
						"name": "gbnxzymbdkxhwjpqkur",
						"type": "EdgeZone",
					},
					"placement": map[string]any{
						"zonePlacementPolicy": "Any",
						"includeZones": []any{
							"inagtbtedobdea",
						},
						"excludeZones": []any{
							"pvvwrhuhdpvbacwmesblpgwzk",
						},
					},
					"tags": map[string]any{
						"key6824": "cefndldgkx",
					},
					"properties": map[string]any{
						"scheduledEventsPolicy": map[string]any{
							"userInitiatedRedeploy": map[string]any{
								"automaticallyApprove": true,
							},
							"userInitiatedReboot": map[string]any{
								"automaticallyApprove": true,
							},
							"scheduledEventsAdditionalPublishingTargets": map[string]any{
								"eventGridAndResourceGraph": map[string]any{
									"enable":                    true,
									"scheduledEventsApiVersion": "lifncbftlkounuyfn",
								},
							},
							"allInstancesDown": map[string]any{
								"automaticallyApprove": true,
							},
						},
						"storageProfile": map[string]any{
							"imageReference": map[string]any{
								"publisher":               "ojlplghybdamadvsrq",
								"offer":                   "uvnqoxhkxefqwbsvjgbswqy",
								"sku":                     "hajdxhjmlkx",
								"version":                 "u",
								"sharedGalleryImageId":    "fz",
								"communityGalleryImageId": "tsfpcq",
								"id":                      "cdbrkpdicibtlliq",
							},
							"osDisk": map[string]any{
								"osType": "Windows",
								"encryptionSettings": map[string]any{
									"diskEncryptionKey": map[string]any{
										"secretUrl": "vzkogocyw",
										"sourceVault": map[string]any{
											"id": "lvzxxyypkeqlflftmfn",
										},
									},
									"keyEncryptionKey": map[string]any{
										"keyUrl": "mjjkvgpoohatw",
										"sourceVault": map[string]any{
											"id": "lvzxxyypkeqlflftmfn",
										},
									},
									"enabled": true,
								},
								"name": "opogpznvctmraoajgizcyrfvpt",
								"vhd": map[string]any{
									"uri": "elpzggtxubepzgjqvdbjmbu",
								},
								"image": map[string]any{
									"uri": "elpzggtxubepzgjqvdbjmbu",
								},
								"caching":                 "None",
								"writeAcceleratorEnabled": true,
								"diffDiskSettings": map[string]any{
									"option":    "Local",
									"placement": "CacheDisk",
								},
								"createOption": "FromImage",
								"diskSizeGB":   2,
								"managedDisk": map[string]any{
									"storageAccountType": "Standard_LRS",
									"diskEncryptionSet": map[string]any{
										"id": "magvkzhdmzhktjlqkkk",
									},
									"securityProfile": map[string]any{
										"securityEncryptionType": "VMGuestStateOnly",
										"diskEncryptionSet": map[string]any{
											"id": "magvkzhdmzhktjlqkkk",
										},
									},
									"id": "numddbqmkxuu",
								},
								"deleteOption": "Delete",
							},
							"dataDisks": []any{
								map[string]any{
									"lun":  7,
									"name": "nbthfzqsxyqvqnbgcljxbwyyoj",
									"vhd": map[string]any{
										"uri": "elpzggtxubepzgjqvdbjmbu",
									},
									"image": map[string]any{
										"uri": "elpzggtxubepzgjqvdbjmbu",
									},
									"caching":                 "None",
									"writeAcceleratorEnabled": true,
									"createOption":            "FromImage",
									"diskSizeGB":              19,
									"managedDisk": map[string]any{
										"storageAccountType": "Standard_LRS",
										"diskEncryptionSet": map[string]any{
											"id": "magvkzhdmzhktjlqkkk",
										},
										"securityProfile": map[string]any{
											"securityEncryptionType": "VMGuestStateOnly",
											"diskEncryptionSet": map[string]any{
												"id": "magvkzhdmzhktjlqkkk",
											},
										},
										"id": "numddbqmkxuu",
									},
									"sourceResource": map[string]any{
										"id": "qnukyordmomtjjqabovlsxl",
									},
									"toBeDetached": true,
									"detachOption": "ForceDetach",
									"deleteOption": "Delete",
								},
							},
							"diskControllerType": "SCSI",
						},
						"hardwareProfile": map[string]any{
							"vmSize": "szrnjqwbruz",
							"vmSizeProperties": map[string]any{
								"vCpusAvailable": 24,
								"vCpusPerCore":   6,
							},
						},
						"additionalCapabilities": map[string]any{
							"ultraSSDEnabled":    true,
							"hibernationEnabled": true,
						},
						"osProfile": map[string]any{
							"computerName":  "bplxnfp",
							"adminUsername": "fxzbi",
							"adminPassword": "<a-password-goes-here>",
							"customData":    "hbdlirohsgnbrahscboc",
							"windowsConfiguration": map[string]any{
								"provisionVMAgent":       true,
								"enableAutomaticUpdates": true,
								"timeZone":               "t",
								"additionalUnattendContent": []any{
									map[string]any{
										"passName":      "OobeSystem",
										"componentName": "Microsoft-Windows-Shell-Setup",
										"settingName":   "AutoLogon",
										"content":       "rguazthnx",
									},
								},
								"patchSettings": map[string]any{
									"patchMode":         "Manual",
									"enableHotpatching": true,
									"assessmentMode":    "ImageDefault",
									"automaticByPlatformSettings": map[string]any{
										"rebootSetting": "Unknown",
										"bypassPlatformSafetyChecksOnUserSchedule": true,
									},
								},
								"winRM": map[string]any{
									"listeners": []any{
										map[string]any{
											"protocol":       "Http",
											"certificateUrl": "quhfapfpyeeocwvwtvuggoqqwt",
										},
									},
								},
							},
							"linuxConfiguration": map[string]any{
								"disablePasswordAuthentication": true,
								"ssh": map[string]any{
									"publicKeys": []any{
										map[string]any{
											"path":    "mrdfxnfjazxog",
											"keyData": "wfhrknkehgesontscqyrewfmhgwt",
										},
									},
								},
								"provisionVMAgent": true,
								"patchSettings": map[string]any{
									"patchMode":      "ImageDefault",
									"assessmentMode": "ImageDefault",
									"automaticByPlatformSettings": map[string]any{
										"rebootSetting": "Unknown",
										"bypassPlatformSafetyChecksOnUserSchedule": true,
									},
								},
								"enableVMAgentPlatformUpdates": true,
							},
							"secrets": []any{
								map[string]any{
									"sourceVault": map[string]any{
										"id": "lvzxxyypkeqlflftmfn",
									},
									"vaultCertificates": []any{
										map[string]any{
											"certificateUrl":   "crgbpfdvlohwkupdjp",
											"certificateStore": "hyx",
										},
									},
								},
							},
							"allowExtensionOperations":    true,
							"requireGuestProvisionSignal": true,
						},
						"networkProfile": map[string]any{
							"networkInterfaces": []any{
								map[string]any{
									"properties": map[string]any{
										"primary":      true,
										"deleteOption": "Delete",
									},
									"id": "ymfxctb",
								},
							},
							"networkApiVersion": "2020-11-01",
							"networkInterfaceConfigurations": []any{
								map[string]any{
									"name": "qrkzoctmzjketostzabnra",
									"properties": map[string]any{
										"primary":                     true,
										"deleteOption":                "Delete",
										"enableAcceleratedNetworking": true,
										"disableTcpStateTracking":     true,
										"enableFpga":                  true,
										"enableIPForwarding":          true,
										"networkSecurityGroup": map[string]any{
											"id": "lvzxxyypkeqlflftmfn",
										},
										"dnsSettings": map[string]any{
											"dnsServers": []any{
												"tqcqopnanyyiavfwhqbkarxtrfqbww",
											},
										},
										"ipConfigurations": []any{
											map[string]any{
												"name": "gqymuvgzzfmxqvdadx",
												"properties": map[string]any{
													"subnet": map[string]any{
														"id": "lvzxxyypkeqlflftmfn",
													},
													"primary": true,
													"publicIPAddressConfiguration": map[string]any{
														"name": "cwxsqjijtwbsyqdwht",
														"properties": map[string]any{
															"idleTimeoutInMinutes": 17,
															"deleteOption":         "Delete",
															"dnsSettings": map[string]any{
																"domainNameLabel":      "fampou",
																"domainNameLabelScope": "TenantReuse",
															},
															"ipTags": []any{
																map[string]any{
																	"ipTagType": "hkjoxhqadudjartwooezaxl",
																	"tag":       "xywunkjglkmmwfpf",
																},
															},
															"publicIPPrefix": map[string]any{
																"id": "lvzxxyypkeqlflftmfn",
															},
															"publicIPAddressVersion":   "IPv4",
															"publicIPAllocationMethod": "Dynamic",
														},
														"sku": map[string]any{
															"name": "Basic",
															"tier": "Regional",
														},
														"tags": map[string]any{
															"key5442": "qhpwpnylvmdthxazhxamnbhdfpf",
														},
													},
													"privateIPAddressVersion": "IPv4",
													"applicationSecurityGroups": []any{
														map[string]any{
															"id": "lvzxxyypkeqlflftmfn",
														},
													},
													"applicationGatewayBackendAddressPools": []any{
														map[string]any{
															"id": "lvzxxyypkeqlflftmfn",
														},
													},
													"loadBalancerBackendAddressPools": []any{
														map[string]any{
															"id": "lvzxxyypkeqlflftmfn",
														},
													},
												},
											},
										},
										"dscpConfiguration": map[string]any{
											"id": "lvzxxyypkeqlflftmfn",
										},
										"auxiliaryMode": "None",
										"auxiliarySku":  "None",
									},
									"tags": map[string]any{
										"key9436": "bjbadzbfvpszbsickv",
									},
								},
							},
						},
						"securityProfile": map[string]any{
							"uefiSettings": map[string]any{
								"secureBootEnabled": true,
								"vTpmEnabled":       true,
							},
							"encryptionAtHost": true,
							"securityType":     "TrustedLaunch",
							"encryptionIdentity": map[string]any{
								"userAssignedIdentityResourceId": "tnajlgbwcepmhytzb",
							},
							"proxyAgentSettings": map[string]any{
								"enabled":          true,
								"mode":             "Audit",
								"keyIncarnationId": 4,
								"wireServer": map[string]any{
									"mode":                                "Audit",
									"inVMAccessControlProfileReferenceId": "xvlzroy",
								},
								"imds": map[string]any{
									"mode":                                "Audit",
									"inVMAccessControlProfileReferenceId": "xvlzroy",
								},
								"addProxyAgentExtension": true,
							},
						},
						"diagnosticsProfile": map[string]any{
							"bootDiagnostics": map[string]any{
								"enabled":    true,
								"storageUri": "pxuhtzehlfsqolbdleirgj",
							},
						},
						"licenseType":          "ymwuemwuntbignqyvzqflvjpcdus",
						"extensionsTimeBudget": "dnyqmcijikzkltjav",
						"scheduledEventsProfile": map[string]any{
							"terminateNotificationProfile": map[string]any{
								"notBeforeTimeout": "owbwifqrlsdmm",
								"enable":           true,
							},
							"osImageNotificationProfile": map[string]any{
								"notBeforeTimeout": "ataqykjdakdvyyzdspaqnhd",
								"enable":           true,
							},
						},
						"userData": "nwjvxe",
						"capacityReservation": map[string]any{
							"capacityReservationGroup": map[string]any{
								"id": "lvzxxyypkeqlflftmfn",
							},
						},
						"applicationProfile": map[string]any{
							"galleryApplications": []any{
								map[string]any{
									"tags":                            "cmygipvpkegyclvpznfu",
									"order":                           8,
									"packageReferenceId":              "afrfkjdrtzftmwramfyu",
									"configurationReference":          "nmfaspclhidtznslsps",
									"treatFailureAsDeploymentFailure": true,
									"enableAutomaticUpgrade":          true,
								},
							},
						},
						"vmExtensions": []any{
							map[string]any{
								"name": "jkpmcxwuahpzwkvexgzpypk",
								"properties": map[string]any{
									"forceUpdateTag":          "dockqxgatsfzhctxrncuw",
									"publisher":               "qesyfldbfoaexyoywhcxafdtdwcg",
									"type":                    "ptlmlzpbpbkfbu",
									"typeHandlerVersion":      "crllsludntz",
									"autoUpgradeMinorVersion": true,
									"enableAutomaticUpgrade":  true,
									"settings":                map[string]any{},
									"protectedSettings":       map[string]any{},
									"suppressFailures":        true,
									"protectedSettingsFromKeyVault": map[string]any{
										"secretUrl": "vzkogocyw",
										"sourceVault": map[string]any{
											"id": "lvzxxyypkeqlflftmfn",
										},
									},
									"provisionAfterExtensions": []any{
										"onbtyoeifafiktrkmal",
									},
								},
							},
						},
					},
					"computeApiVersion": "axcvphjtsdjzcwqczcglmq",
					"name":              "dbozdvegpdvqxltqipvmqsfgunpe",
				},
			},
			ResourceCount:  to.Ptr[int32](23),
			ResourcePrefix: to.Ptr("flivkboavfhjuiucwdjof"),
		},
		ExecutionParameters: &armbulkactions.ExecutionParameters{
			RetryPolicy: &armbulkactions.RetryPolicy{
				RetryCount:           to.Ptr[int32](2),
				RetryWindowInMinutes: to.Ptr[int32](19),
				OnFailureAction:      to.Ptr(armbulkactions.ResourceOperationTypeUnknown),
			},
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armbulkactions.VirtualMachineBulkOperationsClientBulkCreateOperationResponse{
	// 	CreateResourceOperationResponse: armbulkactions.CreateResourceOperationResponse{
	// 		Description: to.Ptr("Create Resource request"),
	// 		Type: to.Ptr("VirtualMachines"),
	// 		Location: to.Ptr("useast2euap"),
	// 		Results: []*armbulkactions.ResourceOperation{
	// 			{
	// 				ResourceID: to.Ptr("/subscriptions/12345678-1234-1234-1234-123456789012/resourceGroups/myResourceGroup/providers/Microsoft.Compute/virtualMachines/myVM"),
	// 				ErrorCode: to.Ptr("TestErrorCode"),
	// 				ErrorDetails: to.Ptr("Test error details"),
	// 				Operation: &armbulkactions.ResourceOperationDetails{
	// 					OperationID: to.Ptr("198fe806-a50d-4a3f-95cf-af162ac59599"),
	// 					ResourceID: to.Ptr("/subscriptions/12345678-1234-1234-1234-123456789012/resourceGroups/myResourceGroup/providers/Microsoft.Compute/virtualMachines/myVM"),
	// 					OpType: to.Ptr(armbulkactions.ResourceOperationTypeCreate),
	// 					SubscriptionID: to.Ptr("1FBA3C66-5C9C-4391-B72F-9F52735FC9F2"),
	// 					Deadline: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2026-06-11T19:35:45.098Z"); return t}()),
	// 					DeadlineType: to.Ptr(armbulkactions.DeadlineTypeInitiateAt),
	// 					State: to.Ptr(armbulkactions.OperationStatePendingScheduling),
	// 					Timezone: to.Ptr("UTC"),
	// 					ResourceOperationError: &armbulkactions.ResourceOperationError{
	// 						ErrorCode: to.Ptr("TestErrorCode"),
	// 						ErrorDetails: to.Ptr("Test error details"),
	// 					},
	// 					FallbackOperationInfo: &armbulkactions.FallbackOperationInfo{
	// 						LastOpType: to.Ptr(armbulkactions.ResourceOperationTypeDeallocate),
	// 						Status: to.Ptr("succeeded"),
	// 						Error: &armbulkactions.ResourceOperationError{
	// 							ErrorCode: to.Ptr("TestErrorCode"),
	// 							ErrorDetails: to.Ptr("Test error details"),
	// 						},
	// 					},
	// 					CompletedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2026-06-11T19:35:45.098Z"); return t}()),
	// 					RetryPolicy: &armbulkactions.RetryPolicy{
	// 						RetryCount: to.Ptr[int32](2),
	// 						RetryWindowInMinutes: to.Ptr[int32](19),
	// 						OnFailureAction: to.Ptr(armbulkactions.ResourceOperationTypeUnknown),
	// 					},
	// 					ResourceNotificationDetails: &armbulkactions.ResourceNotificationDetails{
	// 						ResourceContext: to.Ptr(""),
	// 					},
	// 				},
	// 			},
	// 		},
	// 	},
	// }
}
