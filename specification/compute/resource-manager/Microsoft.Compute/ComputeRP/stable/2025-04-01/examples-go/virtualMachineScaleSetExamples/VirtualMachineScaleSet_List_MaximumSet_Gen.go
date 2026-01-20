package armcompute_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/compute/armcompute/v7"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/ccea12be6cd5e64743499d46f7a9c8b60df52db1/specification/compute/resource-manager/Microsoft.Compute/ComputeRP/stable/2025-04-01/examples/virtualMachineScaleSetExamples/VirtualMachineScaleSet_List_MaximumSet_Gen.json
func ExampleVirtualMachineScaleSetsClient_NewListPager_virtualMachineScaleSetListMaximumSetGen() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcompute.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewVirtualMachineScaleSetsClient().NewListPager("rgcompute", nil)
	for pager.More() {
		page, err := pager.NextPage(ctx)
		if err != nil {
			log.Fatalf("failed to advance page: %v", err)
		}
		for _, v := range page.Value {
			// You could use page here. We use blank identifier for just demo purposes.
			_ = v
		}
		// If the HTTP response code is 200 as defined in example definition, your page structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
		// page.VirtualMachineScaleSetListResult = armcompute.VirtualMachineScaleSetListResult{
		// 	Value: []*armcompute.VirtualMachineScaleSet{
		// 		{
		// 			Name: to.Ptr("{virtualMachineScaleSetName}"),
		// 			Type: to.Ptr("Microsoft.Compute/virtualMachineScaleSets"),
		// 			ID: to.Ptr("/subscriptions/{subscription-id}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/virtualMachineScaleSets/{virtualMachineScaleSetName}"),
		// 			Location: to.Ptr("eastus"),
		// 			Tags: map[string]*string{
		// 			},
		// 			ExtendedLocation: &armcompute.ExtendedLocation{
		// 				Name: to.Ptr("aaaaaaaaaaaaaaaaaaaaa"),
		// 				Type: to.Ptr(armcompute.ExtendedLocationTypesEdgeZone),
		// 			},
		// 			Identity: &armcompute.VirtualMachineScaleSetIdentity{
		// 				Type: to.Ptr(armcompute.ResourceIdentityTypeSystemAssigned),
		// 				PrincipalID: to.Ptr("aaaaaaaaaaaaaaa"),
		// 				TenantID: to.Ptr("aaaaaaaaaaaaaaaa"),
		// 				UserAssignedIdentities: map[string]*armcompute.UserAssignedIdentitiesValue{
		// 					"key3951": &armcompute.UserAssignedIdentitiesValue{
		// 						ClientID: to.Ptr("aaaaaaaaaaaaaaaaaaaaaaaaaaa"),
		// 						PrincipalID: to.Ptr("aaaa"),
		// 					},
		// 				},
		// 			},
		// 			Plan: &armcompute.Plan{
		// 				Name: to.Ptr("aaaaaaaaaa"),
		// 				Product: to.Ptr("aaaaaaaaaaaaaaaaaaaa"),
		// 				PromotionCode: to.Ptr("aaaaaaaaaaaaaaaaaaaa"),
		// 				Publisher: to.Ptr("aaaaaaaaaaaaaaaaaaaaaa"),
		// 			},
		// 			Properties: &armcompute.VirtualMachineScaleSetProperties{
		// 				AdditionalCapabilities: &armcompute.AdditionalCapabilities{
		// 					HibernationEnabled: to.Ptr(true),
		// 					UltraSSDEnabled: to.Ptr(true),
		// 				},
		// 				AutomaticRepairsPolicy: &armcompute.AutomaticRepairsPolicy{
		// 					Enabled: to.Ptr(true),
		// 					GracePeriod: to.Ptr("aaaaaaaaaaaaaaaaaaaaaaaaa"),
		// 				},
		// 				DoNotRunExtensionsOnOverprovisionedVMs: to.Ptr(true),
		// 				HostGroup: &armcompute.SubResource{
		// 					ID: to.Ptr("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/availabilitySets/{availabilitySetName}"),
		// 				},
		// 				OrchestrationMode: to.Ptr(armcompute.OrchestrationModeUniform),
		// 				Overprovision: to.Ptr(true),
		// 				PlatformFaultDomainCount: to.Ptr[int32](1),
		// 				ProvisioningState: to.Ptr("Succeeded"),
		// 				ProximityPlacementGroup: &armcompute.SubResource{
		// 					ID: to.Ptr("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/availabilitySets/{availabilitySetName}"),
		// 				},
		// 				ScaleInPolicy: &armcompute.ScaleInPolicy{
		// 					ForceDeletion: to.Ptr(true),
		// 					Rules: []*armcompute.VirtualMachineScaleSetScaleInRules{
		// 						to.Ptr(armcompute.VirtualMachineScaleSetScaleInRulesDefault)},
		// 					},
		// 					SinglePlacementGroup: to.Ptr(true),
		// 					SpotRestorePolicy: &armcompute.SpotRestorePolicy{
		// 						Enabled: to.Ptr(true),
		// 						RestoreTimeout: to.Ptr("aaaaaaaaaa"),
		// 					},
		// 					UniqueID: to.Ptr("aaaaaaaaaaaaaaaaaaaaaaa"),
		// 					UpgradePolicy: &armcompute.UpgradePolicy{
		// 						AutomaticOSUpgradePolicy: &armcompute.AutomaticOSUpgradePolicy{
		// 							DisableAutomaticRollback: to.Ptr(true),
		// 							EnableAutomaticOSUpgrade: to.Ptr(true),
		// 							OSRollingUpgradeDeferral: to.Ptr(true),
		// 							UseRollingUpgradePolicy: to.Ptr(true),
		// 						},
		// 						Mode: to.Ptr(armcompute.UpgradeModeAutomatic),
		// 						RollingUpgradePolicy: &armcompute.RollingUpgradePolicy{
		// 							EnableCrossZoneUpgrade: to.Ptr(true),
		// 							MaxBatchInstancePercent: to.Ptr[int32](49),
		// 							MaxSurge: to.Ptr(true),
		// 							MaxUnhealthyInstancePercent: to.Ptr[int32](81),
		// 							MaxUnhealthyUpgradedInstancePercent: to.Ptr[int32](98),
		// 							PauseTimeBetweenBatches: to.Ptr("aaaaaaaaaaaaaaa"),
		// 							PrioritizeUnhealthyInstances: to.Ptr(true),
		// 							RollbackFailedInstancesOnPolicyBreach: to.Ptr(true),
		// 						},
		// 					},
		// 					VirtualMachineProfile: &armcompute.VirtualMachineScaleSetVMProfile{
		// 						ApplicationProfile: &armcompute.ApplicationProfile{
		// 							GalleryApplications: []*armcompute.VMGalleryApplication{
		// 								{
		// 									ConfigurationReference: to.Ptr("aaaaa"),
		// 									Order: to.Ptr[int32](29),
		// 									PackageReferenceID: to.Ptr("aaaaaaaaaa"),
		// 									Tags: to.Ptr("aaaaaaaaaaa"),
		// 							}},
		// 						},
		// 						BillingProfile: &armcompute.BillingProfile{
		// 							MaxPrice: to.Ptr[float64](4),
		// 						},
		// 						CapacityReservation: &armcompute.CapacityReservationProfile{
		// 							CapacityReservationGroup: &armcompute.SubResource{
		// 								ID: to.Ptr("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/availabilitySets/{availabilitySetName}"),
		// 							},
		// 						},
		// 						DiagnosticsProfile: &armcompute.DiagnosticsProfile{
		// 							BootDiagnostics: &armcompute.BootDiagnostics{
		// 								Enabled: to.Ptr(true),
		// 								StorageURI: to.Ptr("aaaaaaaaaaaaaaaaaaa"),
		// 							},
		// 						},
		// 						EvictionPolicy: to.Ptr(armcompute.VirtualMachineEvictionPolicyTypesDeallocate),
		// 						ExtensionProfile: &armcompute.VirtualMachineScaleSetExtensionProfile{
		// 							ExtensionsTimeBudget: to.Ptr("aaaaaaaaaaaaaaaaaaaaaaaaaa"),
		// 							Extensions: []*armcompute.VirtualMachineScaleSetExtension{
		// 								{
		// 									ID: to.Ptr("aaaaaaaaaaaaaaaaaaaaaa"),
		// 									Name: to.Ptr("aa"),
		// 									Type: to.Ptr("aaaaa"),
		// 									Properties: &armcompute.VirtualMachineScaleSetExtensionProperties{
		// 										Type: to.Ptr("aaaaaaaa"),
		// 										AutoUpgradeMinorVersion: to.Ptr(true),
		// 										EnableAutomaticUpgrade: to.Ptr(true),
		// 										ForceUpdateTag: to.Ptr("aaaaaaaaaaaaaaaaaaaaaaaaaaa"),
		// 										ProtectedSettings: map[string]any{
		// 										},
		// 										ProvisionAfterExtensions: []*string{
		// 											to.Ptr("aaaaaaaaaaaaaa")},
		// 											ProvisioningState: to.Ptr("Succeeded"),
		// 											Publisher: to.Ptr("aaaaaaaaaaaaa"),
		// 											Settings: map[string]any{
		// 											},
		// 											SuppressFailures: to.Ptr(true),
		// 											TypeHandlerVersion: to.Ptr("aaaaaaaaaaaaaaaaaaaaaaaaaaa"),
		// 										},
		// 								}},
		// 							},
		// 							LicenseType: to.Ptr("aaaa"),
		// 							NetworkProfile: &armcompute.VirtualMachineScaleSetNetworkProfile{
		// 								HealthProbe: &armcompute.APIEntityReference{
		// 									ID: to.Ptr("/subscriptions/{subscription-id}/resourceGroups/myResourceGroup/providers/Microsoft.Compute/restorePointCollections/rpcName/restorePoints/restorePointName/diskRestorePoints/testingexcludedisk_OsDisk_1_74cdaedcea50483d9833c96adefa100f_22b4bdfe-6c54-4f72-84d8-85d8860f0c57"),
		// 								},
		// 								NetworkAPIVersion: to.Ptr(armcompute.NetworkAPIVersionTwoThousandTwenty1101),
		// 								NetworkInterfaceConfigurations: []*armcompute.VirtualMachineScaleSetNetworkConfiguration{
		// 									{
		// 										Name: to.Ptr("myNic"),
		// 										Properties: &armcompute.VirtualMachineScaleSetNetworkConfigurationProperties{
		// 											DeleteOption: to.Ptr(armcompute.DeleteOptionsDelete),
		// 											DNSSettings: &armcompute.VirtualMachineScaleSetNetworkConfigurationDNSSettings{
		// 												DNSServers: []*string{
		// 													to.Ptr("aaaaaaaaaaaa")},
		// 												},
		// 												EnableAcceleratedNetworking: to.Ptr(true),
		// 												EnableFpga: to.Ptr(true),
		// 												EnableIPForwarding: to.Ptr(true),
		// 												IPConfigurations: []*armcompute.VirtualMachineScaleSetIPConfiguration{
		// 													{
		// 														Name: to.Ptr("myIPConfig"),
		// 														Properties: &armcompute.VirtualMachineScaleSetIPConfigurationProperties{
		// 															ApplicationGatewayBackendAddressPools: []*armcompute.SubResource{
		// 																{
		// 																	ID: to.Ptr("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/availabilitySets/{availabilitySetName}"),
		// 															}},
		// 															ApplicationSecurityGroups: []*armcompute.SubResource{
		// 																{
		// 																	ID: to.Ptr("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/availabilitySets/{availabilitySetName}"),
		// 															}},
		// 															LoadBalancerBackendAddressPools: []*armcompute.SubResource{
		// 																{
		// 																	ID: to.Ptr("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/availabilitySets/{availabilitySetName}"),
		// 															}},
		// 															LoadBalancerInboundNatPools: []*armcompute.SubResource{
		// 																{
		// 																	ID: to.Ptr("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/availabilitySets/{availabilitySetName}"),
		// 															}},
		// 															Primary: to.Ptr(true),
		// 															PrivateIPAddressVersion: to.Ptr(armcompute.IPVersionIPv4),
		// 															PublicIPAddressConfiguration: &armcompute.VirtualMachineScaleSetPublicIPAddressConfiguration{
		// 																Name: to.Ptr("aaaaaaaaaaaaaaaaaa"),
		// 																Properties: &armcompute.VirtualMachineScaleSetPublicIPAddressConfigurationProperties{
		// 																	DeleteOption: to.Ptr(armcompute.DeleteOptionsDelete),
		// 																	DNSSettings: &armcompute.VirtualMachineScaleSetPublicIPAddressConfigurationDNSSettings{
		// 																		DomainNameLabel: to.Ptr("aaaaaaaaaaaaaaaaaa"),
		// 																	},
		// 																	IdleTimeoutInMinutes: to.Ptr[int32](18),
		// 																	IPTags: []*armcompute.VirtualMachineScaleSetIPTag{
		// 																		{
		// 																			IPTagType: to.Ptr("aaaaaaa"),
		// 																			Tag: to.Ptr("aaaaaaaaaaaaaaaaaaaaaaaaaaa"),
		// 																	}},
		// 																	PublicIPAddressVersion: to.Ptr(armcompute.IPVersionIPv4),
		// 																	PublicIPPrefix: &armcompute.SubResource{
		// 																		ID: to.Ptr("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/availabilitySets/{availabilitySetName}"),
		// 																	},
		// 																},
		// 																SKU: &armcompute.PublicIPAddressSKU{
		// 																	Name: to.Ptr(armcompute.PublicIPAddressSKUNameBasic),
		// 																	Tier: to.Ptr(armcompute.PublicIPAddressSKUTierRegional),
		// 																},
		// 															},
		// 															Subnet: &armcompute.APIEntityReference{
		// 																ID: to.Ptr("/subscriptions/{subscription-id}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/virtualNetworks/myVNet/subnets/mySubnet"),
		// 															},
		// 														},
		// 												}},
		// 												NetworkSecurityGroup: &armcompute.SubResource{
		// 													ID: to.Ptr("/subscriptions/{subscription-id}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/networkSecurityGroups/myNetworkSecurityGroup"),
		// 												},
		// 												Primary: to.Ptr(true),
		// 											},
		// 									}},
		// 								},
		// 								OSProfile: &armcompute.VirtualMachineScaleSetOSProfile{
		// 									AdminUsername: to.Ptr("admin"),
		// 									ComputerNamePrefix: to.Ptr("{virtualMachineScaleSetName}"),
		// 									CustomData: to.Ptr("aaaaaaaaaaaaaaaaaa"),
		// 									LinuxConfiguration: &armcompute.LinuxConfiguration{
		// 										DisablePasswordAuthentication: to.Ptr(true),
		// 										PatchSettings: &armcompute.LinuxPatchSettings{
		// 											AssessmentMode: to.Ptr(armcompute.LinuxPatchAssessmentModeImageDefault),
		// 											PatchMode: to.Ptr(armcompute.LinuxVMGuestPatchModeImageDefault),
		// 										},
		// 										ProvisionVMAgent: to.Ptr(true),
		// 										SSH: &armcompute.SSHConfiguration{
		// 											PublicKeys: []*armcompute.SSHPublicKey{
		// 												{
		// 													Path: to.Ptr("aaa"),
		// 													KeyData: to.Ptr("aaaaaa"),
		// 											}},
		// 										},
		// 									},
		// 									Secrets: []*armcompute.VaultSecretGroup{
		// 										{
		// 											SourceVault: &armcompute.SubResource{
		// 												ID: to.Ptr("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/availabilitySets/{availabilitySetName}"),
		// 											},
		// 											VaultCertificates: []*armcompute.VaultCertificate{
		// 												{
		// 													CertificateStore: to.Ptr("aaaaaaaaaaaaaaaaaaaaaaaaa"),
		// 													CertificateURL: to.Ptr("aaaaaaa"),
		// 											}},
		// 									}},
		// 									WindowsConfiguration: &armcompute.WindowsConfiguration{
		// 										AdditionalUnattendContent: []*armcompute.AdditionalUnattendContent{
		// 											{
		// 												ComponentName: to.Ptr("Microsoft-Windows-Shell-Setup"),
		// 												Content: to.Ptr("aaaaaaaaaaaaaaaaaaaa"),
		// 												PassName: to.Ptr("OobeSystem"),
		// 												SettingName: to.Ptr(armcompute.SettingNamesAutoLogon),
		// 										}},
		// 										EnableAutomaticUpdates: to.Ptr(true),
		// 										PatchSettings: &armcompute.PatchSettings{
		// 											AssessmentMode: to.Ptr(armcompute.WindowsPatchAssessmentModeImageDefault),
		// 											EnableHotpatching: to.Ptr(true),
		// 											PatchMode: to.Ptr(armcompute.WindowsVMGuestPatchModeManual),
		// 										},
		// 										ProvisionVMAgent: to.Ptr(true),
		// 										TimeZone: to.Ptr("aaaaaaaaaaaaaaaaaaaaaaaaaaaaa"),
		// 										WinRM: &armcompute.WinRMConfiguration{
		// 											Listeners: []*armcompute.WinRMListener{
		// 												{
		// 													CertificateURL: to.Ptr("aaaaaaaaaaaaaaaaaaaaaa"),
		// 													Protocol: to.Ptr(armcompute.ProtocolTypesHTTP),
		// 											}},
		// 										},
		// 									},
		// 								},
		// 								Priority: to.Ptr(armcompute.VirtualMachinePriorityTypesRegular),
		// 								ScheduledEventsProfile: &armcompute.ScheduledEventsProfile{
		// 									TerminateNotificationProfile: &armcompute.TerminateNotificationProfile{
		// 										Enable: to.Ptr(true),
		// 										NotBeforeTimeout: to.Ptr("aa"),
		// 									},
		// 								},
		// 								SecurityProfile: &armcompute.SecurityProfile{
		// 									EncryptionAtHost: to.Ptr(true),
		// 									SecurityType: to.Ptr(armcompute.SecurityTypesTrustedLaunch),
		// 									UefiSettings: &armcompute.UefiSettings{
		// 										SecureBootEnabled: to.Ptr(true),
		// 										VTpmEnabled: to.Ptr(true),
		// 									},
		// 								},
		// 								StorageProfile: &armcompute.VirtualMachineScaleSetStorageProfile{
		// 									DataDisks: []*armcompute.VirtualMachineScaleSetDataDisk{
		// 									},
		// 									ImageReference: &armcompute.ImageReference{
		// 										ID: to.Ptr("aaaaaaaaaa"),
		// 										ExactVersion: to.Ptr("aaaaaaaaaaaaaaaaaaaaaaa"),
		// 										Offer: to.Ptr("databricks"),
		// 										Publisher: to.Ptr("azuredatabricks"),
		// 										SharedGalleryImageID: to.Ptr("aaaaaaaaaaaaaaaaaaaaaa"),
		// 										SKU: to.Ptr("databricksworker"),
		// 										Version: to.Ptr("3.15.2"),
		// 									},
		// 									OSDisk: &armcompute.VirtualMachineScaleSetOSDisk{
		// 										Name: to.Ptr("aaaaaaaaaaaaaaaaaaaaaaa"),
		// 										Caching: to.Ptr(armcompute.CachingTypesReadWrite),
		// 										CreateOption: to.Ptr(armcompute.DiskCreateOptionTypesFromImage),
		// 										DiffDiskSettings: &armcompute.DiffDiskSettings{
		// 											Option: to.Ptr(armcompute.DiffDiskOptionsLocal),
		// 											Placement: to.Ptr(armcompute.DiffDiskPlacementCacheDisk),
		// 										},
		// 										DiskSizeGB: to.Ptr[int32](30),
		// 										Image: &armcompute.VirtualHardDisk{
		// 											URI: to.Ptr("https://{storageAccountName}.blob.core.windows.net/{containerName}/{vhdName}.vhd"),
		// 										},
		// 										ManagedDisk: &armcompute.VirtualMachineScaleSetManagedDiskParameters{
		// 											DiskEncryptionSet: &armcompute.DiskEncryptionSetParameters{
		// 												ID: to.Ptr("aaaaaaaaaaaa"),
		// 											},
		// 											StorageAccountType: to.Ptr(armcompute.StorageAccountTypesPremiumLRS),
		// 										},
		// 										OSType: to.Ptr(armcompute.OperatingSystemTypesWindows),
		// 										VhdContainers: []*string{
		// 											to.Ptr("aaaaaaaaaaaaaaaaaaaaaaaaaaaaa")},
		// 											WriteAcceleratorEnabled: to.Ptr(true),
		// 										},
		// 									},
		// 									UserData: to.Ptr("aaaaaaaaaaaaaaaaaaaaaaaaa"),
		// 								},
		// 								ZoneBalance: to.Ptr(true),
		// 							},
		// 							SKU: &armcompute.SKU{
		// 								Name: to.Ptr("Standard_D2s_v3"),
		// 								Capacity: to.Ptr[int64](4),
		// 								Tier: to.Ptr("Standard"),
		// 							},
		// 							Zones: []*string{
		// 								to.Ptr("aaaaaaa")},
		// 							},
		// 							{
		// 								Name: to.Ptr("{virtualMachineScaleSetName}"),
		// 								Type: to.Ptr("Microsoft.Compute/virtualMachineScaleSets"),
		// 								ID: to.Ptr("/subscriptions/{subscription-id}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/virtualMachineScaleSets/{virtualMachineScaleSetName}1"),
		// 								Location: to.Ptr("eastus"),
		// 								Tags: map[string]*string{
		// 								},
		// 								ExtendedLocation: &armcompute.ExtendedLocation{
		// 									Name: to.Ptr("aaaaaaaaaaaaaaaaaaaaa"),
		// 									Type: to.Ptr(armcompute.ExtendedLocationTypesEdgeZone),
		// 								},
		// 								Identity: &armcompute.VirtualMachineScaleSetIdentity{
		// 									Type: to.Ptr(armcompute.ResourceIdentityTypeSystemAssigned),
		// 									PrincipalID: to.Ptr("aaaaaaaaaaaaaaa"),
		// 									TenantID: to.Ptr("aaaaaaaaaaaaaaaa"),
		// 									UserAssignedIdentities: map[string]*armcompute.UserAssignedIdentitiesValue{
		// 										"key3951": &armcompute.UserAssignedIdentitiesValue{
		// 											ClientID: to.Ptr("aaaaaaaaaaaaaaaaaaaaaaaaaaa"),
		// 											PrincipalID: to.Ptr("aaaa"),
		// 										},
		// 									},
		// 								},
		// 								Plan: &armcompute.Plan{
		// 									Name: to.Ptr("aaaaaaaaaa"),
		// 									Product: to.Ptr("aaaaaaaaaaaaaaaaaaaa"),
		// 									PromotionCode: to.Ptr("aaaaaaaaaaaaaaaaaaaa"),
		// 									Publisher: to.Ptr("aaaaaaaaaaaaaaaaaaaaaa"),
		// 								},
		// 								Properties: &armcompute.VirtualMachineScaleSetProperties{
		// 									AdditionalCapabilities: &armcompute.AdditionalCapabilities{
		// 										HibernationEnabled: to.Ptr(true),
		// 										UltraSSDEnabled: to.Ptr(true),
		// 									},
		// 									AutomaticRepairsPolicy: &armcompute.AutomaticRepairsPolicy{
		// 										Enabled: to.Ptr(true),
		// 										GracePeriod: to.Ptr("aaaaaaaaaaaaaaaaaaaaaaaaa"),
		// 									},
		// 									DoNotRunExtensionsOnOverprovisionedVMs: to.Ptr(true),
		// 									HostGroup: &armcompute.SubResource{
		// 										ID: to.Ptr("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/availabilitySets/{availabilitySetName}"),
		// 									},
		// 									OrchestrationMode: to.Ptr(armcompute.OrchestrationModeUniform),
		// 									Overprovision: to.Ptr(true),
		// 									PlatformFaultDomainCount: to.Ptr[int32](1),
		// 									ProvisioningState: to.Ptr("Succeeded"),
		// 									ProximityPlacementGroup: &armcompute.SubResource{
		// 										ID: to.Ptr("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/availabilitySets/{availabilitySetName}"),
		// 									},
		// 									ScaleInPolicy: &armcompute.ScaleInPolicy{
		// 										ForceDeletion: to.Ptr(true),
		// 										Rules: []*armcompute.VirtualMachineScaleSetScaleInRules{
		// 											to.Ptr(armcompute.VirtualMachineScaleSetScaleInRulesDefault)},
		// 										},
		// 										SinglePlacementGroup: to.Ptr(true),
		// 										SpotRestorePolicy: &armcompute.SpotRestorePolicy{
		// 											Enabled: to.Ptr(true),
		// 											RestoreTimeout: to.Ptr("aaaaaaaaaa"),
		// 										},
		// 										UniqueID: to.Ptr("aaaaa"),
		// 										UpgradePolicy: &armcompute.UpgradePolicy{
		// 											AutomaticOSUpgradePolicy: &armcompute.AutomaticOSUpgradePolicy{
		// 												DisableAutomaticRollback: to.Ptr(true),
		// 												EnableAutomaticOSUpgrade: to.Ptr(true),
		// 												OSRollingUpgradeDeferral: to.Ptr(true),
		// 												UseRollingUpgradePolicy: to.Ptr(true),
		// 											},
		// 											Mode: to.Ptr(armcompute.UpgradeModeAutomatic),
		// 											RollingUpgradePolicy: &armcompute.RollingUpgradePolicy{
		// 												EnableCrossZoneUpgrade: to.Ptr(true),
		// 												MaxBatchInstancePercent: to.Ptr[int32](49),
		// 												MaxSurge: to.Ptr(true),
		// 												MaxUnhealthyInstancePercent: to.Ptr[int32](81),
		// 												MaxUnhealthyUpgradedInstancePercent: to.Ptr[int32](98),
		// 												PauseTimeBetweenBatches: to.Ptr("aaaaaaaaaaaaaaa"),
		// 												PrioritizeUnhealthyInstances: to.Ptr(true),
		// 											},
		// 										},
		// 										VirtualMachineProfile: &armcompute.VirtualMachineScaleSetVMProfile{
		// 											ApplicationProfile: &armcompute.ApplicationProfile{
		// 												GalleryApplications: []*armcompute.VMGalleryApplication{
		// 													{
		// 														ConfigurationReference: to.Ptr("aaaaa"),
		// 														Order: to.Ptr[int32](29),
		// 														PackageReferenceID: to.Ptr("aaaaaaaaaa"),
		// 														Tags: to.Ptr("aaaaaaaaaaa"),
		// 												}},
		// 											},
		// 											BillingProfile: &armcompute.BillingProfile{
		// 												MaxPrice: to.Ptr[float64](4),
		// 											},
		// 											CapacityReservation: &armcompute.CapacityReservationProfile{
		// 												CapacityReservationGroup: &armcompute.SubResource{
		// 													ID: to.Ptr("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/availabilitySets/{availabilitySetName}"),
		// 												},
		// 											},
		// 											DiagnosticsProfile: &armcompute.DiagnosticsProfile{
		// 												BootDiagnostics: &armcompute.BootDiagnostics{
		// 													Enabled: to.Ptr(true),
		// 													StorageURI: to.Ptr("aaaaaaaaaaaaaaaaaaa"),
		// 												},
		// 											},
		// 											EvictionPolicy: to.Ptr(armcompute.VirtualMachineEvictionPolicyTypesDeallocate),
		// 											ExtensionProfile: &armcompute.VirtualMachineScaleSetExtensionProfile{
		// 												ExtensionsTimeBudget: to.Ptr("aaaaaaaaaaaaaaaaaaaaaaaaaa"),
		// 												Extensions: []*armcompute.VirtualMachineScaleSetExtension{
		// 													{
		// 														ID: to.Ptr("aaaaaaaaaaaaaaaaaaaaaa"),
		// 														Name: to.Ptr("aa"),
		// 														Type: to.Ptr("aaaaa"),
		// 														Properties: &armcompute.VirtualMachineScaleSetExtensionProperties{
		// 															Type: to.Ptr("aaaaaaaa"),
		// 															AutoUpgradeMinorVersion: to.Ptr(true),
		// 															EnableAutomaticUpgrade: to.Ptr(true),
		// 															ForceUpdateTag: to.Ptr("aaaaaaaaaaaaaaaaaaaaaaaaaaa"),
		// 															ProtectedSettings: map[string]any{
		// 															},
		// 															ProvisionAfterExtensions: []*string{
		// 																to.Ptr("aaaaaaaaaaaaaa")},
		// 																ProvisioningState: to.Ptr("Succeeded"),
		// 																Publisher: to.Ptr("aaaaaaaaaaaaa"),
		// 																Settings: map[string]any{
		// 																},
		// 																SuppressFailures: to.Ptr(true),
		// 																TypeHandlerVersion: to.Ptr("aaaaaaaaaaaaaaaaaaaaaaaaaaa"),
		// 															},
		// 													}},
		// 												},
		// 												LicenseType: to.Ptr("aaaaaaaaaaaaaaaaaaaaaaa"),
		// 												NetworkProfile: &armcompute.VirtualMachineScaleSetNetworkProfile{
		// 													HealthProbe: &armcompute.APIEntityReference{
		// 														ID: to.Ptr("/subscriptions/{subscription-id}/resourceGroups/myResourceGroup/providers/Microsoft.Compute/restorePointCollections/rpcName/restorePoints/restorePointName/diskRestorePoints/testingexcludedisk_OsDisk_1_74cdaedcea50483d9833c96adefa100f_22b4bdfe-6c54-4f72-84d8-85d8860f0c57"),
		// 													},
		// 													NetworkAPIVersion: to.Ptr(armcompute.NetworkAPIVersionTwoThousandTwenty1101),
		// 													NetworkInterfaceConfigurations: []*armcompute.VirtualMachineScaleSetNetworkConfiguration{
		// 														{
		// 															Name: to.Ptr("myNic1"),
		// 															Properties: &armcompute.VirtualMachineScaleSetNetworkConfigurationProperties{
		// 																DeleteOption: to.Ptr(armcompute.DeleteOptionsDelete),
		// 																DNSSettings: &armcompute.VirtualMachineScaleSetNetworkConfigurationDNSSettings{
		// 																	DNSServers: []*string{
		// 																		to.Ptr("aaaaaaaaaaaa")},
		// 																	},
		// 																	EnableAcceleratedNetworking: to.Ptr(true),
		// 																	EnableFpga: to.Ptr(true),
		// 																	EnableIPForwarding: to.Ptr(true),
		// 																	IPConfigurations: []*armcompute.VirtualMachineScaleSetIPConfiguration{
		// 																		{
		// 																			Name: to.Ptr("myIPConfig"),
		// 																			Properties: &armcompute.VirtualMachineScaleSetIPConfigurationProperties{
		// 																				ApplicationGatewayBackendAddressPools: []*armcompute.SubResource{
		// 																					{
		// 																						ID: to.Ptr("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/availabilitySets/{availabilitySetName}"),
		// 																				}},
		// 																				ApplicationSecurityGroups: []*armcompute.SubResource{
		// 																					{
		// 																						ID: to.Ptr("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/availabilitySets/{availabilitySetName}"),
		// 																				}},
		// 																				LoadBalancerBackendAddressPools: []*armcompute.SubResource{
		// 																					{
		// 																						ID: to.Ptr("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/availabilitySets/{availabilitySetName}"),
		// 																				}},
		// 																				LoadBalancerInboundNatPools: []*armcompute.SubResource{
		// 																					{
		// 																						ID: to.Ptr("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/availabilitySets/{availabilitySetName}"),
		// 																				}},
		// 																				Primary: to.Ptr(true),
		// 																				PrivateIPAddressVersion: to.Ptr(armcompute.IPVersionIPv4),
		// 																				PublicIPAddressConfiguration: &armcompute.VirtualMachineScaleSetPublicIPAddressConfiguration{
		// 																					Name: to.Ptr("aaaaaaaaaaaaaaaaaa"),
		// 																					Properties: &armcompute.VirtualMachineScaleSetPublicIPAddressConfigurationProperties{
		// 																						DeleteOption: to.Ptr(armcompute.DeleteOptionsDelete),
		// 																						DNSSettings: &armcompute.VirtualMachineScaleSetPublicIPAddressConfigurationDNSSettings{
		// 																							DomainNameLabel: to.Ptr("aaaaaaaaaaaaaaaaaa"),
		// 																						},
		// 																						IdleTimeoutInMinutes: to.Ptr[int32](18),
		// 																						IPTags: []*armcompute.VirtualMachineScaleSetIPTag{
		// 																							{
		// 																								IPTagType: to.Ptr("aaaaaaa"),
		// 																								Tag: to.Ptr("aaaaaaaaaaaaaaaaaaaaaaaaaaa"),
		// 																						}},
		// 																						PublicIPAddressVersion: to.Ptr(armcompute.IPVersionIPv4),
		// 																						PublicIPPrefix: &armcompute.SubResource{
		// 																							ID: to.Ptr("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/availabilitySets/{availabilitySetName}"),
		// 																						},
		// 																					},
		// 																					SKU: &armcompute.PublicIPAddressSKU{
		// 																						Name: to.Ptr(armcompute.PublicIPAddressSKUNameBasic),
		// 																						Tier: to.Ptr(armcompute.PublicIPAddressSKUTierRegional),
		// 																					},
		// 																				},
		// 																				Subnet: &armcompute.APIEntityReference{
		// 																					ID: to.Ptr("/subscriptions/{subscription-id}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/virtualNetworks/myVNet/subnets/mySubnet"),
		// 																				},
		// 																			},
		// 																	}},
		// 																	NetworkSecurityGroup: &armcompute.SubResource{
		// 																		ID: to.Ptr("/subscriptions/{subscription-id}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/networkSecurityGroups/myNetworkSecurityGroup"),
		// 																	},
		// 																	Primary: to.Ptr(true),
		// 																},
		// 														}},
		// 													},
		// 													OSProfile: &armcompute.VirtualMachineScaleSetOSProfile{
		// 														AdminUsername: to.Ptr("admin"),
		// 														ComputerNamePrefix: to.Ptr("{virtualMachineScaleSetName}"),
		// 														CustomData: to.Ptr("a"),
		// 														LinuxConfiguration: &armcompute.LinuxConfiguration{
		// 															DisablePasswordAuthentication: to.Ptr(true),
		// 															PatchSettings: &armcompute.LinuxPatchSettings{
		// 																AssessmentMode: to.Ptr(armcompute.LinuxPatchAssessmentModeImageDefault),
		// 																PatchMode: to.Ptr(armcompute.LinuxVMGuestPatchModeImageDefault),
		// 															},
		// 															ProvisionVMAgent: to.Ptr(true),
		// 															SSH: &armcompute.SSHConfiguration{
		// 																PublicKeys: []*armcompute.SSHPublicKey{
		// 																	{
		// 																		Path: to.Ptr("aaa"),
		// 																		KeyData: to.Ptr("aaaaaa"),
		// 																}},
		// 															},
		// 														},
		// 														Secrets: []*armcompute.VaultSecretGroup{
		// 															{
		// 																SourceVault: &armcompute.SubResource{
		// 																	ID: to.Ptr("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/availabilitySets/{availabilitySetName}"),
		// 																},
		// 																VaultCertificates: []*armcompute.VaultCertificate{
		// 																	{
		// 																		CertificateStore: to.Ptr("aaaaaaaaaaaaaaaaaaaaaaaaa"),
		// 																		CertificateURL: to.Ptr("aaaaaaa"),
		// 																}},
		// 														}},
		// 														WindowsConfiguration: &armcompute.WindowsConfiguration{
		// 															AdditionalUnattendContent: []*armcompute.AdditionalUnattendContent{
		// 																{
		// 																	ComponentName: to.Ptr("Microsoft-Windows-Shell-Setup"),
		// 																	Content: to.Ptr("aaaaaaaaaaaaaaaaaaaa"),
		// 																	PassName: to.Ptr("OobeSystem"),
		// 																	SettingName: to.Ptr(armcompute.SettingNamesAutoLogon),
		// 															}},
		// 															EnableAutomaticUpdates: to.Ptr(true),
		// 															PatchSettings: &armcompute.PatchSettings{
		// 																AssessmentMode: to.Ptr(armcompute.WindowsPatchAssessmentModeImageDefault),
		// 																EnableHotpatching: to.Ptr(true),
		// 																PatchMode: to.Ptr(armcompute.WindowsVMGuestPatchModeManual),
		// 															},
		// 															ProvisionVMAgent: to.Ptr(true),
		// 															TimeZone: to.Ptr("aaaaaaaaaaaaaaaaaaaaaaaaaaaaa"),
		// 															WinRM: &armcompute.WinRMConfiguration{
		// 																Listeners: []*armcompute.WinRMListener{
		// 																	{
		// 																		CertificateURL: to.Ptr("aaaaaaaaaaaaaaaaaaaaaa"),
		// 																		Protocol: to.Ptr(armcompute.ProtocolTypesHTTP),
		// 																}},
		// 															},
		// 														},
		// 													},
		// 													Priority: to.Ptr(armcompute.VirtualMachinePriorityTypesRegular),
		// 													ScheduledEventsProfile: &armcompute.ScheduledEventsProfile{
		// 														TerminateNotificationProfile: &armcompute.TerminateNotificationProfile{
		// 															Enable: to.Ptr(true),
		// 															NotBeforeTimeout: to.Ptr("aa"),
		// 														},
		// 													},
		// 													SecurityProfile: &armcompute.SecurityProfile{
		// 														EncryptionAtHost: to.Ptr(true),
		// 														SecurityType: to.Ptr(armcompute.SecurityTypesTrustedLaunch),
		// 														UefiSettings: &armcompute.UefiSettings{
		// 															SecureBootEnabled: to.Ptr(true),
		// 															VTpmEnabled: to.Ptr(true),
		// 														},
		// 													},
		// 													StorageProfile: &armcompute.VirtualMachineScaleSetStorageProfile{
		// 														DataDisks: []*armcompute.VirtualMachineScaleSetDataDisk{
		// 														},
		// 														ImageReference: &armcompute.ImageReference{
		// 															ID: to.Ptr("aaa"),
		// 															ExactVersion: to.Ptr("aa"),
		// 															Offer: to.Ptr("databricks"),
		// 															Publisher: to.Ptr("azuredatabricks"),
		// 															SharedGalleryImageID: to.Ptr("aaaaaaa"),
		// 															SKU: to.Ptr("databricksworker"),
		// 															Version: to.Ptr("3.15.2"),
		// 														},
		// 														OSDisk: &armcompute.VirtualMachineScaleSetOSDisk{
		// 															Name: to.Ptr("a"),
		// 															Caching: to.Ptr(armcompute.CachingTypesReadWrite),
		// 															CreateOption: to.Ptr(armcompute.DiskCreateOptionTypesFromImage),
		// 															DiffDiskSettings: &armcompute.DiffDiskSettings{
		// 																Option: to.Ptr(armcompute.DiffDiskOptionsLocal),
		// 																Placement: to.Ptr(armcompute.DiffDiskPlacementCacheDisk),
		// 															},
		// 															DiskSizeGB: to.Ptr[int32](30),
		// 															Image: &armcompute.VirtualHardDisk{
		// 																URI: to.Ptr("https://{storageAccountName}.blob.core.windows.net/{containerName}/{vhdName}.vhd"),
		// 															},
		// 															ManagedDisk: &armcompute.VirtualMachineScaleSetManagedDiskParameters{
		// 																DiskEncryptionSet: &armcompute.DiskEncryptionSetParameters{
		// 																	ID: to.Ptr("aaaaaaaaaaaa"),
		// 																},
		// 																StorageAccountType: to.Ptr(armcompute.StorageAccountTypesPremiumLRS),
		// 															},
		// 															OSType: to.Ptr(armcompute.OperatingSystemTypesWindows),
		// 															VhdContainers: []*string{
		// 																to.Ptr("aaaaaaaaaaaaaaaa")},
		// 																WriteAcceleratorEnabled: to.Ptr(true),
		// 															},
		// 														},
		// 														UserData: to.Ptr("aaaaaaaaaaaaaaaaaaaaaa"),
		// 													},
		// 													ZoneBalance: to.Ptr(true),
		// 												},
		// 												SKU: &armcompute.SKU{
		// 													Name: to.Ptr("Standard_D2s_v3"),
		// 													Capacity: to.Ptr[int64](4),
		// 													Tier: to.Ptr("Standard"),
		// 												},
		// 												Zones: []*string{
		// 													to.Ptr("aaaaaaaa")},
		// 											}},
		// 										}
	}
}
