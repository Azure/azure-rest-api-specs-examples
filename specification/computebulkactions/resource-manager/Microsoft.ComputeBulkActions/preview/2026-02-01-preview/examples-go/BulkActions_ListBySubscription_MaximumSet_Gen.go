package armcomputebulkactions_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/computebulkactions/armcomputebulkactions"
)

// Generated from example definition: 2026-02-01-preview/BulkActions_ListBySubscription_MaximumSet_Gen.json
func ExampleBulkActionsClient_NewListBySubscriptionPager_bulkActionsListBySubscriptionGeneratedByMaximumSetRule() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcomputebulkactions.NewClientFactory("50352BBD-59F1-4EE2-BA9C-A6E51B0B1B2B", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewBulkActionsClient().NewListBySubscriptionPager("eastus2euap", nil)
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
		// page = armcomputebulkactions.BulkActionsClientListBySubscriptionResponse{
		// 	LaunchBulkInstancesOperationListResult: armcomputebulkactions.LaunchBulkInstancesOperationListResult{
		// 		Value: []*armcomputebulkactions.LocationBasedLaunchBulkInstancesOperation{
		// 			{
		// 				Properties: &armcomputebulkactions.LaunchBulkInstancesOperationProperties{
		// 					Capacity: to.Ptr[int32](24),
		// 					CapacityType: to.Ptr(armcomputebulkactions.CapacityTypeVM),
		// 					PriorityProfile: &armcomputebulkactions.PriorityProfile{
		// 						Type: to.Ptr(armcomputebulkactions.VirtualMachineTypeRegular),
		// 						MaxPricePerVM: to.Ptr[float32](21),
		// 						EvictionPolicy: to.Ptr(armcomputebulkactions.EvictionPolicyDelete),
		// 						AllocationStrategy: to.Ptr(armcomputebulkactions.AllocationStrategyLowestPrice),
		// 					},
		// 					VMSizesProfile: []*armcomputebulkactions.VMSizeProfile{
		// 						{
		// 							Name: to.Ptr("nolktwnfqdwikqiat"),
		// 							Rank: to.Ptr[int32](46189),
		// 						},
		// 					},
		// 					VMAttributes: &armcomputebulkactions.VMAttributes{
		// 						VCPUCount: &armcomputebulkactions.VMAttributeMinMaxInteger{
		// 							Min: to.Ptr[int32](0),
		// 							Max: to.Ptr[int32](0),
		// 						},
		// 						MemoryInGiB: &armcomputebulkactions.VMAttributeMinMaxDouble{
		// 							Min: to.Ptr[float64](0),
		// 							Max: to.Ptr[float64](0),
		// 						},
		// 						ArchitectureTypes: []*armcomputebulkactions.ArchitectureType{
		// 							to.Ptr(armcomputebulkactions.ArchitectureTypeARM64),
		// 						},
		// 						MemoryInGiBPerVCpu: &armcomputebulkactions.VMAttributeMinMaxDouble{
		// 							Min: to.Ptr[float64](0),
		// 							Max: to.Ptr[float64](0),
		// 						},
		// 						LocalStorageSupport: to.Ptr(armcomputebulkactions.VMAttributeSupportExcluded),
		// 						LocalStorageInGiB: &armcomputebulkactions.VMAttributeMinMaxDouble{
		// 							Min: to.Ptr[float64](0),
		// 							Max: to.Ptr[float64](0),
		// 						},
		// 						LocalStorageDiskTypes: []*armcomputebulkactions.LocalStorageDiskType{
		// 							to.Ptr(armcomputebulkactions.LocalStorageDiskTypeHDD),
		// 						},
		// 						DataDiskCount: &armcomputebulkactions.VMAttributeMinMaxInteger{
		// 							Min: to.Ptr[int32](0),
		// 							Max: to.Ptr[int32](0),
		// 						},
		// 						NetworkInterfaceCount: &armcomputebulkactions.VMAttributeMinMaxInteger{
		// 							Min: to.Ptr[int32](0),
		// 							Max: to.Ptr[int32](0),
		// 						},
		// 						NetworkBandwidthInMbps: &armcomputebulkactions.VMAttributeMinMaxDouble{
		// 							Min: to.Ptr[float64](0),
		// 							Max: to.Ptr[float64](0),
		// 						},
		// 						RdmaSupport: to.Ptr(armcomputebulkactions.VMAttributeSupportExcluded),
		// 						RdmaNetworkInterfaceCount: &armcomputebulkactions.VMAttributeMinMaxInteger{
		// 							Min: to.Ptr[int32](0),
		// 							Max: to.Ptr[int32](0),
		// 						},
		// 						AcceleratorSupport: to.Ptr(armcomputebulkactions.VMAttributeSupportExcluded),
		// 						AcceleratorManufacturers: []*armcomputebulkactions.AcceleratorManufacturer{
		// 							to.Ptr(armcomputebulkactions.AcceleratorManufacturerAMD),
		// 						},
		// 						AcceleratorTypes: []*armcomputebulkactions.AcceleratorType{
		// 							to.Ptr(armcomputebulkactions.AcceleratorTypeGPU),
		// 						},
		// 						AcceleratorCount: &armcomputebulkactions.VMAttributeMinMaxInteger{
		// 							Min: to.Ptr[int32](0),
		// 							Max: to.Ptr[int32](0),
		// 						},
		// 						VMCategories: []*armcomputebulkactions.VMCategory{
		// 							to.Ptr(armcomputebulkactions.VMCategoryGeneralPurpose),
		// 						},
		// 						CPUManufacturers: []*armcomputebulkactions.CPUManufacturer{
		// 							to.Ptr(armcomputebulkactions.CPUManufacturerIntel),
		// 						},
		// 						HyperVGenerations: []*armcomputebulkactions.HyperVGeneration{
		// 							to.Ptr(armcomputebulkactions.HyperVGenerationGen1),
		// 						},
		// 						BurstableSupport: to.Ptr(armcomputebulkactions.VMAttributeSupportExcluded),
		// 						AllowedVMSizes: []*string{
		// 							to.Ptr("dmtpdlqphckngwjhvkucfze"),
		// 						},
		// 						ExcludedVMSizes: []*string{
		// 							to.Ptr("yhjhharuhcyfxjnhxmflvsrdmei"),
		// 						},
		// 					},
		// 					ComputeProfile: &armcomputebulkactions.ComputeProfile{
		// 						VirtualMachineProfile: &armcomputebulkactions.VirtualMachineProfile{
		// 							OSProfile: &armcomputebulkactions.OSProfile{
		// 								AdminUsername: to.Ptr("tjdagcdhlpihlhkrz"),
		// 								CustomData: to.Ptr("jemgccf"),
		// 								WindowsConfiguration: &armcomputebulkactions.WindowsConfiguration{
		// 									ProvisionVMAgent: to.Ptr(true),
		// 									EnableAutomaticUpdates: to.Ptr(true),
		// 									TimeZone: to.Ptr("kiibvtut"),
		// 									AdditionalUnattendContent: []*armcomputebulkactions.AdditionalUnattendContent{
		// 										{
		// 											PassName: to.Ptr("OobeSystem"),
		// 											ComponentName: to.Ptr("Microsoft-Windows-Shell-Setup"),
		// 											SettingName: to.Ptr(armcomputebulkactions.SettingNamesAutoLogon),
		// 											Content: to.Ptr("zdpsub"),
		// 										},
		// 									},
		// 									PatchSettings: &armcomputebulkactions.PatchSettings{
		// 										PatchMode: to.Ptr(armcomputebulkactions.WindowsVMGuestPatchModeManual),
		// 										EnableHotpatching: to.Ptr(true),
		// 										AssessmentMode: to.Ptr(armcomputebulkactions.WindowsPatchAssessmentModeImageDefault),
		// 										AutomaticByPlatformSettings: &armcomputebulkactions.WindowsVMGuestPatchAutomaticByPlatformSettings{
		// 											RebootSetting: to.Ptr(armcomputebulkactions.WindowsVMGuestPatchAutomaticByPlatformRebootSettingUnknown),
		// 											BypassPlatformSafetyChecksOnUserSchedule: to.Ptr(true),
		// 										},
		// 									},
		// 									WinRM: &armcomputebulkactions.WinRMConfiguration{
		// 										Listeners: []*armcomputebulkactions.WinRMListener{
		// 											{
		// 												Protocol: to.Ptr(armcomputebulkactions.ProtocolTypesHTTP),
		// 												CertificateURL: to.Ptr("https://microsoft.com/a"),
		// 											},
		// 										},
		// 									},
		// 								},
		// 								LinuxConfiguration: &armcomputebulkactions.LinuxConfiguration{
		// 									DisablePasswordAuthentication: to.Ptr(true),
		// 									SSH: &armcomputebulkactions.SSHConfiguration{
		// 										PublicKeys: []*armcomputebulkactions.SSHPublicKey{
		// 											{
		// 												Path: to.Ptr("xyntsvqsiqsguyegxdvkmwhwz"),
		// 												KeyData: to.Ptr("mfk"),
		// 											},
		// 										},
		// 									},
		// 									ProvisionVMAgent: to.Ptr(true),
		// 									PatchSettings: &armcomputebulkactions.LinuxPatchSettings{
		// 										PatchMode: to.Ptr(armcomputebulkactions.LinuxVMGuestPatchModeImageDefault),
		// 										AssessmentMode: to.Ptr(armcomputebulkactions.LinuxPatchAssessmentModeImageDefault),
		// 										AutomaticByPlatformSettings: &armcomputebulkactions.LinuxVMGuestPatchAutomaticByPlatformSettings{
		// 											RebootSetting: to.Ptr(armcomputebulkactions.LinuxVMGuestPatchAutomaticByPlatformRebootSettingUnknown),
		// 											BypassPlatformSafetyChecksOnUserSchedule: to.Ptr(true),
		// 										},
		// 									},
		// 									EnableVMAgentPlatformUpdates: to.Ptr(true),
		// 								},
		// 								Secrets: []*armcomputebulkactions.VaultSecretGroup{
		// 									{
		// 										SourceVault: &armcomputebulkactions.SubResource{
		// 											ID: to.Ptr("obwiwwsgkdg"),
		// 										},
		// 										VaultCertificates: []*armcomputebulkactions.VaultCertificate{
		// 											{
		// 												CertificateURL: to.Ptr("https://microsoft.com/agmunp"),
		// 												CertificateStore: to.Ptr("zxrjtvfmltdj"),
		// 											},
		// 										},
		// 									},
		// 								},
		// 								AllowExtensionOperations: to.Ptr(true),
		// 								RequireGuestProvisionSignal: to.Ptr(true),
		// 								ComputerName: to.Ptr("jagkikqx"),
		// 							},
		// 							StorageProfile: &armcomputebulkactions.StorageProfile{
		// 								ImageReference: &armcomputebulkactions.ImageReference{
		// 									ID: to.Ptr("iwqrkiccafacxfctrxb"),
		// 									Publisher: to.Ptr("edjvbyisusjhyvnzgyvhixmnfzzsy"),
		// 									Offer: to.Ptr("olkxwdozixpjkjuk"),
		// 									SKU: to.Ptr("qmsq"),
		// 									Version: to.Ptr("hruassyajrafmgmub"),
		// 									SharedGalleryImageID: to.Ptr("ahzweiez"),
		// 									CommunityGalleryImageID: to.Ptr("bysd"),
		// 								},
		// 								OSDisk: &armcomputebulkactions.OSDisk{
		// 									Name: to.Ptr("pccysrjeo"),
		// 									Caching: to.Ptr(armcomputebulkactions.CachingTypesNone),
		// 									WriteAcceleratorEnabled: to.Ptr(true),
		// 									CreateOption: to.Ptr(armcomputebulkactions.DiskCreateOptionTypesFromImage),
		// 									DiffDiskSettings: &armcomputebulkactions.DiffDiskSettings{
		// 										Option: to.Ptr(armcomputebulkactions.DiffDiskOptionsLocal),
		// 										Placement: to.Ptr(armcomputebulkactions.DiffDiskPlacementCacheDisk),
		// 									},
		// 									DiskSizeGB: to.Ptr[int32](18),
		// 									OSType: to.Ptr(armcomputebulkactions.OperatingSystemTypesWindows),
		// 									Image: &armcomputebulkactions.VirtualHardDisk{
		// 										URI: to.Ptr("https://microsoft.com/a"),
		// 									},
		// 									ManagedDisk: &armcomputebulkactions.ManagedDiskParameters{
		// 										StorageAccountType: to.Ptr(armcomputebulkactions.StorageAccountTypesStandardLRS),
		// 										DiskEncryptionSet: &armcomputebulkactions.DiskEncryptionSetParameters{
		// 											ID: to.Ptr("thmks"),
		// 										},
		// 										SecurityProfile: &armcomputebulkactions.VMDiskSecurityProfile{
		// 											SecurityEncryptionType: to.Ptr(armcomputebulkactions.SecurityEncryptionTypesVMGuestStateOnly),
		// 											DiskEncryptionSet: &armcomputebulkactions.DiskEncryptionSetParameters{
		// 												ID: to.Ptr("thmks"),
		// 											},
		// 										},
		// 										ID: to.Ptr("wuqdcyunrkewr"),
		// 									},
		// 									DeleteOption: to.Ptr(armcomputebulkactions.DiskDeleteOptionTypesDelete),
		// 									EncryptionSettings: &armcomputebulkactions.DiskEncryptionSettings{
		// 										DiskEncryptionKey: &armcomputebulkactions.KeyVaultSecretReference{
		// 											SecretURL: to.Ptr("cuatfdkula"),
		// 											SourceVault: &armcomputebulkactions.SubResource{
		// 												ID: to.Ptr("ioypuofzltakyfcomjwfkmyz"),
		// 											},
		// 										},
		// 										KeyEncryptionKey: &armcomputebulkactions.KeyVaultKeyReference{
		// 											KeyURL: to.Ptr("imt"),
		// 											SourceVault: &armcomputebulkactions.SubResource{
		// 												ID: to.Ptr("ioypuofzltakyfcomjwfkmyz"),
		// 											},
		// 										},
		// 										Enabled: to.Ptr(true),
		// 									},
		// 									Vhd: &armcomputebulkactions.VirtualHardDisk{
		// 										URI: to.Ptr("anvtwgmfthxmyhdnbvabmzyrknxwf"),
		// 									},
		// 								},
		// 								DataDisks: []*armcomputebulkactions.DataDisk{
		// 									{
		// 										Name: to.Ptr("aq"),
		// 										Lun: to.Ptr[int32](1),
		// 										Caching: to.Ptr(armcomputebulkactions.CachingTypesNone),
		// 										WriteAcceleratorEnabled: to.Ptr(true),
		// 										CreateOption: to.Ptr(armcomputebulkactions.DiskCreateOptionTypesFromImage),
		// 										DiskSizeGB: to.Ptr[int32](24),
		// 										ManagedDisk: &armcomputebulkactions.ManagedDiskParameters{
		// 											StorageAccountType: to.Ptr(armcomputebulkactions.StorageAccountTypesStandardLRS),
		// 											DiskEncryptionSet: &armcomputebulkactions.DiskEncryptionSetParameters{
		// 												ID: to.Ptr("thmks"),
		// 											},
		// 											SecurityProfile: &armcomputebulkactions.VMDiskSecurityProfile{
		// 												SecurityEncryptionType: to.Ptr(armcomputebulkactions.SecurityEncryptionTypesVMGuestStateOnly),
		// 												DiskEncryptionSet: &armcomputebulkactions.DiskEncryptionSetParameters{
		// 													ID: to.Ptr("thmks"),
		// 												},
		// 											},
		// 											ID: to.Ptr("zcoqnxlomkordbdolkxraqbwgsh"),
		// 										},
		// 										DeleteOption: to.Ptr(armcomputebulkactions.DiskDeleteOptionTypesDelete),
		// 										Vhd: &armcomputebulkactions.VirtualHardDisk{
		// 											URI: to.Ptr("anvtwgmfthxmyhdnbvabmzyrknxwf"),
		// 										},
		// 										Image: &armcomputebulkactions.VirtualHardDisk{
		// 											URI: to.Ptr("anvtwgmfthxmyhdnbvabmzyrknxwf"),
		// 										},
		// 										SourceResource: &armcomputebulkactions.APIEntityReference{
		// 											ID: to.Ptr("fpabycyqmkqqfdfrzqmnykmy"),
		// 										},
		// 										ToBeDetached: to.Ptr(true),
		// 										DetachOption: to.Ptr(armcomputebulkactions.DiskDetachOptionTypesForceDetach),
		// 									},
		// 								},
		// 								DiskControllerType: to.Ptr(armcomputebulkactions.DiskControllerTypesSCSI),
		// 							},
		// 							NetworkProfile: &armcomputebulkactions.NetworkProfile{
		// 								NetworkInterfaceConfigurations: []*armcomputebulkactions.VirtualMachineNetworkInterfaceConfiguration{
		// 									{
		// 										Name: to.Ptr("keppldrpxjgckgsmq"),
		// 										Properties: &armcomputebulkactions.VirtualMachineNetworkInterfaceConfigurationProperties{
		// 											Primary: to.Ptr(true),
		// 											EnableAcceleratedNetworking: to.Ptr(true),
		// 											DisableTCPStateTracking: to.Ptr(true),
		// 											EnableFpga: to.Ptr(true),
		// 											NetworkSecurityGroup: &armcomputebulkactions.SubResource{
		// 												ID: to.Ptr("obwiwwsgkdg"),
		// 											},
		// 											DNSSettings: &armcomputebulkactions.VirtualMachineNetworkInterfaceDNSSettingsConfiguration{
		// 												DNSServers: []*string{
		// 													to.Ptr("pnhvxygytoozxmkt"),
		// 												},
		// 											},
		// 											IPConfigurations: []*armcomputebulkactions.VirtualMachineNetworkInterfaceIPConfiguration{
		// 												{
		// 													Name: to.Ptr("nqjufbencyticmohsdxogwiu"),
		// 													Properties: &armcomputebulkactions.VirtualMachineNetworkInterfaceIPConfigurationProperties{
		// 														Subnet: &armcomputebulkactions.SubResource{
		// 															ID: to.Ptr("djtafmblvomuabwlhlyoxzgdkwkz"),
		// 														},
		// 														Primary: to.Ptr(true),
		// 														PublicIPAddressConfiguration: &armcomputebulkactions.VirtualMachinePublicIPAddressConfiguration{
		// 															Name: to.Ptr("kgvjhctjspzldadcmtgsojglhmj"),
		// 															Properties: &armcomputebulkactions.VirtualMachinePublicIPAddressConfigurationProperties{
		// 																IdleTimeoutInMinutes: to.Ptr[int32](22),
		// 																DNSSettings: &armcomputebulkactions.VirtualMachinePublicIPAddressDNSSettingsConfiguration{
		// 																	DomainNameLabel: to.Ptr("vsvbcpusndz"),
		// 																	DomainNameLabelScope: to.Ptr(armcomputebulkactions.DomainNameLabelScopeTypesTenantReuse),
		// 																},
		// 																IPTags: []*armcomputebulkactions.VirtualMachineIPTag{
		// 																	{
		// 																		IPTagType: to.Ptr("hengwhngakjlsmhuegnlbtpmiihf"),
		// 																		Tag: to.Ptr("zlnuzjdbdnwbtep"),
		// 																	},
		// 																},
		// 																PublicIPPrefix: &armcomputebulkactions.SubResource{
		// 																	ID: to.Ptr("obwiwwsgkdg"),
		// 																},
		// 																PublicIPAddressVersion: to.Ptr(armcomputebulkactions.IPVersionsIPv4),
		// 																DeleteOption: to.Ptr(armcomputebulkactions.DeleteOptionsDelete),
		// 																PublicIPAllocationMethod: to.Ptr(armcomputebulkactions.PublicIPAllocationMethodDynamic),
		// 															},
		// 															SKU: &armcomputebulkactions.PublicIPAddressSKU{
		// 																Name: to.Ptr(armcomputebulkactions.PublicIPAddressSKUNameBasic),
		// 																Tier: to.Ptr(armcomputebulkactions.PublicIPAddressSKUTierRegional),
		// 															},
		// 															Tags: map[string]*string{
		// 															},
		// 														},
		// 														PrivateIPAddressVersion: to.Ptr(armcomputebulkactions.IPVersionsIPv4),
		// 														ApplicationGatewayBackendAddressPools: []*armcomputebulkactions.SubResource{
		// 															{
		// 																ID: to.Ptr("obwiwwsgkdg"),
		// 															},
		// 														},
		// 														ApplicationSecurityGroups: []*armcomputebulkactions.SubResource{
		// 															{
		// 																ID: to.Ptr("obwiwwsgkdg"),
		// 															},
		// 														},
		// 														LoadBalancerBackendAddressPools: []*armcomputebulkactions.SubResource{
		// 															{
		// 																ID: to.Ptr("obwiwwsgkdg"),
		// 															},
		// 														},
		// 													},
		// 												},
		// 											},
		// 											EnableIPForwarding: to.Ptr(true),
		// 											DeleteOption: to.Ptr(armcomputebulkactions.DeleteOptionsDelete),
		// 											AuxiliaryMode: to.Ptr(armcomputebulkactions.NetworkInterfaceAuxiliaryModeNone),
		// 											AuxiliarySKU: to.Ptr(armcomputebulkactions.NetworkInterfaceAuxiliarySKUNone),
		// 											DscpConfiguration: &armcomputebulkactions.SubResource{
		// 												ID: to.Ptr("ioypuofzltakyfcomjwfkmyz"),
		// 											},
		// 										},
		// 										Tags: map[string]*string{
		// 										},
		// 									},
		// 								},
		// 								NetworkAPIVersion: to.Ptr(armcomputebulkactions.NetworkAPIVersion20201101),
		// 								NetworkInterfaces: []*armcomputebulkactions.NetworkInterfaceReference{
		// 									{
		// 										Properties: &armcomputebulkactions.NetworkInterfaceReferenceProperties{
		// 											Primary: to.Ptr(true),
		// 											DeleteOption: to.Ptr(armcomputebulkactions.DeleteOptionsDelete),
		// 										},
		// 										ID: to.Ptr("bmlqsytfgnkwgkibsmsoeh"),
		// 									},
		// 								},
		// 							},
		// 							SecurityProfile: &armcomputebulkactions.SecurityProfile{
		// 								UefiSettings: &armcomputebulkactions.UefiSettings{
		// 									SecureBootEnabled: to.Ptr(true),
		// 									VTpmEnabled: to.Ptr(true),
		// 								},
		// 								EncryptionAtHost: to.Ptr(true),
		// 								SecurityType: to.Ptr(armcomputebulkactions.SecurityTypesTrustedLaunch),
		// 								EncryptionIdentity: &armcomputebulkactions.EncryptionIdentity{
		// 									UserAssignedIdentityResourceID: to.Ptr("wkiykqbrifryaruzokophodpjig"),
		// 								},
		// 								ProxyAgentSettings: &armcomputebulkactions.ProxyAgentSettings{
		// 									Enabled: to.Ptr(true),
		// 									Mode: to.Ptr(armcomputebulkactions.ModeAudit),
		// 									KeyIncarnationID: to.Ptr[int32](17),
		// 									WireServer: &armcomputebulkactions.HostEndpointSettings{
		// 										Mode: to.Ptr(armcomputebulkactions.ModesAudit),
		// 										InVMAccessControlProfileReferenceID: to.Ptr("cubhuucckqkxbifmertj"),
		// 									},
		// 									Imds: &armcomputebulkactions.HostEndpointSettings{
		// 										Mode: to.Ptr(armcomputebulkactions.ModesAudit),
		// 										InVMAccessControlProfileReferenceID: to.Ptr("cubhuucckqkxbifmertj"),
		// 									},
		// 									AddProxyAgentExtension: to.Ptr(true),
		// 								},
		// 							},
		// 							DiagnosticsProfile: &armcomputebulkactions.DiagnosticsProfile{
		// 								BootDiagnostics: &armcomputebulkactions.BootDiagnostics{
		// 									Enabled: to.Ptr(true),
		// 									StorageURI: to.Ptr("https://microsoft.com/a"),
		// 								},
		// 							},
		// 							LicenseType: to.Ptr("iipnwxwfkfbbouzbwicqxnxicjz"),
		// 							ScheduledEventsProfile: &armcomputebulkactions.ScheduledEventsProfile{
		// 								TerminateNotificationProfile: &armcomputebulkactions.TerminateNotificationProfile{
		// 									NotBeforeTimeout: to.Ptr("ypif"),
		// 									Enable: to.Ptr(true),
		// 								},
		// 								OSImageNotificationProfile: &armcomputebulkactions.OSImageNotificationProfile{
		// 									NotBeforeTimeout: to.Ptr("fztbudpjkicyigtvltlbszmivfbmb"),
		// 									Enable: to.Ptr(true),
		// 								},
		// 							},
		// 							UserData: to.Ptr("qcsgczwavs"),
		// 							CapacityReservation: &armcomputebulkactions.CapacityReservationProfile{
		// 								CapacityReservationGroup: &armcomputebulkactions.SubResource{
		// 									ID: to.Ptr("obwiwwsgkdg"),
		// 								},
		// 							},
		// 							ApplicationProfile: &armcomputebulkactions.ApplicationProfile{
		// 								GalleryApplications: []*armcomputebulkactions.VMGalleryApplication{
		// 									{
		// 										Tags: to.Ptr("qgn"),
		// 										Order: to.Ptr[int32](14),
		// 										PackageReferenceID: to.Ptr("soddwzqduyolzz"),
		// 										ConfigurationReference: to.Ptr("mddsvaruvzvblkafsotscupperzu"),
		// 										TreatFailureAsDeploymentFailure: to.Ptr(true),
		// 										EnableAutomaticUpgrade: to.Ptr(true),
		// 									},
		// 								},
		// 							},
		// 							ScheduledEventsPolicy: &armcomputebulkactions.ScheduledEventsPolicy{
		// 								UserInitiatedRedeploy: &armcomputebulkactions.UserInitiatedRedeploy{
		// 									AutomaticallyApprove: to.Ptr(true),
		// 								},
		// 								UserInitiatedReboot: &armcomputebulkactions.UserInitiatedReboot{
		// 									AutomaticallyApprove: to.Ptr(true),
		// 								},
		// 								ScheduledEventsAdditionalPublishingTargets: &armcomputebulkactions.ScheduledEventsAdditionalPublishingTargets{
		// 									EventGridAndResourceGraph: &armcomputebulkactions.EventGridAndResourceGraph{
		// 										Enable: to.Ptr(true),
		// 										ScheduledEventsAPIVersion: to.Ptr("sbzjonqss"),
		// 									},
		// 								},
		// 								AllInstancesDown: &armcomputebulkactions.AllInstancesDown{
		// 									AutomaticallyApprove: to.Ptr(true),
		// 								},
		// 							},
		// 							AdditionalCapabilities: &armcomputebulkactions.AdditionalCapabilities{
		// 								UltraSSDEnabled: to.Ptr(true),
		// 								HibernationEnabled: to.Ptr(true),
		// 							},
		// 							ExtensionsTimeBudget: to.Ptr("hvrimblcqujozpeurohjcn"),
		// 						},
		// 						ComputeAPIVersion: to.Ptr("bccikdfzgrygwpefvmvptutceg"),
		// 						Extensions: []*armcomputebulkactions.VirtualMachineExtension{
		// 							{
		// 								Name: to.Ptr("gj"),
		// 								Properties: &armcomputebulkactions.VirtualMachineExtensionProperties{
		// 									ForceUpdateTag: to.Ptr("mistpuvreycjbhahmcvgkjskeiop"),
		// 									Publisher: to.Ptr("rzsodcysrfxkrgnrjqlpfqe"),
		// 									Type: to.Ptr("eyehf"),
		// 									TypeHandlerVersion: to.Ptr("wezzz"),
		// 									AutoUpgradeMinorVersion: to.Ptr(true),
		// 									EnableAutomaticUpgrade: to.Ptr(true),
		// 									SuppressFailures: to.Ptr(true),
		// 									ProtectedSettingsFromKeyVault: &armcomputebulkactions.KeyVaultSecretReference{
		// 										SecretURL: to.Ptr("cuatfdkula"),
		// 										SourceVault: &armcomputebulkactions.SubResource{
		// 											ID: to.Ptr("ioypuofzltakyfcomjwfkmyz"),
		// 										},
		// 									},
		// 									ProvisionAfterExtensions: []*string{
		// 										to.Ptr("jddcihtuzdczkvkryhktzjlf"),
		// 									},
		// 									Settings: map[string]any{
		// 									},
		// 									ProtectedSettings: map[string]any{
		// 									},
		// 								},
		// 							},
		// 						},
		// 					},
		// 					ZoneAllocationPolicy: &armcomputebulkactions.ZoneAllocationPolicy{
		// 						DistributionStrategy: to.Ptr(armcomputebulkactions.ZoneDistributionStrategyBestEffortSingleZone),
		// 						ZonePreferences: []*armcomputebulkactions.ZonePreference{
		// 							{
		// 								Zone: to.Ptr("voauikerqjpeepaeaokkcybyjd"),
		// 								Rank: to.Ptr[int32](46292),
		// 							},
		// 						},
		// 					},
		// 					RetryPolicy: &armcomputebulkactions.RetryPolicy{
		// 						RetryCount: to.Ptr[int32](9),
		// 						RetryWindowInMinutes: to.Ptr[int32](21),
		// 					},
		// 					ProvisioningState: to.Ptr(armcomputebulkactions.ProvisioningStateSucceeded),
		// 				},
		// 				Zones: []*string{
		// 					to.Ptr("cyriutfcgydtaezeso"),
		// 				},
		// 				Identity: &armcomputebulkactions.ManagedServiceIdentity{
		// 					Type: to.Ptr(armcomputebulkactions.ManagedServiceIdentityTypeNone),
		// 					UserAssignedIdentities: map[string]*armcomputebulkactions.UserAssignedIdentity{
		// 					},
		// 					PrincipalID: to.Ptr("2ebdc740-a760-4406-80e9-6be95ebc9743"),
		// 					TenantID: to.Ptr("3ec2ab23-9f13-4328-85c8-21928acbc7b8"),
		// 				},
		// 				Plan: &armcomputebulkactions.Plan{
		// 					Name: to.Ptr("owvrgjbxrkj"),
		// 					Publisher: to.Ptr("qhybdqbljmztcjujxal"),
		// 					Product: to.Ptr("rlhap"),
		// 					PromotionCode: to.Ptr("agypojbtdxvgqgisautnhcoysgy"),
		// 					Version: to.Ptr("ghmnlomqg"),
		// 				},
		// 				ID: to.Ptr("/subscriptions/50352BBD-59F1-4EE2-BA9C-A6E51B0B1B2B/resourceGroups/rgcomputebulkactions/providers/Microsoft.Compute/bulkActions/3ec2ab23-9f13-4328-85c8-21928acbc7b8"),
		// 				Name: to.Ptr("kv"),
		// 				Type: to.Ptr("v"),
		// 				SystemData: &armcomputebulkactions.SystemData{
		// 					CreatedBy: to.Ptr("kmjag"),
		// 					CreatedByType: to.Ptr(armcomputebulkactions.CreatedByTypeUser),
		// 					CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2025-10-02T21:36:21.174Z"); return t}()),
		// 					LastModifiedBy: to.Ptr("fwcgdzyolflmgnxnnmpoevpwr"),
		// 					LastModifiedByType: to.Ptr(armcomputebulkactions.CreatedByTypeUser),
		// 					LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2025-10-02T21:36:21.174Z"); return t}()),
		// 				},
		// 			},
		// 		},
		// 		NextLink: to.Ptr("https://microsoft.com/a"),
		// 	},
		// }
	}
}
