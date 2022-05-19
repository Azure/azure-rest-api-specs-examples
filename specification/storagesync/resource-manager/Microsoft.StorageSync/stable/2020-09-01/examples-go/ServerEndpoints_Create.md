Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Fstoragesync%2Farmstoragesync%2Fv1.0.0/sdk/resourcemanager/storagesync/armstoragesync/README.md) on how to add the SDK to your project and authenticate.

```go
package armstoragesync_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/storagesync/armstoragesync"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/storagesync/resource-manager/Microsoft.StorageSync/stable/2020-09-01/examples/ServerEndpoints_Create.json
func ExampleServerEndpointsClient_BeginCreate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armstoragesync.NewServerEndpointsClient("52b8da2f-61e0-4a1f-8dde-336911f367fb", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := client.BeginCreate(ctx,
		"SampleResourceGroup_1",
		"SampleStorageSyncService_1",
		"SampleSyncGroup_1",
		"SampleServerEndpoint_1",
		armstoragesync.ServerEndpointCreateParameters{
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
```
