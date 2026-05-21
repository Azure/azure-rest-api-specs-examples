package armstoragesync_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/storagesync/armstoragesync/v2"
)

// Generated from example definition: 2022-09-01/SyncGroups_Create.json
func ExampleSyncGroupsClient_Create() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armstoragesync.NewClientFactory("52b8da2f-61e0-4a1f-8dde-336911f367fb", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewSyncGroupsClient().Create(ctx, "SampleResourceGroup_1", "SampleStorageSyncService_1", "SampleSyncGroup_1", armstoragesync.SyncGroupCreateParameters{
		Properties: map[string]any{},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armstoragesync.SyncGroupsClientCreateResponse{
	// 	SyncGroup: armstoragesync.SyncGroup{
	// 		Name: to.Ptr("SampleSyncGroup_1"),
	// 		Type: to.Ptr("Microsoft.StorageSync/storageSyncServices/syncGroups"),
	// 		ID: to.Ptr("/subscriptions/3a048283-338f-4002-a9dd-a50fdadcb392/resourceGroups/SampleResourceGroup_1/providers/Microsoft.StorageSync/storageSyncServices/SampleStorageSyncService_1/syncGroups/SampleSyncGroup_1"),
	// 		Properties: &armstoragesync.SyncGroupProperties{
	// 			UniqueID: to.Ptr("7868e4ee-8ddd-4a2d-941b-0041f6052a8a"),
	// 		},
	// 	},
	// }
}
