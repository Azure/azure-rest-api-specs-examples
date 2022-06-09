```go
package armrecoveryservicesbackup_test

import (
	"context"
	"log"

	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/recoveryservices/armrecoveryservicesbackup"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/recoveryservicesbackup/resource-manager/Microsoft.RecoveryServices/stable/2021-12-01/examples/PrivateEndpointConnection/PutPrivateEndpointConnection.json
func ExamplePrivateEndpointConnectionClient_BeginPut() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
		return
	}
	ctx := context.Background()
	client, err := armrecoveryservicesbackup.NewPrivateEndpointConnectionClient("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
		return
	}
	poller, err := client.BeginPut(ctx,
		"<vault-name>",
		"<resource-group-name>",
		"<private-endpoint-connection-name>",
		armrecoveryservicesbackup.PrivateEndpointConnectionResource{
			Properties: &armrecoveryservicesbackup.PrivateEndpointConnection{
				PrivateEndpoint: &armrecoveryservicesbackup.PrivateEndpoint{
					ID: to.Ptr("<id>"),
				},
				PrivateLinkServiceConnectionState: &armrecoveryservicesbackup.PrivateLinkServiceConnectionState{
					Description: to.Ptr("<description>"),
					Status:      to.Ptr(armrecoveryservicesbackup.PrivateEndpointConnectionStatusApproved),
				},
				ProvisioningState: to.Ptr(armrecoveryservicesbackup.ProvisioningStateSucceeded),
			},
		},
		&armrecoveryservicesbackup.PrivateEndpointConnectionClientBeginPutOptions{ResumeToken: ""})
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

Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Frecoveryservices%2Farmrecoveryservicesbackup%2Fv0.5.0/sdk/resourcemanager/recoveryservices/armrecoveryservicesbackup/README.md) on how to add the SDK to your project and authenticate.
