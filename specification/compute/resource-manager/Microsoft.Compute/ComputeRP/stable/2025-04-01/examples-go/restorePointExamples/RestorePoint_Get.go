package armcompute_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/compute/armcompute/v7"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/ccea12be6cd5e64743499d46f7a9c8b60df52db1/specification/compute/resource-manager/Microsoft.Compute/ComputeRP/stable/2025-04-01/examples/restorePointExamples/RestorePoint_Get.json
func ExampleRestorePointsClient_Get_getARestorePoint() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcompute.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewRestorePointsClient().Get(ctx, "myResourceGroup", "rpcName", "rpName", &armcompute.RestorePointsClientGetOptions{Expand: nil})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.RestorePoint = armcompute.RestorePoint{
	// 	Name: to.Ptr("rpName"),
	// 	ID: to.Ptr("/subscriptions/{subscription-id}/resourceGroups/myResourceGroup/providers/Microsoft.Compute/restorePointCollections/rpcName/restorePoints/rpName"),
	// 	Properties: &armcompute.RestorePointProperties{
	// 		ConsistencyMode: to.Ptr(armcompute.ConsistencyModeTypesApplicationConsistent),
	// 		ExcludeDisks: []*armcompute.APIEntityReference{
	// 			{
	// 				ID: to.Ptr("/subscriptions/{subscription-id}/resourceGroups/myResourceGroup/providers/Microsoft.Compute/disks/vm8768_disk2_fe6ffde4f69b491ca33fb984d5bcd89f"),
	// 		}},
	// 		InstantAccessDurationMinutes: to.Ptr[int32](120),
	// 		ProvisioningState: to.Ptr("Succeeded"),
	// 		SourceMetadata: &armcompute.RestorePointSourceMetadata{
	// 			DiagnosticsProfile: &armcompute.DiagnosticsProfile{
	// 				BootDiagnostics: &armcompute.BootDiagnostics{
	// 					Enabled: to.Ptr(true),
	// 				},
	// 			},
	// 			HardwareProfile: &armcompute.HardwareProfile{
	// 				VMSize: to.Ptr(armcompute.VirtualMachineSizeTypesStandardB1S),
	// 			},
	// 			Location: to.Ptr("westus"),
	// 			OSProfile: &armcompute.OSProfile{
	// 				AdminUsername: to.Ptr("admin"),
	// 				AllowExtensionOperations: to.Ptr(true),
	// 				ComputerName: to.Ptr("computerName"),
	// 				RequireGuestProvisionSignal: to.Ptr(true),
	// 				Secrets: []*armcompute.VaultSecretGroup{
	// 				},
	// 				WindowsConfiguration: &armcompute.WindowsConfiguration{
	// 					EnableAutomaticUpdates: to.Ptr(true),
	// 					ProvisionVMAgent: to.Ptr(true),
	// 				},
	// 			},
	// 			StorageProfile: &armcompute.RestorePointSourceVMStorageProfile{
	// 				DataDisks: []*armcompute.RestorePointSourceVMDataDisk{
	// 					{
	// 						Name: to.Ptr("testingexcludedisk_DataDisk_1"),
	// 						Caching: to.Ptr(armcompute.CachingTypesNone),
	// 						DiskRestorePoint: &armcompute.DiskRestorePointAttributes{
	// 							ID: to.Ptr("/subscriptions/{subscription-id}/resourceGroups/userdata/providers/Microsoft.Compute/restorePointCollections/mynewrpc/restorePoints/restorepointtwo/diskRestorePoints/testingexcludedisk_DataDisk_1_68785190-1acb-4d5e-a8ae-705b45f3dca5"),
	// 						},
	// 						Lun: to.Ptr[int32](1),
	// 						ManagedDisk: &armcompute.ManagedDiskParameters{
	// 							ID: to.Ptr("/subscriptions/{subscription-id}/resourceGroups/userdata/providers/Microsoft.Compute/disks/testingexcludedisk_DataDisk_1"),
	// 							StorageAccountType: to.Ptr(armcompute.StorageAccountTypesStandardLRS),
	// 						},
	// 				}},
	// 				DiskControllerType: to.Ptr(armcompute.DiskControllerTypesNVMe),
	// 				OSDisk: &armcompute.RestorePointSourceVMOSDisk{
	// 					Name: to.Ptr("testingexcludedisk_OsDisk_1_74cdaedcea50483d9833c96adefa100f"),
	// 					Caching: to.Ptr(armcompute.CachingTypesReadWrite),
	// 					DiskRestorePoint: &armcompute.DiskRestorePointAttributes{
	// 						ID: to.Ptr("/subscriptions/{subscription-id}/resourceGroups/myResourceGroup/providers/Microsoft.Compute/restorePointCollections/rpcName/restorePoints/rpName/diskRestorePoints/testingexcludedisk_OsDisk_1_74cdaedcea50483d9833c96adefa100f_22b4bdfe-6c54-4f72-84d8-85d8860f0c57"),
	// 					},
	// 					ManagedDisk: &armcompute.ManagedDiskParameters{
	// 						ID: to.Ptr("/subscriptions/{subscription-id}/resourceGroups/myResourceGroup/providers/Microsoft.Compute/disks/testingexcludedisk_OsDisk_1_74cdaedcea50483d9833c96adefa100f"),
	// 						StorageAccountType: to.Ptr(armcompute.StorageAccountTypesStandardLRS),
	// 					},
	// 					OSType: to.Ptr(armcompute.OperatingSystemTypeWindows),
	// 				},
	// 			},
	// 			VMID: to.Ptr("76d6541e-80bd-4dc1-932b-3cae4cfb80e7"),
	// 		},
	// 		TimeCreated: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-01-27T20:35:05.840Z"); return t}()),
	// 	},
	// }
}
