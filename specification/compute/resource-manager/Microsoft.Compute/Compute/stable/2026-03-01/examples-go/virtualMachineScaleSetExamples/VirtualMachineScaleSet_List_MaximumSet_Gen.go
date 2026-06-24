package armcompute_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/compute/armcompute/v8"
)

// Generated from example definition: 2026-03-01/virtualMachineScaleSetExamples/VirtualMachineScaleSet_List_MaximumSet_Gen.json
func ExampleVirtualMachineScaleSetsClient_NewListPager_virtualMachineScaleSetListMaximumSetGen() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcompute.NewClientFactory("{subscription-id}", cred, nil)
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
		// page = armcompute.VirtualMachineScaleSetsClientListResponse{
		// 	VirtualMachineScaleSetListResult: armcompute.VirtualMachineScaleSetListResult{
		// 		Value: []*armcompute.VirtualMachineScaleSet{
		// 			{
		// 				Name: to.Ptr("{virtualMachineScaleSetName}"),
		// 				ID: to.Ptr("/subscriptions/{subscription-id}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/virtualMachineScaleSets/{virtualMachineScaleSetName}"),
		// 				Type: to.Ptr("Microsoft.Compute/virtualMachineScaleSets"),
		// 				Location: to.Ptr("eastus"),
		// 				Tags: map[string]*string{
		// 				},
		// 				SKU: &armcompute.SKU{
		// 					Name: to.Ptr("Standard_D2s_v3"),
		// 					Tier: to.Ptr("Standard"),
		// 					Capacity: to.Ptr[int64](4),
		// 				},
		// 				Properties: &armcompute.VirtualMachineScaleSetProperties{
		// 					SinglePlacementGroup: to.Ptr(true),
		// 					UpgradePolicy: &armcompute.UpgradePolicy{
		// 						Mode: to.Ptr(armcompute.UpgradeModeAutomatic),
		// 						AutomaticOSUpgradePolicy: &armcompute.AutomaticOSUpgradePolicy{
		// 							EnableAutomaticOSUpgrade: to.Ptr(true),
		// 							DisableAutomaticRollback: to.Ptr(true),
		// 							UseRollingUpgradePolicy: to.Ptr(true),
		// 							OSRollingUpgradeDeferral: to.Ptr(true),
		// 						},
		// 						RollingUpgradePolicy: &armcompute.RollingUpgradePolicy{
		// 							MaxBatchInstancePercent: to.Ptr[int32](49),
		// 							MaxUnhealthyInstancePercent: to.Ptr[int32](81),
		// 							MaxUnhealthyUpgradedInstancePercent: to.Ptr[int32](98),
		// 							PauseTimeBetweenBatches: to.Ptr("aaaaaaaaaaaaaaa"),
		// 							EnableCrossZoneUpgrade: to.Ptr(true),
		// 							PrioritizeUnhealthyInstances: to.Ptr(true),
		// 							RollbackFailedInstancesOnPolicyBreach: to.Ptr(true),
		// 							MaxSurge: to.Ptr(true),
		// 						},
		// 					},
		// 					VirtualMachineProfile: &armcompute.VirtualMachineScaleSetVMProfile{
		// 						StorageProfile: &armcompute.VirtualMachineScaleSetStorageProfile{
		// 							OSDisk: &armcompute.VirtualMachineScaleSetOSDisk{
		// 								CreateOption: to.Ptr(armcompute.DiskCreateOptionTypesFromImage),
		// 								Caching: to.Ptr(armcompute.CachingTypesReadWrite),
		// 								ManagedDisk: &armcompute.VirtualMachineScaleSetManagedDiskParameters{
		// 									StorageAccountType: to.Ptr(armcompute.StorageAccountTypesPremiumLRS),
		// 									DiskEncryptionSet: &armcompute.DiskEncryptionSetParameters{
		// 										ID: to.Ptr("aaaaaaaaaaaa"),
		// 									},
		// 								},
		// 								DiskSizeGB: to.Ptr[int32](30),
		// 								Name: to.Ptr("aaaaaaaaaaaaaaaaaaaaaaa"),
		// 								WriteAcceleratorEnabled: to.Ptr(true),
		// 								DiffDiskSettings: &armcompute.DiffDiskSettings{
		// 									Option: to.Ptr(armcompute.DiffDiskOptionsLocal),
		// 									Placement: to.Ptr(armcompute.DiffDiskPlacementCacheDisk),
		// 								},
		// 								OSType: to.Ptr(armcompute.OperatingSystemTypesWindows),
		// 								Image: &armcompute.VirtualHardDisk{
		// 									URI: to.Ptr("https://{storageAccountName}.blob.core.windows.net/{containerName}/{vhdName}.vhd"),
		// 								},
		// 								VhdContainers: []*string{
		// 									to.Ptr("aaaaaaaaaaaaaaaaaaaaaaaaaaaaa"),
		// 								},
		// 							},
		// 							ImageReference: &armcompute.ImageReference{
		// 								Publisher: to.Ptr("azuredatabricks"),
		// 								Offer: to.Ptr("databricks"),
		// 								SKU: to.Ptr("databricksworker"),
		// 								Version: to.Ptr("3.15.2"),
		// 								ExactVersion: to.Ptr("aaaaaaaaaaaaaaaaaaaaaaa"),
		// 								SharedGalleryImageID: to.Ptr("aaaaaaaaaaaaaaaaaaaaaa"),
		// 								ID: to.Ptr("aaaaaaaaaa"),
		// 							},
		// 							DataDisks: []*armcompute.VirtualMachineScaleSetDataDisk{
		// 							},
		// 						},
		// 						OSProfile: &armcompute.VirtualMachineScaleSetOSProfile{
		// 							ComputerNamePrefix: to.Ptr("{virtualMachineScaleSetName}"),
		// 							AdminUsername: to.Ptr("admin"),
		// 							LinuxConfiguration: &armcompute.LinuxConfiguration{
		// 								DisablePasswordAuthentication: to.Ptr(true),
		// 								SSH: &armcompute.SSHConfiguration{
		// 									PublicKeys: []*armcompute.SSHPublicKey{
		// 										{
		// 											Path: to.Ptr("aaa"),
		// 											KeyData: to.Ptr("aaaaaa"),
		// 										},
		// 									},
		// 								},
		// 								ProvisionVMAgent: to.Ptr(true),
		// 								PatchSettings: &armcompute.LinuxPatchSettings{
		// 									PatchMode: to.Ptr(armcompute.LinuxVMGuestPatchModeImageDefault),
		// 									AssessmentMode: to.Ptr(armcompute.LinuxPatchAssessmentModeImageDefault),
		// 								},
		// 							},
		// 							CustomData: to.Ptr("aaaaaaaaaaaaaaaaaa"),
		// 							WindowsConfiguration: &armcompute.WindowsConfiguration{
		// 								ProvisionVMAgent: to.Ptr(true),
		// 								EnableAutomaticUpdates: to.Ptr(true),
		// 								TimeZone: to.Ptr("aaaaaaaaaaaaaaaaaaaaaaaaaaaaa"),
		// 								AdditionalUnattendContent: []*armcompute.AdditionalUnattendContent{
		// 									{
		// 										PassName: to.Ptr("OobeSystem"),
		// 										ComponentName: to.Ptr("Microsoft-Windows-Shell-Setup"),
		// 										SettingName: to.Ptr(armcompute.SettingNamesAutoLogon),
		// 										Content: to.Ptr("aaaaaaaaaaaaaaaaaaaa"),
		// 									},
		// 								},
		// 								PatchSettings: &armcompute.PatchSettings{
		// 									PatchMode: to.Ptr(armcompute.WindowsVMGuestPatchModeManual),
		// 									EnableHotpatching: to.Ptr(true),
		// 									AssessmentMode: to.Ptr(armcompute.WindowsPatchAssessmentModeImageDefault),
		// 								},
		// 								WinRM: &armcompute.WinRMConfiguration{
		// 									Listeners: []*armcompute.WinRMListener{
		// 										{
		// 											Protocol: to.Ptr(armcompute.ProtocolTypesHTTP),
		// 											CertificateURL: to.Ptr("aaaaaaaaaaaaaaaaaaaaaa"),
		// 										},
		// 									},
		// 								},
		// 							},
		// 							Secrets: []*armcompute.VaultSecretGroup{
		// 								{
		// 									SourceVault: &armcompute.SubResource{
		// 										ID: to.Ptr("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/availabilitySets/{availabilitySetName}"),
		// 									},
		// 									VaultCertificates: []*armcompute.VaultCertificate{
		// 										{
		// 											CertificateURL: to.Ptr("aaaaaaa"),
		// 											CertificateStore: to.Ptr("aaaaaaaaaaaaaaaaaaaaaaaaa"),
		// 										},
		// 									},
		// 								},
		// 							},
		// 						},
		// 						NetworkProfile: &armcompute.VirtualMachineScaleSetNetworkProfile{
		// 							NetworkInterfaceConfigurations: []*armcompute.VirtualMachineScaleSetNetworkConfiguration{
		// 								{
		// 									Name: to.Ptr("myNic"),
		// 									Properties: &armcompute.VirtualMachineScaleSetNetworkConfigurationProperties{
		// 										Primary: to.Ptr(true),
		// 										IPConfigurations: []*armcompute.VirtualMachineScaleSetIPConfiguration{
		// 											{
		// 												Name: to.Ptr("myIPConfig"),
		// 												Properties: &armcompute.VirtualMachineScaleSetIPConfigurationProperties{
		// 													Primary: to.Ptr(true),
		// 													Subnet: &armcompute.APIEntityReference{
		// 														ID: to.Ptr("/subscriptions/{subscription-id}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/virtualNetworks/myVNet/subnets/mySubnet"),
		// 													},
		// 													PublicIPAddressConfiguration: &armcompute.VirtualMachineScaleSetPublicIPAddressConfiguration{
		// 														Name: to.Ptr("aaaaaaaaaaaaaaaaaa"),
		// 														Properties: &armcompute.VirtualMachineScaleSetPublicIPAddressConfigurationProperties{
		// 															IdleTimeoutInMinutes: to.Ptr[int32](18),
		// 															DNSSettings: &armcompute.VirtualMachineScaleSetPublicIPAddressConfigurationDNSSettings{
		// 																DomainNameLabel: to.Ptr("aaaaaaaaaaaaaaaaaa"),
		// 															},
		// 															IPTags: []*armcompute.VirtualMachineScaleSetIPTag{
		// 																{
		// 																	IPTagType: to.Ptr("aaaaaaa"),
		// 																	Tag: to.Ptr("aaaaaaaaaaaaaaaaaaaaaaaaaaa"),
		// 																},
		// 															},
		// 															PublicIPPrefix: &armcompute.SubResource{
		// 																ID: to.Ptr("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/availabilitySets/{availabilitySetName}"),
		// 															},
		// 															PublicIPAddressVersion: to.Ptr(armcompute.IPVersionIPv4),
		// 															DeleteOption: to.Ptr(armcompute.DeleteOptionsDelete),
		// 														},
		// 														SKU: &armcompute.PublicIPAddressSKU{
		// 															Name: to.Ptr(armcompute.PublicIPAddressSKUNameBasic),
		// 															Tier: to.Ptr(armcompute.PublicIPAddressSKUTierRegional),
		// 														},
		// 													},
		// 													PrivateIPAddressVersion: to.Ptr(armcompute.IPVersionIPv4),
		// 													ApplicationGatewayBackendAddressPools: []*armcompute.SubResource{
		// 														{
		// 															ID: to.Ptr("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/availabilitySets/{availabilitySetName}"),
		// 														},
		// 													},
		// 													ApplicationSecurityGroups: []*armcompute.SubResource{
		// 														{
		// 															ID: to.Ptr("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/availabilitySets/{availabilitySetName}"),
		// 														},
		// 													},
		// 													LoadBalancerBackendAddressPools: []*armcompute.SubResource{
		// 														{
		// 															ID: to.Ptr("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/availabilitySets/{availabilitySetName}"),
		// 														},
		// 													},
		// 													LoadBalancerInboundNatPools: []*armcompute.SubResource{
		// 														{
		// 															ID: to.Ptr("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/availabilitySets/{availabilitySetName}"),
		// 														},
		// 													},
		// 												},
		// 											},
		// 										},
		// 										NetworkSecurityGroup: &armcompute.SubResource{
		// 											ID: to.Ptr("/subscriptions/{subscription-id}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/networkSecurityGroups/myNetworkSecurityGroup"),
		// 										},
		// 										EnableAcceleratedNetworking: to.Ptr(true),
		// 										EnableFpga: to.Ptr(true),
		// 										DNSSettings: &armcompute.VirtualMachineScaleSetNetworkConfigurationDNSSettings{
		// 											DNSServers: []*string{
		// 												to.Ptr("aaaaaaaaaaaa"),
		// 											},
		// 										},
		// 										EnableIPForwarding: to.Ptr(true),
		// 										DeleteOption: to.Ptr(armcompute.DeleteOptionsDelete),
		// 									},
		// 								},
		// 							},
		// 							HealthProbe: &armcompute.APIEntityReference{
		// 								ID: to.Ptr("/subscriptions/{subscription-id}/resourceGroups/myResourceGroup/providers/Microsoft.Compute/restorePointCollections/rpcName/restorePoints/restorePointName/diskRestorePoints/testingexcludedisk_OsDisk_1_74cdaedcea50483d9833c96adefa100f_22b4bdfe-6c54-4f72-84d8-85d8860f0c57"),
		// 							},
		// 							NetworkAPIVersion: to.Ptr(armcompute.NetworkAPIVersionTwoThousandTwenty1101),
		// 						},
		// 						SecurityProfile: &armcompute.SecurityProfile{
		// 							UefiSettings: &armcompute.UefiSettings{
		// 								SecureBootEnabled: to.Ptr(true),
		// 								VTpmEnabled: to.Ptr(true),
		// 							},
		// 							EncryptionAtHost: to.Ptr(true),
		// 							SecurityType: to.Ptr(armcompute.SecurityTypesTrustedLaunch),
		// 						},
		// 						DiagnosticsProfile: &armcompute.DiagnosticsProfile{
		// 							BootDiagnostics: &armcompute.BootDiagnostics{
		// 								Enabled: to.Ptr(true),
		// 								StorageURI: to.Ptr("aaaaaaaaaaaaaaaaaaa"),
		// 							},
		// 						},
		// 						ExtensionProfile: &armcompute.VirtualMachineScaleSetExtensionProfile{
		// 							Extensions: []*armcompute.VirtualMachineScaleSetExtension{
		// 								{
		// 									Name: to.Ptr("aa"),
		// 									Type: to.Ptr("aaaaa"),
		// 									Properties: &armcompute.VirtualMachineScaleSetExtensionProperties{
		// 										ForceUpdateTag: to.Ptr("aaaaaaaaaaaaaaaaaaaaaaaaaaa"),
		// 										Publisher: to.Ptr("aaaaaaaaaaaaa"),
		// 										Type: to.Ptr("aaaaaaaa"),
		// 										TypeHandlerVersion: to.Ptr("aaaaaaaaaaaaaaaaaaaaaaaaaaa"),
		// 										AutoUpgradeMinorVersion: to.Ptr(true),
		// 										EnableAutomaticUpgrade: to.Ptr(true),
		// 										Settings: map[string]any{
		// 										},
		// 										ProtectedSettings: map[string]any{
		// 										},
		// 										ProvisioningState: to.Ptr("aaaaaaaaaaaaaa"),
		// 										ProvisionAfterExtensions: []*string{
		// 											to.Ptr("aaaaaaaaaaaaaa"),
		// 										},
		// 										SuppressFailures: to.Ptr(true),
		// 									},
		// 									ID: to.Ptr("aaaaaaaaaaaaaaaaaaaaaa"),
		// 								},
		// 							},
		// 							ExtensionsTimeBudget: to.Ptr("aaaaaaaaaaaaaaaaaaaaaaaaaa"),
		// 						},
		// 						LicenseType: to.Ptr("aaaa"),
		// 						Priority: to.Ptr(armcompute.VirtualMachinePriorityTypesRegular),
		// 						EvictionPolicy: to.Ptr(armcompute.VirtualMachineEvictionPolicyTypesDeallocate),
		// 						BillingProfile: &armcompute.BillingProfile{
		// 							MaxPrice: to.Ptr[float64](4),
		// 						},
		// 						ScheduledEventsProfile: &armcompute.ScheduledEventsProfile{
		// 							TerminateNotificationProfile: &armcompute.TerminateNotificationProfile{
		// 								NotBeforeTimeout: to.Ptr("aa"),
		// 								Enable: to.Ptr(true),
		// 							},
		// 						},
		// 						UserData: to.Ptr("aaaaaaaaaaaaaaaaaaaaaaaaa"),
		// 						CapacityReservation: &armcompute.CapacityReservationProfile{
		// 							CapacityReservationGroup: &armcompute.SubResource{
		// 								ID: to.Ptr("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/availabilitySets/{availabilitySetName}"),
		// 							},
		// 						},
		// 						ApplicationProfile: &armcompute.ApplicationProfile{
		// 							GalleryApplications: []*armcompute.VMGalleryApplication{
		// 								{
		// 									Tags: to.Ptr("aaaaaaaaaaa"),
		// 									Order: to.Ptr[int32](29),
		// 									PackageReferenceID: to.Ptr("aaaaaaaaaa"),
		// 									ConfigurationReference: to.Ptr("aaaaa"),
		// 								},
		// 							},
		// 						},
		// 					},
		// 					ProvisioningState: to.Ptr("succeeded"),
		// 					Overprovision: to.Ptr(true),
		// 					DoNotRunExtensionsOnOverprovisionedVMs: to.Ptr(true),
		// 					PlatformFaultDomainCount: to.Ptr[int32](1),
		// 					AutomaticRepairsPolicy: &armcompute.AutomaticRepairsPolicy{
		// 						Enabled: to.Ptr(true),
		// 						GracePeriod: to.Ptr("aaaaaaaaaaaaaaaaaaaaaaaaa"),
		// 					},
		// 					UniqueID: to.Ptr("aaaaaaaaaaaaaaaaaaaaaaa"),
		// 					ZoneBalance: to.Ptr(true),
		// 					ProximityPlacementGroup: &armcompute.SubResource{
		// 						ID: to.Ptr("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/availabilitySets/{availabilitySetName}"),
		// 					},
		// 					HostGroup: &armcompute.SubResource{
		// 						ID: to.Ptr("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/availabilitySets/{availabilitySetName}"),
		// 					},
		// 					AdditionalCapabilities: &armcompute.AdditionalCapabilities{
		// 						UltraSSDEnabled: to.Ptr(true),
		// 						HibernationEnabled: to.Ptr(true),
		// 					},
		// 					ScaleInPolicy: &armcompute.ScaleInPolicy{
		// 						Rules: []*armcompute.VirtualMachineScaleSetScaleInRules{
		// 							to.Ptr(armcompute.VirtualMachineScaleSetScaleInRulesDefault),
		// 						},
		// 						ForceDeletion: to.Ptr(true),
		// 					},
		// 					OrchestrationMode: to.Ptr(armcompute.OrchestrationModeUniform),
		// 					SpotRestorePolicy: &armcompute.SpotRestorePolicy{
		// 						Enabled: to.Ptr(true),
		// 						RestoreTimeout: to.Ptr("aaaaaaaaaa"),
		// 					},
		// 				},
		// 				Plan: &armcompute.Plan{
		// 					Name: to.Ptr("aaaaaaaaaa"),
		// 					Publisher: to.Ptr("aaaaaaaaaaaaaaaaaaaaaa"),
		// 					Product: to.Ptr("aaaaaaaaaaaaaaaaaaaa"),
		// 					PromotionCode: to.Ptr("aaaaaaaaaaaaaaaaaaaa"),
		// 				},
		// 				Identity: &armcompute.VirtualMachineScaleSetIdentity{
		// 					PrincipalID: to.Ptr("aaaaaaaaaaaaaaa"),
		// 					TenantID: to.Ptr("aaaaaaaaaaaaaaaa"),
		// 					Type: to.Ptr(armcompute.ResourceIdentityTypeSystemAssigned),
		// 					UserAssignedIdentities: map[string]*armcompute.UserAssignedIdentitiesValue{
		// 						"key3951": &armcompute.UserAssignedIdentitiesValue{
		// 							PrincipalID: to.Ptr("aaaa"),
		// 							ClientID: to.Ptr("aaaaaaaaaaaaaaaaaaaaaaaaaaa"),
		// 						},
		// 					},
		// 				},
		// 				Zones: []*string{
		// 					to.Ptr("aaaaaaa"),
		// 				},
		// 				ExtendedLocation: &armcompute.ExtendedLocation{
		// 					Name: to.Ptr("aaaaaaaaaaaaaaaaaaaaa"),
		// 					Type: to.Ptr(armcompute.ExtendedLocationTypesEdgeZone),
		// 				},
		// 			},
		// 			{
		// 				Name: to.Ptr("{virtualMachineScaleSetName}"),
		// 				ID: to.Ptr("/subscriptions/{subscription-id}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/virtualMachineScaleSets/{virtualMachineScaleSetName}1"),
		// 				Type: to.Ptr("Microsoft.Compute/virtualMachineScaleSets"),
		// 				Location: to.Ptr("eastus"),
		// 				Tags: map[string]*string{
		// 				},
		// 				SKU: &armcompute.SKU{
		// 					Name: to.Ptr("Standard_D2s_v3"),
		// 					Tier: to.Ptr("Standard"),
		// 					Capacity: to.Ptr[int64](4),
		// 				},
		// 				Properties: &armcompute.VirtualMachineScaleSetProperties{
		// 					SinglePlacementGroup: to.Ptr(true),
		// 					UpgradePolicy: &armcompute.UpgradePolicy{
		// 						Mode: to.Ptr(armcompute.UpgradeModeAutomatic),
		// 						AutomaticOSUpgradePolicy: &armcompute.AutomaticOSUpgradePolicy{
		// 							EnableAutomaticOSUpgrade: to.Ptr(true),
		// 							DisableAutomaticRollback: to.Ptr(true),
		// 							UseRollingUpgradePolicy: to.Ptr(true),
		// 							OSRollingUpgradeDeferral: to.Ptr(true),
		// 						},
		// 						RollingUpgradePolicy: &armcompute.RollingUpgradePolicy{
		// 							MaxBatchInstancePercent: to.Ptr[int32](49),
		// 							MaxUnhealthyInstancePercent: to.Ptr[int32](81),
		// 							MaxUnhealthyUpgradedInstancePercent: to.Ptr[int32](98),
		// 							PauseTimeBetweenBatches: to.Ptr("aaaaaaaaaaaaaaa"),
		// 							EnableCrossZoneUpgrade: to.Ptr(true),
		// 							PrioritizeUnhealthyInstances: to.Ptr(true),
		// 							MaxSurge: to.Ptr(true),
		// 						},
		// 					},
		// 					VirtualMachineProfile: &armcompute.VirtualMachineScaleSetVMProfile{
		// 						StorageProfile: &armcompute.VirtualMachineScaleSetStorageProfile{
		// 							OSDisk: &armcompute.VirtualMachineScaleSetOSDisk{
		// 								CreateOption: to.Ptr(armcompute.DiskCreateOptionTypesFromImage),
		// 								Caching: to.Ptr(armcompute.CachingTypesReadWrite),
		// 								ManagedDisk: &armcompute.VirtualMachineScaleSetManagedDiskParameters{
		// 									StorageAccountType: to.Ptr(armcompute.StorageAccountTypesPremiumLRS),
		// 									DiskEncryptionSet: &armcompute.DiskEncryptionSetParameters{
		// 										ID: to.Ptr("aaaaaaaaaaaa"),
		// 									},
		// 								},
		// 								DiskSizeGB: to.Ptr[int32](30),
		// 								Name: to.Ptr("a"),
		// 								WriteAcceleratorEnabled: to.Ptr(true),
		// 								DiffDiskSettings: &armcompute.DiffDiskSettings{
		// 									Option: to.Ptr(armcompute.DiffDiskOptionsLocal),
		// 									Placement: to.Ptr(armcompute.DiffDiskPlacementCacheDisk),
		// 								},
		// 								OSType: to.Ptr(armcompute.OperatingSystemTypesWindows),
		// 								Image: &armcompute.VirtualHardDisk{
		// 									URI: to.Ptr("https://{storageAccountName}.blob.core.windows.net/{containerName}/{vhdName}.vhd"),
		// 								},
		// 								VhdContainers: []*string{
		// 									to.Ptr("aaaaaaaaaaaaaaaa"),
		// 								},
		// 							},
		// 							ImageReference: &armcompute.ImageReference{
		// 								Publisher: to.Ptr("azuredatabricks"),
		// 								Offer: to.Ptr("databricks"),
		// 								SKU: to.Ptr("databricksworker"),
		// 								Version: to.Ptr("3.15.2"),
		// 								ExactVersion: to.Ptr("aa"),
		// 								SharedGalleryImageID: to.Ptr("aaaaaaa"),
		// 								ID: to.Ptr("aaa"),
		// 							},
		// 							DataDisks: []*armcompute.VirtualMachineScaleSetDataDisk{
		// 							},
		// 						},
		// 						OSProfile: &armcompute.VirtualMachineScaleSetOSProfile{
		// 							ComputerNamePrefix: to.Ptr("{virtualMachineScaleSetName}"),
		// 							AdminUsername: to.Ptr("admin"),
		// 							LinuxConfiguration: &armcompute.LinuxConfiguration{
		// 								DisablePasswordAuthentication: to.Ptr(true),
		// 								SSH: &armcompute.SSHConfiguration{
		// 									PublicKeys: []*armcompute.SSHPublicKey{
		// 										{
		// 											Path: to.Ptr("aaa"),
		// 											KeyData: to.Ptr("aaaaaa"),
		// 										},
		// 									},
		// 								},
		// 								ProvisionVMAgent: to.Ptr(true),
		// 								PatchSettings: &armcompute.LinuxPatchSettings{
		// 									PatchMode: to.Ptr(armcompute.LinuxVMGuestPatchModeImageDefault),
		// 									AssessmentMode: to.Ptr(armcompute.LinuxPatchAssessmentModeImageDefault),
		// 								},
		// 							},
		// 							CustomData: to.Ptr("a"),
		// 							WindowsConfiguration: &armcompute.WindowsConfiguration{
		// 								ProvisionVMAgent: to.Ptr(true),
		// 								EnableAutomaticUpdates: to.Ptr(true),
		// 								TimeZone: to.Ptr("aaaaaaaaaaaaaaaaaaaaaaaaaaaaa"),
		// 								AdditionalUnattendContent: []*armcompute.AdditionalUnattendContent{
		// 									{
		// 										PassName: to.Ptr("OobeSystem"),
		// 										ComponentName: to.Ptr("Microsoft-Windows-Shell-Setup"),
		// 										SettingName: to.Ptr(armcompute.SettingNamesAutoLogon),
		// 										Content: to.Ptr("aaaaaaaaaaaaaaaaaaaa"),
		// 									},
		// 								},
		// 								PatchSettings: &armcompute.PatchSettings{
		// 									PatchMode: to.Ptr(armcompute.WindowsVMGuestPatchModeManual),
		// 									EnableHotpatching: to.Ptr(true),
		// 									AssessmentMode: to.Ptr(armcompute.WindowsPatchAssessmentModeImageDefault),
		// 								},
		// 								WinRM: &armcompute.WinRMConfiguration{
		// 									Listeners: []*armcompute.WinRMListener{
		// 										{
		// 											Protocol: to.Ptr(armcompute.ProtocolTypesHTTP),
		// 											CertificateURL: to.Ptr("aaaaaaaaaaaaaaaaaaaaaa"),
		// 										},
		// 									},
		// 								},
		// 							},
		// 							Secrets: []*armcompute.VaultSecretGroup{
		// 								{
		// 									SourceVault: &armcompute.SubResource{
		// 										ID: to.Ptr("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/availabilitySets/{availabilitySetName}"),
		// 									},
		// 									VaultCertificates: []*armcompute.VaultCertificate{
		// 										{
		// 											CertificateURL: to.Ptr("aaaaaaa"),
		// 											CertificateStore: to.Ptr("aaaaaaaaaaaaaaaaaaaaaaaaa"),
		// 										},
		// 									},
		// 								},
		// 							},
		// 						},
		// 						NetworkProfile: &armcompute.VirtualMachineScaleSetNetworkProfile{
		// 							NetworkInterfaceConfigurations: []*armcompute.VirtualMachineScaleSetNetworkConfiguration{
		// 								{
		// 									Name: to.Ptr("myNic1"),
		// 									Properties: &armcompute.VirtualMachineScaleSetNetworkConfigurationProperties{
		// 										Primary: to.Ptr(true),
		// 										IPConfigurations: []*armcompute.VirtualMachineScaleSetIPConfiguration{
		// 											{
		// 												Name: to.Ptr("myIPConfig"),
		// 												Properties: &armcompute.VirtualMachineScaleSetIPConfigurationProperties{
		// 													Primary: to.Ptr(true),
		// 													Subnet: &armcompute.APIEntityReference{
		// 														ID: to.Ptr("/subscriptions/{subscription-id}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/virtualNetworks/myVNet/subnets/mySubnet"),
		// 													},
		// 													PublicIPAddressConfiguration: &armcompute.VirtualMachineScaleSetPublicIPAddressConfiguration{
		// 														Name: to.Ptr("aaaaaaaaaaaaaaaaaa"),
		// 														Properties: &armcompute.VirtualMachineScaleSetPublicIPAddressConfigurationProperties{
		// 															IdleTimeoutInMinutes: to.Ptr[int32](18),
		// 															DNSSettings: &armcompute.VirtualMachineScaleSetPublicIPAddressConfigurationDNSSettings{
		// 																DomainNameLabel: to.Ptr("aaaaaaaaaaaaaaaaaa"),
		// 															},
		// 															IPTags: []*armcompute.VirtualMachineScaleSetIPTag{
		// 																{
		// 																	IPTagType: to.Ptr("aaaaaaa"),
		// 																	Tag: to.Ptr("aaaaaaaaaaaaaaaaaaaaaaaaaaa"),
		// 																},
		// 															},
		// 															PublicIPPrefix: &armcompute.SubResource{
		// 																ID: to.Ptr("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/availabilitySets/{availabilitySetName}"),
		// 															},
		// 															PublicIPAddressVersion: to.Ptr(armcompute.IPVersionIPv4),
		// 															DeleteOption: to.Ptr(armcompute.DeleteOptionsDelete),
		// 														},
		// 														SKU: &armcompute.PublicIPAddressSKU{
		// 															Name: to.Ptr(armcompute.PublicIPAddressSKUNameBasic),
		// 															Tier: to.Ptr(armcompute.PublicIPAddressSKUTierRegional),
		// 														},
		// 													},
		// 													PrivateIPAddressVersion: to.Ptr(armcompute.IPVersionIPv4),
		// 													ApplicationGatewayBackendAddressPools: []*armcompute.SubResource{
		// 														{
		// 															ID: to.Ptr("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/availabilitySets/{availabilitySetName}"),
		// 														},
		// 													},
		// 													ApplicationSecurityGroups: []*armcompute.SubResource{
		// 														{
		// 															ID: to.Ptr("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/availabilitySets/{availabilitySetName}"),
		// 														},
		// 													},
		// 													LoadBalancerBackendAddressPools: []*armcompute.SubResource{
		// 														{
		// 															ID: to.Ptr("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/availabilitySets/{availabilitySetName}"),
		// 														},
		// 													},
		// 													LoadBalancerInboundNatPools: []*armcompute.SubResource{
		// 														{
		// 															ID: to.Ptr("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/availabilitySets/{availabilitySetName}"),
		// 														},
		// 													},
		// 												},
		// 											},
		// 										},
		// 										NetworkSecurityGroup: &armcompute.SubResource{
		// 											ID: to.Ptr("/subscriptions/{subscription-id}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/networkSecurityGroups/myNetworkSecurityGroup"),
		// 										},
		// 										EnableAcceleratedNetworking: to.Ptr(true),
		// 										EnableFpga: to.Ptr(true),
		// 										DNSSettings: &armcompute.VirtualMachineScaleSetNetworkConfigurationDNSSettings{
		// 											DNSServers: []*string{
		// 												to.Ptr("aaaaaaaaaaaa"),
		// 											},
		// 										},
		// 										EnableIPForwarding: to.Ptr(true),
		// 										DeleteOption: to.Ptr(armcompute.DeleteOptionsDelete),
		// 									},
		// 								},
		// 							},
		// 							HealthProbe: &armcompute.APIEntityReference{
		// 								ID: to.Ptr("/subscriptions/{subscription-id}/resourceGroups/myResourceGroup/providers/Microsoft.Compute/restorePointCollections/rpcName/restorePoints/restorePointName/diskRestorePoints/testingexcludedisk_OsDisk_1_74cdaedcea50483d9833c96adefa100f_22b4bdfe-6c54-4f72-84d8-85d8860f0c57"),
		// 							},
		// 							NetworkAPIVersion: to.Ptr(armcompute.NetworkAPIVersionTwoThousandTwenty1101),
		// 						},
		// 						SecurityProfile: &armcompute.SecurityProfile{
		// 							UefiSettings: &armcompute.UefiSettings{
		// 								SecureBootEnabled: to.Ptr(true),
		// 								VTpmEnabled: to.Ptr(true),
		// 							},
		// 							EncryptionAtHost: to.Ptr(true),
		// 							SecurityType: to.Ptr(armcompute.SecurityTypesTrustedLaunch),
		// 						},
		// 						DiagnosticsProfile: &armcompute.DiagnosticsProfile{
		// 							BootDiagnostics: &armcompute.BootDiagnostics{
		// 								Enabled: to.Ptr(true),
		// 								StorageURI: to.Ptr("aaaaaaaaaaaaaaaaaaa"),
		// 							},
		// 						},
		// 						ExtensionProfile: &armcompute.VirtualMachineScaleSetExtensionProfile{
		// 							Extensions: []*armcompute.VirtualMachineScaleSetExtension{
		// 								{
		// 									Name: to.Ptr("aa"),
		// 									Type: to.Ptr("aaaaa"),
		// 									Properties: &armcompute.VirtualMachineScaleSetExtensionProperties{
		// 										ForceUpdateTag: to.Ptr("aaaaaaaaaaaaaaaaaaaaaaaaaaa"),
		// 										Publisher: to.Ptr("aaaaaaaaaaaaa"),
		// 										Type: to.Ptr("aaaaaaaa"),
		// 										TypeHandlerVersion: to.Ptr("aaaaaaaaaaaaaaaaaaaaaaaaaaa"),
		// 										AutoUpgradeMinorVersion: to.Ptr(true),
		// 										EnableAutomaticUpgrade: to.Ptr(true),
		// 										Settings: map[string]any{
		// 										},
		// 										ProtectedSettings: map[string]any{
		// 										},
		// 										ProvisioningState: to.Ptr("aaaaaaaaaaaaaa"),
		// 										ProvisionAfterExtensions: []*string{
		// 											to.Ptr("aaaaaaaaaaaaaa"),
		// 										},
		// 										SuppressFailures: to.Ptr(true),
		// 									},
		// 									ID: to.Ptr("aaaaaaaaaaaaaaaaaaaaaa"),
		// 								},
		// 							},
		// 							ExtensionsTimeBudget: to.Ptr("aaaaaaaaaaaaaaaaaaaaaaaaaa"),
		// 						},
		// 						LicenseType: to.Ptr("aaaaaaaaaaaaaaaaaaaaaaa"),
		// 						Priority: to.Ptr(armcompute.VirtualMachinePriorityTypesRegular),
		// 						EvictionPolicy: to.Ptr(armcompute.VirtualMachineEvictionPolicyTypesDeallocate),
		// 						BillingProfile: &armcompute.BillingProfile{
		// 							MaxPrice: to.Ptr[float64](4),
		// 						},
		// 						ScheduledEventsProfile: &armcompute.ScheduledEventsProfile{
		// 							TerminateNotificationProfile: &armcompute.TerminateNotificationProfile{
		// 								NotBeforeTimeout: to.Ptr("aa"),
		// 								Enable: to.Ptr(true),
		// 							},
		// 						},
		// 						UserData: to.Ptr("aaaaaaaaaaaaaaaaaaaaaa"),
		// 						CapacityReservation: &armcompute.CapacityReservationProfile{
		// 							CapacityReservationGroup: &armcompute.SubResource{
		// 								ID: to.Ptr("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/availabilitySets/{availabilitySetName}"),
		// 							},
		// 						},
		// 						ApplicationProfile: &armcompute.ApplicationProfile{
		// 							GalleryApplications: []*armcompute.VMGalleryApplication{
		// 								{
		// 									Tags: to.Ptr("aaaaaaaaaaa"),
		// 									Order: to.Ptr[int32](29),
		// 									PackageReferenceID: to.Ptr("aaaaaaaaaa"),
		// 									ConfigurationReference: to.Ptr("aaaaa"),
		// 								},
		// 							},
		// 						},
		// 					},
		// 					ProvisioningState: to.Ptr("succeeded"),
		// 					Overprovision: to.Ptr(true),
		// 					DoNotRunExtensionsOnOverprovisionedVMs: to.Ptr(true),
		// 					PlatformFaultDomainCount: to.Ptr[int32](1),
		// 					AutomaticRepairsPolicy: &armcompute.AutomaticRepairsPolicy{
		// 						Enabled: to.Ptr(true),
		// 						GracePeriod: to.Ptr("aaaaaaaaaaaaaaaaaaaaaaaaa"),
		// 					},
		// 					UniqueID: to.Ptr("aaaaa"),
		// 					ZoneBalance: to.Ptr(true),
		// 					ProximityPlacementGroup: &armcompute.SubResource{
		// 						ID: to.Ptr("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/availabilitySets/{availabilitySetName}"),
		// 					},
		// 					HostGroup: &armcompute.SubResource{
		// 						ID: to.Ptr("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/availabilitySets/{availabilitySetName}"),
		// 					},
		// 					AdditionalCapabilities: &armcompute.AdditionalCapabilities{
		// 						UltraSSDEnabled: to.Ptr(true),
		// 						HibernationEnabled: to.Ptr(true),
		// 					},
		// 					ScaleInPolicy: &armcompute.ScaleInPolicy{
		// 						Rules: []*armcompute.VirtualMachineScaleSetScaleInRules{
		// 							to.Ptr(armcompute.VirtualMachineScaleSetScaleInRulesDefault),
		// 						},
		// 						ForceDeletion: to.Ptr(true),
		// 					},
		// 					OrchestrationMode: to.Ptr(armcompute.OrchestrationModeUniform),
		// 					SpotRestorePolicy: &armcompute.SpotRestorePolicy{
		// 						Enabled: to.Ptr(true),
		// 						RestoreTimeout: to.Ptr("aaaaaaaaaa"),
		// 					},
		// 				},
		// 				Plan: &armcompute.Plan{
		// 					Name: to.Ptr("aaaaaaaaaa"),
		// 					Publisher: to.Ptr("aaaaaaaaaaaaaaaaaaaaaa"),
		// 					Product: to.Ptr("aaaaaaaaaaaaaaaaaaaa"),
		// 					PromotionCode: to.Ptr("aaaaaaaaaaaaaaaaaaaa"),
		// 				},
		// 				Identity: &armcompute.VirtualMachineScaleSetIdentity{
		// 					PrincipalID: to.Ptr("aaaaaaaaaaaaaaa"),
		// 					TenantID: to.Ptr("aaaaaaaaaaaaaaaa"),
		// 					Type: to.Ptr(armcompute.ResourceIdentityTypeSystemAssigned),
		// 					UserAssignedIdentities: map[string]*armcompute.UserAssignedIdentitiesValue{
		// 						"key3951": &armcompute.UserAssignedIdentitiesValue{
		// 							PrincipalID: to.Ptr("aaaa"),
		// 							ClientID: to.Ptr("aaaaaaaaaaaaaaaaaaaaaaaaaaa"),
		// 						},
		// 					},
		// 				},
		// 				Zones: []*string{
		// 					to.Ptr("aaaaaaaa"),
		// 				},
		// 				ExtendedLocation: &armcompute.ExtendedLocation{
		// 					Name: to.Ptr("aaaaaaaaaaaaaaaaaaaaa"),
		// 					Type: to.Ptr(armcompute.ExtendedLocationTypesEdgeZone),
		// 				},
		// 			},
		// 		},
		// 		NextLink: to.Ptr("a://example.com/aaaaaaaaaaaaaaaaaaa"),
		// 	},
		// }
	}
}
