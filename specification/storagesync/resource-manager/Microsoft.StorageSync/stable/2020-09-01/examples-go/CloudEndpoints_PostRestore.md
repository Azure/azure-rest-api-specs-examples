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

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/storagesync/resource-manager/Microsoft.StorageSync/stable/2020-09-01/examples/CloudEndpoints_PostRestore.json
func ExampleCloudEndpointsClient_BeginPostRestore() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armstoragesync.NewCloudEndpointsClient("52b8da2f-61e0-4a1f-8dde-336911f367fb", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := client.BeginPostRestore(ctx,
		"SampleResourceGroup_1",
		"SampleStorageSyncService_1",
		"SampleSyncGroup_1",
		"SampleCloudEndpoint_1",
		armstoragesync.PostRestoreRequest{
			AzureFileShareURI: to.Ptr("https://hfsazbackupdevintncus2.file.core.test-cint.azure-test.net/sampleFileShare"),
			RestoreFileSpec: []*armstoragesync.RestoreFileSpec{
				{
					Path:  to.Ptr("text1.txt"),
					Isdir: to.Ptr(false),
				},
				{
					Path:  to.Ptr("MyDir"),
					Isdir: to.Ptr(true),
				},
				{
					Path:  to.Ptr("MyDir/SubDir"),
					Isdir: to.Ptr(false),
				},
				{
					Path:  to.Ptr("MyDir/SubDir/File1.pdf"),
					Isdir: to.Ptr(false),
				}},
			SourceAzureFileShareURI: to.Ptr("https://hfsazbackupdevintncus2.file.core.test-cint.azure-test.net/sampleFileShare"),
			Status:                  to.Ptr("Succeeded"),
		},
		nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	_, err = poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
}
```
