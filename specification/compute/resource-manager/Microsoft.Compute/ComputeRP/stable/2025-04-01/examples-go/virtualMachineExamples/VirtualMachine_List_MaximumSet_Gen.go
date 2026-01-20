package armcompute_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/compute/armcompute/v7"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/ccea12be6cd5e64743499d46f7a9c8b60df52db1/specification/compute/resource-manager/Microsoft.Compute/ComputeRP/stable/2025-04-01/examples/virtualMachineExamples/VirtualMachine_List_MaximumSet_Gen.json
func ExampleVirtualMachinesClient_NewListPager_virtualMachineListMaximumSetGen() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcompute.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewVirtualMachinesClient().NewListPager("rgcompute", &armcompute.VirtualMachinesClientListOptions{Filter: to.Ptr("aaaaaaaaaaaaaaaaaaaaaaa"),
		Expand: nil,
	})
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
		// page.VirtualMachineListResult = armcompute.VirtualMachineListResult{
		// 	Value: []*armcompute.VirtualMachine{
		// 		{
		// 			Name: to.Ptr("{virtualMachineName}"),
		// 			Type: to.Ptr("Microsoft.Compute/virtualMachines"),
		// 			ID: to.Ptr("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/virtualMachines/{virtualMachineName}"),
		// 			Location: to.Ptr("eastus"),
		// 			Tags: map[string]*string{
		// 			},
		// 			ExtendedLocation: &armcompute.ExtendedLocation{
		// 				Name: to.Ptr("aaaa"),
		// 				Type: to.Ptr(armcompute.ExtendedLocationTypesEdgeZone),
		// 			},
		// 			Identity: &armcompute.VirtualMachineIdentity{
		// 				Type: to.Ptr(armcompute.ResourceIdentityTypeSystemAssigned),
		// 				PrincipalID: to.Ptr("aaaaaaaaaaaaaaaa"),
		// 				TenantID: to.Ptr("aaaaa"),
		// 				UserAssignedIdentities: map[string]*armcompute.UserAssignedIdentitiesValue{
		// 					"key5688": &armcompute.UserAssignedIdentitiesValue{
		// 						ClientID: to.Ptr("aaaaaaaaaaa"),
		// 						PrincipalID: to.Ptr("aaaaaaaaaaaaaaa"),
		// 					},
		// 				},
		// 			},
		// 			Plan: &armcompute.Plan{
		// 				Name: to.Ptr("aaaaaaaaaaaaaaaaaa"),
		// 				Product: to.Ptr("aaaaaaaaaaaaaaaaaaaaaa"),
		// 				PromotionCode: to.Ptr("aaaaaaaaaaaaaa"),
		// 				Publisher: to.Ptr("aaaaaaaaaaaaaaaaa"),
		// 			},
		// 			Properties: &armcompute.VirtualMachineProperties{
		// 				AdditionalCapabilities: &armcompute.AdditionalCapabilities{
		// 					HibernationEnabled: to.Ptr(true),
		// 					UltraSSDEnabled: to.Ptr(true),
		// 				},
		// 				ApplicationProfile: &armcompute.ApplicationProfile{
		// 					GalleryApplications: []*armcompute.VMGalleryApplication{
		// 						{
		// 							ConfigurationReference: to.Ptr("aaaaaaaaaaaaaaaaaaaaaaaaaaaaaa"),
		// 							Order: to.Ptr[int32](4),
		// 							PackageReferenceID: to.Ptr("aaaaaaaaaaaaaaaaaaaaaa"),
		// 							Tags: to.Ptr("aaaaa"),
		// 					}},
		// 				},
		// 				AvailabilitySet: &armcompute.SubResource{
		// 					ID: to.Ptr("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/availabilitySets/{availabilitySetName}"),
		// 				},
		// 				BillingProfile: &armcompute.BillingProfile{
		// 					MaxPrice: to.Ptr[float64](26),
		// 				},
		// 				CapacityReservation: &armcompute.CapacityReservationProfile{
		// 					CapacityReservationGroup: &armcompute.SubResource{
		// 						ID: to.Ptr("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/availabilitySets/{availabilitySetName}"),
		// 					},
		// 				},
		// 				DiagnosticsProfile: &armcompute.DiagnosticsProfile{
		// 					BootDiagnostics: &armcompute.BootDiagnostics{
		// 						Enabled: to.Ptr(true),
		// 						StorageURI: to.Ptr("aaaaaaaaaaaaaaaaaaaaa"),
		// 					},
		// 				},
		// 				EvictionPolicy: to.Ptr(armcompute.VirtualMachineEvictionPolicyTypesDeallocate),
		// 				ExtensionsTimeBudget: to.Ptr("aaaaaaaaaaaaaaaaaaaaaaaaaaaa"),
		// 				HardwareProfile: &armcompute.HardwareProfile{
		// 					VMSize: to.Ptr(armcompute.VirtualMachineSizeTypesStandardA0),
		// 					VMSizeProperties: &armcompute.VMSizeProperties{
		// 						VCPUsAvailable: to.Ptr[int32](7),
		// 						VCPUsPerCore: to.Ptr[int32](14),
		// 					},
		// 				},
		// 				Host: &armcompute.SubResource{
		// 					ID: to.Ptr("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/availabilitySets/{availabilitySetName}"),
		// 				},
		// 				HostGroup: &armcompute.SubResource{
		// 					ID: to.Ptr("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/availabilitySets/{availabilitySetName}"),
		// 				},
		// 				LicenseType: to.Ptr("aaaaaaaaaaaaaaa"),
		// 				NetworkProfile: &armcompute.NetworkProfile{
		// 					NetworkAPIVersion: to.Ptr(armcompute.NetworkAPIVersion("2022-05-01")),
		// 					NetworkInterfaceConfigurations: []*armcompute.VirtualMachineNetworkInterfaceConfiguration{
		// 						{
		// 							Name: to.Ptr("aaaaaaaa"),
		// 							Properties: &armcompute.VirtualMachineNetworkInterfaceConfigurationProperties{
		// 								DeleteOption: to.Ptr(armcompute.DeleteOptionsDelete),
		// 								DisableTCPStateTracking: to.Ptr(true),
		// 								DNSSettings: &armcompute.VirtualMachineNetworkInterfaceDNSSettingsConfiguration{
		// 									DNSServers: []*string{
		// 										to.Ptr("aaaaaaaaaaaaaaaaaaaaaaaa")},
		// 									},
		// 									DscpConfiguration: &armcompute.SubResource{
		// 										ID: to.Ptr("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/availabilitySets/{availabilitySetName}"),
		// 									},
		// 									EnableAcceleratedNetworking: to.Ptr(true),
		// 									EnableFpga: to.Ptr(true),
		// 									EnableIPForwarding: to.Ptr(true),
		// 									IPConfigurations: []*armcompute.VirtualMachineNetworkInterfaceIPConfiguration{
		// 										{
		// 											Name: to.Ptr("aaaaaaaa"),
		// 											Properties: &armcompute.VirtualMachineNetworkInterfaceIPConfigurationProperties{
		// 												ApplicationGatewayBackendAddressPools: []*armcompute.SubResource{
		// 													{
		// 														ID: to.Ptr("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/availabilitySets/{availabilitySetName}"),
		// 												}},
		// 												ApplicationSecurityGroups: []*armcompute.SubResource{
		// 													{
		// 														ID: to.Ptr("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/availabilitySets/{availabilitySetName}"),
		// 												}},
		// 												LoadBalancerBackendAddressPools: []*armcompute.SubResource{
		// 													{
		// 														ID: to.Ptr("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/availabilitySets/{availabilitySetName}"),
		// 												}},
		// 												Primary: to.Ptr(true),
		// 												PrivateIPAddressVersion: to.Ptr(armcompute.IPVersionsIPv4),
		// 												PublicIPAddressConfiguration: &armcompute.VirtualMachinePublicIPAddressConfiguration{
		// 													Name: to.Ptr("aaaaaaaaaaaaaaaaaa"),
		// 													Properties: &armcompute.VirtualMachinePublicIPAddressConfigurationProperties{
		// 														DeleteOption: to.Ptr(armcompute.DeleteOptionsDelete),
		// 														DNSSettings: &armcompute.VirtualMachinePublicIPAddressDNSSettingsConfiguration{
		// 															DomainNameLabel: to.Ptr("aaaaa"),
		// 														},
		// 														IdleTimeoutInMinutes: to.Ptr[int32](23),
		// 														IPTags: []*armcompute.VirtualMachineIPTag{
		// 															{
		// 																IPTagType: to.Ptr("aaaaa"),
		// 																Tag: to.Ptr("aaaaaaaaaaaaaaaaaaaaaaaaaaa"),
		// 														}},
		// 														PublicIPAddressVersion: to.Ptr(armcompute.IPVersionsIPv4),
		// 														PublicIPAllocationMethod: to.Ptr(armcompute.PublicIPAllocationMethodDynamic),
		// 														PublicIPPrefix: &armcompute.SubResource{
		// 															ID: to.Ptr("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/availabilitySets/{availabilitySetName}"),
		// 														},
		// 													},
		// 													SKU: &armcompute.PublicIPAddressSKU{
		// 														Name: to.Ptr(armcompute.PublicIPAddressSKUNameBasic),
		// 														Tier: to.Ptr(armcompute.PublicIPAddressSKUTierRegional),
		// 													},
		// 												},
		// 												Subnet: &armcompute.SubResource{
		// 													ID: to.Ptr("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/availabilitySets/{availabilitySetName}"),
		// 												},
		// 											},
		// 									}},
		// 									NetworkSecurityGroup: &armcompute.SubResource{
		// 										ID: to.Ptr("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/availabilitySets/{availabilitySetName}"),
		// 									},
		// 									Primary: to.Ptr(true),
		// 								},
		// 						}},
		// 						NetworkInterfaces: []*armcompute.NetworkInterfaceReference{
		// 							{
		// 								ID: to.Ptr("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/networkInterfaces/{networkInterfaceName}"),
		// 								Properties: &armcompute.NetworkInterfaceReferenceProperties{
		// 									DeleteOption: to.Ptr(armcompute.DeleteOptionsDelete),
		// 									Primary: to.Ptr(true),
		// 								},
		// 						}},
		// 					},
		// 					OSProfile: &armcompute.OSProfile{
		// 						AdminUsername: to.Ptr("Foo12"),
		// 						AllowExtensionOperations: to.Ptr(true),
		// 						ComputerName: to.Ptr("Test"),
		// 						CustomData: to.Ptr("aaaa"),
		// 						LinuxConfiguration: &armcompute.LinuxConfiguration{
		// 							DisablePasswordAuthentication: to.Ptr(true),
		// 							PatchSettings: &armcompute.LinuxPatchSettings{
		// 								AssessmentMode: to.Ptr(armcompute.LinuxPatchAssessmentModeImageDefault),
		// 								PatchMode: to.Ptr(armcompute.LinuxVMGuestPatchModeImageDefault),
		// 							},
		// 							ProvisionVMAgent: to.Ptr(true),
		// 							SSH: &armcompute.SSHConfiguration{
		// 								PublicKeys: []*armcompute.SSHPublicKey{
		// 									{
		// 										Path: to.Ptr("aaaaaaaaaaaaaaaaaaaaaa"),
		// 										KeyData: to.Ptr("aaa"),
		// 								}},
		// 							},
		// 						},
		// 						RequireGuestProvisionSignal: to.Ptr(true),
		// 						Secrets: []*armcompute.VaultSecretGroup{
		// 						},
		// 						WindowsConfiguration: &armcompute.WindowsConfiguration{
		// 							AdditionalUnattendContent: []*armcompute.AdditionalUnattendContent{
		// 								{
		// 									ComponentName: to.Ptr("Microsoft-Windows-Shell-Setup"),
		// 									Content: to.Ptr("aaaaaaaaaaaaaaaaaaaaaaaaaa"),
		// 									PassName: to.Ptr("OobeSystem"),
		// 									SettingName: to.Ptr(armcompute.SettingNamesAutoLogon),
		// 							}},
		// 							EnableAutomaticUpdates: to.Ptr(true),
		// 							PatchSettings: &armcompute.PatchSettings{
		// 								AssessmentMode: to.Ptr(armcompute.WindowsPatchAssessmentModeImageDefault),
		// 								EnableHotpatching: to.Ptr(true),
		// 								PatchMode: to.Ptr(armcompute.WindowsVMGuestPatchModeManual),
		// 							},
		// 							ProvisionVMAgent: to.Ptr(true),
		// 							TimeZone: to.Ptr("aaaaaaaaaaaaaaaaaaaaaa"),
		// 							WinRM: &armcompute.WinRMConfiguration{
		// 								Listeners: []*armcompute.WinRMListener{
		// 									{
		// 										CertificateURL: to.Ptr("aaaaaaaaaaaaaaaaaaaaa"),
		// 										Protocol: to.Ptr(armcompute.ProtocolTypesHTTP),
		// 								}},
		// 							},
		// 						},
		// 					},
		// 					PlatformFaultDomain: to.Ptr[int32](8),
		// 					Priority: to.Ptr(armcompute.VirtualMachinePriorityTypesRegular),
		// 					ProvisioningState: to.Ptr("Succeeded"),
		// 					ProximityPlacementGroup: &armcompute.SubResource{
		// 						ID: to.Ptr("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/availabilitySets/{availabilitySetName}"),
		// 					},
		// 					ScheduledEventsProfile: &armcompute.ScheduledEventsProfile{
		// 						OSImageNotificationProfile: &armcompute.OSImageNotificationProfile{
		// 							Enable: to.Ptr(true),
		// 							NotBeforeTimeout: to.Ptr("PT15M"),
		// 						},
		// 						TerminateNotificationProfile: &armcompute.TerminateNotificationProfile{
		// 							Enable: to.Ptr(true),
		// 							NotBeforeTimeout: to.Ptr("PT10M"),
		// 						},
		// 					},
		// 					SecurityProfile: &armcompute.SecurityProfile{
		// 						EncryptionAtHost: to.Ptr(true),
		// 						SecurityType: to.Ptr(armcompute.SecurityTypesTrustedLaunch),
		// 						UefiSettings: &armcompute.UefiSettings{
		// 							SecureBootEnabled: to.Ptr(true),
		// 							VTpmEnabled: to.Ptr(true),
		// 						},
		// 					},
		// 					StorageProfile: &armcompute.StorageProfile{
		// 						DataDisks: []*armcompute.DataDisk{
		// 						},
		// 						ImageReference: &armcompute.ImageReference{
		// 							ID: to.Ptr("aaaaaaaaaaaaaaaaaaaaaaaaaaaaaa"),
		// 							CommunityGalleryImageID: to.Ptr("aaaa"),
		// 							ExactVersion: to.Ptr("aaaaaaaaaaaaa"),
		// 							Offer: to.Ptr("WindowsServer"),
		// 							Publisher: to.Ptr("MicrosoftWindowsServer"),
		// 							SharedGalleryImageID: to.Ptr("aaaaaaaaaaaaaaa"),
		// 							SKU: to.Ptr("2012-R2-Datacenter"),
		// 							Version: to.Ptr("4.127.20170406"),
		// 						},
		// 						OSDisk: &armcompute.OSDisk{
		// 							Name: to.Ptr("test"),
		// 							Caching: to.Ptr(armcompute.CachingTypesNone),
		// 							CreateOption: to.Ptr(armcompute.DiskCreateOptionTypesFromImage),
		// 							DeleteOption: to.Ptr(armcompute.DiskDeleteOptionTypesDelete),
		// 							DiffDiskSettings: &armcompute.DiffDiskSettings{
		// 								Option: to.Ptr(armcompute.DiffDiskOptionsLocal),
		// 								Placement: to.Ptr(armcompute.DiffDiskPlacementCacheDisk),
		// 							},
		// 							DiskSizeGB: to.Ptr[int32](127),
		// 							EncryptionSettings: &armcompute.DiskEncryptionSettings{
		// 								DiskEncryptionKey: &armcompute.KeyVaultSecretReference{
		// 									SecretURL: to.Ptr("aaaaaaaaa"),
		// 									SourceVault: &armcompute.SubResource{
		// 										ID: to.Ptr("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/availabilitySets/{availabilitySetName}"),
		// 									},
		// 								},
		// 								Enabled: to.Ptr(true),
		// 								KeyEncryptionKey: &armcompute.KeyVaultKeyReference{
		// 									KeyURL: to.Ptr("aaaaaaaaaaaaa"),
		// 									SourceVault: &armcompute.SubResource{
		// 										ID: to.Ptr("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/availabilitySets/{availabilitySetName}"),
		// 									},
		// 								},
		// 							},
		// 							Image: &armcompute.VirtualHardDisk{
		// 								URI: to.Ptr("https://{storageAccountName}.blob.core.windows.net/{containerName}/{vhdName}.vhd"),
		// 							},
		// 							ManagedDisk: &armcompute.ManagedDiskParameters{
		// 								ID: to.Ptr("/subscriptions/{subscription-id}/resourceGroups/myResourceGroup/providers/Microsoft.Compute/disks/testingexcludedisk_OsDisk_1_74cdaedcea50483d9833c96adefa100f"),
		// 								DiskEncryptionSet: &armcompute.DiskEncryptionSetParameters{
		// 									ID: to.Ptr("aaaaaaaaaaaaaaaaaaaaaaaaaaaa"),
		// 								},
		// 								SecurityProfile: &armcompute.VMDiskSecurityProfile{
		// 									DiskEncryptionSet: &armcompute.DiskEncryptionSetParameters{
		// 										ID: to.Ptr("aaaaaaaaaaaaaaaaaaaaaaaaaaaa"),
		// 									},
		// 									SecurityEncryptionType: to.Ptr(armcompute.SecurityEncryptionTypesVMGuestStateOnly),
		// 								},
		// 								StorageAccountType: to.Ptr(armcompute.StorageAccountTypesStandardLRS),
		// 							},
		// 							OSType: to.Ptr(armcompute.OperatingSystemTypesWindows),
		// 							Vhd: &armcompute.VirtualHardDisk{
		// 								URI: to.Ptr("https://{storageAccountName}.blob.core.windows.net/{containerName}/{vhdName}.vhd"),
		// 							},
		// 							WriteAcceleratorEnabled: to.Ptr(true),
		// 						},
		// 					},
		// 					TimeCreated: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2022-01-14T16:43:41.683Z"); return t}()),
		// 					UserData: to.Ptr("aaa"),
		// 					VirtualMachineScaleSet: &armcompute.SubResource{
		// 						ID: to.Ptr("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/availabilitySets/{availabilitySetName}"),
		// 					},
		// 					VMID: to.Ptr("{vmId}"),
		// 				},
		// 				Resources: []*armcompute.VirtualMachineExtension{
		// 					{
		// 						Name: to.Ptr("aaaaaaaaaaaaa"),
		// 						Type: to.Ptr("aaaaaaaaaaaaaaaaaaaaaaaaaaa"),
		// 						ID: to.Ptr("aaaaaaaaaaaaaaaaaaaaaaaaaaaaa"),
		// 						Location: to.Ptr("aaaaaaaaaaaaaaaa"),
		// 						Tags: map[string]*string{
		// 							"key9428": to.Ptr("aaaaaaa"),
		// 						},
		// 						Properties: &armcompute.VirtualMachineExtensionProperties{
		// 							Type: to.Ptr("aaaaaaaa"),
		// 							AutoUpgradeMinorVersion: to.Ptr(true),
		// 							EnableAutomaticUpgrade: to.Ptr(true),
		// 							ForceUpdateTag: to.Ptr("aaaaaaaaaaaaaaa"),
		// 							InstanceView: &armcompute.VirtualMachineExtensionInstanceView{
		// 								Name: to.Ptr("aaaaaaaaaaaaaaaaaaaaaaaa"),
		// 								Type: to.Ptr("aaaaaaaaaaaaaaaaaaaaaaaa"),
		// 								Statuses: []*armcompute.InstanceViewStatus{
		// 									{
		// 										Code: to.Ptr("aaa"),
		// 										DisplayStatus: to.Ptr("aaaaaaaaaaaaaaaaaaaaaaaa"),
		// 										Level: to.Ptr(armcompute.StatusLevelTypesInfo),
		// 										Message: to.Ptr("aaaaaaaaaaaaaaaaaaa"),
		// 										Time: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2022-01-14T16:43:41.657Z"); return t}()),
		// 								}},
		// 								Substatuses: []*armcompute.InstanceViewStatus{
		// 									{
		// 										Code: to.Ptr("aaa"),
		// 										DisplayStatus: to.Ptr("aaaaaaaaaaaaaaaaaaaaaaaa"),
		// 										Level: to.Ptr(armcompute.StatusLevelTypesInfo),
		// 										Message: to.Ptr("aaaaaaaaaaaaaaaaaaa"),
		// 										Time: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2022-01-14T16:43:41.657Z"); return t}()),
		// 								}},
		// 								TypeHandlerVersion: to.Ptr("aaaaaa"),
		// 							},
		// 							ProtectedSettings: map[string]any{
		// 							},
		// 							ProtectedSettingsFromKeyVault: &armcompute.KeyVaultSecretReference{
		// 								SecretURL: to.Ptr("https://kvName.vault.azure.net/secrets/secretName/79b88b3a6f5440ffb2e73e44a0db712e"),
		// 								SourceVault: &armcompute.SubResource{
		// 									ID: to.Ptr("/subscriptions/a53f7094-a16c-47af-abe4-b05c05d0d79a/resourceGroups/myResourceGroup/providers/Microsoft.KeyVault/vaults/kvName"),
		// 								},
		// 							},
		// 							ProvisioningState: to.Ptr("Succeeded"),
		// 							Publisher: to.Ptr("aaaaaaaaaaaaaaaa"),
		// 							Settings: map[string]any{
		// 							},
		// 							SuppressFailures: to.Ptr(true),
		// 							TypeHandlerVersion: to.Ptr("aaaaaaaaaaaa"),
		// 						},
		// 				}},
		// 				Zones: []*string{
		// 					to.Ptr("aaaaaaaaaaaaaaaaaaaaaaaaaaa")},
		// 				},
		// 				{
		// 					Name: to.Ptr("{virtualMachineName}"),
		// 					Type: to.Ptr("Microsoft.Compute/virtualMachines"),
		// 					ID: to.Ptr("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/virtualMachines/{virtualMachineName}"),
		// 					Location: to.Ptr("eastus"),
		// 					Tags: map[string]*string{
		// 					},
		// 					ExtendedLocation: &armcompute.ExtendedLocation{
		// 						Name: to.Ptr("aaaa"),
		// 						Type: to.Ptr(armcompute.ExtendedLocationTypesEdgeZone),
		// 					},
		// 					Identity: &armcompute.VirtualMachineIdentity{
		// 						Type: to.Ptr(armcompute.ResourceIdentityTypeSystemAssigned),
		// 						PrincipalID: to.Ptr("aaaaaaaaaaaaaaaa"),
		// 						TenantID: to.Ptr("aaaaa"),
		// 						UserAssignedIdentities: map[string]*armcompute.UserAssignedIdentitiesValue{
		// 							"key5688": &armcompute.UserAssignedIdentitiesValue{
		// 								ClientID: to.Ptr("aaaaaaaaaaa"),
		// 								PrincipalID: to.Ptr("aaaaaaaaaaaaaaa"),
		// 							},
		// 						},
		// 					},
		// 					Plan: &armcompute.Plan{
		// 						Name: to.Ptr("aaaaaaaaaaaaaaaaaa"),
		// 						Product: to.Ptr("aaaaaaaaaaaaaaaaaaaaaa"),
		// 						PromotionCode: to.Ptr("aaaaaaaaaaaaaa"),
		// 						Publisher: to.Ptr("aaaaaaaaaaaaaaaaa"),
		// 					},
		// 					Properties: &armcompute.VirtualMachineProperties{
		// 						AdditionalCapabilities: &armcompute.AdditionalCapabilities{
		// 							HibernationEnabled: to.Ptr(true),
		// 							UltraSSDEnabled: to.Ptr(true),
		// 						},
		// 						ApplicationProfile: &armcompute.ApplicationProfile{
		// 							GalleryApplications: []*armcompute.VMGalleryApplication{
		// 								{
		// 									ConfigurationReference: to.Ptr("aaaaaaaaaaaaaaaaaaaaaaaaaaaaaa"),
		// 									Order: to.Ptr[int32](4),
		// 									PackageReferenceID: to.Ptr("aaaaaaaaaaaaaaaaaaaaaa"),
		// 									Tags: to.Ptr("aaaaa"),
		// 							}},
		// 						},
		// 						AvailabilitySet: &armcompute.SubResource{
		// 							ID: to.Ptr("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/availabilitySets/{availabilitySetName}"),
		// 						},
		// 						BillingProfile: &armcompute.BillingProfile{
		// 							MaxPrice: to.Ptr[float64](26),
		// 						},
		// 						CapacityReservation: &armcompute.CapacityReservationProfile{
		// 							CapacityReservationGroup: &armcompute.SubResource{
		// 								ID: to.Ptr("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/availabilitySets/{availabilitySetName}"),
		// 							},
		// 						},
		// 						DiagnosticsProfile: &armcompute.DiagnosticsProfile{
		// 							BootDiagnostics: &armcompute.BootDiagnostics{
		// 								Enabled: to.Ptr(true),
		// 								StorageURI: to.Ptr("aaaaaaaaaaaaaaaaaaaaa"),
		// 							},
		// 						},
		// 						EvictionPolicy: to.Ptr(armcompute.VirtualMachineEvictionPolicyTypesDeallocate),
		// 						ExtensionsTimeBudget: to.Ptr("aaaaaaaaaaaaaaaaaaaaaa"),
		// 						HardwareProfile: &armcompute.HardwareProfile{
		// 							VMSize: to.Ptr(armcompute.VirtualMachineSizeTypesStandardA0),
		// 							VMSizeProperties: &armcompute.VMSizeProperties{
		// 								VCPUsAvailable: to.Ptr[int32](7),
		// 								VCPUsPerCore: to.Ptr[int32](14),
		// 							},
		// 						},
		// 						Host: &armcompute.SubResource{
		// 							ID: to.Ptr("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/availabilitySets/{availabilitySetName}"),
		// 						},
		// 						HostGroup: &armcompute.SubResource{
		// 							ID: to.Ptr("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/availabilitySets/{availabilitySetName}"),
		// 						},
		// 						InstanceView: &armcompute.VirtualMachineInstanceView{
		// 							AssignedHost: to.Ptr("aaaaaaaaaaaaaaaaaaaaaaaaaaaaa"),
		// 							BootDiagnostics: &armcompute.BootDiagnosticsInstanceView{
		// 								ConsoleScreenshotBlobURI: to.Ptr("aaaaaaaaaaaaaaaaaaaaaaa"),
		// 								SerialConsoleLogBlobURI: to.Ptr("aaaaaaaaaaaaaaaaaaa"),
		// 								Status: &armcompute.InstanceViewStatus{
		// 									Code: to.Ptr("aaa"),
		// 									DisplayStatus: to.Ptr("aaaaaaaaaaaaaaaaaaaaaaaa"),
		// 									Level: to.Ptr(armcompute.StatusLevelTypesInfo),
		// 									Message: to.Ptr("aaaaaaaaaaaaaaaaaaa"),
		// 									Time: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2022-01-14T16:43:41.657Z"); return t}()),
		// 								},
		// 							},
		// 							ComputerName: to.Ptr("aaaaaaaaaaaaaaaaaaaaaaaaaaaaa"),
		// 							Disks: []*armcompute.DiskInstanceView{
		// 								{
		// 									Name: to.Ptr("a"),
		// 									EncryptionSettings: []*armcompute.DiskEncryptionSettings{
		// 										{
		// 											DiskEncryptionKey: &armcompute.KeyVaultSecretReference{
		// 												SecretURL: to.Ptr("aaaaaaaaa"),
		// 												SourceVault: &armcompute.SubResource{
		// 													ID: to.Ptr("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/availabilitySets/{availabilitySetName}"),
		// 												},
		// 											},
		// 											Enabled: to.Ptr(true),
		// 											KeyEncryptionKey: &armcompute.KeyVaultKeyReference{
		// 												KeyURL: to.Ptr("aaaaaaaaaaaaa"),
		// 												SourceVault: &armcompute.SubResource{
		// 													ID: to.Ptr("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/availabilitySets/{availabilitySetName}"),
		// 												},
		// 											},
		// 									}},
		// 									Statuses: []*armcompute.InstanceViewStatus{
		// 										{
		// 											Code: to.Ptr("aaa"),
		// 											DisplayStatus: to.Ptr("aaaaaaaaaaaaaaaaaaaaaaaa"),
		// 											Level: to.Ptr(armcompute.StatusLevelTypesInfo),
		// 											Message: to.Ptr("aaaaaaaaaaaaaaaaaaa"),
		// 											Time: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2022-01-14T16:43:41.657Z"); return t}()),
		// 									}},
		// 							}},
		// 							HyperVGeneration: to.Ptr(armcompute.HyperVGenerationTypeV1),
		// 							IsVMInStandbyPool: to.Ptr(false),
		// 							MaintenanceRedeployStatus: &armcompute.MaintenanceRedeployStatus{
		// 								IsCustomerInitiatedMaintenanceAllowed: to.Ptr(true),
		// 								LastOperationMessage: to.Ptr("aa"),
		// 								LastOperationResultCode: to.Ptr(armcompute.MaintenanceOperationResultCodeTypesNone),
		// 								MaintenanceWindowEndTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2022-01-14T16:43:41.682Z"); return t}()),
		// 								MaintenanceWindowStartTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2022-01-14T16:43:41.682Z"); return t}()),
		// 								PreMaintenanceWindowEndTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2022-01-14T16:43:41.682Z"); return t}()),
		// 								PreMaintenanceWindowStartTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2022-01-14T16:43:41.682Z"); return t}()),
		// 							},
		// 							OSName: to.Ptr("aaaaaaaaaaa"),
		// 							OSVersion: to.Ptr("aaaaaaaaaaaaaa"),
		// 							PatchStatus: &armcompute.VirtualMachinePatchStatus{
		// 								AvailablePatchSummary: &armcompute.AvailablePatchSummary{
		// 									AssessmentActivityID: to.Ptr("aaaaaaaaaaaaaaaaaaa"),
		// 									CriticalAndSecurityPatchCount: to.Ptr[int32](22),
		// 									Error: &armcompute.APIError{
		// 										Code: to.Ptr("aaaaaa"),
		// 										Innererror: &armcompute.InnerError{
		// 											Errordetail: to.Ptr("aaaaaaaaaaaaaaaaaaaaaaaaaaaa"),
		// 											Exceptiontype: to.Ptr("aaaaaaaaaaaaaaaaaaaaaaaaaaaaa"),
		// 										},
		// 										Message: to.Ptr("aaa"),
		// 										Target: to.Ptr("aaaaaaaaa"),
		// 										Details: []*armcompute.APIErrorBase{
		// 											{
		// 												Code: to.Ptr("aaaaaaaaaaaaaaaaaaaaa"),
		// 												Message: to.Ptr("aaaaaaaaaaaaa"),
		// 												Target: to.Ptr("aaaaaa"),
		// 										}},
		// 									},
		// 									LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2022-01-14T16:43:41.683Z"); return t}()),
		// 									OtherPatchCount: to.Ptr[int32](25),
		// 									RebootPending: to.Ptr(true),
		// 									StartTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2022-01-14T16:43:41.682Z"); return t}()),
		// 									Status: to.Ptr(armcompute.PatchOperationStatusUnknown),
		// 								},
		// 								ConfigurationStatuses: []*armcompute.InstanceViewStatus{
		// 									{
		// 										Code: to.Ptr("aaa"),
		// 										DisplayStatus: to.Ptr("aaaaaaaaaaaaaaaaaaaaaaaa"),
		// 										Level: to.Ptr(armcompute.StatusLevelTypesInfo),
		// 										Message: to.Ptr("aaaaaaaaaaaaaaaaaaa"),
		// 										Time: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2022-01-14T16:43:41.657Z"); return t}()),
		// 								}},
		// 								LastPatchInstallationSummary: &armcompute.LastPatchInstallationSummary{
		// 									Error: &armcompute.APIError{
		// 										Code: to.Ptr("aaaaaa"),
		// 										Innererror: &armcompute.InnerError{
		// 											Errordetail: to.Ptr("aaaaaaaaaaaaaaaaaaaaaaaaaaaa"),
		// 											Exceptiontype: to.Ptr("aaaaaaaaaaaaaaaaaaaaaaaaaaaaa"),
		// 										},
		// 										Message: to.Ptr("aaa"),
		// 										Target: to.Ptr("aaaaaaaaa"),
		// 										Details: []*armcompute.APIErrorBase{
		// 											{
		// 												Code: to.Ptr("aaaaaaaaaaaaaaaaaaaaa"),
		// 												Message: to.Ptr("aaaaaaaaaaaaa"),
		// 												Target: to.Ptr("aaaaaa"),
		// 										}},
		// 									},
		// 									ExcludedPatchCount: to.Ptr[int32](1),
		// 									FailedPatchCount: to.Ptr[int32](30),
		// 									InstallationActivityID: to.Ptr("aaaaaa"),
		// 									InstalledPatchCount: to.Ptr[int32](28),
		// 									LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2022-01-14T16:43:41.683Z"); return t}()),
		// 									MaintenanceWindowExceeded: to.Ptr(true),
		// 									NotSelectedPatchCount: to.Ptr[int32](20),
		// 									PendingPatchCount: to.Ptr[int32](2),
		// 									StartTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2022-01-14T16:43:41.683Z"); return t}()),
		// 									Status: to.Ptr(armcompute.PatchOperationStatusUnknown),
		// 								},
		// 							},
		// 							PlatformFaultDomain: to.Ptr[int32](29),
		// 							PlatformUpdateDomain: to.Ptr[int32](1),
		// 							RdpThumbPrint: to.Ptr("aaaaaaaaaaaaaaaaaaaaaaa"),
		// 							Statuses: []*armcompute.InstanceViewStatus{
		// 								{
		// 									Code: to.Ptr("aaa"),
		// 									DisplayStatus: to.Ptr("aaaaaaaaaaaaaaaaaaaaaaaa"),
		// 									Level: to.Ptr(armcompute.StatusLevelTypesInfo),
		// 									Message: to.Ptr("aaaaaaaaaaaaaaaaaaa"),
		// 									Time: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2022-01-14T16:43:41.657Z"); return t}()),
		// 							}},
		// 							VMAgent: &armcompute.VirtualMachineAgentInstanceView{
		// 								ExtensionHandlers: []*armcompute.VirtualMachineExtensionHandlerInstanceView{
		// 									{
		// 										Type: to.Ptr("aaaaa"),
		// 										Status: &armcompute.InstanceViewStatus{
		// 											Code: to.Ptr("aaa"),
		// 											DisplayStatus: to.Ptr("aaaaaaaaaaaaaaaaaaaaaaaa"),
		// 											Level: to.Ptr(armcompute.StatusLevelTypesInfo),
		// 											Message: to.Ptr("aaaaaaaaaaaaaaaaaaa"),
		// 											Time: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2022-01-14T16:43:41.657Z"); return t}()),
		// 										},
		// 										TypeHandlerVersion: to.Ptr("aaaaaaaaaaaaaaaaaaa"),
		// 								}},
		// 								Statuses: []*armcompute.InstanceViewStatus{
		// 									{
		// 										Code: to.Ptr("aaa"),
		// 										DisplayStatus: to.Ptr("aaaaaaaaaaaaaaaaaaaaaaaa"),
		// 										Level: to.Ptr(armcompute.StatusLevelTypesInfo),
		// 										Message: to.Ptr("aaaaaaaaaaaaaaaaaaa"),
		// 										Time: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2022-01-14T16:43:41.657Z"); return t}()),
		// 								}},
		// 								VMAgentVersion: to.Ptr("aaaaaaaaaaaaaaaaaaaaaaaaaaaaa"),
		// 							},
		// 							VMHealth: &armcompute.VirtualMachineHealthStatus{
		// 								Status: &armcompute.InstanceViewStatus{
		// 									Code: to.Ptr("aaa"),
		// 									DisplayStatus: to.Ptr("aaaaaaaaaaaaaaaaaaaaaaaa"),
		// 									Level: to.Ptr(armcompute.StatusLevelTypesInfo),
		// 									Message: to.Ptr("aaaaaaaaaaaaaaaaaaa"),
		// 									Time: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2022-01-14T16:43:41.657Z"); return t}()),
		// 								},
		// 							},
		// 							Extensions: []*armcompute.VirtualMachineExtensionInstanceView{
		// 								{
		// 									Name: to.Ptr("aaaaaaaaaaaaaaaaaaaaaaaa"),
		// 									Type: to.Ptr("aaaaaaaaaaaaaaaaaaaaaaaa"),
		// 									Statuses: []*armcompute.InstanceViewStatus{
		// 										{
		// 											Code: to.Ptr("aaa"),
		// 											DisplayStatus: to.Ptr("aaaaaaaaaaaaaaaaaaaaaaaa"),
		// 											Level: to.Ptr(armcompute.StatusLevelTypesInfo),
		// 											Message: to.Ptr("aaaaaaaaaaaaaaaaaaa"),
		// 											Time: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2022-01-14T16:43:41.657Z"); return t}()),
		// 									}},
		// 									Substatuses: []*armcompute.InstanceViewStatus{
		// 										{
		// 											Code: to.Ptr("aaa"),
		// 											DisplayStatus: to.Ptr("aaaaaaaaaaaaaaaaaaaaaaaa"),
		// 											Level: to.Ptr(armcompute.StatusLevelTypesInfo),
		// 											Message: to.Ptr("aaaaaaaaaaaaaaaaaaa"),
		// 											Time: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2022-01-14T16:43:41.657Z"); return t}()),
		// 									}},
		// 									TypeHandlerVersion: to.Ptr("aaaaaa"),
		// 							}},
		// 						},
		// 						LicenseType: to.Ptr("aaaaaaaaaaaaaa"),
		// 						NetworkProfile: &armcompute.NetworkProfile{
		// 							NetworkAPIVersion: to.Ptr(armcompute.NetworkAPIVersion("2022-05-01")),
		// 							NetworkInterfaceConfigurations: []*armcompute.VirtualMachineNetworkInterfaceConfiguration{
		// 								{
		// 									Name: to.Ptr("aaaaaaaa"),
		// 									Properties: &armcompute.VirtualMachineNetworkInterfaceConfigurationProperties{
		// 										DeleteOption: to.Ptr(armcompute.DeleteOptionsDelete),
		// 										DisableTCPStateTracking: to.Ptr(true),
		// 										DNSSettings: &armcompute.VirtualMachineNetworkInterfaceDNSSettingsConfiguration{
		// 											DNSServers: []*string{
		// 												to.Ptr("aaaaaaaaaaaaaaaaaaaaaaaa")},
		// 											},
		// 											DscpConfiguration: &armcompute.SubResource{
		// 												ID: to.Ptr("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/availabilitySets/{availabilitySetName}"),
		// 											},
		// 											EnableAcceleratedNetworking: to.Ptr(true),
		// 											EnableFpga: to.Ptr(true),
		// 											EnableIPForwarding: to.Ptr(true),
		// 											IPConfigurations: []*armcompute.VirtualMachineNetworkInterfaceIPConfiguration{
		// 												{
		// 													Name: to.Ptr("aaaaaaaa"),
		// 													Properties: &armcompute.VirtualMachineNetworkInterfaceIPConfigurationProperties{
		// 														ApplicationGatewayBackendAddressPools: []*armcompute.SubResource{
		// 															{
		// 																ID: to.Ptr("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/availabilitySets/{availabilitySetName}"),
		// 														}},
		// 														ApplicationSecurityGroups: []*armcompute.SubResource{
		// 															{
		// 																ID: to.Ptr("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/availabilitySets/{availabilitySetName}"),
		// 														}},
		// 														LoadBalancerBackendAddressPools: []*armcompute.SubResource{
		// 															{
		// 																ID: to.Ptr("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/availabilitySets/{availabilitySetName}"),
		// 														}},
		// 														Primary: to.Ptr(true),
		// 														PrivateIPAddressVersion: to.Ptr(armcompute.IPVersionsIPv4),
		// 														PublicIPAddressConfiguration: &armcompute.VirtualMachinePublicIPAddressConfiguration{
		// 															Name: to.Ptr("aaaaaaaaaaaaaaaaaa"),
		// 															Properties: &armcompute.VirtualMachinePublicIPAddressConfigurationProperties{
		// 																DeleteOption: to.Ptr(armcompute.DeleteOptionsDelete),
		// 																DNSSettings: &armcompute.VirtualMachinePublicIPAddressDNSSettingsConfiguration{
		// 																	DomainNameLabel: to.Ptr("aaaaa"),
		// 																},
		// 																IdleTimeoutInMinutes: to.Ptr[int32](23),
		// 																IPTags: []*armcompute.VirtualMachineIPTag{
		// 																	{
		// 																		IPTagType: to.Ptr("aaaaa"),
		// 																		Tag: to.Ptr("aaaaaaaaaaaaaaaaaaaaaaaaaaa"),
		// 																}},
		// 																PublicIPAddressVersion: to.Ptr(armcompute.IPVersionsIPv4),
		// 																PublicIPAllocationMethod: to.Ptr(armcompute.PublicIPAllocationMethodDynamic),
		// 																PublicIPPrefix: &armcompute.SubResource{
		// 																	ID: to.Ptr("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/availabilitySets/{availabilitySetName}"),
		// 																},
		// 															},
		// 															SKU: &armcompute.PublicIPAddressSKU{
		// 																Name: to.Ptr(armcompute.PublicIPAddressSKUNameBasic),
		// 																Tier: to.Ptr(armcompute.PublicIPAddressSKUTierRegional),
		// 															},
		// 														},
		// 														Subnet: &armcompute.SubResource{
		// 															ID: to.Ptr("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/availabilitySets/{availabilitySetName}"),
		// 														},
		// 													},
		// 											}},
		// 											NetworkSecurityGroup: &armcompute.SubResource{
		// 												ID: to.Ptr("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/availabilitySets/{availabilitySetName}"),
		// 											},
		// 											Primary: to.Ptr(true),
		// 										},
		// 								}},
		// 								NetworkInterfaces: []*armcompute.NetworkInterfaceReference{
		// 									{
		// 										ID: to.Ptr("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/networkInterfaces/{networkInterfaceName}"),
		// 										Properties: &armcompute.NetworkInterfaceReferenceProperties{
		// 											DeleteOption: to.Ptr(armcompute.DeleteOptionsDelete),
		// 											Primary: to.Ptr(true),
		// 										},
		// 								}},
		// 							},
		// 							OSProfile: &armcompute.OSProfile{
		// 								AdminUsername: to.Ptr("Foo12"),
		// 								AllowExtensionOperations: to.Ptr(true),
		// 								ComputerName: to.Ptr("Test"),
		// 								CustomData: to.Ptr("aaaaaaaaaaaaaaaaaaaaaaa"),
		// 								LinuxConfiguration: &armcompute.LinuxConfiguration{
		// 									DisablePasswordAuthentication: to.Ptr(true),
		// 									PatchSettings: &armcompute.LinuxPatchSettings{
		// 										AssessmentMode: to.Ptr(armcompute.LinuxPatchAssessmentModeImageDefault),
		// 										PatchMode: to.Ptr(armcompute.LinuxVMGuestPatchModeImageDefault),
		// 									},
		// 									ProvisionVMAgent: to.Ptr(true),
		// 									SSH: &armcompute.SSHConfiguration{
		// 										PublicKeys: []*armcompute.SSHPublicKey{
		// 											{
		// 												Path: to.Ptr("aaaaaaaaaaaaaaaaaaaaaa"),
		// 												KeyData: to.Ptr("aaa"),
		// 										}},
		// 									},
		// 								},
		// 								RequireGuestProvisionSignal: to.Ptr(true),
		// 								Secrets: []*armcompute.VaultSecretGroup{
		// 								},
		// 								WindowsConfiguration: &armcompute.WindowsConfiguration{
		// 									AdditionalUnattendContent: []*armcompute.AdditionalUnattendContent{
		// 										{
		// 											ComponentName: to.Ptr("Microsoft-Windows-Shell-Setup"),
		// 											Content: to.Ptr("aaaaaaaaaaaaaaaaaaaaaaaaaa"),
		// 											PassName: to.Ptr("OobeSystem"),
		// 											SettingName: to.Ptr(armcompute.SettingNamesAutoLogon),
		// 									}},
		// 									EnableAutomaticUpdates: to.Ptr(true),
		// 									PatchSettings: &armcompute.PatchSettings{
		// 										AssessmentMode: to.Ptr(armcompute.WindowsPatchAssessmentModeImageDefault),
		// 										EnableHotpatching: to.Ptr(true),
		// 										PatchMode: to.Ptr(armcompute.WindowsVMGuestPatchModeManual),
		// 									},
		// 									ProvisionVMAgent: to.Ptr(true),
		// 									TimeZone: to.Ptr("aaaaaaaaaaaaaaaaaaaa"),
		// 									WinRM: &armcompute.WinRMConfiguration{
		// 										Listeners: []*armcompute.WinRMListener{
		// 											{
		// 												CertificateURL: to.Ptr("aaaaaaaaaaaaaaaaaaaaa"),
		// 												Protocol: to.Ptr(armcompute.ProtocolTypesHTTP),
		// 										}},
		// 									},
		// 								},
		// 							},
		// 							PlatformFaultDomain: to.Ptr[int32](11),
		// 							Priority: to.Ptr(armcompute.VirtualMachinePriorityTypesRegular),
		// 							ProvisioningState: to.Ptr("Succeeded"),
		// 							ProximityPlacementGroup: &armcompute.SubResource{
		// 								ID: to.Ptr("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/availabilitySets/{availabilitySetName}"),
		// 							},
		// 							ScheduledEventsProfile: &armcompute.ScheduledEventsProfile{
		// 								OSImageNotificationProfile: &armcompute.OSImageNotificationProfile{
		// 									Enable: to.Ptr(true),
		// 									NotBeforeTimeout: to.Ptr("PT15M"),
		// 								},
		// 								TerminateNotificationProfile: &armcompute.TerminateNotificationProfile{
		// 									Enable: to.Ptr(true),
		// 									NotBeforeTimeout: to.Ptr("PT10M"),
		// 								},
		// 							},
		// 							SecurityProfile: &armcompute.SecurityProfile{
		// 								EncryptionAtHost: to.Ptr(true),
		// 								SecurityType: to.Ptr(armcompute.SecurityTypesTrustedLaunch),
		// 								UefiSettings: &armcompute.UefiSettings{
		// 									SecureBootEnabled: to.Ptr(true),
		// 									VTpmEnabled: to.Ptr(true),
		// 								},
		// 							},
		// 							StorageProfile: &armcompute.StorageProfile{
		// 								DataDisks: []*armcompute.DataDisk{
		// 								},
		// 								ImageReference: &armcompute.ImageReference{
		// 									ID: to.Ptr("aaaaaaaa"),
		// 									CommunityGalleryImageID: to.Ptr("aaaaaaaaaaaaaaaaaaaaaaaaaaaa"),
		// 									ExactVersion: to.Ptr("aa"),
		// 									Offer: to.Ptr("WindowsServer"),
		// 									Publisher: to.Ptr("MicrosoftWindowsServer"),
		// 									SharedGalleryImageID: to.Ptr("aaaaaaaaaaaaaaaaaaaaaaaaa"),
		// 									SKU: to.Ptr("2012-R2-Datacenter"),
		// 									Version: to.Ptr("4.127.20170406"),
		// 								},
		// 								OSDisk: &armcompute.OSDisk{
		// 									Name: to.Ptr("test"),
		// 									Caching: to.Ptr(armcompute.CachingTypesNone),
		// 									CreateOption: to.Ptr(armcompute.DiskCreateOptionTypesFromImage),
		// 									DeleteOption: to.Ptr(armcompute.DiskDeleteOptionTypesDelete),
		// 									DiffDiskSettings: &armcompute.DiffDiskSettings{
		// 										Option: to.Ptr(armcompute.DiffDiskOptionsLocal),
		// 										Placement: to.Ptr(armcompute.DiffDiskPlacementCacheDisk),
		// 									},
		// 									DiskSizeGB: to.Ptr[int32](127),
		// 									EncryptionSettings: &armcompute.DiskEncryptionSettings{
		// 										DiskEncryptionKey: &armcompute.KeyVaultSecretReference{
		// 											SecretURL: to.Ptr("aaaaaaaaa"),
		// 											SourceVault: &armcompute.SubResource{
		// 												ID: to.Ptr("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/availabilitySets/{availabilitySetName}"),
		// 											},
		// 										},
		// 										Enabled: to.Ptr(true),
		// 										KeyEncryptionKey: &armcompute.KeyVaultKeyReference{
		// 											KeyURL: to.Ptr("aaaaaaaaaaaaa"),
		// 											SourceVault: &armcompute.SubResource{
		// 												ID: to.Ptr("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/availabilitySets/{availabilitySetName}"),
		// 											},
		// 										},
		// 									},
		// 									Image: &armcompute.VirtualHardDisk{
		// 										URI: to.Ptr("https://{storageAccountName}.blob.core.windows.net/{containerName}/{vhdName}.vhd"),
		// 									},
		// 									ManagedDisk: &armcompute.ManagedDiskParameters{
		// 										ID: to.Ptr("/subscriptions/{subscription-id}/resourceGroups/myResourceGroup/providers/Microsoft.Compute/disks/testingexcludedisk_OsDisk_1_74cdaedcea50483d9833c96adefa100f"),
		// 										DiskEncryptionSet: &armcompute.DiskEncryptionSetParameters{
		// 											ID: to.Ptr("aaaaaaaaaaaaaaaaaaaaaaaaaaaa"),
		// 										},
		// 										SecurityProfile: &armcompute.VMDiskSecurityProfile{
		// 											DiskEncryptionSet: &armcompute.DiskEncryptionSetParameters{
		// 												ID: to.Ptr("aaaaaaaaaaaaaaaaaaaaaaaaaaaa"),
		// 											},
		// 											SecurityEncryptionType: to.Ptr(armcompute.SecurityEncryptionTypesVMGuestStateOnly),
		// 										},
		// 										StorageAccountType: to.Ptr(armcompute.StorageAccountTypesStandardLRS),
		// 									},
		// 									OSType: to.Ptr(armcompute.OperatingSystemTypesWindows),
		// 									Vhd: &armcompute.VirtualHardDisk{
		// 										URI: to.Ptr("https://{storageAccountName}.blob.core.windows.net/{containerName}/{vhdName}.vhd"),
		// 									},
		// 									WriteAcceleratorEnabled: to.Ptr(true),
		// 								},
		// 							},
		// 							TimeCreated: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2022-01-14T16:43:41.685Z"); return t}()),
		// 							UserData: to.Ptr("aaaaaaaaaaaaaaaaaaa"),
		// 							VirtualMachineScaleSet: &armcompute.SubResource{
		// 								ID: to.Ptr("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/availabilitySets/{availabilitySetName}"),
		// 							},
		// 							VMID: to.Ptr("{vmId}"),
		// 						},
		// 						Resources: []*armcompute.VirtualMachineExtension{
		// 							{
		// 								Name: to.Ptr("aaaaaaaaaaaaa"),
		// 								Type: to.Ptr("aaaaaaaaaaaaaaaaaaaaaaaaaaa"),
		// 								ID: to.Ptr("aaaaaaaaaaaaaaaaaaaaaaaaaaaaa"),
		// 								Location: to.Ptr("aaaaaaaaaaaaaaaa"),
		// 								Tags: map[string]*string{
		// 									"key9428": to.Ptr("aaaaaaa"),
		// 								},
		// 								Properties: &armcompute.VirtualMachineExtensionProperties{
		// 									Type: to.Ptr("aaaaaaaa"),
		// 									AutoUpgradeMinorVersion: to.Ptr(true),
		// 									EnableAutomaticUpgrade: to.Ptr(true),
		// 									ForceUpdateTag: to.Ptr("aaaaaaaaaaaaaaa"),
		// 									InstanceView: &armcompute.VirtualMachineExtensionInstanceView{
		// 										Name: to.Ptr("aaaaaaaaaaaaaaaaaaaaaaaa"),
		// 										Type: to.Ptr("aaaaaaaaaaaaaaaaaaaaaaaa"),
		// 										Statuses: []*armcompute.InstanceViewStatus{
		// 											{
		// 												Code: to.Ptr("aaa"),
		// 												DisplayStatus: to.Ptr("aaaaaaaaaaaaaaaaaaaaaaaa"),
		// 												Level: to.Ptr(armcompute.StatusLevelTypesInfo),
		// 												Message: to.Ptr("aaaaaaaaaaaaaaaaaaa"),
		// 												Time: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2022-01-14T16:43:41.657Z"); return t}()),
		// 										}},
		// 										Substatuses: []*armcompute.InstanceViewStatus{
		// 											{
		// 												Code: to.Ptr("aaa"),
		// 												DisplayStatus: to.Ptr("aaaaaaaaaaaaaaaaaaaaaaaa"),
		// 												Level: to.Ptr(armcompute.StatusLevelTypesInfo),
		// 												Message: to.Ptr("aaaaaaaaaaaaaaaaaaa"),
		// 												Time: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2022-01-14T16:43:41.657Z"); return t}()),
		// 										}},
		// 										TypeHandlerVersion: to.Ptr("aaaaaa"),
		// 									},
		// 									ProtectedSettings: map[string]any{
		// 									},
		// 									ProtectedSettingsFromKeyVault: &armcompute.KeyVaultSecretReference{
		// 										SecretURL: to.Ptr("https://kvName.vault.azure.net/secrets/secretName/79b88b3a6f5440ffb2e73e44a0db712e"),
		// 										SourceVault: &armcompute.SubResource{
		// 											ID: to.Ptr("/subscriptions/a53f7094-a16c-47af-abe4-b05c05d0d79a/resourceGroups/myResourceGroup/providers/Microsoft.KeyVault/vaults/kvName"),
		// 										},
		// 									},
		// 									ProvisioningState: to.Ptr("Succeeded"),
		// 									Publisher: to.Ptr("aaaaaaaaaaaaaaaa"),
		// 									Settings: map[string]any{
		// 									},
		// 									SuppressFailures: to.Ptr(true),
		// 									TypeHandlerVersion: to.Ptr("aaaaaaaaaaaa"),
		// 								},
		// 						}},
		// 						Zones: []*string{
		// 							to.Ptr("aaaaaa")},
		// 					}},
		// 				}
	}
}
