package armcompute_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/compute/armcompute/v7"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/ccea12be6cd5e64743499d46f7a9c8b60df52db1/specification/compute/resource-manager/Microsoft.Compute/DiskRP/stable/2025-01-02/examples/snapshotExamples/Snapshot_Update_WithAcceleratedNetwork.json
func ExampleSnapshotsClient_BeginUpdate_updateASnapshotWithAcceleratedNetworking() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcompute.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewSnapshotsClient().BeginUpdate(ctx, "myResourceGroup", "mySnapshot", armcompute.SnapshotUpdate{
		Properties: &armcompute.SnapshotUpdateProperties{
			DiskSizeGB: to.Ptr[int32](20),
			SupportedCapabilities: &armcompute.SupportedCapabilities{
				AcceleratedNetwork: to.Ptr(false),
			},
		},
		Tags: map[string]*string{
			"department": to.Ptr("Development"),
			"project":    to.Ptr("UpdateSnapshots"),
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
	// res.Snapshot = armcompute.Snapshot{
	// 	Name: to.Ptr("mySnapshot"),
	// 	Location: to.Ptr("West US"),
	// 	Tags: map[string]*string{
	// 		"department": to.Ptr("Development"),
	// 		"project": to.Ptr("UpdateSnapshots"),
	// 	},
	// 	Properties: &armcompute.SnapshotProperties{
	// 		CreationData: &armcompute.CreationData{
	// 			CreateOption: to.Ptr(armcompute.DiskCreateOptionCopy),
	// 			SourceResourceID: to.Ptr("subscriptions/{subscription-id}/resourceGroups/myResourceGroup/providers/Microsoft.Compute/snapshots/mySnapshot1"),
	// 		},
	// 		DiskSizeGB: to.Ptr[int32](20),
	// 		ProvisioningState: to.Ptr("Succeeded"),
	// 		SupportedCapabilities: &armcompute.SupportedCapabilities{
	// 			AcceleratedNetwork: to.Ptr(false),
	// 		},
	// 	},
	// }
}
