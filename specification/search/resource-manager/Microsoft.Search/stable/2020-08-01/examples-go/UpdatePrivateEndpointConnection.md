```go
package armsearch_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/search/armsearch"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/search/resource-manager/Microsoft.Search/stable/2020-08-01/examples/UpdatePrivateEndpointConnection.json
func ExamplePrivateEndpointConnectionsClient_Update() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armsearch.NewPrivateEndpointConnectionsClient("subid", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := client.Update(ctx,
		"rg1",
		"mysearchservice",
		"testEndpoint.50bf4fbe-d7c1-4b48-a642-4f5892642546",
		armsearch.PrivateEndpointConnection{
			Properties: &armsearch.PrivateEndpointConnectionProperties{
				PrivateLinkServiceConnectionState: &armsearch.PrivateEndpointConnectionPropertiesPrivateLinkServiceConnectionState{
					Description: to.Ptr("Rejected for some reason"),
					Status:      to.Ptr(armsearch.PrivateLinkServiceConnectionStatusRejected),
				},
			},
		},
		&armsearch.SearchManagementRequestOptions{ClientRequestID: nil},
		nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// TODO: use response item
	_ = res
}
```

Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Fsearch%2Farmsearch%2Fv1.0.0/sdk/resourcemanager/search/armsearch/README.md) on how to add the SDK to your project and authenticate.
