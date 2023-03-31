package armstoragesync_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/storagesync/armstoragesync"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/d55b8005f05b040b852c15e74a0f3e36494a15e1/specification/storagesync/resource-manager/Microsoft.StorageSync/stable/2020-09-01/examples/SyncGroups_ListByStorageSyncService.json
func ExampleSyncGroupsClient_NewListByStorageSyncServicePager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armstoragesync.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewSyncGroupsClient().NewListByStorageSyncServicePager("SampleResourceGroup_1", "SampleStorageSyncService_1", nil)
	for pager.More() {
		page, err := pager.NextPage(ctx)
		if err != nil {
			log.Fatalf("failed to advance page: %v", err)
		}
		for _, v := range page.Value {
			// You could use page here. We use blank identifier for just demo purposes.
			_ = v
		}
		// If the HTTP response code is 200 as defined in example definition, your page structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
		// page.SyncGroupArray = armstoragesync.SyncGroupArray{
		// 	Value: []*armstoragesync.SyncGroup{
		// 		{
		// 			Name: to.Ptr("SampleSyncGroup_1"),
		// 			Type: to.Ptr("Microsoft.StorageSync/storageSyncServices/syncGroups"),
		// 			ID: to.Ptr("/subscriptions/3a048283-338f-4002-a9dd-a50fdadcb392/resourceGroups/SampleResourceGroup_1/providers/Microsoft.StorageSync/storageSyncServices/SSS_Restore_08-08_Test112/syncGroups/SampleSyncGroup_1"),
		// 			Properties: &armstoragesync.SyncGroupProperties{
		// 				SyncGroupStatus: to.Ptr("0"),
		// 				UniqueID: to.Ptr("191660cd-6a1a-4f8c-9787-a6bed206a1dd"),
		// 			},
		// 	}},
		// }
	}
}
