Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Frecoveryservices%2Farmrecoveryservicesbackup%2Fv0.3.1/sdk/resourcemanager/recoveryservices/armrecoveryservicesbackup/README.md) on how to add the SDK to your project and authenticate.

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

// x-ms-original-file: specification/recoveryservicesbackup/resource-manager/Microsoft.RecoveryServices/stable/2021-12-01/examples/PrivateEndpointConnection/PutPrivateEndpointConnection.json
func ExamplePrivateEndpointConnectionClient_BeginPut() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armrecoveryservicesbackup.NewPrivateEndpointConnectionClient("<subscription-id>", cred, nil)
	poller, err := client.BeginPut(ctx,
		"<vault-name>",
		"<resource-group-name>",
		"<private-endpoint-connection-name>",
		armrecoveryservicesbackup.PrivateEndpointConnectionResource{
			Properties: &armrecoveryservicesbackup.PrivateEndpointConnection{
				PrivateEndpoint: &armrecoveryservicesbackup.PrivateEndpoint{
					ID: to.StringPtr("<id>"),
				},
				PrivateLinkServiceConnectionState: &armrecoveryservicesbackup.PrivateLinkServiceConnectionState{
					Description: to.StringPtr("<description>"),
					Status:      armrecoveryservicesbackup.PrivateEndpointConnectionStatus("Approved").ToPtr(),
				},
				ProvisioningState: armrecoveryservicesbackup.ProvisioningState("Succeeded").ToPtr(),
			},
		},
		nil)
	if err != nil {
		log.Fatal(err)
	}
	res, err := poller.PollUntilDone(ctx, 30*time.Second)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Response result: %#v\n", res.PrivateEndpointConnectionClientPutResult)
}
```
