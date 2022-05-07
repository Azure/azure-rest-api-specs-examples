Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Faad%2Farmaad%2Fv0.4.0/sdk/resourcemanager/aad/armaad/README.md) on how to add the SDK to your project and authenticate.

```go
package armaad_test

import (
	"context"
	"log"

	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/aad/armaad"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/azureactivedirectory/resource-manager/Microsoft.Aadiam/stable/2020-03-01/examples/AzureADPrivateEndpointConnectionsCreate.json
func ExamplePrivateEndpointConnectionsClient_BeginCreate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
		return
	}
	ctx := context.Background()
	client, err := armaad.NewPrivateEndpointConnectionsClient("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
		return
	}
	poller, err := client.BeginCreate(ctx,
		"<resource-group-name>",
		"<policy-name>",
		"<private-endpoint-connection-name>",
		armaad.PrivateEndpointConnection{
			Properties: &armaad.PrivateEndpointConnectionProperties{
				PrivateEndpoint: &armaad.PrivateEndpoint{
					ID: to.Ptr("<id>"),
				},
				PrivateLinkServiceConnectionState: &armaad.PrivateLinkServiceConnectionState{
					Description:     to.Ptr("<description>"),
					ActionsRequired: to.Ptr("<actions-required>"),
					Status:          to.Ptr(armaad.PrivateEndpointServiceConnectionStatusApproved),
				},
			},
		},
		&armaad.PrivateEndpointConnectionsClientBeginCreateOptions{ResumeToken: ""})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
		return
	}
	_, err = poller.PollUntilDone(ctx, 30*time.Second)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
		return
	}
}
```
