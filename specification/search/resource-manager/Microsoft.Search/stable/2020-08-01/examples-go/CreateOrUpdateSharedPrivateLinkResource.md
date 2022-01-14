Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Fsearch%2Farmsearch%2Fv0.3.0/sdk/resourcemanager/search/armsearch/README.md) on how to add the SDK to your project and authenticate.

```go
package armsearch_test

import (
	"context"
	"log"

	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/search/armsearch"
)

// x-ms-original-file: specification/search/resource-manager/Microsoft.Search/stable/2020-08-01/examples/CreateOrUpdateSharedPrivateLinkResource.json
func ExampleSharedPrivateLinkResourcesClient_BeginCreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armsearch.NewSharedPrivateLinkResourcesClient("<subscription-id>", cred, nil)
	poller, err := client.BeginCreateOrUpdate(ctx,
		"<resource-group-name>",
		"<search-service-name>",
		"<shared-private-link-resource-name>",
		armsearch.SharedPrivateLinkResource{
			Properties: &armsearch.SharedPrivateLinkResourceProperties{
				GroupID:               to.StringPtr("<group-id>"),
				PrivateLinkResourceID: to.StringPtr("<private-link-resource-id>"),
				RequestMessage:        to.StringPtr("<request-message>"),
			},
		},
		&armsearch.SearchManagementRequestOptions{ClientRequestID: nil})
	if err != nil {
		log.Fatal(err)
	}
	res, err := poller.PollUntilDone(ctx, 30*time.Second)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Response result: %#v\n", res.SharedPrivateLinkResourcesClientCreateOrUpdateResult)
}
```
