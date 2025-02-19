package armcomputefleet_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/computefleet/armcomputefleet"
)

// Generated from example definition: D:/w/Azure/azure-sdk-for-go/sdk/resourcemanager/computefleet/armcomputefleet/TempTypeSpecFiles/AzureFleet.Management/examples/2024-11-01/Fleets_Get.json
func ExampleFleetsClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcomputefleet.NewClientFactory("1DC2F28C-A625-4B0E-9748-9885A3C9E9EB", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewFleetsClient().Get(ctx, "rgazurefleet", "testFleet", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armcomputefleet.FleetsClientGetResponse{
	// 	Fleet: &armcomputefleet.Fleet{
	// 		Properties: &armcomputefleet.FleetProperties{
	// 			ProvisioningState: to.Ptr(armcomputefleet.ProvisioningStateCreating),
	// 			SpotPriorityProfile: &armcomputefleet.SpotPriorityProfile{
	// 				Capacity: to.Ptr[int32](20),
	// 				MinCapacity: to.Ptr[int32](10),
	// 				MaxPricePerVM: to.Ptr[float32](0.00865),
	// 				EvictionPolicy: to.Ptr(armcomputefleet.EvictionPolicyDelete),
	// 				AllocationStrategy: to.Ptr(armcomputefleet.SpotAllocationStrategyPriceCapacityOptimized),
	// 				Maintain: to.Ptr(true),
	// 			},
	// 			RegularPriorityProfile: &armcomputefleet.RegularPriorityProfile{
	// 				Capacity: to.Ptr[int32](20),
	// 				MinCapacity: to.Ptr[int32](10),
	// 				AllocationStrategy: to.Ptr(armcomputefleet.RegularPriorityAllocationStrategyLowestPrice),
	// 			},
	// 			VMSizesProfile: []*armcomputefleet.VMSizeProfile{
	// 				{
	// 					Name: to.Ptr("Standard_d1_v2"),
	// 					Rank: to.Ptr[int32](19225),
	// 				},
	// 			},
	// 			ComputeProfile: &armcomputefleet.ComputeProfile{
	// 				BaseVirtualMachineProfile: &armcomputefleet.BaseVirtualMachineProfile{
	// 					OSProfile: &armcomputefleet.VirtualMachineScaleSetOSProfile{
	// 						ComputerNamePrefix: to.Ptr("o"),
	// 						AdminUsername: to.Ptr("nrgzqciiaaxjrqldbmjbqkyhntp"),
	// 						WindowsConfiguration: &armcomputefleet.WindowsConfiguration{
	// 							ProvisionVMAgent: to.Ptr(true),
	// 							EnableAutomaticUpdates: to.Ptr(true),
	// 							TimeZone: to.Ptr("hlyjiqcfksgrpjrct"),
	// 							AdditionalUnattendContent: []*armcomputefleet.AdditionalUnattendContent{
	// 								{
	// 									PassName: to.Ptr("OobeSystem"),
	// 									ComponentName: to.Ptr("Microsoft-Windows-Shell-Setup"),
	// 									SettingName: to.Ptr(armcomputefleet.SettingNamesAutoLogon),
	// 								},
	// 							},
	// 							PatchSettings: &armcomputefleet.PatchSettings{
	// 								PatchMode: to.Ptr(armcomputefleet.WindowsVMGuestPatchModeManual),
	// 								EnableHotpatching: to.Ptr(true),
	// 								AssessmentMode: to.Ptr(armcomputefleet.WindowsPatchAssessmentModeImageDefault),
	// 								AutomaticByPlatformSettings: &armcomputefleet.WindowsVMGuestPatchAutomaticByPlatformSettings{
	// 									RebootSetting: to.Ptr(armcomputefleet.WindowsVMGuestPatchAutomaticByPlatformRebootSettingUnknown),
	// 									BypassPlatformSafetyChecksOnUserSchedule: to.Ptr(true),
	// 								},
	// 							},
	// 							WinRM: &armcomputefleet.WinRMConfiguration{
	// 								Listeners: []*armcomputefleet.WinRMListener{
	// 									{
	// 										Protocol: to.Ptr(armcomputefleet.ProtocolTypesHTTP),
	// 										CertificateURL: to.Ptr("https://myVaultName.vault.azure.net/secrets/myCertName"),
	// 									},
	// 								},
	// 							},
	// 							EnableVMAgentPlatformUpdates: to.Ptr(true),
	// 						},
	// 						LinuxConfiguration: &armcomputefleet.LinuxConfiguration{
	// 							DisablePasswordAuthentication: to.Ptr(true),
	// 							SSH: &armcomputefleet.SSHConfiguration{
	// 								PublicKeys: []*armcomputefleet.SSHPublicKey{
	// 									{
	// 										Path: to.Ptr("kmqz"),
	// 										KeyData: to.Ptr("kivgsubusvpprwqaqpjcmhsv"),
	// 									},
	// 								},
	// 							},
	// 							ProvisionVMAgent: to.Ptr(true),
	// 							PatchSettings: &armcomputefleet.LinuxPatchSettings{
	// 								PatchMode: to.Ptr(armcomputefleet.LinuxVMGuestPatchModeImageDefault),
	// 								AssessmentMode: to.Ptr(armcomputefleet.LinuxPatchAssessmentModeImageDefault),
	// 								AutomaticByPlatformSettings: &armcomputefleet.LinuxVMGuestPatchAutomaticByPlatformSettings{
	// 									RebootSetting: to.Ptr(armcomputefleet.LinuxVMGuestPatchAutomaticByPlatformRebootSettingUnknown),
	// 									BypassPlatformSafetyChecksOnUserSchedule: to.Ptr(true),
	// 								},
	// 							},
	// 							EnableVMAgentPlatformUpdates: to.Ptr(true),
	// 						},
	// 						Secrets: []*armcomputefleet.VaultSecretGroup{
	// 							{
	// 								SourceVault: &armcomputefleet.SubResource{
	// 									ID: to.Ptr("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.KeyVault/vaults/{vaultName}"),
	// 								},
	// 								VaultCertificates: []*armcomputefleet.VaultCertificate{
	// 									{
	// 										CertificateURL: to.Ptr("https://myVaultName.vault.azure.net/secrets/myCertName"),
	// 										CertificateStore: to.Ptr("nlxrwavpzhueffxsshlun"),
	// 									},
	// 								},
	// 							},
	// 						},
	// 						AllowExtensionOperations: to.Ptr(true),
	// 						RequireGuestProvisionSignal: to.Ptr(true),
	// 					},
	// 					StorageProfile: &armcomputefleet.VirtualMachineScaleSetStorageProfile{
	// 						ImageReference: &armcomputefleet.ImageReference{
	// 							Publisher: to.Ptr("mqxgwbiyjzmxavhbkd"),
	// 							Offer: to.Ptr("isxgumkarlkomp"),
	// 							SKU: to.Ptr("eojmppqcrnpmxirtp"),
	// 							Version: to.Ptr("wvpcqefgtmqdgltiuz"),
	// 							ExactVersion: to.Ptr("zjbntmiskjexlr"),
	// 							SharedGalleryImageID: to.Ptr("kmkgihoxwlawuuhcinfirktdwkmx"),
	// 							CommunityGalleryImageID: to.Ptr("vlqe"),
	// 							ID: to.Ptr("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/galleries/{galleryName}/images/{imageName}/versions/{versionName}"),
	// 						},
	// 						OSDisk: &armcomputefleet.VirtualMachineScaleSetOSDisk{
	// 							Name: to.Ptr("wfttw"),
	// 							Caching: to.Ptr(armcomputefleet.CachingTypesNone),
	// 							WriteAcceleratorEnabled: to.Ptr(true),
	// 							CreateOption: to.Ptr(armcomputefleet.DiskCreateOptionTypesFromImage),
	// 							DiffDiskSettings: &armcomputefleet.DiffDiskSettings{
	// 								Option: to.Ptr(armcomputefleet.DiffDiskOptionsLocal),
	// 								Placement: to.Ptr(armcomputefleet.DiffDiskPlacementCacheDisk),
	// 							},
	// 							DiskSizeGB: to.Ptr[int32](14),
	// 							OSType: to.Ptr(armcomputefleet.OperatingSystemTypesWindows),
	// 							Image: &armcomputefleet.VirtualHardDisk{
	// 								URI: to.Ptr("https://myStorageAccountName.blob.core.windows.net/myContainerName/myVhdName.vhd"),
	// 							},
	// 							VhdContainers: []*string{
	// 								to.Ptr("tkzcwddtinkfpnfklatw"),
	// 							},
	// 							ManagedDisk: &armcomputefleet.VirtualMachineScaleSetManagedDiskParameters{
	// 								StorageAccountType: to.Ptr(armcomputefleet.StorageAccountTypesStandardLRS),
	// 								DiskEncryptionSet: &armcomputefleet.DiskEncryptionSetParameters{
	// 									ID: to.Ptr("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/diskEncryptionSets/{diskEncryptionSetName}"),
	// 								},
	// 								SecurityProfile: &armcomputefleet.VMDiskSecurityProfile{
	// 									SecurityEncryptionType: to.Ptr(armcomputefleet.SecurityEncryptionTypesVMGuestStateOnly),
	// 									DiskEncryptionSet: &armcomputefleet.DiskEncryptionSetParameters{
	// 										ID: to.Ptr("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/diskEncryptionSets/{diskEncryptionSetName}"),
	// 									},
	// 								},
	// 							},
	// 							DeleteOption: to.Ptr(armcomputefleet.DiskDeleteOptionTypesDelete),
	// 						},
	// 						DataDisks: []*armcomputefleet.VirtualMachineScaleSetDataDisk{
	// 							{
	// 								Name: to.Ptr("eogiykmdmeikswxmigjws"),
	// 								Lun: to.Ptr[int32](14),
	// 								Caching: to.Ptr(armcomputefleet.CachingTypesNone),
	// 								WriteAcceleratorEnabled: to.Ptr(true),
	// 								CreateOption: to.Ptr(armcomputefleet.DiskCreateOptionTypesFromImage),
	// 								DiskSizeGB: to.Ptr[int32](6),
	// 								ManagedDisk: &armcomputefleet.VirtualMachineScaleSetManagedDiskParameters{
	// 									StorageAccountType: to.Ptr(armcomputefleet.StorageAccountTypesStandardLRS),
	// 									DiskEncryptionSet: &armcomputefleet.DiskEncryptionSetParameters{
	// 										ID: to.Ptr("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/diskEncryptionSets/{diskEncryptionSetName}"),
	// 									},
	// 									SecurityProfile: &armcomputefleet.VMDiskSecurityProfile{
	// 										SecurityEncryptionType: to.Ptr(armcomputefleet.SecurityEncryptionTypesVMGuestStateOnly),
	// 										DiskEncryptionSet: &armcomputefleet.DiskEncryptionSetParameters{
	// 											ID: to.Ptr("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/diskEncryptionSets/{diskEncryptionSetName}"),
	// 										},
	// 									},
	// 								},
	// 								DiskIOPSReadWrite: to.Ptr[int64](27),
	// 								DiskMBpsReadWrite: to.Ptr[int64](2),
	// 								DeleteOption: to.Ptr(armcomputefleet.DiskDeleteOptionTypesDelete),
	// 							},
	// 						},
	// 					},
	// 					NetworkProfile: &armcomputefleet.VirtualMachineScaleSetNetworkProfile{
	// 						HealthProbe: &armcomputefleet.APIEntityReference{
	// 							ID: to.Ptr("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/loadBalancers/{loadBalancerName}/probes/{probeName}"),
	// 						},
	// 						NetworkInterfaceConfigurations: []*armcomputefleet.VirtualMachineScaleSetNetworkConfiguration{
	// 							{
	// 								Name: to.Ptr("i"),
	// 								Properties: &armcomputefleet.VirtualMachineScaleSetNetworkConfigurationProperties{
	// 									Primary: to.Ptr(true),
	// 									EnableAcceleratedNetworking: to.Ptr(true),
	// 									DisableTCPStateTracking: to.Ptr(true),
	// 									EnableFpga: to.Ptr(true),
	// 									NetworkSecurityGroup: &armcomputefleet.SubResource{
	// 										ID: to.Ptr("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/networkSecurityGroups/{networkSecurityGroupName}"),
	// 									},
	// 									DNSSettings: &armcomputefleet.VirtualMachineScaleSetNetworkConfigurationDNSSettings{
	// 										DNSServers: []*string{
	// 											to.Ptr("nxmmfolhclsesu"),
	// 										},
	// 									},
	// 									IPConfigurations: []*armcomputefleet.VirtualMachineScaleSetIPConfiguration{
	// 										{
	// 											Name: to.Ptr("oezqhkidfhyywlfzwuotilrpbqnjg"),
	// 											Properties: &armcomputefleet.VirtualMachineScaleSetIPConfigurationProperties{
	// 												Subnet: &armcomputefleet.APIEntityReference{
	// 													ID: to.Ptr("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/virtualNetworks/{virtualNetworkName}/subnets/{subnetName}"),
	// 												},
	// 												Primary: to.Ptr(true),
	// 												PublicIPAddressConfiguration: &armcomputefleet.VirtualMachineScaleSetPublicIPAddressConfiguration{
	// 													Name: to.Ptr("fvpqf"),
	// 													Properties: &armcomputefleet.VirtualMachineScaleSetPublicIPAddressConfigurationProperties{
	// 														IdleTimeoutInMinutes: to.Ptr[int32](9),
	// 														DNSSettings: &armcomputefleet.VirtualMachineScaleSetPublicIPAddressConfigurationDNSSettings{
	// 															DomainNameLabel: to.Ptr("ukrddzvmorpmfsczjwtbvp"),
	// 															DomainNameLabelScope: to.Ptr(armcomputefleet.DomainNameLabelScopeTypesTenantReuse),
	// 														},
	// 														IPTags: []*armcomputefleet.VirtualMachineScaleSetIPTag{
	// 															{
	// 																IPTagType: to.Ptr("sddgsoemnzgqizale"),
	// 																Tag: to.Ptr("wufmhrjsakbiaetyara"),
	// 															},
	// 														},
	// 														PublicIPPrefix: &armcomputefleet.SubResource{
	// 															ID: to.Ptr("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/publicIPPrefixes/{publicIPPrefixName}"),
	// 														},
	// 														PublicIPAddressVersion: to.Ptr(armcomputefleet.IPVersionIPv4),
	// 														DeleteOption: to.Ptr(armcomputefleet.DeleteOptionsDelete),
	// 													},
	// 													SKU: &armcomputefleet.PublicIPAddressSKU{
	// 														Name: to.Ptr(armcomputefleet.PublicIPAddressSKUNameBasic),
	// 														Tier: to.Ptr(armcomputefleet.PublicIPAddressSKUTierRegional),
	// 													},
	// 												},
	// 												PrivateIPAddressVersion: to.Ptr(armcomputefleet.IPVersionIPv4),
	// 												ApplicationGatewayBackendAddressPools: []*armcomputefleet.SubResource{
	// 													{
	// 														ID: to.Ptr("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/applicationGateways/{applicationGatewayName}/backendAddressPools/{backendAddressPoolName}"),
	// 													},
	// 												},
	// 												ApplicationSecurityGroups: []*armcomputefleet.SubResource{
	// 													{
	// 														ID: to.Ptr("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/applicationSecurityGroups/{applicationSecurityGroupName}"),
	// 													},
	// 												},
	// 												LoadBalancerBackendAddressPools: []*armcomputefleet.SubResource{
	// 													{
	// 														ID: to.Ptr("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/loadBalancers/{loadBalancerName}/backendAddressPools/{backendAddressPoolName}"),
	// 													},
	// 												},
	// 												LoadBalancerInboundNatPools: []*armcomputefleet.SubResource{
	// 													{
	// 														ID: to.Ptr("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/loadBalancers/{loadBalancerName}/inboundNatPools/{inboundNatPoolName}"),
	// 													},
	// 												},
	// 											},
	// 										},
	// 									},
	// 									EnableIPForwarding: to.Ptr(true),
	// 									DeleteOption: to.Ptr(armcomputefleet.DeleteOptionsDelete),
	// 									AuxiliaryMode: to.Ptr(armcomputefleet.NetworkInterfaceAuxiliaryModeNone),
	// 									AuxiliarySKU: to.Ptr(armcomputefleet.NetworkInterfaceAuxiliarySKUNone),
	// 								},
	// 							},
	// 						},
	// 						NetworkAPIVersion: to.Ptr(armcomputefleet.NetworkAPIVersionV20201101),
	// 					},
	// 					SecurityProfile: &armcomputefleet.SecurityProfile{
	// 						UefiSettings: &armcomputefleet.UefiSettings{
	// 							SecureBootEnabled: to.Ptr(true),
	// 							VTpmEnabled: to.Ptr(true),
	// 						},
	// 						EncryptionAtHost: to.Ptr(true),
	// 						SecurityType: to.Ptr(armcomputefleet.SecurityTypesTrustedLaunch),
	// 						EncryptionIdentity: &armcomputefleet.EncryptionIdentity{
	// 							UserAssignedIdentityResourceID: to.Ptr("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ManagedIdentity/userAssignedIdentities/{userAssignedIdentityName}"),
	// 						},
	// 						ProxyAgentSettings: &armcomputefleet.ProxyAgentSettings{
	// 							Enabled: to.Ptr(true),
	// 							Mode: to.Ptr(armcomputefleet.ModeAudit),
	// 							KeyIncarnationID: to.Ptr[int32](20),
	// 						},
	// 					},
	// 					DiagnosticsProfile: &armcomputefleet.DiagnosticsProfile{
	// 						BootDiagnostics: &armcomputefleet.BootDiagnostics{
	// 							Enabled: to.Ptr(true),
	// 							StorageURI: to.Ptr("http://myStorageAccountName.blob.core.windows.net"),
	// 						},
	// 					},
	// 					ExtensionProfile: &armcomputefleet.VirtualMachineScaleSetExtensionProfile{
	// 						Extensions: []*armcomputefleet.VirtualMachineScaleSetExtension{
	// 							{
	// 								Name: to.Ptr("bndxuxx"),
	// 								Type: to.Ptr("cmeam"),
	// 								Properties: &armcomputefleet.VirtualMachineScaleSetExtensionProperties{
	// 									ForceUpdateTag: to.Ptr("yhgxw"),
	// 									Publisher: to.Ptr("kpxtirxjfprhs"),
	// 									Type: to.Ptr("pgjilctjjwaa"),
	// 									TypeHandlerVersion: to.Ptr("zevivcoilxmbwlrihhhibq"),
	// 									AutoUpgradeMinorVersion: to.Ptr(true),
	// 									EnableAutomaticUpgrade: to.Ptr(true),
	// 									Settings: map[string]any{
	// 									},
	// 									ProvisioningState: to.Ptr("Succeeded"),
	// 									ProvisionAfterExtensions: []*string{
	// 										to.Ptr("nftzosroolbcwmpupujzqwqe"),
	// 									},
	// 									SuppressFailures: to.Ptr(true),
	// 									ProtectedSettingsFromKeyVault: &armcomputefleet.KeyVaultSecretReference{
	// 										SecretURL: to.Ptr("https://myVaultName.vault.azure.net/secrets/secret/mySecretName"),
	// 										SourceVault: &armcomputefleet.SubResource{
	// 											ID: to.Ptr("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.KeyVault/vaults/{vaultName}"),
	// 										},
	// 									},
	// 								},
	// 								ID: to.Ptr("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/virtualMachines/{vmName}/extensions/{extensionName}"),
	// 							},
	// 						},
	// 						ExtensionsTimeBudget: to.Ptr("mbhjahtdygwgyszdwjtvlvtgchdwil"),
	// 					},
	// 					LicenseType: to.Ptr("v"),
	// 					ScheduledEventsProfile: &armcomputefleet.ScheduledEventsProfile{
	// 						TerminateNotificationProfile: &armcomputefleet.TerminateNotificationProfile{
	// 							NotBeforeTimeout: to.Ptr("iljppmmw"),
	// 							Enable: to.Ptr(true),
	// 						},
	// 						OSImageNotificationProfile: &armcomputefleet.OSImageNotificationProfile{
	// 							NotBeforeTimeout: to.Ptr("olbpadmevekyczfokodtfprxti"),
	// 							Enable: to.Ptr(true),
	// 						},
	// 					},
	// 					UserData: to.Ptr("s"),
	// 					CapacityReservation: &armcomputefleet.CapacityReservationProfile{
	// 						CapacityReservationGroup: &armcomputefleet.SubResource{
	// 							ID: to.Ptr("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/capacityReservationGroups/{capacityReservationGroupName}"),
	// 						},
	// 					},
	// 					ApplicationProfile: &armcomputefleet.ApplicationProfile{
	// 						GalleryApplications: []*armcomputefleet.VMGalleryApplication{
	// 							{
	// 								Tags: to.Ptr("eyrqjbib"),
	// 								Order: to.Ptr[int32](5),
	// 								PackageReferenceID: to.Ptr("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/galleries/{galleryName}/applications/{applicationName}/versions/{versionName}"),
	// 								TreatFailureAsDeploymentFailure: to.Ptr(true),
	// 								EnableAutomaticUpgrade: to.Ptr(true),
	// 							},
	// 						},
	// 					},
	// 					HardwareProfile: &armcomputefleet.VirtualMachineScaleSetHardwareProfile{
	// 						VMSizeProperties: &armcomputefleet.VMSizeProperties{
	// 							VCPUsAvailable: to.Ptr[int32](16),
	// 							VCPUsPerCore: to.Ptr[int32](23),
	// 						},
	// 					},
	// 					ServiceArtifactReference: &armcomputefleet.ServiceArtifactReference{
	// 						ID: to.Ptr("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/galleries/{galleryName}/serviceArtifacts/{serviceArtifactsName}/vmArtifactsProfiles/{vmArtifactsProfileName}"),
	// 					},
	// 					SecurityPostureReference: &armcomputefleet.SecurityPostureReference{
	// 						ID: to.Ptr("mubredelfbshboaxrsxiajihahaa"),
	// 						ExcludeExtensions: []*string{
	// 							to.Ptr("{securityPostureVMExtensionName}"),
	// 						},
	// 						IsOverridable: to.Ptr(true),
	// 					},
	// 					TimeCreated: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2024-04-29T21:51:44.043Z"); return t}()),
	// 				},
	// 				ComputeAPIVersion: to.Ptr("2023-07-01"),
	// 				PlatformFaultDomainCount: to.Ptr[int32](1),
	// 			},
	// 			TimeCreated: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2024-07-10T00:16:34.8590555+00:00"); return t}()),
	// 			UniqueID: to.Ptr("a2f7fabd-bbc2-4a20-afe1-49fdb3885a28"),
	// 		},
	// 		Zones: []*string{
	// 			to.Ptr("zone1"),
	// 			to.Ptr("zone2"),
	// 		},
	// 		Identity: &armcomputefleet.ManagedServiceIdentity{
	// 			PrincipalID: to.Ptr("4d508e5b-374b-4382-9a1c-01fb8b6cb37c"),
	// 			TenantID: to.Ptr("5d508e5b-374b-4382-9a1c-01fb8b6cb37c"),
	// 			Type: to.Ptr(armcomputefleet.ManagedServiceIdentityTypeUserAssigned),
	// 			UserAssignedIdentities: map[string]*armcomputefleet.UserAssignedIdentity{
	// 				"key9851": &armcomputefleet.UserAssignedIdentity{
	// 					PrincipalID: to.Ptr("6d508e5b-374b-4382-9a1c-01fb8b6cb37c"),
	// 					ClientID: to.Ptr("7d508e5b-374b-4382-9a1c-01fb8b6cb37c"),
	// 				},
	// 			},
	// 		},
	// 		Tags: map[string]*string{
	// 			"key3518": to.Ptr("luvrnuvsgdpbuofdskkcoqhfh"),
	// 		},
	// 		Plan: &armcomputefleet.Plan{
	// 			Name: to.Ptr("jwgrcrnrtfoxn"),
	// 			Publisher: to.Ptr("iozjbiqqckqm"),
	// 			Product: to.Ptr("cgopbyvdyqikahwyxfpzwaqk"),
	// 			PromotionCode: to.Ptr("naglezezplcaruqogtxnuizslqnnbr"),
	// 			Version: to.Ptr("wa"),
	// 		},
	// 		Location: to.Ptr("westus"),
	// 		ID: to.Ptr("/subscriptions/7B0CD4DB-3381-4013-9B31-FB6E6FD0FF1C/resourceGroups/rgazurefleet/providers/Microsoft.AzureFleet/fleets/testFleet"),
	// 		Name: to.Ptr("testFleet"),
	// 		Type: to.Ptr("Microsoft.AzureFleet/fleets"),
	// 		SystemData: &armcomputefleet.SystemData{
	// 			CreatedBy: to.Ptr("rowegentrpoajsv"),
	// 			CreatedByType: to.Ptr(armcomputefleet.CreatedByTypeUser),
	// 			CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2024-04-29T21:51:44.043Z"); return t}()),
	// 			LastModifiedBy: to.Ptr("edwuayhhaoepxzisfaqjhmrxjq"),
	// 			LastModifiedByType: to.Ptr(armcomputefleet.CreatedByTypeUser),
	// 			LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2024-04-29T21:51:44.043Z"); return t}()),
	// 		},
	// 	},
	// }
}
