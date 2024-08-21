package armcompute_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/compute/armcompute/v6"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/c7b98b36e4023331545051284d8500adf98f02fe/specification/compute/resource-manager/Microsoft.Compute/DiskRP/stable/2024-03-02/examples/diskExamples/Disk_CreateOrUpdate_BurstingEnabled.json
func ExampleDisksClient_BeginUpdate_createOrUpdateABurstingEnabledManagedDisk() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcompute.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewDisksClient().BeginUpdate(ctx, "myResourceGroup", "myDisk", armcompute.DiskUpdate{
		Properties: &armcompute.DiskUpdateProperties{
			BurstingEnabled: to.Ptr(true),
			DiskSizeGB:      to.Ptr[int32](1024),
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
	// 	Name: to.Ptr("myDisk"),
	// 	ID: to.Ptr("/subscriptions/{subscriptionId}/resourceGroups/myResourceGroup/providers/Microsoft.Compute/disks/myDisk"),
	// 	Location: to.Ptr("West US"),
	// 	Properties: &armcompute.DiskProperties{
	// 		BurstingEnabled: to.Ptr(true),
	// 		CreationData: &armcompute.CreationData{
	// 			CreateOption: to.Ptr(armcompute.DiskCreateOptionEmpty),
	// 		},
	// 		DiskSizeGB: to.Ptr[int32](1024),
	// 		ProvisioningState: to.Ptr("Succeeded"),
	// 	},
	// }
}
