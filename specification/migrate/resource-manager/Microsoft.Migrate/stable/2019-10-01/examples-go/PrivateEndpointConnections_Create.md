```go
package armmigrate_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/migrate/armmigrate"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/migrate/resource-manager/Microsoft.Migrate/stable/2019-10-01/examples/PrivateEndpointConnections_Create.json
func ExamplePrivateEndpointConnectionClient_Update() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armmigrate.NewPrivateEndpointConnectionClient("6393a73f-8d55-47ef-b6dd-179b3e0c7910", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := client.Update(ctx,
		"abgoyal-westEurope",
		"abgoyalWEselfhostb72bproject",
		"custestpece80project3980pe.7e35576b-3df4-478e-9759-f64351cf4f43",
		&armmigrate.PrivateEndpointConnectionClientUpdateOptions{PrivateEndpointConnectionBody: &armmigrate.PrivateEndpointConnection{
			ETag: to.Ptr("\"00009300-0000-0300-0000-602b967b0000\""),
			Properties: &armmigrate.PrivateEndpointConnectionProperties{
				PrivateLinkServiceConnectionState: &armmigrate.PrivateLinkServiceConnectionState{
					ActionsRequired: to.Ptr(""),
					Status:          to.Ptr(armmigrate.PrivateLinkServiceConnectionStateStatusApproved),
				},
			},
		},
		})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// TODO: use response item
	_ = res
}
```

Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Fmigrate%2Farmmigrate%2Fv1.0.0/sdk/resourcemanager/migrate/armmigrate/README.md) on how to add the SDK to your project and authenticate.
