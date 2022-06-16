package armstoragepool_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/storagepool/armstoragepool"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/storagepool/resource-manager/Microsoft.StoragePool/stable/2021-08-01/examples/IscsiTargets_Patch.json
func ExampleIscsiTargetsClient_BeginUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armstoragepool.NewIscsiTargetsClient("00000000-0000-0000-0000-000000000000", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := client.BeginUpdate(ctx,
		"myResourceGroup",
		"myDiskPool",
		"myIscsiTarget",
		armstoragepool.IscsiTargetUpdate{
			Properties: &armstoragepool.IscsiTargetUpdateProperties{
				Luns: []*armstoragepool.IscsiLun{
					{
						Name:                       to.Ptr("lun0"),
						ManagedDiskAzureResourceID: to.Ptr("/subscriptions/11111111-1111-1111-1111-111111111111/resourceGroups/myResourceGroup/providers/Microsoft.Compute/disks/vm-name_DataDisk_1"),
					}},
				StaticACLs: []*armstoragepool.ACL{
					{
						InitiatorIqn: to.Ptr("iqn.2005-03.org.iscsi:client"),
						MappedLuns: []*string{
							to.Ptr("lun0")},
					}},
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
