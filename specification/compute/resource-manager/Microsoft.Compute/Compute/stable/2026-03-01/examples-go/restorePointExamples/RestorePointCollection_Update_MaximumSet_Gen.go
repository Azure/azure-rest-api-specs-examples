package armcompute_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/compute/armcompute/v8"
)

// Generated from example definition: 2026-03-01/restorePointExamples/RestorePointCollection_Update_MaximumSet_Gen.json
func ExampleRestorePointCollectionsClient_Update_restorePointCollectionUpdateMaximumSetGen() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcompute.NewClientFactory("{subscription-id}", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewRestorePointCollectionsClient().Update(ctx, "rgcompute", "aaaaaaaaaaaaaaaaaaaa", armcompute.RestorePointCollectionUpdate{
		Properties: &armcompute.RestorePointCollectionProperties{
			Source: &armcompute.RestorePointCollectionSourceProperties{
				ID: to.Ptr("/subscriptions/{subscription-id}/resourceGroups/myResourceGroup/providers/Microsoft.Compute/virtualMachines/myVM"),
			},
			InstantAccess: to.Ptr(true),
		},
		Tags: map[string]*string{
			"key8536": to.Ptr("aaaaaaaaaaaaaaaaaaa"),
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armcompute.RestorePointCollectionsClientUpdateResponse{
	// 	RestorePointCollection: armcompute.RestorePointCollection{
	// 		Location: to.Ptr("norwayeast"),
	// 		Properties: &armcompute.RestorePointCollectionProperties{
	// 			Source: &armcompute.RestorePointCollectionSourceProperties{
	// 				ID: to.Ptr("/subscriptions/{subscription-id}/resourceGroups/myResourceGroup/providers/Microsoft.Compute/virtualMachines/myVM"),
	// 				Location: to.Ptr("eastus"),
	// 			},
	// 			InstantAccess: to.Ptr(true),
	// 			ProvisioningState: to.Ptr("Successful"),
	// 			RestorePointCollectionID: to.Ptr("638f052b-a7c2-450c-89e7-6a3b8f1d6a7c"),
	// 			RestorePoints: []*armcompute.RestorePoint{
	// 				{
	// 					Properties: &armcompute.RestorePointProperties{
	// 						ExcludeDisks: []*armcompute.APIEntityReference{
	// 							{
	// 								ID: to.Ptr("/subscriptions/{subscription-id}/resourceGroups/myResourceGroup/providers/Microsoft.Compute/restorePointCollections/rpcName/restorePoints/restorePointName/diskRestorePoints/testingexcludedisk_OsDisk_1_74cdaedcea50483d9833c96adefa100f_22b4bdfe-6c54-4f72-84d8-85d8860f0c57"),
	// 							},
	// 						},
	// 						SourceMetadata: &armcompute.RestorePointSourceMetadata{
	// 							HardwareProfile: &armcompute.HardwareProfile{
	// 								VMSize: to.Ptr(armcompute.VirtualMachineSizeTypesStandardB1S),
	// 								VMSizeProperties: &armcompute.VMSizeProperties{
	// 									VCPUsAvailable: to.Ptr[int32](9),
	// 									VCPUsPerCore: to.Ptr[int32](12),
	// 								},
	// 							},
	// 							StorageProfile: &armcompute.RestorePointSourceVMStorageProfile{
	// 								OSDisk: &armcompute.RestorePointSourceVMOSDisk{
	// 									OSType: to.Ptr(armcompute.OperatingSystemTypeWindows),
	// 									EncryptionSettings: &armcompute.DiskEncryptionSettings{
	// 										DiskEncryptionKey: &armcompute.KeyVaultSecretReference{
	// 											SecretURL: to.Ptr("aaaaaaaa"),
	// 											SourceVault: &armcompute.SubResource{
	// 												ID: to.Ptr("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/availabilitySets/{availabilitySetName}"),
	// 											},
	// 										},
	// 										KeyEncryptionKey: &armcompute.KeyVaultKeyReference{
	// 											KeyURL: to.Ptr("aaaaaaaaaaaaaa"),
	// 											SourceVault: &armcompute.SubResource{
	// 												ID: to.Ptr("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/availabilitySets/{availabilitySetName}"),
	// 											},
	// 										},
	// 										Enabled: to.Ptr(true),
	// 									},
	// 									Name: to.Ptr("testingexcludedisk_OsDisk_1_74cdaedcea50483d9833c96adefa100f"),
	// 									Caching: to.Ptr(armcompute.CachingTypesReadWrite),
	// 									DiskSizeGB: to.Ptr[int32](3),
	// 									ManagedDisk: &armcompute.ManagedDiskParameters{
	// 										StorageAccountType: to.Ptr(armcompute.StorageAccountTypesStandardLRS),
	// 										DiskEncryptionSet: &armcompute.DiskEncryptionSetParameters{
	// 											ID: to.Ptr("aaaaaaaaaaaa"),
	// 										},
	// 										ID: to.Ptr("/subscriptions/{subscription-id}/resourceGroups/myResourceGroup/providers/Microsoft.Compute/disks/testingexcludedisk_OsDisk_1_74cdaedcea50483d9833c96adefa100f"),
	// 									},
	// 									DiskRestorePoint: &armcompute.DiskRestorePointAttributes{
	// 										ID: to.Ptr("/subscriptions/{subscription-id}/resourceGroups/myResourceGroup/providers/Microsoft.Compute/restorePointCollections/rpcName/restorePoints/restorePointName/diskRestorePoints/testingexcludedisk_OsDisk_1_74cdaedcea50483d9833c96adefa100f_22b4bdfe-6c54-4f72-84d8-85d8860f0c57"),
	// 									},
	// 								},
	// 								DataDisks: []*armcompute.RestorePointSourceVMDataDisk{
	// 									{
	// 										Lun: to.Ptr[int32](1),
	// 										Name: to.Ptr("testingexcludedisk_DataDisk_1"),
	// 										Caching: to.Ptr(armcompute.CachingTypesNone),
	// 										ManagedDisk: &armcompute.ManagedDiskParameters{
	// 											StorageAccountType: to.Ptr(armcompute.StorageAccountTypesStandardLRS),
	// 											ID: to.Ptr("/subscriptions/{subscription-id}/resourceGroups/myResourceGroup/providers/Microsoft.Compute/disks/testingexcludedisk_DataDisk_1"),
	// 											DiskEncryptionSet: &armcompute.DiskEncryptionSetParameters{
	// 												ID: to.Ptr("aaaaaaaaaaaa"),
	// 											},
	// 										},
	// 										DiskRestorePoint: &armcompute.DiskRestorePointAttributes{
	// 											ID: to.Ptr("/subscriptions/{subscription-id}/resourceGroups/myResourceGroup/providers/Microsoft.Compute/restorePointCollections/rpcName/restorePoints/restorePointName/diskRestorePoints/testingexcludedisk_DataDisk_1_68785190-1acb-4d5e-a8ae-705b45f3dca5"),
	// 										},
	// 										DiskSizeGB: to.Ptr[int32](24),
	// 									},
	// 								},
	// 								DiskControllerType: to.Ptr(armcompute.DiskControllerTypesNVMe),
	// 							},
	// 							OSProfile: &armcompute.OSProfile{
	// 								ComputerName: to.Ptr("computerName"),
	// 								AdminUsername: to.Ptr("admin"),
	// 								CustomData: to.Ptr("aaaaaaaaaaaaaaaaaaaaaaaaaaa"),
	// 								WindowsConfiguration: &armcompute.WindowsConfiguration{
	// 									ProvisionVMAgent: to.Ptr(true),
	// 									EnableAutomaticUpdates: to.Ptr(true),
	// 									TimeZone: to.Ptr("aaaaaaaaaaaaaaaaaaaaaaaaaaaaa"),
	// 									AdditionalUnattendContent: []*armcompute.AdditionalUnattendContent{
	// 										{
	// 											PassName: to.Ptr("OobeSystem"),
	// 											ComponentName: to.Ptr("Microsoft-Windows-Shell-Setup"),
	// 											SettingName: to.Ptr(armcompute.SettingNamesAutoLogon),
	// 											Content: to.Ptr("aaaaaaaaaaaaaaaaaaaa"),
	// 										},
	// 									},
	// 									PatchSettings: &armcompute.PatchSettings{
	// 										PatchMode: to.Ptr(armcompute.WindowsVMGuestPatchModeManual),
	// 										EnableHotpatching: to.Ptr(true),
	// 										AssessmentMode: to.Ptr(armcompute.WindowsPatchAssessmentModeImageDefault),
	// 									},
	// 									WinRM: &armcompute.WinRMConfiguration{
	// 										Listeners: []*armcompute.WinRMListener{
	// 											{
	// 												Protocol: to.Ptr(armcompute.ProtocolTypesHTTP),
	// 												CertificateURL: to.Ptr("aaaaaaaaaaaaaaaaaaaaaa"),
	// 											},
	// 										},
	// 									},
	// 								},
	// 								LinuxConfiguration: &armcompute.LinuxConfiguration{
	// 									DisablePasswordAuthentication: to.Ptr(true),
	// 									SSH: &armcompute.SSHConfiguration{
	// 										PublicKeys: []*armcompute.SSHPublicKey{
	// 											{
	// 												Path: to.Ptr("aaa"),
	// 												KeyData: to.Ptr("aaaaaa"),
	// 											},
	// 										},
	// 									},
	// 									ProvisionVMAgent: to.Ptr(true),
	// 									PatchSettings: &armcompute.LinuxPatchSettings{
	// 										PatchMode: to.Ptr(armcompute.LinuxVMGuestPatchModeImageDefault),
	// 										AssessmentMode: to.Ptr(armcompute.LinuxPatchAssessmentModeImageDefault),
	// 									},
	// 								},
	// 								Secrets: []*armcompute.VaultSecretGroup{
	// 									{
	// 										SourceVault: &armcompute.SubResource{
	// 											ID: to.Ptr("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/availabilitySets/{availabilitySetName}"),
	// 										},
	// 										VaultCertificates: []*armcompute.VaultCertificate{
	// 											{
	// 												CertificateURL: to.Ptr("aaaaaaa"),
	// 												CertificateStore: to.Ptr("aaaaaaaaaaaaaaaaaaaaaaaaa"),
	// 											},
	// 										},
	// 									},
	// 								},
	// 								AllowExtensionOperations: to.Ptr(true),
	// 								RequireGuestProvisionSignal: to.Ptr(true),
	// 							},
	// 							DiagnosticsProfile: &armcompute.DiagnosticsProfile{
	// 								BootDiagnostics: &armcompute.BootDiagnostics{
	// 									Enabled: to.Ptr(true),
	// 									StorageURI: to.Ptr("aaaaaaaaaaaaaaaaaaa"),
	// 								},
	// 							},
	// 							LicenseType: to.Ptr("aaaaaaaaaaaaaaaaaaaaaaaaaaa"),
	// 							VMID: to.Ptr("76d6541e-80bd-4dc1-932b-3cae4cfb80e7"),
	// 							SecurityProfile: &armcompute.SecurityProfile{
	// 								UefiSettings: &armcompute.UefiSettings{
	// 									SecureBootEnabled: to.Ptr(true),
	// 									VTpmEnabled: to.Ptr(true),
	// 								},
	// 								EncryptionAtHost: to.Ptr(true),
	// 								SecurityType: to.Ptr(armcompute.SecurityTypesTrustedLaunch),
	// 							},
	// 							Location: to.Ptr("westus"),
	// 						},
	// 						ProvisioningState: to.Ptr("aaaaaaaaaaaaaaaaa"),
	// 						ConsistencyMode: to.Ptr(armcompute.ConsistencyModeTypesCrashConsistent),
	// 						TimeCreated: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-11-30T12:58:26.593Z"); return t}()),
	// 					},
	// 					ID: to.Ptr("aaaaaaaaaaa"),
	// 					Name: to.Ptr("aaaaaaaaaaaaaaaaaa"),
	// 					Type: to.Ptr("aaaaaaaaaaaaaaaaaaaaaaaaa"),
	// 				},
	// 			},
	// 		},
	// 		Tags: map[string]*string{
	// 		},
	// 		ID: to.Ptr("/subscriptions/{subscription-id}/resourceGroups/myResourceGroup/providers/Microsoft.Compute/restorePointCollections/myRpc"),
	// 		Name: to.Ptr("myRpc"),
	// 		Type: to.Ptr("Microsoft.Compute/restorePointCollections"),
	// 	},
	// }
}
