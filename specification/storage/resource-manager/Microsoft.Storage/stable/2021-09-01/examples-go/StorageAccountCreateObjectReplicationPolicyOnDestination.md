Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Fstorage%2Farmstorage%2Fv0.6.0/sdk/resourcemanager/storage/armstorage/README.md) on how to add the SDK to your project and authenticate.

```go
package armstorage_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/storage/armstorage"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/storage/resource-manager/Microsoft.Storage/stable/2021-09-01/examples/StorageAccountCreateObjectReplicationPolicyOnDestination.json
func ExampleObjectReplicationPoliciesClient_CreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
		return
	}
	ctx := context.Background()
	client, err := armstorage.NewObjectReplicationPoliciesClient("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
		return
	}
	res, err := client.CreateOrUpdate(ctx,
		"<resource-group-name>",
		"<account-name>",
		"<object-replication-policy-id>",
		armstorage.ObjectReplicationPolicy{
			Properties: &armstorage.ObjectReplicationPolicyProperties{
				DestinationAccount: to.Ptr("<destination-account>"),
				Rules: []*armstorage.ObjectReplicationPolicyRule{
					{
						DestinationContainer: to.Ptr("<destination-container>"),
						Filters: &armstorage.ObjectReplicationPolicyFilter{
							PrefixMatch: []*string{
								to.Ptr("blobA"),
								to.Ptr("blobB")},
						},
						SourceContainer: to.Ptr("<source-container>"),
					}},
				SourceAccount: to.Ptr("<source-account>"),
			},
		},
		nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
		return
	}
	// TODO: use response item
	_ = res
}
```
