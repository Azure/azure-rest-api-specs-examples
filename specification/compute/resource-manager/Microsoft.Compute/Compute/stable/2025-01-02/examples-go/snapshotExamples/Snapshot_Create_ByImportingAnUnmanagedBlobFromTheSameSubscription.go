package armcompute_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/compute/armcompute/v8"
)

// Generated from example definition: 2025-01-02/snapshotExamples/Snapshot_Create_ByImportingAnUnmanagedBlobFromTheSameSubscription.json
func ExampleSnapshotsClient_BeginCreateOrUpdate_createASnapshotByImportingAnUnmanagedBlobFromTheSameSubscription() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcompute.NewClientFactory("{subscription-id}", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewSnapshotsClient().BeginCreateOrUpdate(ctx, "myResourceGroup", "mySnapshot1", armcompute.Snapshot{
		Location: to.Ptr("West US"),
		Properties: &armcompute.SnapshotProperties{
			CreationData: &armcompute.CreationData{
				CreateOption: to.Ptr(armcompute.DiskCreateOptionImport),
				SourceURI:    to.Ptr("https://mystorageaccount.blob.core.windows.net/osimages/osimage.vhd"),
			},
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	res, err := poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to poll the result: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armcompute.SnapshotsClientCreateOrUpdateResponse{
	// 	Snapshot: armcompute.Snapshot{
	// 		ID: to.Ptr("/subscriptions/{subscription-id}/resourceGroups/myResourceGroup/providers/Microsoft.Compute/snapshots/mySnapshot"),
	// 		Properties: &armcompute.SnapshotProperties{
	// 			CreationData: &armcompute.CreationData{
	// 				CreateOption: to.Ptr(armcompute.DiskCreateOptionImport),
	// 				SourceURI: to.Ptr("https://mystorageaccount.blob.core.windows.net/osimages/osimage.vhd"),
	// 			},
	// 			ProvisioningState: to.Ptr("Updating"),
	// 		},
	// 		Location: to.Ptr("West US"),
	// 		Name: to.Ptr("mySnapshot1"),
	// 	},
	// }
}
