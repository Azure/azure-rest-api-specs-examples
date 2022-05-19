Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Frecoveryservices%2Farmrecoveryservicesbackup%2Fv1.0.0/sdk/resourcemanager/recoveryservices/armrecoveryservicesbackup/README.md) on how to add the SDK to your project and authenticate.

```go
package armrecoveryservicesbackup_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/recoveryservices/armrecoveryservicesbackup"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/recoveryservicesbackup/resource-manager/Microsoft.RecoveryServices/stable/2022-02-01/examples/PrivateEndpointConnection/PutPrivateEndpointConnection.json
func ExamplePrivateEndpointConnectionClient_BeginPut() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armrecoveryservicesbackup.NewPrivateEndpointConnectionClient("04cf684a-d41f-4550-9f70-7708a3a2283b", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := client.BeginPut(ctx,
		"gaallavaultbvtd2msi",
		"gaallaRG",
		"gaallatestpe2.5704c932-249a-490b-a142-1396838cd3b",
		armrecoveryservicesbackup.PrivateEndpointConnectionResource{
			Properties: &armrecoveryservicesbackup.PrivateEndpointConnection{
				PrivateEndpoint: &armrecoveryservicesbackup.PrivateEndpoint{
					ID: to.Ptr("/subscriptions/04cf684a-d41f-4550-9f70-7708a3a2283b/resourceGroups/gaallaRG/providers/Microsoft.Network/privateEndpoints/gaallatestpe3"),
				},
				PrivateLinkServiceConnectionState: &armrecoveryservicesbackup.PrivateLinkServiceConnectionState{
					Description: to.Ptr("Approved by johndoe@company.com"),
					Status:      to.Ptr(armrecoveryservicesbackup.PrivateEndpointConnectionStatusApproved),
				},
				ProvisioningState: to.Ptr(armrecoveryservicesbackup.ProvisioningStateSucceeded),
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
