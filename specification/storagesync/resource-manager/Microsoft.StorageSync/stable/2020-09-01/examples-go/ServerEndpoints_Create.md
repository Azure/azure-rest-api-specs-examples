Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Fstoragesync%2Farmstoragesync%2Fv0.4.0/sdk/resourcemanager/storagesync/armstoragesync/README.md) on how to add the SDK to your project and authenticate.

```go
package armstoragesync_test

import (
	"context"
	"log"

	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/storagesync/armstoragesync"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/storagesync/resource-manager/Microsoft.StorageSync/stable/2020-09-01/examples/ServerEndpoints_Create.json
func ExampleServerEndpointsClient_BeginCreate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
		return
	}
	ctx := context.Background()
	client, err := armstoragesync.NewServerEndpointsClient("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
		return
	}
	poller, err := client.BeginCreate(ctx,
		"<resource-group-name>",
		"<storage-sync-service-name>",
		"<sync-group-name>",
		"<server-endpoint-name>",
		armstoragesync.ServerEndpointCreateParameters{
			Properties: &armstoragesync.ServerEndpointCreateParametersProperties{
				CloudTiering:                 to.Ptr(armstoragesync.FeatureStatusOff),
				InitialDownloadPolicy:        to.Ptr(armstoragesync.InitialDownloadPolicyNamespaceThenModifiedFiles),
				InitialUploadPolicy:          to.Ptr(armstoragesync.InitialUploadPolicyServerAuthoritative),
				LocalCacheMode:               to.Ptr(armstoragesync.LocalCacheModeUpdateLocallyCachedFiles),
				OfflineDataTransfer:          to.Ptr(armstoragesync.FeatureStatusOn),
				OfflineDataTransferShareName: to.Ptr("<offline-data-transfer-share-name>"),
				ServerLocalPath:              to.Ptr("<server-local-path>"),
				ServerResourceID:             to.Ptr("<server-resource-id>"),
				TierFilesOlderThanDays:       to.Ptr[int32](0),
				VolumeFreeSpacePercent:       to.Ptr[int32](100),
			},
		},
		&armstoragesync.ServerEndpointsClientBeginCreateOptions{ResumeToken: ""})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
		return
	}
	res, err := poller.PollUntilDone(ctx, 30*time.Second)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
		return
	}
	// TODO: use response item
	_ = res
}
```
