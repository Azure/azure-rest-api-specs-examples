Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Fstorage%2Farmstorage%2Fv1.0.0/sdk/resourcemanager/storage/armstorage/README.md) on how to add the SDK to your project and authenticate.

```go
package armstorage_test

import (
	"context"
	"log"

	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/storage/armstorage"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/storage/resource-manager/Microsoft.Storage/stable/2021-09-01/examples/BlobRangesRestore.json
func ExampleAccountsClient_BeginRestoreBlobRanges() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armstorage.NewAccountsClient("{subscription-id}", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := client.BeginRestoreBlobRanges(ctx,
		"res9101",
		"sto4445",
		armstorage.BlobRestoreParameters{
			BlobRanges: []*armstorage.BlobRestoreRange{
				{
					EndRange:   to.Ptr("container/blobpath2"),
					StartRange: to.Ptr("container/blobpath1"),
				},
				{
					EndRange:   to.Ptr(""),
					StartRange: to.Ptr("container2/blobpath3"),
				}},
			TimeToRestore: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2019-04-20T15:30:00.0000000Z"); return t }()),
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
