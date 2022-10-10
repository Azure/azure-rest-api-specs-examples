package armcompute_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/compute/armcompute/v4"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/compute/resource-manager/Microsoft.Compute/DiskRP/stable/2022-07-02/examples/diskExamples/Disk_Create_WithUltraSSD_ReadOnly.json
func ExampleDisksClient_BeginCreateOrUpdate_createAManagedDiskWithUltraAccountTypeWithReadOnlyPropertySet() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armcompute.NewDisksClient("{subscription-id}", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := client.BeginCreateOrUpdate(ctx, "myResourceGroup", "myUltraReadOnlyDisk", armcompute.Disk{
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
	// TODO: use response item
	_ = res
}
