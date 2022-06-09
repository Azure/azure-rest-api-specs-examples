```go
package armiotcentral_test

import (
	"context"
	"log"

	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/iotcentral/armiotcentral"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/iotcentral/resource-manager/Microsoft.IoTCentral/preview/2021-11-01-preview/examples/PrivateEndpointConnections_Update.json
func ExamplePrivateEndpointConnectionsClient_BeginCreate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
		return
	}
	ctx := context.Background()
	client, err := armiotcentral.NewPrivateEndpointConnectionsClient("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
		return
	}
	poller, err := client.BeginCreate(ctx,
		"<resource-group-name>",
		"<resource-name>",
		"<private-endpoint-connection-name>",
		armiotcentral.PrivateEndpointConnection{
			Properties: &armiotcentral.PrivateEndpointConnectionProperties{
				PrivateLinkServiceConnectionState: &armiotcentral.PrivateLinkServiceConnectionState{
					Description:     to.Ptr("<description>"),
					ActionsRequired: to.Ptr("<actions-required>"),
					Status:          to.Ptr(armiotcentral.PrivateEndpointServiceConnectionStatusApproved),
				},
			},
		},
		&armiotcentral.PrivateEndpointConnectionsClientBeginCreateOptions{ResumeToken: ""})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
		return
	}
	res, err := poller.PollUntilDone(ctx, 30*time.Second)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
		return
	}
	// TODO: use response item
	_ = res
}
```

Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Fiotcentral%2Farmiotcentral%2Fv0.4.0/sdk/resourcemanager/iotcentral/armiotcentral/README.md) on how to add the SDK to your project and authenticate.
