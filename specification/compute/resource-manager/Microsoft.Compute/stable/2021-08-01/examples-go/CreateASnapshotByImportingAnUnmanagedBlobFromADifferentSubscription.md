Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Fcompute%2Farmcompute%2Fv0.4.0/sdk/resourcemanager/compute/armcompute/README.md) on how to add the SDK to your project and authenticate.

```go
package armcompute_test

import (
	"context"
	"log"

	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/compute/armcompute"
)

// x-ms-original-file: specification/compute/resource-manager/Microsoft.Compute/stable/2021-08-01/examples/CreateASnapshotByImportingAnUnmanagedBlobFromADifferentSubscription.json
func ExampleSnapshotsClient_BeginCreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armcompute.NewSnapshotsClient("<subscription-id>", cred, nil)
	poller, err := client.BeginCreateOrUpdate(ctx,
		"<resource-group-name>",
		"<snapshot-name>",
		armcompute.Snapshot{
			Location: to.StringPtr("<location>"),
			Properties: &armcompute.SnapshotProperties{
				CreationData: &armcompute.CreationData{
					CreateOption:     armcompute.DiskCreateOption("Import").ToPtr(),
					SourceURI:        to.StringPtr("<source-uri>"),
					StorageAccountID: to.StringPtr("<storage-account-id>"),
				},
			},
		},
		nil)
	if err != nil {
		log.Fatal(err)
	}
	res, err := poller.PollUntilDone(ctx, 30*time.Second)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Response result: %#v\n", res.SnapshotsClientCreateOrUpdateResult)
}
```
