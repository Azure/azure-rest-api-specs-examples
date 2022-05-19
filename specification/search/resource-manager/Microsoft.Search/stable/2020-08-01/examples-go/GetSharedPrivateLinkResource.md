Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Fsearch%2Farmsearch%2Fv1.0.0/sdk/resourcemanager/search/armsearch/README.md) on how to add the SDK to your project and authenticate.

```go
package armsearch_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/search/armsearch"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/search/resource-manager/Microsoft.Search/stable/2020-08-01/examples/GetSharedPrivateLinkResource.json
func ExampleSharedPrivateLinkResourcesClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armsearch.NewSharedPrivateLinkResourcesClient("subid", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := client.Get(ctx,
		"rg1",
		"mysearchservice",
		"testResource",
		&armsearch.SearchManagementRequestOptions{ClientRequestID: nil},
		nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// TODO: use response item
	_ = res
}
```
