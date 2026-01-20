package armcompute_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/compute/armcompute/v7"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/ccea12be6cd5e64743499d46f7a9c8b60df52db1/specification/compute/resource-manager/Microsoft.Compute/DiskRP/stable/2025-01-02/examples/diskExamples/Disk_Create_WithUltraSSD_ReadOnly.json
func ExampleDisksClient_BeginCreateOrUpdate_createAManagedDiskWithUltraAccountTypeWithReadOnlyPropertySet() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcompute.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewDisksClient().BeginCreateOrUpdate(ctx, "myResourceGroup", "myUltraReadOnlyDisk", armcompute.Disk{
		Location: to.Ptr("West US"),
		Properties: &armcompute.DiskProperties{
			CreationData: &armcompute.CreationData{
				CreateOption:      to.Ptr(armcompute.DiskCreateOptionEmpty),
				LogicalSectorSize: to.Ptr[int32](4096),
			},
			DiskIOPSReadWrite: to.Ptr[int64](125),
			DiskMBpsReadWrite: to.Ptr[int64](3000),
			DiskSizeGB:        to.Ptr[int32](200),
			Encryption: &armcompute.Encryption{
				Type: to.Ptr(armcompute.EncryptionTypeEncryptionAtRestWithPlatformKey),
			},
		},
		SKU: &armcompute.DiskSKU{
			Name: to.Ptr(armcompute.DiskStorageAccountTypesUltraSSDLRS),
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	res, err := poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.Disk = armcompute.Disk{
	// 	Name: to.Ptr("myUltraReadOnlyDisk"),
	// 	ID: to.Ptr("/subscriptions/{subscriptionId}/resourceGroups/myResourceGroup/providers/Microsoft.Compute/disks/myDisk"),
	// 	Location: to.Ptr("West US"),
	// 	Properties: &armcompute.DiskProperties{
	// 		CreationData: &armcompute.CreationData{
	// 			CreateOption: to.Ptr(armcompute.DiskCreateOptionEmpty),
	// 		},
	// 		DiskSizeGB: to.Ptr[int32](200),
	// 		Encryption: &armcompute.Encryption{
	// 			Type: to.Ptr(armcompute.EncryptionTypeEncryptionAtRestWithPlatformKey),
	// 		},
	// 		ProvisioningState: to.Ptr("Succeeded"),
	// 	},
	// 	SKU: &armcompute.DiskSKU{
	// 		Name: to.Ptr(armcompute.DiskStorageAccountTypesUltraSSDLRS),
	// 		Tier: to.Ptr("Ultra"),
	// 	},
	// }
}
