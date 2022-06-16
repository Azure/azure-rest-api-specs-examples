package armstoragepool_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/storagepool/armstoragepool"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/storagepool/resource-manager/Microsoft.StoragePool/stable/2021-08-01/examples/DiskPools_Patch.json
func ExampleDiskPoolsClient_BeginUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armstoragepool.NewDiskPoolsClient("00000000-0000-0000-0000-000000000000", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := client.BeginUpdate(ctx,
		"myResourceGroup",
		"myDiskPool",
		armstoragepool.DiskPoolUpdate{
			Properties: &armstoragepool.DiskPoolUpdateProperties{
				Disks: []*armstoragepool.Disk{
					{
						ID: to.Ptr("/subscriptions/11111111-1111-1111-1111-111111111111/resourceGroups/myResourceGroup/providers/Microsoft.Compute/disks/vm-name_DataDisk_0"),
					},
					{
						ID: to.Ptr("/subscriptions/11111111-1111-1111-1111-111111111111/resourceGroups/myResourceGroup/providers/Microsoft.Compute/disks/vm-name_DataDisk_1"),
					}},
			},
			SKU: &armstoragepool.SKU{
				Name: to.Ptr("Basic_B1"),
				Tier: to.Ptr("Basic"),
			},
			Tags: map[string]*string{
				"key": to.Ptr("value"),
			},
		},
		nil)
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
