package armstoragepool_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/storagepool/armstoragepool"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/d55b8005f05b040b852c15e74a0f3e36494a15e1/specification/storagepool/resource-manager/Microsoft.StoragePool/stable/2021-08-01/examples/DiskPools_Put.json
func ExampleDiskPoolsClient_BeginCreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armstoragepool.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewDiskPoolsClient().BeginCreateOrUpdate(ctx, "myResourceGroup", "myDiskPool", armstoragepool.DiskPoolCreate{
		Location: to.Ptr("westus"),
		Properties: &armstoragepool.DiskPoolCreateProperties{
			AvailabilityZones: []*string{
				to.Ptr("1")},
			Disks: []*armstoragepool.Disk{
				{
					ID: to.Ptr("/subscriptions/11111111-1111-1111-1111-111111111111/resourceGroups/myResourceGroup/providers/Microsoft.Compute/disks/vm-name_DataDisk_0"),
				},
				{
					ID: to.Ptr("/subscriptions/11111111-1111-1111-1111-111111111111/resourceGroups/myResourceGroup/providers/Microsoft.Compute/disks/vm-name_DataDisk_1"),
				}},
			SubnetID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/myResourceGroup/providers/Microsoft.Network/virtualNetworks/myvnet/subnets/mysubnet"),
		},
		SKU: &armstoragepool.SKU{
			Name: to.Ptr("Basic_V1"),
			Tier: to.Ptr("Basic"),
		},
		Tags: map[string]*string{
			"key": to.Ptr("value"),
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
	// res.DiskPool = armstoragepool.DiskPool{
	// 	Name: to.Ptr("myDiskPool"),
	// 	Type: to.Ptr("Microsoft.StoragePool/diskPools"),
	// 	ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/myResourceGroup/providers/Microsoft.StoragePool/diskPools/myDiskPool"),
	// 	Location: to.Ptr("westus2"),
	// 	Tags: map[string]*string{
	// 		"env": to.Ptr("int"),
	// 	},
	// 	ManagedBy: to.Ptr("/subscriptions/11111111-1111-1111-1111-111111111111/resourceGroups/myResourceGroup/providers/Microsoft.AVS/privateClouds/myPrivateCloud/clusters/Cluster-1/datastores/datastore1"),
	// 	ManagedByExtended: []*string{
	// 		to.Ptr("/subscriptions/11111111-1111-1111-1111-111111111111/resourceGroups/myResourceGroup/providers/Microsoft.AVS/privateClouds/myPrivateCloud/clusters/Cluster-1/datastores/datastore1"),
	// 		to.Ptr("/subscriptions/11111111-1111-1111-1111-111111111111/resourceGroups/myResourceGroup/providers/Microsoft.AVS/privateClouds/myPrivateCloud/clusters/Cluster-1/datastores/datastore2")},
	// 		Properties: &armstoragepool.DiskPoolProperties{
	// 			AvailabilityZones: []*string{
	// 				to.Ptr("1")},
	// 				Disks: []*armstoragepool.Disk{
	// 					{
	// 						ID: to.Ptr("/subscriptions/11111111-1111-1111-1111-111111111111/resourceGroups/myResourceGroup/providers/Microsoft.Compute/disks/vm-name_DataDisk_0"),
	// 					},
	// 					{
	// 						ID: to.Ptr("/subscriptions/11111111-1111-1111-1111-111111111111/resourceGroups/myResourceGroup/providers/Microsoft.Compute/disks/vm-name_DataDisk_1"),
	// 				}},
	// 				ProvisioningState: to.Ptr(armstoragepool.ProvisioningStatesSucceeded),
	// 				Status: to.Ptr(armstoragepool.OperationalStatusUnknown),
	// 				SubnetID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/myResourceGroup/providers/Microsoft.Network/virtualNetworks/myvnet/subnets/mysubnet"),
	// 			},
	// 			SKU: &armstoragepool.SKU{
	// 				Name: to.Ptr("Basic_0"),
	// 				Tier: to.Ptr("Basic"),
	// 			},
	// 			SystemData: &armstoragepool.SystemMetadata{
	// 				CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-06-24T06:53:57.000Z"); return t}()),
	// 				CreatedBy: to.Ptr("alias"),
	// 				CreatedByType: to.Ptr(armstoragepool.CreatedByTypeUser),
	// 				LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-06-24T06:53:57.000Z"); return t}()),
	// 				LastModifiedBy: to.Ptr("alias"),
	// 				LastModifiedByType: to.Ptr(armstoragepool.CreatedByTypeUser),
	// 			},
	// 		}
}
