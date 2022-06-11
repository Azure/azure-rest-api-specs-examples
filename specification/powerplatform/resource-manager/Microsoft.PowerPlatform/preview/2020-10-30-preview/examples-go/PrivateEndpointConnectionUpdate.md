```go
package armpowerplatform_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/powerplatform/armpowerplatform"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/powerplatform/resource-manager/Microsoft.PowerPlatform/preview/2020-10-30-preview/examples/PrivateEndpointConnectionUpdate.json
func ExamplePrivateEndpointConnectionsClient_BeginCreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armpowerplatform.NewPrivateEndpointConnectionsClient("00000000-1111-2222-3333-444444444444", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := client.BeginCreateOrUpdate(ctx,
		"rg1",
		"ddb1",
		"privateEndpointConnectionName",
		armpowerplatform.PrivateEndpointConnection{
			Properties: &armpowerplatform.PrivateEndpointConnectionProperties{
				PrivateLinkServiceConnectionState: &armpowerplatform.PrivateLinkServiceConnectionState{
					Description: to.Ptr("Approved by johndoe@contoso.com"),
					Status:      to.Ptr(armpowerplatform.PrivateEndpointServiceConnectionStatusApproved),
				},
			},
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

Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Fpowerplatform%2Farmpowerplatform%2Fv0.1.0/sdk/resourcemanager/powerplatform/armpowerplatform/README.md) on how to add the SDK to your project and authenticate.
