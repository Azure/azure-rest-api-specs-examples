package armcompute_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/compute/armcompute/v8"
)

// Generated from example definition: 2026-03-01/restorePointExamples/RestorePointCollection_Get_WithContainedRestorePoints.json
func ExampleRestorePointCollectionsClient_Get_getARestorePointCollectionIncludingTheRestorePointsContainedInTheRestorePointCollection() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcompute.NewClientFactory("{subscription-id}", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewRestorePointCollectionsClient().Get(ctx, "myResourceGroup", "rpcName", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armcompute.RestorePointCollectionsClientGetResponse{
	// 	RestorePointCollection: armcompute.RestorePointCollection{
	// 		Name: to.Ptr("rpcName"),
	// 		ID: to.Ptr("/subscriptions/{subscription-id}/resourceGroups/myResourceGroup/providers/Microsoft.Compute/restorePointCollections/rpcName"),
	// 		Type: to.Ptr("Microsoft.Compute/restorePointCollections"),
	// 		Location: to.Ptr("westus"),
	// 		Tags: map[string]*string{
	// 			"myTag1": to.Ptr("tagValue1"),
	// 		},
	// 		Properties: &armcompute.RestorePointCollectionProperties{
	// 			Source: &armcompute.RestorePointCollectionSourceProperties{
	// 				ID: to.Ptr("/subscriptions/{subscription-id}/resourceGroups/myResourceGroup/providers/Microsoft.Compute/virtualMachines/myVM"),
	// 				Location: to.Ptr("eastus"),
	// 			},
	// 			InstantAccess: to.Ptr(true),
	// 			RestorePointCollectionID: to.Ptr("59f04a5d-f783-4200-a1bd-d3f464e8c4b4"),
	// 			ProvisioningState: to.Ptr("Succeeded"),
	// 			RestorePoints: []*armcompute.RestorePoint{
	// 				{
	// 					Name: to.Ptr("restorePointName"),
	// 					ID: to.Ptr("/subscriptions/{subscription-id}/resourceGroups/myResourceGroup/providers/Microsoft.Compute/restorePointCollections/rpcName/restorePoints/restorePointName"),
	// 					Properties: &armcompute.RestorePointProperties{
	// 						ExcludeDisks: []*armcompute.APIEntityReference{
	// 							{
	// 								ID: to.Ptr("/subscriptions/{subscription-id}/resourceGroups/myResourceGroup/providers/Microsoft.Compute/disks/vm8768_disk2_fe6ffde4f69b491ca33fb984d5bcd89f"),
	// 							},
	// 						},
	// 						SourceMetadata: &armcompute.RestorePointSourceMetadata{
	// 							VMID: to.Ptr("76d6541e-80bd-4dc1-932b-3cae4cfb80e7"),
	// 							HardwareProfile: &armcompute.HardwareProfile{
	// 								VMSize: to.Ptr(armcompute.VirtualMachineSizeTypesStandardB1S),
	// 							},
	// 							StorageProfile: &armcompute.RestorePointSourceVMStorageProfile{
	// 								OSDisk: &armcompute.RestorePointSourceVMOSDisk{
	// 									OSType: to.Ptr(armcompute.OperatingSystemTypeWindows),
	// 									Name: to.Ptr("testingexcludedisk_OsDisk_1_74cdaedcea50483d9833c96adefa100f"),
	// 									Caching: to.Ptr(armcompute.CachingTypesReadWrite),
	// 									ManagedDisk: &armcompute.ManagedDiskParameters{
	// 										StorageAccountType: to.Ptr(armcompute.StorageAccountTypesStandardLRS),
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
	// 										},
	// 										DiskRestorePoint: &armcompute.DiskRestorePointAttributes{
	// 											ID: to.Ptr("/subscriptions/{subscription-id}/resourceGroups/myResourceGroup/providers/Microsoft.Compute/restorePointCollections/rpcName/restorePoints/restorePointName/diskRestorePoints/testingexcludedisk_DataDisk_1_68785190-1acb-4d5e-a8ae-705b45f3dca5"),
	// 										},
	// 									},
	// 								},
	// 								DiskControllerType: to.Ptr(armcompute.DiskControllerTypesNVMe),
	// 							},
	// 							OSProfile: &armcompute.OSProfile{
	// 								ComputerName: to.Ptr("computerName"),
	// 								AdminUsername: to.Ptr("admin"),
	// 								WindowsConfiguration: &armcompute.WindowsConfiguration{
	// 									ProvisionVMAgent: to.Ptr(true),
	// 									EnableAutomaticUpdates: to.Ptr(true),
	// 								},
	// 								Secrets: []*armcompute.VaultSecretGroup{
	// 								},
	// 								AllowExtensionOperations: to.Ptr(true),
	// 								RequireGuestProvisionSignal: to.Ptr(true),
	// 							},
	// 							DiagnosticsProfile: &armcompute.DiagnosticsProfile{
	// 								BootDiagnostics: &armcompute.BootDiagnostics{
	// 									Enabled: to.Ptr(true),
	// 								},
	// 							},
	// 							Location: to.Ptr("westus"),
	// 						},
	// 						InstantAccessDurationMinutes: to.Ptr[int32](120),
	// 						ProvisioningState: to.Ptr("Succeeded"),
	// 						ConsistencyMode: to.Ptr(armcompute.ConsistencyModeTypesApplicationConsistent),
	// 						TimeCreated: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-01-27T20:35:05.8401519+00:00"); return t}()),
	// 					},
	// 				},
	// 			},
	// 		},
	// 	},
	// }
}
