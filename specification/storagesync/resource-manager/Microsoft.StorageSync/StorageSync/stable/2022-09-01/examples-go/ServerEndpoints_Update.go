package armstoragesync_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/storagesync/armstoragesync/v2"
)

// Generated from example definition: 2022-09-01/ServerEndpoints_Update.json
func ExampleServerEndpointsClient_BeginUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armstoragesync.NewClientFactory("52b8da2f-61e0-4a1f-8dde-336911f367fb", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewServerEndpointsClient().BeginUpdate(ctx, "SampleResourceGroup_1", "SampleStorageSyncService_1", "SampleSyncGroup_1", "SampleServerEndpoint_1", armstoragesync.ServerEndpointUpdateParameters{
		Properties: &armstoragesync.ServerEndpointUpdateProperties{
			CloudTiering:           to.Ptr(armstoragesync.FeatureStatusOff),
			LocalCacheMode:         to.Ptr(armstoragesync.LocalCacheModeUpdateLocallyCachedFiles),
			OfflineDataTransfer:    to.Ptr(armstoragesync.FeatureStatusOff),
			TierFilesOlderThanDays: to.Ptr[int32](0),
			VolumeFreeSpacePercent: to.Ptr[int32](100),
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
	// res = armstoragesync.ServerEndpointsClientUpdateResponse{
	// 	ServerEndpoint: armstoragesync.ServerEndpoint{
	// 		Name: to.Ptr("SampleServerEndpoint_1"),
	// 		Type: to.Ptr("Microsoft.StorageSync/storageSyncServices/syncGroups/serverEndpoints"),
	// 		ID: to.Ptr("/subscriptions/52b8da2f-61e0-4a1f-8dde-336911f367fb/resourceGroups/SampleResourceGroup_1/providers/Microsoft.StorageSync/storageSyncServices/SampleStorageSyncService_1/syncGroups/SampleSyncGroup_1/serverEndpoints/SampleServerEndpoint_1"),
	// 		Properties: &armstoragesync.ServerEndpointProperties{
	// 			CloudTiering: to.Ptr(armstoragesync.FeatureStatusOff),
	// 			FriendlyName: to.Ptr("somemachine2.redmond.corp.microsoft.com"),
	// 			InitialDownloadPolicy: to.Ptr(armstoragesync.InitialDownloadPolicyNamespaceThenModifiedFiles),
	// 			InitialUploadPolicy: to.Ptr(armstoragesync.InitialUploadPolicyMerge),
	// 			LastOperationName: to.Ptr("ICreateServerEndpointWorkflow"),
	// 			LastWorkflowID: to.Ptr("storageSyncServices/healthDemo1/workflows/569afb5c-7172-4cf8-a8ea-d987f727f11b"),
	// 			LocalCacheMode: to.Ptr(armstoragesync.LocalCacheModeUpdateLocallyCachedFiles),
	// 			OfflineDataTransfer: to.Ptr(armstoragesync.FeatureStatusOff),
	// 			OfflineDataTransferShareName: to.Ptr("myfileshare"),
	// 			OfflineDataTransferStorageAccountResourceID: to.Ptr("/subscriptions/744f4d70-6d17-4921-8970-a765d14f763f/resourceGroups/myRG/providers/Microsoft.Storage/storageAccounts/mysa"),
	// 			OfflineDataTransferStorageAccountTenantID: to.Ptr("72f988bf-86f1-41af-91ab-2d7cd011db47"),
	// 			ProvisioningState: to.Ptr("Succeeded"),
	// 			ServerLocalPath: to.Ptr("C:\\data_quota2"),
	// 			ServerName: to.Ptr("somemachine2.redmond.corp.microsoft.com"),
	// 			ServerResourceID: to.Ptr("/subscriptions/3a048283-338f-4002-a9dd-a50fdadcb392/resourceGroups/anpintDemoRG/providers/kailanitest07.one.microsoft.com/storageSyncServices/healthdemo1/registeredServers/f94e2944-b48d-4e5b-bdc7-c48ab3712659"),
	// 			SyncStatus: &armstoragesync.ServerEndpointSyncStatus{
	// 				CombinedHealth: to.Ptr(armstoragesync.ServerEndpointHealthStateError),
	// 				DownloadHealth: to.Ptr(armstoragesync.ServerEndpointHealthStateHealthy),
	// 				DownloadStatus: &armstoragesync.ServerEndpointSyncSessionStatus{
	// 					LastSyncPerItemErrorCount: to.Ptr[int64](0),
	// 					LastSyncResult: to.Ptr[int32](0),
	// 					LastSyncSuccessTimestamp: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-06-11T23:28:33.9217334Z"); return t}()),
	// 					LastSyncTimestamp: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-06-11T23:28:33.9217334Z"); return t}()),
	// 				},
	// 				LastUpdatedTimestamp: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-06-11T23:32:51.1057915Z"); return t}()),
	// 				OfflineDataTransferStatus: to.Ptr(armstoragesync.ServerEndpointOfflineDataTransferStateStopping),
	// 				UploadHealth: to.Ptr(armstoragesync.ServerEndpointHealthStateError),
	// 				UploadStatus: &armstoragesync.ServerEndpointSyncSessionStatus{
	// 					LastSyncPerItemErrorCount: to.Ptr[int64](0),
	// 					LastSyncResult: to.Ptr[int32](-2134351810),
	// 					LastSyncTimestamp: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-06-11T23:32:51.1057915Z"); return t}()),
	// 				},
	// 			},
	// 			TierFilesOlderThanDays: to.Ptr[int32](0),
	// 			VolumeFreeSpacePercent: to.Ptr[int32](20),
	// 		},
	// 	},
	// }
}
