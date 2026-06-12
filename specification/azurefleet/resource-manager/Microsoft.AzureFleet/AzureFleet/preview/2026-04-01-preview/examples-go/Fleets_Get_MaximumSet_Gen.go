package armcomputefleet_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/computefleet/armcomputefleet/v2"
)

// Generated from example definition: 2026-04-01-preview/Fleets_Get_MaximumSet_Gen.json
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
	res, err := clientFactory.NewFleetsClient().Get(ctx, "rgazurefleet", "myFleet", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armcomputefleet.FleetsClientGetResponse{
	// 	Fleet: armcomputefleet.Fleet{
	// 		Properties: &armcomputefleet.FleetProperties{
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
	// 					Name: to.Ptr("Standard_D1_v2"),
	// 				},
	// 				{
	// 					Name: to.Ptr("Standard_D2_v2"),
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
	// 										Protocol: to.Ptr(armcomputefleet.ProtocolTypesHTTPS),
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
	// 							SharedGalleryImageID: to.Ptr("kmkgihoxwlawuuhcinfirktdwkmx"),
	// 							CommunityGalleryImageID: to.Ptr("vlqe"),
	// 							ID: to.Ptr("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/galleries/{galleryName}/images/{imageName}/versions/{versionName}"),
	// 							ExactVersion: to.Ptr("zjbntmiskjexlr"),
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
	// 						DiskControllerType: to.Ptr(armcomputefleet.DiskControllerTypes("uzb")),
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
	// 								Properties: &armcomputefleet.VirtualMachineScaleSetExtensionProperties{
	// 									ForceUpdateTag: to.Ptr("yhgxw"),
	// 									Publisher: to.Ptr("kpxtirxjfprhs"),
	// 									Type: to.Ptr("pgjilctjjwaa"),
	// 									TypeHandlerVersion: to.Ptr("zevivcoilxmbwlrihhhibq"),
	// 									AutoUpgradeMinorVersion: to.Ptr(true),
	// 									EnableAutomaticUpgrade: to.Ptr(true),
	// 									Settings: map[string]any{
	// 									},
	// 									ProvisionAfterExtensions: []*string{
	// 										to.Ptr("nftzosroolbcwmpupujzqwqe"),
	// 									},
	// 									SuppressFailures: to.Ptr(true),
	// 									ProtectedSettingsFromKeyVault: &armcomputefleet.KeyVaultSecretReference{
	// 										SecretURL: to.Ptr("https://myvaultName.vault.azure.net/secrets/secret/mySecretName"),
	// 										SourceVault: &armcomputefleet.SubResource{
	// 											ID: to.Ptr("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.KeyVault/vaults/{vaultName}"),
	// 										},
	// 									},
	// 									ProvisioningState: to.Ptr("Succeeded"),
	// 								},
	// 								ID: to.Ptr("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/virtualMachines/{vmName}/extensions/{extensionName}"),
	// 								Type: to.Ptr("cmeam"),
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
	// 								ConfigurationReference: to.Ptr("ulztmiavpojpbpbddgnuuiimxcpau"),
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
	// 						ID: to.Ptr("/CommunityGalleries/{communityGalleryName}/securityPostures/{securityPostureName}/versions/{major.minor.patch}|{major.*}|latest"),
	// 						ExcludeExtensions: []*string{
	// 							to.Ptr("{securityPostureVMExtensionName}"),
	// 						},
	// 						IsOverridable: to.Ptr(true),
	// 					},
	// 					TimeCreated: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2024-04-29T21:51:44.043Z"); return t}()),
	// 				},
	// 				ComputeAPIVersion: to.Ptr("2023-07-01"),
	// 				PlatformFaultDomainCount: to.Ptr[int32](1),
	// 				AdditionalVirtualMachineCapabilities: &armcomputefleet.AdditionalCapabilities{
	// 					UltraSSDEnabled: to.Ptr(true),
	// 					HibernationEnabled: to.Ptr(true),
	// 				},
	// 			},
	// 			Mode: to.Ptr(armcomputefleet.FleetModeLaunch),
	// 			VMNamePrefix: to.Ptr("test-vm"),
	// 			CapacityType: to.Ptr(armcomputefleet.CapacityTypeVM),
	// 			ZoneAllocationPolicy: &armcomputefleet.ZoneAllocationPolicy{
	// 				DistributionStrategy: to.Ptr(armcomputefleet.ZoneDistributionStrategyBestEffortSingleZone),
	// 			},
	// 			ProvisioningState: to.Ptr(armcomputefleet.ProvisioningStateCreating),
	// 			VMAttributes: &armcomputefleet.VMAttributes{
	// 				VCPUCount: &armcomputefleet.VMAttributeMinMaxInteger{
	// 					Min: to.Ptr[int32](2),
	// 					Max: to.Ptr[int32](4),
	// 				},
	// 				MemoryInGiB: &armcomputefleet.VMAttributeMinMaxDouble{
	// 					Min: to.Ptr[float64](2),
	// 					Max: to.Ptr[float64](4),
	// 				},
	// 				MemoryInGiBPerVCpu: &armcomputefleet.VMAttributeMinMaxDouble{
	// 					Min: to.Ptr[float64](2),
	// 					Max: to.Ptr[float64](4),
	// 				},
	// 				LocalStorageSupport: to.Ptr(armcomputefleet.VMAttributeSupportExcluded),
	// 				LocalStorageInGiB: &armcomputefleet.VMAttributeMinMaxDouble{
	// 					Min: to.Ptr[float64](2),
	// 					Max: to.Ptr[float64](4),
	// 				},
	// 				LocalStorageDiskTypes: []*armcomputefleet.LocalStorageDiskType{
	// 					to.Ptr(armcomputefleet.LocalStorageDiskTypeHDD),
	// 				},
	// 				DataDiskCount: &armcomputefleet.VMAttributeMinMaxInteger{
	// 					Min: to.Ptr[int32](2),
	// 					Max: to.Ptr[int32](4),
	// 				},
	// 				NetworkInterfaceCount: &armcomputefleet.VMAttributeMinMaxInteger{
	// 					Min: to.Ptr[int32](2),
	// 					Max: to.Ptr[int32](4),
	// 				},
	// 				NetworkBandwidthInMbps: &armcomputefleet.VMAttributeMinMaxDouble{
	// 					Min: to.Ptr[float64](2),
	// 					Max: to.Ptr[float64](4),
	// 				},
	// 				RdmaSupport: to.Ptr(armcomputefleet.VMAttributeSupportExcluded),
	// 				RdmaNetworkInterfaceCount: &armcomputefleet.VMAttributeMinMaxInteger{
	// 					Min: to.Ptr[int32](2),
	// 					Max: to.Ptr[int32](4),
	// 				},
	// 				AcceleratorSupport: to.Ptr(armcomputefleet.VMAttributeSupportExcluded),
	// 				AcceleratorManufacturers: []*armcomputefleet.AcceleratorManufacturer{
	// 					to.Ptr(armcomputefleet.AcceleratorManufacturerAMD),
	// 				},
	// 				AcceleratorTypes: []*armcomputefleet.AcceleratorType{
	// 					to.Ptr(armcomputefleet.AcceleratorTypeGPU),
	// 				},
	// 				AcceleratorCount: &armcomputefleet.VMAttributeMinMaxInteger{
	// 					Min: to.Ptr[int32](2),
	// 					Max: to.Ptr[int32](4),
	// 				},
	// 				VMCategories: []*armcomputefleet.VMCategory{
	// 					to.Ptr(armcomputefleet.VMCategoryGeneralPurpose),
	// 				},
	// 				ArchitectureTypes: []*armcomputefleet.ArchitectureType{
	// 					to.Ptr(armcomputefleet.ArchitectureTypeARM64),
	// 				},
	// 				CPUManufacturers: []*armcomputefleet.CPUManufacturer{
	// 					to.Ptr(armcomputefleet.CPUManufacturerIntel),
	// 				},
	// 				BurstableSupport: to.Ptr(armcomputefleet.VMAttributeSupportExcluded),
	// 				ExcludedVMSizes: []*string{
	// 					to.Ptr("Standard_A1"),
	// 				},
	// 			},
	// 			AdditionalLocationsProfile: &armcomputefleet.AdditionalLocationsProfile{
	// 				LocationProfiles: []*armcomputefleet.LocationProfile{
	// 					{
	// 						Location: to.Ptr("v"),
	// 						VirtualMachineProfileOverride: &armcomputefleet.BaseVirtualMachineProfile{
	// 							OSProfile: &armcomputefleet.VirtualMachineScaleSetOSProfile{
	// 								ComputerNamePrefix: to.Ptr("tec"),
	// 								AdminUsername: to.Ptr("xdgnnqymtamdyqxy"),
	// 								WindowsConfiguration: &armcomputefleet.WindowsConfiguration{
	// 									ProvisionVMAgent: to.Ptr(true),
	// 									EnableAutomaticUpdates: to.Ptr(true),
	// 									TimeZone: to.Ptr("ktf"),
	// 									AdditionalUnattendContent: []*armcomputefleet.AdditionalUnattendContent{
	// 										{
	// 											PassName: to.Ptr("OobeSystem"),
	// 											ComponentName: to.Ptr("Microsoft-Windows-Shell-Setup"),
	// 											SettingName: to.Ptr(armcomputefleet.SettingNamesAutoLogon),
	// 										},
	// 									},
	// 									PatchSettings: &armcomputefleet.PatchSettings{
	// 										PatchMode: to.Ptr(armcomputefleet.WindowsVMGuestPatchModeManual),
	// 										EnableHotpatching: to.Ptr(true),
	// 										AssessmentMode: to.Ptr(armcomputefleet.WindowsPatchAssessmentModeImageDefault),
	// 										AutomaticByPlatformSettings: &armcomputefleet.WindowsVMGuestPatchAutomaticByPlatformSettings{
	// 											RebootSetting: to.Ptr(armcomputefleet.WindowsVMGuestPatchAutomaticByPlatformRebootSettingUnknown),
	// 											BypassPlatformSafetyChecksOnUserSchedule: to.Ptr(true),
	// 										},
	// 									},
	// 									WinRM: &armcomputefleet.WinRMConfiguration{
	// 										Listeners: []*armcomputefleet.WinRMListener{
	// 											{
	// 												Protocol: to.Ptr(armcomputefleet.ProtocolTypesHTTP),
	// 												CertificateURL: to.Ptr("https://microsoft.com/apzd"),
	// 											},
	// 										},
	// 									},
	// 									EnableVMAgentPlatformUpdates: to.Ptr(true),
	// 								},
	// 								LinuxConfiguration: &armcomputefleet.LinuxConfiguration{
	// 									DisablePasswordAuthentication: to.Ptr(true),
	// 									SSH: &armcomputefleet.SSHConfiguration{
	// 										PublicKeys: []*armcomputefleet.SSHPublicKey{
	// 											{
	// 												Path: to.Ptr("ebeglujkldnntlpmazrg"),
	// 												KeyData: to.Ptr("vmgnwtwjcodavmu"),
	// 											},
	// 										},
	// 									},
	// 									ProvisionVMAgent: to.Ptr(true),
	// 									PatchSettings: &armcomputefleet.LinuxPatchSettings{
	// 										PatchMode: to.Ptr(armcomputefleet.LinuxVMGuestPatchModeImageDefault),
	// 										AssessmentMode: to.Ptr(armcomputefleet.LinuxPatchAssessmentModeImageDefault),
	// 										AutomaticByPlatformSettings: &armcomputefleet.LinuxVMGuestPatchAutomaticByPlatformSettings{
	// 											RebootSetting: to.Ptr(armcomputefleet.LinuxVMGuestPatchAutomaticByPlatformRebootSettingUnknown),
	// 											BypassPlatformSafetyChecksOnUserSchedule: to.Ptr(true),
	// 										},
	// 									},
	// 									EnableVMAgentPlatformUpdates: to.Ptr(true),
	// 								},
	// 								Secrets: []*armcomputefleet.VaultSecretGroup{
	// 									{
	// 										SourceVault: &armcomputefleet.SubResource{
	// 											ID: to.Ptr("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.KeyVault/vaults/{vaultName}"),
	// 										},
	// 										VaultCertificates: []*armcomputefleet.VaultCertificate{
	// 											{
	// 												CertificateURL: to.Ptr("https://microsoft.com/a"),
	// 												CertificateStore: to.Ptr("yycyfwpymjtwzza"),
	// 											},
	// 										},
	// 									},
	// 								},
	// 								AllowExtensionOperations: to.Ptr(true),
	// 								RequireGuestProvisionSignal: to.Ptr(true),
	// 							},
	// 							StorageProfile: &armcomputefleet.VirtualMachineScaleSetStorageProfile{
	// 								ImageReference: &armcomputefleet.ImageReference{
	// 									Publisher: to.Ptr("mqxgwbiyjzmxavhbkd"),
	// 									Offer: to.Ptr("isxgumkarlkomp"),
	// 									SKU: to.Ptr("eojmppqcrnpmxirtp"),
	// 									Version: to.Ptr("wvpcqefgtmqdgltiuz"),
	// 									ExactVersion: to.Ptr("zjbntmiskjexlr"),
	// 									SharedGalleryImageID: to.Ptr("kmkgihoxwlawuuhcinfirktdwkmx"),
	// 									CommunityGalleryImageID: to.Ptr("vlqe"),
	// 									ID: to.Ptr("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/galleries/{galleryName}/images/{imageName}/versions/{versionName}"),
	// 								},
	// 								OSDisk: &armcomputefleet.VirtualMachineScaleSetOSDisk{
	// 									Name: to.Ptr("dt"),
	// 									Caching: to.Ptr(armcomputefleet.CachingTypesNone),
	// 									WriteAcceleratorEnabled: to.Ptr(true),
	// 									CreateOption: to.Ptr(armcomputefleet.DiskCreateOptionTypesFromImage),
	// 									DiffDiskSettings: &armcomputefleet.DiffDiskSettings{
	// 										Option: to.Ptr(armcomputefleet.DiffDiskOptionsLocal),
	// 										Placement: to.Ptr(armcomputefleet.DiffDiskPlacementCacheDisk),
	// 									},
	// 									DiskSizeGB: to.Ptr[int32](9),
	// 									OSType: to.Ptr(armcomputefleet.OperatingSystemTypesWindows),
	// 									Image: &armcomputefleet.VirtualHardDisk{
	// 										URI: to.Ptr("https://myStorageAccountName.blob.core.windows.net/myContainerName/myVhdName.vhd"),
	// 									},
	// 									VhdContainers: []*string{
	// 										to.Ptr("kdagj"),
	// 									},
	// 									ManagedDisk: &armcomputefleet.VirtualMachineScaleSetManagedDiskParameters{
	// 										StorageAccountType: to.Ptr(armcomputefleet.StorageAccountTypesStandardLRS),
	// 										DiskEncryptionSet: &armcomputefleet.DiskEncryptionSetParameters{
	// 											ID: to.Ptr("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/diskEncryptionSets/{diskEncryptionSetName}"),
	// 										},
	// 										SecurityProfile: &armcomputefleet.VMDiskSecurityProfile{
	// 											SecurityEncryptionType: to.Ptr(armcomputefleet.SecurityEncryptionTypesVMGuestStateOnly),
	// 											DiskEncryptionSet: &armcomputefleet.DiskEncryptionSetParameters{
	// 												ID: to.Ptr("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/diskEncryptionSets/{diskEncryptionSetName}"),
	// 											},
	// 										},
	// 									},
	// 									DeleteOption: to.Ptr(armcomputefleet.DiskDeleteOptionTypesDelete),
	// 								},
	// 								DataDisks: []*armcomputefleet.VirtualMachineScaleSetDataDisk{
	// 									{
	// 										Name: to.Ptr("mhljivkyryuomrapmmxx"),
	// 										Lun: to.Ptr[int32](6),
	// 										Caching: to.Ptr(armcomputefleet.CachingTypesNone),
	// 										WriteAcceleratorEnabled: to.Ptr(true),
	// 										CreateOption: to.Ptr(armcomputefleet.DiskCreateOptionTypesFromImage),
	// 										DiskSizeGB: to.Ptr[int32](9),
	// 										ManagedDisk: &armcomputefleet.VirtualMachineScaleSetManagedDiskParameters{
	// 											StorageAccountType: to.Ptr(armcomputefleet.StorageAccountTypesStandardLRS),
	// 											DiskEncryptionSet: &armcomputefleet.DiskEncryptionSetParameters{
	// 												ID: to.Ptr("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/diskEncryptionSets/{diskEncryptionSetName}"),
	// 											},
	// 											SecurityProfile: &armcomputefleet.VMDiskSecurityProfile{
	// 												SecurityEncryptionType: to.Ptr(armcomputefleet.SecurityEncryptionTypesVMGuestStateOnly),
	// 												DiskEncryptionSet: &armcomputefleet.DiskEncryptionSetParameters{
	// 													ID: to.Ptr("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/diskEncryptionSets/{diskEncryptionSetName}"),
	// 												},
	// 											},
	// 										},
	// 										DiskIOPSReadWrite: to.Ptr[int64](24),
	// 										DiskMBpsReadWrite: to.Ptr[int64](4),
	// 										DeleteOption: to.Ptr(armcomputefleet.DiskDeleteOptionTypesDelete),
	// 									},
	// 								},
	// 								DiskControllerType: to.Ptr(armcomputefleet.DiskControllerTypesSCSI),
	// 							},
	// 							NetworkProfile: &armcomputefleet.VirtualMachineScaleSetNetworkProfile{
	// 								HealthProbe: &armcomputefleet.APIEntityReference{
	// 									ID: to.Ptr("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/loadBalancers/{loadBalancerName}/probes/{probeName}"),
	// 								},
	// 								NetworkInterfaceConfigurations: []*armcomputefleet.VirtualMachineScaleSetNetworkConfiguration{
	// 									{
	// 										Name: to.Ptr("gpunpcdsdphgspvgwwbnk"),
	// 										Properties: &armcomputefleet.VirtualMachineScaleSetNetworkConfigurationProperties{
	// 											Primary: to.Ptr(true),
	// 											EnableAcceleratedNetworking: to.Ptr(true),
	// 											DisableTCPStateTracking: to.Ptr(true),
	// 											EnableFpga: to.Ptr(true),
	// 											NetworkSecurityGroup: &armcomputefleet.SubResource{
	// 												ID: to.Ptr("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/networkSecurityGroups/{networkSecurityGroupName}"),
	// 											},
	// 											DNSSettings: &armcomputefleet.VirtualMachineScaleSetNetworkConfigurationDNSSettings{
	// 												DNSServers: []*string{
	// 													to.Ptr("sjpmlu"),
	// 												},
	// 											},
	// 											IPConfigurations: []*armcomputefleet.VirtualMachineScaleSetIPConfiguration{
	// 												{
	// 													Name: to.Ptr("fweiphgkyhbcsbfjmxzczkpg"),
	// 													Properties: &armcomputefleet.VirtualMachineScaleSetIPConfigurationProperties{
	// 														Subnet: &armcomputefleet.APIEntityReference{
	// 															ID: to.Ptr("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/virtualNetworks/{virtualNetworkName}/subnets/{subnetName}"),
	// 														},
	// 														Primary: to.Ptr(true),
	// 														PublicIPAddressConfiguration: &armcomputefleet.VirtualMachineScaleSetPublicIPAddressConfiguration{
	// 															Name: to.Ptr("dvnoamqjyshquvtmf"),
	// 															Properties: &armcomputefleet.VirtualMachineScaleSetPublicIPAddressConfigurationProperties{
	// 																IdleTimeoutInMinutes: to.Ptr[int32](1),
	// 																DNSSettings: &armcomputefleet.VirtualMachineScaleSetPublicIPAddressConfigurationDNSSettings{
	// 																	DomainNameLabel: to.Ptr("ayofnb"),
	// 																	DomainNameLabelScope: to.Ptr(armcomputefleet.DomainNameLabelScopeTypesTenantReuse),
	// 																},
	// 																IPTags: []*armcomputefleet.VirtualMachineScaleSetIPTag{
	// 																	{
	// 																		IPTagType: to.Ptr("zqpznczmc"),
	// 																		Tag: to.Ptr("ugnfzikniqjisffrbvryavenhmtd"),
	// 																	},
	// 																},
	// 																PublicIPPrefix: &armcomputefleet.SubResource{
	// 																	ID: to.Ptr("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/publicIPPrefixes/{publicIPPrefixName}"),
	// 																},
	// 																PublicIPAddressVersion: to.Ptr(armcomputefleet.IPVersionIPv4),
	// 																DeleteOption: to.Ptr(armcomputefleet.DeleteOptionsDelete),
	// 															},
	// 															SKU: &armcomputefleet.PublicIPAddressSKU{
	// 																Name: to.Ptr(armcomputefleet.PublicIPAddressSKUNameBasic),
	// 																Tier: to.Ptr(armcomputefleet.PublicIPAddressSKUTierRegional),
	// 															},
	// 														},
	// 														PrivateIPAddressVersion: to.Ptr(armcomputefleet.IPVersionIPv4),
	// 														ApplicationGatewayBackendAddressPools: []*armcomputefleet.SubResource{
	// 															{
	// 																ID: to.Ptr("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/applicationGateways/{applicationGatewayName}/backendAddressPools/{backendAddressPoolName}"),
	// 															},
	// 														},
	// 														ApplicationSecurityGroups: []*armcomputefleet.SubResource{
	// 															{
	// 																ID: to.Ptr("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/applicationSecurityGroups/{applicationSecurityGroupName}"),
	// 															},
	// 														},
	// 														LoadBalancerBackendAddressPools: []*armcomputefleet.SubResource{
	// 															{
	// 																ID: to.Ptr("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/loadBalancers/{loadBalancerName}/backendAddressPools/{backendAddressPoolName}"),
	// 															},
	// 														},
	// 														LoadBalancerInboundNatPools: []*armcomputefleet.SubResource{
	// 															{
	// 																ID: to.Ptr("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/loadBalancers/{loadBalancerName}/inboundNatPools/{inboundNatPoolName}"),
	// 															},
	// 														},
	// 													},
	// 												},
	// 											},
	// 											EnableIPForwarding: to.Ptr(true),
	// 											DeleteOption: to.Ptr(armcomputefleet.DeleteOptionsDelete),
	// 											AuxiliaryMode: to.Ptr(armcomputefleet.NetworkInterfaceAuxiliaryModeNone),
	// 											AuxiliarySKU: to.Ptr(armcomputefleet.NetworkInterfaceAuxiliarySKUNone),
	// 										},
	// 									},
	// 								},
	// 								NetworkAPIVersion: to.Ptr(armcomputefleet.NetworkAPIVersionV20201101),
	// 							},
	// 							SecurityProfile: &armcomputefleet.SecurityProfile{
	// 								UefiSettings: &armcomputefleet.UefiSettings{
	// 									SecureBootEnabled: to.Ptr(true),
	// 									VTpmEnabled: to.Ptr(true),
	// 								},
	// 								EncryptionAtHost: to.Ptr(true),
	// 								SecurityType: to.Ptr(armcomputefleet.SecurityTypesTrustedLaunch),
	// 								EncryptionIdentity: &armcomputefleet.EncryptionIdentity{
	// 									UserAssignedIdentityResourceID: to.Ptr("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ManagedIdentity/userAssignedIdentities/{userAssignedIdentityName}"),
	// 								},
	// 								ProxyAgentSettings: &armcomputefleet.ProxyAgentSettings{
	// 									Enabled: to.Ptr(true),
	// 									Mode: to.Ptr(armcomputefleet.ModeAudit),
	// 									KeyIncarnationID: to.Ptr[int32](6),
	// 								},
	// 							},
	// 							DiagnosticsProfile: &armcomputefleet.DiagnosticsProfile{
	// 								BootDiagnostics: &armcomputefleet.BootDiagnostics{
	// 									Enabled: to.Ptr(true),
	// 									StorageURI: to.Ptr("https://microsoft.com/a"),
	// 								},
	// 							},
	// 							ExtensionProfile: &armcomputefleet.VirtualMachineScaleSetExtensionProfile{
	// 								Extensions: []*armcomputefleet.VirtualMachineScaleSetExtension{
	// 									{
	// 										ID: to.Ptr("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/virtualMachines/{vmName}/extensions/{extensionName}"),
	// 										Name: to.Ptr("oredyuufsd"),
	// 										Type: to.Ptr("qbmhvtcrufkhovrweggioqlcld"),
	// 										Properties: &armcomputefleet.VirtualMachineScaleSetExtensionProperties{
	// 											ForceUpdateTag: to.Ptr("muglieujh"),
	// 											Publisher: to.Ptr("ccbiyfuveemaaopgxbjpm"),
	// 											Type: to.Ptr("yorumzkbfpxnrdwgczwwaeaxmda"),
	// 											TypeHandlerVersion: to.Ptr("nlnqbmgzwubbc"),
	// 											AutoUpgradeMinorVersion: to.Ptr(true),
	// 											EnableAutomaticUpgrade: to.Ptr(true),
	// 											Settings: map[string]any{
	// 											},
	// 											ProvisioningState: to.Ptr("lcbvhjidublayqhmuzxzdvzgzslsys"),
	// 											ProvisionAfterExtensions: []*string{
	// 												to.Ptr("xuefrutmgzsxrpjjayvy"),
	// 											},
	// 											SuppressFailures: to.Ptr(true),
	// 											ProtectedSettingsFromKeyVault: &armcomputefleet.KeyVaultSecretReference{
	// 												SecretURL: to.Ptr("https://microsoft.com/a"),
	// 												SourceVault: &armcomputefleet.SubResource{
	// 													ID: to.Ptr("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.KeyVault/vaults/{vaultName}"),
	// 												},
	// 											},
	// 										},
	// 									},
	// 								},
	// 								ExtensionsTimeBudget: to.Ptr("trluxrynunvdnukztblhbnsubr"),
	// 							},
	// 							LicenseType: to.Ptr("ginsqshzwimjteiyfxhnjxfrcaat"),
	// 							ScheduledEventsProfile: &armcomputefleet.ScheduledEventsProfile{
	// 								TerminateNotificationProfile: &armcomputefleet.TerminateNotificationProfile{
	// 									NotBeforeTimeout: to.Ptr("plbazenobaeueixatewbey"),
	// 									Enable: to.Ptr(true),
	// 								},
	// 								OSImageNotificationProfile: &armcomputefleet.OSImageNotificationProfile{
	// 									NotBeforeTimeout: to.Ptr("ednjvcedpjmczw"),
	// 									Enable: to.Ptr(true),
	// 								},
	// 							},
	// 							UserData: to.Ptr("zekdr"),
	// 							CapacityReservation: &armcomputefleet.CapacityReservationProfile{
	// 								CapacityReservationGroup: &armcomputefleet.SubResource{
	// 									ID: to.Ptr("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/capacityReservationGroups/{capacityReservationGroupName}"),
	// 								},
	// 							},
	// 							ApplicationProfile: &armcomputefleet.ApplicationProfile{
	// 								GalleryApplications: []*armcomputefleet.VMGalleryApplication{
	// 									{
	// 										Tags: to.Ptr("eomzidad"),
	// 										Order: to.Ptr[int32](22),
	// 										PackageReferenceID: to.Ptr("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/galleries/{galleryName}/applications/{applicationName}/versions/{versionName}"),
	// 										ConfigurationReference: to.Ptr("zdqfcpvt"),
	// 										TreatFailureAsDeploymentFailure: to.Ptr(true),
	// 										EnableAutomaticUpgrade: to.Ptr(true),
	// 									},
	// 								},
	// 							},
	// 							HardwareProfile: &armcomputefleet.VirtualMachineScaleSetHardwareProfile{
	// 								VMSizeProperties: &armcomputefleet.VMSizeProperties{
	// 									VCPUsAvailable: to.Ptr[int32](8),
	// 									VCPUsPerCore: to.Ptr[int32](17),
	// 								},
	// 							},
	// 							ServiceArtifactReference: &armcomputefleet.ServiceArtifactReference{
	// 								ID: to.Ptr("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/galleries/{galleryName}/serviceArtifacts/{serviceArtifactsName}/vmArtifactsProfiles/{vmArtifactsProfileName}"),
	// 							},
	// 							SecurityPostureReference: &armcomputefleet.SecurityPostureReference{
	// 								ID: to.Ptr("/CommunityGalleries/{communityGalleryName}/securityPostures/{securityPostureName}/versions/{major.minor.patch}|{major.*}|latest"),
	// 								ExcludeExtensions: []*string{
	// 									to.Ptr("ragwgzswxzzz"),
	// 								},
	// 								IsOverridable: to.Ptr(true),
	// 							},
	// 							TimeCreated: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2025-07-29T18:51:38.722Z"); return t}()),
	// 						},
	// 					},
	// 				},
	// 			},
	// 			TimeCreated: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2024-07-10T00:16:34.8590555+00:00"); return t}()),
	// 			UniqueID: to.Ptr("a2f7fabd-bbc2-4a20-afe1-49fdb3885a28"),
	// 		},
	// 		Zones: []*string{
	// 			to.Ptr("1"),
	// 			to.Ptr("2"),
	// 		},
	// 		Identity: &armcomputefleet.ManagedServiceIdentity{
	// 			Type: to.Ptr(armcomputefleet.ManagedServiceIdentityTypeUserAssigned),
	// 			UserAssignedIdentities: map[string]*armcomputefleet.UserAssignedIdentity{
	// 			},
	// 			PrincipalID: to.Ptr("4d508e5b-374b-4382-9a1c-01fb8b6cb37c"),
	// 			TenantID: to.Ptr("5d508e5b-374b-4382-9a1c-01fb8b6cb37c"),
	// 		},
	// 		Tags: map[string]*string{
	// 		},
	// 		Location: to.Ptr("westus"),
	// 		Plan: &armcomputefleet.Plan{
	// 			Name: to.Ptr("jwgrcrnrtfoxn"),
	// 			Publisher: to.Ptr("iozjbiqqckqm"),
	// 			Product: to.Ptr("cgopbyvdyqikahwyxfpzwaqk"),
	// 			PromotionCode: to.Ptr("naglezezplcaruqogtxnuizslqnnbr"),
	// 			Version: to.Ptr("wa"),
	// 		},
	// 		ID: to.Ptr("/subscriptions/7B0CD4DB-3381-4013-9B31-FB6E6FD0FF1C/resourceGroups/rgazurefleet/providers/Microsoft.AzureFleet/fleets/myFleet"),
	// 		Name: to.Ptr("myFleet"),
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
