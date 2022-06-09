```go
package armstorage_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/storage/armstorage"
)

// x-ms-original-file: specification/storage/resource-manager/Microsoft.Storage/stable/2021-08-01/examples/StorageAccountCreateObjectReplicationPolicyOnDestination.json
func ExampleObjectReplicationPoliciesClient_CreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armstorage.NewObjectReplicationPoliciesClient("<subscription-id>", cred, nil)
	res, err := client.CreateOrUpdate(ctx,
		"<resource-group-name>",
		"<account-name>",
		"<object-replication-policy-id>",
		armstorage.ObjectReplicationPolicy{
			Properties: &armstorage.ObjectReplicationPolicyProperties{
				DestinationAccount: to.StringPtr("<destination-account>"),
				Rules: []*armstorage.ObjectReplicationPolicyRule{
					{
						DestinationContainer: to.StringPtr("<destination-container>"),
						Filters: &armstorage.ObjectReplicationPolicyFilter{
							PrefixMatch: []*string{
								to.StringPtr("blobA"),
								to.StringPtr("blobB")},
						},
						SourceContainer: to.StringPtr("<source-container>"),
					}},
				SourceAccount: to.StringPtr("<source-account>"),
			},
		},
		nil)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Response result: %#v\n", res.ObjectReplicationPoliciesClientCreateOrUpdateResult)
}
```

Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Fstorage%2Farmstorage%2Fv0.4.1/sdk/resourcemanager/storage/armstorage/README.md) on how to add the SDK to your project and authenticate.
