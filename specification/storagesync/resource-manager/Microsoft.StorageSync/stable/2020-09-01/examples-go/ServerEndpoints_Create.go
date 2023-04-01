package armstoragesync_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/storagesync/armstoragesync"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/d55b8005f05b040b852c15e74a0f3e36494a15e1/specification/storagesync/resource-manager/Microsoft.StorageSync/stable/2020-09-01/examples/ServerEndpoints_Create.json
func ExampleServerEndpointsClient_BeginCreate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armstoragesync.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewServerEndpointsClient().BeginCreate(ctx, "SampleResourceGroup_1", "SampleStorageSyncService_1", "SampleSyncGroup_1", "SampleServerEndpoint_1", armstoragesync.ServerEndpointCreateParameters{
		Properties: &armstoragesync.ServerEndpointCreateParametersProperties{
			CloudTiering:                 to.Ptr(armstoragesync.FeatureStatusOff),
			InitialDownloadPolicy:        to.Ptr(armstoragesync.InitialDownloadPolicyNamespaceThenModifiedFiles),
			InitialUploadPolicy:          to.Ptr(armstoragesync.InitialUploadPolicyServerAuthoritative),
			LocalCacheMode:               to.Ptr(armstoragesync.LocalCacheModeUpdateLocallyCachedFiles),
			OfflineDataTransfer:          to.Ptr(armstoragesync.FeatureStatusOn),
			OfflineDataTransferShareName: to.Ptr("myfileshare"),
			ServerLocalPath:              to.Ptr("D:\\SampleServerEndpoint_1"),
			ServerResourceID:             to.Ptr("/subscriptions/52b8da2f-61e0-4a1f-8dde-336911f367fb/resourceGroups/SampleResourceGroup_1/providers/Microsoft.StorageSync/storageSyncServices/SampleStorageSyncService_1/registeredServers/080d4133-bdb5-40a0-96a0-71a6057bfe9a"),
			TierFilesOlderThanDays:       to.Ptr[int32](0),
			VolumeFreeSpacePercent:       to.Ptr[int32](100),
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
	// res.ServerEndpoint = armstoragesync.ServerEndpoint{
	// 	Name: to.Ptr("SampleServerEndpoint_1"),
	// 	Type: to.Ptr("Microsoft.StorageSync/storageSyncServices/syncGroups/serverEndpoints"),
	// 	ID: to.Ptr("/subscriptions/52b8da2f-61e0-4a1f-8dde-336911f367fb/resourceGroups/SampleResourceGroup_1/providers/Microsoft.StorageSync/storageSyncServices/SampleStorageSyncService_1/syncGroups/SampleSyncGroup_1/serverEndpoints/SampleServerEndpoint_1"),
	// 	Properties: &armstoragesync.ServerEndpointProperties{
	// 		CloudTiering: to.Ptr(armstoragesync.FeatureStatusOff),
	// 		FriendlyName: to.Ptr("somemachine.redmond.corp.microsoft.com"),
	// 		InitialDownloadPolicy: to.Ptr(armstoragesync.InitialDownloadPolicyNamespaceThenModifiedFiles),
	// 		InitialUploadPolicy: to.Ptr(armstoragesync.InitialUploadPolicyServerAuthoritative),
	// 		LastOperationName: to.Ptr("ICreateServerEndpointWorkflow"),
	// 		LastWorkflowID: to.Ptr("storageSyncServices/healthDemo1/workflows/569afb5c-7172-4cf8-a8ea-d987f727f11a"),
	// 		LocalCacheMode: to.Ptr(armstoragesync.LocalCacheModeUpdateLocallyCachedFiles),
	// 		OfflineDataTransfer: to.Ptr(armstoragesync.FeatureStatusOn),
	// 		OfflineDataTransferShareName: to.Ptr("myfileshare"),
	// 		OfflineDataTransferStorageAccountResourceID: to.Ptr("/subscriptions/744f4d70-6d17-4921-8970-a765d14f763f/resourceGroups/myRG/providers/Microsoft.Storage/storageAccounts/mysa"),
	// 		OfflineDataTransferStorageAccountTenantID: to.Ptr("72f988bf-86f1-41af-91ab-2d7cd011db47"),
	// 		ProvisioningState: to.Ptr("Succeeded"),
	// 		ServerLocalPath: to.Ptr("D:\\SampleServerEndpoint_1"),
	// 		ServerName: to.Ptr("somemachine.redmond.corp.microsoft.com"),
	// 		ServerResourceID: to.Ptr("/subscriptions/52b8da2f-61e0-4a1f-8dde-336911f367fb/resourceGroups/SampleResourceGroup_1/providers/Microsoft.StorageSync/storageSyncServices/SampleStorageSyncService_1/registeredServers/080d4133-bdb5-40a0-96a0-71a6057bfe9a"),
	// 		TierFilesOlderThanDays: to.Ptr[int32](0),
	// 		VolumeFreeSpacePercent: to.Ptr[int32](100),
	// 	},
	// }
}
