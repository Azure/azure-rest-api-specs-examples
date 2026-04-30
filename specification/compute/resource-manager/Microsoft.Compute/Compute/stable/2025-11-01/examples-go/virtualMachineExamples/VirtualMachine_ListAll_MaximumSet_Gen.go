package armcompute_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/compute/armcompute/v8"
)

// Generated from example definition: 2025-11-01/virtualMachineExamples/VirtualMachine_ListAll_MaximumSet_Gen.json
func ExampleVirtualMachinesClient_NewListAllPager_virtualMachineListAllMaximumSetGen() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcompute.NewClientFactory("{subscription-id}", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewVirtualMachinesClient().NewListAllPager(&armcompute.VirtualMachinesClientListAllOptions{
		StatusOnly: to.Ptr("aaaaaa"),
		Filter:     to.Ptr("aaaaaaaaaaaaaaaaaaaaaaaaaaaa")})
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
		// page = armcompute.VirtualMachinesClientListAllResponse{
		// 	VirtualMachineListResult: armcompute.VirtualMachineListResult{
		// 		Value: []*armcompute.VirtualMachine{
		// 			{
		// 				Properties: &armcompute.VirtualMachineProperties{
		// 					VMID: to.Ptr("{vmId}"),
		// 					AvailabilitySet: &armcompute.SubResource{
		// 						ID: to.Ptr("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/availabilitySets/{availabilitySetName}"),
		// 					},
		// 					HardwareProfile: &armcompute.HardwareProfile{
		// 						VMSize: to.Ptr(armcompute.VirtualMachineSizeTypesStandardA0),
		// 						VMSizeProperties: &armcompute.VMSizeProperties{
		// 							VCPUsAvailable: to.Ptr[int32](7),
		// 							VCPUsPerCore: to.Ptr[int32](14),
		// 						},
		// 					},
		// 					StorageProfile: &armcompute.StorageProfile{
		// 						ImageReference: &armcompute.ImageReference{
		// 							Publisher: to.Ptr("MicrosoftWindowsServer"),
		// 							Offer: to.Ptr("WindowsServer"),
		// 							SKU: to.Ptr("2012-R2-Datacenter"),
		// 							Version: to.Ptr("4.127.20170406"),
		// 							ExactVersion: to.Ptr("aaaaaaaaaaaaa"),
		// 							SharedGalleryImageID: to.Ptr("aaaaaaaaaaaaaaa"),
		// 							CommunityGalleryImageID: to.Ptr("aaaa"),
		// 							ID: to.Ptr("aaaaaaaaaaaaaaaaaaaaaaaaaaaaaa"),
		// 						},
		// 						OSDisk: &armcompute.OSDisk{
		// 							OSType: to.Ptr(armcompute.OperatingSystemTypesWindows),
		// 							Name: to.Ptr("test"),
		// 							CreateOption: to.Ptr(armcompute.DiskCreateOptionTypesFromImage),
		// 							Vhd: &armcompute.VirtualHardDisk{
		// 								URI: to.Ptr("https://{storageAccountName}.blob.core.windows.net/{containerName}/{vhdName}.vhd"),
		// 							},
		// 							Caching: to.Ptr(armcompute.CachingTypesNone),
		// 							DiskSizeGB: to.Ptr[int32](127),
		// 							EncryptionSettings: &armcompute.DiskEncryptionSettings{
		// 								DiskEncryptionKey: &armcompute.KeyVaultSecretReference{
		// 									SecretURL: to.Ptr("aaaaaaaaa"),
		// 									SourceVault: &armcompute.SubResource{
		// 										ID: to.Ptr("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/availabilitySets/{availabilitySetName}"),
		// 									},
		// 								},
		// 								KeyEncryptionKey: &armcompute.KeyVaultKeyReference{
		// 									KeyURL: to.Ptr("aaaaaaaaaaaaa"),
		// 									SourceVault: &armcompute.SubResource{
		// 										ID: to.Ptr("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/availabilitySets/{availabilitySetName}"),
		// 									},
		// 								},
		// 								Enabled: to.Ptr(true),
		// 							},
		// 							Image: &armcompute.VirtualHardDisk{
		// 								URI: to.Ptr("https://{storageAccountName}.blob.core.windows.net/{containerName}/{vhdName}.vhd"),
		// 							},
		// 							WriteAcceleratorEnabled: to.Ptr(true),
		// 							DiffDiskSettings: &armcompute.DiffDiskSettings{
		// 								Option: to.Ptr(armcompute.DiffDiskOptionsLocal),
		// 								Placement: to.Ptr(armcompute.DiffDiskPlacementCacheDisk),
		// 							},
		// 							ManagedDisk: &armcompute.ManagedDiskParameters{
		// 								StorageAccountType: to.Ptr(armcompute.StorageAccountTypesStandardLRS),
		// 								DiskEncryptionSet: &armcompute.DiskEncryptionSetParameters{
		// 									ID: to.Ptr("aaaaaaaaaaaaaaaaaaaaaaaaaaaa"),
		// 								},
		// 								SecurityProfile: &armcompute.VMDiskSecurityProfile{
		// 									SecurityEncryptionType: to.Ptr(armcompute.SecurityEncryptionTypesVMGuestStateOnly),
		// 									DiskEncryptionSet: &armcompute.DiskEncryptionSetParameters{
		// 										ID: to.Ptr("aaaaaaaaaaaaaaaaaaaaaaaaaaaa"),
		// 									},
		// 								},
		// 								ID: to.Ptr("/subscriptions/{subscription-id}/resourceGroups/myResourceGroup/providers/Microsoft.Compute/disks/testingexcludedisk_OsDisk_1_74cdaedcea50483d9833c96adefa100f"),
		// 							},
		// 							DeleteOption: to.Ptr(armcompute.DiskDeleteOptionTypesDelete),
		// 						},
		// 						DataDisks: []*armcompute.DataDisk{
		// 						},
		// 					},
		// 					OSProfile: &armcompute.OSProfile{
		// 						ComputerName: to.Ptr("Test"),
		// 						AdminUsername: to.Ptr("Foo12"),
		// 						WindowsConfiguration: &armcompute.WindowsConfiguration{
		// 							ProvisionVMAgent: to.Ptr(true),
		// 							EnableAutomaticUpdates: to.Ptr(true),
		// 							TimeZone: to.Ptr("aaaaaaaaaaaaaaaaaaaaaa"),
		// 							AdditionalUnattendContent: []*armcompute.AdditionalUnattendContent{
		// 								{
		// 									PassName: to.Ptr("OobeSystem"),
		// 									ComponentName: to.Ptr("Microsoft-Windows-Shell-Setup"),
		// 									SettingName: to.Ptr(armcompute.SettingNamesAutoLogon),
		// 									Content: to.Ptr("aaaaaaaaaaaaaaaaaaaaaaaaaa"),
		// 								},
		// 							},
		// 							PatchSettings: &armcompute.PatchSettings{
		// 								PatchMode: to.Ptr(armcompute.WindowsVMGuestPatchModeManual),
		// 								EnableHotpatching: to.Ptr(true),
		// 								AssessmentMode: to.Ptr(armcompute.WindowsPatchAssessmentModeImageDefault),
		// 							},
		// 							WinRM: &armcompute.WinRMConfiguration{
		// 								Listeners: []*armcompute.WinRMListener{
		// 									{
		// 										Protocol: to.Ptr(armcompute.ProtocolTypesHTTP),
		// 										CertificateURL: to.Ptr("aaaaaaaaaaaaaaaaaaaaa"),
		// 									},
		// 								},
		// 							},
		// 						},
		// 						Secrets: []*armcompute.VaultSecretGroup{
		// 						},
		// 						AllowExtensionOperations: to.Ptr(true),
		// 						CustomData: to.Ptr("aaaa"),
		// 						LinuxConfiguration: &armcompute.LinuxConfiguration{
		// 							DisablePasswordAuthentication: to.Ptr(true),
		// 							SSH: &armcompute.SSHConfiguration{
		// 								PublicKeys: []*armcompute.SSHPublicKey{
		// 									{
		// 										Path: to.Ptr("aaaaaaaaaaaaaaaaaaaaaa"),
		// 										KeyData: to.Ptr("aaa"),
		// 									},
		// 								},
		// 							},
		// 							ProvisionVMAgent: to.Ptr(true),
		// 							PatchSettings: &armcompute.LinuxPatchSettings{
		// 								PatchMode: to.Ptr(armcompute.LinuxVMGuestPatchModeImageDefault),
		// 								AssessmentMode: to.Ptr(armcompute.LinuxPatchAssessmentModeImageDefault),
		// 							},
		// 						},
		// 						RequireGuestProvisionSignal: to.Ptr(true),
		// 					},
		// 					NetworkProfile: &armcompute.NetworkProfile{
		// 						NetworkInterfaces: []*armcompute.NetworkInterfaceReference{
		// 							{
		// 								ID: to.Ptr("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/networkInterfaces/{networkInterfaceName}"),
		// 								Properties: &armcompute.NetworkInterfaceReferenceProperties{
		// 									Primary: to.Ptr(true),
		// 									DeleteOption: to.Ptr(armcompute.DeleteOptionsDelete),
		// 								},
		// 							},
		// 						},
		// 						NetworkAPIVersion: to.Ptr(armcompute.NetworkAPIVersion("2022-05-01")),
		// 						NetworkInterfaceConfigurations: []*armcompute.VirtualMachineNetworkInterfaceConfiguration{
		// 							{
		// 								Name: to.Ptr("aaaaaaaa"),
		// 								Properties: &armcompute.VirtualMachineNetworkInterfaceConfigurationProperties{
		// 									Primary: to.Ptr(true),
		// 									DeleteOption: to.Ptr(armcompute.DeleteOptionsDelete),
		// 									EnableAcceleratedNetworking: to.Ptr(true),
		// 									DisableTCPStateTracking: to.Ptr(true),
		// 									EnableFpga: to.Ptr(true),
		// 									EnableIPForwarding: to.Ptr(true),
		// 									NetworkSecurityGroup: &armcompute.SubResource{
		// 										ID: to.Ptr("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/availabilitySets/{availabilitySetName}"),
		// 									},
		// 									DNSSettings: &armcompute.VirtualMachineNetworkInterfaceDNSSettingsConfiguration{
		// 										DNSServers: []*string{
		// 											to.Ptr("aaaaaaaaaaaaaaaaaaaaaaaa"),
		// 										},
		// 									},
		// 									IPConfigurations: []*armcompute.VirtualMachineNetworkInterfaceIPConfiguration{
		// 										{
		// 											Name: to.Ptr("aaaaaaaa"),
		// 											Properties: &armcompute.VirtualMachineNetworkInterfaceIPConfigurationProperties{
		// 												Subnet: &armcompute.SubResource{
		// 													ID: to.Ptr("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/availabilitySets/{availabilitySetName}"),
		// 												},
		// 												Primary: to.Ptr(true),
		// 												PublicIPAddressConfiguration: &armcompute.VirtualMachinePublicIPAddressConfiguration{
		// 													Name: to.Ptr("aaaaaaaaaaaaaaaaaa"),
		// 													Properties: &armcompute.VirtualMachinePublicIPAddressConfigurationProperties{
		// 														IdleTimeoutInMinutes: to.Ptr[int32](23),
		// 														DeleteOption: to.Ptr(armcompute.DeleteOptionsDelete),
		// 														DNSSettings: &armcompute.VirtualMachinePublicIPAddressDNSSettingsConfiguration{
		// 															DomainNameLabel: to.Ptr("aaaaa"),
		// 															DomainNameLabelScope: to.Ptr(armcompute.DomainNameLabelScopeTypesTenantReuse),
		// 														},
		// 														IPTags: []*armcompute.VirtualMachineIPTag{
		// 															{
		// 																IPTagType: to.Ptr("aaaaa"),
		// 																Tag: to.Ptr("aaaaaaaaaaaaaaaaaaaaaaaaaaa"),
		// 															},
		// 														},
		// 														PublicIPPrefix: &armcompute.SubResource{
		// 															ID: to.Ptr("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/availabilitySets/{availabilitySetName}"),
		// 														},
		// 														PublicIPAddressVersion: to.Ptr(armcompute.IPVersionsIPv4),
		// 														PublicIPAllocationMethod: to.Ptr(armcompute.PublicIPAllocationMethodDynamic),
		// 													},
		// 													SKU: &armcompute.PublicIPAddressSKU{
		// 														Name: to.Ptr(armcompute.PublicIPAddressSKUNameBasic),
		// 														Tier: to.Ptr(armcompute.PublicIPAddressSKUTierRegional),
		// 													},
		// 												},
		// 												PrivateIPAddressVersion: to.Ptr(armcompute.IPVersionsIPv4),
		// 												ApplicationSecurityGroups: []*armcompute.SubResource{
		// 													{
		// 														ID: to.Ptr("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/availabilitySets/{availabilitySetName}"),
		// 													},
		// 												},
		// 												ApplicationGatewayBackendAddressPools: []*armcompute.SubResource{
		// 													{
		// 														ID: to.Ptr("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/availabilitySets/{availabilitySetName}"),
		// 													},
		// 												},
		// 												LoadBalancerBackendAddressPools: []*armcompute.SubResource{
		// 													{
		// 														ID: to.Ptr("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/availabilitySets/{availabilitySetName}"),
		// 													},
		// 												},
		// 											},
		// 										},
		// 									},
		// 									DscpConfiguration: &armcompute.SubResource{
		// 										ID: to.Ptr("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/availabilitySets/{availabilitySetName}"),
		// 									},
		// 								},
		// 							},
		// 						},
		// 					},
		// 					ProvisioningState: to.Ptr("Succeeded"),
		// 					AdditionalCapabilities: &armcompute.AdditionalCapabilities{
		// 						UltraSSDEnabled: to.Ptr(true),
		// 						HibernationEnabled: to.Ptr(true),
		// 					},
		// 					SecurityProfile: &armcompute.SecurityProfile{
		// 						UefiSettings: &armcompute.UefiSettings{
		// 							SecureBootEnabled: to.Ptr(true),
		// 							VTpmEnabled: to.Ptr(true),
		// 						},
		// 						EncryptionAtHost: to.Ptr(true),
		// 						SecurityType: to.Ptr(armcompute.SecurityTypesTrustedLaunch),
		// 					},
		// 					DiagnosticsProfile: &armcompute.DiagnosticsProfile{
		// 						BootDiagnostics: &armcompute.BootDiagnostics{
		// 							Enabled: to.Ptr(true),
		// 							StorageURI: to.Ptr("aaaaaaaaaaaaaaaaaaaaa"),
		// 						},
		// 					},
		// 					VirtualMachineScaleSet: &armcompute.SubResource{
		// 						ID: to.Ptr("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/availabilitySets/{availabilitySetName}"),
		// 					},
		// 					ProximityPlacementGroup: &armcompute.SubResource{
		// 						ID: to.Ptr("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/availabilitySets/{availabilitySetName}"),
		// 					},
		// 					Priority: to.Ptr(armcompute.VirtualMachinePriorityTypesRegular),
		// 					EvictionPolicy: to.Ptr(armcompute.VirtualMachineEvictionPolicyTypesDeallocate),
		// 					BillingProfile: &armcompute.BillingProfile{
		// 						MaxPrice: to.Ptr[float64](26),
		// 					},
		// 					Host: &armcompute.SubResource{
		// 						ID: to.Ptr("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/availabilitySets/{availabilitySetName}"),
		// 					},
		// 					HostGroup: &armcompute.SubResource{
		// 						ID: to.Ptr("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/availabilitySets/{availabilitySetName}"),
		// 					},
		// 					LicenseType: to.Ptr("aaaaaaaaaaaaaaa"),
		// 					ExtensionsTimeBudget: to.Ptr("aaaaaaaaaaaaaaaaaaaaaaaaaaaa"),
		// 					PlatformFaultDomain: to.Ptr[int32](8),
		// 					ScheduledEventsProfile: &armcompute.ScheduledEventsProfile{
		// 						TerminateNotificationProfile: &armcompute.TerminateNotificationProfile{
		// 							NotBeforeTimeout: to.Ptr("PT10M"),
		// 							Enable: to.Ptr(true),
		// 						},
		// 						OSImageNotificationProfile: &armcompute.OSImageNotificationProfile{
		// 							NotBeforeTimeout: to.Ptr("PT15M"),
		// 							Enable: to.Ptr(true),
		// 						},
		// 					},
		// 					UserData: to.Ptr("aaa"),
		// 					CapacityReservation: &armcompute.CapacityReservationProfile{
		// 						CapacityReservationGroup: &armcompute.SubResource{
		// 							ID: to.Ptr("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/availabilitySets/{availabilitySetName}"),
		// 						},
		// 					},
		// 					ApplicationProfile: &armcompute.ApplicationProfile{
		// 						GalleryApplications: []*armcompute.VMGalleryApplication{
		// 							{
		// 								Tags: to.Ptr("aaaaa"),
		// 								Order: to.Ptr[int32](4),
		// 								PackageReferenceID: to.Ptr("aaaaaaaaaaaaaaaaaaaaaa"),
		// 								ConfigurationReference: to.Ptr("aaaaaaaaaaaaaaaaaaaaaaaaaaaaaa"),
		// 							},
		// 						},
		// 					},
		// 					TimeCreated: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2022-01-14T16:43:41.683Z"); return t}()),
		// 				},
		// 				Type: to.Ptr("Microsoft.Compute/virtualMachines"),
		// 				Location: to.Ptr("eastus"),
		// 				Tags: map[string]*string{
		// 				},
		// 				ID: to.Ptr("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/virtualMachines/{virtualMachineName}"),
		// 				Name: to.Ptr("{virtualMachineName}"),
		// 				Plan: &armcompute.Plan{
		// 					Name: to.Ptr("aaaaaaaaaaaaaaaaaa"),
		// 					Publisher: to.Ptr("aaaaaaaaaaaaaaaaa"),
		// 					Product: to.Ptr("aaaaaaaaaaaaaaaaaaaaaa"),
		// 					PromotionCode: to.Ptr("aaaaaaaaaaaaaa"),
		// 				},
		// 				Resources: []*armcompute.VirtualMachineExtension{
		// 					{
		// 						Properties: &armcompute.VirtualMachineExtensionProperties{
		// 							ForceUpdateTag: to.Ptr("aaaaaaaaaaaaaaa"),
		// 							Publisher: to.Ptr("aaaaaaaaaaaaaaaa"),
		// 							Type: to.Ptr("aaaaaaaa"),
		// 							TypeHandlerVersion: to.Ptr("aaaaaaaaaaaa"),
		// 							AutoUpgradeMinorVersion: to.Ptr(true),
		// 							EnableAutomaticUpgrade: to.Ptr(true),
		// 							Settings: map[string]any{
		// 							},
		// 							ProtectedSettings: map[string]any{
		// 							},
		// 							ProvisioningState: to.Ptr("aaa"),
		// 							SuppressFailures: to.Ptr(true),
		// 							ProtectedSettingsFromKeyVault: &armcompute.KeyVaultSecretReference{
		// 								SourceVault: &armcompute.SubResource{
		// 									ID: to.Ptr("/subscriptions/a53f7094-a16c-47af-abe4-b05c05d0d79a/resourceGroups/myResourceGroup/providers/Microsoft.KeyVault/vaults/kvName"),
		// 								},
		// 								SecretURL: to.Ptr("https://kvName.vault.azure.net/secrets/secretName/79b88b3a6f5440ffb2e73e44a0db712e"),
		// 							},
		// 						},
		// 						ID: to.Ptr("aaaaaaaaaaaaaaaaaaaaaaaaaaaaa"),
		// 						Name: to.Ptr("aaaaaaaaaaaaa"),
		// 						Type: to.Ptr("aaaaaaaaaaaaaaaaaaaaaaaaaaa"),
		// 						Location: to.Ptr("aaaaaaaaaaaaaaaa"),
		// 						Tags: map[string]*string{
		// 							"key9428": to.Ptr("aaaaaaa"),
		// 						},
		// 					},
		// 				},
		// 				Identity: &armcompute.VirtualMachineIdentity{
		// 					PrincipalID: to.Ptr("aaaaaaaaaaaaaaaa"),
		// 					TenantID: to.Ptr("aaaaa"),
		// 					Type: to.Ptr(armcompute.ResourceIdentityTypeSystemAssigned),
		// 					UserAssignedIdentities: map[string]*armcompute.UserAssignedIdentitiesValue{
		// 						"key5688": &armcompute.UserAssignedIdentitiesValue{
		// 							PrincipalID: to.Ptr("aaaaaaaaaaaaaaa"),
		// 							ClientID: to.Ptr("aaaaaaaaaaa"),
		// 						},
		// 					},
		// 				},
		// 				Zones: []*string{
		// 					to.Ptr("aaaaaaaaaaaaaaaaaaaaaaaaaaa"),
		// 				},
		// 				ExtendedLocation: &armcompute.ExtendedLocation{
		// 					Name: to.Ptr("aaaa"),
		// 					Type: to.Ptr(armcompute.ExtendedLocationTypesEdgeZone),
		// 				},
		// 			},
		// 			{
		// 				Properties: &armcompute.VirtualMachineProperties{
		// 					VMID: to.Ptr("{vmId}"),
		// 					AvailabilitySet: &armcompute.SubResource{
		// 						ID: to.Ptr("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/availabilitySets/{availabilitySetName}"),
		// 					},
		// 					HardwareProfile: &armcompute.HardwareProfile{
		// 						VMSize: to.Ptr(armcompute.VirtualMachineSizeTypesStandardA0),
		// 						VMSizeProperties: &armcompute.VMSizeProperties{
		// 							VCPUsAvailable: to.Ptr[int32](7),
		// 							VCPUsPerCore: to.Ptr[int32](14),
		// 						},
		// 					},
		// 					StorageProfile: &armcompute.StorageProfile{
		// 						ImageReference: &armcompute.ImageReference{
		// 							Publisher: to.Ptr("MicrosoftWindowsServer"),
		// 							Offer: to.Ptr("WindowsServer"),
		// 							SKU: to.Ptr("2012-R2-Datacenter"),
		// 							Version: to.Ptr("4.127.20170406"),
		// 							ExactVersion: to.Ptr("aa"),
		// 							SharedGalleryImageID: to.Ptr("aaaaaaaaaaaaaaaaaaaaaaaaa"),
		// 							CommunityGalleryImageID: to.Ptr("aaaaaaaaaaaaaaaaaaaaaaaaaaaa"),
		// 							ID: to.Ptr("aaaaaaaa"),
		// 						},
		// 						OSDisk: &armcompute.OSDisk{
		// 							OSType: to.Ptr(armcompute.OperatingSystemTypesWindows),
		// 							Name: to.Ptr("test"),
		// 							CreateOption: to.Ptr(armcompute.DiskCreateOptionTypesFromImage),
		// 							Vhd: &armcompute.VirtualHardDisk{
		// 								URI: to.Ptr("https://{storageAccountName}.blob.core.windows.net/{containerName}/{vhdName}.vhd"),
		// 							},
		// 							Caching: to.Ptr(armcompute.CachingTypesNone),
		// 							DiskSizeGB: to.Ptr[int32](127),
		// 							EncryptionSettings: &armcompute.DiskEncryptionSettings{
		// 								DiskEncryptionKey: &armcompute.KeyVaultSecretReference{
		// 									SecretURL: to.Ptr("aaaaaaaaa"),
		// 									SourceVault: &armcompute.SubResource{
		// 										ID: to.Ptr("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/availabilitySets/{availabilitySetName}"),
		// 									},
		// 								},
		// 								KeyEncryptionKey: &armcompute.KeyVaultKeyReference{
		// 									KeyURL: to.Ptr("aaaaaaaaaaaaa"),
		// 									SourceVault: &armcompute.SubResource{
		// 										ID: to.Ptr("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/availabilitySets/{availabilitySetName}"),
		// 									},
		// 								},
		// 								Enabled: to.Ptr(true),
		// 							},
		// 							Image: &armcompute.VirtualHardDisk{
		// 								URI: to.Ptr("https://{storageAccountName}.blob.core.windows.net/{containerName}/{vhdName}.vhd"),
		// 							},
		// 							WriteAcceleratorEnabled: to.Ptr(true),
		// 							DiffDiskSettings: &armcompute.DiffDiskSettings{
		// 								Option: to.Ptr(armcompute.DiffDiskOptionsLocal),
		// 								Placement: to.Ptr(armcompute.DiffDiskPlacementCacheDisk),
		// 							},
		// 							ManagedDisk: &armcompute.ManagedDiskParameters{
		// 								StorageAccountType: to.Ptr(armcompute.StorageAccountTypesStandardLRS),
		// 								DiskEncryptionSet: &armcompute.DiskEncryptionSetParameters{
		// 									ID: to.Ptr("aaaaaaaaaaaaaaaaaaaaaaaaaaaa"),
		// 								},
		// 								SecurityProfile: &armcompute.VMDiskSecurityProfile{
		// 									SecurityEncryptionType: to.Ptr(armcompute.SecurityEncryptionTypesVMGuestStateOnly),
		// 									DiskEncryptionSet: &armcompute.DiskEncryptionSetParameters{
		// 										ID: to.Ptr("aaaaaaaaaaaaaaaaaaaaaaaaaaaa"),
		// 									},
		// 								},
		// 								ID: to.Ptr("/subscriptions/{subscription-id}/resourceGroups/myResourceGroup/providers/Microsoft.Compute/disks/testingexcludedisk_OsDisk_1_74cdaedcea50483d9833c96adefa100f"),
		// 							},
		// 							DeleteOption: to.Ptr(armcompute.DiskDeleteOptionTypesDelete),
		// 						},
		// 						DataDisks: []*armcompute.DataDisk{
		// 						},
		// 					},
		// 					OSProfile: &armcompute.OSProfile{
		// 						ComputerName: to.Ptr("Test"),
		// 						AdminUsername: to.Ptr("Foo12"),
		// 						WindowsConfiguration: &armcompute.WindowsConfiguration{
		// 							ProvisionVMAgent: to.Ptr(true),
		// 							EnableAutomaticUpdates: to.Ptr(true),
		// 							TimeZone: to.Ptr("aaaaaaaaaaaaaaaaaaaa"),
		// 							AdditionalUnattendContent: []*armcompute.AdditionalUnattendContent{
		// 								{
		// 									PassName: to.Ptr("OobeSystem"),
		// 									ComponentName: to.Ptr("Microsoft-Windows-Shell-Setup"),
		// 									SettingName: to.Ptr(armcompute.SettingNamesAutoLogon),
		// 									Content: to.Ptr("aaaaaaaaaaaaaaaaaaaaaaaaaa"),
		// 								},
		// 							},
		// 							PatchSettings: &armcompute.PatchSettings{
		// 								PatchMode: to.Ptr(armcompute.WindowsVMGuestPatchModeManual),
		// 								EnableHotpatching: to.Ptr(true),
		// 								AssessmentMode: to.Ptr(armcompute.WindowsPatchAssessmentModeImageDefault),
		// 							},
		// 							WinRM: &armcompute.WinRMConfiguration{
		// 								Listeners: []*armcompute.WinRMListener{
		// 									{
		// 										Protocol: to.Ptr(armcompute.ProtocolTypesHTTP),
		// 										CertificateURL: to.Ptr("aaaaaaaaaaaaaaaaaaaaa"),
		// 									},
		// 								},
		// 							},
		// 						},
		// 						Secrets: []*armcompute.VaultSecretGroup{
		// 						},
		// 						AllowExtensionOperations: to.Ptr(true),
		// 						CustomData: to.Ptr("aaaaaaaaaaaaaaaaaaaaaaa"),
		// 						LinuxConfiguration: &armcompute.LinuxConfiguration{
		// 							DisablePasswordAuthentication: to.Ptr(true),
		// 							SSH: &armcompute.SSHConfiguration{
		// 								PublicKeys: []*armcompute.SSHPublicKey{
		// 									{
		// 										Path: to.Ptr("aaaaaaaaaaaaaaaaaaaaaa"),
		// 										KeyData: to.Ptr("aaa"),
		// 									},
		// 								},
		// 							},
		// 							ProvisionVMAgent: to.Ptr(true),
		// 							PatchSettings: &armcompute.LinuxPatchSettings{
		// 								PatchMode: to.Ptr(armcompute.LinuxVMGuestPatchModeImageDefault),
		// 								AssessmentMode: to.Ptr(armcompute.LinuxPatchAssessmentModeImageDefault),
		// 							},
		// 						},
		// 						RequireGuestProvisionSignal: to.Ptr(true),
		// 					},
		// 					NetworkProfile: &armcompute.NetworkProfile{
		// 						NetworkInterfaces: []*armcompute.NetworkInterfaceReference{
		// 							{
		// 								ID: to.Ptr("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/networkInterfaces/{networkInterfaceName}"),
		// 								Properties: &armcompute.NetworkInterfaceReferenceProperties{
		// 									Primary: to.Ptr(true),
		// 									DeleteOption: to.Ptr(armcompute.DeleteOptionsDelete),
		// 								},
		// 							},
		// 						},
		// 						NetworkAPIVersion: to.Ptr(armcompute.NetworkAPIVersion("2022-05-01")),
		// 						NetworkInterfaceConfigurations: []*armcompute.VirtualMachineNetworkInterfaceConfiguration{
		// 							{
		// 								Name: to.Ptr("aaaaaaaa"),
		// 								Properties: &armcompute.VirtualMachineNetworkInterfaceConfigurationProperties{
		// 									Primary: to.Ptr(true),
		// 									DeleteOption: to.Ptr(armcompute.DeleteOptionsDelete),
		// 									EnableAcceleratedNetworking: to.Ptr(true),
		// 									DisableTCPStateTracking: to.Ptr(true),
		// 									EnableFpga: to.Ptr(true),
		// 									EnableIPForwarding: to.Ptr(true),
		// 									NetworkSecurityGroup: &armcompute.SubResource{
		// 										ID: to.Ptr("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/availabilitySets/{availabilitySetName}"),
		// 									},
		// 									DNSSettings: &armcompute.VirtualMachineNetworkInterfaceDNSSettingsConfiguration{
		// 										DNSServers: []*string{
		// 											to.Ptr("aaaaaaaaaaaaaaaaaaaaaaaa"),
		// 										},
		// 									},
		// 									IPConfigurations: []*armcompute.VirtualMachineNetworkInterfaceIPConfiguration{
		// 										{
		// 											Name: to.Ptr("aaaaaaaa"),
		// 											Properties: &armcompute.VirtualMachineNetworkInterfaceIPConfigurationProperties{
		// 												Subnet: &armcompute.SubResource{
		// 													ID: to.Ptr("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/availabilitySets/{availabilitySetName}"),
		// 												},
		// 												Primary: to.Ptr(true),
		// 												PublicIPAddressConfiguration: &armcompute.VirtualMachinePublicIPAddressConfiguration{
		// 													Name: to.Ptr("aaaaaaaaaaaaaaaaaa"),
		// 													Properties: &armcompute.VirtualMachinePublicIPAddressConfigurationProperties{
		// 														IdleTimeoutInMinutes: to.Ptr[int32](23),
		// 														DeleteOption: to.Ptr(armcompute.DeleteOptionsDelete),
		// 														DNSSettings: &armcompute.VirtualMachinePublicIPAddressDNSSettingsConfiguration{
		// 															DomainNameLabel: to.Ptr("aaaaa"),
		// 															DomainNameLabelScope: to.Ptr(armcompute.DomainNameLabelScopeTypesSubscriptionReuse),
		// 														},
		// 														IPTags: []*armcompute.VirtualMachineIPTag{
		// 															{
		// 																IPTagType: to.Ptr("aaaaa"),
		// 																Tag: to.Ptr("aaaaaaaaaaaaaaaaaaaaaaaaaaa"),
		// 															},
		// 														},
		// 														PublicIPPrefix: &armcompute.SubResource{
		// 															ID: to.Ptr("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/availabilitySets/{availabilitySetName}"),
		// 														},
		// 														PublicIPAddressVersion: to.Ptr(armcompute.IPVersionsIPv4),
		// 														PublicIPAllocationMethod: to.Ptr(armcompute.PublicIPAllocationMethodDynamic),
		// 													},
		// 													SKU: &armcompute.PublicIPAddressSKU{
		// 														Name: to.Ptr(armcompute.PublicIPAddressSKUNameBasic),
		// 														Tier: to.Ptr(armcompute.PublicIPAddressSKUTierRegional),
		// 													},
		// 												},
		// 												PrivateIPAddressVersion: to.Ptr(armcompute.IPVersionsIPv4),
		// 												ApplicationSecurityGroups: []*armcompute.SubResource{
		// 													{
		// 														ID: to.Ptr("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/availabilitySets/{availabilitySetName}"),
		// 													},
		// 												},
		// 												ApplicationGatewayBackendAddressPools: []*armcompute.SubResource{
		// 													{
		// 														ID: to.Ptr("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/availabilitySets/{availabilitySetName}"),
		// 													},
		// 												},
		// 												LoadBalancerBackendAddressPools: []*armcompute.SubResource{
		// 													{
		// 														ID: to.Ptr("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/availabilitySets/{availabilitySetName}"),
		// 													},
		// 												},
		// 											},
		// 										},
		// 									},
		// 									DscpConfiguration: &armcompute.SubResource{
		// 										ID: to.Ptr("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/availabilitySets/{availabilitySetName}"),
		// 									},
		// 								},
		// 							},
		// 						},
		// 					},
		// 					ProvisioningState: to.Ptr("Succeeded"),
		// 					AdditionalCapabilities: &armcompute.AdditionalCapabilities{
		// 						UltraSSDEnabled: to.Ptr(true),
		// 						HibernationEnabled: to.Ptr(true),
		// 					},
		// 					SecurityProfile: &armcompute.SecurityProfile{
		// 						UefiSettings: &armcompute.UefiSettings{
		// 							SecureBootEnabled: to.Ptr(true),
		// 							VTpmEnabled: to.Ptr(true),
		// 						},
		// 						EncryptionAtHost: to.Ptr(true),
		// 						SecurityType: to.Ptr(armcompute.SecurityTypesTrustedLaunch),
		// 					},
		// 					DiagnosticsProfile: &armcompute.DiagnosticsProfile{
		// 						BootDiagnostics: &armcompute.BootDiagnostics{
		// 							Enabled: to.Ptr(true),
		// 							StorageURI: to.Ptr("aaaaaaaaaaaaaaaaaaaaa"),
		// 						},
		// 					},
		// 					VirtualMachineScaleSet: &armcompute.SubResource{
		// 						ID: to.Ptr("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/availabilitySets/{availabilitySetName}"),
		// 					},
		// 					ProximityPlacementGroup: &armcompute.SubResource{
		// 						ID: to.Ptr("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/availabilitySets/{availabilitySetName}"),
		// 					},
		// 					Priority: to.Ptr(armcompute.VirtualMachinePriorityTypesRegular),
		// 					EvictionPolicy: to.Ptr(armcompute.VirtualMachineEvictionPolicyTypesDeallocate),
		// 					BillingProfile: &armcompute.BillingProfile{
		// 						MaxPrice: to.Ptr[float64](26),
		// 					},
		// 					Host: &armcompute.SubResource{
		// 						ID: to.Ptr("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/availabilitySets/{availabilitySetName}"),
		// 					},
		// 					HostGroup: &armcompute.SubResource{
		// 						ID: to.Ptr("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/availabilitySets/{availabilitySetName}"),
		// 					},
		// 					LicenseType: to.Ptr("aaaaaaaaaaaaaa"),
		// 					ExtensionsTimeBudget: to.Ptr("aaaaaaaaaaaaaaaaaaaaaa"),
		// 					PlatformFaultDomain: to.Ptr[int32](11),
		// 					ScheduledEventsProfile: &armcompute.ScheduledEventsProfile{
		// 						TerminateNotificationProfile: &armcompute.TerminateNotificationProfile{
		// 							NotBeforeTimeout: to.Ptr("PT10M"),
		// 							Enable: to.Ptr(true),
		// 						},
		// 						OSImageNotificationProfile: &armcompute.OSImageNotificationProfile{
		// 							NotBeforeTimeout: to.Ptr("PT15M"),
		// 							Enable: to.Ptr(true),
		// 						},
		// 					},
		// 					UserData: to.Ptr("aaaaaaaaaaaaaaaaaaa"),
		// 					CapacityReservation: &armcompute.CapacityReservationProfile{
		// 						CapacityReservationGroup: &armcompute.SubResource{
		// 							ID: to.Ptr("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/availabilitySets/{availabilitySetName}"),
		// 						},
		// 					},
		// 					ApplicationProfile: &armcompute.ApplicationProfile{
		// 						GalleryApplications: []*armcompute.VMGalleryApplication{
		// 							{
		// 								Tags: to.Ptr("aaaaa"),
		// 								Order: to.Ptr[int32](4),
		// 								PackageReferenceID: to.Ptr("aaaaaaaaaaaaaaaaaaaaaa"),
		// 								ConfigurationReference: to.Ptr("aaaaaaaaaaaaaaaaaaaaaaaaaaaaaa"),
		// 							},
		// 						},
		// 					},
		// 					TimeCreated: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2022-01-14T16:43:41.685Z"); return t}()),
		// 				},
		// 				Type: to.Ptr("Microsoft.Compute/virtualMachines"),
		// 				Location: to.Ptr("eastus"),
		// 				Tags: map[string]*string{
		// 				},
		// 				ID: to.Ptr("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/virtualMachines/{virtualMachineName}"),
		// 				Name: to.Ptr("{virtualMachineName}"),
		// 				Plan: &armcompute.Plan{
		// 					Name: to.Ptr("aaaaaaaaaaaaaaaaaa"),
		// 					Publisher: to.Ptr("aaaaaaaaaaaaaaaaa"),
		// 					Product: to.Ptr("aaaaaaaaaaaaaaaaaaaaaa"),
		// 					PromotionCode: to.Ptr("aaaaaaaaaaaaaa"),
		// 				},
		// 				Resources: []*armcompute.VirtualMachineExtension{
		// 					{
		// 						Properties: &armcompute.VirtualMachineExtensionProperties{
		// 							ForceUpdateTag: to.Ptr("aaaaaaaaaaaaaaa"),
		// 							Publisher: to.Ptr("aaaaaaaaaaaaaaaa"),
		// 							Type: to.Ptr("aaaaaaaa"),
		// 							TypeHandlerVersion: to.Ptr("aaaaaaaaaaaa"),
		// 							AutoUpgradeMinorVersion: to.Ptr(true),
		// 							EnableAutomaticUpgrade: to.Ptr(true),
		// 							Settings: map[string]any{
		// 							},
		// 							ProtectedSettings: map[string]any{
		// 							},
		// 							ProvisioningState: to.Ptr("aaa"),
		// 							SuppressFailures: to.Ptr(true),
		// 							ProtectedSettingsFromKeyVault: &armcompute.KeyVaultSecretReference{
		// 								SourceVault: &armcompute.SubResource{
		// 									ID: to.Ptr("/subscriptions/a53f7094-a16c-47af-abe4-b05c05d0d79a/resourceGroups/myResourceGroup/providers/Microsoft.KeyVault/vaults/kvName"),
		// 								},
		// 								SecretURL: to.Ptr("https://kvName.vault.azure.net/secrets/secretName/79b88b3a6f5440ffb2e73e44a0db712e"),
		// 							},
		// 						},
		// 						ID: to.Ptr("aaaaaaaaaaaaaaaaaaaaaaaaaaaaa"),
		// 						Name: to.Ptr("aaaaaaaaaaaaa"),
		// 						Type: to.Ptr("aaaaaaaaaaaaaaaaaaaaaaaaaaa"),
		// 						Location: to.Ptr("aaaaaaaaaaaaaaaa"),
		// 						Tags: map[string]*string{
		// 							"key9428": to.Ptr("aaaaaaa"),
		// 						},
		// 					},
		// 				},
		// 				Identity: &armcompute.VirtualMachineIdentity{
		// 					PrincipalID: to.Ptr("aaaaaaaaaaaaaaaa"),
		// 					TenantID: to.Ptr("aaaaa"),
		// 					Type: to.Ptr(armcompute.ResourceIdentityTypeSystemAssigned),
		// 					UserAssignedIdentities: map[string]*armcompute.UserAssignedIdentitiesValue{
		// 						"key5688": &armcompute.UserAssignedIdentitiesValue{
		// 							PrincipalID: to.Ptr("aaaaaaaaaaaaaaa"),
		// 							ClientID: to.Ptr("aaaaaaaaaaa"),
		// 						},
		// 					},
		// 				},
		// 				Zones: []*string{
		// 					to.Ptr("aaaaaa"),
		// 				},
		// 				ExtendedLocation: &armcompute.ExtendedLocation{
		// 					Name: to.Ptr("aaaa"),
		// 					Type: to.Ptr(armcompute.ExtendedLocationTypesEdgeZone),
		// 				},
		// 			},
		// 		},
		// 		NextLink: to.Ptr("a://example.com/a"),
		// 	},
		// }
	}
}
