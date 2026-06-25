package armcompute_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/compute/armcompute/v8"
)

// Generated from example definition: 2026-03-01/virtualMachineScaleSetExamples/VirtualMachineScaleSet_ListAll_MaximumSet_Gen.json
func ExampleVirtualMachineScaleSetsClient_NewListAllPager_virtualMachineScaleSetListAllMaximumSetGen() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcompute.NewClientFactory("{subscription-id}", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewVirtualMachineScaleSetsClient().NewListAllPager(nil)
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
		// page = armcompute.VirtualMachineScaleSetsClientListAllResponse{
		// 	VirtualMachineScaleSetListWithLinkResult: armcompute.VirtualMachineScaleSetListWithLinkResult{
		// 		Value: []*armcompute.VirtualMachineScaleSet{
		// 			{
		// 				SKU: &armcompute.SKU{
		// 					Tier: to.Ptr("Standard"),
		// 					Capacity: to.Ptr[int64](3),
		// 					Name: to.Ptr("Standard_D1_v2"),
		// 				},
		// 				Location: to.Ptr("westus"),
		// 				Properties: &armcompute.VirtualMachineScaleSetProperties{
		// 					Overprovision: to.Ptr(true),
		// 					VirtualMachineProfile: &armcompute.VirtualMachineScaleSetVMProfile{
		// 						StorageProfile: &armcompute.VirtualMachineScaleSetStorageProfile{
		// 							ImageReference: &armcompute.ImageReference{
		// 								SKU: to.Ptr("2016-Datacenter"),
		// 								Publisher: to.Ptr("MicrosoftWindowsServer"),
		// 								Version: to.Ptr("latest"),
		// 								Offer: to.Ptr("WindowsServer"),
		// 								ExactVersion: to.Ptr("aaaaaaaaaaaaaaaaaaaaaaaaaaaaaa"),
		// 								SharedGalleryImageID: to.Ptr("aaaaaaaaaaaaaaaaaaaaaa"),
		// 								ID: to.Ptr("a"),
		// 							},
		// 							OSDisk: &armcompute.VirtualMachineScaleSetOSDisk{
		// 								Caching: to.Ptr(armcompute.CachingTypesReadWrite),
		// 								ManagedDisk: &armcompute.VirtualMachineScaleSetManagedDiskParameters{
		// 									StorageAccountType: to.Ptr(armcompute.StorageAccountTypesStandardLRS),
		// 									DiskEncryptionSet: &armcompute.DiskEncryptionSetParameters{
		// 										ID: to.Ptr("aaaaaaaaaaaa"),
		// 									},
		// 								},
		// 								CreateOption: to.Ptr(armcompute.DiskCreateOptionTypesFromImage),
		// 								Name: to.Ptr("aaaaaaaaaaaaaaa"),
		// 								WriteAcceleratorEnabled: to.Ptr(true),
		// 								DiffDiskSettings: &armcompute.DiffDiskSettings{
		// 									Option: to.Ptr(armcompute.DiffDiskOptionsLocal),
		// 									Placement: to.Ptr(armcompute.DiffDiskPlacementCacheDisk),
		// 								},
		// 								DiskSizeGB: to.Ptr[int32](30),
		// 								OSType: to.Ptr(armcompute.OperatingSystemTypesWindows),
		// 								Image: &armcompute.VirtualHardDisk{
		// 									URI: to.Ptr("https://{storageAccountName}.blob.core.windows.net/{containerName}/{vhdName}.vhd"),
		// 								},
		// 								VhdContainers: []*string{
		// 									to.Ptr("aaaaaaaaaaaaaaaaaaaaaaaaa"),
		// 								},
		// 							},
		// 							DataDisks: []*armcompute.VirtualMachineScaleSetDataDisk{
		// 								{
		// 									Name: to.Ptr("aaaaaaaaaaaaaa"),
		// 									Lun: to.Ptr[int32](24),
		// 									Caching: to.Ptr(armcompute.CachingTypesNone),
		// 									WriteAcceleratorEnabled: to.Ptr(true),
		// 									CreateOption: to.Ptr(armcompute.DiskCreateOptionTypesFromImage),
		// 									DiskSizeGB: to.Ptr[int32](11),
		// 									ManagedDisk: &armcompute.VirtualMachineScaleSetManagedDiskParameters{
		// 										StorageAccountType: to.Ptr(armcompute.StorageAccountTypesStandardLRS),
		// 										DiskEncryptionSet: &armcompute.DiskEncryptionSetParameters{
		// 											ID: to.Ptr("aaaaaaaaaaaa"),
		// 										},
		// 									},
		// 									DiskIOPSReadWrite: to.Ptr[int64](11),
		// 									DiskMBpsReadWrite: to.Ptr[int64](13),
		// 								},
		// 							},
		// 						},
		// 						OSProfile: &armcompute.VirtualMachineScaleSetOSProfile{
		// 							ComputerNamePrefix: to.Ptr("{vmss-name}"),
		// 							AdminUsername: to.Ptr("{your-username}"),
		// 							CustomData: to.Ptr("aaaaaaaaaaaaaaaaaaaaaaaaaa"),
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
		// 									Name: to.Ptr("{vmss-name}"),
		// 									Properties: &armcompute.VirtualMachineScaleSetNetworkConfigurationProperties{
		// 										Primary: to.Ptr(true),
		// 										EnableIPForwarding: to.Ptr(true),
		// 										IPConfigurations: []*armcompute.VirtualMachineScaleSetIPConfiguration{
		// 											{
		// 												Name: to.Ptr("{vmss-name}"),
		// 												Properties: &armcompute.VirtualMachineScaleSetIPConfigurationProperties{
		// 													Subnet: &armcompute.APIEntityReference{
		// 														ID: to.Ptr("/subscriptions/{subscription-id}/resourceGroups/myResourceGroup/providers/Microsoft.Network/virtualNetworks/{existing-virtual-network-name}/subnets/{existing-subnet-name}"),
		// 													},
		// 													Primary: to.Ptr(true),
		// 													PublicIPAddressConfiguration: &armcompute.VirtualMachineScaleSetPublicIPAddressConfiguration{
		// 														Name: to.Ptr("aaaaaaaaaaaaaaaaaa"),
		// 														Properties: &armcompute.VirtualMachineScaleSetPublicIPAddressConfigurationProperties{
		// 															IdleTimeoutInMinutes: to.Ptr[int32](18),
		// 															DNSSettings: &armcompute.VirtualMachineScaleSetPublicIPAddressConfigurationDNSSettings{
		// 																DomainNameLabel: to.Ptr("aaaaaaaaaaaaaaaaaa"),
		// 																DomainNameLabelScope: to.Ptr(armcompute.DomainNameLabelScopeTypesTenantReuse),
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
		// 										EnableAcceleratedNetworking: to.Ptr(true),
		// 										EnableFpga: to.Ptr(true),
		// 										NetworkSecurityGroup: &armcompute.SubResource{
		// 											ID: to.Ptr("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/availabilitySets/{availabilitySetName}"),
		// 										},
		// 										DNSSettings: &armcompute.VirtualMachineScaleSetNetworkConfigurationDNSSettings{
		// 											DNSServers: []*string{
		// 												to.Ptr("aaaaaaaaaaaa"),
		// 											},
		// 										},
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
		// 						LicenseType: to.Ptr("aaaaaaaaaa"),
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
		// 						UserData: to.Ptr("aaa"),
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
		// 					UpgradePolicy: &armcompute.UpgradePolicy{
		// 						Mode: to.Ptr(armcompute.UpgradeModeManual),
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
		// 						AutomaticOSUpgradePolicy: &armcompute.AutomaticOSUpgradePolicy{
		// 							EnableAutomaticOSUpgrade: to.Ptr(true),
		// 							DisableAutomaticRollback: to.Ptr(true),
		// 							OSRollingUpgradeDeferral: to.Ptr(true),
		// 						},
		// 					},
		// 					AutomaticRepairsPolicy: &armcompute.AutomaticRepairsPolicy{
		// 						Enabled: to.Ptr(true),
		// 						GracePeriod: to.Ptr("aaaaaaaaaaaaaaaaaaaaaaaaa"),
		// 					},
		// 					ProvisioningState: to.Ptr("succeeded"),
		// 					DoNotRunExtensionsOnOverprovisionedVMs: to.Ptr(true),
		// 					UniqueID: to.Ptr("aaaaaaaa"),
		// 					SinglePlacementGroup: to.Ptr(true),
		// 					ZoneBalance: to.Ptr(true),
		// 					PlatformFaultDomainCount: to.Ptr[int32](1),
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
		// 					to.Ptr("aaaaaaaaaaaaaaaaaaaa"),
		// 				},
		// 				ExtendedLocation: &armcompute.ExtendedLocation{
		// 					Name: to.Ptr("aaaaaaaaaaaaaaaaaaaaa"),
		// 					Type: to.Ptr(armcompute.ExtendedLocationTypesEdgeZone),
		// 				},
		// 				ID: to.Ptr("/subscriptions/{subscription-id}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/virtualMachineScaleSets/{virtualMachineScaleSetName}"),
		// 				Name: to.Ptr("{virtualMachineScaleSetName}"),
		// 				Type: to.Ptr("Microsoft.Compute/virtualMachineScaleSets"),
		// 				Tags: map[string]*string{
		// 					"key8425": to.Ptr("aaa"),
		// 				},
		// 			},
		// 		},
		// 		NextLink: to.Ptr("a://example.com/aaaa"),
		// 	},
		// }
	}
}
