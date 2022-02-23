Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Fstorage%2Farmstorage%2Fv0.4.1/sdk/resourcemanager/storage/armstorage/README.md) on how to add the SDK to your project and authenticate.

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

// x-ms-original-file: specification/storage/resource-manager/Microsoft.Storage/stable/2021-08-01/examples/FileShareAclsPatch.json
func ExampleFileSharesClient_Update() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armstorage.NewFileSharesClient("<subscription-id>", cred, nil)
	res, err := client.Update(ctx,
		"<resource-group-name>",
		"<account-name>",
		"<share-name>",
		armstorage.FileShare{
			FileShareProperties: &armstorage.FileShareProperties{
				SignedIdentifiers: []*armstorage.SignedIdentifier{
					{
						AccessPolicy: &armstorage.AccessPolicy{
							ExpiryTime: to.TimePtr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-05-01T08:49:37.0000000Z"); return t }()),
							Permission: to.StringPtr("<permission>"),
							StartTime:  to.TimePtr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-04-01T08:49:37.0000000Z"); return t }()),
						},
						ID: to.StringPtr("<id>"),
					}},
			},
		},
		nil)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Response result: %#v\n", res.FileSharesClientUpdateResult)
}
```
